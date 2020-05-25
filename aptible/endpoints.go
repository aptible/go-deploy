package aptible

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type EndpointUpdates struct {
	ContainerPort int64
	IPWhitelist   []string
	Platform      string
}

type Endpoint struct {
	ID            int64
	Hostname      string
	ContainerPort int64
	IPWhitelist   []string
	Platform      string
}

type EndpointCreateAttrs struct {
	ResourceType  string
	Type          *string
	Default       bool
	Internal      bool
	ContainerPort int64
	IPWhitelist   []string
	Platform      string
}

// CreateEndpoint() creates Vhost API object + provision operation on the app.
func (c *Client) CreateEndpoint(resourceId int64, attrs EndpointCreateAttrs) (Endpoint, error) {
	serviceId, err := c.GetServiceID(resourceId, attrs.ResourceType)
	if err != nil {
		return Endpoint{}, err
	}

	// Create Vhost API object
	appRequest := models.AppRequest33{
		Type:        attrs.Type,
		Default:     attrs.Default,
		Internal:    attrs.Internal,
		IPWhitelist: attrs.IPWhitelist,
		Platform:    attrs.Platform,
	}
	if *attrs.Type != "tcp" {
		appRequest.ContainerPort = attrs.ContainerPort
	}
	params := operations.NewPostServicesServiceIDVhostsParams().WithServiceID(serviceId).WithAppRequest(&appRequest)
	resp, err := c.Client.Operations.PostServicesServiceIDVhosts(params, c.Token)
	if err != nil {
		return Endpoint{}, err
	}

	// Create "provision" operation
	payload := resp.Payload
	endpointId := *payload.ID
	requestType := "provision"
	operationRequest := models.AppRequest26{Type: &requestType}
	operationParams := operations.NewPostVhostsVhostIDOperationsParams().WithVhostID(endpointId).WithAppRequest(&operationRequest)
	operationResponse, err := c.Client.Operations.PostVhostsVhostIDOperations(operationParams, c.Token)
	if err != nil {
		return Endpoint{}, err
	}

	// Wait on provision operation to complete.
	if operationResponse.Payload.ID == nil {
		return Endpoint{}, fmt.Errorf("operation ID is a nil pointer")
	}
	operationId := *operationResponse.Payload.ID

	_, err = c.WaitForOperation(operationId)
	if err != nil {
		return Endpoint{}, err
	}

	endpoint, _, err := c.GetEndpoint(endpointId, attrs.ResourceType)

	return endpoint, err
}

// GetEndpoint() returns the response's payload, a bool saying whether or not the endpoint
// has been deprovisioned, and an error.
func (c *Client) GetEndpoint(endpointId int64, resourceType string) (Endpoint, bool, error) {
	ep := Endpoint{}
	params := operations.NewGetVhostsIDParams().WithID(endpointId)
	resp, err := c.Client.Operations.GetVhostsID(params, c.Token)
	if err != nil {
		errStruct := err.(*operations.GetVhostsIDDefault)
		switch errStruct.Code() {
		case 404:
			// If deleted == true, then the endpoint needs to be removed from Terraform.
			return Endpoint{}, true, nil
		case 401:
			e := fmt.Errorf("make sure you have the correct auth token")
			return Endpoint{}, false, e
		default:
			e := fmt.Errorf("there was an error when completing the request to get the app \n[ERROR] -%s", err)
			return Endpoint{}, false, e
		}
	}
	// Setting fields of Endpoint struct
	payload := resp.Payload
	// hostname, container port
	if resourceType == "app" {
		if payload.VirtualDomain == nil {
			return Endpoint{}, false, fmt.Errorf("payload.VirtualDomain is a nil pointer")
		}
		ep.Hostname = *payload.VirtualDomain
		if payload.ContainerPort == nil {
			return Endpoint{}, false, fmt.Errorf("payload.ContainerPort is a nil pointer")
		}
		ep.ContainerPort = *payload.ContainerPort
	} else {
		if payload.ExternalHost == nil {
			return Endpoint{}, false, fmt.Errorf("payload.ExternalHost is a nil pointer")
		}
		ep.Hostname = *payload.ExternalHost
	}
	// id
	if payload.ID == nil {
		return Endpoint{}, false, fmt.Errorf("payload.ID is a nil pointer")
	}
	ep.ID = *payload.ID
	// ip whitelist
	ep.IPWhitelist = payload.IPWhitelist
	// platform
	if payload.Platform == nil {
		return Endpoint{}, false, fmt.Errorf("payload.Platform is a nil pointer")
	}
	ep.Platform = *payload.Platform

	return ep, false, nil
}

// UpdateEndpoint() takes in an endpointId and updates needed, and updates the endpoint.
func (c *Client) UpdateEndpoint(endpointId int64, up EndpointUpdates) error {
	appRequest := models.AppRequest34{
		ContainerPort: up.ContainerPort,
		IPWhitelist:   up.IPWhitelist,
		Platform:      up.Platform,
	}

	params := operations.NewPutVhostsIDParams().WithID(endpointId).WithAppRequest(&appRequest)
	_, err := c.Client.Operations.PutVhostsID(params, c.Token)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEndpoint() deletes the endpoint.
func (c *Client) DeleteEndpoint(endpointId int64) error {
	params := operations.NewDeleteVhostsIDParams().WithID(endpointId)
	_, err := c.Client.Operations.DeleteVhostsID(params, c.Token)
	return err
}

// GetServiceID() Gets the service ID + acts as a helper for GetEndpoint().
func (c *Client) GetServiceID(resourceId int64, resourceType string) (int64, error) {
	if resourceType == "app" {
		params := operations.NewGetAppsAppIDServicesParams().WithAppID(resourceId)
		resp, err := c.Client.Operations.GetAppsAppIDServices(params, c.Token)
		if err != nil {
			return 0, err
		}
		services := resp.Payload.Embedded.Services
		if len(services) <= 0 {
			return 0, fmt.Errorf("the app must be deployed before creating an endpoint for it")
		}
		// TODO: add logic for finding "right" service
		return services[0].ID, nil
	} else if resourceType == "database" {
		params := operations.NewGetDatabasesIDParams().WithID(resourceId)
		resp, err := c.Client.Operations.GetDatabasesID(params, c.Token)
		if err != nil {
			return 0, err
		}

		serviceHref := resp.Payload.Links.Service.Href.String()
		serviceString := strings.Split(serviceHref, "/")[4]
		serviceId, _ := strconv.Atoi(serviceString)
		return int64(serviceId), nil
	}

	return 0, nil
}

func GetEndpointType(t string) (string, error) {
	switch t {
	case "HTTPS", "https":
		return "http_proxy_protocol", nil
	case "TCP", "tcp":
		return "tcp", nil
	case "TLS", "tls":
		return "tls", nil
	default:
		e := fmt.Errorf("invalid endpoint type, please use HTTPS, TLS, or TCP")
		return "", e
	}
}
