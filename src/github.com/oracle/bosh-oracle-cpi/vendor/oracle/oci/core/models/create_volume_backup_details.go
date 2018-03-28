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

// CreateVolumeBackupDetails create volume backup details
// swagger:model CreateVolumeBackupDetails
type CreateVolumeBackupDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Operations": {"CostCenter": "42"}}`
	//
	DefinedTags map[string]map[string]interface{} `json:"definedTags,omitempty"`

	// A user-friendly name for the volume backup. Does not have to be unique and it's changeable.
	// Avoid entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Department": "Finance"}`
	//
	FreeformTags map[string]string `json:"freeformTags,omitempty"`

	// The type of backup to create. If omitted, defaults to INCREMENTAL.
	Type string `json:"type,omitempty"`

	// The OCID of the volume that needs to be backed up.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	VolumeID *string `json:"volumeId"`
}

// Validate validates this create volume backup details
func (m *CreateVolumeBackupDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVolumeBackupDetails) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

var createVolumeBackupDetailsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FULL","INCREMENTAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createVolumeBackupDetailsTypeTypePropEnum = append(createVolumeBackupDetailsTypeTypePropEnum, v)
	}
}

const (

	// CreateVolumeBackupDetailsTypeFULL captures enum value "FULL"
	CreateVolumeBackupDetailsTypeFULL string = "FULL"

	// CreateVolumeBackupDetailsTypeINCREMENTAL captures enum value "INCREMENTAL"
	CreateVolumeBackupDetailsTypeINCREMENTAL string = "INCREMENTAL"
)

// prop value enum
func (m *CreateVolumeBackupDetails) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createVolumeBackupDetailsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateVolumeBackupDetails) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CreateVolumeBackupDetails) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volumeId", "body", m.VolumeID); err != nil {
		return err
	}

	if err := validate.MinLength("volumeId", "body", string(*m.VolumeID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("volumeId", "body", string(*m.VolumeID), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateVolumeBackupDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateVolumeBackupDetails) UnmarshalBinary(b []byte) error {
	var res CreateVolumeBackupDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
