// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBankidsParams creates a new GetBankidsParams object
// with the default values initialized.
func NewGetBankidsParams() *GetBankidsParams {
	var ()
	return &GetBankidsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBankidsParamsWithTimeout creates a new GetBankidsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBankidsParamsWithTimeout(timeout time.Duration) *GetBankidsParams {
	var ()
	return &GetBankidsParams{

		timeout: timeout,
	}
}

// NewGetBankidsParamsWithContext creates a new GetBankidsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBankidsParamsWithContext(ctx context.Context) *GetBankidsParams {
	var ()
	return &GetBankidsParams{

		Context: ctx,
	}
}

// NewGetBankidsParamsWithHTTPClient creates a new GetBankidsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBankidsParamsWithHTTPClient(client *http.Client) *GetBankidsParams {
	var ()
	return &GetBankidsParams{
		HTTPClient: client,
	}
}

/*GetBankidsParams contains all the parameters to send to the API endpoint
for the get bankids operation typically these are written to a http.Request
*/
type GetBankidsParams struct {

	/*PageNumber
	  Which page to select

	*/
	PageNumber *int64
	/*PageSize
	  Number of items to select

	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bankids params
func (o *GetBankidsParams) WithTimeout(timeout time.Duration) *GetBankidsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bankids params
func (o *GetBankidsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bankids params
func (o *GetBankidsParams) WithContext(ctx context.Context) *GetBankidsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bankids params
func (o *GetBankidsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bankids params
func (o *GetBankidsParams) WithHTTPClient(client *http.Client) *GetBankidsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bankids params
func (o *GetBankidsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get bankids params
func (o *GetBankidsParams) WithPageNumber(pageNumber *int64) *GetBankidsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get bankids params
func (o *GetBankidsParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get bankids params
func (o *GetBankidsParams) WithPageSize(pageSize *int64) *GetBankidsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get bankids params
func (o *GetBankidsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetBankidsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber int64
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt64(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
