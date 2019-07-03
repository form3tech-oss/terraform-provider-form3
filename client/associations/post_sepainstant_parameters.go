// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostSepainstantParams creates a new PostSepainstantParams object
// with the default values initialized.
func NewPostSepainstantParams() *PostSepainstantParams {
	var ()
	return &PostSepainstantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSepainstantParamsWithTimeout creates a new PostSepainstantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSepainstantParamsWithTimeout(timeout time.Duration) *PostSepainstantParams {
	var ()
	return &PostSepainstantParams{

		timeout: timeout,
	}
}

// NewPostSepainstantParamsWithContext creates a new PostSepainstantParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSepainstantParamsWithContext(ctx context.Context) *PostSepainstantParams {
	var ()
	return &PostSepainstantParams{

		Context: ctx,
	}
}

// NewPostSepainstantParamsWithHTTPClient creates a new PostSepainstantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSepainstantParamsWithHTTPClient(client *http.Client) *PostSepainstantParams {
	var ()
	return &PostSepainstantParams{
		HTTPClient: client,
	}
}

/*PostSepainstantParams contains all the parameters to send to the API endpoint
for the post sepainstant operation typically these are written to a http.Request
*/
type PostSepainstantParams struct {

	/*CreationRequest*/
	CreationRequest *models.SepaInstantAssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sepainstant params
func (o *PostSepainstantParams) WithTimeout(timeout time.Duration) *PostSepainstantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sepainstant params
func (o *PostSepainstantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sepainstant params
func (o *PostSepainstantParams) WithContext(ctx context.Context) *PostSepainstantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sepainstant params
func (o *PostSepainstantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sepainstant params
func (o *PostSepainstantParams) WithHTTPClient(client *http.Client) *PostSepainstantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sepainstant params
func (o *PostSepainstantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post sepainstant params
func (o *PostSepainstantParams) WithCreationRequest(creationRequest *models.SepaInstantAssociationCreation) *PostSepainstantParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post sepainstant params
func (o *PostSepainstantParams) SetCreationRequest(creationRequest *models.SepaInstantAssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostSepainstantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
