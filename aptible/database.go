package aptible

import (
	"errors"
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type Database struct {
	ID                int64
	DefaultConnection string
	ConnectionURLs    []string
	ContainerSize     int64
	DiskSize          int64
	ContainerProfile  string
	Deleted           bool
	Handle            string
	Type              string
	EnvironmentID     int64
	InitializeFromID  int64
	Service           Service
	DatabaseImage     DatabaseImage
}

// DBUpdates - struct to define what operations you contain your DB update to. Add values to this struct
// 			   to eventually pass it around for consumption by the go sdk
type DBUpdates struct {
	ContainerSize    int64
	DiskSize         int64
	ContainerProfile string
	Handle           string
	// SkipOperationUpdate - changing a DB can incur multiple API calls. To contain it to only update the handle
	// 						 set this to true then you only make one API call (to update database itself) without
	// 						 triggering an operation
	SkipOperationUpdate bool
}

type DBCreateAttrs struct {
	Handle           *string
	Type             string
	ContainerSize    int64
	DiskSize         int64
	ContainerProfile string
	DatabaseImageID  int64
}

func (c *Client) CreateDatabase(accountID int64, attrs DBCreateAttrs) (Database, error) {
	// creates API object
	request := models.AppRequest13{
		Handle: attrs.Handle,
		Type:   attrs.Type,
	}

	if attrs.DatabaseImageID != 0 {
		request.DatabaseImageID = &attrs.DatabaseImageID
	}

	params := operations.NewPostAccountsAccountIDDatabasesParams().WithAccountID(accountID).WithAppRequest(&request)
	resp, err := c.Client.Operations.PostAccountsAccountIDDatabases(params, c.Token)
	if err != nil {
		return Database{}, err
	}

	// provisions database
	requestType := "provision"
	provisionRequest := models.AppRequest24{
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

	// Setting the container profile on provision is not currently supported so
	// if a non-default container profile is requested, restart with the desired profile
	if attrs.ContainerProfile != "" && attrs.ContainerProfile != "m4" {
		requestType := "restart"
		request := models.AppRequest24{
			Type:            &requestType,
			InstanceProfile: attrs.ContainerProfile,
		}

		params := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseID).WithAppRequest(&request)
		op, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(params, c.Token)
		if err != nil {
			return Database{}, err
		}
		if op.Payload.ID != nil {
			operationID := *op.Payload.ID
			_, err = c.WaitForOperation(operationID)
			if err != nil {
				return Database{}, err
			}
		} else {
			return Database{}, fmt.Errorf("restart operation id is a nil pointer")
		}
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
		var e *operations.GetDatabasesIDDefault
		if errors.As(err, &e) {
			if e.Code() == 404 {
				err = nil
			}
			database.Deleted = true
			return database, err
		}
		return Database{}, err
	}

	defaultConnection := resp.Payload.ConnectionURL
	if defaultConnection == nil {
		return Database{}, fmt.Errorf("defaultConnection is a nil pointer")
	}
	database.DefaultConnection = *defaultConnection

	connectionUrls := resp.Payload.Embedded.DatabaseCredentials
	for _, u := range connectionUrls {
		if u == nil {
			continue
		}

		database.ConnectionURLs = append(database.ConnectionURLs, u.ConnectionURL)
	}

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
	database.ContainerProfile = service.ContainerProfile

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

	if resp.Payload.Links.DatabaseImage != nil {
		imageHref := resp.Payload.Links.DatabaseImage.Href.String()
		dbImage, err := c.GetImageFromHref(imageHref)
		if err != nil {
			return database, err
		}
		database.DatabaseImage = dbImage
	}

	return database, nil
}

func (c *Client) UpdateDatabase(databaseID int64, updates DBUpdates) error {
	requestType := "restart"
	request := models.AppRequest24{
		Type:            &requestType,
		InstanceProfile: updates.ContainerProfile,
	}

	if updates.ContainerSize >= 512 {
		request.ContainerSize = updates.ContainerSize
	}
	if updates.DiskSize >= 10 {
		request.DiskSize = updates.DiskSize
	}

	if !updates.SkipOperationUpdate {
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
	}

	if updates.Handle != "" {
		// need to update the database directly
		databaseUpdateRequest := models.AppRequest14{
			Handle: updates.Handle,
		}
		updateDatabaseParams := operations.NewPutDatabasesIDParams().WithID(databaseID).WithAppRequest(&databaseUpdateRequest)
		_, err := c.Client.Operations.PutDatabasesID(updateDatabaseParams, c.Token)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) DeleteDatabase(databaseID int64) error {
	requestType := "deprovision"
	request := models.AppRequest24{
		Type: &requestType,
	}
	deprovisionParams := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseID).WithAppRequest(&request)
	op, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(deprovisionParams, c.Token)
	if err != nil {
		return err
	}
	operationID := *op.Payload.ID
	_, err = c.WaitForOperation(operationID)
	return err
}

func (c *Client) GetDatabaseOperations(databaseID int64, page int64) (*models.InlineResponse20031, error) {
	params := operations.NewGetDatabasesDatabaseIDOperationsParams().WithDatabaseID(databaseID).WithPage(&page)
	resp, err := c.Client.Operations.GetDatabasesDatabaseIDOperations(params, c.Token)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
