package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/core/models"
)

// DeleteCpeReader is a Reader for the DeleteCpe structure.
type DeleteCpeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCpeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCpeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteCpeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCpeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteCpePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteCpeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteCpeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCpeNoContent creates a DeleteCpeNoContent with default headers values
func NewDeleteCpeNoContent() *DeleteCpeNoContent {
	return &DeleteCpeNoContent{}
}

/*DeleteCpeNoContent handles this case with default header values.

The CPE is being deleted.
*/
type DeleteCpeNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeleteCpeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cpes/{cpeId}][%d] deleteCpeNoContent ", 204)
}

func (o *DeleteCpeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeleteCpeUnauthorized creates a DeleteCpeUnauthorized with default headers values
func NewDeleteCpeUnauthorized() *DeleteCpeUnauthorized {
	return &DeleteCpeUnauthorized{}
}

/*DeleteCpeUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCpeUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteCpeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cpes/{cpeId}][%d] deleteCpeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCpeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCpeNotFound creates a DeleteCpeNotFound with default headers values
func NewDeleteCpeNotFound() *DeleteCpeNotFound {
	return &DeleteCpeNotFound{}
}

/*DeleteCpeNotFound handles this case with default header values.

Not Found
*/
type DeleteCpeNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteCpeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /cpes/{cpeId}][%d] deleteCpeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCpeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCpePreconditionFailed creates a DeleteCpePreconditionFailed with default headers values
func NewDeleteCpePreconditionFailed() *DeleteCpePreconditionFailed {
	return &DeleteCpePreconditionFailed{}
}

/*DeleteCpePreconditionFailed handles this case with default header values.

Precondition Failed
*/
type DeleteCpePreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteCpePreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /cpes/{cpeId}][%d] deleteCpePreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteCpePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCpeInternalServerError creates a DeleteCpeInternalServerError with default headers values
func NewDeleteCpeInternalServerError() *DeleteCpeInternalServerError {
	return &DeleteCpeInternalServerError{}
}

/*DeleteCpeInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteCpeInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteCpeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cpes/{cpeId}][%d] deleteCpeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCpeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCpeDefault creates a DeleteCpeDefault with default headers values
func NewDeleteCpeDefault(code int) *DeleteCpeDefault {
	return &DeleteCpeDefault{
		_statusCode: code,
	}
}

/*DeleteCpeDefault handles this case with default header values.

An error has occurred.
*/
type DeleteCpeDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete cpe default response
func (o *DeleteCpeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCpeDefault) Error() string {
	return fmt.Sprintf("[DELETE /cpes/{cpeId}][%d] DeleteCpe default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCpeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
