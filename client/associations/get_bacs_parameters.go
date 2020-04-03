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

// NewGetBacsParams creates a new GetBacsParams object
// with the default values initialized.
func NewGetBacsParams() *GetBacsParams {
	var ()
	return &GetBacsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBacsParamsWithTimeout creates a new GetBacsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBacsParamsWithTimeout(timeout time.Duration) *GetBacsParams {
	var ()
	return &GetBacsParams{

		timeout: timeout,
	}
}

// NewGetBacsParamsWithContext creates a new GetBacsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBacsParamsWithContext(ctx context.Context) *GetBacsParams {
	var ()
	return &GetBacsParams{

		Context: ctx,
	}
}

// NewGetBacsParamsWithHTTPClient creates a new GetBacsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBacsParamsWithHTTPClient(client *http.Client) *GetBacsParams {
	var ()
	return &GetBacsParams{
		HTTPClient: client,
	}
}

/*GetBacsParams contains all the parameters to send to the API endpoint
for the get bacs operation typically these are written to a http.Request
*/
type GetBacsParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bacs params
func (o *GetBacsParams) WithTimeout(timeout time.Duration) *GetBacsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bacs params
func (o *GetBacsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bacs params
func (o *GetBacsParams) WithContext(ctx context.Context) *GetBacsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bacs params
func (o *GetBacsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bacs params
func (o *GetBacsParams) WithHTTPClient(client *http.Client) *GetBacsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bacs params
func (o *GetBacsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get bacs params
func (o *GetBacsParams) WithFilterOrganisationID(filterOrganisationID *strfmt.UUID) *GetBacsParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get bacs params
func (o *GetBacsParams) SetFilterOrganisationID(filterOrganisationID *strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBacsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
