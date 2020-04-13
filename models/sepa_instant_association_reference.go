// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SepaInstantAssociationReference sepa instant association reference
// swagger:model SepaInstantAssociationReference
type SepaInstantAssociationReference struct {

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// type
	// Required: true
	// Enum: [sepainstant_associations]
	Type string `json:"type"`
}

// Validate validates this sepa instant association reference
func (m *SepaInstantAssociationReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaInstantAssociationReference) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var sepaInstantAssociationReferenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sepainstant_associations"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sepaInstantAssociationReferenceTypeTypePropEnum = append(sepaInstantAssociationReferenceTypeTypePropEnum, v)
	}
}

const (

	// SepaInstantAssociationReferenceTypeSepainstantAssociations captures enum value "sepainstant_associations"
	SepaInstantAssociationReferenceTypeSepainstantAssociations string = "sepainstant_associations"
)

// prop value enum
func (m *SepaInstantAssociationReference) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sepaInstantAssociationReferenceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SepaInstantAssociationReference) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepaInstantAssociationReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaInstantAssociationReference) UnmarshalBinary(b []byte) error {
	var res SepaInstantAssociationReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
