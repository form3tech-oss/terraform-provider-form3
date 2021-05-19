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

// PaymentSubmissionStatus payment submission status
//
// swagger:model PaymentSubmissionStatus
type PaymentSubmissionStatus string

const (

	// PaymentSubmissionStatusAccepted captures enum value "accepted"
	PaymentSubmissionStatusAccepted PaymentSubmissionStatus = "accepted"

	// PaymentSubmissionStatusLimitCheckPending captures enum value "limit_check_pending"
	PaymentSubmissionStatusLimitCheckPending PaymentSubmissionStatus = "limit_check_pending"

	// PaymentSubmissionStatusLimitCheckFailed captures enum value "limit_check_failed"
	PaymentSubmissionStatusLimitCheckFailed PaymentSubmissionStatus = "limit_check_failed"

	// PaymentSubmissionStatusLimitCheckPassed captures enum value "limit_check_passed"
	PaymentSubmissionStatusLimitCheckPassed PaymentSubmissionStatus = "limit_check_passed"

	// PaymentSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	PaymentSubmissionStatusReleasedToGateway PaymentSubmissionStatus = "released_to_gateway"

	// PaymentSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	PaymentSubmissionStatusQueuedForDelivery PaymentSubmissionStatus = "queued_for_delivery"

	// PaymentSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	PaymentSubmissionStatusDeliveryConfirmed PaymentSubmissionStatus = "delivery_confirmed"

	// PaymentSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	PaymentSubmissionStatusDeliveryFailed PaymentSubmissionStatus = "delivery_failed"

	// PaymentSubmissionStatusSubmitted captures enum value "submitted"
	PaymentSubmissionStatusSubmitted PaymentSubmissionStatus = "submitted"

	// PaymentSubmissionStatusValidationPending captures enum value "validation_pending"
	PaymentSubmissionStatusValidationPending PaymentSubmissionStatus = "validation_pending"
)

// for schema
var paymentSubmissionStatusEnum []interface{}

func init() {
	var res []PaymentSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","limit_check_pending","limit_check_failed","limit_check_passed","released_to_gateway","queued_for_delivery","delivery_confirmed","delivery_failed","submitted","validation_pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSubmissionStatusEnum = append(paymentSubmissionStatusEnum, v)
	}
}

func (m PaymentSubmissionStatus) validatePaymentSubmissionStatusEnum(path, location string, value PaymentSubmissionStatus) error {
	if err := validate.EnumCase(path, location, value, paymentSubmissionStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment submission status
func (m PaymentSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
