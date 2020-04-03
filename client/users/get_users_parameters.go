// Code generated by go-swagger; DO NOT EDIT.

package users

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
)

// NewGetUsersParams creates a new GetUsersParams object
// with the default values initialized.
func NewGetUsersParams() *GetUsersParams {
	var ()
	return &GetUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersParamsWithTimeout creates a new GetUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersParamsWithTimeout(timeout time.Duration) *GetUsersParams {
	var ()
	return &GetUsersParams{

		timeout: timeout,
	}
}

// NewGetUsersParamsWithContext creates a new GetUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersParamsWithContext(ctx context.Context) *GetUsersParams {
	var ()
	return &GetUsersParams{

		Context: ctx,
	}
}

// NewGetUsersParamsWithHTTPClient creates a new GetUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersParamsWithHTTPClient(client *http.Client) *GetUsersParams {
	var ()
	return &GetUsersParams{
		HTTPClient: client,
	}
}

/*GetUsersParams contains all the parameters to send to the API endpoint
for the get users operation typically these are written to a http.Request
*/
type GetUsersParams struct {

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

// WithTimeout adds the timeout to the get users params
func (o *GetUsersParams) WithTimeout(timeout time.Duration) *GetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users params
func (o *GetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users params
func (o *GetUsersParams) WithContext(ctx context.Context) *GetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users params
func (o *GetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) WithHTTPClient(client *http.Client) *GetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get users params
func (o *GetUsersParams) WithPageNumber(pageNumber *int64) *GetUsersParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get users params
func (o *GetUsersParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get users params
func (o *GetUsersParams) WithPageSize(pageSize *int64) *GetUsersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get users params
func (o *GetUsersParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
