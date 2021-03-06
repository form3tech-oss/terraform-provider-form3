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

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostSepasctParams creates a new PostSepasctParams object
// with the default values initialized.
func NewPostSepasctParams() *PostSepasctParams {
	var ()
	return &PostSepasctParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSepasctParamsWithTimeout creates a new PostSepasctParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSepasctParamsWithTimeout(timeout time.Duration) *PostSepasctParams {
	var ()
	return &PostSepasctParams{

		timeout: timeout,
	}
}

// NewPostSepasctParamsWithContext creates a new PostSepasctParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSepasctParamsWithContext(ctx context.Context) *PostSepasctParams {
	var ()
	return &PostSepasctParams{

		Context: ctx,
	}
}

// NewPostSepasctParamsWithHTTPClient creates a new PostSepasctParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSepasctParamsWithHTTPClient(client *http.Client) *PostSepasctParams {
	var ()
	return &PostSepasctParams{
		HTTPClient: client,
	}
}

/*PostSepasctParams contains all the parameters to send to the API endpoint
for the post sepasct operation typically these are written to a http.Request
*/
type PostSepasctParams struct {

	/*CreationRequest*/
	CreationRequest *models.SepaSctAssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sepasct params
func (o *PostSepasctParams) WithTimeout(timeout time.Duration) *PostSepasctParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sepasct params
func (o *PostSepasctParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sepasct params
func (o *PostSepasctParams) WithContext(ctx context.Context) *PostSepasctParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sepasct params
func (o *PostSepasctParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sepasct params
func (o *PostSepasctParams) WithHTTPClient(client *http.Client) *PostSepasctParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sepasct params
func (o *PostSepasctParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post sepasct params
func (o *PostSepasctParams) WithCreationRequest(creationRequest *models.SepaSctAssociationCreation) *PostSepasctParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post sepasct params
func (o *PostSepasctParams) SetCreationRequest(creationRequest *models.SepaSctAssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostSepasctParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreationRequest != nil {
		if err := r.SetBodyParam(o.CreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
