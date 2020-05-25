package aptible

import (
	"fmt"
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
	Deleted       bool
}

type EndpointCreateAttrs struct {
	ResourceType       string
	ServiceProcessType string
	Type               *string
	Default            bool
	Internal           bool
	ContainerPort      int64
	IPWhitelist        []string
	Platform           string
}

// CreateEndpoint() creates Vhost API object + provision operation on the app.
func (c *Client) CreateEndpoint(service Service, attrs EndpointCreateAttrs) (Endpoint, error) {
	// Create Vhost API object
	request := models.AppRequest33{
		Type:        attrs.Type,
		Default:     attrs.Default,
		Internal:    attrs.Internal,
		IPWhitelist: attrs.IPWhitelist,
		Platform:    attrs.Platform,
	}
	if *attrs.Type != "tcp" {
		request.ContainerPort = attrs.ContainerPort
	}
	params := operations.NewPostServicesServiceIDVhostsParams().WithServiceID(service.ID).WithAppRequest(&request)
	resp, err := c.Client.Operations.PostServicesServiceIDVhosts(params, c.Token)
	if err != nil {
		return Endpoint{}, err
	}

	// Create "provision" operation
	payload := resp.Payload
	endpointID := *payload.ID
	requestType := "provision"
	operationRequest := models.AppRequest26{Type: &requestType}
	operationParams := operations.NewPostVhostsVhostIDOperationsParams().WithVhostID(endpointID).WithAppRequest(&operationRequest)
	operationResponse, err := c.Client.Operations.PostVhostsVhostIDOperations(operationParams, c.Token)
	if err != nil {
		return Endpoint{}, err
	}

	// Wait on provision operation to complete.
	if operationResponse.Payload.ID == nil {
		return Endpoint{}, fmt.Errorf("operation ID is a nil pointer")
	}
	operationID := *operationResponse.Payload.ID

	_, err = c.WaitForOperation(operationID)
	if err != nil {
		return Endpoint{}, err
	}

	endpoint, err := c.GetEndpoint(endpointID)

	return endpoint, err
}

// GetEndpoint() returns the response's payload, a bool saying whether or not the endpoint
// has been deprovisioned, and an error.
func (c *Client) GetEndpoint(endpointID int64) (Endpoint, error) {
	endpoint := Endpoint{
		Deleted: false,
	}
	params := operations.NewGetVhostsIDParams().WithID(endpointID)
	response, err := c.Client.Operations.GetVhostsID(params, c.Token)
	if err != nil {
		errStruct := err.(*operations.GetVhostsIDDefault)
		switch errStruct.Code() {
		case 404:
			// If deleted == true, then the endpoint needs to be removed from Terraform.
			endpoint.Deleted = true
			return endpoint, nil
		case 401:
			e := fmt.Errorf("make sure you have the correct auth token")
			return endpoint, e
		default:
			e := fmt.Errorf("there was an error when completing the request to get the app \n[ERROR] -%s", err)
			return endpoint, e
		}
	}

	if response.Payload.VirtualDomain == nil {
		return endpoint, fmt.Errorf("payload.VirtualDomain is a nil pointer")
	}

	endpoint.Hostname = *response.Payload.VirtualDomain
	if response.Payload.ContainerPort == nil {
		return endpoint, fmt.Errorf("payload.ContainerPort is a nil pointer")
	}
	endpoint.ContainerPort = *response.Payload.ContainerPort

	if response.Payload.ExternalHost == nil {
		return endpoint, fmt.Errorf("payload.ExternalHost is a nil pointer")
	}
	endpoint.Hostname = *response.Payload.ExternalHost


	if response.Payload.ID == nil {
		return endpoint, fmt.Errorf("payload.ID is a nil pointer")
	}
	endpoint.ID = *response.Payload.ID

	endpoint.IPWhitelist = response.Payload.IPWhitelist

	if response.Payload.Platform == nil {
		return endpoint, fmt.Errorf("payload.Platform is a nil pointer")
	}
	endpoint.Platform = *response.Payload.Platform

	return endpoint, nil
}

// UpdateEndpoint() takes in an endpointID and updates needed, and updates the endpoint.
func (c *Client) UpdateEndpoint(endpointID int64, up EndpointUpdates) error {
	appRequest := models.AppRequest34{
		ContainerPort: up.ContainerPort,
		IPWhitelist:   up.IPWhitelist,
		Platform:      up.Platform,
	}

	params := operations.NewPutVhostsIDParams().WithID(endpointID).WithAppRequest(&appRequest)
	_, err := c.Client.Operations.PutVhostsID(params, c.Token)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEndpoint() deletes the endpoint.
func (c *Client) DeleteEndpoint(endpointID int64) error {
	params := operations.NewDeleteVhostsIDParams().WithID(endpointID)
	_, err := c.Client.Operations.DeleteVhostsID(params, c.Token)
	return err
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
