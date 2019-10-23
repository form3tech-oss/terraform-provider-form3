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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetKeysKeyIDCertificatesCertificateIDParams creates a new GetKeysKeyIDCertificatesCertificateIDParams object
// with the default values initialized.
func NewGetKeysKeyIDCertificatesCertificateIDParams() *GetKeysKeyIDCertificatesCertificateIDParams {
	var ()
	return &GetKeysKeyIDCertificatesCertificateIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeysKeyIDCertificatesCertificateIDParamsWithTimeout creates a new GetKeysKeyIDCertificatesCertificateIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeysKeyIDCertificatesCertificateIDParamsWithTimeout(timeout time.Duration) *GetKeysKeyIDCertificatesCertificateIDParams {
	var ()
	return &GetKeysKeyIDCertificatesCertificateIDParams{

		timeout: timeout,
	}
}

// NewGetKeysKeyIDCertificatesCertificateIDParamsWithContext creates a new GetKeysKeyIDCertificatesCertificateIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeysKeyIDCertificatesCertificateIDParamsWithContext(ctx context.Context) *GetKeysKeyIDCertificatesCertificateIDParams {
	var ()
	return &GetKeysKeyIDCertificatesCertificateIDParams{

		Context: ctx,
	}
}

// NewGetKeysKeyIDCertificatesCertificateIDParamsWithHTTPClient creates a new GetKeysKeyIDCertificatesCertificateIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeysKeyIDCertificatesCertificateIDParamsWithHTTPClient(client *http.Client) *GetKeysKeyIDCertificatesCertificateIDParams {
	var ()
	return &GetKeysKeyIDCertificatesCertificateIDParams{
		HTTPClient: client,
	}
}

/*GetKeysKeyIDCertificatesCertificateIDParams contains all the parameters to send to the API endpoint
for the get keys key ID certificates certificate ID operation typically these are written to a http.Request
*/
type GetKeysKeyIDCertificatesCertificateIDParams struct {

	/*CertificateID
	  Certificate Id

	*/
	CertificateID strfmt.UUID
	/*KeyID
	  Key Id

	*/
	KeyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) WithTimeout(timeout time.Duration) *GetKeysKeyIDCertificatesCertificateIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) WithContext(ctx context.Context) *GetKeysKeyIDCertificatesCertificateIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) WithHTTPClient(client *http.Client) *GetKeysKeyIDCertificatesCertificateIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateID adds the certificateID to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) WithCertificateID(certificateID strfmt.UUID) *GetKeysKeyIDCertificatesCertificateIDParams {
	o.SetCertificateID(certificateID)
	return o
}

// SetCertificateID adds the certificateId to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) SetCertificateID(certificateID strfmt.UUID) {
	o.CertificateID = certificateID
}

// WithKeyID adds the keyID to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) WithKeyID(keyID strfmt.UUID) *GetKeysKeyIDCertificatesCertificateIDParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the get keys key ID certificates certificate ID params
func (o *GetKeysKeyIDCertificatesCertificateIDParams) SetKeyID(keyID strfmt.UUID) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeysKeyIDCertificatesCertificateIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param certificate_id
	if err := r.SetPathParam("certificate_id", o.CertificateID.String()); err != nil {
		return err
	}

	// path param key_id
	if err := r.SetPathParam("key_id", o.KeyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
