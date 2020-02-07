package helpers

import (
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	deploy "github.com/reggregory/go-deploy/client"
)

// sets up client and gets auth token used for API requests
func SetUpClient() (*deploy.DeployAPIV1, runtime.ClientAuthInfoWriter) {
	rt := httptransport.New(
		"api-rachel.aptible-sandbox.com",
		deploy.DefaultBasePath,
		deploy.DefaultSchemes)
	rt.Consumers["application/hal+json"] = runtime.JSONConsumer()
	rt.Producers["application/hal+json"] = runtime.JSONProducer()
	client := deploy.New(rt, strfmt.Default)

	token, _ := GetToken()
	bearerTokenAuth := httptransport.BearerToken(token)

	return client, bearerTokenAuth
}
