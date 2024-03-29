package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type Service struct {
	ID                     int64
	ContainerCount         int64
	ContainerMemoryLimitMb int64
	ProcessType            string
	Command                string
	ResourceType           string
	ResourceID             int64
	EnvironmentID          int64
	ContainerProfile       string
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

	if response.Payload.InstanceClass == nil {
		return service, fmt.Errorf("instance class is a nil pointer")
	}
	service.ContainerProfile = *response.Payload.InstanceClass

	if response.Payload.Command == nil {
		service.Command = ""
	} else {
		service.Command = *response.Payload.Command
	}

	if response.Payload.Links.Database == nil && response.Payload.Links.App == nil {
		return service, fmt.Errorf("something is wrong: both app and database are nil for service")
	}

	if response.Payload.Links.Database != nil {
		service.ResourceType = "database"
		resourceHref := response.Payload.Links.Database.Href.String()
		resourceID, err := GetIDFromHref(resourceHref)
		if err != nil {
			return service, err
		}
		service.ResourceID = resourceID
	} else {
		service.ResourceType = "app"
		resourceHref := response.Payload.Links.App.Href.String()
		resourceID, err := GetIDFromHref(resourceHref)
		if err != nil {
			return service, err
		}
		service.ResourceID = resourceID
	}

	envHref := response.Payload.Links.Account.Href.String()
	envID, err := GetIDFromHref(envHref)
	if err != nil {
		return service, err
	}
	service.EnvironmentID = envID

	return service, nil
}

func (c *Client) ScaleService(serviceID int64, containerCount int64, memoryLimit int64, containerProfile string) error {
	requestType := "scale"
	request := models.AppRequest26{
		Type:            &requestType,
		ContainerSize:   memoryLimit,
		ContainerCount:  containerCount,
		InstanceProfile: containerProfile,
	}
	appParams := operations.NewPostServicesServiceIDOperationsParams().WithServiceID(serviceID).WithAppRequest(&request)
	response, err := c.Client.Operations.PostServicesServiceIDOperations(appParams, c.Token)
	if err != nil {
		return err
	}

	operationID := *response.Payload.ID
	_, err = c.WaitForOperation(operationID)

	return err
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
			s := Service{
				ID:               service.ID,
				ContainerCount:   service.ContainerCount,
				ProcessType:      service.ProcessType,
				Command:          service.Command,
				ContainerProfile: service.InstanceClass,
			}
			if service.ContainerMemoryLimitMb != nil {
				s.ContainerMemoryLimitMb = *service.ContainerMemoryLimitMb
			}
			return s, nil
		}
	}

	return Service{}, fmt.Errorf("service for app %v with process type %s cannot be found", appID, serviceName)
}
