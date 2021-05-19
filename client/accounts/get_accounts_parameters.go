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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAccountsParams creates a new GetAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountsParams() *GetAccountsParams {
	return &GetAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsParamsWithTimeout creates a new GetAccountsParams object
// with the ability to set a timeout on a request.
func NewGetAccountsParamsWithTimeout(timeout time.Duration) *GetAccountsParams {
	return &GetAccountsParams{
		timeout: timeout,
	}
}

// NewGetAccountsParamsWithContext creates a new GetAccountsParams object
// with the ability to set a context for a request.
func NewGetAccountsParamsWithContext(ctx context.Context) *GetAccountsParams {
	return &GetAccountsParams{
		Context: ctx,
	}
}

// NewGetAccountsParamsWithHTTPClient creates a new GetAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountsParamsWithHTTPClient(client *http.Client) *GetAccountsParams {
	return &GetAccountsParams{
		HTTPClient: client,
	}
}

/* GetAccountsParams contains all the parameters to send to the API endpoint
   for the get accounts operation.

   Typically these are written to a http.Request.
*/
type GetAccountsParams struct {

	/* FilterAccountNumber.

	   Filter by account number
	*/
	FilterAccountNumber []string

	/* FilterBankID.

	   Filter by bank id e.g. sort code or bic
	*/
	FilterBankID []string

	/* FilterBankIDCode.

	   Filter by type of bank id e.g. "GBDSC"
	*/
	FilterBankIDCode []string

	/* FilterCountry.

	   Filter by country e.g. FR,GB
	*/
	FilterCountry []string

	/* FilterCustomerID.

	   Filter by customer_id
	*/
	FilterCustomerID []string

	/* FilterIban.

	   Filter by IBAN
	*/
	FilterIban []string

	/* FilterOrganisationID.

	   Filter by organisation id
	*/
	FilterOrganisationID []strfmt.UUID

	/* PageNumber.

	   Which page to select
	*/
	PageNumber *int64

	/* PageSize.

	   Number of items to select
	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsParams) WithDefaults() *GetAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) WithTimeout(timeout time.Duration) *GetAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts params
func (o *GetAccountsParams) WithContext(ctx context.Context) *GetAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts params
func (o *GetAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts params
func (o *GetAccountsParams) WithHTTPClient(client *http.Client) *GetAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts params
func (o *GetAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterAccountNumber adds the filterAccountNumber to the get accounts params
func (o *GetAccountsParams) WithFilterAccountNumber(filterAccountNumber []string) *GetAccountsParams {
	o.SetFilterAccountNumber(filterAccountNumber)
	return o
}

// SetFilterAccountNumber adds the filterAccountNumber to the get accounts params
func (o *GetAccountsParams) SetFilterAccountNumber(filterAccountNumber []string) {
	o.FilterAccountNumber = filterAccountNumber
}

// WithFilterBankID adds the filterBankID to the get accounts params
func (o *GetAccountsParams) WithFilterBankID(filterBankID []string) *GetAccountsParams {
	o.SetFilterBankID(filterBankID)
	return o
}

// SetFilterBankID adds the filterBankId to the get accounts params
func (o *GetAccountsParams) SetFilterBankID(filterBankID []string) {
	o.FilterBankID = filterBankID
}

// WithFilterBankIDCode adds the filterBankIDCode to the get accounts params
func (o *GetAccountsParams) WithFilterBankIDCode(filterBankIDCode []string) *GetAccountsParams {
	o.SetFilterBankIDCode(filterBankIDCode)
	return o
}

// SetFilterBankIDCode adds the filterBankIdCode to the get accounts params
func (o *GetAccountsParams) SetFilterBankIDCode(filterBankIDCode []string) {
	o.FilterBankIDCode = filterBankIDCode
}

// WithFilterCountry adds the filterCountry to the get accounts params
func (o *GetAccountsParams) WithFilterCountry(filterCountry []string) *GetAccountsParams {
	o.SetFilterCountry(filterCountry)
	return o
}

// SetFilterCountry adds the filterCountry to the get accounts params
func (o *GetAccountsParams) SetFilterCountry(filterCountry []string) {
	o.FilterCountry = filterCountry
}

// WithFilterCustomerID adds the filterCustomerID to the get accounts params
func (o *GetAccountsParams) WithFilterCustomerID(filterCustomerID []string) *GetAccountsParams {
	o.SetFilterCustomerID(filterCustomerID)
	return o
}

// SetFilterCustomerID adds the filterCustomerId to the get accounts params
func (o *GetAccountsParams) SetFilterCustomerID(filterCustomerID []string) {
	o.FilterCustomerID = filterCustomerID
}

// WithFilterIban adds the filterIban to the get accounts params
func (o *GetAccountsParams) WithFilterIban(filterIban []string) *GetAccountsParams {
	o.SetFilterIban(filterIban)
	return o
}

// SetFilterIban adds the filterIban to the get accounts params
func (o *GetAccountsParams) SetFilterIban(filterIban []string) {
	o.FilterIban = filterIban
}

// WithFilterOrganisationID adds the filterOrganisationID to the get accounts params
func (o *GetAccountsParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetAccountsParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get accounts params
func (o *GetAccountsParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WithPageNumber adds the pageNumber to the get accounts params
func (o *GetAccountsParams) WithPageNumber(pageNumber *int64) *GetAccountsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get accounts params
func (o *GetAccountsParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get accounts params
func (o *GetAccountsParams) WithPageSize(pageSize *int64) *GetAccountsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get accounts params
func (o *GetAccountsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterAccountNumber != nil {

		// binding items for filter[account_number]
		joinedFilterAccountNumber := o.bindParamFilterAccountNumber(reg)

		// query array param filter[account_number]
		if err := r.SetQueryParam("filter[account_number]", joinedFilterAccountNumber...); err != nil {
			return err
		}
	}

	if o.FilterBankID != nil {

		// binding items for filter[bank_id]
		joinedFilterBankID := o.bindParamFilterBankID(reg)

		// query array param filter[bank_id]
		if err := r.SetQueryParam("filter[bank_id]", joinedFilterBankID...); err != nil {
			return err
		}
	}

	if o.FilterBankIDCode != nil {

		// binding items for filter[bank_id_code]
		joinedFilterBankIDCode := o.bindParamFilterBankIDCode(reg)

		// query array param filter[bank_id_code]
		if err := r.SetQueryParam("filter[bank_id_code]", joinedFilterBankIDCode...); err != nil {
			return err
		}
	}

	if o.FilterCountry != nil {

		// binding items for filter[country]
		joinedFilterCountry := o.bindParamFilterCountry(reg)

		// query array param filter[country]
		if err := r.SetQueryParam("filter[country]", joinedFilterCountry...); err != nil {
			return err
		}
	}

	if o.FilterCustomerID != nil {

		// binding items for filter[customer_id]
		joinedFilterCustomerID := o.bindParamFilterCustomerID(reg)

		// query array param filter[customer_id]
		if err := r.SetQueryParam("filter[customer_id]", joinedFilterCustomerID...); err != nil {
			return err
		}
	}

	if o.FilterIban != nil {

		// binding items for filter[iban]
		joinedFilterIban := o.bindParamFilterIban(reg)

		// query array param filter[iban]
		if err := r.SetQueryParam("filter[iban]", joinedFilterIban...); err != nil {
			return err
		}
	}

	if o.FilterOrganisationID != nil {

		// binding items for filter[organisation_id]
		joinedFilterOrganisationID := o.bindParamFilterOrganisationID(reg)

		// query array param filter[organisation_id]
		if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
			return err
		}
	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber int64

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt64(qrPageNumber)
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

// bindParamGetAccounts binds the parameter filter[account_number]
func (o *GetAccountsParams) bindParamFilterAccountNumber(formats strfmt.Registry) []string {
	filterAccountNumberIR := o.FilterAccountNumber

	var filterAccountNumberIC []string
	for _, filterAccountNumberIIR := range filterAccountNumberIR { // explode []string

		filterAccountNumberIIV := filterAccountNumberIIR // string as string
		filterAccountNumberIC = append(filterAccountNumberIC, filterAccountNumberIIV)
	}

	// items.CollectionFormat: "csv"
	filterAccountNumberIS := swag.JoinByFormat(filterAccountNumberIC, "csv")

	return filterAccountNumberIS
}

// bindParamGetAccounts binds the parameter filter[bank_id]
func (o *GetAccountsParams) bindParamFilterBankID(formats strfmt.Registry) []string {
	filterBankIDIR := o.FilterBankID

	var filterBankIDIC []string
	for _, filterBankIDIIR := range filterBankIDIR { // explode []string

		filterBankIDIIV := filterBankIDIIR // string as string
		filterBankIDIC = append(filterBankIDIC, filterBankIDIIV)
	}

	// items.CollectionFormat: "csv"
	filterBankIDIS := swag.JoinByFormat(filterBankIDIC, "csv")

	return filterBankIDIS
}

// bindParamGetAccounts binds the parameter filter[bank_id_code]
func (o *GetAccountsParams) bindParamFilterBankIDCode(formats strfmt.Registry) []string {
	filterBankIDCodeIR := o.FilterBankIDCode

	var filterBankIDCodeIC []string
	for _, filterBankIDCodeIIR := range filterBankIDCodeIR { // explode []string

		filterBankIDCodeIIV := filterBankIDCodeIIR // string as string
		filterBankIDCodeIC = append(filterBankIDCodeIC, filterBankIDCodeIIV)
	}

	// items.CollectionFormat: "csv"
	filterBankIDCodeIS := swag.JoinByFormat(filterBankIDCodeIC, "csv")

	return filterBankIDCodeIS
}

// bindParamGetAccounts binds the parameter filter[country]
func (o *GetAccountsParams) bindParamFilterCountry(formats strfmt.Registry) []string {
	filterCountryIR := o.FilterCountry

	var filterCountryIC []string
	for _, filterCountryIIR := range filterCountryIR { // explode []string

		filterCountryIIV := filterCountryIIR // string as string
		filterCountryIC = append(filterCountryIC, filterCountryIIV)
	}

	// items.CollectionFormat: "csv"
	filterCountryIS := swag.JoinByFormat(filterCountryIC, "csv")

	return filterCountryIS
}

// bindParamGetAccounts binds the parameter filter[customer_id]
func (o *GetAccountsParams) bindParamFilterCustomerID(formats strfmt.Registry) []string {
	filterCustomerIDIR := o.FilterCustomerID

	var filterCustomerIDIC []string
	for _, filterCustomerIDIIR := range filterCustomerIDIR { // explode []string

		filterCustomerIDIIV := filterCustomerIDIIR // string as string
		filterCustomerIDIC = append(filterCustomerIDIC, filterCustomerIDIIV)
	}

	// items.CollectionFormat: "csv"
	filterCustomerIDIS := swag.JoinByFormat(filterCustomerIDIC, "csv")

	return filterCustomerIDIS
}

// bindParamGetAccounts binds the parameter filter[iban]
func (o *GetAccountsParams) bindParamFilterIban(formats strfmt.Registry) []string {
	filterIbanIR := o.FilterIban

	var filterIbanIC []string
	for _, filterIbanIIR := range filterIbanIR { // explode []string

		filterIbanIIV := filterIbanIIR // string as string
		filterIbanIC = append(filterIbanIC, filterIbanIIV)
	}

	// items.CollectionFormat: "csv"
	filterIbanIS := swag.JoinByFormat(filterIbanIC, "csv")

	return filterIbanIS
}

// bindParamGetAccounts binds the parameter filter[organisation_id]
func (o *GetAccountsParams) bindParamFilterOrganisationID(formats strfmt.Registry) []string {
	filterOrganisationIDIR := o.FilterOrganisationID

	var filterOrganisationIDIC []string
	for _, filterOrganisationIDIIR := range filterOrganisationIDIR { // explode []strfmt.UUID

		filterOrganisationIDIIV := filterOrganisationIDIIR.String() // strfmt.UUID as string
		filterOrganisationIDIC = append(filterOrganisationIDIC, filterOrganisationIDIIV)
	}

	// items.CollectionFormat: "csv"
	filterOrganisationIDIS := swag.JoinByFormat(filterOrganisationIDIC, "csv")

	return filterOrganisationIDIS
}
