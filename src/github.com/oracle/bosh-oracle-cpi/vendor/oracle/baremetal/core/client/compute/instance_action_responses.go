package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/core/models"
)

// InstanceActionReader is a Reader for the InstanceAction structure.
type InstanceActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstanceActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInstanceActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewInstanceActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewInstanceActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewInstanceActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewInstanceActionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewInstanceActionPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewInstanceActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewInstanceActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInstanceActionOK creates a InstanceActionOK with default headers values
func NewInstanceActionOK() *InstanceActionOK {
	return &InstanceActionOK{}
}

/*InstanceActionOK handles this case with default header values.

The action is being performed.
*/
type InstanceActionOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Instance
}

func (o *InstanceActionOK) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceId}][%d] instanceActionOK  %+v", 200, o.Payload)
}

func (o *InstanceActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Instance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstanceActionBadRequest creates a InstanceActionBadRequest with default headers values
func NewInstanceActionBadRequest() *InstanceActionBadRequest {
	return &InstanceActionBadRequest{}
}

/*InstanceActionBadRequest handles this case with default header values.

Bad Request
*/
type InstanceActionBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *InstanceActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceId}][%d] instanceActionBadRequest  %+v", 400, o.Payload)
}

func (o *InstanceActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstanceActionUnauthorized creates a InstanceActionUnauthorized with default headers values
func NewInstanceActionUnauthorized() *InstanceActionUnauthorized {
	return &InstanceActionUnauthorized{}
}

/*InstanceActionUnauthorized handles this case with default header values.

Unauthorized
*/
type InstanceActionUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *InstanceActionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceId}][%d] instanceActionUnauthorized  %+v", 401, o.Payload)
}

func (o *InstanceActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstanceActionNotFound creates a InstanceActionNotFound with default headers values
func NewInstanceActionNotFound() *InstanceActionNotFound {
	return &InstanceActionNotFound{}
}

/*InstanceActionNotFound handles this case with default header values.

Not Found
*/
type InstanceActionNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *InstanceActionNotFound) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceId}][%d] instanceActionNotFound  %+v", 404, o.Payload)
}

func (o *InstanceActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstanceActionConflict creates a InstanceActionConflict with default headers values
func NewInstanceActionConflict() *InstanceActionConflict {
	return &InstanceActionConflict{}
}

/*InstanceActionConflict handles this case with default header values.

Conflict
*/
type InstanceActionConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *InstanceActionConflict) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceId}][%d] instanceActionConflict  %+v", 409, o.Payload)
}

func (o *InstanceActionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstanceActionPreconditionFailed creates a InstanceActionPreconditionFailed with default headers values
func NewInstanceActionPreconditionFailed() *InstanceActionPreconditionFailed {
	return &InstanceActionPreconditionFailed{}
}

/*InstanceActionPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type InstanceActionPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *InstanceActionPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceId}][%d] instanceActionPreconditionFailed  %+v", 412, o.Payload)
}

func (o *InstanceActionPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstanceActionInternalServerError creates a InstanceActionInternalServerError with default headers values
func NewInstanceActionInternalServerError() *InstanceActionInternalServerError {
	return &InstanceActionInternalServerError{}
}

/*InstanceActionInternalServerError handles this case with default header values.

Internal Server Error
*/
type InstanceActionInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *InstanceActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceId}][%d] instanceActionInternalServerError  %+v", 500, o.Payload)
}

func (o *InstanceActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstanceActionDefault creates a InstanceActionDefault with default headers values
func NewInstanceActionDefault(code int) *InstanceActionDefault {
	return &InstanceActionDefault{
		_statusCode: code,
	}
}

/*InstanceActionDefault handles this case with default header values.

An error has occurred.
*/
type InstanceActionDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the instance action default response
func (o *InstanceActionDefault) Code() int {
	return o._statusCode
}

func (o *InstanceActionDefault) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceId}][%d] InstanceAction default  %+v", o._statusCode, o.Payload)
}

func (o *InstanceActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
