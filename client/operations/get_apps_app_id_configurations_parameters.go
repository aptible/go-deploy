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

// NewGetAppsAppIDConfigurationsParams creates a new GetAppsAppIDConfigurationsParams object
// with the default values initialized.
func NewGetAppsAppIDConfigurationsParams() *GetAppsAppIDConfigurationsParams {
	var ()
	return &GetAppsAppIDConfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppsAppIDConfigurationsParamsWithTimeout creates a new GetAppsAppIDConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppsAppIDConfigurationsParamsWithTimeout(timeout time.Duration) *GetAppsAppIDConfigurationsParams {
	var ()
	return &GetAppsAppIDConfigurationsParams{

		timeout: timeout,
	}
}

// NewGetAppsAppIDConfigurationsParamsWithContext creates a new GetAppsAppIDConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppsAppIDConfigurationsParamsWithContext(ctx context.Context) *GetAppsAppIDConfigurationsParams {
	var ()
	return &GetAppsAppIDConfigurationsParams{

		Context: ctx,
	}
}

// NewGetAppsAppIDConfigurationsParamsWithHTTPClient creates a new GetAppsAppIDConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppsAppIDConfigurationsParamsWithHTTPClient(client *http.Client) *GetAppsAppIDConfigurationsParams {
	var ()
	return &GetAppsAppIDConfigurationsParams{
		HTTPClient: client,
	}
}

/*GetAppsAppIDConfigurationsParams contains all the parameters to send to the API endpoint
for the get apps app ID configurations operation typically these are written to a http.Request
*/
type GetAppsAppIDConfigurationsParams struct {

	/*AppID
	  app_id

	*/
	AppID int64
	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) WithTimeout(timeout time.Duration) *GetAppsAppIDConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) WithContext(ctx context.Context) *GetAppsAppIDConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) WithHTTPClient(client *http.Client) *GetAppsAppIDConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) WithAppID(appID int64) *GetAppsAppIDConfigurationsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) SetAppID(appID int64) {
	o.AppID = appID
}

// WithPage adds the page to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) WithPage(page *int64) *GetAppsAppIDConfigurationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get apps app ID configurations params
func (o *GetAppsAppIDConfigurationsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppsAppIDConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", swag.FormatInt64(o.AppID)); err != nil {
		return err
	}

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
