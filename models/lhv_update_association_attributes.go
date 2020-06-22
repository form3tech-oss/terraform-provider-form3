// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LhvUpdateAssociationAttributes lhv update association attributes
// swagger:model LhvUpdateAssociationAttributes
type LhvUpdateAssociationAttributes struct {

	// client code
	// Min Length: 1
	ClientCode string `json:"client_code,omitempty"`

	// client country
	// Min Length: 1
	ClientCountry string `json:"client_country,omitempty"`

	// name
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// use simulator
	UseSimulator *bool `json:"use_simulator,omitempty"`
}

// Validate validates this lhv update association attributes
func (m *LhvUpdateAssociationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LhvUpdateAssociationAttributes) validateClientCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientCode) { // not required
		return nil
	}

	if err := validate.MinLength("client_code", "body", string(m.ClientCode), 1); err != nil {
		return err
	}

	return nil
}

func (m *LhvUpdateAssociationAttributes) validateClientCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientCountry) { // not required
		return nil
	}

	if err := validate.MinLength("client_country", "body", string(m.ClientCountry), 1); err != nil {
		return err
	}

	return nil
}

func (m *LhvUpdateAssociationAttributes) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LhvUpdateAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LhvUpdateAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res LhvUpdateAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
