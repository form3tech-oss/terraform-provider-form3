// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ValidationStatus validation status
//
// swagger:model ValidationStatus
type ValidationStatus string

const (

	// ValidationStatusFailed captures enum value "failed"
	ValidationStatusFailed ValidationStatus = "failed"

	// ValidationStatusPassed captures enum value "passed"
	ValidationStatusPassed ValidationStatus = "passed"

	// ValidationStatusNotAccepted captures enum value "not_accepted"
	ValidationStatusNotAccepted ValidationStatus = "not_accepted"
)

// for schema
var validationStatusEnum []interface{}

func init() {
	var res []ValidationStatus
	if err := json.Unmarshal([]byte(`["failed","passed","not_accepted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		validationStatusEnum = append(validationStatusEnum, v)
	}
}

func (m ValidationStatus) validateValidationStatusEnum(path, location string, value ValidationStatus) error {
	if err := validate.Enum(path, location, value, validationStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this validation status
func (m ValidationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateValidationStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
