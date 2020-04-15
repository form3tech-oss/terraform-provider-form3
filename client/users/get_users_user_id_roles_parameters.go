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

// NewGetUsersUserIDRolesParams creates a new GetUsersUserIDRolesParams object
// with the default values initialized.
func NewGetUsersUserIDRolesParams() *GetUsersUserIDRolesParams {
	var ()
	return &GetUsersUserIDRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUserIDRolesParamsWithTimeout creates a new GetUsersUserIDRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersUserIDRolesParamsWithTimeout(timeout time.Duration) *GetUsersUserIDRolesParams {
	var ()
	return &GetUsersUserIDRolesParams{

		timeout: timeout,
	}
}

// NewGetUsersUserIDRolesParamsWithContext creates a new GetUsersUserIDRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersUserIDRolesParamsWithContext(ctx context.Context) *GetUsersUserIDRolesParams {
	var ()
	return &GetUsersUserIDRolesParams{

		Context: ctx,
	}
}

// NewGetUsersUserIDRolesParamsWithHTTPClient creates a new GetUsersUserIDRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersUserIDRolesParamsWithHTTPClient(client *http.Client) *GetUsersUserIDRolesParams {
	var ()
	return &GetUsersUserIDRolesParams{
		HTTPClient: client,
	}
}

/*GetUsersUserIDRolesParams contains all the parameters to send to the API endpoint
for the get users user ID roles operation typically these are written to a http.Request
*/
type GetUsersUserIDRolesParams struct {

	/*UserID
	  User Id

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users user ID roles params
func (o *GetUsersUserIDRolesParams) WithTimeout(timeout time.Duration) *GetUsersUserIDRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users user ID roles params
func (o *GetUsersUserIDRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users user ID roles params
func (o *GetUsersUserIDRolesParams) WithContext(ctx context.Context) *GetUsersUserIDRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users user ID roles params
func (o *GetUsersUserIDRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users user ID roles params
func (o *GetUsersUserIDRolesParams) WithHTTPClient(client *http.Client) *GetUsersUserIDRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users user ID roles params
func (o *GetUsersUserIDRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get users user ID roles params
func (o *GetUsersUserIDRolesParams) WithUserID(userID strfmt.UUID) *GetUsersUserIDRolesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get users user ID roles params
func (o *GetUsersUserIDRolesParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUserIDRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
