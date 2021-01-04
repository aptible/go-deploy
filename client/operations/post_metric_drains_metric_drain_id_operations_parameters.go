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

// NewPostMetricDrainsMetricDrainIDOperationsParams creates a new PostMetricDrainsMetricDrainIDOperationsParams object
// with the default values initialized.
func NewPostMetricDrainsMetricDrainIDOperationsParams() *PostMetricDrainsMetricDrainIDOperationsParams {
	var ()
	return &PostMetricDrainsMetricDrainIDOperationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMetricDrainsMetricDrainIDOperationsParamsWithTimeout creates a new PostMetricDrainsMetricDrainIDOperationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMetricDrainsMetricDrainIDOperationsParamsWithTimeout(timeout time.Duration) *PostMetricDrainsMetricDrainIDOperationsParams {
	var ()
	return &PostMetricDrainsMetricDrainIDOperationsParams{

		timeout: timeout,
	}
}

// NewPostMetricDrainsMetricDrainIDOperationsParamsWithContext creates a new PostMetricDrainsMetricDrainIDOperationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMetricDrainsMetricDrainIDOperationsParamsWithContext(ctx context.Context) *PostMetricDrainsMetricDrainIDOperationsParams {
	var ()
	return &PostMetricDrainsMetricDrainIDOperationsParams{

		Context: ctx,
	}
}

// NewPostMetricDrainsMetricDrainIDOperationsParamsWithHTTPClient creates a new PostMetricDrainsMetricDrainIDOperationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMetricDrainsMetricDrainIDOperationsParamsWithHTTPClient(client *http.Client) *PostMetricDrainsMetricDrainIDOperationsParams {
	var ()
	return &PostMetricDrainsMetricDrainIDOperationsParams{
		HTTPClient: client,
	}
}

/*PostMetricDrainsMetricDrainIDOperationsParams contains all the parameters to send to the API endpoint
for the post metric drains metric drain ID operations operation typically these are written to a http.Request
*/
type PostMetricDrainsMetricDrainIDOperationsParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest30
	/*MetricDrainID
	  metric_drain_id

	*/
	MetricDrainID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) WithTimeout(timeout time.Duration) *PostMetricDrainsMetricDrainIDOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) WithContext(ctx context.Context) *PostMetricDrainsMetricDrainIDOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) WithHTTPClient(client *http.Client) *PostMetricDrainsMetricDrainIDOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) WithAppRequest(appRequest *models.AppRequest30) *PostMetricDrainsMetricDrainIDOperationsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) SetAppRequest(appRequest *models.AppRequest30) {
	o.AppRequest = appRequest
}

// WithMetricDrainID adds the metricDrainID to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) WithMetricDrainID(metricDrainID int64) *PostMetricDrainsMetricDrainIDOperationsParams {
	o.SetMetricDrainID(metricDrainID)
	return o
}

// SetMetricDrainID adds the metricDrainId to the post metric drains metric drain ID operations params
func (o *PostMetricDrainsMetricDrainIDOperationsParams) SetMetricDrainID(metricDrainID int64) {
	o.MetricDrainID = metricDrainID
}

// WriteToRequest writes these params to a swagger request
func (o *PostMetricDrainsMetricDrainIDOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param metric_drain_id
	if err := r.SetPathParam("metric_drain_id", swag.FormatInt64(o.MetricDrainID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
