// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewPostRolesParams creates a new PostRolesParams object
// with the default values initialized.
func NewPostRolesParams() *PostRolesParams {
	var ()
	return &PostRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRolesParamsWithTimeout creates a new PostRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRolesParamsWithTimeout(timeout time.Duration) *PostRolesParams {
	var ()
	return &PostRolesParams{

		timeout: timeout,
	}
}

// NewPostRolesParamsWithContext creates a new PostRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRolesParamsWithContext(ctx context.Context) *PostRolesParams {
	var ()
	return &PostRolesParams{

		Context: ctx,
	}
}

// NewPostRolesParamsWithHTTPClient creates a new PostRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRolesParamsWithHTTPClient(client *http.Client) *PostRolesParams {
	var ()
	return &PostRolesParams{
		HTTPClient: client,
	}
}

/*PostRolesParams contains all the parameters to send to the API endpoint
for the post roles operation typically these are written to a http.Request
*/
type PostRolesParams struct {

	/*RoleCreationRequest*/
	RoleCreationRequest *models.RoleCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post roles params
func (o *PostRolesParams) WithTimeout(timeout time.Duration) *PostRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post roles params
func (o *PostRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post roles params
func (o *PostRolesParams) WithContext(ctx context.Context) *PostRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post roles params
func (o *PostRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post roles params
func (o *PostRolesParams) WithHTTPClient(client *http.Client) *PostRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post roles params
func (o *PostRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleCreationRequest adds the roleCreationRequest to the post roles params
func (o *PostRolesParams) WithRoleCreationRequest(roleCreationRequest *models.RoleCreation) *PostRolesParams {
	o.SetRoleCreationRequest(roleCreationRequest)
	return o
}

// SetRoleCreationRequest adds the roleCreationRequest to the post roles params
func (o *PostRolesParams) SetRoleCreationRequest(roleCreationRequest *models.RoleCreation) {
	o.RoleCreationRequest = roleCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RoleCreationRequest != nil {
		if err := r.SetBodyParam(o.RoleCreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
