package aptible

import (
	"fmt"

	"github.com/reggregory/go-deploy/client/operations"
)

// Gets environment id associated with a given handle.
func (c *Client) GetEnvironmentIDFromHandle(handle string) (int64, error) {
	params := operations.NewGetAccountsParams()
	resp, err := c.Client.Operations.GetAccounts(params, c.Token)
	if err != nil {
		return 0, err
	}

	num_accts := *resp.Payload.TotalCount
	per_pg := *resp.Payload.PerPage
	page := *resp.Payload.CurrentPage

	for num_accts > 0 {
		accounts := resp.Payload.Embedded.Accounts
		for i := range accounts {
			if accounts[i].Handle == handle {
				a := accounts[i]
				return a.ID, nil
			}
		}
		if num_accts-per_pg > 0 {
			num_accts -= per_pg
			page += 1
		} else {
			return 0, fmt.Errorf("There are no accounts with handle: %s", handle)
		}
		params := operations.NewGetAccountsParams().WithPage(&page)
		resp, err = c.Client.Operations.GetAccounts(params, c.Token)
		if err != nil {
			return 0, err
		}
	}
	return 0, fmt.Errorf("There are no accounts with handle: %s", handle)
}

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
