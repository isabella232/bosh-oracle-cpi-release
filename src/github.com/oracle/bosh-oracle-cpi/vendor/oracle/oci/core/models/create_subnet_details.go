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

// CreateSubnetDetails create subnet details
// swagger:model CreateSubnetDetails
type CreateSubnetDetails struct {

	// The Availability Domain to contain the subnet.
	//
	// Example: `Uocm:PHX-AD-1`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	AvailabilityDomain *string `json:"availabilityDomain"`

	// The CIDR IP address range of the subnet.
	//
	// Example: `172.16.1.0/24`
	//
	// Required: true
	// Max Length: 32
	// Min Length: 1
	CidrBlock *string `json:"cidrBlock"`

	// The OCID of the compartment to contain the subnet.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// The OCID of the set of DHCP options the subnet will use. If you don't
	// provide a value, the subnet will use the VCN's default set of DHCP options.
	//
	// Max Length: 255
	// Min Length: 1
	DhcpOptionsID string `json:"dhcpOptionsId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// A DNS label for the subnet, used in conjunction with the VNIC's hostname and
	// VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (e.g., `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be an alphanumeric string that begins with a letter and is unique within the VCN.
	// The value cannot be changed.
	//
	// This value must be set if you want to use the Internet and VCN Resolver to resolve the
	// hostnames of instances in the subnet. It can only be set if the VCN itself
	// was created with a DNS label.
	//
	// For more information, see
	// [DNS in Your Virtual Cloud Network](/Content/Network/Concepts/dns.htm).
	//
	// Example: `subnet123`
	//
	// Max Length: 15
	// Min Length: 1
	DNSLabel string `json:"dnsLabel,omitempty"`

	// Whether VNICs within this subnet can have public IP addresses.
	// Defaults to false, which means VNICs created in this subnet will
	// automatically be assigned public IP addresses unless specified
	// otherwise during instance launch or VNIC creation (with the
	// `assignPublicIp` flag in [CreateVnicDetails](#/en/iaas/20160918/CreateVnicDetails/)).
	// If `prohibitPublicIpOnVnic` is set to true, VNICs created in this
	// subnet cannot have public IP addresses (i.e., it's a private
	// subnet).
	//
	// Example: `true`
	//
	ProhibitPublicIPOnVnic bool `json:"prohibitPublicIpOnVnic,omitempty"`

	// The OCID of the route table the subnet will use. If you don't provide a value,
	// the subnet will use the VCN's default route table.
	//
	// Max Length: 255
	// Min Length: 1
	RouteTableID string `json:"routeTableId,omitempty"`

	// OCIDs for the security lists to associate with the subnet. If you don't
	// provide a value, the VCN's default security list will be associated with
	// the subnet. Remember that security lists are associated at the subnet
	// level, but the rules are applied to the individual VNICs in the subnet.
	//
	SecurityListIds []string `json:"securityListIds"`

	// The OCID of the VCN to contain the subnet.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	VcnID *string `json:"vcnId"`
}

// Validate validates this create subnet details
func (m *CreateSubnetDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCidrBlock(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDhcpOptionsID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDNSLabel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRouteTableID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecurityListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVcnID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSubnetDetails) validateAvailabilityDomain(formats strfmt.Registry) error {

	if err := validate.Required("availabilityDomain", "body", m.AvailabilityDomain); err != nil {
		return err
	}

	if err := validate.MinLength("availabilityDomain", "body", string(*m.AvailabilityDomain), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("availabilityDomain", "body", string(*m.AvailabilityDomain), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubnetDetails) validateCidrBlock(formats strfmt.Registry) error {

	if err := validate.Required("cidrBlock", "body", m.CidrBlock); err != nil {
		return err
	}

	if err := validate.MinLength("cidrBlock", "body", string(*m.CidrBlock), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("cidrBlock", "body", string(*m.CidrBlock), 32); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubnetDetails) validateCompartmentID(formats strfmt.Registry) error {

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

func (m *CreateSubnetDetails) validateDhcpOptionsID(formats strfmt.Registry) error {

	if swag.IsZero(m.DhcpOptionsID) { // not required
		return nil
	}

	if err := validate.MinLength("dhcpOptionsId", "body", string(m.DhcpOptionsID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("dhcpOptionsId", "body", string(m.DhcpOptionsID), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubnetDetails) validateDisplayName(formats strfmt.Registry) error {

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

func (m *CreateSubnetDetails) validateDNSLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSLabel) { // not required
		return nil
	}

	if err := validate.MinLength("dnsLabel", "body", string(m.DNSLabel), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("dnsLabel", "body", string(m.DNSLabel), 15); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubnetDetails) validateRouteTableID(formats strfmt.Registry) error {

	if swag.IsZero(m.RouteTableID) { // not required
		return nil
	}

	if err := validate.MinLength("routeTableId", "body", string(m.RouteTableID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("routeTableId", "body", string(m.RouteTableID), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubnetDetails) validateSecurityListIds(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityListIds) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityListIds); i++ {

		if err := validate.MinLength("securityListIds"+"."+strconv.Itoa(i), "body", string(m.SecurityListIds[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("securityListIds"+"."+strconv.Itoa(i), "body", string(m.SecurityListIds[i]), 255); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateSubnetDetails) validateVcnID(formats strfmt.Registry) error {

	if err := validate.Required("vcnId", "body", m.VcnID); err != nil {
		return err
	}

	if err := validate.MinLength("vcnId", "body", string(*m.VcnID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("vcnId", "body", string(*m.VcnID), 255); err != nil {
		return err
	}

	return nil
}
