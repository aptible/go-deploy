// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// NewPostDatabasesDatabaseIDOperationsParams creates a new PostDatabasesDatabaseIDOperationsParams object
// with the default values initialized.
func NewPostDatabasesDatabaseIDOperationsParams() *PostDatabasesDatabaseIDOperationsParams {
	var ()
	return &PostDatabasesDatabaseIDOperationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDatabasesDatabaseIDOperationsParamsWithTimeout creates a new PostDatabasesDatabaseIDOperationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDatabasesDatabaseIDOperationsParamsWithTimeout(timeout time.Duration) *PostDatabasesDatabaseIDOperationsParams {
	var ()
	return &PostDatabasesDatabaseIDOperationsParams{

		timeout: timeout,
	}
}

// NewPostDatabasesDatabaseIDOperationsParamsWithContext creates a new PostDatabasesDatabaseIDOperationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDatabasesDatabaseIDOperationsParamsWithContext(ctx context.Context) *PostDatabasesDatabaseIDOperationsParams {
	var ()
	return &PostDatabasesDatabaseIDOperationsParams{

		Context: ctx,
	}
}

// NewPostDatabasesDatabaseIDOperationsParamsWithHTTPClient creates a new PostDatabasesDatabaseIDOperationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDatabasesDatabaseIDOperationsParamsWithHTTPClient(client *http.Client) *PostDatabasesDatabaseIDOperationsParams {
	var ()
	return &PostDatabasesDatabaseIDOperationsParams{
		HTTPClient: client,
	}
}

/*PostDatabasesDatabaseIDOperationsParams contains all the parameters to send to the API endpoint
for the post databases database ID operations operation typically these are written to a http.Request
*/
type PostDatabasesDatabaseIDOperationsParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest24
	/*DatabaseID
	  database_id

	*/
	DatabaseID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) WithTimeout(timeout time.Duration) *PostDatabasesDatabaseIDOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) WithContext(ctx context.Context) *PostDatabasesDatabaseIDOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) WithHTTPClient(client *http.Client) *PostDatabasesDatabaseIDOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) WithAppRequest(appRequest *models.AppRequest24) *PostDatabasesDatabaseIDOperationsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) SetAppRequest(appRequest *models.AppRequest24) {
	o.AppRequest = appRequest
}

// WithDatabaseID adds the databaseID to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) WithDatabaseID(databaseID int64) *PostDatabasesDatabaseIDOperationsParams {
	o.SetDatabaseID(databaseID)
	return o
}

// SetDatabaseID adds the databaseId to the post databases database ID operations params
func (o *PostDatabasesDatabaseIDOperationsParams) SetDatabaseID(databaseID int64) {
	o.DatabaseID = databaseID
}

// WriteToRequest writes these params to a swagger request
func (o *PostDatabasesDatabaseIDOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param database_id
	if err := r.SetPathParam("database_id", swag.FormatInt64(o.DatabaseID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
