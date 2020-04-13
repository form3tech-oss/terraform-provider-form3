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

// KeyAttributes key attributes
// swagger:model KeyAttributes
type KeyAttributes struct {

	// certificate signing request
	CertificateSigningRequest string `json:"certificate_signing_request,omitempty"`

	// curve
	// Enum: [PRIME256V1 SECP256R1 SECP384R1 SECP256K1]
	Curve string `json:"curve,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// private key
	PrivateKey string `json:"private_key,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// type
	// Enum: [RSA EC AES 3DES]
	Type *string `json:"type,omitempty"`
}

// Validate validates this key attributes
func (m *KeyAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurve(formats); err != nil {
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

var keyAttributesTypeCurvePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRIME256V1","SECP256R1","SECP384R1","SECP256K1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		keyAttributesTypeCurvePropEnum = append(keyAttributesTypeCurvePropEnum, v)
	}
}

const (

	// KeyAttributesCurvePRIME256V1 captures enum value "PRIME256V1"
	KeyAttributesCurvePRIME256V1 string = "PRIME256V1"

	// KeyAttributesCurveSECP256R1 captures enum value "SECP256R1"
	KeyAttributesCurveSECP256R1 string = "SECP256R1"

	// KeyAttributesCurveSECP384R1 captures enum value "SECP384R1"
	KeyAttributesCurveSECP384R1 string = "SECP384R1"

	// KeyAttributesCurveSECP256K1 captures enum value "SECP256K1"
	KeyAttributesCurveSECP256K1 string = "SECP256K1"
)

// prop value enum
func (m *KeyAttributes) validateCurveEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, keyAttributesTypeCurvePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KeyAttributes) validateCurve(formats strfmt.Registry) error {

	if swag.IsZero(m.Curve) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurveEnum("curve", "body", m.Curve); err != nil {
		return err
	}

	return nil
}

var keyAttributesTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RSA","EC","AES","3DES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		keyAttributesTypeTypePropEnum = append(keyAttributesTypeTypePropEnum, v)
	}
}

const (

	// KeyAttributesTypeRSA captures enum value "RSA"
	KeyAttributesTypeRSA string = "RSA"

	// KeyAttributesTypeEC captures enum value "EC"
	KeyAttributesTypeEC string = "EC"

	// KeyAttributesTypeAES captures enum value "AES"
	KeyAttributesTypeAES string = "AES"

	// KeyAttributesTypeNr3DES captures enum value "3DES"
	KeyAttributesTypeNr3DES string = "3DES"
)

// prop value enum
func (m *KeyAttributes) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, keyAttributesTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KeyAttributes) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyAttributes) UnmarshalBinary(b []byte) error {
	var res KeyAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
