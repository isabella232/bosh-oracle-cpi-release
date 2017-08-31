package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateImageDetails create image details
// swagger:model CreateImageDetails
type CreateImageDetails struct {

	// The OCID of the compartment containing the instance you want to use as the basis for the image.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// A user-friendly name for the image. It does not have to be unique, and it's changeable. You cannot use an
	// Oracle-provided image name as a custom image name.
	//
	// Example: `My Oracle Linux image`
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// The OCID of the instance you want to use as the basis for the image.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	InstanceID *string `json:"instanceId"`
}

// Validate validates this create image details
func (m *CreateImageDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateImageDetails) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageDetails) validateDisplayName(formats strfmt.Registry) error {

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

func (m *CreateImageDetails) validateInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("instanceId", "body", m.InstanceID); err != nil {
		return err
	}

	if err := validate.MinLength("instanceId", "body", string(*m.InstanceID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("instanceId", "body", string(*m.InstanceID), 255); err != nil {
		return err
	}

	return nil
}
