package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type App struct {
	ID      int64
	GitRepo string
}

func (c *Client) CreateApp(handle string, accountId int64) (App, error) {
	app := App{}
	appRequest := models.AppRequest3{Handle: &handle}
	params := operations.NewPostAccountsAccountIDAppsParams().WithAccountID(accountId).WithAppRequest(&appRequest)
	response, err := c.Client.Operations.PostAccountsAccountIDApps(params, c.Token)
	if err != nil {
		return app, err
	}

	payload := response.Payload
	if payload.ID == nil {
		return app, fmt.Errorf("app ID is a nil pointer")
	}
	app.ID = *payload.ID

	if payload.GitRepo == nil {
		return app, fmt.Errorf("app GitRepo is a nil pointer")
	}
	app.GitRepo = *payload.GitRepo

	return app, err
}

func (c *Client) DeployApp(appId int64, config map[string]interface{}) error {
	requestType := "deploy"
	appRequest := models.AppRequest21{Type: &requestType, Env: config, ContainerCount: 1, ContainerSize: 1024}
	appParams := operations.NewPostAppsAppIDOperationsParams().WithAppID(appId).WithAppRequest(&appRequest)
	app, err := c.Client.Operations.PostAppsAppIDOperations(appParams, c.Token)
	operationId := *app.Payload.ID
	_, err = c.WaitForOperation(operationId)
	return err
}

func (c *Client) GetApp(appId int64) (bool, error) {
	params := operations.NewGetAppsIDParams().WithID(appId)
	_, err := c.Client.Operations.GetAppsID(params, c.Token)

	if err != nil {
		errStruct := err.(*operations.GetAppsIDDefault)
		switch errStruct.Code() {
		case 404:
			// If deleted == true, then the app needs to be removed from Terraform.
			return true, nil
		case 401:
			e := fmt.Errorf("make sure you have the correct auth token")
			return false, e
		default:
			e := fmt.Errorf("there was an error when completing the request to get the app \n[ERROR] -%s", err)
			return false, e
		}
	}
	return false, err
}

// Updates the `config` based on changes made in the config file
func (c *Client) UpdateApp(config map[string]interface{}, appId int64) error {
	appRequest := models.AppRequest21{}
	if _, ok := config["APTIBLE_DOCKER_IMAGE"]; ok {
		// Deploying app
		requestType := "deploy"
		appRequest = models.AppRequest21{Type: &requestType, Env: config, ContainerCount: 1, ContainerSize: 1024}
	} else {
		// Configuring app
		requestType := "configure"
		appRequest = models.AppRequest21{Type: &requestType, Env: config}
	}

	appParams := operations.NewPostAppsAppIDOperationsParams().WithAppID(appId).WithAppRequest(&appRequest)
	_, err := c.Client.Operations.PostAppsAppIDOperations(appParams, c.Token)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteApp(appId int64) (bool, error) {
	requestType := "deprovision"
	appRequest := models.AppRequest21{Type: &requestType}
	appParams := operations.NewPostAppsAppIDOperationsParams().WithAppID(appId).WithAppRequest(&appRequest)
	op, err := c.Client.Operations.PostAppsAppIDOperations(appParams, c.Token)
	if err != nil {
		return false, err
	}
	operationId := *op.Payload.ID
	return c.WaitForOperation(operationId)
}
