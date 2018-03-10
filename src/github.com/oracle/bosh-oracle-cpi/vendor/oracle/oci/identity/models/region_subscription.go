package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RegionSubscription region subscription
// swagger:model RegionSubscription
type RegionSubscription struct {

	// Indicates the region is home region or not.
	// Required: true
	IsHomeRegion *bool `json:"isHomeRegion"`

	// The key of the region such as PHX, IAD.
	// Required: true
	// Max Length: 16
	// Min Length: 1
	RegionKey *string `json:"regionKey"`

	// The name of the region such as us-phoenix-1.
	// Required: true
	// Max Length: 16
	// Min Length: 1
	RegionName *string `json:"regionName"`

	// The region subscription status such as Ready, InProgress
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this region subscription
func (m *RegionSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsHomeRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegionKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegionName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegionSubscription) validateIsHomeRegion(formats strfmt.Registry) error {

	if err := validate.Required("isHomeRegion", "body", m.IsHomeRegion); err != nil {
		return err
	}

	return nil
}

func (m *RegionSubscription) validateRegionKey(formats strfmt.Registry) error {

	if err := validate.Required("regionKey", "body", m.RegionKey); err != nil {
		return err
	}

	if err := validate.MinLength("regionKey", "body", string(*m.RegionKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("regionKey", "body", string(*m.RegionKey), 16); err != nil {
		return err
	}

	return nil
}

func (m *RegionSubscription) validateRegionName(formats strfmt.Registry) error {

	if err := validate.Required("regionName", "body", m.RegionName); err != nil {
		return err
	}

	if err := validate.MinLength("regionName", "body", string(*m.RegionName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("regionName", "body", string(*m.RegionName), 16); err != nil {
		return err
	}

	return nil
}

var regionSubscriptionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["READY","IN_PROGRESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		regionSubscriptionTypeStatusPropEnum = append(regionSubscriptionTypeStatusPropEnum, v)
	}
}

const (
	// RegionSubscriptionStatusREADY captures enum value "READY"
	RegionSubscriptionStatusREADY string = "READY"
	// RegionSubscriptionStatusINPROGRESS captures enum value "IN_PROGRESS"
	RegionSubscriptionStatusINPROGRESS string = "IN_PROGRESS"
)

// prop value enum
func (m *RegionSubscription) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, regionSubscriptionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RegionSubscription) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}
