package aptible

import (
	"fmt"
	"github.com/aptible/go-deploy/client/operations"
)

type Service struct {
	ID                     int64
	ContainerCount         int64
	ContainerMemoryLimitMb int64
	ProcessType            string
	Command                string
	ResourceType           string
}

func (c *Client) GetService(serviceID int64) (Service, error) {
	service := Service{}

	params := operations.NewGetServicesIDParams().WithID(serviceID)
	response, err := c.Client.Operations.GetServicesID(params, c.Token)
	if err != nil {
		return service, err
	}

	if response.Payload.ID == nil {
		return service, fmt.Errorf("disk size is a nil pointer")
	}
	service.ID = *response.Payload.ID

	if response.Payload.ContainerCount == nil {
		return service, fmt.Errorf("container count is a nil pointer")
	}
	service.ContainerCount = *response.Payload.ContainerCount

	if response.Payload.ContainerMemoryLimitMb == nil {
		return service, fmt.Errorf("container memory limit mb is a nil pointer")
	}
	service.ContainerMemoryLimitMb = *response.Payload.ContainerMemoryLimitMb

	if response.Payload.ProcessType == nil {
		return service, fmt.Errorf("process type is a nil pointer")
	}
	service.ProcessType = *response.Payload.ProcessType

	if response.Payload.Command == nil {
		service.Command = ""
	} else {
		service.Command = *response.Payload.Command
	}

	if response.Payload.Links.App != nil {
		service.ResourceType = "app"
	} else {
		service.ResourceType = "database"
	}

	return service, nil
}

func (c *Client) GetServiceFromHref(href string) (Service, error) {
	serviceID, err := GetIDFromHref(href)
	if err != nil {
		return Service{}, err
	}
	return c.GetService(serviceID)
}

func (c *Client) GetServiceForAppByName(appID int64, serviceName string) (Service, error) {

	params := operations.NewGetAppsAppIDServicesParams().WithAppID(appID)
	response, err := c.Client.Operations.GetAppsAppIDServices(params, c.Token)
	if err != nil {
		return Service{}, err
	}

	services := response.Payload.Embedded.Services
	if len(services) <= 0 {
		return Service{}, fmt.Errorf("the app must be deployed before creating an endpoint for it")
	}

	for _, service := range services {
		if service.ProcessType == serviceName {
			return Service{
				ID:                     service.ID,
				ContainerCount:         service.ContainerCount,
				ContainerMemoryLimitMb: *service.ContainerMemoryLimitMb,
				ProcessType:            service.ProcessType,
				Command:                service.Command,
			}, nil
		}
	}

	return Service{}, fmt.Errorf("service for app %v with process type %s cannot be found", appID, serviceName)
}
