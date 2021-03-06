// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SepaSctAssociationRelationships sepa sct association relationships
//
// swagger:model SepaSctAssociationRelationships
type SepaSctAssociationRelationships struct {

	// sponsor
	Sponsor *SepaSctAssociationRelationshipsSponsor `json:"sponsor,omitempty"`
}

// Validate validates this sepa sct association relationships
func (m *SepaSctAssociationRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSponsor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaSctAssociationRelationships) validateSponsor(formats strfmt.Registry) error {

	if swag.IsZero(m.Sponsor) { // not required
		return nil
	}

	if m.Sponsor != nil {
		if err := m.Sponsor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sponsor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepaSctAssociationRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaSctAssociationRelationships) UnmarshalBinary(b []byte) error {
	var res SepaSctAssociationRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SepaSctAssociationRelationshipsSponsor sepa sct association relationships sponsor
//
// swagger:model SepaSctAssociationRelationshipsSponsor
type SepaSctAssociationRelationshipsSponsor struct {

	// data
	Data *SepaSctSponsorAssociationReference `json:"data,omitempty"`
}

// Validate validates this sepa sct association relationships sponsor
func (m *SepaSctAssociationRelationshipsSponsor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaSctAssociationRelationshipsSponsor) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sponsor" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepaSctAssociationRelationshipsSponsor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaSctAssociationRelationshipsSponsor) UnmarshalBinary(b []byte) error {
	var res SepaSctAssociationRelationshipsSponsor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
