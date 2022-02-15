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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/aptible/go-deploy/models"
)

// NewPutDatabasesIDParams creates a new PutDatabasesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDatabasesIDParams() *PutDatabasesIDParams {
	return &PutDatabasesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDatabasesIDParamsWithTimeout creates a new PutDatabasesIDParams object
// with the ability to set a timeout on a request.
func NewPutDatabasesIDParamsWithTimeout(timeout time.Duration) *PutDatabasesIDParams {
	return &PutDatabasesIDParams{
		timeout: timeout,
	}
}

// NewPutDatabasesIDParamsWithContext creates a new PutDatabasesIDParams object
// with the ability to set a context for a request.
func NewPutDatabasesIDParamsWithContext(ctx context.Context) *PutDatabasesIDParams {
	return &PutDatabasesIDParams{
		Context: ctx,
	}
}

// NewPutDatabasesIDParamsWithHTTPClient creates a new PutDatabasesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDatabasesIDParamsWithHTTPClient(client *http.Client) *PutDatabasesIDParams {
	return &PutDatabasesIDParams{
		HTTPClient: client,
	}
}

/* PutDatabasesIDParams contains all the parameters to send to the API endpoint
   for the put databases ID operation.

   Typically these are written to a http.Request.
*/
type PutDatabasesIDParams struct {

	// AppRequest.
	AppRequest *models.AppRequest14

	/* ID.

	   id
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put databases ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDatabasesIDParams) WithDefaults() *PutDatabasesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put databases ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDatabasesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put databases ID params
func (o *PutDatabasesIDParams) WithTimeout(timeout time.Duration) *PutDatabasesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put databases ID params
func (o *PutDatabasesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put databases ID params
func (o *PutDatabasesIDParams) WithContext(ctx context.Context) *PutDatabasesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put databases ID params
func (o *PutDatabasesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put databases ID params
func (o *PutDatabasesIDParams) WithHTTPClient(client *http.Client) *PutDatabasesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put databases ID params
func (o *PutDatabasesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the put databases ID params
func (o *PutDatabasesIDParams) WithAppRequest(appRequest *models.AppRequest14) *PutDatabasesIDParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the put databases ID params
func (o *PutDatabasesIDParams) SetAppRequest(appRequest *models.AppRequest14) {
	o.AppRequest = appRequest
}

// WithID adds the id to the put databases ID params
func (o *PutDatabasesIDParams) WithID(id int64) *PutDatabasesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put databases ID params
func (o *PutDatabasesIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDatabasesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
