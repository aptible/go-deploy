package helpers

import (
	// "fmt"

	"github.com/go-openapi/runtime"
	deploy "github.com/reggregory/go-deploy/client"
	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
)

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
