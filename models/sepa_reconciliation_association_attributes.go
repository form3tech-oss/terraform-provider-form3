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

// SepaReconciliationAssociationAttributes sepa reconciliation association attributes
//
// swagger:model SepaReconciliationAssociationAttributes
type SepaReconciliationAssociationAttributes struct {

	// address
	// Required: true
	Address SepaReconciliationAssociationAttributesAddress `json:"address"`

	// name
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// reconciliation bic
	// Required: true
	// Min Length: 1
	ReconciliationBic string `json:"reconciliation_bic"`

	// reconciliation iban
	// Required: true
	// Min Length: 1
	ReconciliationIban string `json:"reconciliation_iban"`

	// technical bic
	// Required: true
	// Min Length: 1
	TechnicalBic string `json:"technical_bic"`
}

// Validate validates this sepa reconciliation association attributes
func (m *SepaReconciliationAssociationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReconciliationBic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReconciliationIban(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTechnicalBic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaReconciliationAssociationAttributes) validateAddress(formats strfmt.Registry) error {

	if err := m.Address.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("address")
		}
		return err
	}

	return nil
}

func (m *SepaReconciliationAssociationAttributes) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *SepaReconciliationAssociationAttributes) validateReconciliationBic(formats strfmt.Registry) error {

	if err := validate.RequiredString("reconciliation_bic", "body", m.ReconciliationBic); err != nil {
		return err
	}

	if err := validate.MinLength("reconciliation_bic", "body", m.ReconciliationBic, 1); err != nil {
		return err
	}

	return nil
}

func (m *SepaReconciliationAssociationAttributes) validateReconciliationIban(formats strfmt.Registry) error {

	if err := validate.RequiredString("reconciliation_iban", "body", m.ReconciliationIban); err != nil {
		return err
	}

	if err := validate.MinLength("reconciliation_iban", "body", m.ReconciliationIban, 1); err != nil {
		return err
	}

	return nil
}

func (m *SepaReconciliationAssociationAttributes) validateTechnicalBic(formats strfmt.Registry) error {

	if err := validate.RequiredString("technical_bic", "body", m.TechnicalBic); err != nil {
		return err
	}

	if err := validate.MinLength("technical_bic", "body", m.TechnicalBic, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sepa reconciliation association attributes based on the context it is used
func (m *SepaReconciliationAssociationAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaReconciliationAssociationAttributes) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Address.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("address")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepaReconciliationAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaReconciliationAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res SepaReconciliationAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SepaReconciliationAssociationAttributesAddress sepa reconciliation association attributes address
//
// swagger:model SepaReconciliationAssociationAttributesAddress
type SepaReconciliationAssociationAttributesAddress struct {

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

// Validate validates this sepa reconciliation association attributes address
func (m *SepaReconciliationAssociationAttributesAddress) Validate(formats strfmt.Registry) error {
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

func (m *SepaReconciliationAssociationAttributesAddress) validateBuildingNumber(formats strfmt.Registry) error {

	if err := validate.RequiredString("address"+"."+"building_number", "body", m.BuildingNumber); err != nil {
		return err
	}

	if err := validate.MinLength("address"+"."+"building_number", "body", m.BuildingNumber, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("address"+"."+"building_number", "body", m.BuildingNumber, 16); err != nil {
		return err
	}

	return nil
}

func (m *SepaReconciliationAssociationAttributesAddress) validateCity(formats strfmt.Registry) error {

	if err := validate.RequiredString("address"+"."+"city", "body", m.City); err != nil {
		return err
	}

	if err := validate.MinLength("address"+"."+"city", "body", m.City, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("address"+"."+"city", "body", m.City, 35); err != nil {
		return err
	}

	return nil
}

func (m *SepaReconciliationAssociationAttributesAddress) validateCountry(formats strfmt.Registry) error {

	if err := validate.RequiredString("address"+"."+"country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.MinLength("address"+"."+"country", "body", m.Country, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("address"+"."+"country", "body", m.Country, 70); err != nil {
		return err
	}

	return nil
}

func (m *SepaReconciliationAssociationAttributesAddress) validateStreet(formats strfmt.Registry) error {

	if err := validate.RequiredString("address"+"."+"street", "body", m.Street); err != nil {
		return err
	}

	if err := validate.MinLength("address"+"."+"street", "body", m.Street, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("address"+"."+"street", "body", m.Street, 70); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sepa reconciliation association attributes address based on context it is used
func (m *SepaReconciliationAssociationAttributesAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SepaReconciliationAssociationAttributesAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaReconciliationAssociationAttributesAddress) UnmarshalBinary(b []byte) error {
	var res SepaReconciliationAssociationAttributesAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
