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

// NewPostAccountsAccountIDMetricDrainsParams creates a new PostAccountsAccountIDMetricDrainsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAccountsAccountIDMetricDrainsParams() *PostAccountsAccountIDMetricDrainsParams {
	return &PostAccountsAccountIDMetricDrainsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountsAccountIDMetricDrainsParamsWithTimeout creates a new PostAccountsAccountIDMetricDrainsParams object
// with the ability to set a timeout on a request.
func NewPostAccountsAccountIDMetricDrainsParamsWithTimeout(timeout time.Duration) *PostAccountsAccountIDMetricDrainsParams {
	return &PostAccountsAccountIDMetricDrainsParams{
		timeout: timeout,
	}
}

// NewPostAccountsAccountIDMetricDrainsParamsWithContext creates a new PostAccountsAccountIDMetricDrainsParams object
// with the ability to set a context for a request.
func NewPostAccountsAccountIDMetricDrainsParamsWithContext(ctx context.Context) *PostAccountsAccountIDMetricDrainsParams {
	return &PostAccountsAccountIDMetricDrainsParams{
		Context: ctx,
	}
}

// NewPostAccountsAccountIDMetricDrainsParamsWithHTTPClient creates a new PostAccountsAccountIDMetricDrainsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAccountsAccountIDMetricDrainsParamsWithHTTPClient(client *http.Client) *PostAccountsAccountIDMetricDrainsParams {
	return &PostAccountsAccountIDMetricDrainsParams{
		HTTPClient: client,
	}
}

/* PostAccountsAccountIDMetricDrainsParams contains all the parameters to send to the API endpoint
   for the post accounts account ID metric drains operation.

   Typically these are written to a http.Request.
*/
type PostAccountsAccountIDMetricDrainsParams struct {

	/* AccountID.

	   account_id
	*/
	AccountID int64

	// AppRequest.
	AppRequest *models.AppRequest19

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post accounts account ID metric drains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAccountsAccountIDMetricDrainsParams) WithDefaults() *PostAccountsAccountIDMetricDrainsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post accounts account ID metric drains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAccountsAccountIDMetricDrainsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) WithTimeout(timeout time.Duration) *PostAccountsAccountIDMetricDrainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) WithContext(ctx context.Context) *PostAccountsAccountIDMetricDrainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) WithHTTPClient(client *http.Client) *PostAccountsAccountIDMetricDrainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) WithAccountID(accountID int64) *PostAccountsAccountIDMetricDrainsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) SetAccountID(accountID int64) {
	o.AccountID = accountID
}

// WithAppRequest adds the appRequest to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) WithAppRequest(appRequest *models.AppRequest19) *PostAccountsAccountIDMetricDrainsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post accounts account ID metric drains params
func (o *PostAccountsAccountIDMetricDrainsParams) SetAppRequest(appRequest *models.AppRequest19) {
	o.AppRequest = appRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountsAccountIDMetricDrainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
