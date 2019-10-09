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

// NewPostConfirmationOfPayeeParams creates a new PostConfirmationOfPayeeParams object
// with the default values initialized.
func NewPostConfirmationOfPayeeParams() *PostConfirmationOfPayeeParams {
	var ()
	return &PostConfirmationOfPayeeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostConfirmationOfPayeeParamsWithTimeout creates a new PostConfirmationOfPayeeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostConfirmationOfPayeeParamsWithTimeout(timeout time.Duration) *PostConfirmationOfPayeeParams {
	var ()
	return &PostConfirmationOfPayeeParams{

		timeout: timeout,
	}
}

// NewPostConfirmationOfPayeeParamsWithContext creates a new PostConfirmationOfPayeeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostConfirmationOfPayeeParamsWithContext(ctx context.Context) *PostConfirmationOfPayeeParams {
	var ()
	return &PostConfirmationOfPayeeParams{

		Context: ctx,
	}
}

// NewPostConfirmationOfPayeeParamsWithHTTPClient creates a new PostConfirmationOfPayeeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostConfirmationOfPayeeParamsWithHTTPClient(client *http.Client) *PostConfirmationOfPayeeParams {
	var ()
	return &PostConfirmationOfPayeeParams{
		HTTPClient: client,
	}
}

/*PostConfirmationOfPayeeParams contains all the parameters to send to the API endpoint
for the post confirmation of payee operation typically these are written to a http.Request
*/
type PostConfirmationOfPayeeParams struct {

	/*CreationRequest*/
	CreationRequest *models.CoPAssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post confirmation of payee params
func (o *PostConfirmationOfPayeeParams) WithTimeout(timeout time.Duration) *PostConfirmationOfPayeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post confirmation of payee params
func (o *PostConfirmationOfPayeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post confirmation of payee params
func (o *PostConfirmationOfPayeeParams) WithContext(ctx context.Context) *PostConfirmationOfPayeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post confirmation of payee params
func (o *PostConfirmationOfPayeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post confirmation of payee params
func (o *PostConfirmationOfPayeeParams) WithHTTPClient(client *http.Client) *PostConfirmationOfPayeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post confirmation of payee params
func (o *PostConfirmationOfPayeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post confirmation of payee params
func (o *PostConfirmationOfPayeeParams) WithCreationRequest(creationRequest *models.CoPAssociationCreation) *PostConfirmationOfPayeeParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post confirmation of payee params
func (o *PostConfirmationOfPayeeParams) SetCreationRequest(creationRequest *models.CoPAssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostConfirmationOfPayeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
