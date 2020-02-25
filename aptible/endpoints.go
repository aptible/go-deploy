package aptible

import (
	"fmt"
	"time"

	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
)

// CreateEndpoint() creates Vhost API object + provision operation on the app.
func (c *Client) CreateEndpoint(app_id int64) (*models.InlineResponse2019, error) {
	service_id, err := c.GetServiceID(app_id)
	if err != nil {
		return nil, err
	}

	// Create Vhost API object
	req_type := "http_proxy_protocol"
	app_req := models.AppRequest33{Type: &req_type, Default: true}
	params := operations.NewPostServicesServiceIDVhostsParams().WithServiceID(service_id).WithAppRequest(&app_req)
	resp, err := c.Client.Operations.PostServicesServiceIDVhosts(params, c.Token)
	if err != nil {
		return nil, err
	}

	payload := resp.Payload
	endpoint_id := *payload.ID

	// Create "provision" operation
	req_type = "provision"
	op_req := models.AppRequest26{Type: &req_type}
	op_params := operations.NewPostVhostsVhostIDOperationsParams().WithVhostID(endpoint_id).WithAppRequest(&op_req)
	_, err = c.Client.Operations.PostVhostsVhostIDOperations(op_params, c.Token)
	if err != nil {
		return nil, err
	}

	// Wait for endpoint to be provisioned.
	for *payload.Status != "provisioned" {
		time.Sleep(1 * time.Second)
		payload, _, err = c.GetEndpoint(endpoint_id)
	}

	return payload, nil
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
func (c *Client) UpdateEndpoint(endpoint_id int64, updates map[string]interface{}) error {
	app_req := models.AppRequest34{}

	if _, ok := updates["container_port"]; ok {
		app_req.ContainerPort = updates["container_port"].(int64)
	}

	if _, ok := updates["ip_filtering"]; ok {
		app_req.IPWhitelist = updates["ip_filtering"].([]string)
	}

	if _, ok := updates["platform"]; ok {
		app_req.Platform = updates["platform"].(string)
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
func (c *Client) GetServiceID(app_id int64) (int64, error) {
	params := operations.NewGetAppsAppIDServicesParams().WithAppID(app_id)
	resp, err := c.Client.Operations.GetAppsAppIDServices(params, c.Token)
	if err != nil {
		return 0, err
	}
	services := resp.Payload.Embedded.Services
	if len(services) <= 0 {
		return 0, fmt.Errorf("The app has no services.")
	}
	// for i := 0; i < len(services); i++ {
	// TODO: add logic for finding "right" service
	// }

	id := services[0].ID
	return id, err
}
