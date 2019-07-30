// Code generated by go-swagger; DO NOT EDIT.

package ace

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

// NewPostRolesRoleIDAcesParams creates a new PostRolesRoleIDAcesParams object
// with the default values initialized.
func NewPostRolesRoleIDAcesParams() *PostRolesRoleIDAcesParams {
	var ()
	return &PostRolesRoleIDAcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRolesRoleIDAcesParamsWithTimeout creates a new PostRolesRoleIDAcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRolesRoleIDAcesParamsWithTimeout(timeout time.Duration) *PostRolesRoleIDAcesParams {
	var ()
	return &PostRolesRoleIDAcesParams{

		timeout: timeout,
	}
}

// NewPostRolesRoleIDAcesParamsWithContext creates a new PostRolesRoleIDAcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRolesRoleIDAcesParamsWithContext(ctx context.Context) *PostRolesRoleIDAcesParams {
	var ()
	return &PostRolesRoleIDAcesParams{

		Context: ctx,
	}
}

// NewPostRolesRoleIDAcesParamsWithHTTPClient creates a new PostRolesRoleIDAcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRolesRoleIDAcesParamsWithHTTPClient(client *http.Client) *PostRolesRoleIDAcesParams {
	var ()
	return &PostRolesRoleIDAcesParams{
		HTTPClient: client,
	}
}

/*PostRolesRoleIDAcesParams contains all the parameters to send to the API endpoint
for the post roles role ID aces operation typically these are written to a http.Request
*/
type PostRolesRoleIDAcesParams struct {

	/*AceCreationRequest*/
	AceCreationRequest *models.AceCreation
	/*RoleID
	  Role Id

	*/
	RoleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) WithTimeout(timeout time.Duration) *PostRolesRoleIDAcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) WithContext(ctx context.Context) *PostRolesRoleIDAcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) WithHTTPClient(client *http.Client) *PostRolesRoleIDAcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAceCreationRequest adds the aceCreationRequest to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) WithAceCreationRequest(aceCreationRequest *models.AceCreation) *PostRolesRoleIDAcesParams {
	o.SetAceCreationRequest(aceCreationRequest)
	return o
}

// SetAceCreationRequest adds the aceCreationRequest to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) SetAceCreationRequest(aceCreationRequest *models.AceCreation) {
	o.AceCreationRequest = aceCreationRequest
}

// WithRoleID adds the roleID to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) WithRoleID(roleID strfmt.UUID) *PostRolesRoleIDAcesParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the post roles role ID aces params
func (o *PostRolesRoleIDAcesParams) SetRoleID(roleID strfmt.UUID) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *PostRolesRoleIDAcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AceCreationRequest != nil {
		if err := r.SetBodyParam(o.AceCreationRequest); err != nil {
			return err
		}
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
