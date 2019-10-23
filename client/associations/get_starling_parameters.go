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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetStarlingParams creates a new GetStarlingParams object
// with the default values initialized.
func NewGetStarlingParams() *GetStarlingParams {
	var ()
	return &GetStarlingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStarlingParamsWithTimeout creates a new GetStarlingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStarlingParamsWithTimeout(timeout time.Duration) *GetStarlingParams {
	var ()
	return &GetStarlingParams{

		timeout: timeout,
	}
}

// NewGetStarlingParamsWithContext creates a new GetStarlingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStarlingParamsWithContext(ctx context.Context) *GetStarlingParams {
	var ()
	return &GetStarlingParams{

		Context: ctx,
	}
}

// NewGetStarlingParamsWithHTTPClient creates a new GetStarlingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStarlingParamsWithHTTPClient(client *http.Client) *GetStarlingParams {
	var ()
	return &GetStarlingParams{
		HTTPClient: client,
	}
}

/*GetStarlingParams contains all the parameters to send to the API endpoint
for the get starling operation typically these are written to a http.Request
*/
type GetStarlingParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get starling params
func (o *GetStarlingParams) WithTimeout(timeout time.Duration) *GetStarlingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get starling params
func (o *GetStarlingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get starling params
func (o *GetStarlingParams) WithContext(ctx context.Context) *GetStarlingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get starling params
func (o *GetStarlingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get starling params
func (o *GetStarlingParams) WithHTTPClient(client *http.Client) *GetStarlingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get starling params
func (o *GetStarlingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get starling params
func (o *GetStarlingParams) WithFilterOrganisationID(filterOrganisationID *strfmt.UUID) *GetStarlingParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get starling params
func (o *GetStarlingParams) SetFilterOrganisationID(filterOrganisationID *strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStarlingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterOrganisationID != nil {

		// query param filter[organisation_id]
		var qrFilterOrganisationID strfmt.UUID
		if o.FilterOrganisationID != nil {
			qrFilterOrganisationID = *o.FilterOrganisationID
		}
		qFilterOrganisationID := qrFilterOrganisationID.String()
		if qFilterOrganisationID != "" {
			if err := r.SetQueryParam("filter[organisation_id]", qFilterOrganisationID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
