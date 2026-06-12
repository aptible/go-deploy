package aptible

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	deploy "github.com/aptible/go-deploy/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type userAgentTransport struct {
	wrapped   http.RoundTripper
	userAgent string
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("User-Agent", t.userAgent)
	return t.wrapped.RoundTrip(req)
}

type Client struct {
	Client   *deploy.DeployAPIV1
	Token    runtime.ClientAuthInfoWriter
	RawToken string
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
	rt.Transport = &userAgentTransport{
		wrapped:   http.DefaultTransport,
		userAgent: "aptible/go-deploy/" + version(),
	}
	client := deploy.New(rt, strfmt.Default)

	token, err := GetToken()
	if err != nil {
		return nil, err
	}

	bearerTokenAuth := httptransport.BearerToken(token)

	c := Client{}
	c.Client = client
	c.Token = bearerTokenAuth
	c.RawToken = token
	return &c, nil
}

// Gets host from env var or sets default.
func GetHost() (string, error) {
	host, ok := os.LookupEnv("APTIBLE_API_ROOT_URL")
	if !ok {
		return deploy.DefaultHost, nil
	}
	host = strings.Trim(host, " ")

	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimPrefix(host, "https://")

	if !validHost(host) {
		return "", fmt.Errorf("[ERROR] Host must be of the form xxx.xxx.xxx. Inputted host: %s", host)
	}

	return host, nil
}

func validHost(host string) bool {
	re, _ := regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	return re.MatchString(host)
}
