// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

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

// NewGetDirectdebitsParams creates a new GetDirectdebitsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDirectdebitsParams() *GetDirectdebitsParams {
	return &GetDirectdebitsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDirectdebitsParamsWithTimeout creates a new GetDirectdebitsParams object
// with the ability to set a timeout on a request.
func NewGetDirectdebitsParamsWithTimeout(timeout time.Duration) *GetDirectdebitsParams {
	return &GetDirectdebitsParams{
		timeout: timeout,
	}
}

// NewGetDirectdebitsParamsWithContext creates a new GetDirectdebitsParams object
// with the ability to set a context for a request.
func NewGetDirectdebitsParamsWithContext(ctx context.Context) *GetDirectdebitsParams {
	return &GetDirectdebitsParams{
		Context: ctx,
	}
}

// NewGetDirectdebitsParamsWithHTTPClient creates a new GetDirectdebitsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDirectdebitsParamsWithHTTPClient(client *http.Client) *GetDirectdebitsParams {
	return &GetDirectdebitsParams{
		HTTPClient: client,
	}
}

/* GetDirectdebitsParams contains all the parameters to send to the API endpoint
   for the get directdebits operation.

   Typically these are written to a http.Request.
*/
type GetDirectdebitsParams struct {

	// FilterAdmissionAdmissionDateFrom.
	//
	// Format: date-time
	FilterAdmissionAdmissionDateFrom *strfmt.DateTime

	// FilterAdmissionAdmissionDateTo.
	//
	// Format: date-time
	FilterAdmissionAdmissionDateTo *strfmt.DateTime

	/* FilterAdmissionSchemeStatusCode.

	   Filter by admission scheme status code
	*/
	FilterAdmissionSchemeStatusCode *string

	/* FilterAdmissionStatus.

	   Filter by admission status
	*/
	FilterAdmissionStatus *string

	/* FilterAmount.

	   Filter by amount
	*/
	FilterAmount *string

	// FilterBeneficiaryPartyAccountNumber.
	FilterBeneficiaryPartyAccountNumber *string

	// FilterBeneficiaryPartyBankID.
	FilterBeneficiaryPartyBankID *string

	// FilterClearingID.
	FilterClearingID *string

	// FilterCurrency.
	FilterCurrency *string

	// FilterDebtorPartyAccountNumber.
	FilterDebtorPartyAccountNumber *string

	// FilterDebtorPartyBankID.
	FilterDebtorPartyBankID *string

	/* FilterOrganisationID.

	   Filter by organisation id
	*/
	FilterOrganisationID []strfmt.UUID

	// FilterPaymentScheme.
	FilterPaymentScheme *string

	// FilterPaymentType.
	FilterPaymentType *string

	// FilterProcessingDateFrom.
	//
	// Format: date
	FilterProcessingDateFrom *strfmt.Date

	// FilterProcessingDateTo.
	//
	// Format: date
	FilterProcessingDateTo *strfmt.Date

	/* FilterReference.

	   Filter by reference
	*/
	FilterReference *string

	/* FilterUniqueSchemeID.

	   Filter by unique scheme id
	*/
	FilterUniqueSchemeID *string

	/* PageNumber.

	   Which page to select
	*/
	PageNumber *string

	/* PageSize.

	   Number of items to select
	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get directdebits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDirectdebitsParams) WithDefaults() *GetDirectdebitsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get directdebits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDirectdebitsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get directdebits params
func (o *GetDirectdebitsParams) WithTimeout(timeout time.Duration) *GetDirectdebitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get directdebits params
func (o *GetDirectdebitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get directdebits params
func (o *GetDirectdebitsParams) WithContext(ctx context.Context) *GetDirectdebitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get directdebits params
func (o *GetDirectdebitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get directdebits params
func (o *GetDirectdebitsParams) WithHTTPClient(client *http.Client) *GetDirectdebitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get directdebits params
func (o *GetDirectdebitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterAdmissionAdmissionDateFrom adds the filterAdmissionAdmissionDateFrom to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterAdmissionAdmissionDateFrom(filterAdmissionAdmissionDateFrom *strfmt.DateTime) *GetDirectdebitsParams {
	o.SetFilterAdmissionAdmissionDateFrom(filterAdmissionAdmissionDateFrom)
	return o
}

// SetFilterAdmissionAdmissionDateFrom adds the filterAdmissionAdmissionDateFrom to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterAdmissionAdmissionDateFrom(filterAdmissionAdmissionDateFrom *strfmt.DateTime) {
	o.FilterAdmissionAdmissionDateFrom = filterAdmissionAdmissionDateFrom
}

// WithFilterAdmissionAdmissionDateTo adds the filterAdmissionAdmissionDateTo to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterAdmissionAdmissionDateTo(filterAdmissionAdmissionDateTo *strfmt.DateTime) *GetDirectdebitsParams {
	o.SetFilterAdmissionAdmissionDateTo(filterAdmissionAdmissionDateTo)
	return o
}

// SetFilterAdmissionAdmissionDateTo adds the filterAdmissionAdmissionDateTo to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterAdmissionAdmissionDateTo(filterAdmissionAdmissionDateTo *strfmt.DateTime) {
	o.FilterAdmissionAdmissionDateTo = filterAdmissionAdmissionDateTo
}

// WithFilterAdmissionSchemeStatusCode adds the filterAdmissionSchemeStatusCode to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterAdmissionSchemeStatusCode(filterAdmissionSchemeStatusCode *string) *GetDirectdebitsParams {
	o.SetFilterAdmissionSchemeStatusCode(filterAdmissionSchemeStatusCode)
	return o
}

// SetFilterAdmissionSchemeStatusCode adds the filterAdmissionSchemeStatusCode to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterAdmissionSchemeStatusCode(filterAdmissionSchemeStatusCode *string) {
	o.FilterAdmissionSchemeStatusCode = filterAdmissionSchemeStatusCode
}

// WithFilterAdmissionStatus adds the filterAdmissionStatus to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterAdmissionStatus(filterAdmissionStatus *string) *GetDirectdebitsParams {
	o.SetFilterAdmissionStatus(filterAdmissionStatus)
	return o
}

// SetFilterAdmissionStatus adds the filterAdmissionStatus to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterAdmissionStatus(filterAdmissionStatus *string) {
	o.FilterAdmissionStatus = filterAdmissionStatus
}

// WithFilterAmount adds the filterAmount to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterAmount(filterAmount *string) *GetDirectdebitsParams {
	o.SetFilterAmount(filterAmount)
	return o
}

// SetFilterAmount adds the filterAmount to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterAmount(filterAmount *string) {
	o.FilterAmount = filterAmount
}

// WithFilterBeneficiaryPartyAccountNumber adds the filterBeneficiaryPartyAccountNumber to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterBeneficiaryPartyAccountNumber(filterBeneficiaryPartyAccountNumber *string) *GetDirectdebitsParams {
	o.SetFilterBeneficiaryPartyAccountNumber(filterBeneficiaryPartyAccountNumber)
	return o
}

// SetFilterBeneficiaryPartyAccountNumber adds the filterBeneficiaryPartyAccountNumber to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterBeneficiaryPartyAccountNumber(filterBeneficiaryPartyAccountNumber *string) {
	o.FilterBeneficiaryPartyAccountNumber = filterBeneficiaryPartyAccountNumber
}

// WithFilterBeneficiaryPartyBankID adds the filterBeneficiaryPartyBankID to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterBeneficiaryPartyBankID(filterBeneficiaryPartyBankID *string) *GetDirectdebitsParams {
	o.SetFilterBeneficiaryPartyBankID(filterBeneficiaryPartyBankID)
	return o
}

// SetFilterBeneficiaryPartyBankID adds the filterBeneficiaryPartyBankId to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterBeneficiaryPartyBankID(filterBeneficiaryPartyBankID *string) {
	o.FilterBeneficiaryPartyBankID = filterBeneficiaryPartyBankID
}

// WithFilterClearingID adds the filterClearingID to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterClearingID(filterClearingID *string) *GetDirectdebitsParams {
	o.SetFilterClearingID(filterClearingID)
	return o
}

// SetFilterClearingID adds the filterClearingId to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterClearingID(filterClearingID *string) {
	o.FilterClearingID = filterClearingID
}

// WithFilterCurrency adds the filterCurrency to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterCurrency(filterCurrency *string) *GetDirectdebitsParams {
	o.SetFilterCurrency(filterCurrency)
	return o
}

// SetFilterCurrency adds the filterCurrency to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterCurrency(filterCurrency *string) {
	o.FilterCurrency = filterCurrency
}

// WithFilterDebtorPartyAccountNumber adds the filterDebtorPartyAccountNumber to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterDebtorPartyAccountNumber(filterDebtorPartyAccountNumber *string) *GetDirectdebitsParams {
	o.SetFilterDebtorPartyAccountNumber(filterDebtorPartyAccountNumber)
	return o
}

// SetFilterDebtorPartyAccountNumber adds the filterDebtorPartyAccountNumber to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterDebtorPartyAccountNumber(filterDebtorPartyAccountNumber *string) {
	o.FilterDebtorPartyAccountNumber = filterDebtorPartyAccountNumber
}

// WithFilterDebtorPartyBankID adds the filterDebtorPartyBankID to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterDebtorPartyBankID(filterDebtorPartyBankID *string) *GetDirectdebitsParams {
	o.SetFilterDebtorPartyBankID(filterDebtorPartyBankID)
	return o
}

// SetFilterDebtorPartyBankID adds the filterDebtorPartyBankId to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterDebtorPartyBankID(filterDebtorPartyBankID *string) {
	o.FilterDebtorPartyBankID = filterDebtorPartyBankID
}

// WithFilterOrganisationID adds the filterOrganisationID to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetDirectdebitsParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WithFilterPaymentScheme adds the filterPaymentScheme to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterPaymentScheme(filterPaymentScheme *string) *GetDirectdebitsParams {
	o.SetFilterPaymentScheme(filterPaymentScheme)
	return o
}

// SetFilterPaymentScheme adds the filterPaymentScheme to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterPaymentScheme(filterPaymentScheme *string) {
	o.FilterPaymentScheme = filterPaymentScheme
}

// WithFilterPaymentType adds the filterPaymentType to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterPaymentType(filterPaymentType *string) *GetDirectdebitsParams {
	o.SetFilterPaymentType(filterPaymentType)
	return o
}

// SetFilterPaymentType adds the filterPaymentType to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterPaymentType(filterPaymentType *string) {
	o.FilterPaymentType = filterPaymentType
}

// WithFilterProcessingDateFrom adds the filterProcessingDateFrom to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterProcessingDateFrom(filterProcessingDateFrom *strfmt.Date) *GetDirectdebitsParams {
	o.SetFilterProcessingDateFrom(filterProcessingDateFrom)
	return o
}

// SetFilterProcessingDateFrom adds the filterProcessingDateFrom to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterProcessingDateFrom(filterProcessingDateFrom *strfmt.Date) {
	o.FilterProcessingDateFrom = filterProcessingDateFrom
}

// WithFilterProcessingDateTo adds the filterProcessingDateTo to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterProcessingDateTo(filterProcessingDateTo *strfmt.Date) *GetDirectdebitsParams {
	o.SetFilterProcessingDateTo(filterProcessingDateTo)
	return o
}

// SetFilterProcessingDateTo adds the filterProcessingDateTo to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterProcessingDateTo(filterProcessingDateTo *strfmt.Date) {
	o.FilterProcessingDateTo = filterProcessingDateTo
}

// WithFilterReference adds the filterReference to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterReference(filterReference *string) *GetDirectdebitsParams {
	o.SetFilterReference(filterReference)
	return o
}

// SetFilterReference adds the filterReference to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterReference(filterReference *string) {
	o.FilterReference = filterReference
}

// WithFilterUniqueSchemeID adds the filterUniqueSchemeID to the get directdebits params
func (o *GetDirectdebitsParams) WithFilterUniqueSchemeID(filterUniqueSchemeID *string) *GetDirectdebitsParams {
	o.SetFilterUniqueSchemeID(filterUniqueSchemeID)
	return o
}

// SetFilterUniqueSchemeID adds the filterUniqueSchemeId to the get directdebits params
func (o *GetDirectdebitsParams) SetFilterUniqueSchemeID(filterUniqueSchemeID *string) {
	o.FilterUniqueSchemeID = filterUniqueSchemeID
}

// WithPageNumber adds the pageNumber to the get directdebits params
func (o *GetDirectdebitsParams) WithPageNumber(pageNumber *string) *GetDirectdebitsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get directdebits params
func (o *GetDirectdebitsParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get directdebits params
func (o *GetDirectdebitsParams) WithPageSize(pageSize *int64) *GetDirectdebitsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get directdebits params
func (o *GetDirectdebitsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetDirectdebitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterAdmissionAdmissionDateFrom != nil {

		// query param filter[admission.admission_date_from]
		var qrFilterAdmissionAdmissionDateFrom strfmt.DateTime

		if o.FilterAdmissionAdmissionDateFrom != nil {
			qrFilterAdmissionAdmissionDateFrom = *o.FilterAdmissionAdmissionDateFrom
		}
		qFilterAdmissionAdmissionDateFrom := qrFilterAdmissionAdmissionDateFrom.String()
		if qFilterAdmissionAdmissionDateFrom != "" {

			if err := r.SetQueryParam("filter[admission.admission_date_from]", qFilterAdmissionAdmissionDateFrom); err != nil {
				return err
			}
		}
	}

	if o.FilterAdmissionAdmissionDateTo != nil {

		// query param filter[admission.admission_date_to]
		var qrFilterAdmissionAdmissionDateTo strfmt.DateTime

		if o.FilterAdmissionAdmissionDateTo != nil {
			qrFilterAdmissionAdmissionDateTo = *o.FilterAdmissionAdmissionDateTo
		}
		qFilterAdmissionAdmissionDateTo := qrFilterAdmissionAdmissionDateTo.String()
		if qFilterAdmissionAdmissionDateTo != "" {

			if err := r.SetQueryParam("filter[admission.admission_date_to]", qFilterAdmissionAdmissionDateTo); err != nil {
				return err
			}
		}
	}

	if o.FilterAdmissionSchemeStatusCode != nil {

		// query param filter[admission.scheme_status_code]
		var qrFilterAdmissionSchemeStatusCode string

		if o.FilterAdmissionSchemeStatusCode != nil {
			qrFilterAdmissionSchemeStatusCode = *o.FilterAdmissionSchemeStatusCode
		}
		qFilterAdmissionSchemeStatusCode := qrFilterAdmissionSchemeStatusCode
		if qFilterAdmissionSchemeStatusCode != "" {

			if err := r.SetQueryParam("filter[admission.scheme_status_code]", qFilterAdmissionSchemeStatusCode); err != nil {
				return err
			}
		}
	}

	if o.FilterAdmissionStatus != nil {

		// query param filter[admission.status]
		var qrFilterAdmissionStatus string

		if o.FilterAdmissionStatus != nil {
			qrFilterAdmissionStatus = *o.FilterAdmissionStatus
		}
		qFilterAdmissionStatus := qrFilterAdmissionStatus
		if qFilterAdmissionStatus != "" {

			if err := r.SetQueryParam("filter[admission.status]", qFilterAdmissionStatus); err != nil {
				return err
			}
		}
	}

	if o.FilterAmount != nil {

		// query param filter[amount]
		var qrFilterAmount string

		if o.FilterAmount != nil {
			qrFilterAmount = *o.FilterAmount
		}
		qFilterAmount := qrFilterAmount
		if qFilterAmount != "" {

			if err := r.SetQueryParam("filter[amount]", qFilterAmount); err != nil {
				return err
			}
		}
	}

	if o.FilterBeneficiaryPartyAccountNumber != nil {

		// query param filter[beneficiary_party.account_number]
		var qrFilterBeneficiaryPartyAccountNumber string

		if o.FilterBeneficiaryPartyAccountNumber != nil {
			qrFilterBeneficiaryPartyAccountNumber = *o.FilterBeneficiaryPartyAccountNumber
		}
		qFilterBeneficiaryPartyAccountNumber := qrFilterBeneficiaryPartyAccountNumber
		if qFilterBeneficiaryPartyAccountNumber != "" {

			if err := r.SetQueryParam("filter[beneficiary_party.account_number]", qFilterBeneficiaryPartyAccountNumber); err != nil {
				return err
			}
		}
	}

	if o.FilterBeneficiaryPartyBankID != nil {

		// query param filter[beneficiary_party.bank_id]
		var qrFilterBeneficiaryPartyBankID string

		if o.FilterBeneficiaryPartyBankID != nil {
			qrFilterBeneficiaryPartyBankID = *o.FilterBeneficiaryPartyBankID
		}
		qFilterBeneficiaryPartyBankID := qrFilterBeneficiaryPartyBankID
		if qFilterBeneficiaryPartyBankID != "" {

			if err := r.SetQueryParam("filter[beneficiary_party.bank_id]", qFilterBeneficiaryPartyBankID); err != nil {
				return err
			}
		}
	}

	if o.FilterClearingID != nil {

		// query param filter[clearing_id]
		var qrFilterClearingID string

		if o.FilterClearingID != nil {
			qrFilterClearingID = *o.FilterClearingID
		}
		qFilterClearingID := qrFilterClearingID
		if qFilterClearingID != "" {

			if err := r.SetQueryParam("filter[clearing_id]", qFilterClearingID); err != nil {
				return err
			}
		}
	}

	if o.FilterCurrency != nil {

		// query param filter[currency]
		var qrFilterCurrency string

		if o.FilterCurrency != nil {
			qrFilterCurrency = *o.FilterCurrency
		}
		qFilterCurrency := qrFilterCurrency
		if qFilterCurrency != "" {

			if err := r.SetQueryParam("filter[currency]", qFilterCurrency); err != nil {
				return err
			}
		}
	}

	if o.FilterDebtorPartyAccountNumber != nil {

		// query param filter[debtor_party.account_number]
		var qrFilterDebtorPartyAccountNumber string

		if o.FilterDebtorPartyAccountNumber != nil {
			qrFilterDebtorPartyAccountNumber = *o.FilterDebtorPartyAccountNumber
		}
		qFilterDebtorPartyAccountNumber := qrFilterDebtorPartyAccountNumber
		if qFilterDebtorPartyAccountNumber != "" {

			if err := r.SetQueryParam("filter[debtor_party.account_number]", qFilterDebtorPartyAccountNumber); err != nil {
				return err
			}
		}
	}

	if o.FilterDebtorPartyBankID != nil {

		// query param filter[debtor_party.bank_id]
		var qrFilterDebtorPartyBankID string

		if o.FilterDebtorPartyBankID != nil {
			qrFilterDebtorPartyBankID = *o.FilterDebtorPartyBankID
		}
		qFilterDebtorPartyBankID := qrFilterDebtorPartyBankID
		if qFilterDebtorPartyBankID != "" {

			if err := r.SetQueryParam("filter[debtor_party.bank_id]", qFilterDebtorPartyBankID); err != nil {
				return err
			}
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

	if o.FilterPaymentScheme != nil {

		// query param filter[payment_scheme]
		var qrFilterPaymentScheme string

		if o.FilterPaymentScheme != nil {
			qrFilterPaymentScheme = *o.FilterPaymentScheme
		}
		qFilterPaymentScheme := qrFilterPaymentScheme
		if qFilterPaymentScheme != "" {

			if err := r.SetQueryParam("filter[payment_scheme]", qFilterPaymentScheme); err != nil {
				return err
			}
		}
	}

	if o.FilterPaymentType != nil {

		// query param filter[payment_type]
		var qrFilterPaymentType string

		if o.FilterPaymentType != nil {
			qrFilterPaymentType = *o.FilterPaymentType
		}
		qFilterPaymentType := qrFilterPaymentType
		if qFilterPaymentType != "" {

			if err := r.SetQueryParam("filter[payment_type]", qFilterPaymentType); err != nil {
				return err
			}
		}
	}

	if o.FilterProcessingDateFrom != nil {

		// query param filter[processing_date_from]
		var qrFilterProcessingDateFrom strfmt.Date

		if o.FilterProcessingDateFrom != nil {
			qrFilterProcessingDateFrom = *o.FilterProcessingDateFrom
		}
		qFilterProcessingDateFrom := qrFilterProcessingDateFrom.String()
		if qFilterProcessingDateFrom != "" {

			if err := r.SetQueryParam("filter[processing_date_from]", qFilterProcessingDateFrom); err != nil {
				return err
			}
		}
	}

	if o.FilterProcessingDateTo != nil {

		// query param filter[processing_date_to]
		var qrFilterProcessingDateTo strfmt.Date

		if o.FilterProcessingDateTo != nil {
			qrFilterProcessingDateTo = *o.FilterProcessingDateTo
		}
		qFilterProcessingDateTo := qrFilterProcessingDateTo.String()
		if qFilterProcessingDateTo != "" {

			if err := r.SetQueryParam("filter[processing_date_to]", qFilterProcessingDateTo); err != nil {
				return err
			}
		}
	}

	if o.FilterReference != nil {

		// query param filter[reference]
		var qrFilterReference string

		if o.FilterReference != nil {
			qrFilterReference = *o.FilterReference
		}
		qFilterReference := qrFilterReference
		if qFilterReference != "" {

			if err := r.SetQueryParam("filter[reference]", qFilterReference); err != nil {
				return err
			}
		}
	}

	if o.FilterUniqueSchemeID != nil {

		// query param filter[unique_scheme_id]
		var qrFilterUniqueSchemeID string

		if o.FilterUniqueSchemeID != nil {
			qrFilterUniqueSchemeID = *o.FilterUniqueSchemeID
		}
		qFilterUniqueSchemeID := qrFilterUniqueSchemeID
		if qFilterUniqueSchemeID != "" {

			if err := r.SetQueryParam("filter[unique_scheme_id]", qFilterUniqueSchemeID); err != nil {
				return err
			}
		}
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

// bindParamGetDirectdebits binds the parameter filter[organisation_id]
func (o *GetDirectdebitsParams) bindParamFilterOrganisationID(formats strfmt.Registry) []string {
	filterOrganisationIDIR := o.FilterOrganisationID

	var filterOrganisationIDIC []string
	for _, filterOrganisationIDIIR := range filterOrganisationIDIR { // explode []strfmt.UUID

		filterOrganisationIDIIV := filterOrganisationIDIIR.String() // strfmt.UUID as string
		filterOrganisationIDIC = append(filterOrganisationIDIC, filterOrganisationIDIIV)
	}

	// items.CollectionFormat: ""
	filterOrganisationIDIS := swag.JoinByFormat(filterOrganisationIDIC, "")

	return filterOrganisationIDIS
}
