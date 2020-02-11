package helpers

import (
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
	if err != nil {
		return err
	}
	return err
}
