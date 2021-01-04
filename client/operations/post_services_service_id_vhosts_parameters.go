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

// NewPostServicesServiceIDVhostsParams creates a new PostServicesServiceIDVhostsParams object
// with the default values initialized.
func NewPostServicesServiceIDVhostsParams() *PostServicesServiceIDVhostsParams {
	var ()
	return &PostServicesServiceIDVhostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostServicesServiceIDVhostsParamsWithTimeout creates a new PostServicesServiceIDVhostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostServicesServiceIDVhostsParamsWithTimeout(timeout time.Duration) *PostServicesServiceIDVhostsParams {
	var ()
	return &PostServicesServiceIDVhostsParams{

		timeout: timeout,
	}
}

// NewPostServicesServiceIDVhostsParamsWithContext creates a new PostServicesServiceIDVhostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostServicesServiceIDVhostsParamsWithContext(ctx context.Context) *PostServicesServiceIDVhostsParams {
	var ()
	return &PostServicesServiceIDVhostsParams{

		Context: ctx,
	}
}

// NewPostServicesServiceIDVhostsParamsWithHTTPClient creates a new PostServicesServiceIDVhostsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostServicesServiceIDVhostsParamsWithHTTPClient(client *http.Client) *PostServicesServiceIDVhostsParams {
	var ()
	return &PostServicesServiceIDVhostsParams{
		HTTPClient: client,
	}
}

/*PostServicesServiceIDVhostsParams contains all the parameters to send to the API endpoint
for the post services service ID vhosts operation typically these are written to a http.Request
*/
type PostServicesServiceIDVhostsParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest34
	/*ServiceID
	  service_id

	*/
	ServiceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) WithTimeout(timeout time.Duration) *PostServicesServiceIDVhostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) WithContext(ctx context.Context) *PostServicesServiceIDVhostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) WithHTTPClient(client *http.Client) *PostServicesServiceIDVhostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) WithAppRequest(appRequest *models.AppRequest34) *PostServicesServiceIDVhostsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) SetAppRequest(appRequest *models.AppRequest34) {
	o.AppRequest = appRequest
}

// WithServiceID adds the serviceID to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) WithServiceID(serviceID int64) *PostServicesServiceIDVhostsParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the post services service ID vhosts params
func (o *PostServicesServiceIDVhostsParams) SetServiceID(serviceID int64) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostServicesServiceIDVhostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param service_id
	if err := r.SetPathParam("service_id", swag.FormatInt64(o.ServiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
