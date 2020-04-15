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
	"github.com/go-openapi/strfmt"
)

// NewDeleteUsersUserIDRolesRoleIDParams creates a new DeleteUsersUserIDRolesRoleIDParams object
// with the default values initialized.
func NewDeleteUsersUserIDRolesRoleIDParams() *DeleteUsersUserIDRolesRoleIDParams {
	var ()
	return &DeleteUsersUserIDRolesRoleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsersUserIDRolesRoleIDParamsWithTimeout creates a new DeleteUsersUserIDRolesRoleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUsersUserIDRolesRoleIDParamsWithTimeout(timeout time.Duration) *DeleteUsersUserIDRolesRoleIDParams {
	var ()
	return &DeleteUsersUserIDRolesRoleIDParams{

		timeout: timeout,
	}
}

// NewDeleteUsersUserIDRolesRoleIDParamsWithContext creates a new DeleteUsersUserIDRolesRoleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUsersUserIDRolesRoleIDParamsWithContext(ctx context.Context) *DeleteUsersUserIDRolesRoleIDParams {
	var ()
	return &DeleteUsersUserIDRolesRoleIDParams{

		Context: ctx,
	}
}

// NewDeleteUsersUserIDRolesRoleIDParamsWithHTTPClient creates a new DeleteUsersUserIDRolesRoleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUsersUserIDRolesRoleIDParamsWithHTTPClient(client *http.Client) *DeleteUsersUserIDRolesRoleIDParams {
	var ()
	return &DeleteUsersUserIDRolesRoleIDParams{
		HTTPClient: client,
	}
}

/*DeleteUsersUserIDRolesRoleIDParams contains all the parameters to send to the API endpoint
for the delete users user ID roles role ID operation typically these are written to a http.Request
*/
type DeleteUsersUserIDRolesRoleIDParams struct {

	/*RoleID
	  Role Id

	*/
	RoleID strfmt.UUID
	/*UserID
	  User Id

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) WithTimeout(timeout time.Duration) *DeleteUsersUserIDRolesRoleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) WithContext(ctx context.Context) *DeleteUsersUserIDRolesRoleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) WithHTTPClient(client *http.Client) *DeleteUsersUserIDRolesRoleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) WithRoleID(roleID strfmt.UUID) *DeleteUsersUserIDRolesRoleIDParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) SetRoleID(roleID strfmt.UUID) {
	o.RoleID = roleID
}

// WithUserID adds the userID to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) WithUserID(userID strfmt.UUID) *DeleteUsersUserIDRolesRoleIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete users user ID roles role ID params
func (o *DeleteUsersUserIDRolesRoleIDParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsersUserIDRolesRoleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID.String()); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
