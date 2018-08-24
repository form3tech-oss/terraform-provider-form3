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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPaymentsParams creates a new GetPaymentsParams object
// with the default values initialized.
func NewGetPaymentsParams() *GetPaymentsParams {
	var ()
	return &GetPaymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsParamsWithTimeout creates a new GetPaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsParamsWithTimeout(timeout time.Duration) *GetPaymentsParams {
	var ()
	return &GetPaymentsParams{

		timeout: timeout,
	}
}

// NewGetPaymentsParamsWithContext creates a new GetPaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsParamsWithContext(ctx context.Context) *GetPaymentsParams {
	var ()
	return &GetPaymentsParams{

		Context: ctx,
	}
}

// NewGetPaymentsParamsWithHTTPClient creates a new GetPaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsParamsWithHTTPClient(client *http.Client) *GetPaymentsParams {
	var ()
	return &GetPaymentsParams{
		HTTPClient: client,
	}
}

/*GetPaymentsParams contains all the parameters to send to the API endpoint
for the get payments operation typically these are written to a http.Request
*/
type GetPaymentsParams struct {

	/*FilterAdmissionAdmissionDateFrom*/
	FilterAdmissionAdmissionDateFrom *strfmt.DateTime
	/*FilterAdmissionAdmissionDateTo*/
	FilterAdmissionAdmissionDateTo *strfmt.DateTime
	/*FilterAdmissionSchemeStatusCode
	  Filter by admission scheme status code

	*/
	FilterAdmissionSchemeStatusCode *string
	/*FilterAdmissionStatus
	  Filter by admission status

	*/
	FilterAdmissionStatus *string
	/*FilterBeneficiaryPartyAccountNumber*/
	FilterBeneficiaryPartyAccountNumber *string
	/*FilterBeneficiaryPartyBankID*/
	FilterBeneficiaryPartyBankID *string
	/*FilterCurrency*/
	FilterCurrency *string
	/*FilterDebtorPartyAccountNumber*/
	FilterDebtorPartyAccountNumber *string
	/*FilterDebtorPartyBankID*/
	FilterDebtorPartyBankID *string
	/*FilterOrganisationID
	  Filter by organisation id

	*/
	FilterOrganisationID []strfmt.UUID
	/*FilterPaymentScheme*/
	FilterPaymentScheme *string
	/*FilterPaymentType*/
	FilterPaymentType *string
	/*FilterProcessingDateFrom*/
	FilterProcessingDateFrom *strfmt.Date
	/*FilterProcessingDateTo*/
	FilterProcessingDateTo *strfmt.Date
	/*FilterSchemeTransactionID*/
	FilterSchemeTransactionID *string
	/*FilterSubmissionSchemeStatusCode
	  Filter by submission scheme status code

	*/
	FilterSubmissionSchemeStatusCode *string
	/*FilterSubmissionStatus
	  Filter by submission status

	*/
	FilterSubmissionStatus *string
	/*FilterSubmissionSubmissionDateFrom*/
	FilterSubmissionSubmissionDateFrom *strfmt.DateTime
	/*FilterSubmissionSubmissionDateTo*/
	FilterSubmissionSubmissionDateTo *strfmt.DateTime
	/*FilterUniqueSchemeID*/
	FilterUniqueSchemeID *string
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

// WithTimeout adds the timeout to the get payments params
func (o *GetPaymentsParams) WithTimeout(timeout time.Duration) *GetPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments params
func (o *GetPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments params
func (o *GetPaymentsParams) WithContext(ctx context.Context) *GetPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments params
func (o *GetPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments params
func (o *GetPaymentsParams) WithHTTPClient(client *http.Client) *GetPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments params
func (o *GetPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterAdmissionAdmissionDateFrom adds the filterAdmissionAdmissionDateFrom to the get payments params
func (o *GetPaymentsParams) WithFilterAdmissionAdmissionDateFrom(filterAdmissionAdmissionDateFrom *strfmt.DateTime) *GetPaymentsParams {
	o.SetFilterAdmissionAdmissionDateFrom(filterAdmissionAdmissionDateFrom)
	return o
}

// SetFilterAdmissionAdmissionDateFrom adds the filterAdmissionAdmissionDateFrom to the get payments params
func (o *GetPaymentsParams) SetFilterAdmissionAdmissionDateFrom(filterAdmissionAdmissionDateFrom *strfmt.DateTime) {
	o.FilterAdmissionAdmissionDateFrom = filterAdmissionAdmissionDateFrom
}

// WithFilterAdmissionAdmissionDateTo adds the filterAdmissionAdmissionDateTo to the get payments params
func (o *GetPaymentsParams) WithFilterAdmissionAdmissionDateTo(filterAdmissionAdmissionDateTo *strfmt.DateTime) *GetPaymentsParams {
	o.SetFilterAdmissionAdmissionDateTo(filterAdmissionAdmissionDateTo)
	return o
}

// SetFilterAdmissionAdmissionDateTo adds the filterAdmissionAdmissionDateTo to the get payments params
func (o *GetPaymentsParams) SetFilterAdmissionAdmissionDateTo(filterAdmissionAdmissionDateTo *strfmt.DateTime) {
	o.FilterAdmissionAdmissionDateTo = filterAdmissionAdmissionDateTo
}

// WithFilterAdmissionSchemeStatusCode adds the filterAdmissionSchemeStatusCode to the get payments params
func (o *GetPaymentsParams) WithFilterAdmissionSchemeStatusCode(filterAdmissionSchemeStatusCode *string) *GetPaymentsParams {
	o.SetFilterAdmissionSchemeStatusCode(filterAdmissionSchemeStatusCode)
	return o
}

// SetFilterAdmissionSchemeStatusCode adds the filterAdmissionSchemeStatusCode to the get payments params
func (o *GetPaymentsParams) SetFilterAdmissionSchemeStatusCode(filterAdmissionSchemeStatusCode *string) {
	o.FilterAdmissionSchemeStatusCode = filterAdmissionSchemeStatusCode
}

// WithFilterAdmissionStatus adds the filterAdmissionStatus to the get payments params
func (o *GetPaymentsParams) WithFilterAdmissionStatus(filterAdmissionStatus *string) *GetPaymentsParams {
	o.SetFilterAdmissionStatus(filterAdmissionStatus)
	return o
}

// SetFilterAdmissionStatus adds the filterAdmissionStatus to the get payments params
func (o *GetPaymentsParams) SetFilterAdmissionStatus(filterAdmissionStatus *string) {
	o.FilterAdmissionStatus = filterAdmissionStatus
}

// WithFilterBeneficiaryPartyAccountNumber adds the filterBeneficiaryPartyAccountNumber to the get payments params
func (o *GetPaymentsParams) WithFilterBeneficiaryPartyAccountNumber(filterBeneficiaryPartyAccountNumber *string) *GetPaymentsParams {
	o.SetFilterBeneficiaryPartyAccountNumber(filterBeneficiaryPartyAccountNumber)
	return o
}

// SetFilterBeneficiaryPartyAccountNumber adds the filterBeneficiaryPartyAccountNumber to the get payments params
func (o *GetPaymentsParams) SetFilterBeneficiaryPartyAccountNumber(filterBeneficiaryPartyAccountNumber *string) {
	o.FilterBeneficiaryPartyAccountNumber = filterBeneficiaryPartyAccountNumber
}

// WithFilterBeneficiaryPartyBankID adds the filterBeneficiaryPartyBankID to the get payments params
func (o *GetPaymentsParams) WithFilterBeneficiaryPartyBankID(filterBeneficiaryPartyBankID *string) *GetPaymentsParams {
	o.SetFilterBeneficiaryPartyBankID(filterBeneficiaryPartyBankID)
	return o
}

// SetFilterBeneficiaryPartyBankID adds the filterBeneficiaryPartyBankId to the get payments params
func (o *GetPaymentsParams) SetFilterBeneficiaryPartyBankID(filterBeneficiaryPartyBankID *string) {
	o.FilterBeneficiaryPartyBankID = filterBeneficiaryPartyBankID
}

// WithFilterCurrency adds the filterCurrency to the get payments params
func (o *GetPaymentsParams) WithFilterCurrency(filterCurrency *string) *GetPaymentsParams {
	o.SetFilterCurrency(filterCurrency)
	return o
}

// SetFilterCurrency adds the filterCurrency to the get payments params
func (o *GetPaymentsParams) SetFilterCurrency(filterCurrency *string) {
	o.FilterCurrency = filterCurrency
}

// WithFilterDebtorPartyAccountNumber adds the filterDebtorPartyAccountNumber to the get payments params
func (o *GetPaymentsParams) WithFilterDebtorPartyAccountNumber(filterDebtorPartyAccountNumber *string) *GetPaymentsParams {
	o.SetFilterDebtorPartyAccountNumber(filterDebtorPartyAccountNumber)
	return o
}

// SetFilterDebtorPartyAccountNumber adds the filterDebtorPartyAccountNumber to the get payments params
func (o *GetPaymentsParams) SetFilterDebtorPartyAccountNumber(filterDebtorPartyAccountNumber *string) {
	o.FilterDebtorPartyAccountNumber = filterDebtorPartyAccountNumber
}

// WithFilterDebtorPartyBankID adds the filterDebtorPartyBankID to the get payments params
func (o *GetPaymentsParams) WithFilterDebtorPartyBankID(filterDebtorPartyBankID *string) *GetPaymentsParams {
	o.SetFilterDebtorPartyBankID(filterDebtorPartyBankID)
	return o
}

// SetFilterDebtorPartyBankID adds the filterDebtorPartyBankId to the get payments params
func (o *GetPaymentsParams) SetFilterDebtorPartyBankID(filterDebtorPartyBankID *string) {
	o.FilterDebtorPartyBankID = filterDebtorPartyBankID
}

// WithFilterOrganisationID adds the filterOrganisationID to the get payments params
func (o *GetPaymentsParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetPaymentsParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get payments params
func (o *GetPaymentsParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WithFilterPaymentScheme adds the filterPaymentScheme to the get payments params
func (o *GetPaymentsParams) WithFilterPaymentScheme(filterPaymentScheme *string) *GetPaymentsParams {
	o.SetFilterPaymentScheme(filterPaymentScheme)
	return o
}

// SetFilterPaymentScheme adds the filterPaymentScheme to the get payments params
func (o *GetPaymentsParams) SetFilterPaymentScheme(filterPaymentScheme *string) {
	o.FilterPaymentScheme = filterPaymentScheme
}

// WithFilterPaymentType adds the filterPaymentType to the get payments params
func (o *GetPaymentsParams) WithFilterPaymentType(filterPaymentType *string) *GetPaymentsParams {
	o.SetFilterPaymentType(filterPaymentType)
	return o
}

// SetFilterPaymentType adds the filterPaymentType to the get payments params
func (o *GetPaymentsParams) SetFilterPaymentType(filterPaymentType *string) {
	o.FilterPaymentType = filterPaymentType
}

// WithFilterProcessingDateFrom adds the filterProcessingDateFrom to the get payments params
func (o *GetPaymentsParams) WithFilterProcessingDateFrom(filterProcessingDateFrom *strfmt.Date) *GetPaymentsParams {
	o.SetFilterProcessingDateFrom(filterProcessingDateFrom)
	return o
}

// SetFilterProcessingDateFrom adds the filterProcessingDateFrom to the get payments params
func (o *GetPaymentsParams) SetFilterProcessingDateFrom(filterProcessingDateFrom *strfmt.Date) {
	o.FilterProcessingDateFrom = filterProcessingDateFrom
}

// WithFilterProcessingDateTo adds the filterProcessingDateTo to the get payments params
func (o *GetPaymentsParams) WithFilterProcessingDateTo(filterProcessingDateTo *strfmt.Date) *GetPaymentsParams {
	o.SetFilterProcessingDateTo(filterProcessingDateTo)
	return o
}

// SetFilterProcessingDateTo adds the filterProcessingDateTo to the get payments params
func (o *GetPaymentsParams) SetFilterProcessingDateTo(filterProcessingDateTo *strfmt.Date) {
	o.FilterProcessingDateTo = filterProcessingDateTo
}

// WithFilterSchemeTransactionID adds the filterSchemeTransactionID to the get payments params
func (o *GetPaymentsParams) WithFilterSchemeTransactionID(filterSchemeTransactionID *string) *GetPaymentsParams {
	o.SetFilterSchemeTransactionID(filterSchemeTransactionID)
	return o
}

// SetFilterSchemeTransactionID adds the filterSchemeTransactionId to the get payments params
func (o *GetPaymentsParams) SetFilterSchemeTransactionID(filterSchemeTransactionID *string) {
	o.FilterSchemeTransactionID = filterSchemeTransactionID
}

// WithFilterSubmissionSchemeStatusCode adds the filterSubmissionSchemeStatusCode to the get payments params
func (o *GetPaymentsParams) WithFilterSubmissionSchemeStatusCode(filterSubmissionSchemeStatusCode *string) *GetPaymentsParams {
	o.SetFilterSubmissionSchemeStatusCode(filterSubmissionSchemeStatusCode)
	return o
}

// SetFilterSubmissionSchemeStatusCode adds the filterSubmissionSchemeStatusCode to the get payments params
func (o *GetPaymentsParams) SetFilterSubmissionSchemeStatusCode(filterSubmissionSchemeStatusCode *string) {
	o.FilterSubmissionSchemeStatusCode = filterSubmissionSchemeStatusCode
}

// WithFilterSubmissionStatus adds the filterSubmissionStatus to the get payments params
func (o *GetPaymentsParams) WithFilterSubmissionStatus(filterSubmissionStatus *string) *GetPaymentsParams {
	o.SetFilterSubmissionStatus(filterSubmissionStatus)
	return o
}

// SetFilterSubmissionStatus adds the filterSubmissionStatus to the get payments params
func (o *GetPaymentsParams) SetFilterSubmissionStatus(filterSubmissionStatus *string) {
	o.FilterSubmissionStatus = filterSubmissionStatus
}

// WithFilterSubmissionSubmissionDateFrom adds the filterSubmissionSubmissionDateFrom to the get payments params
func (o *GetPaymentsParams) WithFilterSubmissionSubmissionDateFrom(filterSubmissionSubmissionDateFrom *strfmt.DateTime) *GetPaymentsParams {
	o.SetFilterSubmissionSubmissionDateFrom(filterSubmissionSubmissionDateFrom)
	return o
}

// SetFilterSubmissionSubmissionDateFrom adds the filterSubmissionSubmissionDateFrom to the get payments params
func (o *GetPaymentsParams) SetFilterSubmissionSubmissionDateFrom(filterSubmissionSubmissionDateFrom *strfmt.DateTime) {
	o.FilterSubmissionSubmissionDateFrom = filterSubmissionSubmissionDateFrom
}

// WithFilterSubmissionSubmissionDateTo adds the filterSubmissionSubmissionDateTo to the get payments params
func (o *GetPaymentsParams) WithFilterSubmissionSubmissionDateTo(filterSubmissionSubmissionDateTo *strfmt.DateTime) *GetPaymentsParams {
	o.SetFilterSubmissionSubmissionDateTo(filterSubmissionSubmissionDateTo)
	return o
}

// SetFilterSubmissionSubmissionDateTo adds the filterSubmissionSubmissionDateTo to the get payments params
func (o *GetPaymentsParams) SetFilterSubmissionSubmissionDateTo(filterSubmissionSubmissionDateTo *strfmt.DateTime) {
	o.FilterSubmissionSubmissionDateTo = filterSubmissionSubmissionDateTo
}

// WithFilterUniqueSchemeID adds the filterUniqueSchemeID to the get payments params
func (o *GetPaymentsParams) WithFilterUniqueSchemeID(filterUniqueSchemeID *string) *GetPaymentsParams {
	o.SetFilterUniqueSchemeID(filterUniqueSchemeID)
	return o
}

// SetFilterUniqueSchemeID adds the filterUniqueSchemeId to the get payments params
func (o *GetPaymentsParams) SetFilterUniqueSchemeID(filterUniqueSchemeID *string) {
	o.FilterUniqueSchemeID = filterUniqueSchemeID
}

// WithPageNumber adds the pageNumber to the get payments params
func (o *GetPaymentsParams) WithPageNumber(pageNumber *string) *GetPaymentsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get payments params
func (o *GetPaymentsParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get payments params
func (o *GetPaymentsParams) WithPageSize(pageSize *int64) *GetPaymentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get payments params
func (o *GetPaymentsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
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

	if o.FilterSchemeTransactionID != nil {

		// query param filter[scheme_transaction_id]
		var qrFilterSchemeTransactionID string
		if o.FilterSchemeTransactionID != nil {
			qrFilterSchemeTransactionID = *o.FilterSchemeTransactionID
		}
		qFilterSchemeTransactionID := qrFilterSchemeTransactionID
		if qFilterSchemeTransactionID != "" {
			if err := r.SetQueryParam("filter[scheme_transaction_id]", qFilterSchemeTransactionID); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionSchemeStatusCode != nil {

		// query param filter[submission.scheme_status_code]
		var qrFilterSubmissionSchemeStatusCode string
		if o.FilterSubmissionSchemeStatusCode != nil {
			qrFilterSubmissionSchemeStatusCode = *o.FilterSubmissionSchemeStatusCode
		}
		qFilterSubmissionSchemeStatusCode := qrFilterSubmissionSchemeStatusCode
		if qFilterSubmissionSchemeStatusCode != "" {
			if err := r.SetQueryParam("filter[submission.scheme_status_code]", qFilterSubmissionSchemeStatusCode); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionStatus != nil {

		// query param filter[submission.status]
		var qrFilterSubmissionStatus string
		if o.FilterSubmissionStatus != nil {
			qrFilterSubmissionStatus = *o.FilterSubmissionStatus
		}
		qFilterSubmissionStatus := qrFilterSubmissionStatus
		if qFilterSubmissionStatus != "" {
			if err := r.SetQueryParam("filter[submission.status]", qFilterSubmissionStatus); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionSubmissionDateFrom != nil {

		// query param filter[submission.submission_date_from]
		var qrFilterSubmissionSubmissionDateFrom strfmt.DateTime
		if o.FilterSubmissionSubmissionDateFrom != nil {
			qrFilterSubmissionSubmissionDateFrom = *o.FilterSubmissionSubmissionDateFrom
		}
		qFilterSubmissionSubmissionDateFrom := qrFilterSubmissionSubmissionDateFrom.String()
		if qFilterSubmissionSubmissionDateFrom != "" {
			if err := r.SetQueryParam("filter[submission.submission_date_from]", qFilterSubmissionSubmissionDateFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionSubmissionDateTo != nil {

		// query param filter[submission.submission_date_to]
		var qrFilterSubmissionSubmissionDateTo strfmt.DateTime
		if o.FilterSubmissionSubmissionDateTo != nil {
			qrFilterSubmissionSubmissionDateTo = *o.FilterSubmissionSubmissionDateTo
		}
		qFilterSubmissionSubmissionDateTo := qrFilterSubmissionSubmissionDateTo.String()
		if qFilterSubmissionSubmissionDateTo != "" {
			if err := r.SetQueryParam("filter[submission.submission_date_to]", qFilterSubmissionSubmissionDateTo); err != nil {
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
