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

// NewPatchSepactLiquidityAssociationIDParams creates a new PatchSepactLiquidityAssociationIDParams object
// with the default values initialized.
func NewPatchSepactLiquidityAssociationIDParams() *PatchSepactLiquidityAssociationIDParams {
	var ()
	return &PatchSepactLiquidityAssociationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSepactLiquidityAssociationIDParamsWithTimeout creates a new PatchSepactLiquidityAssociationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSepactLiquidityAssociationIDParamsWithTimeout(timeout time.Duration) *PatchSepactLiquidityAssociationIDParams {
	var ()
	return &PatchSepactLiquidityAssociationIDParams{

		timeout: timeout,
	}
}

// NewPatchSepactLiquidityAssociationIDParamsWithContext creates a new PatchSepactLiquidityAssociationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSepactLiquidityAssociationIDParamsWithContext(ctx context.Context) *PatchSepactLiquidityAssociationIDParams {
	var ()
	return &PatchSepactLiquidityAssociationIDParams{

		Context: ctx,
	}
}

// NewPatchSepactLiquidityAssociationIDParamsWithHTTPClient creates a new PatchSepactLiquidityAssociationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSepactLiquidityAssociationIDParamsWithHTTPClient(client *http.Client) *PatchSepactLiquidityAssociationIDParams {
	var ()
	return &PatchSepactLiquidityAssociationIDParams{
		HTTPClient: client,
	}
}

/*PatchSepactLiquidityAssociationIDParams contains all the parameters to send to the API endpoint
for the patch sepact liquidity association ID operation typically these are written to a http.Request
*/
type PatchSepactLiquidityAssociationIDParams struct {

	/*AssociationID
	  Association Id

	*/
	AssociationID strfmt.UUID
	/*UpdateRequest*/
	UpdateRequest *models.SepactLiquidityAssociationUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) WithTimeout(timeout time.Duration) *PatchSepactLiquidityAssociationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) WithContext(ctx context.Context) *PatchSepactLiquidityAssociationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) WithHTTPClient(client *http.Client) *PatchSepactLiquidityAssociationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) WithAssociationID(associationID strfmt.UUID) *PatchSepactLiquidityAssociationIDParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithUpdateRequest adds the updateRequest to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) WithUpdateRequest(updateRequest *models.SepactLiquidityAssociationUpdate) *PatchSepactLiquidityAssociationIDParams {
	o.SetUpdateRequest(updateRequest)
	return o
}

// SetUpdateRequest adds the updateRequest to the patch sepact liquidity association ID params
func (o *PatchSepactLiquidityAssociationIDParams) SetUpdateRequest(updateRequest *models.SepactLiquidityAssociationUpdate) {
	o.UpdateRequest = updateRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSepactLiquidityAssociationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associationId
	if err := r.SetPathParam("associationId", o.AssociationID.String()); err != nil {
		return err
	}

	if o.UpdateRequest != nil {
		if err := r.SetBodyParam(o.UpdateRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
