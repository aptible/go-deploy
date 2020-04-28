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

// NewPostServicesServiceIDOperationsParams creates a new PostServicesServiceIDOperationsParams object
// with the default values initialized.
func NewPostServicesServiceIDOperationsParams() *PostServicesServiceIDOperationsParams {
	var ()
	return &PostServicesServiceIDOperationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostServicesServiceIDOperationsParamsWithTimeout creates a new PostServicesServiceIDOperationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostServicesServiceIDOperationsParamsWithTimeout(timeout time.Duration) *PostServicesServiceIDOperationsParams {
	var ()
	return &PostServicesServiceIDOperationsParams{

		timeout: timeout,
	}
}

// NewPostServicesServiceIDOperationsParamsWithContext creates a new PostServicesServiceIDOperationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostServicesServiceIDOperationsParamsWithContext(ctx context.Context) *PostServicesServiceIDOperationsParams {
	var ()
	return &PostServicesServiceIDOperationsParams{

		Context: ctx,
	}
}

// NewPostServicesServiceIDOperationsParamsWithHTTPClient creates a new PostServicesServiceIDOperationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostServicesServiceIDOperationsParamsWithHTTPClient(client *http.Client) *PostServicesServiceIDOperationsParams {
	var ()
	return &PostServicesServiceIDOperationsParams{
		HTTPClient: client,
	}
}

/*PostServicesServiceIDOperationsParams contains all the parameters to send to the API endpoint
for the post services service ID operations operation typically these are written to a http.Request
*/
type PostServicesServiceIDOperationsParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest25
	/*ServiceID
	  service_id

	*/
	ServiceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) WithTimeout(timeout time.Duration) *PostServicesServiceIDOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) WithContext(ctx context.Context) *PostServicesServiceIDOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) WithHTTPClient(client *http.Client) *PostServicesServiceIDOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) WithAppRequest(appRequest *models.AppRequest25) *PostServicesServiceIDOperationsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) SetAppRequest(appRequest *models.AppRequest25) {
	o.AppRequest = appRequest
}

// WithServiceID adds the serviceID to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) WithServiceID(serviceID int64) *PostServicesServiceIDOperationsParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the post services service ID operations params
func (o *PostServicesServiceIDOperationsParams) SetServiceID(serviceID int64) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostServicesServiceIDOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
