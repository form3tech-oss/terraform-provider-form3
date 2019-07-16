// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBankidsParams creates a new GetBankidsParams object
// with the default values initialized.
func NewGetBankidsParams() *GetBankidsParams {
	var ()
	return &GetBankidsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBankidsParamsWithTimeout creates a new GetBankidsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBankidsParamsWithTimeout(timeout time.Duration) *GetBankidsParams {
	var ()
	return &GetBankidsParams{

		timeout: timeout,
	}
}

// NewGetBankidsParamsWithContext creates a new GetBankidsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBankidsParamsWithContext(ctx context.Context) *GetBankidsParams {
	var ()
	return &GetBankidsParams{

		Context: ctx,
	}
}

// NewGetBankidsParamsWithHTTPClient creates a new GetBankidsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBankidsParamsWithHTTPClient(client *http.Client) *GetBankidsParams {
	var ()
	return &GetBankidsParams{
		HTTPClient: client,
	}
}

/*GetBankidsParams contains all the parameters to send to the API endpoint
for the get bankids operation typically these are written to a http.Request
*/
type GetBankidsParams struct {

	/*FilterBankID
	  Filter by bank id e.g. sort code or bic

	*/
	FilterBankID []string
	/*FilterBankIDCode
	  Filter by type of bank id e.g. "GBDSC"

	*/
	FilterBankIDCode []string
	/*FilterCountry
	  Filter by country e.g. FR,GB

	*/
	FilterCountry []string
	/*FilterOrganisationID
	  Filter by organisation id

	*/
	FilterOrganisationID []strfmt.UUID
	/*PageNumber
	  Which page to select

	*/
	PageNumber *string
	/*PageSize
	  Number of items to select

	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bankids params
func (o *GetBankidsParams) WithTimeout(timeout time.Duration) *GetBankidsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bankids params
func (o *GetBankidsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bankids params
func (o *GetBankidsParams) WithContext(ctx context.Context) *GetBankidsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bankids params
func (o *GetBankidsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bankids params
func (o *GetBankidsParams) WithHTTPClient(client *http.Client) *GetBankidsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bankids params
func (o *GetBankidsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterBankID adds the filterBankID to the get bankids params
func (o *GetBankidsParams) WithFilterBankID(filterBankID []string) *GetBankidsParams {
	o.SetFilterBankID(filterBankID)
	return o
}

// SetFilterBankID adds the filterBankId to the get bankids params
func (o *GetBankidsParams) SetFilterBankID(filterBankID []string) {
	o.FilterBankID = filterBankID
}

// WithFilterBankIDCode adds the filterBankIDCode to the get bankids params
func (o *GetBankidsParams) WithFilterBankIDCode(filterBankIDCode []string) *GetBankidsParams {
	o.SetFilterBankIDCode(filterBankIDCode)
	return o
}

// SetFilterBankIDCode adds the filterBankIdCode to the get bankids params
func (o *GetBankidsParams) SetFilterBankIDCode(filterBankIDCode []string) {
	o.FilterBankIDCode = filterBankIDCode
}

// WithFilterCountry adds the filterCountry to the get bankids params
func (o *GetBankidsParams) WithFilterCountry(filterCountry []string) *GetBankidsParams {
	o.SetFilterCountry(filterCountry)
	return o
}

// SetFilterCountry adds the filterCountry to the get bankids params
func (o *GetBankidsParams) SetFilterCountry(filterCountry []string) {
	o.FilterCountry = filterCountry
}

// WithFilterOrganisationID adds the filterOrganisationID to the get bankids params
func (o *GetBankidsParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetBankidsParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get bankids params
func (o *GetBankidsParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WithPageNumber adds the pageNumber to the get bankids params
func (o *GetBankidsParams) WithPageNumber(pageNumber *string) *GetBankidsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get bankids params
func (o *GetBankidsParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get bankids params
func (o *GetBankidsParams) WithPageSize(pageSize *int64) *GetBankidsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get bankids params
func (o *GetBankidsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetBankidsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesFilterBankID := o.FilterBankID

	joinedFilterBankID := swag.JoinByFormat(valuesFilterBankID, "csv")
	// query array param filter[bank_id]
	if err := r.SetQueryParam("filter[bank_id]", joinedFilterBankID...); err != nil {
		return err
	}

	valuesFilterBankIDCode := o.FilterBankIDCode

	joinedFilterBankIDCode := swag.JoinByFormat(valuesFilterBankIDCode, "csv")
	// query array param filter[bank_id_code]
	if err := r.SetQueryParam("filter[bank_id_code]", joinedFilterBankIDCode...); err != nil {
		return err
	}

	valuesFilterCountry := o.FilterCountry

	joinedFilterCountry := swag.JoinByFormat(valuesFilterCountry, "csv")
	// query array param filter[country]
	if err := r.SetQueryParam("filter[country]", joinedFilterCountry...); err != nil {
		return err
	}

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "csv")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}