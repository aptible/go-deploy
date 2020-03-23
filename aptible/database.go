package aptible

import (
	"fmt"
	"time"

	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
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

func (c *Client) CreateDatabase(env_id int64, attrs DBCreateAttrs) (*models.InlineResponse2014, error) {
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

	op_id := *op_resp.Payload.ID
	err = c.WaitForOperation(op_id)
	if err != nil {
		return nil, err
	}

	payload, _, err := c.GetDatabase(db_id)
	return payload, nil
}

func (c *Client) GetDatabase(db_id int64) (*models.InlineResponse2014, bool, error) {
	params := operations.NewGetDatabasesIDParams().WithID(db_id)
	resp, err := c.Client.Operations.GetDatabasesID(params, c.Token)
	if err != nil {
		switch err.(type) {
		case *operations.GetDatabasesIDDefault:
			err_struct := err.(*operations.GetDatabasesIDDefault)
			switch err_struct.Code() {
			case 404:
				// If deleted == true, then the database needs to be removed from Terraform.
				return nil, true, nil
			case 401:
				e := fmt.Errorf("Make sure you have the correct auth token.")
				return nil, false, e
			default:
				e := fmt.Errorf("There was an error when completing the request to get the database. \n[ERROR] -%s", err)
				return nil, false, e
			}
		default:
			e := fmt.Errorf("There was an unknown error. \n[ERROR] -%s", err)
			return nil, false, e
		}
	}
	return resp.Payload, false, nil
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
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteDatabase(db_id int64) error {
	params := operations.NewDeleteDatabasesIDParams().WithID(db_id)
	_, err := c.Client.Operations.DeleteDatabasesID(params, c.Token)
	if err != nil {
		return err
	}

	return nil
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

// Gets database with specific handle.
func (c *Client) GetDatabaseFromHandle(env_id int64, handle string) (*models.InlineResponse20014EmbeddedDatabases, error) {
	params := operations.NewGetAccountsAccountIDDatabasesParams().WithAccountID(env_id)
	resp, err := c.Client.Operations.GetAccountsAccountIDDatabases(params, c.Token)
	if err != nil {
		return nil, err
	}

	databases := resp.Payload.Embedded.Databases
	for i := range databases {
		if databases[i].Handle == handle {
			return databases[i], nil
		}
	}

	return nil, fmt.Errorf("There are no databases with handle: %s", handle)
}
