package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
)

type Configuration struct {
	ID  int64
	Env interface{}
}

func (c *Client) GetConfiguration(configID int64) (Configuration, error) {
	config := Configuration{}

	params := operations.NewGetConfigurationsIDParams().WithID(configID)
	response, err := c.Client.Operations.GetConfigurationsID(params, c.Token)
	if err != nil {
		return config, err
	}

	if response.Payload.Env == nil {
		return config, fmt.Errorf("env is a nil pointer")
	}
	config.Env = response.Payload.Env

	return config, nil
}

func (c *Client) GetConfigurationFromHref(href string) (Configuration, error) {
	configID, err := GetIDFromHref(href)
	if err != nil {
		return Configuration{}, err
	}
	return c.GetConfiguration(configID)
}
