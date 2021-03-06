// Code generated by go-swagger; DO NOT EDIT.

package account_routings

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

// NewDeleteAccountRoutingsIDParams creates a new DeleteAccountRoutingsIDParams object
// with the default values initialized.
func NewDeleteAccountRoutingsIDParams() *DeleteAccountRoutingsIDParams {
	var ()
	return &DeleteAccountRoutingsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccountRoutingsIDParamsWithTimeout creates a new DeleteAccountRoutingsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAccountRoutingsIDParamsWithTimeout(timeout time.Duration) *DeleteAccountRoutingsIDParams {
	var ()
	return &DeleteAccountRoutingsIDParams{

		timeout: timeout,
	}
}

// NewDeleteAccountRoutingsIDParamsWithContext creates a new DeleteAccountRoutingsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAccountRoutingsIDParamsWithContext(ctx context.Context) *DeleteAccountRoutingsIDParams {
	var ()
	return &DeleteAccountRoutingsIDParams{

		Context: ctx,
	}
}

// NewDeleteAccountRoutingsIDParamsWithHTTPClient creates a new DeleteAccountRoutingsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAccountRoutingsIDParamsWithHTTPClient(client *http.Client) *DeleteAccountRoutingsIDParams {
	var ()
	return &DeleteAccountRoutingsIDParams{
		HTTPClient: client,
	}
}

/*DeleteAccountRoutingsIDParams contains all the parameters to send to the API endpoint
for the delete account routings ID operation typically these are written to a http.Request
*/
type DeleteAccountRoutingsIDParams struct {

	/*ID
	  Account Routing Id

	*/
	ID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) WithTimeout(timeout time.Duration) *DeleteAccountRoutingsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) WithContext(ctx context.Context) *DeleteAccountRoutingsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) WithHTTPClient(client *http.Client) *DeleteAccountRoutingsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) WithID(id strfmt.UUID) *DeleteAccountRoutingsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) WithVersion(version int64) *DeleteAccountRoutingsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete account routings ID params
func (o *DeleteAccountRoutingsIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccountRoutingsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
