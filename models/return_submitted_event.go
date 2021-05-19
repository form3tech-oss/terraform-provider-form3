// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReturnSubmittedEvent return submitted event
//
// swagger:model ReturnSubmittedEvent
type ReturnSubmittedEvent struct {

	// payment
	Payment *Payment `json:"payment,omitempty"`

	// return payment
	ReturnPayment *ReturnPayment `json:"return_payment,omitempty"`

	// return submission
	ReturnSubmission *ReturnSubmission `json:"return_submission,omitempty"`
}

// Validate validates this return submitted event
func (m *ReturnSubmittedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnPayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmittedEvent) validatePayment(formats strfmt.Registry) error {
	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnSubmittedEvent) validateReturnPayment(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnPayment) { // not required
		return nil
	}

	if m.ReturnPayment != nil {
		if err := m.ReturnPayment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("return_payment")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnSubmittedEvent) validateReturnSubmission(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnSubmission) { // not required
		return nil
	}

	if m.ReturnSubmission != nil {
		if err := m.ReturnSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("return_submission")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this return submitted event based on the context it is used
func (m *ReturnSubmittedEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePayment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReturnPayment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReturnSubmission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmittedEvent) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

	if m.Payment != nil {
		if err := m.Payment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnSubmittedEvent) contextValidateReturnPayment(ctx context.Context, formats strfmt.Registry) error {

	if m.ReturnPayment != nil {
		if err := m.ReturnPayment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("return_payment")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnSubmittedEvent) contextValidateReturnSubmission(ctx context.Context, formats strfmt.Registry) error {

	if m.ReturnSubmission != nil {
		if err := m.ReturnSubmission.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("return_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmittedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmittedEvent) UnmarshalBinary(b []byte) error {
	var res ReturnSubmittedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
