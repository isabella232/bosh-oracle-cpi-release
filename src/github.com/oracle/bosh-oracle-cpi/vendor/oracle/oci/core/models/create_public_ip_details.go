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

// CreatePublicIPDetails create public Ip details
// swagger:model CreatePublicIpDetails
type CreatePublicIPDetails struct {

	// The OCID of the compartment to contain the public IP. For ephemeral public IPs,
	// you must set this to the private IP's compartment OCID.
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// Defines when the public IP is deleted and released back to the Oracle Cloud
	// Infrastructure public IP pool. For more information, see
	// [Public IP Addresses](/Content/Network/Tasks/managingpublicIPs.htm).
	//
	// Required: true
	Lifetime *string `json:"lifetime"`

	// The OCID of the private IP to assign the public IP to.
	//
	// Required for an ephemeral public IP because it must always be assigned to a private IP
	// (specifically a *primary* private IP).
	//
	// Optional for a reserved public IP. If you don't provide it, the public IP is created but not
	// assigned to a private IP. You can later assign the public IP with
	// [UpdatePublicIp](#/en/iaas/20160918/PublicIp/UpdatePublicIp).
	//
	// Max Length: 255
	// Min Length: 1
	PrivateIPID string `json:"privateIpId,omitempty"`
}

// Validate validates this create public Ip details
func (m *CreatePublicIPDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifetime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrivateIPID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePublicIPDetails) validateCompartmentID(formats strfmt.Registry) error {

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

func (m *CreatePublicIPDetails) validateDisplayName(formats strfmt.Registry) error {

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

var createPublicIpDetailsTypeLifetimePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EPHEMERAL","RESERVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createPublicIpDetailsTypeLifetimePropEnum = append(createPublicIpDetailsTypeLifetimePropEnum, v)
	}
}

const (

	// CreatePublicIPDetailsLifetimeEPHEMERAL captures enum value "EPHEMERAL"
	CreatePublicIPDetailsLifetimeEPHEMERAL string = "EPHEMERAL"

	// CreatePublicIPDetailsLifetimeRESERVED captures enum value "RESERVED"
	CreatePublicIPDetailsLifetimeRESERVED string = "RESERVED"
)

// prop value enum
func (m *CreatePublicIPDetails) validateLifetimeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createPublicIpDetailsTypeLifetimePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreatePublicIPDetails) validateLifetime(formats strfmt.Registry) error {

	if err := validate.Required("lifetime", "body", m.Lifetime); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifetimeEnum("lifetime", "body", *m.Lifetime); err != nil {
		return err
	}

	return nil
}

func (m *CreatePublicIPDetails) validatePrivateIPID(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateIPID) { // not required
		return nil
	}

	if err := validate.MinLength("privateIpId", "body", string(m.PrivateIPID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("privateIpId", "body", string(m.PrivateIPID), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePublicIPDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePublicIPDetails) UnmarshalBinary(b []byte) error {
	var res CreatePublicIPDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
