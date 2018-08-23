// Code generated by go-swagger; DO NOT EDIT.

package limits

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

	models "github.com/form3tech-oss/go-form3/models"
)

// NewPostLimitsParams creates a new PostLimitsParams object
// with the default values initialized.
func NewPostLimitsParams() *PostLimitsParams {
	var ()
	return &PostLimitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLimitsParamsWithTimeout creates a new PostLimitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLimitsParamsWithTimeout(timeout time.Duration) *PostLimitsParams {
	var ()
	return &PostLimitsParams{

		timeout: timeout,
	}
}

// NewPostLimitsParamsWithContext creates a new PostLimitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLimitsParamsWithContext(ctx context.Context) *PostLimitsParams {
	var ()
	return &PostLimitsParams{

		Context: ctx,
	}
}

// NewPostLimitsParamsWithHTTPClient creates a new PostLimitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLimitsParamsWithHTTPClient(client *http.Client) *PostLimitsParams {
	var ()
	return &PostLimitsParams{
		HTTPClient: client,
	}
}

/*PostLimitsParams contains all the parameters to send to the API endpoint
for the post limits operation typically these are written to a http.Request
*/
type PostLimitsParams struct {

	/*LimitCreationRequest*/
	LimitCreationRequest *models.LimitCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post limits params
func (o *PostLimitsParams) WithTimeout(timeout time.Duration) *PostLimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post limits params
func (o *PostLimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post limits params
func (o *PostLimitsParams) WithContext(ctx context.Context) *PostLimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post limits params
func (o *PostLimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post limits params
func (o *PostLimitsParams) WithHTTPClient(client *http.Client) *PostLimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post limits params
func (o *PostLimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimitCreationRequest adds the limitCreationRequest to the post limits params
func (o *PostLimitsParams) WithLimitCreationRequest(limitCreationRequest *models.LimitCreation) *PostLimitsParams {
	o.SetLimitCreationRequest(limitCreationRequest)
	return o
}

// SetLimitCreationRequest adds the limitCreationRequest to the post limits params
func (o *PostLimitsParams) SetLimitCreationRequest(limitCreationRequest *models.LimitCreation) {
	o.LimitCreationRequest = limitCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostLimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LimitCreationRequest != nil {
		if err := r.SetBodyParam(o.LimitCreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
