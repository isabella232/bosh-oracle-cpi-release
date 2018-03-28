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

// IScsiVolumeAttachment An ISCSI volume attachment.
// swagger:model IScsiVolumeAttachment
type IScsiVolumeAttachment struct {
	attachmentTypeField string

	availabilityDomainField *string

	compartmentIdField *string

	displayNameField string

	idField *string

	instanceIdField *string

	isReadOnlyField bool

	lifecycleStateField *string

	timeCreatedField *strfmt.DateTime

	volumeIdField *string

	chapSecretField string

	chapUsernameField string

	ipv4Field *string

	iqnField *string

	portField *int64
}

func (m *IScsiVolumeAttachment) AttachmentType() string {
	//return m.attachmentTypeField
	return DiscriminatorTypeValues["IScsiVolumeAttachment"]
}
func (m *IScsiVolumeAttachment) SetAttachmentType(val string) {
	m.attachmentTypeField = val
}

func (m *IScsiVolumeAttachment) AvailabilityDomain() *string {
	return m.availabilityDomainField
}
func (m *IScsiVolumeAttachment) SetAvailabilityDomain(val *string) {
	m.availabilityDomainField = val
}

func (m *IScsiVolumeAttachment) CompartmentID() *string {
	return m.compartmentIdField
}
func (m *IScsiVolumeAttachment) SetCompartmentID(val *string) {
	m.compartmentIdField = val
}

func (m *IScsiVolumeAttachment) DisplayName() string {
	return m.displayNameField
}
func (m *IScsiVolumeAttachment) SetDisplayName(val string) {
	m.displayNameField = val
}

func (m *IScsiVolumeAttachment) ID() *string {
	return m.idField
}
func (m *IScsiVolumeAttachment) SetID(val *string) {
	m.idField = val
}

func (m *IScsiVolumeAttachment) InstanceID() *string {
	return m.instanceIdField
}
func (m *IScsiVolumeAttachment) SetInstanceID(val *string) {
	m.instanceIdField = val
}

func (m *IScsiVolumeAttachment) IsReadOnly() bool {
	return m.isReadOnlyField
}
func (m *IScsiVolumeAttachment) SetIsReadOnly(val bool) {
	m.isReadOnlyField = val
}

func (m *IScsiVolumeAttachment) LifecycleState() *string {
	return m.lifecycleStateField
}
func (m *IScsiVolumeAttachment) SetLifecycleState(val *string) {
	m.lifecycleStateField = val
}

func (m *IScsiVolumeAttachment) TimeCreated() *strfmt.DateTime {
	return m.timeCreatedField
}
func (m *IScsiVolumeAttachment) SetTimeCreated(val *strfmt.DateTime) {
	m.timeCreatedField = val
}

func (m *IScsiVolumeAttachment) VolumeID() *string {
	return m.volumeIdField
}
func (m *IScsiVolumeAttachment) SetVolumeID(val *string) {
	m.volumeIdField = val
}

func (m *IScsiVolumeAttachment) ChapSecret() string {
	return m.chapSecretField
}
func (m *IScsiVolumeAttachment) SetChapSecret(val string) {
	m.chapSecretField = val
}

func (m *IScsiVolumeAttachment) ChapUsername() string {
	return m.chapUsernameField
}
func (m *IScsiVolumeAttachment) SetChapUsername(val string) {
	m.chapUsernameField = val
}

func (m *IScsiVolumeAttachment) IPV4() *string {
	return m.ipv4Field
}
func (m *IScsiVolumeAttachment) SetIPV4(val *string) {
	m.ipv4Field = val
}

func (m *IScsiVolumeAttachment) Iqn() *string {
	return m.iqnField
}
func (m *IScsiVolumeAttachment) SetIqn(val *string) {
	m.iqnField = val
}

func (m *IScsiVolumeAttachment) Port() *int64 {
	return m.portField
}
func (m *IScsiVolumeAttachment) SetPort(val *int64) {
	m.portField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *IScsiVolumeAttachment) UnmarshalJSON(raw []byte) error {
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

		AttachmentType string `json:"attachmentType"`

		AvailabilityDomain *string `json:"availabilityDomain"`

		CompartmentID *string `json:"compartmentId"`

		DisplayName string `json:"displayName,omitempty"`

		ID *string `json:"id"`

		InstanceID *string `json:"instanceId"`

		IsReadOnly bool `json:"isReadOnly,omitempty"`

		LifecycleState *string `json:"lifecycleState"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`

		VolumeID *string `json:"volumeId"`

		ChapSecret string `json:"chapSecret,omitempty"`

		ChapUsername string `json:"chapUsername,omitempty"`

		IPV4 *string `json:"ipv4"`

		Iqn *string `json:"iqn"`

		Port *int64 `json:"port"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result IScsiVolumeAttachment

	if base.AttachmentType != result.AttachmentType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid attachmentType value: %q", base.AttachmentType)
	}

	result.availabilityDomainField = base.AvailabilityDomain

	result.compartmentIdField = base.CompartmentID

	result.displayNameField = base.DisplayName

	result.idField = base.ID

	result.instanceIdField = base.InstanceID

	result.isReadOnlyField = base.IsReadOnly

	result.lifecycleStateField = base.LifecycleState

	result.timeCreatedField = base.TimeCreated

	result.volumeIdField = base.VolumeID

	result.chapSecretField = base.ChapSecret

	result.chapUsernameField = base.ChapUsername

	result.ipv4Field = base.IPV4

	result.iqnField = base.Iqn

	result.portField = base.Port

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m IScsiVolumeAttachment) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AttachmentType string `json:"attachmentType"`

		AvailabilityDomain *string `json:"availabilityDomain"`

		CompartmentID *string `json:"compartmentId"`

		DisplayName string `json:"displayName,omitempty"`

		ID *string `json:"id"`

		InstanceID *string `json:"instanceId"`

		IsReadOnly bool `json:"isReadOnly,omitempty"`

		LifecycleState *string `json:"lifecycleState"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`

		VolumeID *string `json:"volumeId"`

		ChapSecret string `json:"chapSecret,omitempty"`

		ChapUsername string `json:"chapUsername,omitempty"`

		IPV4 *string `json:"ipv4"`

		Iqn *string `json:"iqn"`

		Port *int64 `json:"port"`
	}{

		AttachmentType: m.AttachmentType(),

		AvailabilityDomain: m.AvailabilityDomain(),

		CompartmentID: m.CompartmentID(),

		DisplayName: m.DisplayName(),

		ID: m.ID(),

		InstanceID: m.InstanceID(),

		IsReadOnly: m.IsReadOnly(),

		LifecycleState: m.LifecycleState(),

		TimeCreated: m.TimeCreated(),

		VolumeID: m.VolumeID(),

		ChapSecret: m.ChapSecret(),

		ChapUsername: m.ChapUsername(),

		IPV4: m.IPV4(),

		Iqn: m.Iqn(),

		Port: m.Port(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this i scsi volume attachment
func (m *IScsiVolumeAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailabilityDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompartmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIqn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IScsiVolumeAttachment) validateAttachmentType(formats strfmt.Registry) error {

	if err := validate.RequiredString("attachmentType", "body", string(m.AttachmentType())); err != nil {
		return err
	}

	if err := validate.MinLength("attachmentType", "body", string(m.AttachmentType()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("attachmentType", "body", string(m.AttachmentType()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateAvailabilityDomain(formats strfmt.Registry) error {

	if err := validate.Required("availabilityDomain", "body", m.AvailabilityDomain()); err != nil {
		return err
	}

	if err := validate.MinLength("availabilityDomain", "body", string(*m.AvailabilityDomain()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("availabilityDomain", "body", string(*m.AvailabilityDomain()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID()); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName()) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateID(formats strfmt.Registry) error {

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

func (m *IScsiVolumeAttachment) validateInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("instanceId", "body", m.InstanceID()); err != nil {
		return err
	}

	if err := validate.MinLength("instanceId", "body", string(*m.InstanceID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("instanceId", "body", string(*m.InstanceID()), 255); err != nil {
		return err
	}

	return nil
}

var iScsiVolumeAttachmentTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ATTACHING","ATTACHED","DETACHING","DETACHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iScsiVolumeAttachmentTypeLifecycleStatePropEnum = append(iScsiVolumeAttachmentTypeLifecycleStatePropEnum, v)
	}
}

// property enum
func (m *IScsiVolumeAttachment) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iScsiVolumeAttachmentTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IScsiVolumeAttachment) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState()); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState()); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("timeCreated", "body", m.TimeCreated()); err != nil {
		return err
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volumeId", "body", m.VolumeID()); err != nil {
		return err
	}

	if err := validate.MinLength("volumeId", "body", string(*m.VolumeID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("volumeId", "body", string(*m.VolumeID()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateIPV4(formats strfmt.Registry) error {

	if err := validate.Required("ipv4", "body", m.IPV4()); err != nil {
		return err
	}

	if err := validate.MinLength("ipv4", "body", string(*m.IPV4()), 7); err != nil {
		return err
	}

	if err := validate.MaxLength("ipv4", "body", string(*m.IPV4()), 15); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateIqn(formats strfmt.Registry) error {

	if err := validate.Required("iqn", "body", m.Iqn()); err != nil {
		return err
	}

	if err := validate.MinLength("iqn", "body", string(*m.Iqn()), 1); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port()); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port()), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port()), 65535, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IScsiVolumeAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IScsiVolumeAttachment) UnmarshalBinary(b []byte) error {
	var res IScsiVolumeAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
