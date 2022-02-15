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

// NewPostImagesImageIDOperationsParams creates a new PostImagesImageIDOperationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostImagesImageIDOperationsParams() *PostImagesImageIDOperationsParams {
	return &PostImagesImageIDOperationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostImagesImageIDOperationsParamsWithTimeout creates a new PostImagesImageIDOperationsParams object
// with the ability to set a timeout on a request.
func NewPostImagesImageIDOperationsParamsWithTimeout(timeout time.Duration) *PostImagesImageIDOperationsParams {
	return &PostImagesImageIDOperationsParams{
		timeout: timeout,
	}
}

// NewPostImagesImageIDOperationsParamsWithContext creates a new PostImagesImageIDOperationsParams object
// with the ability to set a context for a request.
func NewPostImagesImageIDOperationsParamsWithContext(ctx context.Context) *PostImagesImageIDOperationsParams {
	return &PostImagesImageIDOperationsParams{
		Context: ctx,
	}
}

// NewPostImagesImageIDOperationsParamsWithHTTPClient creates a new PostImagesImageIDOperationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostImagesImageIDOperationsParamsWithHTTPClient(client *http.Client) *PostImagesImageIDOperationsParams {
	return &PostImagesImageIDOperationsParams{
		HTTPClient: client,
	}
}

/* PostImagesImageIDOperationsParams contains all the parameters to send to the API endpoint
   for the post images image ID operations operation.

   Typically these are written to a http.Request.
*/
type PostImagesImageIDOperationsParams struct {

	// AppRequest.
	AppRequest *models.AppRequest28

	/* ImageID.

	   image_id
	*/
	ImageID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post images image ID operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostImagesImageIDOperationsParams) WithDefaults() *PostImagesImageIDOperationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post images image ID operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostImagesImageIDOperationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) WithTimeout(timeout time.Duration) *PostImagesImageIDOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) WithContext(ctx context.Context) *PostImagesImageIDOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) WithHTTPClient(client *http.Client) *PostImagesImageIDOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) WithAppRequest(appRequest *models.AppRequest28) *PostImagesImageIDOperationsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) SetAppRequest(appRequest *models.AppRequest28) {
	o.AppRequest = appRequest
}

// WithImageID adds the imageID to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) WithImageID(imageID int64) *PostImagesImageIDOperationsParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the post images image ID operations params
func (o *PostImagesImageIDOperationsParams) SetImageID(imageID int64) {
	o.ImageID = imageID
}

// WriteToRequest writes these params to a swagger request
func (o *PostImagesImageIDOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param image_id
	if err := r.SetPathParam("image_id", swag.FormatInt64(o.ImageID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
