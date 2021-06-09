// Code generated by go-swagger; DO NOT EDIT.

package platformsecurityapi

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

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostPlatformSecuritySigningKeysParams creates a new PostPlatformSecuritySigningKeysParams object
// with the default values initialized.
func NewPostPlatformSecuritySigningKeysParams() *PostPlatformSecuritySigningKeysParams {
	var ()
	return &PostPlatformSecuritySigningKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPlatformSecuritySigningKeysParamsWithTimeout creates a new PostPlatformSecuritySigningKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPlatformSecuritySigningKeysParamsWithTimeout(timeout time.Duration) *PostPlatformSecuritySigningKeysParams {
	var ()
	return &PostPlatformSecuritySigningKeysParams{

		timeout: timeout,
	}
}

// NewPostPlatformSecuritySigningKeysParamsWithContext creates a new PostPlatformSecuritySigningKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPlatformSecuritySigningKeysParamsWithContext(ctx context.Context) *PostPlatformSecuritySigningKeysParams {
	var ()
	return &PostPlatformSecuritySigningKeysParams{

		Context: ctx,
	}
}

// NewPostPlatformSecuritySigningKeysParamsWithHTTPClient creates a new PostPlatformSecuritySigningKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPlatformSecuritySigningKeysParamsWithHTTPClient(client *http.Client) *PostPlatformSecuritySigningKeysParams {
	var ()
	return &PostPlatformSecuritySigningKeysParams{
		HTTPClient: client,
	}
}

/*PostPlatformSecuritySigningKeysParams contains all the parameters to send to the API endpoint
for the post platform security signing keys operation typically these are written to a http.Request
*/
type PostPlatformSecuritySigningKeysParams struct {

	/*Data*/
	Data *models.SigningKeysCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post platform security signing keys params
func (o *PostPlatformSecuritySigningKeysParams) WithTimeout(timeout time.Duration) *PostPlatformSecuritySigningKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post platform security signing keys params
func (o *PostPlatformSecuritySigningKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post platform security signing keys params
func (o *PostPlatformSecuritySigningKeysParams) WithContext(ctx context.Context) *PostPlatformSecuritySigningKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post platform security signing keys params
func (o *PostPlatformSecuritySigningKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post platform security signing keys params
func (o *PostPlatformSecuritySigningKeysParams) WithHTTPClient(client *http.Client) *PostPlatformSecuritySigningKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post platform security signing keys params
func (o *PostPlatformSecuritySigningKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the post platform security signing keys params
func (o *PostPlatformSecuritySigningKeysParams) WithData(data *models.SigningKeysCreation) *PostPlatformSecuritySigningKeysParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the post platform security signing keys params
func (o *PostPlatformSecuritySigningKeysParams) SetData(data *models.SigningKeysCreation) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PostPlatformSecuritySigningKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
