// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPSecConnectionDeviceStatus Status of the IPSec connection.
// swagger:model IPSecConnectionDeviceStatus
type IPSecConnectionDeviceStatus struct {

	// The OCID of the compartment containing the IPSec connection.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// The IPSec connection's Oracle ID (OCID).
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// The date and time the IPSec connection was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`

	// Two [TunnelStatus](#/en/iaas/20160918/TunnelStatus/) objects.
	Tunnels []*TunnelStatus `json:"tunnels"`
}

// Validate validates this IP sec connection device status
func (m *IPSecConnectionDeviceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTunnels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPSecConnectionDeviceStatus) validateCompartmentID(formats strfmt.Registry) error {

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

func (m *IPSecConnectionDeviceStatus) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 255); err != nil {
		return err
	}

	return nil
}

func (m *IPSecConnectionDeviceStatus) validateTimeCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPSecConnectionDeviceStatus) validateTunnels(formats strfmt.Registry) error {

	if swag.IsZero(m.Tunnels) { // not required
		return nil
	}

	for i := 0; i < len(m.Tunnels); i++ {

		if swag.IsZero(m.Tunnels[i]) { // not required
			continue
		}

		if m.Tunnels[i] != nil {

			if err := m.Tunnels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tunnels" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPSecConnectionDeviceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPSecConnectionDeviceStatus) UnmarshalBinary(b []byte) error {
	var res IPSecConnectionDeviceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
