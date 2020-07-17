package aptible

import (
	"errors"
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
)

// Gets environment id associated with a given handle.
func (c *Client) GetEnvironmentIDFromHandle(handle string) (int64, error) {
	params := operations.NewGetAccountsParams()
	response, err := c.Client.Operations.GetAccounts(params, c.Token)
	if err != nil {
		return 0, err
	}

	if response.Payload.TotalCount == nil {
		return 0, fmt.Errorf("TotalCount is a nil pointer")
	}
	numAccounts := *response.Payload.TotalCount

	if response.Payload.PerPage == nil {
		return 0, fmt.Errorf("PerPage is a nil pointer")
	}
	perPage := *response.Payload.PerPage

	if response.Payload.TotalCount == nil {
		return 0, fmt.Errorf("CurrentPage is a nil pointer")
	}
	page := *response.Payload.CurrentPage

	for numAccounts > 0 {
		accounts := response.Payload.Embedded.Accounts
		for i := range accounts {
			if accounts[i].Handle == handle {
				a := accounts[i]
				return a.ID, nil
			}
		}
		if numAccounts-perPage > 0 {
			numAccounts -= perPage
			page += 1
		} else {
			return 0, fmt.Errorf("there are no accounts with handle: %s", handle)
		}
		params := operations.NewGetAccountsParams().WithPage(&page)
		response, err = c.Client.Operations.GetAccounts(params, c.Token)
		if err != nil {
			return 0, err
		}
	}
	return 0, fmt.Errorf("there are no accounts with handle: %s", handle)
}

// Gets database id associated with a given handle.
func (c *Client) GetDatabaseIDFromHandle(accountID int64, handle string) (int64, bool, error) {
	deleted := false
	params := operations.NewGetAccountsAccountIDDatabasesParams().WithAccountID(accountID)
	response, err := c.Client.Operations.GetAccountsAccountIDDatabases(params, c.Token)
	if err != nil {
		var e *operations.GetAccountsAccountIDDatabasesDefault
		if errors.As(err, &e) {
			if e.Code() == 404 {
				deleted = true
			}
		}
		return 0, deleted, err
	}

	if response.Payload.TotalCount == nil {
		return 0, false, fmt.Errorf("TotalCount is a nil pointer")
	}
	numOps := *response.Payload.TotalCount

	if response.Payload.PerPage == nil {
		return 0, false, fmt.Errorf("PerPage is a nil pointer")
	}
	perPage := *response.Payload.PerPage

	if response.Payload.TotalCount == nil {
		return 0, false, fmt.Errorf("CurrentPage is a nil pointer")
	}
	page := *response.Payload.CurrentPage

	for numOps > 0 {
		databases := response.Payload.Embedded.Databases
		for i := range databases {
			if databases[i].Handle == handle {
				d := databases[i]
				return d.ID, deleted, nil
			}
		}
		if numOps-perPage > 0 {
			numOps -= perPage
			page += 1
		} else {
			return 0, deleted, fmt.Errorf("there are no databases with handle: %s", handle)
		}
		params := operations.NewGetAccountsAccountIDDatabasesParams().WithAccountID(accountID).WithPage(&page)
		response, err = c.Client.Operations.GetAccountsAccountIDDatabases(params, c.Token)
		if err != nil {
			var e *operations.GetAccountsAccountIDDatabasesDefault
			if errors.As(err, &e) {
				if e.Code() == 404 {
					deleted = true
				}
			}
			return 0, deleted, err
		}
	}
	return 0, deleted, fmt.Errorf("there are no databases with handle: %s", handle)
}
