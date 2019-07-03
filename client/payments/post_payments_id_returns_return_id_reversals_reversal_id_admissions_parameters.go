// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams creates a new PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams object
// with the default values initialized.
func NewPostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams() *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParamsWithTimeout creates a new PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParamsWithContext creates a new PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParamsWithContext(ctx context.Context) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParamsWithHTTPClient creates a new PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams contains all the parameters to send to the API endpoint
for the post payments ID returns return ID reversals reversal ID admissions operation typically these are written to a http.Request
*/
type PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams struct {

	/*ReturnReversalAdmissionCreationRequest*/
	ReturnReversalAdmissionCreationRequest *models.ReturnReversalAdmissionCreation
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*ReturnID
	  Return Id

	*/
	ReturnID strfmt.UUID
	/*ReversalID
	  Reversal Id

	*/
	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) WithContext(ctx context.Context) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnReversalAdmissionCreationRequest adds the returnReversalAdmissionCreationRequest to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) WithReturnReversalAdmissionCreationRequest(returnReversalAdmissionCreationRequest *models.ReturnReversalAdmissionCreation) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	o.SetReturnReversalAdmissionCreationRequest(returnReversalAdmissionCreationRequest)
	return o
}

// SetReturnReversalAdmissionCreationRequest adds the returnReversalAdmissionCreationRequest to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) SetReturnReversalAdmissionCreationRequest(returnReversalAdmissionCreationRequest *models.ReturnReversalAdmissionCreation) {
	o.ReturnReversalAdmissionCreationRequest = returnReversalAdmissionCreationRequest
}

// WithID adds the id to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) WithID(id strfmt.UUID) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithReturnID adds the returnID to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) WithReturnID(returnID strfmt.UUID) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	o.SetReturnID(returnID)
	return o
}

// SetReturnID adds the returnId to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) SetReturnID(returnID strfmt.UUID) {
	o.ReturnID = returnID
}

// WithReversalID adds the reversalID to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) WithReversalID(reversalID strfmt.UUID) *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams {
	o.SetReversalID(reversalID)
	return o
}

// SetReversalID adds the reversalId to the post payments ID returns return ID reversals reversal ID admissions params
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) SetReversalID(reversalID strfmt.UUID) {
	o.ReversalID = reversalID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDReturnsReturnIDReversalsReversalIDAdmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReturnReversalAdmissionCreationRequest != nil {
		if err := r.SetBodyParam(o.ReturnReversalAdmissionCreationRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
