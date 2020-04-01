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

	models "github.com/reggregory/go-deploy/models"
)

// NewPatchAppsIDParams creates a new PatchAppsIDParams object
// with the default values initialized.
func NewPatchAppsIDParams() *PatchAppsIDParams {
	var ()
	return &PatchAppsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAppsIDParamsWithTimeout creates a new PatchAppsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAppsIDParamsWithTimeout(timeout time.Duration) *PatchAppsIDParams {
	var ()
	return &PatchAppsIDParams{

		timeout: timeout,
	}
}

// NewPatchAppsIDParamsWithContext creates a new PatchAppsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAppsIDParamsWithContext(ctx context.Context) *PatchAppsIDParams {
	var ()
	return &PatchAppsIDParams{

		Context: ctx,
	}
}

// NewPatchAppsIDParamsWithHTTPClient creates a new PatchAppsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAppsIDParamsWithHTTPClient(client *http.Client) *PatchAppsIDParams {
	var ()
	return &PatchAppsIDParams{
		HTTPClient: client,
	}
}

/*PatchAppsIDParams contains all the parameters to send to the API endpoint
for the patch apps ID operation typically these are written to a http.Request
*/
type PatchAppsIDParams struct {

	/*ID
	  id

	*/
	ID int64
	/*PatchRequest*/
	PatchRequest *models.PatchRequest1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch apps ID params
func (o *PatchAppsIDParams) WithTimeout(timeout time.Duration) *PatchAppsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch apps ID params
func (o *PatchAppsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch apps ID params
func (o *PatchAppsIDParams) WithContext(ctx context.Context) *PatchAppsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch apps ID params
func (o *PatchAppsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch apps ID params
func (o *PatchAppsIDParams) WithHTTPClient(client *http.Client) *PatchAppsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch apps ID params
func (o *PatchAppsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch apps ID params
func (o *PatchAppsIDParams) WithID(id int64) *PatchAppsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch apps ID params
func (o *PatchAppsIDParams) SetID(id int64) {
	o.ID = id
}

// WithPatchRequest adds the patchRequest to the patch apps ID params
func (o *PatchAppsIDParams) WithPatchRequest(patchRequest *models.PatchRequest1) *PatchAppsIDParams {
	o.SetPatchRequest(patchRequest)
	return o
}

// SetPatchRequest adds the patchRequest to the patch apps ID params
func (o *PatchAppsIDParams) SetPatchRequest(patchRequest *models.PatchRequest1) {
	o.PatchRequest = patchRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAppsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.PatchRequest != nil {
		if err := r.SetBodyParam(o.PatchRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
