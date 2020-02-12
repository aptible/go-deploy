package helpers

import (
	"fmt"

	"github.com/go-openapi/runtime"
	deploy "github.com/reggregory/go-deploy/client"
	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
)

func CreateApp(client *deploy.DeployAPIV1, token runtime.ClientAuthInfoWriter, handle string, account_id int64) (*models.InlineResponse2011, error) {
	appreq := models.AppRequest3{Handle: &handle}
	params := operations.NewPostAccountsAccountIDAppsParams().WithAccountID(account_id).WithAppRequest(&appreq)
	resp, err := client.Operations.PostAccountsAccountIDApps(params, token)
	if err != nil {
		return nil, err
	}
	return resp.Payload, err
}

func DeployApp(client *deploy.DeployAPIV1, token runtime.ClientAuthInfoWriter, app_id int64, env map[string]interface{}) error {
	req_type := "deploy"
	app_req := models.AppRequest21{Type: &req_type, Env: env, ContainerCount: 1, ContainerSize: 1024}
	app_params := operations.NewPostAppsAppIDOperationsParams().WithAppID(app_id).WithAppRequest(&app_req)
	_, err := client.Operations.PostAppsAppIDOperations(app_params, token)
	return err
}

func GetApp(client *deploy.DeployAPIV1, token runtime.ClientAuthInfoWriter, app_id int64) (bool, error) {
	params := operations.NewGetAppsIDParams().WithID(app_id)
	_, err := client.Operations.GetAppsID(params, token)

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

// Updates the `env` based on changes made in the config file
func UpdateEnv(env map[string]interface{}, client *deploy.DeployAPIV1, app_id int64, bearerTokenAuth runtime.ClientAuthInfoWriter) error {
	app_req := models.AppRequest21{}
	if _, ok := env["APTIBLE_DOCKER_IMAGE"]; ok {
		// Deploying app
		req_type := "deploy"
		app_req = models.AppRequest21{Type: &req_type, Env: env, ContainerCount: 1, ContainerSize: 1024}
	} else {
		// Configuring app
		req_type := "configure"
		app_req = models.AppRequest21{Type: &req_type, Env: env}
	}

	app_params := operations.NewPostAppsAppIDOperationsParams().WithAppID(app_id).WithAppRequest(&app_req)
	_, err := client.Operations.PostAppsAppIDOperations(app_params, bearerTokenAuth)
	if err != nil {
		return err
	}
	return nil
}

func DestroyApp(client *deploy.DeployAPIV1, token runtime.ClientAuthInfoWriter, app_id int64) error {
	req_type := "deprovision"
	app_req := models.AppRequest21{Type: &req_type}
	app_params := operations.NewPostAppsAppIDOperationsParams().WithAppID(app_id).WithAppRequest(&app_req)
	_, err := client.Operations.PostAppsAppIDOperations(app_params, token)
	return err
}
