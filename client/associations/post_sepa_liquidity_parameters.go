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

// NewPostSepaLiquidityParams creates a new PostSepaLiquidityParams object
// with the default values initialized.
func NewPostSepaLiquidityParams() *PostSepaLiquidityParams {
	var ()
	return &PostSepaLiquidityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSepaLiquidityParamsWithTimeout creates a new PostSepaLiquidityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSepaLiquidityParamsWithTimeout(timeout time.Duration) *PostSepaLiquidityParams {
	var ()
	return &PostSepaLiquidityParams{

		timeout: timeout,
	}
}

// NewPostSepaLiquidityParamsWithContext creates a new PostSepaLiquidityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSepaLiquidityParamsWithContext(ctx context.Context) *PostSepaLiquidityParams {
	var ()
	return &PostSepaLiquidityParams{

		Context: ctx,
	}
}

// NewPostSepaLiquidityParamsWithHTTPClient creates a new PostSepaLiquidityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSepaLiquidityParamsWithHTTPClient(client *http.Client) *PostSepaLiquidityParams {
	var ()
	return &PostSepaLiquidityParams{
		HTTPClient: client,
	}
}

/*PostSepaLiquidityParams contains all the parameters to send to the API endpoint
for the post sepa liquidity operation typically these are written to a http.Request
*/
type PostSepaLiquidityParams struct {

	/*CreationRequest*/
	CreationRequest *models.SepaLiquidityAssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sepa liquidity params
func (o *PostSepaLiquidityParams) WithTimeout(timeout time.Duration) *PostSepaLiquidityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sepa liquidity params
func (o *PostSepaLiquidityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sepa liquidity params
func (o *PostSepaLiquidityParams) WithContext(ctx context.Context) *PostSepaLiquidityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sepa liquidity params
func (o *PostSepaLiquidityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sepa liquidity params
func (o *PostSepaLiquidityParams) WithHTTPClient(client *http.Client) *PostSepaLiquidityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sepa liquidity params
func (o *PostSepaLiquidityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post sepa liquidity params
func (o *PostSepaLiquidityParams) WithCreationRequest(creationRequest *models.SepaLiquidityAssociationCreation) *PostSepaLiquidityParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post sepa liquidity params
func (o *PostSepaLiquidityParams) SetCreationRequest(creationRequest *models.SepaLiquidityAssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostSepaLiquidityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
