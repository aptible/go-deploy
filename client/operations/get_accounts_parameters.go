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
)

// NewGetAccountsParams creates a new GetAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountsParams() *GetAccountsParams {
	return &GetAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsParamsWithTimeout creates a new GetAccountsParams object
// with the ability to set a timeout on a request.
func NewGetAccountsParamsWithTimeout(timeout time.Duration) *GetAccountsParams {
	return &GetAccountsParams{
		timeout: timeout,
	}
}

// NewGetAccountsParamsWithContext creates a new GetAccountsParams object
// with the ability to set a context for a request.
func NewGetAccountsParamsWithContext(ctx context.Context) *GetAccountsParams {
	return &GetAccountsParams{
		Context: ctx,
	}
}

// NewGetAccountsParamsWithHTTPClient creates a new GetAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountsParamsWithHTTPClient(client *http.Client) *GetAccountsParams {
	return &GetAccountsParams{
		HTTPClient: client,
	}
}

/* GetAccountsParams contains all the parameters to send to the API endpoint
   for the get accounts operation.

   Typically these are written to a http.Request.
*/
type GetAccountsParams struct {

	/* Page.

	   current page of results for pagination
	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsParams) WithDefaults() *GetAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) WithTimeout(timeout time.Duration) *GetAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts params
func (o *GetAccountsParams) WithContext(ctx context.Context) *GetAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts params
func (o *GetAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts params
func (o *GetAccountsParams) WithHTTPClient(client *http.Client) *GetAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts params
func (o *GetAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get accounts params
func (o *GetAccountsParams) WithPage(page *int64) *GetAccountsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get accounts params
func (o *GetAccountsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
