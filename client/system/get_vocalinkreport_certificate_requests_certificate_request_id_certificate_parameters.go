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

// NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams creates a new GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams object
// with the default values initialized.
func NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams() *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams {
	var ()
	return &GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParamsWithTimeout creates a new GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParamsWithTimeout(timeout time.Duration) *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams {
	var ()
	return &GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams{

		timeout: timeout,
	}
}

// NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParamsWithContext creates a new GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParamsWithContext(ctx context.Context) *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams {
	var ()
	return &GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams{

		Context: ctx,
	}
}

// NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParamsWithHTTPClient creates a new GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParamsWithHTTPClient(client *http.Client) *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams {
	var ()
	return &GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams{
		HTTPClient: client,
	}
}

/*GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams contains all the parameters to send to the API endpoint
for the get vocalinkreport certificate requests certificate request ID certificate operation typically these are written to a http.Request
*/
type GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams struct {

	/*CertificateRequestID
	  Certificate Request Id

	*/
	CertificateRequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vocalinkreport certificate requests certificate request ID certificate params
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) WithTimeout(timeout time.Duration) *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vocalinkreport certificate requests certificate request ID certificate params
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vocalinkreport certificate requests certificate request ID certificate params
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) WithContext(ctx context.Context) *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vocalinkreport certificate requests certificate request ID certificate params
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vocalinkreport certificate requests certificate request ID certificate params
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) WithHTTPClient(client *http.Client) *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vocalinkreport certificate requests certificate request ID certificate params
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateRequestID adds the certificateRequestID to the get vocalinkreport certificate requests certificate request ID certificate params
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) WithCertificateRequestID(certificateRequestID strfmt.UUID) *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams {
	o.SetCertificateRequestID(certificateRequestID)
	return o
}

// SetCertificateRequestID adds the certificateRequestId to the get vocalinkreport certificate requests certificate request ID certificate params
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) SetCertificateRequestID(certificateRequestID strfmt.UUID) {
	o.CertificateRequestID = certificateRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param certificate_request_id
	if err := r.SetPathParam("certificate_request_id", o.CertificateRequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
