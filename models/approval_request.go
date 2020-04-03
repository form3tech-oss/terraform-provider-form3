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

// ApprovalRequest approval request
//
// swagger:model ApprovalRequest
type ApprovalRequest struct {

	// attributes
	Attributes *ApprovalRequestAttributes `json:"attributes,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// type
	// Pattern: ^[A-Za-z]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this approval request
func (m *ApprovalRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApprovalRequest) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *ApprovalRequest) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequest) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequest) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequest) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApprovalRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApprovalRequest) UnmarshalBinary(b []byte) error {
	var res ApprovalRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApprovalRequestAttributes approval request attributes
//
// swagger:model ApprovalRequestAttributes
type ApprovalRequestAttributes struct {

	// action
	// Pattern: ^[A-Za-z]*$
	Action string `json:"action,omitempty"`

	// action time
	// Format: date-time
	ActionTime *strfmt.DateTime `json:"action_time,omitempty"`

	// actioned by
	// Format: uuid
	ActionedBy strfmt.UUID `json:"actioned_by,omitempty"`

	// after data
	AfterData interface{} `json:"after_data,omitempty"`

	// before data
	BeforeData interface{} `json:"before_data,omitempty"`

	// record id
	// Format: uuid
	RecordID strfmt.UUID `json:"record_id,omitempty"`

	// record orgid
	// Format: uuid
	RecordOrgid strfmt.UUID `json:"record_orgid,omitempty"`

	// record type
	// Pattern: ^[A-Za-z]*$
	RecordType string `json:"record_type,omitempty"`

	// record version
	// Minimum: 0
	RecordVersion *int64 `json:"record_version,omitempty"`

	// status
	// Pattern: ^[A-Za-z]*$
	Status string `json:"status,omitempty"`
}

// Validate validates this approval request attributes
func (m *ApprovalRequestAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordOrgid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApprovalRequestAttributes) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"action", "body", string(m.Action), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateActionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"action_time", "body", "date-time", m.ActionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateActionedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"actioned_by", "body", "uuid", m.ActionedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateRecordID(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordID) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"record_id", "body", "uuid", m.RecordID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateRecordOrgid(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordOrgid) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"record_orgid", "body", "uuid", m.RecordOrgid.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateRecordType(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordType) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"record_type", "body", string(m.RecordType), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateRecordVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordVersion) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"record_version", "body", int64(*m.RecordVersion), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"status", "body", string(m.Status), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApprovalRequestAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApprovalRequestAttributes) UnmarshalBinary(b []byte) error {
	var res ApprovalRequestAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
