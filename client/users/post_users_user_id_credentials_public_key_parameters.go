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

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostUsersUserIDCredentialsPublicKeyParams creates a new PostUsersUserIDCredentialsPublicKeyParams object
// with the default values initialized.
func NewPostUsersUserIDCredentialsPublicKeyParams() *PostUsersUserIDCredentialsPublicKeyParams {
	var ()
	return &PostUsersUserIDCredentialsPublicKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsersUserIDCredentialsPublicKeyParamsWithTimeout creates a new PostUsersUserIDCredentialsPublicKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUsersUserIDCredentialsPublicKeyParamsWithTimeout(timeout time.Duration) *PostUsersUserIDCredentialsPublicKeyParams {
	var ()
	return &PostUsersUserIDCredentialsPublicKeyParams{

		timeout: timeout,
	}
}

// NewPostUsersUserIDCredentialsPublicKeyParamsWithContext creates a new PostUsersUserIDCredentialsPublicKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUsersUserIDCredentialsPublicKeyParamsWithContext(ctx context.Context) *PostUsersUserIDCredentialsPublicKeyParams {
	var ()
	return &PostUsersUserIDCredentialsPublicKeyParams{

		Context: ctx,
	}
}

// NewPostUsersUserIDCredentialsPublicKeyParamsWithHTTPClient creates a new PostUsersUserIDCredentialsPublicKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUsersUserIDCredentialsPublicKeyParamsWithHTTPClient(client *http.Client) *PostUsersUserIDCredentialsPublicKeyParams {
	var ()
	return &PostUsersUserIDCredentialsPublicKeyParams{
		HTTPClient: client,
	}
}

/*PostUsersUserIDCredentialsPublicKeyParams contains all the parameters to send to the API endpoint
for the post users user ID credentials public key operation typically these are written to a http.Request
*/
type PostUsersUserIDCredentialsPublicKeyParams struct {

	/*PublicKey
	  The public key to create.

	*/
	PublicKey *models.PublicKey
	/*UserID
	  User Id

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) WithTimeout(timeout time.Duration) *PostUsersUserIDCredentialsPublicKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) WithContext(ctx context.Context) *PostUsersUserIDCredentialsPublicKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) WithHTTPClient(client *http.Client) *PostUsersUserIDCredentialsPublicKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPublicKey adds the publicKey to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) WithPublicKey(publicKey *models.PublicKey) *PostUsersUserIDCredentialsPublicKeyParams {
	o.SetPublicKey(publicKey)
	return o
}

// SetPublicKey adds the publicKey to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) SetPublicKey(publicKey *models.PublicKey) {
	o.PublicKey = publicKey
}

// WithUserID adds the userID to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) WithUserID(userID strfmt.UUID) *PostUsersUserIDCredentialsPublicKeyParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post users user ID credentials public key params
func (o *PostUsersUserIDCredentialsPublicKeyParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsersUserIDCredentialsPublicKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PublicKey != nil {
		if err := r.SetBodyParam(o.PublicKey); err != nil {
			return err
		}
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
