// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewGetAccountconfigurationsIDParams creates a new GetAccountconfigurationsIDParams object
// with the default values initialized.
func NewGetAccountconfigurationsIDParams() *GetAccountconfigurationsIDParams {
	var ()
	return &GetAccountconfigurationsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountconfigurationsIDParamsWithTimeout creates a new GetAccountconfigurationsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountconfigurationsIDParamsWithTimeout(timeout time.Duration) *GetAccountconfigurationsIDParams {
	var ()
	return &GetAccountconfigurationsIDParams{

		timeout: timeout,
	}
}

// NewGetAccountconfigurationsIDParamsWithContext creates a new GetAccountconfigurationsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountconfigurationsIDParamsWithContext(ctx context.Context) *GetAccountconfigurationsIDParams {
	var ()
	return &GetAccountconfigurationsIDParams{

		Context: ctx,
	}
}

// NewGetAccountconfigurationsIDParamsWithHTTPClient creates a new GetAccountconfigurationsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountconfigurationsIDParamsWithHTTPClient(client *http.Client) *GetAccountconfigurationsIDParams {
	var ()
	return &GetAccountconfigurationsIDParams{
		HTTPClient: client,
	}
}

/*GetAccountconfigurationsIDParams contains all the parameters to send to the API endpoint
for the get accountconfigurations ID operation typically these are written to a http.Request
*/
type GetAccountconfigurationsIDParams struct {

	/*ID
	  AccountConfiguration Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get accountconfigurations ID params
func (o *GetAccountconfigurationsIDParams) WithTimeout(timeout time.Duration) *GetAccountconfigurationsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accountconfigurations ID params
func (o *GetAccountconfigurationsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accountconfigurations ID params
func (o *GetAccountconfigurationsIDParams) WithContext(ctx context.Context) *GetAccountconfigurationsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accountconfigurations ID params
func (o *GetAccountconfigurationsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accountconfigurations ID params
func (o *GetAccountconfigurationsIDParams) WithHTTPClient(client *http.Client) *GetAccountconfigurationsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accountconfigurations ID params
func (o *GetAccountconfigurationsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get accountconfigurations ID params
func (o *GetAccountconfigurationsIDParams) WithID(id strfmt.UUID) *GetAccountconfigurationsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get accountconfigurations ID params
func (o *GetAccountconfigurationsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountconfigurationsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
