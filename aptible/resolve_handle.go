package aptible

import (
	"fmt"

	"github.com/reggregory/go-deploy/client/operations"
	"github.com/reggregory/go-deploy/models"
)

// Gets database with specific handle.
func (c *Client) GetDatabaseFromHandle(env_id int64, handle string) (*models.InlineResponse20014EmbeddedDatabases, error) {
	params := operations.NewGetAccountsAccountIDDatabasesParams().WithAccountID(env_id)
	resp, err := c.Client.Operations.GetAccountsAccountIDDatabases(params, c.Token)
	if err != nil {
		return nil, err
	}

	num_ops := *resp.Payload.TotalCount
	per_pg := *resp.Payload.PerPage
	page := *resp.Payload.CurrentPage

	for num_ops > 0 {
		databases := resp.Payload.Embedded.Databases
		for i := range databases {
			if databases[i].Handle == handle {
				return databases[i], nil
			}
		}
		if num_ops-per_pg > 0 {
			num_ops -= per_pg
			page += 1
		} else {
			return nil, fmt.Errorf("There are no databases with handle: %s", handle)
		}
		params := operations.NewGetAccountsAccountIDDatabasesParams().WithAccountID(env_id).WithPage(&page)
		resp, err = c.Client.Operations.GetAccountsAccountIDDatabases(params, c.Token)
		if err != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("There are no databases with handle: %s", handle)
}
