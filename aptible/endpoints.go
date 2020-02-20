package aptible

import (
	"fmt"

	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
)

func (c *Client) CreateEndpoint(app_id int64) (int64, error) {
	service_id, err := c.GetServiceID(app_id)
	if err != nil {
		return 0, err
	}

	// Create Vhost API object
	req_type := "http_proxy_protocol"
	app_req := models.AppRequest33{Type: &req_type, Default: true}
	params := operations.NewPostServicesServiceIDVhostsParams().WithServiceID(service_id).WithAppRequest(&app_req)
	resp, err := c.Client.Operations.PostServicesServiceIDVhosts(params, c.Token)
	if err != nil {
		return 0, err
	}

	// Create "provision" operation
	vhost_id := *resp.Payload.ID
	req_type = "provision"
	op_req := models.AppRequest26{Type: &req_type}
	op_params := operations.NewPostVhostsVhostIDOperationsParams().WithVhostID(vhost_id).WithAppRequest(&op_req)
	_, err = c.Client.Operations.PostVhostsVhostIDOperations(op_params, c.Token)
	if err != nil {
		return 0, err
	}

	return vhost_id, nil
}

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
	// TODO: add logic for finding "right" service (whatever that means)
	// }

	id := services[0].ID
	return id, err
}
