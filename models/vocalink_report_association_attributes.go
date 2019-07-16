// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VocalinkReportAssociationAttributes vocalink report association attributes
// swagger:model VocalinkReportAssociationAttributes
type VocalinkReportAssociationAttributes struct {

	// bacs service user number
	BacsServiceUserNumber string `json:"bacs_service_user_number,omitempty"`
}

// Validate validates this vocalink report association attributes
func (m *VocalinkReportAssociationAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VocalinkReportAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VocalinkReportAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res VocalinkReportAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}