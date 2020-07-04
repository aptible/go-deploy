package aptible

import (
	"fmt"
	"log"
	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type App struct {
	ID            int64
	GitRepo       string
	Deleted       bool
	EnvironmentID int64
	Handle        string
	Env           interface{}
}

func (c *Client) CreateApp(handle string, accountID int64) (App, error) {
	app := App{}
	appRequest := models.AppRequest3{Handle: &handle}
	params := operations.NewPostAccountsAccountIDAppsParams().WithAccountID(accountID).WithAppRequest(&appRequest)
	response, err := c.Client.Operations.PostAccountsAccountIDApps(params, c.Token)
	if err != nil {
		return app, err
	}

	if response.Payload.ID == nil {
		return app, fmt.Errorf("app ID is a nil pointer")
	}
	app.ID = *response.Payload.ID

	if response.Payload.GitRepo == nil {
		return app, fmt.Errorf("app GitRepo is a nil pointer")
	}
	app.GitRepo = *response.Payload.GitRepo

	return app, err
}

func (c *Client) DeployApp(appID int64, config map[string]interface{}) error {
	requestType := "deploy"
	request := models.AppRequest21{Type: &requestType, Env: config, ContainerCount: 1, ContainerSize: 1024}
	params := operations.NewPostAppsAppIDOperationsParams().WithAppID(appID).WithAppRequest(&request)
	response, err := c.Client.Operations.PostAppsAppIDOperations(params, c.Token)

	operationID := *response.Payload.ID
	_, err = c.WaitForOperation(operationID)

	return err
}

func (c *Client) GetApp(appID int64) (App, error) {
	app := App{
		ID:      appID,
		Deleted: false,
	}

	params := operations.NewGetAppsIDParams().WithID(appID)
	response, err := c.Client.Operations.GetAppsID(params, c.Token)

	if err != nil {
		errStruct := err.(*operations.GetAppsIDDefault)
		switch errStruct.Code() {
		case 404:

			// If deleted == true, then the app needs to be removed from Terraform.
			app.Deleted = true
			return app, nil
		case 401:
			e := fmt.Errorf("make sure you have the correct auth token")
			return app, e
		default:
			e := fmt.Errorf("there was an error when completing the request to get the app \n[ERROR] -%s", err)
			return app, e
		}
	}

	if response.Payload.GitRepo == nil {
		return app, fmt.Errorf("app GitRepo is a nil pointer")
	}
	app.GitRepo = *response.Payload.GitRepo

	if response.Payload.Handle == nil {
		return app, fmt.Errorf("app Handle is a nil pointer")
	}
	app.Handle = *response.Payload.Handle

	envHref := response.Payload.Links.Account.Href.String()
	envID, err := GetIDFromHref(envHref)
	if err != nil {
		return app, err
	}
	app.EnvironmentID = envID

	log.Println(response.Payload.Links.CurrentConfiguration)

	if response.Payload.Links.CurrentConfiguration != nil {
		configHref := response.Payload.Links.CurrentConfiguration.Href.String()
		config, err := c.GetConfigurationFromHref(configHref)
		if err != nil {
			return app, err
		}
		app.Env = config.Env
	}

	return app, err
}

// Updates the `config` based on changes made in the config file
func (c *Client) UpdateApp(config map[string]interface{}, appID int64) error {
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

	appParams := operations.NewPostAppsAppIDOperationsParams().WithAppID(appID).WithAppRequest(&appRequest)
	_, err := c.Client.Operations.PostAppsAppIDOperations(appParams, c.Token)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteApp(appID int64) (bool, error) {
	requestType := "deprovision"
	appRequest := models.AppRequest21{Type: &requestType}
	appParams := operations.NewPostAppsAppIDOperationsParams().WithAppID(appID).WithAppRequest(&appRequest)
	op, err := c.Client.Operations.PostAppsAppIDOperations(appParams, c.Token)
	if err != nil {
		return false, err
	}
	operationID := *op.Payload.ID
	return c.WaitForOperation(operationID)
}
