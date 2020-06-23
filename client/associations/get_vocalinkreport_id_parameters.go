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

// NewGetVocalinkreportIDParams creates a new GetVocalinkreportIDParams object
// with the default values initialized.
func NewGetVocalinkreportIDParams() *GetVocalinkreportIDParams {
	var ()
	return &GetVocalinkreportIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVocalinkreportIDParamsWithTimeout creates a new GetVocalinkreportIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVocalinkreportIDParamsWithTimeout(timeout time.Duration) *GetVocalinkreportIDParams {
	var ()
	return &GetVocalinkreportIDParams{

		timeout: timeout,
	}
}

// NewGetVocalinkreportIDParamsWithContext creates a new GetVocalinkreportIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVocalinkreportIDParamsWithContext(ctx context.Context) *GetVocalinkreportIDParams {
	var ()
	return &GetVocalinkreportIDParams{

		Context: ctx,
	}
}

// NewGetVocalinkreportIDParamsWithHTTPClient creates a new GetVocalinkreportIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVocalinkreportIDParamsWithHTTPClient(client *http.Client) *GetVocalinkreportIDParams {
	var ()
	return &GetVocalinkreportIDParams{
		HTTPClient: client,
	}
}

/*GetVocalinkreportIDParams contains all the parameters to send to the API endpoint
for the get vocalinkreport ID operation typically these are written to a http.Request
*/
type GetVocalinkreportIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vocalinkreport ID params
func (o *GetVocalinkreportIDParams) WithTimeout(timeout time.Duration) *GetVocalinkreportIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vocalinkreport ID params
func (o *GetVocalinkreportIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vocalinkreport ID params
func (o *GetVocalinkreportIDParams) WithContext(ctx context.Context) *GetVocalinkreportIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vocalinkreport ID params
func (o *GetVocalinkreportIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vocalinkreport ID params
func (o *GetVocalinkreportIDParams) WithHTTPClient(client *http.Client) *GetVocalinkreportIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vocalinkreport ID params
func (o *GetVocalinkreportIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vocalinkreport ID params
func (o *GetVocalinkreportIDParams) WithID(id strfmt.UUID) *GetVocalinkreportIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vocalinkreport ID params
func (o *GetVocalinkreportIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVocalinkreportIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
