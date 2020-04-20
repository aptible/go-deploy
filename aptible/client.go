package aptible

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	deploy "github.com/reggregory/go-deploy/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client struct {
	Client *deploy.DeployAPIV1
	Token  runtime.ClientAuthInfoWriter
}

// sets up client and gets auth token used for API requests
func SetUpClient() (*Client, error) {
	host, err := GetHost()
	if err != nil {
		return nil, err
	}

	rt := httptransport.New(
		host,
		deploy.DefaultBasePath,
		deploy.DefaultSchemes)
	rt.Consumers["application/hal+json"] = runtime.JSONConsumer()
	rt.Producers["application/hal+json"] = runtime.JSONProducer()
	client := deploy.New(rt, strfmt.Default)

	token, err := GetToken()
	if err != nil {
		return nil, err
	}
	bearerTokenAuth := httptransport.BearerToken(token)

	c := Client{}
	c.Client = client
	c.Token = bearerTokenAuth
	return &c, nil
}

// Gets host from env var, ensures it ends in ".com", or sets default.
func GetHost() (string, error) {
	host, ok := os.LookupEnv("APTIBLE_API_ROOT_URL")
	if !ok {
		return deploy.DefaultHost, nil
	}
	host = strings.Trim(host, " ")

	if strings.HasPrefix(host, "http://") {
		host = host[7:]
	} else if strings.HasPrefix(host, "https://") {
		host = host[8:]
	}

	if !validHost(host) {
		return "", fmt.Errorf("[ERROR] Host must be of the form xxx.xxx.com. Inputted host: %s", host)
	}

	return host, nil
}

func validHost(host string) bool {
	re, _ := regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	if re.MatchString(host) && strings.HasSuffix(host, ".com") {
		return true
	}
	return false
}
