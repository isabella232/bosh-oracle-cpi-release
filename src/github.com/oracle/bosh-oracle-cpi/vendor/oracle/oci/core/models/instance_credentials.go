package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// InstanceCredentials The credentials for a particular instance.
// swagger:model InstanceCredentials
type InstanceCredentials struct {

	// The password for the username.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Password *string `json:"password"`

	// The username.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Username *string `json:"username"`
}

// Validate validates this instance credentials
func (m *InstanceCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceCredentials) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(*m.Password), 255); err != nil {
		return err
	}

	return nil
}

func (m *InstanceCredentials) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", string(*m.Username), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", string(*m.Username), 255); err != nil {
		return err
	}

	return nil
}
