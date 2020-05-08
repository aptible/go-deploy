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

func (c *Client) CreateDatabase(env_id int64, attrs DBCreateAttrs) (Database, error) {
	// creates API object
	app_req := models.AppRequest12{
		Handle: attrs.Handle,
		Type:   attrs.Type,
	}
	params := operations.NewPostAccountsAccountIDDatabasesParams().WithAccountID(env_id).WithAppRequest(&app_req)
	resp, err := c.Client.Operations.PostAccountsAccountIDDatabases(params, c.Token)
	if err != nil {
		return Database{}, err
	}
	// provisions database
	req_type := "provision"
	prov_req := models.AppRequest23{
		Type:          &req_type,
		ContainerSize: attrs.ContainerSize,
		DiskSize:      attrs.DiskSize,
	}
	db_id := *resp.Payload.ID
	provision_params := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(db_id).WithAppRequest(&prov_req)
	op_resp, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(provision_params, c.Token)
	if err != nil {
		return Database{}, err
	}
	// waits for provision operation to finish
	op_id := *op_resp.Payload.ID
	_, err = c.WaitForOperation(op_id)
	if err != nil {
		return Database{}, err
	}
	// gets database
	db, _, err := c.GetDatabase(db_id)
	return db, nil
}

func (c *Client) GetDatabase(db_id int64) (Database, bool, error) {
	deleted := false
	db := Database{}
	db.ID = db_id

	// get database, error checking
	params := operations.NewGetDatabasesIDParams().WithID(db_id)
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
	c_url := resp.Payload.ConnectionURL
	if c_url == nil {
		return Database{}, deleted, fmt.Errorf("ConnectionURL is a nil pointer.")
	}
	db.ConnectionURL = *c_url

	// get updates to container size
	serv_href := resp.Payload.Links.Service.Href.String()
	c_size, err := c.GetContainerSize(serv_href)
	if err != nil {
		return Database{}, false, err
	}
	db.ContainerSize = c_size

	// get updates to disk size
	disk_href := resp.Payload.Links.Disk.Href.String()
	d_size, err := c.GetDiskSize(disk_href)
	if err != nil {
		return Database{}, false, err
	}
	db.DiskSize = d_size

	return db, false, nil
}

func (c *Client) UpdateDatabase(db_id int64, updates DBUpdates) error {
	req_type := "restart"
	app_req := models.AppRequest23{
		Type: &req_type,
	}

	if updates.ContainerSize >= 512 {
		app_req.ContainerSize = updates.ContainerSize
	}
	if updates.DiskSize >= 10 {
		app_req.DiskSize = updates.DiskSize
	}

	params := operations.NewPostDatabasesDatabaseIDOperationsParams().WithDatabaseID(db_id).WithAppRequest(&app_req)
	op, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(params, c.Token)
	if err != nil {
		return err
	}
	if op.Payload.ID != nil {
		op_id := *op.Payload.ID
		_, err = c.WaitForOperation(op_id)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("ID is a nil pointer.")
	}

	return nil
}

func (c *Client) DeleteDatabase(db_id int64) error {
	params := operations.NewDeleteDatabasesIDParams().WithID(db_id)
	_, err := c.Client.Operations.DeleteDatabasesID(params, c.Token)
	return err
}

// HELPERS //

// Gets operations of a database on a per page basis
func (c *Client) GetDatabaseOperations(db_id int64, page int64) (*models.InlineResponse20029, error) {
	params := operations.NewGetDatabasesDatabaseIDOperationsParams().WithDatabaseID(db_id).WithPage(&page)
	resp, err := c.Client.Operations.GetDatabasesDatabaseIDOperations(params, c.Token)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (c *Client) GetContainerSize(href string) (int64, error) {
	service_id, err := GetIDFromHref(href)
	if err != nil {
		return 0, err
	}
	serv_params := operations.NewGetServicesIDParams().WithID(service_id)
	serv_resp, err := c.Client.Operations.GetServicesID(serv_params, c.Token)
	if err != nil {
		return 0, err
	}
	container_ptr := serv_resp.Payload.ContainerMemoryLimitMb
	if container_ptr == nil {
		return 0, fmt.Errorf("Container size is a nil pointer.")
	}
	return *container_ptr, nil
}

func (c *Client) GetDiskSize(href string) (int64, error) {
	disk_id, err := GetIDFromHref(href)
	if err != nil {
		return 0, err
	}
	disk_params := operations.NewGetDisksIDParams().WithID(disk_id)
	disk_resp, err := c.Client.Operations.GetDisksID(disk_params, c.Token)
	if err != nil {
		return 0, err
	}
	disk_ptr := disk_resp.Payload.Size
	if disk_ptr == nil {
		return 0, fmt.Errorf("Disk size is a nil pointer.")
	}
	return *disk_ptr, nil
}
