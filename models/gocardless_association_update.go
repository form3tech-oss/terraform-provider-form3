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

// GocardlessAssociationUpdate gocardless association update
//
// swagger:model GocardlessAssociationUpdate
type GocardlessAssociationUpdate struct {

	// attributes
	Attributes *GocardlessAssociationPatchAttributes `json:"attributes,omitempty"`
}

// Validate validates this gocardless association update
func (m *GocardlessAssociationUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GocardlessAssociationUpdate) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gocardless association update based on the context it is used
func (m *GocardlessAssociationUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GocardlessAssociationUpdate) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {
		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GocardlessAssociationUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GocardlessAssociationUpdate) UnmarshalBinary(b []byte) error {
	var res GocardlessAssociationUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
