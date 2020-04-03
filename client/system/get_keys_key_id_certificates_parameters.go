// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetKeysKeyIDCertificatesParams creates a new GetKeysKeyIDCertificatesParams object
// with the default values initialized.
func NewGetKeysKeyIDCertificatesParams() *GetKeysKeyIDCertificatesParams {
	var ()
	return &GetKeysKeyIDCertificatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeysKeyIDCertificatesParamsWithTimeout creates a new GetKeysKeyIDCertificatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeysKeyIDCertificatesParamsWithTimeout(timeout time.Duration) *GetKeysKeyIDCertificatesParams {
	var ()
	return &GetKeysKeyIDCertificatesParams{

		timeout: timeout,
	}
}

// NewGetKeysKeyIDCertificatesParamsWithContext creates a new GetKeysKeyIDCertificatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeysKeyIDCertificatesParamsWithContext(ctx context.Context) *GetKeysKeyIDCertificatesParams {
	var ()
	return &GetKeysKeyIDCertificatesParams{

		Context: ctx,
	}
}

// NewGetKeysKeyIDCertificatesParamsWithHTTPClient creates a new GetKeysKeyIDCertificatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeysKeyIDCertificatesParamsWithHTTPClient(client *http.Client) *GetKeysKeyIDCertificatesParams {
	var ()
	return &GetKeysKeyIDCertificatesParams{
		HTTPClient: client,
	}
}

/*GetKeysKeyIDCertificatesParams contains all the parameters to send to the API endpoint
for the get keys key ID certificates operation typically these are written to a http.Request
*/
type GetKeysKeyIDCertificatesParams struct {

	/*KeyID
	  Key Id

	*/
	KeyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get keys key ID certificates params
func (o *GetKeysKeyIDCertificatesParams) WithTimeout(timeout time.Duration) *GetKeysKeyIDCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keys key ID certificates params
func (o *GetKeysKeyIDCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keys key ID certificates params
func (o *GetKeysKeyIDCertificatesParams) WithContext(ctx context.Context) *GetKeysKeyIDCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keys key ID certificates params
func (o *GetKeysKeyIDCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keys key ID certificates params
func (o *GetKeysKeyIDCertificatesParams) WithHTTPClient(client *http.Client) *GetKeysKeyIDCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keys key ID certificates params
func (o *GetKeysKeyIDCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyID adds the keyID to the get keys key ID certificates params
func (o *GetKeysKeyIDCertificatesParams) WithKeyID(keyID strfmt.UUID) *GetKeysKeyIDCertificatesParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the get keys key ID certificates params
func (o *GetKeysKeyIDCertificatesParams) SetKeyID(keyID strfmt.UUID) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeysKeyIDCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key_id
	if err := r.SetPathParam("key_id", o.KeyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
