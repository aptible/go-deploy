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

// NewPatchCertificatesIDParams creates a new PatchCertificatesIDParams object
// with the default values initialized.
func NewPatchCertificatesIDParams() *PatchCertificatesIDParams {
	var ()
	return &PatchCertificatesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCertificatesIDParamsWithTimeout creates a new PatchCertificatesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCertificatesIDParamsWithTimeout(timeout time.Duration) *PatchCertificatesIDParams {
	var ()
	return &PatchCertificatesIDParams{

		timeout: timeout,
	}
}

// NewPatchCertificatesIDParamsWithContext creates a new PatchCertificatesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCertificatesIDParamsWithContext(ctx context.Context) *PatchCertificatesIDParams {
	var ()
	return &PatchCertificatesIDParams{

		Context: ctx,
	}
}

// NewPatchCertificatesIDParamsWithHTTPClient creates a new PatchCertificatesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCertificatesIDParamsWithHTTPClient(client *http.Client) *PatchCertificatesIDParams {
	var ()
	return &PatchCertificatesIDParams{
		HTTPClient: client,
	}
}

/*PatchCertificatesIDParams contains all the parameters to send to the API endpoint
for the patch certificates ID operation typically these are written to a http.Request
*/
type PatchCertificatesIDParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest6
	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch certificates ID params
func (o *PatchCertificatesIDParams) WithTimeout(timeout time.Duration) *PatchCertificatesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch certificates ID params
func (o *PatchCertificatesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch certificates ID params
func (o *PatchCertificatesIDParams) WithContext(ctx context.Context) *PatchCertificatesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch certificates ID params
func (o *PatchCertificatesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch certificates ID params
func (o *PatchCertificatesIDParams) WithHTTPClient(client *http.Client) *PatchCertificatesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch certificates ID params
func (o *PatchCertificatesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the patch certificates ID params
func (o *PatchCertificatesIDParams) WithAppRequest(appRequest *models.AppRequest6) *PatchCertificatesIDParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the patch certificates ID params
func (o *PatchCertificatesIDParams) SetAppRequest(appRequest *models.AppRequest6) {
	o.AppRequest = appRequest
}

// WithID adds the id to the patch certificates ID params
func (o *PatchCertificatesIDParams) WithID(id int64) *PatchCertificatesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch certificates ID params
func (o *PatchCertificatesIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCertificatesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
