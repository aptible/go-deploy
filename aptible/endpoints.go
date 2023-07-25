package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type EndpointUpdates struct {
	ContainerPort  int64
	ContainerPorts []int64
	IPWhitelist    []string
	Platform       string
}

type Endpoint struct {
	ID             int64
	ExternalHost   string
	ContainerPort  int64
	ContainerPorts []int64
	IPWhitelist    []string
	Platform       string
	Deleted        bool
	Acme           bool
	Default        bool
	UserDomain     string
	VirtualDomain  string
	Service        Service
	Type           string
	Internal       bool
	AcmeChallenges []AcmeChallenge
}

type AcmeChallenge struct {
	Method string
	From   string
	To     string
}

type EndpointCreateAttrs struct {
	ResourceType       string
	ServiceProcessType string
	Type               *string
	Default            bool
	Internal           bool
	ContainerPort      int64
	ContainerPorts     []int64
	IPWhitelist        []string
	Platform           string
	Acme               bool
	UserDomain         string
}

// CreateEndpoint() creates Vhost API object + provision operation on the app.
func (c *Client) CreateEndpoint(service Service, attrs EndpointCreateAttrs) (Endpoint, error) {
	// Create Vhost API object
	request := models.AppRequest34{
		Acme:        attrs.Acme,
		Type:        attrs.Type,
		Default:     attrs.Default,
		Internal:    attrs.Internal,
		IPWhitelist: attrs.IPWhitelist,
		Platform:    attrs.Platform,
	}

	if *attrs.Type == "tcp" {
		request.ContainerPort = attrs.ContainerPort
	}

	if *attrs.Type != "http_proxy_protocol" && *attrs.Type != "http" {
		request.ContainerPorts = attrs.ContainerPorts
	}

	if attrs.UserDomain != "" {
		request.UserDomain = attrs.UserDomain
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
	operationRequest := models.AppRequest27{Type: &requestType}
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

	return c.GetEndpoint(endpointID)
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

	if response.Payload.ID == nil {
		return endpoint, fmt.Errorf("payload.ID is a nil pointer")
	}
	endpoint.ID = *response.Payload.ID

	if response.Payload.VirtualDomain == nil {
		endpoint.VirtualDomain = ""
	} else {
		endpoint.VirtualDomain = *response.Payload.VirtualDomain
	}

	if response.Payload.UserDomain == nil {
		endpoint.UserDomain = ""
	} else {
		endpoint.UserDomain = *response.Payload.UserDomain
	}

	if response.Payload.ContainerPort == nil {
		endpoint.ContainerPort = 0
	} else {
		endpoint.ContainerPort = *response.Payload.ContainerPort
	}

	if response.Payload.ContainerPorts == nil {
		endpoint.ContainerPorts = []int64{}
	} else {
		endpoint.ContainerPorts = response.Payload.ContainerPorts
	}

	if response.Payload.ExternalHost == nil {
		return endpoint, fmt.Errorf("payload.ExternalHost is a nil pointer")
	}
	endpoint.ExternalHost = *response.Payload.ExternalHost

	if response.Payload.Internal == nil {
		return endpoint, fmt.Errorf("payload.Internal is a nil pointer")
	}
	endpoint.Internal = *response.Payload.Internal

	if response.Payload.Acme == nil {
		return endpoint, fmt.Errorf("payload.Acme is a nil pointer")
	}
	endpoint.Acme = *response.Payload.Acme

	if response.Payload.Default == nil {
		return endpoint, fmt.Errorf("payload.Default is a nil pointer")
	}
	endpoint.Default = *response.Payload.Default

	endpoint.IPWhitelist = response.Payload.IPWhitelist

	if response.Payload.Platform == nil {
		return endpoint, fmt.Errorf("payload.Platform is a nil pointer")
	}
	endpoint.Platform = *response.Payload.Platform

	if response.Payload.Type == nil {
		return endpoint, fmt.Errorf("payload.Type is a nil pointer")
	}
	endpointType, err := GetHumanReadableEndpointType(*response.Payload.Type)
	if err != nil {
		return endpoint, err
	}
	endpoint.Type = endpointType

	// Optional parameter we process if available
	if response.Payload.AcmeConfiguration != nil {
		var cs []AcmeChallenge
		for _, challenge := range (*response.Payload.AcmeConfiguration).Challenges {
			if challenge != nil && (*challenge).From != nil && (*challenge).To != nil {
				c := *challenge
				for _, to := range c.To {
					// These are deprecated challenge names so we aren't returning them
					if to.Legacy {
						continue
					}
					cs = append(cs, AcmeChallenge{Method: c.Method, From: c.From.Name, To: to.Name})
				}
			}
		}
		if len(cs) > 0 {
			endpoint.AcmeChallenges = cs
		}
	}

	serviceHref := response.Payload.Links.Service.Href.String()
	service, err := c.GetServiceFromHref(serviceHref)
	if err != nil {
		return endpoint, err
	}
	endpoint.Service = service

	return endpoint, nil
}

// UpdateEndpoint() takes in an endpointID and updates needed, and updates the endpoint.
func (c *Client) UpdateEndpoint(endpointID int64, up EndpointUpdates) error {
	appRequest := models.AppRequest35{
		ContainerPort:  up.ContainerPort,
		ContainerPorts: up.ContainerPorts,
		IPWhitelist:    up.IPWhitelist,
		Platform:       up.Platform,
	}

	params := operations.NewPutVhostsIDParams().WithID(endpointID).WithAppRequest(&appRequest)
	_, err := c.Client.Operations.PutVhostsID(params, c.Token)
	if err != nil {
		return err
	}

	requestType := "provision"
	operationRequest := models.AppRequest27{Type: &requestType}
	operationParams := operations.NewPostVhostsVhostIDOperationsParams().WithVhostID(endpointID).WithAppRequest(&operationRequest)
	operationResponse, err := c.Client.Operations.PostVhostsVhostIDOperations(operationParams, c.Token)
	if err != nil {
		return err
	}

	// Wait on provision operation to complete.
	if operationResponse.Payload.ID == nil {
		return fmt.Errorf("operation ID is a nil pointer")
	}
	operationID := *operationResponse.Payload.ID

	_, err = c.WaitForOperation(operationID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEndpoint() deletes the endpoint.
func (c *Client) DeleteEndpoint(endpointID int64) error {
	requestType := "deprovision"
	operationRequest := models.AppRequest27{Type: &requestType}
	operationParams := operations.NewPostVhostsVhostIDOperationsParams().WithVhostID(endpointID).WithAppRequest(&operationRequest)
	op, err := c.Client.Operations.PostVhostsVhostIDOperations(operationParams, c.Token)
	if err != nil {
		return err
	}

	operationID := *op.Payload.ID
	_, err = c.WaitForOperation(operationID)
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

func GetHumanReadableEndpointType(t string) (string, error) {
	switch t {
	case "http_proxy_protocol":
		return "https", nil
	case "tcp":
		return "tcp", nil
	case "tls":
		return "tls", nil
	default:
		e := fmt.Errorf("invalid endpoint type - %s", t)
		return "", e
	}
}
