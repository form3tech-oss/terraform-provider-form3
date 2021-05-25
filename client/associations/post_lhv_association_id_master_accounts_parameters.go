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

// NewPostLhvAssociationIDMasterAccountsParams creates a new PostLhvAssociationIDMasterAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLhvAssociationIDMasterAccountsParams() *PostLhvAssociationIDMasterAccountsParams {
	return &PostLhvAssociationIDMasterAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLhvAssociationIDMasterAccountsParamsWithTimeout creates a new PostLhvAssociationIDMasterAccountsParams object
// with the ability to set a timeout on a request.
func NewPostLhvAssociationIDMasterAccountsParamsWithTimeout(timeout time.Duration) *PostLhvAssociationIDMasterAccountsParams {
	return &PostLhvAssociationIDMasterAccountsParams{
		timeout: timeout,
	}
}

// NewPostLhvAssociationIDMasterAccountsParamsWithContext creates a new PostLhvAssociationIDMasterAccountsParams object
// with the ability to set a context for a request.
func NewPostLhvAssociationIDMasterAccountsParamsWithContext(ctx context.Context) *PostLhvAssociationIDMasterAccountsParams {
	return &PostLhvAssociationIDMasterAccountsParams{
		Context: ctx,
	}
}

// NewPostLhvAssociationIDMasterAccountsParamsWithHTTPClient creates a new PostLhvAssociationIDMasterAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLhvAssociationIDMasterAccountsParamsWithHTTPClient(client *http.Client) *PostLhvAssociationIDMasterAccountsParams {
	return &PostLhvAssociationIDMasterAccountsParams{
		HTTPClient: client,
	}
}

/* PostLhvAssociationIDMasterAccountsParams contains all the parameters to send to the API endpoint
   for the post lhv association ID master accounts operation.

   Typically these are written to a http.Request.
*/
type PostLhvAssociationIDMasterAccountsParams struct {

	/* AssociationID.

	   Association Id

	   Format: uuid
	*/
	AssociationID strfmt.UUID

	// CreationRequest.
	CreationRequest *models.LhvMasterAccountCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post lhv association ID master accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLhvAssociationIDMasterAccountsParams) WithDefaults() *PostLhvAssociationIDMasterAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post lhv association ID master accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLhvAssociationIDMasterAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) WithTimeout(timeout time.Duration) *PostLhvAssociationIDMasterAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) WithContext(ctx context.Context) *PostLhvAssociationIDMasterAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) WithHTTPClient(client *http.Client) *PostLhvAssociationIDMasterAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) WithAssociationID(associationID strfmt.UUID) *PostLhvAssociationIDMasterAccountsParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithCreationRequest adds the creationRequest to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) WithCreationRequest(creationRequest *models.LhvMasterAccountCreation) *PostLhvAssociationIDMasterAccountsParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post lhv association ID master accounts params
func (o *PostLhvAssociationIDMasterAccountsParams) SetCreationRequest(creationRequest *models.LhvMasterAccountCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostLhvAssociationIDMasterAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associationId
	if err := r.SetPathParam("associationId", o.AssociationID.String()); err != nil {
		return err
	}
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
