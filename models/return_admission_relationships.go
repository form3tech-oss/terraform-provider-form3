// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReturnAdmissionRelationships return admission relationships
//
// swagger:model ReturnAdmissionRelationships
type ReturnAdmissionRelationships struct {

	// payment return
	PaymentReturn *RelationshipLinks `json:"payment_return,omitempty"`

	// validations
	Validations *RelationshipLinks `json:"validations,omitempty"`
}

// Validate validates this return admission relationships
func (m *ReturnAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnAdmissionRelationships) validatePaymentReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReturn) { // not required
		return nil
	}

	if m.PaymentReturn != nil {
		if err := m.PaymentReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_return")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnAdmissionRelationships) validateValidations(formats strfmt.Registry) error {

	if swag.IsZero(m.Validations) { // not required
		return nil
	}

	if m.Validations != nil {
		if err := m.Validations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ReturnAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
