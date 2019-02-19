// Code generated by go-swagger; DO NOT EDIT.

package organisations

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
)

// NewGetUnitsParams creates a new GetUnitsParams object
// with the default values initialized.
func NewGetUnitsParams() *GetUnitsParams {
	var ()
	return &GetUnitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUnitsParamsWithTimeout creates a new GetUnitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUnitsParamsWithTimeout(timeout time.Duration) *GetUnitsParams {
	var ()
	return &GetUnitsParams{

		timeout: timeout,
	}
}

// NewGetUnitsParamsWithContext creates a new GetUnitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUnitsParamsWithContext(ctx context.Context) *GetUnitsParams {
	var ()
	return &GetUnitsParams{

		Context: ctx,
	}
}

// NewGetUnitsParamsWithHTTPClient creates a new GetUnitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUnitsParamsWithHTTPClient(client *http.Client) *GetUnitsParams {
	var ()
	return &GetUnitsParams{
		HTTPClient: client,
	}
}

/*GetUnitsParams contains all the parameters to send to the API endpoint
for the get units operation typically these are written to a http.Request
*/
type GetUnitsParams struct {

	/*FilterChildOrganisationID
	  Child org id

	*/
	FilterChildOrganisationID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get units params
func (o *GetUnitsParams) WithTimeout(timeout time.Duration) *GetUnitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get units params
func (o *GetUnitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get units params
func (o *GetUnitsParams) WithContext(ctx context.Context) *GetUnitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get units params
func (o *GetUnitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get units params
func (o *GetUnitsParams) WithHTTPClient(client *http.Client) *GetUnitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get units params
func (o *GetUnitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterChildOrganisationID adds the filterChildOrganisationID to the get units params
func (o *GetUnitsParams) WithFilterChildOrganisationID(filterChildOrganisationID *strfmt.UUID) *GetUnitsParams {
	o.SetFilterChildOrganisationID(filterChildOrganisationID)
	return o
}

// SetFilterChildOrganisationID adds the filterChildOrganisationId to the get units params
func (o *GetUnitsParams) SetFilterChildOrganisationID(filterChildOrganisationID *strfmt.UUID) {
	o.FilterChildOrganisationID = filterChildOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUnitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterChildOrganisationID != nil {

		// query param filter[child_organisation_id]
		var qrFilterChildOrganisationID strfmt.UUID
		if o.FilterChildOrganisationID != nil {
			qrFilterChildOrganisationID = *o.FilterChildOrganisationID
		}
		qFilterChildOrganisationID := qrFilterChildOrganisationID.String()
		if qFilterChildOrganisationID != "" {
			if err := r.SetQueryParam("filter[child_organisation_id]", qFilterChildOrganisationID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
