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

// ExportImageDetails The destination details for the image export.
//
// Set `destinationType` to `objectStorageTuple`
// and use [ExportImageViaObjectStorageTupleDetails](#/en/iaas/latest/requests/ExportImageViaObjectStorageTupleDetails)
// when specifying the namespace, bucket name, and object name.
//
// Set `destinationType` to `objectStorageUri` and
// use [ExportImageViaObjectStorageUriDetails](#/en/iaas/latest/requests/ExportImageViaObjectStorageUriDetails)
// when specifying the Object Storage URL.
//
// swagger:discriminator ExportImageDetails destinationType
type ExportImageDetails interface {
	runtime.Validatable

	// The destination type. Use `objectStorageTuple` when specifying the namespace, bucket name, and object name.
	// Use `objectStorageUri` when specifying the Object Storage URL.
	//
	// Required: true
	DestinationType() string
	SetDestinationType(string)
}

type exportImageDetails struct {
	destinationTypeField string
}

func (m *exportImageDetails) DestinationType() string {
	return "ExportImageDetails"
}
func (m *exportImageDetails) SetDestinationType(val string) {

}

// Validate validates this export image details
func (m *exportImageDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// UnmarshalExportImageDetailsSlice unmarshals polymorphic slices of ExportImageDetails
func UnmarshalExportImageDetailsSlice(reader io.Reader, consumer runtime.Consumer) ([]ExportImageDetails, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ExportImageDetails
	for _, element := range elements {
		obj, err := unmarshalExportImageDetails(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalExportImageDetails unmarshals polymorphic ExportImageDetails
func UnmarshalExportImageDetails(reader io.Reader, consumer runtime.Consumer) (ExportImageDetails, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalExportImageDetails(data, consumer)
}

func unmarshalExportImageDetails(data []byte, consumer runtime.Consumer) (ExportImageDetails, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the destinationType property.
	var getType struct {
		DestinationType string `json:"destinationType"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("destinationType", "body", getType.DestinationType); err != nil {
		return nil, err
	}

	// The value of destinationType is used to determine which type to create and unmarshal the data into
	switch getType.DestinationType {
	case "ExportImageDetails":
		var result exportImageDetails
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case DiscriminatorTypeValues["ExportImageViaObjectStorageTupleDetails"]:
		var result ExportImageViaObjectStorageTupleDetails
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case DiscriminatorTypeValues["ExportImageViaObjectStorageUriDetails"]:
		var result ExportImageViaObjectStorageURIDetails
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid destinationType value: %q", getType.DestinationType)

}
