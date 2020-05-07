package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
)

// Gets database id associated with a given handle.
func (c *Client) GetDatabaseIDFromHandle(env_id int64, handle string) (int64, bool, error) {
	deleted := false
	params := operations.NewGetAccountsAccountIDDatabasesParams().WithAccountID(env_id)
	resp, err := c.Client.Operations.GetAccountsAccountIDDatabases(params, c.Token)
	if err != nil {
		switch err.(type) {
		case *operations.GetAccountsAccountIDDatabasesDefault:
			if err.(*operations.GetAccountsAccountIDDatabasesDefault).Code() == 404 {
				deleted = true
			}
			return 0, deleted, err
		default:
			return 0, deleted, err
		}
	}

	num_ops := *resp.Payload.TotalCount
	per_pg := *resp.Payload.PerPage
	page := *resp.Payload.CurrentPage

	for num_ops > 0 {
		databases := resp.Payload.Embedded.Databases
		for i := range databases {
			if databases[i].Handle == handle {
				d := databases[i]
				return d.ID, deleted, nil
			}
		}
		if num_ops-per_pg > 0 {
			num_ops -= per_pg
			page += 1
		} else {
			return 0, deleted, fmt.Errorf("There are no databases with handle: %s", handle)
		}
		params := operations.NewGetAccountsAccountIDDatabasesParams().WithAccountID(env_id).WithPage(&page)
		resp, err = c.Client.Operations.GetAccountsAccountIDDatabases(params, c.Token)
		if err != nil {
			switch err.(type) {
			case *operations.GetAccountsAccountIDDatabasesDefault:
				if err.(*operations.GetAccountsAccountIDDatabasesDefault).Code() == 404 {
					deleted = true
				}
				return 0, deleted, err
			default:
				return 0, deleted, err
			}
		}
	}
	return 0, deleted, fmt.Errorf("There are no databases with handle: %s", handle)
}
