// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// VolumeSourceFromVolumeBackupDetailsAllOf1 volume source from volume backup details all of1
// swagger:discriminator volumeSourceFromVolumeBackupDetailsAllOf1 volumeBackup
type VolumeSourceFromVolumeBackupDetailsAllOf1 interface {
	runtime.Validatable

	// The OCID of the volume backup.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID() *string
	SetID(*string)
}

type volumeSourceFromVolumeBackupDetailsAllOf1 struct {
	idField *string
}

func (m *volumeSourceFromVolumeBackupDetailsAllOf1) ID() *string {
	return m.idField
}
func (m *volumeSourceFromVolumeBackupDetailsAllOf1) SetID(val *string) {
	m.idField = val
}

// Validate validates this volume source from volume backup details all of1
func (m *volumeSourceFromVolumeBackupDetailsAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *volumeSourceFromVolumeBackupDetailsAllOf1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID()), 255); err != nil {
		return err
	}

	return nil
}

// UnmarshalVolumeSourceFromVolumeBackupDetailsAllOf1Slice unmarshals polymorphic slices of VolumeSourceFromVolumeBackupDetailsAllOf1
func UnmarshalVolumeSourceFromVolumeBackupDetailsAllOf1Slice(reader io.Reader, consumer runtime.Consumer) ([]VolumeSourceFromVolumeBackupDetailsAllOf1, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []VolumeSourceFromVolumeBackupDetailsAllOf1
	for _, element := range elements {
		obj, err := unmarshalVolumeSourceFromVolumeBackupDetailsAllOf1(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalVolumeSourceFromVolumeBackupDetailsAllOf1 unmarshals polymorphic VolumeSourceFromVolumeBackupDetailsAllOf1
func UnmarshalVolumeSourceFromVolumeBackupDetailsAllOf1(reader io.Reader, consumer runtime.Consumer) (VolumeSourceFromVolumeBackupDetailsAllOf1, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalVolumeSourceFromVolumeBackupDetailsAllOf1(data, consumer)
}

func unmarshalVolumeSourceFromVolumeBackupDetailsAllOf1(data []byte, consumer runtime.Consumer) (VolumeSourceFromVolumeBackupDetailsAllOf1, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the volumeBackup property.
	var getType struct {
		VolumeBackup string `json:"volumeBackup"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("volumeBackup", "body", getType.VolumeBackup); err != nil {
		return nil, err
	}

	// The value of volumeBackup is used to determine which type to create and unmarshal the data into
	switch getType.VolumeBackup {
	case "VolumeSourceFromVolumeBackupDetails":
		var result VolumeSourceFromVolumeBackupDetails
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "volumeSourceFromVolumeBackupDetailsAllOf1":
		var result volumeSourceFromVolumeBackupDetailsAllOf1
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid volumeBackup value: %q", getType.VolumeBackup)

}
