package aptible

import (
	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type Environment struct {
	Handle *string
	ID     *int64
	Type   *string
}

type EnvironmentUpdates struct {
	Handle string
}

type EnvironmentCreateAttrs struct {
	Handle  string
	Type    string
	StackID int64
}

func (c *Client) CreateEnvironment(organizationID string, attrs EnvironmentCreateAttrs) (Environment, error) {
	params := operations.NewPostAccountsParams().WithAppRequest(&models.AppRequest{
		Handle:         &attrs.Handle,
		OrganizationID: &organizationID,
		Type:           &attrs.Type,
		StackID:        attrs.StackID,
	})
	environment, err := c.Client.Operations.PostAccounts(params, c.Token)
	if err != nil {
		return Environment{}, err
	}
	return Environment{
		Handle: environment.Payload.Handle,
		ID:     environment.Payload.ID,
		Type:   environment.Payload.Type,
	}, nil
}

func (c *Client) GetEnvironment(environmentID int64) (Environment, error) {
	params := operations.NewGetAccountsIDParams().WithID(environmentID)
	environment, err := c.Client.Operations.GetAccountsID(params, c.Token)
	if err != nil {
		return Environment{}, nil
	}
	return Environment{
		Handle: environment.Payload.Handle,
		ID:     environment.Payload.ID,
		Type:   environment.Payload.Type,
	}, nil
}

func (c *Client) UpdateEnvironment(environmentID int64, environmentUpdates EnvironmentUpdates) error {
	params := operations.NewPutAccountsIDParams().WithAppRequest(&models.AppRequest1{
		Handle: environmentUpdates.Handle,
	}).WithID(environmentID)
	_, err := c.Client.Operations.PutAccountsID(params, c.Token)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteEnvironment(environmentID int64) error {
	params := operations.NewDeleteAccountsIDParams().WithID(environmentID)
	_, err := c.Client.Operations.DeleteAccountsID(params, c.Token)
	if err != nil {
		return err
	}
	return nil
}