package aptible

import (
	"fmt"
	"time"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

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

type DBCreated struct {
	ID            int64
	ConnectionURL string
	ContainerSize string
	DiskSize      string
}

func (c *Client) CreateDatabase(env_id int64, attrs DBCreateAttrs) (*models.InlineResponse20014EmbeddedDatabases, error) {
	// creates API object
	app_req := models.AppRequest12{
		Handle: attrs.Handle,
		Type:   attrs.Type,
	}
	params := operations.NewPostAccountsAccountIDDatabasesParams().WithAccountID(env_id).WithAppRequest(&app_req)
	resp, err := c.Client.Operations.PostAccountsAccountIDDatabases(params, c.Token)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	// waits for provision operation to finish
	op_id := *op_resp.Payload.ID
	err = c.WaitForOperation(op_id)
	if err != nil {
		return nil, err
	}
	// gets database
	payload, _, err := c.GetDatabaseFromHandle(env_id, *attrs.Handle)
	return payload, nil
}

func (c *Client) GetDatabase(db_id int64) (DBUpdates, bool, error) {
	deleted := false
	updates := DBUpdates{}
	page := int64(1)
	payload, err := c.GetDatabaseOperations(db_id, page)
	if err != nil {
		switch err.(type) {
		case *operations.GetDatabasesDatabaseIDOperationsDefault:
			if err.(*operations.GetDatabasesDatabaseIDOperationsDefault).Code() == 404 {
				deleted = true
			}
			return updates, deleted, err
		default:
			return updates, deleted, err
		}
	}
	// get updates to container size + disk size from operations
	num_ops := *payload.TotalCount
	per_pg := *payload.PerPage
	for num_ops > 0 {
		ops := payload.Embedded.Operations
		GetUpdatesFromOperations(ops, &updates)
		// if more pages left
		if num_ops-per_pg > 0 {
			num_ops -= per_pg
			page += 1
			// get new page of ops
			payload, err = c.GetDatabaseOperations(db_id, page)
			if err != nil {
				switch err.(type) {
				case *operations.GetDatabasesDatabaseIDOperationsDefault:
					if err.(*operations.GetDatabasesDatabaseIDOperationsDefault).Code() == 404 {
						deleted = true
					}
					return updates, deleted, err
				default:
					return updates, deleted, err
				}
			}
		} else {
			return updates, deleted, nil
		}
	}
	return updates, deleted, fmt.Errorf("Unknown error occurred.")
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
	_, err := c.Client.Operations.PostDatabasesDatabaseIDOperations(params, c.Token)
	return err
}

func (c *Client) DeleteDatabase(db_id int64) error {
	params := operations.NewDeleteDatabasesIDParams().WithID(db_id)
	_, err := c.Client.Operations.DeleteDatabasesID(params, c.Token)
	return err
}

// HELPERS //

// Waits for operation to succeed.
func (c *Client) WaitForOperation(op_id int64) error {

	params := operations.NewGetOperationsIDParams().WithID(op_id)
	op, err := c.Client.Operations.GetOperationsID(params, c.Token)
	if err != nil {
		return err
	}
	status := *op.Payload.Status

	for status != "succeeded" {
		time.Sleep(2 * time.Second)
		op, err = c.Client.Operations.GetOperationsID(params, c.Token)
		if err != nil {
			return err
		}
		status = *op.Payload.Status
		fmt.Println("Still creating...")
	}
	fmt.Println("Done!")

	return nil
}

// Gets operations of a database on a per page basis
func (c *Client) GetDatabaseOperations(db_id int64, page int64) (*models.InlineResponse20029, error) {
	params := operations.NewGetDatabasesDatabaseIDOperationsParams().WithDatabaseID(db_id).WithPage(&page)
	resp, err := c.Client.Operations.GetDatabasesDatabaseIDOperations(params, c.Token)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// Gets updates to container and disk size
func GetUpdatesFromOperations(ops []*models.InlineResponse2003EmbeddedEmbeddedLastOperation, updates *DBUpdates) {
	for i := range ops {
		if ops[i].Type == "provision" || ops[i].Type == "restart" {
			if ops[i].ContainerSize != nil && updates.ContainerSize == 0 {
				updates.ContainerSize = *ops[i].ContainerSize
			}
			if ops[i].DiskSize != 0 && updates.DiskSize == 0 {
				updates.DiskSize = ops[i].DiskSize
			}
		}
	}
}
