// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SigningKeysResourceType signing keys resource type
// swagger:model SigningKeysResourceType
type SigningKeysResourceType string

const (

	// SigningKeysResourceTypeSigningKeys captures enum value "signing_keys"
	SigningKeysResourceTypeSigningKeys SigningKeysResourceType = "signing_keys"
)

// for schema
var signingKeysResourceTypeEnum []interface{}

func init() {
	var res []SigningKeysResourceType
	if err := json.Unmarshal([]byte(`["signing_keys"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		signingKeysResourceTypeEnum = append(signingKeysResourceTypeEnum, v)
	}
}

func (m SigningKeysResourceType) validateSigningKeysResourceTypeEnum(path, location string, value SigningKeysResourceType) error {
	if err := validate.Enum(path, location, value, signingKeysResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this signing keys resource type
func (m SigningKeysResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSigningKeysResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
