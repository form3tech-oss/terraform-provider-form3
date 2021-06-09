// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PayportAssociationAttributes payport association attributes
//
// swagger:model PayportAssociationAttributes
type PayportAssociationAttributes struct {

	// customer sending fps institution
	CustomerSendingFpsInstitution string `json:"customer_sending_fps_institution,omitempty"`

	// participant id
	ParticipantID string `json:"participant_id,omitempty"`

	// participant type
	ParticipantType PayportParticipantType `json:"participant_type,omitempty"`

	// sponsor account number
	SponsorAccountNumber string `json:"sponsor_account_number,omitempty"`

	// sponsor bank id
	SponsorBankID string `json:"sponsor_bank_id,omitempty"`
}

// Validate validates this payport association attributes
func (m *PayportAssociationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParticipantType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayportAssociationAttributes) validateParticipantType(formats strfmt.Registry) error {

	if swag.IsZero(m.ParticipantType) { // not required
		return nil
	}

	if err := m.ParticipantType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("participant_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PayportAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayportAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res PayportAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
