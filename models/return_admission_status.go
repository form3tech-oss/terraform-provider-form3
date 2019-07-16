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

// ReturnAdmissionStatus return admission status
// swagger:model ReturnAdmissionStatus
type ReturnAdmissionStatus string

const (

	// ReturnAdmissionStatusConfirmed captures enum value "confirmed"
	ReturnAdmissionStatusConfirmed ReturnAdmissionStatus = "confirmed"

	// ReturnAdmissionStatusFailed captures enum value "failed"
	ReturnAdmissionStatusFailed ReturnAdmissionStatus = "failed"
)

// for schema
var returnAdmissionStatusEnum []interface{}

func init() {
	var res []ReturnAdmissionStatus
	if err := json.Unmarshal([]byte(`["confirmed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnAdmissionStatusEnum = append(returnAdmissionStatusEnum, v)
	}
}

func (m ReturnAdmissionStatus) validateReturnAdmissionStatusEnum(path, location string, value ReturnAdmissionStatus) error {
	if err := validate.Enum(path, location, value, returnAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this return admission status
func (m ReturnAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReturnAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}