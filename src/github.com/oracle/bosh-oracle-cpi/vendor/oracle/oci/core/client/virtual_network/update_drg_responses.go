// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// UpdateDrgReader is a Reader for the UpdateDrg structure.
type UpdateDrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateDrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateDrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateDrgNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateDrgConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateDrgPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateDrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateDrgDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDrgOK creates a UpdateDrgOK with default headers values
func NewUpdateDrgOK() *UpdateDrgOK {
	return &UpdateDrgOK{}
}

/*UpdateDrgOK handles this case with default header values.

The DRG was updated.
*/
type UpdateDrgOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Drg
}

func (o *UpdateDrgOK) Error() string {
	return fmt.Sprintf("[PUT /drgs/{drgId}][%d] updateDrgOK  %+v", 200, o.Payload)
}

func (o *UpdateDrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Drg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDrgBadRequest creates a UpdateDrgBadRequest with default headers values
func NewUpdateDrgBadRequest() *UpdateDrgBadRequest {
	return &UpdateDrgBadRequest{}
}

/*UpdateDrgBadRequest handles this case with default header values.

Bad Request
*/
type UpdateDrgBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDrgBadRequest) Error() string {
	return fmt.Sprintf("[PUT /drgs/{drgId}][%d] updateDrgBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDrgUnauthorized creates a UpdateDrgUnauthorized with default headers values
func NewUpdateDrgUnauthorized() *UpdateDrgUnauthorized {
	return &UpdateDrgUnauthorized{}
}

/*UpdateDrgUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateDrgUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDrgUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /drgs/{drgId}][%d] updateDrgUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateDrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDrgNotFound creates a UpdateDrgNotFound with default headers values
func NewUpdateDrgNotFound() *UpdateDrgNotFound {
	return &UpdateDrgNotFound{}
}

/*UpdateDrgNotFound handles this case with default header values.

Not Found
*/
type UpdateDrgNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDrgNotFound) Error() string {
	return fmt.Sprintf("[PUT /drgs/{drgId}][%d] updateDrgNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDrgNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDrgConflict creates a UpdateDrgConflict with default headers values
func NewUpdateDrgConflict() *UpdateDrgConflict {
	return &UpdateDrgConflict{}
}

/*UpdateDrgConflict handles this case with default header values.

Conflict
*/
type UpdateDrgConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDrgConflict) Error() string {
	return fmt.Sprintf("[PUT /drgs/{drgId}][%d] updateDrgConflict  %+v", 409, o.Payload)
}

func (o *UpdateDrgConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDrgPreconditionFailed creates a UpdateDrgPreconditionFailed with default headers values
func NewUpdateDrgPreconditionFailed() *UpdateDrgPreconditionFailed {
	return &UpdateDrgPreconditionFailed{}
}

/*UpdateDrgPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateDrgPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDrgPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /drgs/{drgId}][%d] updateDrgPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateDrgPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDrgInternalServerError creates a UpdateDrgInternalServerError with default headers values
func NewUpdateDrgInternalServerError() *UpdateDrgInternalServerError {
	return &UpdateDrgInternalServerError{}
}

/*UpdateDrgInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateDrgInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDrgInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /drgs/{drgId}][%d] updateDrgInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDrgDefault creates a UpdateDrgDefault with default headers values
func NewUpdateDrgDefault(code int) *UpdateDrgDefault {
	return &UpdateDrgDefault{
		_statusCode: code,
	}
}

/*UpdateDrgDefault handles this case with default header values.

An error has occurred.
*/
type UpdateDrgDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update drg default response
func (o *UpdateDrgDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDrgDefault) Error() string {
	return fmt.Sprintf("[PUT /drgs/{drgId}][%d] UpdateDrg default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDrgDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
