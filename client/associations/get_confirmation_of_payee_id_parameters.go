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

// NewGetConfirmationOfPayeeIDParams creates a new GetConfirmationOfPayeeIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConfirmationOfPayeeIDParams() *GetConfirmationOfPayeeIDParams {
	return &GetConfirmationOfPayeeIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfirmationOfPayeeIDParamsWithTimeout creates a new GetConfirmationOfPayeeIDParams object
// with the ability to set a timeout on a request.
func NewGetConfirmationOfPayeeIDParamsWithTimeout(timeout time.Duration) *GetConfirmationOfPayeeIDParams {
	return &GetConfirmationOfPayeeIDParams{
		timeout: timeout,
	}
}

// NewGetConfirmationOfPayeeIDParamsWithContext creates a new GetConfirmationOfPayeeIDParams object
// with the ability to set a context for a request.
func NewGetConfirmationOfPayeeIDParamsWithContext(ctx context.Context) *GetConfirmationOfPayeeIDParams {
	return &GetConfirmationOfPayeeIDParams{
		Context: ctx,
	}
}

// NewGetConfirmationOfPayeeIDParamsWithHTTPClient creates a new GetConfirmationOfPayeeIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConfirmationOfPayeeIDParamsWithHTTPClient(client *http.Client) *GetConfirmationOfPayeeIDParams {
	return &GetConfirmationOfPayeeIDParams{
		HTTPClient: client,
	}
}

/* GetConfirmationOfPayeeIDParams contains all the parameters to send to the API endpoint
   for the get confirmation of payee ID operation.

   Typically these are written to a http.Request.
*/
type GetConfirmationOfPayeeIDParams struct {

	/* ID.

	   Association Id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get confirmation of payee ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfirmationOfPayeeIDParams) WithDefaults() *GetConfirmationOfPayeeIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get confirmation of payee ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfirmationOfPayeeIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get confirmation of payee ID params
func (o *GetConfirmationOfPayeeIDParams) WithTimeout(timeout time.Duration) *GetConfirmationOfPayeeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get confirmation of payee ID params
func (o *GetConfirmationOfPayeeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get confirmation of payee ID params
func (o *GetConfirmationOfPayeeIDParams) WithContext(ctx context.Context) *GetConfirmationOfPayeeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get confirmation of payee ID params
func (o *GetConfirmationOfPayeeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get confirmation of payee ID params
func (o *GetConfirmationOfPayeeIDParams) WithHTTPClient(client *http.Client) *GetConfirmationOfPayeeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get confirmation of payee ID params
func (o *GetConfirmationOfPayeeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get confirmation of payee ID params
func (o *GetConfirmationOfPayeeIDParams) WithID(id strfmt.UUID) *GetConfirmationOfPayeeIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get confirmation of payee ID params
func (o *GetConfirmationOfPayeeIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfirmationOfPayeeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
