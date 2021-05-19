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

// NewGetUsersUserIDCredentialsParams creates a new GetUsersUserIDCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersUserIDCredentialsParams() *GetUsersUserIDCredentialsParams {
	return &GetUsersUserIDCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUserIDCredentialsParamsWithTimeout creates a new GetUsersUserIDCredentialsParams object
// with the ability to set a timeout on a request.
func NewGetUsersUserIDCredentialsParamsWithTimeout(timeout time.Duration) *GetUsersUserIDCredentialsParams {
	return &GetUsersUserIDCredentialsParams{
		timeout: timeout,
	}
}

// NewGetUsersUserIDCredentialsParamsWithContext creates a new GetUsersUserIDCredentialsParams object
// with the ability to set a context for a request.
func NewGetUsersUserIDCredentialsParamsWithContext(ctx context.Context) *GetUsersUserIDCredentialsParams {
	return &GetUsersUserIDCredentialsParams{
		Context: ctx,
	}
}

// NewGetUsersUserIDCredentialsParamsWithHTTPClient creates a new GetUsersUserIDCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersUserIDCredentialsParamsWithHTTPClient(client *http.Client) *GetUsersUserIDCredentialsParams {
	return &GetUsersUserIDCredentialsParams{
		HTTPClient: client,
	}
}

/* GetUsersUserIDCredentialsParams contains all the parameters to send to the API endpoint
   for the get users user ID credentials operation.

   Typically these are written to a http.Request.
*/
type GetUsersUserIDCredentialsParams struct {

	/* UserID.

	   User Id

	   Format: uuid
	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users user ID credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersUserIDCredentialsParams) WithDefaults() *GetUsersUserIDCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users user ID credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersUserIDCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get users user ID credentials params
func (o *GetUsersUserIDCredentialsParams) WithTimeout(timeout time.Duration) *GetUsersUserIDCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users user ID credentials params
func (o *GetUsersUserIDCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users user ID credentials params
func (o *GetUsersUserIDCredentialsParams) WithContext(ctx context.Context) *GetUsersUserIDCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users user ID credentials params
func (o *GetUsersUserIDCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users user ID credentials params
func (o *GetUsersUserIDCredentialsParams) WithHTTPClient(client *http.Client) *GetUsersUserIDCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users user ID credentials params
func (o *GetUsersUserIDCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get users user ID credentials params
func (o *GetUsersUserIDCredentialsParams) WithUserID(userID strfmt.UUID) *GetUsersUserIDCredentialsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get users user ID credentials params
func (o *GetUsersUserIDCredentialsParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUserIDCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
