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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// NewPostAccountsParams creates a new PostAccountsParams object
// with the default values initialized.
func NewPostAccountsParams() *PostAccountsParams {
	var ()
	return &PostAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountsParamsWithTimeout creates a new PostAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAccountsParamsWithTimeout(timeout time.Duration) *PostAccountsParams {
	var ()
	return &PostAccountsParams{

		timeout: timeout,
	}
}

// NewPostAccountsParamsWithContext creates a new PostAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAccountsParamsWithContext(ctx context.Context) *PostAccountsParams {
	var ()
	return &PostAccountsParams{

		Context: ctx,
	}
}

// NewPostAccountsParamsWithHTTPClient creates a new PostAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAccountsParamsWithHTTPClient(client *http.Client) *PostAccountsParams {
	var ()
	return &PostAccountsParams{
		HTTPClient: client,
	}
}

/*PostAccountsParams contains all the parameters to send to the API endpoint
for the post accounts operation typically these are written to a http.Request
*/
type PostAccountsParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post accounts params
func (o *PostAccountsParams) WithTimeout(timeout time.Duration) *PostAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post accounts params
func (o *PostAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post accounts params
func (o *PostAccountsParams) WithContext(ctx context.Context) *PostAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post accounts params
func (o *PostAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post accounts params
func (o *PostAccountsParams) WithHTTPClient(client *http.Client) *PostAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post accounts params
func (o *PostAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post accounts params
func (o *PostAccountsParams) WithAppRequest(appRequest *models.AppRequest) *PostAccountsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post accounts params
func (o *PostAccountsParams) SetAppRequest(appRequest *models.AppRequest) {
	o.AppRequest = appRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
