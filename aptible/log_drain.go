package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

type LogDrain struct {
	Deleted                bool
	ID                     int64
	Handle                 string
	DrainType              string
	DrainApps              bool
	DrainDatabases         bool
	DrainEphemeralSessions bool
	DrainProxies           bool
	DrainHost              string
	DrainPort              int64
	LoggingToken           string
	URL                    string
	DatabaseID             int64
	AccountID              int64
}

type LogDrainCreateAttrs struct {
	Database               strfmt.URI
	DatabaseID             int64
	DrainApps              bool
	DrainDatabases         bool
	DrainEphemeralSessions bool
	DrainProxies           bool
	DrainHost              strfmt.URI
	DrainPassword          string
	DrainPort              int64
	DrainType              *string
	LoggingToken           string
	URL                    strfmt.URI
}

func (c *Client) CreateLogDrain(handle string, accountID int64, attrs *LogDrainCreateAttrs) (*LogDrain, error) {
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
		Handle:                 &handle,
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

	requestType := "provision"
	operationRequest := &models.AppRequest29{Type: &requestType}
	operationParams := operations.NewPostLogDrainsLogDrainIDOperationsParams().WithLogDrainID(logDrain.ID).WithAppRequest(operationRequest)
	operationResponse, err := c.Client.Operations.PostLogDrainsLogDrainIDOperations(operationParams, c.Token)
	if err != nil {
		return nil, err
	}

	// Wait on provision operation to complete.
	if operationResponse.Payload.ID == nil {
		return nil, fmt.Errorf("operation ID is a nil pointer")
	}
	operationID := *operationResponse.Payload.ID
	_, err = c.WaitForOperation(operationID)
	if err != nil {
		return nil, err
	}

	return c.GetLogDrain(logDrain.ID)
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
			// If deleted == true, then the log_drain needs to be removed from Terraform.
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

	logDrain.ID = swag.Int64Value(response.Payload.ID)
	logDrain.DrainType = swag.StringValue(response.Payload.DrainType)
	logDrain.Handle = swag.StringValue(response.Payload.Handle)
	logDrain.DrainType = swag.StringValue(response.Payload.DrainType)
	logDrain.DrainHost = swag.StringValue(response.Payload.DrainHost)
	logDrain.DrainPort = swag.Int64Value(response.Payload.DrainPort)
	logDrain.LoggingToken = swag.StringValue(response.Payload.LoggingToken)
	logDrain.URL = swag.StringValue(response.Payload.URL)
	logDrain.DrainApps = swag.BoolValue(response.Payload.DrainApps)
	logDrain.DrainDatabases = swag.BoolValue(response.Payload.DrainDatabases)
	logDrain.DrainEphemeralSessions = swag.BoolValue(response.Payload.DrainEphemeralSessions)
	logDrain.DrainProxies = swag.BoolValue(response.Payload.DrainProxies)
	logDrain.DatabaseID, _ = GetIDFromHref(response.Payload.Links.Database.Href.String())
	logDrain.AccountID, _ = GetIDFromHref(response.Payload.Links.Account.Href.String())
	return logDrain, nil
}

func (c *Client) DeleteLogDrain(logDrainID int64) (bool, error) {
	requestType := "deprovision"
	request := models.AppRequest29{Type: &requestType}
	params := operations.NewPostLogDrainsLogDrainIDOperationsParams().WithLogDrainID(logDrainID).WithAppRequest(&request)
	op, err := c.Client.Operations.PostLogDrainsLogDrainIDOperations(params, c.Token)
	if err != nil {
		return false, err
	}

	operationID := *op.Payload.ID
	return c.WaitForOperation(operationID)
}
