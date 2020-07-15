package aptible

import (
	"fmt"
	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type Database struct {
	ID               int64
	ConnectionURL    string
	ContainerSize    int64
	DiskSize         int64
	Deleted          bool
	Handle           string
	Type             string
	EnvironmentID    int64
	InitializeFromID int64
	Service		     Service
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

func (c *Client) CreateDatabase(accountID int64, attrs DBCreateAttrs) (Database, error) {
	// creates API object
	request := models.AppRequest12{
		Handle: attrs.Handle,
		Type:   attrs.Type,
	}
	params := operations.NewPostAccountsAccountIDDatabasesParams().WithAccountID(accountID).WithAppRequest(&request)
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
	databaseID := *resp.Payload.ID

	provisionParams := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseID).WithAppRequest(&provisionRequest)
	operationResponse, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(provisionParams, c.Token)
	if err != nil {
		return Database{}, err
	}

	// waits for provision operation to finish
	operationID := *operationResponse.Payload.ID
	_, err = c.WaitForOperation(operationID)
	if err != nil {
		return Database{}, err
	}

	// gets database
	return c.GetDatabase(databaseID)
}

func (c *Client) GetDatabase(databaseID int64) (Database, error) {
	database := Database{
		ID:      databaseID,
		Deleted: false,
	}

	params := operations.NewGetDatabasesIDParams().WithID(databaseID)
	resp, err := c.Client.Operations.GetDatabasesID(params, c.Token)
	if err != nil {
		switch err.(type) {
		case *operations.GetDatabasesIDDefault:
			if err.(*operations.GetDatabasesIDDefault).Code() == 404 {
				err = nil
			}
			database.Deleted = true
			return database, err
		default:
			return Database{}, err
		}
	}

	connectionUrl := resp.Payload.ConnectionURL
	if connectionUrl == nil {
		return Database{}, fmt.Errorf("connectionUrl is a nil pointer")
	}
	database.ConnectionURL = *connectionUrl

	databaseType := resp.Payload.Type
	if databaseType == nil {
		return Database{}, fmt.Errorf("databaseType is a nil pointer")
	}
	database.Type = *databaseType

	handle := resp.Payload.Handle
	if handle == nil {
		return Database{}, fmt.Errorf("handle is a nil pointer")
	}
	database.Handle = *handle

	// get updates to container size
	serviceHref := resp.Payload.Links.Service.Href.String()
	service, err := c.GetServiceFromHref(serviceHref)
	if err != nil {
		return Database{}, err
	}
	database.ContainerSize = service.ContainerMemoryLimitMb

	diskHref := resp.Payload.Links.Disk.Href.String()
	disk, err := c.GetDiskFromHref(diskHref)
	if err != nil {
		return Database{}, err
	}
	database.DiskSize = disk.Size

	envHref := resp.Payload.Links.Account.Href.String()
	envID, err := GetIDFromHref(envHref)
	if err != nil {
		return Database{}, err
	}
	database.EnvironmentID = envID

	if resp.Payload.Links.InitializeFrom != nil {
		initializeFromHref := resp.Payload.Links.InitializeFrom.Href.String()
		initializeFromID, err := GetIDFromHref(initializeFromHref)
		if err != nil {
			return database, err
		}
		database.InitializeFromID = initializeFromID
	}

	if resp.Payload.Links.Service != nil {
		serviceHref := resp.Payload.Links.Service.Href.String()
		service, err := c.GetServiceFromHref(serviceHref)
		if err != nil {
			return database, err
		}
		database.Service = service
	}

	return database, nil
}

func (c *Client) UpdateDatabase(databaseID int64, updates DBUpdates) error {
	requestType := "restart"
	request := models.AppRequest23{
		Type: &requestType,
	}

	if updates.ContainerSize >= 512 {
		request.ContainerSize = updates.ContainerSize
	}
	if updates.DiskSize >= 10 {
		request.DiskSize = updates.DiskSize
	}

	params := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseID).WithAppRequest(&request)
	op, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(params, c.Token)
	if err != nil {
		return err
	}
	if op.Payload.ID != nil {
		operationID := *op.Payload.ID
		_, err = c.WaitForOperation(operationID)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("id is a nil pointer")
	}

	return nil
}

func (c *Client) DeleteDatabase(databaseID int64) error {
	params := operations.NewDeleteDatabasesIDParams().WithID(databaseID)
	_, err := c.Client.Operations.DeleteDatabasesID(params, c.Token)
	return err
}

func (c *Client) GetDatabaseOperations(databaseID int64, page int64) (*models.InlineResponse20029, error) {
	params := operations.NewGetDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseID).WithPage(&page)
	resp, err := c.Client.Operations.GetDatabasesDatabaseIDOperations(params, c.Token)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
