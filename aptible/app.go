package aptible

import (
	"fmt"

	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
)

type App struct {
	ID      int64
	GitRepo string
}

func (c *Client) CreateApp(handle string, env_id int64) (App, error) {
	app := App{}
	appreq := models.AppRequest3{Handle: &handle}
	params := operations.NewPostAccountsAccountIDAppsParams().WithAccountID(env_id).WithAppRequest(&appreq)
	resp, err := c.Client.Operations.PostAccountsAccountIDApps(params, c.Token)
	if err != nil {
		return app, err
	}

	payload := resp.Payload
	if payload.ID == nil {
		return app, fmt.Errorf("App ID is a nil pointer")
	}
	app.ID = *payload.ID

	if payload.GitRepo == nil {
		return app, fmt.Errorf("App GitRepo is a nil pointer")
	}
	app.GitRepo = *payload.GitRepo

	return app, err
}

func (c *Client) DeployApp(app_id int64, config map[string]interface{}) error {
	req_type := "deploy"
	app_req := models.AppRequest21{Type: &req_type, Env: config, ContainerCount: 1, ContainerSize: 1024}
	app_params := operations.NewPostAppsAppIDOperationsParams().WithAppID(app_id).WithAppRequest(&app_req)
	app, err := c.Client.Operations.PostAppsAppIDOperations(app_params, c.Token)
	op_id := *app.Payload.ID
	_, err = c.WaitForOperation(op_id)
	return err
}

func (c *Client) GetApp(app_id int64) (bool, error) {
	params := operations.NewGetAppsIDParams().WithID(app_id)
	_, err := c.Client.Operations.GetAppsID(params, c.Token)

	if err != nil {
		err_struct := err.(*operations.GetAppsIDDefault)
		switch err_struct.Code() {
		case 404:
			// If deleted == true, then the app needs to be removed from Terraform.
			return true, nil
		case 401:
			e := fmt.Errorf("Make sure you have the correct auth token.")
			return false, e
		default:
			e := fmt.Errorf("There was an error when completing the request to get the app. \n[ERROR] -%s", err)
			return false, e
		}
	}
	return false, err
}

// Updates the `config` based on changes made in the config file
func (c *Client) UpdateApp(config map[string]interface{}, app_id int64) error {
	app_req := models.AppRequest21{}
	if _, ok := config["APTIBLE_DOCKER_IMAGE"]; ok {
		// Deploying app
		req_type := "deploy"
		app_req = models.AppRequest21{Type: &req_type, Env: config, ContainerCount: 1, ContainerSize: 1024}
	} else {
		// Configuring app
		req_type := "configure"
		app_req = models.AppRequest21{Type: &req_type, Env: config}
	}

	app_params := operations.NewPostAppsAppIDOperationsParams().WithAppID(app_id).WithAppRequest(&app_req)
	_, err := c.Client.Operations.PostAppsAppIDOperations(app_params, c.Token)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteApp(app_id int64) (bool, error) {
	req_type := "deprovision"
	app_req := models.AppRequest21{Type: &req_type}
	app_params := operations.NewPostAppsAppIDOperationsParams().WithAppID(app_id).WithAppRequest(&app_req)
	op, err := c.Client.Operations.PostAppsAppIDOperations(app_params, c.Token)
	if err != nil {
		return false, err
	}
	op_id := *op.Payload.ID
	return c.WaitForOperation(op_id)
}
