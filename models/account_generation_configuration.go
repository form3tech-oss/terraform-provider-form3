// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountGenerationConfiguration account generation configuration
// swagger:model AccountGenerationConfiguration
type AccountGenerationConfiguration struct {

	// bank id
	BankID string `json:"bank_id,omitempty"`

	// base currency
	BaseCurrency string `json:"base_currency,omitempty"`

	// bic
	Bic string `json:"bic,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// mod check enabled
	ModCheckEnabled bool `json:"mod_check_enabled,omitempty"`

	// valid account ranges
	ValidAccountRanges []*AccountGenerationConfigurationValidAccountRangesItems `json:"valid_account_ranges"`
}

// Validate validates this account generation configuration
func (m *AccountGenerationConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidAccountRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountGenerationConfiguration) validateValidAccountRanges(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidAccountRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidAccountRanges); i++ {
		if swag.IsZero(m.ValidAccountRanges[i]) { // not required
			continue
		}

		if m.ValidAccountRanges[i] != nil {
			if err := m.ValidAccountRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("valid_account_ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountGenerationConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountGenerationConfiguration) UnmarshalBinary(b []byte) error {
	var res AccountGenerationConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
