// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SigningKeysRequestData signing keys request data
//
// swagger:model SigningKeysRequestData
type SigningKeysRequestData struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// organisation id
	// Required: true
	OrganisationID *string `json:"organisation_id"`

	// type
	// Required: true
	// Enum: [signing_keys]
	Type *string `json:"type"`
}

// Validate validates this signing keys request data
func (m *SigningKeysRequestData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
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

func (m *SigningKeysRequestData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SigningKeysRequestData) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	return nil
}

var signingKeysRequestDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["signing_keys"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		signingKeysRequestDataTypeTypePropEnum = append(signingKeysRequestDataTypeTypePropEnum, v)
	}
}

const (

	// SigningKeysRequestDataTypeSigningKeys captures enum value "signing_keys"
	SigningKeysRequestDataTypeSigningKeys string = "signing_keys"
)

// prop value enum
func (m *SigningKeysRequestData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, signingKeysRequestDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SigningKeysRequestData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SigningKeysRequestData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigningKeysRequestData) UnmarshalBinary(b []byte) error {
	var res SigningKeysRequestData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
