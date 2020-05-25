package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type Database struct {
	ID            int64
	ConnectionURL string
	ContainerSize int64
	DiskSize      int64
}

type DBUpdates struct {
	ContainerSize int64
	DiskSize      int64
}

type DBCreateAttrs struct {
	Handle        *string
	Type          string
	ContainerSize int64
	DiskSize      int64
}

func (c *Client) CreateDatabase(accountId int64, attrs DBCreateAttrs) (Database, error) {
	// creates API object
	appRequest := models.AppRequest12{
		Handle: attrs.Handle,
		Type:   attrs.Type,
	}
	params := operations.NewPostAccountsAccountIDDatabasesParams().WithAccountID(accountId).WithAppRequest(&appRequest)
	resp, err := c.Client.Operations.PostAccountsAccountIDDatabases(params, c.Token)
	if err != nil {
		return Database{}, err
	}
	// provisions database
	requestType := "provision"
	provisionRequest := models.AppRequest23{
		Type:          &requestType,
		ContainerSize: attrs.ContainerSize,
		DiskSize:      attrs.DiskSize,
	}
	databaseId := *resp.Payload.ID
	provisionParams := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseId).WithAppRequest(&provisionRequest)
	operationResponse, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(provisionParams, c.Token)
	if err != nil {
		return Database{}, err
	}
	// waits for provision operation to finish
	operationId := *operationResponse.Payload.ID
	_, err = c.WaitForOperation(operationId)
	if err != nil {
		return Database{}, err
	}
	// gets database
	database, _, err := c.GetDatabase(databaseId)
	return database, nil
}

func (c *Client) GetDatabase(databaseId int64) (Database, bool, error) {
	deleted := false
	database := Database{}
	database.ID = databaseId

	// get database, error checking
	params := operations.NewGetDatabasesIDParams().WithID(databaseId)
	resp, err := c.Client.Operations.GetDatabasesID(params, c.Token)
	if err != nil {
		switch err.(type) {
		case *operations.GetDatabasesIDDefault:
			if err.(*operations.GetDatabasesIDDefault).Code() == 404 {
				deleted = true
				err = nil
			}
			return Database{}, deleted, err
		default:
			return Database{}, deleted, err
		}
	}

	// get ConnectionURL
	connectionUrl := resp.Payload.ConnectionURL
	if connectionUrl == nil {
		return Database{}, deleted, fmt.Errorf("connectionUrl is a nil pointer")
	}
	database.ConnectionURL = *connectionUrl

	// get updates to container size
	serviceHref := resp.Payload.Links.Service.Href.String()
	containerSize, err := c.GetContainerSize(serviceHref)
	if err != nil {
		return Database{}, false, err
	}
	database.ContainerSize = containerSize

	// get updates to disk size
	diskHref := resp.Payload.Links.Disk.Href.String()
	diskSize, err := c.GetDiskSize(diskHref)
	if err != nil {
		return Database{}, false, err
	}
	database.DiskSize = diskSize

	return database, false, nil
}

func (c *Client) UpdateDatabase(databaseId int64, updates DBUpdates) error {
	requestType := "restart"
	appRequest := models.AppRequest23{
		Type: &requestType,
	}

	if updates.ContainerSize >= 512 {
		appRequest.ContainerSize = updates.ContainerSize
	}
	if updates.DiskSize >= 10 {
		appRequest.DiskSize = updates.DiskSize
	}

	params := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseId).WithAppRequest(&appRequest)
	op, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(params, c.Token)
	if err != nil {
		return err
	}
	if op.Payload.ID != nil {
		operationId := *op.Payload.ID
		_, err = c.WaitForOperation(operationId)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("id is a nil pointer")
	}

	return nil
}

func (c *Client) DeleteDatabase(databaseId int64) error {
	params := operations.NewDeleteDatabasesIDParams().WithID(databaseId)
	_, err := c.Client.Operations.DeleteDatabasesID(params, c.Token)
	return err
}

// HELPERS //

// Gets operations of a database on a per page basis
func (c *Client) GetDatabaseOperations(databaseId int64, page int64) (*models.InlineResponse20029, error) {
	params := operations.NewGetDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseId).WithPage(&page)
	resp, err := c.Client.Operations.GetDatabasesDatabaseIDOperations(params, c.Token)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (c *Client) GetContainerSize(href string) (int64, error) {
	serviceId, err := GetIDFromHref(href)
	if err != nil {
		return 0, err
	}
	serviceParams := operations.NewGetServicesIDParams().WithID(serviceId)
	serviceResponse, err := c.Client.Operations.GetServicesID(serviceParams, c.Token)
	if err != nil {
		return 0, err
	}
	containerPtr := serviceResponse.Payload.ContainerMemoryLimitMb
	if containerPtr == nil {
		return 0, fmt.Errorf("container size is a nil pointer")
	}
	return *containerPtr, nil
}

func (c *Client) GetDiskSize(href string) (int64, error) {
	diskId, err := GetIDFromHref(href)
	if err != nil {
		return 0, err
	}
	diskParams := operations.NewGetDisksIDParams().WithID(diskId)
	diskResp, err := c.Client.Operations.GetDisksID(diskParams, c.Token)
	if err != nil {
		return 0, err
	}
	diskPtr := diskResp.Payload.Size
	if diskPtr == nil {
		return 0, fmt.Errorf("disk size is a nil pointer")
	}
	return *diskPtr, nil
}
