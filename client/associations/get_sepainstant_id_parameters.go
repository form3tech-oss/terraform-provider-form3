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

// NewGetSepainstantIDParams creates a new GetSepainstantIDParams object
// with the default values initialized.
func NewGetSepainstantIDParams() *GetSepainstantIDParams {
	var ()
	return &GetSepainstantIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSepainstantIDParamsWithTimeout creates a new GetSepainstantIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSepainstantIDParamsWithTimeout(timeout time.Duration) *GetSepainstantIDParams {
	var ()
	return &GetSepainstantIDParams{

		timeout: timeout,
	}
}

// NewGetSepainstantIDParamsWithContext creates a new GetSepainstantIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSepainstantIDParamsWithContext(ctx context.Context) *GetSepainstantIDParams {
	var ()
	return &GetSepainstantIDParams{

		Context: ctx,
	}
}

// NewGetSepainstantIDParamsWithHTTPClient creates a new GetSepainstantIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSepainstantIDParamsWithHTTPClient(client *http.Client) *GetSepainstantIDParams {
	var ()
	return &GetSepainstantIDParams{
		HTTPClient: client,
	}
}

/*GetSepainstantIDParams contains all the parameters to send to the API endpoint
for the get sepainstant ID operation typically these are written to a http.Request
*/
type GetSepainstantIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sepainstant ID params
func (o *GetSepainstantIDParams) WithTimeout(timeout time.Duration) *GetSepainstantIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sepainstant ID params
func (o *GetSepainstantIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sepainstant ID params
func (o *GetSepainstantIDParams) WithContext(ctx context.Context) *GetSepainstantIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sepainstant ID params
func (o *GetSepainstantIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sepainstant ID params
func (o *GetSepainstantIDParams) WithHTTPClient(client *http.Client) *GetSepainstantIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sepainstant ID params
func (o *GetSepainstantIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get sepainstant ID params
func (o *GetSepainstantIDParams) WithID(id strfmt.UUID) *GetSepainstantIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get sepainstant ID params
func (o *GetSepainstantIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSepainstantIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
