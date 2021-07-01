package aptible

import (
	"fmt"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
)

type LogDrain struct {
	Deleted   bool
	ID        int64
	Handle    string
	DrainType string
}

type LogDrainCreateAttrs struct {
	Database               strfmt.URI
	DatabaseID             int64
	DrainApps              bool
	DrainDatabases         bool
	DrainEphemeralSessions bool
	DrainHost              strfmt.URI
	DrainPassword          string
	DrainPort              int64
	DrainProxies           bool
	DrainType              *string
	Handle                 *string
	LoggingToken           string
	URL                    strfmt.URI
}

func (c *Client) CreateLogDrain(handle *string, accountID int64, attrs *LogDrainCreateAttrs) (*LogDrain, error) {
	request := &models.AppRequest16{
		Database:               attrs.Database,
		DatabaseID:             attrs.DatabaseID,
		DrainApps:              attrs.DrainApps,
		DrainDatabases:         attrs.DrainDatabases,
		DrainEphemeralSessions: attrs.DrainEphemeralSessions,
		DrainHost:              attrs.DrainHost,
		DrainPassword:          attrs.DrainPassword,
		DrainPort:              attrs.DrainPort,
		DrainType:              attrs.DrainType,
		Handle:                 handle,
		LoggingToken:           attrs.LoggingToken,
		URL:                    attrs.URL,
	}
	logDrain := &LogDrain{}
	params := operations.NewPostAccountsAccountIDLogDrainsParams().WithAccountID(accountID).WithAppRequest(request)
	response, err := c.Client.Operations.PostAccountsAccountIDLogDrains(params, c.Token)
	if err != nil {
		return logDrain, err
	}

	if response.Payload.ID == nil {
		return logDrain, fmt.Errorf("log drain ID is a nil pointer")
	}
	logDrain.ID = *response.Payload.ID
	logDrain.Handle = *response.Payload.Handle
	logDrain.DrainType = *response.Payload.DrainType
	logDrain.Deleted = false

	return logDrain, nil
}

func (c *Client) GetLogDrain(logDrainID int64) (*LogDrain, error) {
	logDrain := &LogDrain{
		Deleted: false,
	}

	params := operations.NewGetLogDrainsIDParams().WithID(logDrainID)
	response, err := c.Client.Operations.GetLogDrainsID(params, c.Token)
	if err != nil {
		errStruct := err.(*operations.GetLogDrainsIDDefault)
		switch errStruct.Code() {
		case 404:
			// If deleted == true, then the endpoint needs to be removed from Terraform.
			logDrain.Deleted = true
			return logDrain, nil
		case 401:
			e := fmt.Errorf("make sure you have the correct auth token")
			return logDrain, e
		default:
			e := fmt.Errorf("there was an error when completing the request to get the app \n[ERROR] -%s", err)
			return logDrain, e
		}
	}

	if response.Payload.ID == nil {
		return logDrain, fmt.Errorf("payload.ID is a nil pointer")
	}
	logDrain.ID = *response.Payload.ID
	logDrain.DrainType = *response.Payload.DrainType
	logDrain.Handle = *response.Payload.Handle
	return logDrain, nil
}

func (c *Client) DeleteLogDrain(logDrainID int64) error {
	requestType := "deprovision"
	request := models.AppRequest29{Type: &requestType}
	params := operations.NewPostLogDrainsLogDrainIDOperationsParams().WithLogDrainID(logDrainID).WithAppRequest(&request)
	op, err := c.Client.Operations.PostLogDrainsLogDrainIDOperations(params, c.Token)
	if err != nil {
		return err
	}

	operationID := *op.Payload.ID
	_, err = c.WaitForOperation(operationID)
	return err
}
