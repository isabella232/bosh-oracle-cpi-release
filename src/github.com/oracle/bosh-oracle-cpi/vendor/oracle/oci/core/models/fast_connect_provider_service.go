package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FastConnectProviderService A service offering from a supported provider. For more information,
// see [FastConnect Overview](/Content/Network/Concepts/fastconnect.htm).
//
// swagger:model FastConnectProviderService
type FastConnectProviderService struct {

	// A description of the service offered by the provider.
	//
	// Max Length: 255
	// Min Length: 1
	Description string `json:"description,omitempty"`

	// The name of the provider.
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ProviderName *string `json:"providerName"`

	// The name of the service offered by the provider.
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ProviderServiceName *string `json:"providerServiceName"`
}

// Validate validates this fast connect provider service
func (m *FastConnectProviderService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProviderName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProviderServiceName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FastConnectProviderService) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *FastConnectProviderService) validateProviderName(formats strfmt.Registry) error {

	if err := validate.Required("providerName", "body", m.ProviderName); err != nil {
		return err
	}

	if err := validate.MinLength("providerName", "body", string(*m.ProviderName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("providerName", "body", string(*m.ProviderName), 255); err != nil {
		return err
	}

	return nil
}

func (m *FastConnectProviderService) validateProviderServiceName(formats strfmt.Registry) error {

	if err := validate.Required("providerServiceName", "body", m.ProviderServiceName); err != nil {
		return err
	}

	if err := validate.MinLength("providerServiceName", "body", string(*m.ProviderServiceName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("providerServiceName", "body", string(*m.ProviderServiceName), 255); err != nil {
		return err
	}

	return nil
}
