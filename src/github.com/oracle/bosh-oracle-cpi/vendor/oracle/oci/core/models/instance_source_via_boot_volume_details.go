// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceSourceViaBootVolumeDetails instance source via boot volume details
// swagger:model InstanceSourceViaBootVolumeDetails
type InstanceSourceViaBootVolumeDetails struct {
	sourceTypeField string

	bootVolumeIdField *string
}

func (m *InstanceSourceViaBootVolumeDetails) SourceType() string {
	//return m.sourceTypeField
	return DiscriminatorTypeValues["InstanceSourceViaBootVolumeDetails"]
}
func (m *InstanceSourceViaBootVolumeDetails) SetSourceType(val string) {
	m.sourceTypeField = val
}

func (m *InstanceSourceViaBootVolumeDetails) BootVolumeID() *string {
	return m.bootVolumeIdField
}
func (m *InstanceSourceViaBootVolumeDetails) SetBootVolumeID(val *string) {
	m.bootVolumeIdField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *InstanceSourceViaBootVolumeDetails) UnmarshalJSON(raw []byte) error {
	var data struct {
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		SourceType string `json:"sourceType"`

		BootVolumeID *string `json:"bootVolumeId"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result InstanceSourceViaBootVolumeDetails

	if base.SourceType != result.SourceType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid sourceType value: %q", base.SourceType)
	}

	result.bootVolumeIdField = base.BootVolumeID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m InstanceSourceViaBootVolumeDetails) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		SourceType string `json:"sourceType"`

		BootVolumeID *string `json:"bootVolumeId"`
	}{

		SourceType: m.SourceType(),

		BootVolumeID: m.BootVolumeID(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this instance source via boot volume details
func (m *InstanceSourceViaBootVolumeDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceSourceViaBootVolumeDetails) validateSourceType(formats strfmt.Registry) error {

	if err := validate.RequiredString("sourceType", "body", string(m.SourceType())); err != nil {
		return err
	}

	return nil
}

func (m *InstanceSourceViaBootVolumeDetails) validateBootVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("bootVolumeId", "body", m.BootVolumeID()); err != nil {
		return err
	}

	if err := validate.MinLength("bootVolumeId", "body", string(*m.BootVolumeID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("bootVolumeId", "body", string(*m.BootVolumeID()), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceSourceViaBootVolumeDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceSourceViaBootVolumeDetails) UnmarshalBinary(b []byte) error {
	var res InstanceSourceViaBootVolumeDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
