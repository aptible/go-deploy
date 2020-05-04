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

// NewPatchVhostsIDParams creates a new PatchVhostsIDParams object
// with the default values initialized.
func NewPatchVhostsIDParams() *PatchVhostsIDParams {
	var ()
	return &PatchVhostsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchVhostsIDParamsWithTimeout creates a new PatchVhostsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchVhostsIDParamsWithTimeout(timeout time.Duration) *PatchVhostsIDParams {
	var ()
	return &PatchVhostsIDParams{

		timeout: timeout,
	}
}

// NewPatchVhostsIDParamsWithContext creates a new PatchVhostsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchVhostsIDParamsWithContext(ctx context.Context) *PatchVhostsIDParams {
	var ()
	return &PatchVhostsIDParams{

		Context: ctx,
	}
}

// NewPatchVhostsIDParamsWithHTTPClient creates a new PatchVhostsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchVhostsIDParamsWithHTTPClient(client *http.Client) *PatchVhostsIDParams {
	var ()
	return &PatchVhostsIDParams{
		HTTPClient: client,
	}
}

/*PatchVhostsIDParams contains all the parameters to send to the API endpoint
for the patch vhosts ID operation typically these are written to a http.Request
*/
type PatchVhostsIDParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest35
	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch vhosts ID params
func (o *PatchVhostsIDParams) WithTimeout(timeout time.Duration) *PatchVhostsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch vhosts ID params
func (o *PatchVhostsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch vhosts ID params
func (o *PatchVhostsIDParams) WithContext(ctx context.Context) *PatchVhostsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch vhosts ID params
func (o *PatchVhostsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch vhosts ID params
func (o *PatchVhostsIDParams) WithHTTPClient(client *http.Client) *PatchVhostsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch vhosts ID params
func (o *PatchVhostsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the patch vhosts ID params
func (o *PatchVhostsIDParams) WithAppRequest(appRequest *models.AppRequest35) *PatchVhostsIDParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the patch vhosts ID params
func (o *PatchVhostsIDParams) SetAppRequest(appRequest *models.AppRequest35) {
	o.AppRequest = appRequest
}

// WithID adds the id to the patch vhosts ID params
func (o *PatchVhostsIDParams) WithID(id int64) *PatchVhostsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch vhosts ID params
func (o *PatchVhostsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchVhostsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
