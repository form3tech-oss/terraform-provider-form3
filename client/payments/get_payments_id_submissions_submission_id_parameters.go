// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// NewGetPaymentsIDSubmissionsSubmissionIDParams creates a new GetPaymentsIDSubmissionsSubmissionIDParams object
// with the default values initialized.
func NewGetPaymentsIDSubmissionsSubmissionIDParams() *GetPaymentsIDSubmissionsSubmissionIDParams {
	var ()
	return &GetPaymentsIDSubmissionsSubmissionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsIDSubmissionsSubmissionIDParamsWithTimeout creates a new GetPaymentsIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsIDSubmissionsSubmissionIDParamsWithTimeout(timeout time.Duration) *GetPaymentsIDSubmissionsSubmissionIDParams {
	var ()
	return &GetPaymentsIDSubmissionsSubmissionIDParams{

		timeout: timeout,
	}
}

// NewGetPaymentsIDSubmissionsSubmissionIDParamsWithContext creates a new GetPaymentsIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsIDSubmissionsSubmissionIDParamsWithContext(ctx context.Context) *GetPaymentsIDSubmissionsSubmissionIDParams {
	var ()
	return &GetPaymentsIDSubmissionsSubmissionIDParams{

		Context: ctx,
	}
}

// NewGetPaymentsIDSubmissionsSubmissionIDParamsWithHTTPClient creates a new GetPaymentsIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsIDSubmissionsSubmissionIDParamsWithHTTPClient(client *http.Client) *GetPaymentsIDSubmissionsSubmissionIDParams {
	var ()
	return &GetPaymentsIDSubmissionsSubmissionIDParams{
		HTTPClient: client,
	}
}

/*GetPaymentsIDSubmissionsSubmissionIDParams contains all the parameters to send to the API endpoint
for the get payments ID submissions submission ID operation typically these are written to a http.Request
*/
type GetPaymentsIDSubmissionsSubmissionIDParams struct {

	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*SubmissionID
	  Submission Id

	*/
	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) WithTimeout(timeout time.Duration) *GetPaymentsIDSubmissionsSubmissionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) WithContext(ctx context.Context) *GetPaymentsIDSubmissionsSubmissionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) WithHTTPClient(client *http.Client) *GetPaymentsIDSubmissionsSubmissionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) WithID(id strfmt.UUID) *GetPaymentsIDSubmissionsSubmissionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithSubmissionID adds the submissionID to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) WithSubmissionID(submissionID strfmt.UUID) *GetPaymentsIDSubmissionsSubmissionIDParams {
	o.SetSubmissionID(submissionID)
	return o
}

// SetSubmissionID adds the submissionId to the get payments ID submissions submission ID params
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) SetSubmissionID(submissionID strfmt.UUID) {
	o.SubmissionID = submissionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsIDSubmissionsSubmissionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
