package aptible

import (
	"fmt"

	"github.com/aptible/go-deploy/client/operations"
	"github.com/aptible/go-deploy/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

type MetricDrain struct {
	Deleted       bool
	ID            int64
	Handle        string
	DrainType     string
	DatabaseID    int64
	Address       strfmt.URI
	DrainUsername string
	DrainPassword string
	DrainDatabase string
	ApiKey        string
	SeriesURL     strfmt.URI
	AccountID     int64
}

type MetricDrainCreateAttrs struct {
	DatabaseID    int64
	DrainType     *string
	LoggingToken  string
	Address       strfmt.URI
	DrainUsername string
	DrainPassword string
	DrainDatabase string
	ApiKey        string
	SeriesURL     strfmt.URI
}

func (c *Client) CreateMetricDrain(handle string, accountID int64, attrs *MetricDrainCreateAttrs) (*MetricDrain, error) {
	request := &models.AppRequest19{
		DatabaseID: attrs.DatabaseID,
		DrainType:  attrs.DrainType,
		Handle:     &handle,
	}

	// influxdb_database drains cannot have a DrainConfiguration
	if *attrs.DrainType != "influxdb_database" {
		request.DrainConfiguration = &models.AccountsaccountIdmetricDrainsDrainConfiguration{
			// TODO: Update address to a URI
			Address:  attrs.Address.String(),
			Username: attrs.DrainUsername,
			Password: attrs.DrainPassword,
			Database: attrs.DrainDatabase,
			// TODO: Add api_key and series_url to swagger schema
			//ApiKey: attrs.ApiKey,
			//SeriesURL: attrs.ApiKey,
		}
	}

	metricDrain := &MetricDrain{}
	params := operations.NewPostAccountsAccountIDMetricDrainsParams().WithAccountID(accountID).WithAppRequest(request)
	response, err := c.Client.Operations.PostAccountsAccountIDMetricDrains(params, c.Token)
	if err != nil {
		return metricDrain, err
	}

	if response.Payload.ID == nil {
		return metricDrain, fmt.Errorf("metric drain ID is a nil pointer")
	}
	metricDrain.ID = *response.Payload.ID

	requestType := "provision"
	operationRequest := &models.AppRequest30{Type: &requestType}
	operationParams := operations.NewPostMetricDrainsMetricDrainIDOperationsParams().WithMetricDrainID(metricDrain.ID).WithAppRequest(operationRequest)
	operationResponse, err := c.Client.Operations.PostMetricDrainsMetricDrainIDOperations(operationParams, c.Token)
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

	return c.GetMetricDrain(metricDrain.ID)
}

func (c *Client) GetMetricDrain(metricDrainID int64) (*MetricDrain, error) {
	metricDrain := &MetricDrain{
		Deleted: false,
	}

	params := operations.NewGetMetricDrainsIDParams().WithID(metricDrainID)
	response, err := c.Client.Operations.GetMetricDrainsID(params, c.Token)
	if err != nil {
		errStruct := err.(*operations.GetMetricDrainsIDDefault)
		switch errStruct.Code() {
		case 404:
			// If deleted == true, then the metric_drain needs to be removed from Terraform.
			metricDrain.Deleted = true
			return metricDrain, nil
		case 401:
			e := fmt.Errorf("make sure you have the correct auth token")
			return metricDrain, e
		default:
			e := fmt.Errorf("there was an error when completing the request to get the app \n[ERROR] -%s", err)
			return metricDrain, e
		}
	}

	metricDrain.ID = swag.Int64Value(response.Payload.ID)
	metricDrain.Handle = swag.StringValue(response.Payload.Handle)
	metricDrain.DrainType = swag.StringValue(response.Payload.DrainType)
	metricDrain.AccountID, _ = GetIDFromHref(response.Payload.Links.Account.Href.String())

	if response.Payload.Links.Database != nil {
		metricDrain.DatabaseID, _ = GetIDFromHref(response.Payload.Links.Database.Href.String())
	}
	return metricDrain, nil
}

func (c *Client) DeleteMetricDrain(metricDrainID int64) (bool, error) {
	requestType := "deprovision"
	request := models.AppRequest30{Type: &requestType}
	params := operations.NewPostMetricDrainsMetricDrainIDOperationsParams().WithMetricDrainID(metricDrainID).WithAppRequest(&request)
	op, err := c.Client.Operations.PostMetricDrainsMetricDrainIDOperations(params, c.Token)
	if err != nil {
		return false, err
	}

	operationID := *op.Payload.ID
	return c.WaitForOperation(operationID)
}
