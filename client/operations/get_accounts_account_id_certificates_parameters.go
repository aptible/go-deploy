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

// NewGetAccountsAccountIDCertificatesParams creates a new GetAccountsAccountIDCertificatesParams object
// with the default values initialized.
func NewGetAccountsAccountIDCertificatesParams() *GetAccountsAccountIDCertificatesParams {
	var ()
	return &GetAccountsAccountIDCertificatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsAccountIDCertificatesParamsWithTimeout creates a new GetAccountsAccountIDCertificatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountsAccountIDCertificatesParamsWithTimeout(timeout time.Duration) *GetAccountsAccountIDCertificatesParams {
	var ()
	return &GetAccountsAccountIDCertificatesParams{

		timeout: timeout,
	}
}

// NewGetAccountsAccountIDCertificatesParamsWithContext creates a new GetAccountsAccountIDCertificatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountsAccountIDCertificatesParamsWithContext(ctx context.Context) *GetAccountsAccountIDCertificatesParams {
	var ()
	return &GetAccountsAccountIDCertificatesParams{

		Context: ctx,
	}
}

// NewGetAccountsAccountIDCertificatesParamsWithHTTPClient creates a new GetAccountsAccountIDCertificatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountsAccountIDCertificatesParamsWithHTTPClient(client *http.Client) *GetAccountsAccountIDCertificatesParams {
	var ()
	return &GetAccountsAccountIDCertificatesParams{
		HTTPClient: client,
	}
}

/*GetAccountsAccountIDCertificatesParams contains all the parameters to send to the API endpoint
for the get accounts account ID certificates operation typically these are written to a http.Request
*/
type GetAccountsAccountIDCertificatesParams struct {

	/*AccountID
	  account_id

	*/
	AccountID int64
	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) WithTimeout(timeout time.Duration) *GetAccountsAccountIDCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) WithContext(ctx context.Context) *GetAccountsAccountIDCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) WithHTTPClient(client *http.Client) *GetAccountsAccountIDCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) WithAccountID(accountID int64) *GetAccountsAccountIDCertificatesParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) SetAccountID(accountID int64) {
	o.AccountID = accountID
}

// WithPage adds the page to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) WithPage(page *int64) *GetAccountsAccountIDCertificatesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get accounts account ID certificates params
func (o *GetAccountsAccountIDCertificatesParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsAccountIDCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", swag.FormatInt64(o.AccountID)); err != nil {
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
