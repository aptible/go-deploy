package aptible

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
)

type Updates struct {
	ContainerPort int64
	IPWhitelist   []string
	Platform      string
}

type CreateAttrs struct {
	ResourceType  string
	Type          *string
	Default       bool
	Internal      bool
	ContainerPort int64
	IPWhitelist   []string
	Platform      string
}

// CreateEndpoint() creates Vhost API object + provision operation on the app.
func (c *Client) CreateEndpoint(resource_id int64, attrs CreateAttrs) (*models.InlineResponse2019, error) {
	service_id, err := c.GetServiceID(resource_id, attrs.ResourceType)
	if err != nil {
		return nil, err
	}

	// Create Vhost API object
	app_req := models.AppRequest33{
		Type:        attrs.Type,
		Internal:    attrs.Internal,
		IPWhitelist: attrs.IPWhitelist,
		Platform:    attrs.Platform,
	}
	if *attrs.Type != "tcp" {
		app_req.ContainerPort = attrs.ContainerPort
	}
	params := operations.NewPostServicesServiceIDVhostsParams().WithServiceID(service_id).WithAppRequest(&app_req)
	resp, err := c.Client.Operations.PostServicesServiceIDVhosts(params, c.Token)
	if err != nil {
		return nil, err
	}

	// Create "provision" operation
	payload := resp.Payload
	endpoint_id := *payload.ID
	req_type := "provision"
	op_req := models.AppRequest26{Type: &req_type}
	op_params := operations.NewPostVhostsVhostIDOperationsParams().WithVhostID(endpoint_id).WithAppRequest(&op_req)
	op_resp, err := c.Client.Operations.PostVhostsVhostIDOperations(op_params, c.Token)
	if err != nil {
		return nil, err
	}

	// Wait on provision operation to complete.
	op_id := *op_resp.Payload.ID
	err = c.WaitForOperation(op_id)
	if err != nil {
		return nil, err
	}
	payload, _, err = c.GetEndpoint(endpoint_id)
	return payload, err
}

// GetEndpoint() returns the response's payload, a bool saying whether or not the endpoint
// has been deprovisioned, and an error.
func (c *Client) GetEndpoint(endpoint_id int64) (*models.InlineResponse2019, bool, error) {
	params := operations.NewGetVhostsIDParams().WithID(endpoint_id)
	resp, err := c.Client.Operations.GetVhostsID(params, c.Token)
	if err != nil {
		err_struct := err.(*operations.GetVhostsIDDefault)
		switch err_struct.Code() {
		case 404:
			// If deleted == true, then the endpoint needs to be removed from Terraform.
			return nil, true, nil
		case 401:
			e := fmt.Errorf("Make sure you have the correct auth token.")
			return nil, false, e
		default:
			e := fmt.Errorf("There was an error when completing the request to get the app. \n[ERROR] -%s", err)
			return nil, false, e
		}
	}
	return resp.Payload, false, nil
}

// UpdateEndpoint() takes in an endpoint_id and updates needed, and updates the endpoint.
func (c *Client) UpdateEndpoint(endpoint_id int64, up Updates) error {
	app_req := models.AppRequest34{
		ContainerPort: up.ContainerPort,
		IPWhitelist:   up.IPWhitelist,
		Platform:      up.Platform,
	}

	params := operations.NewPutVhostsIDParams().WithID(endpoint_id).WithAppRequest(&app_req)
	_, err := c.Client.Operations.PutVhostsID(params, c.Token)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEndpoint() deletes the endpoint.
func (c *Client) DeleteEndpoint(endpoint_id int64) error {
	params := operations.NewDeleteVhostsIDParams().WithID(endpoint_id)
	_, err := c.Client.Operations.DeleteVhostsID(params, c.Token)
	if err != nil {
		return err
	}
	return nil
}

// GetServiceID() Gets the service ID + acts as a helper for GetEndpoint().
func (c *Client) GetServiceID(resource_id int64, resource_type string) (int64, error) {
	if resource_type == "app" {
		params := operations.NewGetAppsAppIDServicesParams().WithAppID(resource_id)
		resp, err := c.Client.Operations.GetAppsAppIDServices(params, c.Token)
		if err != nil {
			return 0, err
		}
		services := resp.Payload.Embedded.Services
		if len(services) <= 0 {
			return 0, fmt.Errorf("The app has no services.")
		}
		// TODO: add logic for finding "right" service
		return services[0].ID, nil
	} else if resource_type == "database" {
		params := operations.NewGetDatabasesIDParams().WithID(resource_id)
		resp, err := c.Client.Operations.GetDatabasesID(params, c.Token)
		if err != nil {
			return 0, err
		}

		serv_href := resp.Payload.Links.Service.Href.String()
		serv_str := strings.Split(serv_href, "/")[4]
		service_id, _ := strconv.Atoi(serv_str)
		return int64(service_id), nil
	}

	return 0, nil
}

func GetEndpointType(t string) (string, error) {
	switch t {
	case "HTTPS":
		return "http_proxy_protocol", nil
	default:
		e := fmt.Errorf("Invalid endpoint type. The only valid types are HTTPS, TLS, and TCP.")
		return "", e
	}
}

func MakeStringSlice(if_slice []interface{}) ([]string, error) {
	str_slice := make([]string, len(if_slice))
	for i := 0; i < len(if_slice); i++ {
		str_slice[i] = (if_slice[i].(string))
	}
	return str_slice, nil
}
