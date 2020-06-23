// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MatchingCriteria matching criteria
//
// swagger:model MatchingCriteria
type MatchingCriteria struct {

	// close match threshold
	// Required: true
	CloseMatchThreshold *string `json:"close_match_threshold"`

	// exact match threshold
	// Required: true
	ExactMatchThreshold *string `json:"exact_match_threshold"`
}

// Validate validates this matching criteria
func (m *MatchingCriteria) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloseMatchThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExactMatchThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MatchingCriteria) validateCloseMatchThreshold(formats strfmt.Registry) error {

	if err := validate.Required("close_match_threshold", "body", m.CloseMatchThreshold); err != nil {
		return err
	}

	return nil
}

func (m *MatchingCriteria) validateExactMatchThreshold(formats strfmt.Registry) error {

	if err := validate.Required("exact_match_threshold", "body", m.ExactMatchThreshold); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MatchingCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchingCriteria) UnmarshalBinary(b []byte) error {
	var res MatchingCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
