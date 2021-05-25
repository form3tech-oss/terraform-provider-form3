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

// NewGetSepaLiquidityIDParams creates a new GetSepaLiquidityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSepaLiquidityIDParams() *GetSepaLiquidityIDParams {
	return &GetSepaLiquidityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSepaLiquidityIDParamsWithTimeout creates a new GetSepaLiquidityIDParams object
// with the ability to set a timeout on a request.
func NewGetSepaLiquidityIDParamsWithTimeout(timeout time.Duration) *GetSepaLiquidityIDParams {
	return &GetSepaLiquidityIDParams{
		timeout: timeout,
	}
}

// NewGetSepaLiquidityIDParamsWithContext creates a new GetSepaLiquidityIDParams object
// with the ability to set a context for a request.
func NewGetSepaLiquidityIDParamsWithContext(ctx context.Context) *GetSepaLiquidityIDParams {
	return &GetSepaLiquidityIDParams{
		Context: ctx,
	}
}

// NewGetSepaLiquidityIDParamsWithHTTPClient creates a new GetSepaLiquidityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSepaLiquidityIDParamsWithHTTPClient(client *http.Client) *GetSepaLiquidityIDParams {
	return &GetSepaLiquidityIDParams{
		HTTPClient: client,
	}
}

/* GetSepaLiquidityIDParams contains all the parameters to send to the API endpoint
   for the get sepa liquidity ID operation.

   Typically these are written to a http.Request.
*/
type GetSepaLiquidityIDParams struct {

	/* ID.

	   Association Id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get sepa liquidity ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSepaLiquidityIDParams) WithDefaults() *GetSepaLiquidityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sepa liquidity ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSepaLiquidityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sepa liquidity ID params
func (o *GetSepaLiquidityIDParams) WithTimeout(timeout time.Duration) *GetSepaLiquidityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sepa liquidity ID params
func (o *GetSepaLiquidityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sepa liquidity ID params
func (o *GetSepaLiquidityIDParams) WithContext(ctx context.Context) *GetSepaLiquidityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sepa liquidity ID params
func (o *GetSepaLiquidityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sepa liquidity ID params
func (o *GetSepaLiquidityIDParams) WithHTTPClient(client *http.Client) *GetSepaLiquidityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sepa liquidity ID params
func (o *GetSepaLiquidityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get sepa liquidity ID params
func (o *GetSepaLiquidityIDParams) WithID(id strfmt.UUID) *GetSepaLiquidityIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get sepa liquidity ID params
func (o *GetSepaLiquidityIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSepaLiquidityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
