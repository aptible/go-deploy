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
)

// NewGetAppsParams creates a new GetAppsParams object
// with the default values initialized.
func NewGetAppsParams() *GetAppsParams {
	var ()
	return &GetAppsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppsParamsWithTimeout creates a new GetAppsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppsParamsWithTimeout(timeout time.Duration) *GetAppsParams {
	var ()
	return &GetAppsParams{

		timeout: timeout,
	}
}

// NewGetAppsParamsWithContext creates a new GetAppsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppsParamsWithContext(ctx context.Context) *GetAppsParams {
	var ()
	return &GetAppsParams{

		Context: ctx,
	}
}

// NewGetAppsParamsWithHTTPClient creates a new GetAppsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppsParamsWithHTTPClient(client *http.Client) *GetAppsParams {
	var ()
	return &GetAppsParams{
		HTTPClient: client,
	}
}

/*GetAppsParams contains all the parameters to send to the API endpoint
for the get apps operation typically these are written to a http.Request
*/
type GetAppsParams struct {

	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get apps params
func (o *GetAppsParams) WithTimeout(timeout time.Duration) *GetAppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apps params
func (o *GetAppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apps params
func (o *GetAppsParams) WithContext(ctx context.Context) *GetAppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apps params
func (o *GetAppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get apps params
func (o *GetAppsParams) WithHTTPClient(client *http.Client) *GetAppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get apps params
func (o *GetAppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get apps params
func (o *GetAppsParams) WithPage(page *int64) *GetAppsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get apps params
func (o *GetAppsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
