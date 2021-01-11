package aptible

import (
	"fmt"

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
	Services      []Service
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

	if response.Payload.Links.CurrentConfiguration != nil {
		configHref := response.Payload.Links.CurrentConfiguration.Href.String()
		config, err := c.GetConfigurationFromHref(configHref)
		if err != nil {
			return app, err
		}
		app.Env = config.Env
	}

	if response.Payload.Embedded.Services != nil {
		for _, s := range response.Payload.Embedded.Services {
			service := Service{
				ID:                     s.ID,
				ContainerCount:         s.ContainerCount,
				ContainerMemoryLimitMb: *s.ContainerMemoryLimitMb,
				ProcessType:            s.ProcessType,
				Command:                s.Command,
				ResourceType:           s.ResourceType,
				ResourceID:             app.ID,
				EnvironmentID:          app.EnvironmentID,
			}
			app.Services = append(app.Services, service)
		}
	}

	return app, err
}

func (c *Client) DeployApp(config map[string]interface{}, appID int64) error {
	requestType := "configure"
	if _, ok := config["APTIBLE_DOCKER_IMAGE"]; ok {
		requestType = "deploy"
	}
	appRequest := models.AppRequest22{Type: &requestType, Env: config}
	appParams := operations.NewPostAppsAppIDOperationsParams().WithAppID(appID).WithAppRequest(&appRequest)
	response, err := c.Client.Operations.PostAppsAppIDOperations(appParams, c.Token)
	if err != nil {
		return err
	}

	operationID := *response.Payload.ID
	_, err = c.WaitForOperation(operationID)

	return err
}

func (c *Client) DeleteApp(appID int64) (bool, error) {
	requestType := "deprovision"
	appRequest := models.AppRequest22{Type: &requestType}
	appParams := operations.NewPostAppsAppIDOperationsParams().WithAppID(appID).WithAppRequest(&appRequest)
	op, err := c.Client.Operations.PostAppsAppIDOperations(appParams, c.Token)
	if err != nil {
		return false, err
	}
	operationID := *op.Payload.ID
	return c.WaitForOperation(operationID)
}
