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
)

// NewGetLhvAssociationIDMasterAccountsMasterAccountIDParams creates a new GetLhvAssociationIDMasterAccountsMasterAccountIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLhvAssociationIDMasterAccountsMasterAccountIDParams() *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	return &GetLhvAssociationIDMasterAccountsMasterAccountIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLhvAssociationIDMasterAccountsMasterAccountIDParamsWithTimeout creates a new GetLhvAssociationIDMasterAccountsMasterAccountIDParams object
// with the ability to set a timeout on a request.
func NewGetLhvAssociationIDMasterAccountsMasterAccountIDParamsWithTimeout(timeout time.Duration) *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	return &GetLhvAssociationIDMasterAccountsMasterAccountIDParams{
		timeout: timeout,
	}
}

// NewGetLhvAssociationIDMasterAccountsMasterAccountIDParamsWithContext creates a new GetLhvAssociationIDMasterAccountsMasterAccountIDParams object
// with the ability to set a context for a request.
func NewGetLhvAssociationIDMasterAccountsMasterAccountIDParamsWithContext(ctx context.Context) *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	return &GetLhvAssociationIDMasterAccountsMasterAccountIDParams{
		Context: ctx,
	}
}

// NewGetLhvAssociationIDMasterAccountsMasterAccountIDParamsWithHTTPClient creates a new GetLhvAssociationIDMasterAccountsMasterAccountIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLhvAssociationIDMasterAccountsMasterAccountIDParamsWithHTTPClient(client *http.Client) *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	return &GetLhvAssociationIDMasterAccountsMasterAccountIDParams{
		HTTPClient: client,
	}
}

/* GetLhvAssociationIDMasterAccountsMasterAccountIDParams contains all the parameters to send to the API endpoint
   for the get lhv association ID master accounts master account ID operation.

   Typically these are written to a http.Request.
*/
type GetLhvAssociationIDMasterAccountsMasterAccountIDParams struct {

	/* AssociationID.

	   Association Id

	   Format: uuid
	*/
	AssociationID strfmt.UUID

	/* MasterAccountID.

	   Master Account Id

	   Format: uuid
	*/
	MasterAccountID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get lhv association ID master accounts master account ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) WithDefaults() *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get lhv association ID master accounts master account ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) WithTimeout(timeout time.Duration) *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) WithContext(ctx context.Context) *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) WithHTTPClient(client *http.Client) *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) WithAssociationID(associationID strfmt.UUID) *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithMasterAccountID adds the masterAccountID to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) WithMasterAccountID(masterAccountID strfmt.UUID) *GetLhvAssociationIDMasterAccountsMasterAccountIDParams {
	o.SetMasterAccountID(masterAccountID)
	return o
}

// SetMasterAccountID adds the masterAccountId to the get lhv association ID master accounts master account ID params
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) SetMasterAccountID(masterAccountID strfmt.UUID) {
	o.MasterAccountID = masterAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLhvAssociationIDMasterAccountsMasterAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associationId
	if err := r.SetPathParam("associationId", o.AssociationID.String()); err != nil {
		return err
	}

	// path param masterAccountId
	if err := r.SetPathParam("masterAccountId", o.MasterAccountID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
