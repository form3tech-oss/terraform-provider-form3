// Code generated by go-swagger; DO NOT EDIT.

package users

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

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostUsersParams creates a new PostUsersParams object
// with the default values initialized.
func NewPostUsersParams() *PostUsersParams {
	var ()
	return &PostUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsersParamsWithTimeout creates a new PostUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUsersParamsWithTimeout(timeout time.Duration) *PostUsersParams {
	var ()
	return &PostUsersParams{

		timeout: timeout,
	}
}

// NewPostUsersParamsWithContext creates a new PostUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUsersParamsWithContext(ctx context.Context) *PostUsersParams {
	var ()
	return &PostUsersParams{

		Context: ctx,
	}
}

// NewPostUsersParamsWithHTTPClient creates a new PostUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUsersParamsWithHTTPClient(client *http.Client) *PostUsersParams {
	var ()
	return &PostUsersParams{
		HTTPClient: client,
	}
}

/*PostUsersParams contains all the parameters to send to the API endpoint
for the post users operation typically these are written to a http.Request
*/
type PostUsersParams struct {

	/*UserCreationRequest*/
	UserCreationRequest *models.UserCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post users params
func (o *PostUsersParams) WithTimeout(timeout time.Duration) *PostUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post users params
func (o *PostUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post users params
func (o *PostUsersParams) WithContext(ctx context.Context) *PostUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post users params
func (o *PostUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post users params
func (o *PostUsersParams) WithHTTPClient(client *http.Client) *PostUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post users params
func (o *PostUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserCreationRequest adds the userCreationRequest to the post users params
func (o *PostUsersParams) WithUserCreationRequest(userCreationRequest *models.UserCreation) *PostUsersParams {
	o.SetUserCreationRequest(userCreationRequest)
	return o
}

// SetUserCreationRequest adds the userCreationRequest to the post users params
func (o *PostUsersParams) SetUserCreationRequest(userCreationRequest *models.UserCreation) {
	o.UserCreationRequest = userCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserCreationRequest != nil {
		if err := r.SetBodyParam(o.UserCreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
