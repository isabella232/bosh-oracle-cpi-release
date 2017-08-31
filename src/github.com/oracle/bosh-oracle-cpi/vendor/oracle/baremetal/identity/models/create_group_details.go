package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CreateGroupDetails create group details
// swagger:model CreateGroupDetails
type CreateGroupDetails struct {

	// The OCID of the tenancy containing the group.
	// Required: true
	CompartmentID *string `json:"compartmentId"`

	// The description you assign to the group during creation. Does not have to be unique, and it's changeable.
	// Required: true
	// Max Length: 400
	// Min Length: 1
	Description *string `json:"description"`

	// The name you assign to the group during creation. The name must be unique across all groups
	// in the tenancy and cannot be changed.
	//
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this create group details
func (m *CreateGroupDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGroupDetails) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	return nil
}

func (m *CreateGroupDetails) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 400); err != nil {
		return err
	}

	return nil
}

func (m *CreateGroupDetails) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}
