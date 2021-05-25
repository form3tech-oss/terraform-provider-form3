// Code generated by go-swagger; DO NOT EDIT.

package associations

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
)

// NewGetBacsIDParams creates a new GetBacsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBacsIDParams() *GetBacsIDParams {
	return &GetBacsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBacsIDParamsWithTimeout creates a new GetBacsIDParams object
// with the ability to set a timeout on a request.
func NewGetBacsIDParamsWithTimeout(timeout time.Duration) *GetBacsIDParams {
	return &GetBacsIDParams{
		timeout: timeout,
	}
}

// NewGetBacsIDParamsWithContext creates a new GetBacsIDParams object
// with the ability to set a context for a request.
func NewGetBacsIDParamsWithContext(ctx context.Context) *GetBacsIDParams {
	return &GetBacsIDParams{
		Context: ctx,
	}
}

// NewGetBacsIDParamsWithHTTPClient creates a new GetBacsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBacsIDParamsWithHTTPClient(client *http.Client) *GetBacsIDParams {
	return &GetBacsIDParams{
		HTTPClient: client,
	}
}

/* GetBacsIDParams contains all the parameters to send to the API endpoint
   for the get bacs ID operation.

   Typically these are written to a http.Request.
*/
type GetBacsIDParams struct {

	/* ID.

	   Association Id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bacs ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBacsIDParams) WithDefaults() *GetBacsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bacs ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBacsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bacs ID params
func (o *GetBacsIDParams) WithTimeout(timeout time.Duration) *GetBacsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bacs ID params
func (o *GetBacsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bacs ID params
func (o *GetBacsIDParams) WithContext(ctx context.Context) *GetBacsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bacs ID params
func (o *GetBacsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bacs ID params
func (o *GetBacsIDParams) WithHTTPClient(client *http.Client) *GetBacsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bacs ID params
func (o *GetBacsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get bacs ID params
func (o *GetBacsIDParams) WithID(id strfmt.UUID) *GetBacsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get bacs ID params
func (o *GetBacsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBacsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
