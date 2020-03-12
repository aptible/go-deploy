package aptible

import (
	"os"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	deploy "github.com/aptible/go-deploy/client"
)

type Client struct {
	Client *deploy.DeployAPIV1
	Token  runtime.ClientAuthInfoWriter
}

// sets up client and gets auth token used for API requests
func SetUpClient() *Client {
	rt := httptransport.New(
		GetHost(),
		deploy.DefaultBasePath,
		deploy.DefaultSchemes)
	rt.Consumers["application/hal+json"] = runtime.JSONConsumer()
	rt.Producers["application/hal+json"] = runtime.JSONProducer()
	client := deploy.New(rt, strfmt.Default)

	token, _ := GetToken()
	bearerTokenAuth := httptransport.BearerToken(token)

	c := Client{}
	c.Client = client
	c.Token = bearerTokenAuth
	return &c
}

func GetHost() string {
	host, ok := os.LookupEnv("APTIBLE_API_ROOT_URL")
	if !ok {
		host = deploy.DefaultHost
	}

	if strings.HasPrefix(host, "https://") {
		host = host[8:]
	}

	return host
}
