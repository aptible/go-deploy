package aptible

import (
	"fmt"
	
	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
	"github.com/go-openapi/swag"
)

type Environment struct {
	Deleted bool
	Handle  string
	ID      int64
}

type EnvironmentUpdates struct {
	Handle string
}

type EnvironmentCreateAttrs struct {
	Handle string
}

func (c *Client) CreateEnvironment(organizationID string, stackID int64, attrs EnvironmentCreateAttrs) (Environment, error) {
	stack, err := c.GetStack(stackID)
	if err != nil {
		return Environment{}, err
	}
	environmentType := "production"
	if stack.isShared() {
		environmentType = "development"
	}
	params := operations.NewPostAccountsParams().WithAppRequest(&models.AppRequest{
		Handle:         &attrs.Handle,
		OrganizationID: &organizationID,
		StackID:        stack.ID,
		Type:           &environmentType,
	})
	environment, err := c.Client.Operations.PostAccounts(params, c.Token)
	if err != nil {
		return Environment{}, err
	}
	return Environment{
		Handle: *environment.Payload.Handle,
		ID:     *environment.Payload.ID,
	}, nil
}

func (c *Client) GetEnvironment(environmentID int64) (Environment, error) {
	environment := Environment{
		Deleted: false,
	}
	params := operations.NewGetAccountsIDParams().WithID(environmentID)
	environmentData, err := c.Client.Operations.GetAccountsID(params, c.Token)
	if err != nil {
		switch err.(*operations.GetAccountsIDDefault).Code() {
		case 404:
			environment.Deleted = true
			return environment, nil
		case 401:
			e := fmt.Errorf("make sure you have the correct auth token")
			return Environment{}, e
		default:
			e := fmt.Errorf("there was an error when completing the request to get the environment \n[ERROR] -%s", err)
			return Environment{}, e
		}
	}
	environment.Handle = swag.StringValue(environmentData.Payload.Handle)
	environment.ID = swag.Int64Value(environmentData.Payload.ID)
	return environment, nil
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
