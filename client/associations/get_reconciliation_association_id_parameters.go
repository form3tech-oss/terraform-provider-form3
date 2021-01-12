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

// NewGetReconciliationAssociationIDParams creates a new GetReconciliationAssociationIDParams object
// with the default values initialized.
func NewGetReconciliationAssociationIDParams() *GetReconciliationAssociationIDParams {
	var ()
	return &GetReconciliationAssociationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReconciliationAssociationIDParamsWithTimeout creates a new GetReconciliationAssociationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReconciliationAssociationIDParamsWithTimeout(timeout time.Duration) *GetReconciliationAssociationIDParams {
	var ()
	return &GetReconciliationAssociationIDParams{

		timeout: timeout,
	}
}

// NewGetReconciliationAssociationIDParamsWithContext creates a new GetReconciliationAssociationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReconciliationAssociationIDParamsWithContext(ctx context.Context) *GetReconciliationAssociationIDParams {
	var ()
	return &GetReconciliationAssociationIDParams{

		Context: ctx,
	}
}

// NewGetReconciliationAssociationIDParamsWithHTTPClient creates a new GetReconciliationAssociationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReconciliationAssociationIDParamsWithHTTPClient(client *http.Client) *GetReconciliationAssociationIDParams {
	var ()
	return &GetReconciliationAssociationIDParams{
		HTTPClient: client,
	}
}

/*GetReconciliationAssociationIDParams contains all the parameters to send to the API endpoint
for the get reconciliation association ID operation typically these are written to a http.Request
*/
type GetReconciliationAssociationIDParams struct {

	/*AssociationID
	  Association id

	*/
	AssociationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get reconciliation association ID params
func (o *GetReconciliationAssociationIDParams) WithTimeout(timeout time.Duration) *GetReconciliationAssociationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reconciliation association ID params
func (o *GetReconciliationAssociationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reconciliation association ID params
func (o *GetReconciliationAssociationIDParams) WithContext(ctx context.Context) *GetReconciliationAssociationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reconciliation association ID params
func (o *GetReconciliationAssociationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reconciliation association ID params
func (o *GetReconciliationAssociationIDParams) WithHTTPClient(client *http.Client) *GetReconciliationAssociationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reconciliation association ID params
func (o *GetReconciliationAssociationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the get reconciliation association ID params
func (o *GetReconciliationAssociationIDParams) WithAssociationID(associationID strfmt.UUID) *GetReconciliationAssociationIDParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the get reconciliation association ID params
func (o *GetReconciliationAssociationIDParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReconciliationAssociationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associationId
	if err := r.SetPathParam("associationId", o.AssociationID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}