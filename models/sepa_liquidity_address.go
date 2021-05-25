// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SepaLiquidityAddress sepa liquidity address
//
// swagger:model SepaLiquidityAddress
type SepaLiquidityAddress struct {

	// building number
	// Required: true
	// Max Length: 16
	// Min Length: 1
	BuildingNumber string `json:"building_number"`

	// city
	// Required: true
	// Max Length: 35
	// Min Length: 1
	City string `json:"city"`

	// country
	// Required: true
	// Max Length: 70
	// Min Length: 1
	Country string `json:"country"`

	// street
	// Required: true
	// Max Length: 70
	// Min Length: 1
	Street string `json:"street"`
}

// Validate validates this sepa liquidity address
func (m *SepaLiquidityAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildingNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaLiquidityAddress) validateBuildingNumber(formats strfmt.Registry) error {

	if err := validate.RequiredString("building_number", "body", m.BuildingNumber); err != nil {
		return err
	}

	if err := validate.MinLength("building_number", "body", m.BuildingNumber, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("building_number", "body", m.BuildingNumber, 16); err != nil {
		return err
	}

	return nil
}

func (m *SepaLiquidityAddress) validateCity(formats strfmt.Registry) error {

	if err := validate.RequiredString("city", "body", m.City); err != nil {
		return err
	}

	if err := validate.MinLength("city", "body", m.City, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", m.City, 35); err != nil {
		return err
	}

	return nil
}

func (m *SepaLiquidityAddress) validateCountry(formats strfmt.Registry) error {

	if err := validate.RequiredString("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.MinLength("country", "body", m.Country, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("country", "body", m.Country, 70); err != nil {
		return err
	}

	return nil
}

func (m *SepaLiquidityAddress) validateStreet(formats strfmt.Registry) error {

	if err := validate.RequiredString("street", "body", m.Street); err != nil {
		return err
	}

	if err := validate.MinLength("street", "body", m.Street, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("street", "body", m.Street, 70); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sepa liquidity address based on context it is used
func (m *SepaLiquidityAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SepaLiquidityAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaLiquidityAddress) UnmarshalBinary(b []byte) error {
	var res SepaLiquidityAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
