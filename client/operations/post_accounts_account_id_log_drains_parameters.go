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

// NewPostAccountsAccountIDLogDrainsParams creates a new PostAccountsAccountIDLogDrainsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAccountsAccountIDLogDrainsParams() *PostAccountsAccountIDLogDrainsParams {
	return &PostAccountsAccountIDLogDrainsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountsAccountIDLogDrainsParamsWithTimeout creates a new PostAccountsAccountIDLogDrainsParams object
// with the ability to set a timeout on a request.
func NewPostAccountsAccountIDLogDrainsParamsWithTimeout(timeout time.Duration) *PostAccountsAccountIDLogDrainsParams {
	return &PostAccountsAccountIDLogDrainsParams{
		timeout: timeout,
	}
}

// NewPostAccountsAccountIDLogDrainsParamsWithContext creates a new PostAccountsAccountIDLogDrainsParams object
// with the ability to set a context for a request.
func NewPostAccountsAccountIDLogDrainsParamsWithContext(ctx context.Context) *PostAccountsAccountIDLogDrainsParams {
	return &PostAccountsAccountIDLogDrainsParams{
		Context: ctx,
	}
}

// NewPostAccountsAccountIDLogDrainsParamsWithHTTPClient creates a new PostAccountsAccountIDLogDrainsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAccountsAccountIDLogDrainsParamsWithHTTPClient(client *http.Client) *PostAccountsAccountIDLogDrainsParams {
	return &PostAccountsAccountIDLogDrainsParams{
		HTTPClient: client,
	}
}

/* PostAccountsAccountIDLogDrainsParams contains all the parameters to send to the API endpoint
   for the post accounts account ID log drains operation.

   Typically these are written to a http.Request.
*/
type PostAccountsAccountIDLogDrainsParams struct {

	/* AccountID.

	   account_id
	*/
	AccountID int64

	// AppRequest.
	AppRequest *models.AppRequest16

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post accounts account ID log drains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAccountsAccountIDLogDrainsParams) WithDefaults() *PostAccountsAccountIDLogDrainsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post accounts account ID log drains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAccountsAccountIDLogDrainsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) WithTimeout(timeout time.Duration) *PostAccountsAccountIDLogDrainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) WithContext(ctx context.Context) *PostAccountsAccountIDLogDrainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) WithHTTPClient(client *http.Client) *PostAccountsAccountIDLogDrainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) WithAccountID(accountID int64) *PostAccountsAccountIDLogDrainsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) SetAccountID(accountID int64) {
	o.AccountID = accountID
}

// WithAppRequest adds the appRequest to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) WithAppRequest(appRequest *models.AppRequest16) *PostAccountsAccountIDLogDrainsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post accounts account ID log drains params
func (o *PostAccountsAccountIDLogDrainsParams) SetAppRequest(appRequest *models.AppRequest16) {
	o.AppRequest = appRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountsAccountIDLogDrainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", swag.FormatInt64(o.AccountID)); err != nil {
		return err
	}
	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
