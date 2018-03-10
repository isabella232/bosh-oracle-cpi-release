package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// UpdateSubnetReader is a Reader for the UpdateSubnet structure.
type UpdateSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateSubnetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateSubnetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateSubnetPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateSubnetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSubnetOK creates a UpdateSubnetOK with default headers values
func NewUpdateSubnetOK() *UpdateSubnetOK {
	return &UpdateSubnetOK{}
}

/*UpdateSubnetOK handles this case with default header values.

The subnet was updated.
*/
type UpdateSubnetOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Subnet
}

func (o *UpdateSubnetOK) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnetId}][%d] updateSubnetOK  %+v", 200, o.Payload)
}

func (o *UpdateSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Subnet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubnetBadRequest creates a UpdateSubnetBadRequest with default headers values
func NewUpdateSubnetBadRequest() *UpdateSubnetBadRequest {
	return &UpdateSubnetBadRequest{}
}

/*UpdateSubnetBadRequest handles this case with default header values.

Bad Request
*/
type UpdateSubnetBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSubnetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnetId}][%d] updateSubnetBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubnetUnauthorized creates a UpdateSubnetUnauthorized with default headers values
func NewUpdateSubnetUnauthorized() *UpdateSubnetUnauthorized {
	return &UpdateSubnetUnauthorized{}
}

/*UpdateSubnetUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateSubnetUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSubnetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnetId}][%d] updateSubnetUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateSubnetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubnetNotFound creates a UpdateSubnetNotFound with default headers values
func NewUpdateSubnetNotFound() *UpdateSubnetNotFound {
	return &UpdateSubnetNotFound{}
}

/*UpdateSubnetNotFound handles this case with default header values.

Not Found
*/
type UpdateSubnetNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSubnetNotFound) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnetId}][%d] updateSubnetNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubnetConflict creates a UpdateSubnetConflict with default headers values
func NewUpdateSubnetConflict() *UpdateSubnetConflict {
	return &UpdateSubnetConflict{}
}

/*UpdateSubnetConflict handles this case with default header values.

Conflict
*/
type UpdateSubnetConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSubnetConflict) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnetId}][%d] updateSubnetConflict  %+v", 409, o.Payload)
}

func (o *UpdateSubnetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubnetPreconditionFailed creates a UpdateSubnetPreconditionFailed with default headers values
func NewUpdateSubnetPreconditionFailed() *UpdateSubnetPreconditionFailed {
	return &UpdateSubnetPreconditionFailed{}
}

/*UpdateSubnetPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateSubnetPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSubnetPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnetId}][%d] updateSubnetPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateSubnetPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubnetInternalServerError creates a UpdateSubnetInternalServerError with default headers values
func NewUpdateSubnetInternalServerError() *UpdateSubnetInternalServerError {
	return &UpdateSubnetInternalServerError{}
}

/*UpdateSubnetInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateSubnetInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnetId}][%d] updateSubnetInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubnetDefault creates a UpdateSubnetDefault with default headers values
func NewUpdateSubnetDefault(code int) *UpdateSubnetDefault {
	return &UpdateSubnetDefault{
		_statusCode: code,
	}
}

/*UpdateSubnetDefault handles this case with default header values.

An error has occurred.
*/
type UpdateSubnetDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update subnet default response
func (o *UpdateSubnetDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSubnetDefault) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnetId}][%d] UpdateSubnet default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSubnetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
