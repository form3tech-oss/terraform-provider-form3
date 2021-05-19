// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PaymentAdmissionStatusReason payment admission status reason
//
// swagger:model PaymentAdmissionStatusReason
type PaymentAdmissionStatusReason string

const (

	// PaymentAdmissionStatusReasonAccepted captures enum value "accepted"
	PaymentAdmissionStatusReasonAccepted PaymentAdmissionStatusReason = "accepted"

	// PaymentAdmissionStatusReasonInvalidBeneficiaryDetails captures enum value "invalid_beneficiary_details"
	PaymentAdmissionStatusReasonInvalidBeneficiaryDetails PaymentAdmissionStatusReason = "invalid_beneficiary_details"

	// PaymentAdmissionStatusReasonBankidNotProvisioned captures enum value "bankid_not_provisioned"
	PaymentAdmissionStatusReasonBankidNotProvisioned PaymentAdmissionStatusReason = "bankid_not_provisioned"

	// PaymentAdmissionStatusReasonUnknownAccountnumber captures enum value "unknown_accountnumber"
	PaymentAdmissionStatusReasonUnknownAccountnumber PaymentAdmissionStatusReason = "unknown_accountnumber"
)

// for schema
var paymentAdmissionStatusReasonEnum []interface{}

func init() {
	var res []PaymentAdmissionStatusReason
	if err := json.Unmarshal([]byte(`["accepted","invalid_beneficiary_details","bankid_not_provisioned","unknown_accountnumber"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAdmissionStatusReasonEnum = append(paymentAdmissionStatusReasonEnum, v)
	}
}

func (m PaymentAdmissionStatusReason) validatePaymentAdmissionStatusReasonEnum(path, location string, value PaymentAdmissionStatusReason) error {
	if err := validate.EnumCase(path, location, value, paymentAdmissionStatusReasonEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment admission status reason
func (m PaymentAdmissionStatusReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentAdmissionStatusReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
