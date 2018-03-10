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

// UpdateCrossConnectReader is a Reader for the UpdateCrossConnect structure.
type UpdateCrossConnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCrossConnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCrossConnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateCrossConnectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateCrossConnectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateCrossConnectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateCrossConnectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateCrossConnectPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateCrossConnectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateCrossConnectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCrossConnectOK creates a UpdateCrossConnectOK with default headers values
func NewUpdateCrossConnectOK() *UpdateCrossConnectOK {
	return &UpdateCrossConnectOK{}
}

/*UpdateCrossConnectOK handles this case with default header values.

The cross-connect was updated.
*/
type UpdateCrossConnectOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.CrossConnect
}

func (o *UpdateCrossConnectOK) Error() string {
	return fmt.Sprintf("[PUT /crossConnects/{crossConnectId}][%d] updateCrossConnectOK  %+v", 200, o.Payload)
}

func (o *UpdateCrossConnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.CrossConnect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCrossConnectBadRequest creates a UpdateCrossConnectBadRequest with default headers values
func NewUpdateCrossConnectBadRequest() *UpdateCrossConnectBadRequest {
	return &UpdateCrossConnectBadRequest{}
}

/*UpdateCrossConnectBadRequest handles this case with default header values.

Bad Request
*/
type UpdateCrossConnectBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCrossConnectBadRequest) Error() string {
	return fmt.Sprintf("[PUT /crossConnects/{crossConnectId}][%d] updateCrossConnectBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCrossConnectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCrossConnectUnauthorized creates a UpdateCrossConnectUnauthorized with default headers values
func NewUpdateCrossConnectUnauthorized() *UpdateCrossConnectUnauthorized {
	return &UpdateCrossConnectUnauthorized{}
}

/*UpdateCrossConnectUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCrossConnectUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCrossConnectUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /crossConnects/{crossConnectId}][%d] updateCrossConnectUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateCrossConnectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCrossConnectNotFound creates a UpdateCrossConnectNotFound with default headers values
func NewUpdateCrossConnectNotFound() *UpdateCrossConnectNotFound {
	return &UpdateCrossConnectNotFound{}
}

/*UpdateCrossConnectNotFound handles this case with default header values.

Not Found
*/
type UpdateCrossConnectNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCrossConnectNotFound) Error() string {
	return fmt.Sprintf("[PUT /crossConnects/{crossConnectId}][%d] updateCrossConnectNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCrossConnectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCrossConnectConflict creates a UpdateCrossConnectConflict with default headers values
func NewUpdateCrossConnectConflict() *UpdateCrossConnectConflict {
	return &UpdateCrossConnectConflict{}
}

/*UpdateCrossConnectConflict handles this case with default header values.

Conflict
*/
type UpdateCrossConnectConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCrossConnectConflict) Error() string {
	return fmt.Sprintf("[PUT /crossConnects/{crossConnectId}][%d] updateCrossConnectConflict  %+v", 409, o.Payload)
}

func (o *UpdateCrossConnectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCrossConnectPreconditionFailed creates a UpdateCrossConnectPreconditionFailed with default headers values
func NewUpdateCrossConnectPreconditionFailed() *UpdateCrossConnectPreconditionFailed {
	return &UpdateCrossConnectPreconditionFailed{}
}

/*UpdateCrossConnectPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateCrossConnectPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCrossConnectPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /crossConnects/{crossConnectId}][%d] updateCrossConnectPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateCrossConnectPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCrossConnectInternalServerError creates a UpdateCrossConnectInternalServerError with default headers values
func NewUpdateCrossConnectInternalServerError() *UpdateCrossConnectInternalServerError {
	return &UpdateCrossConnectInternalServerError{}
}

/*UpdateCrossConnectInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateCrossConnectInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCrossConnectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /crossConnects/{crossConnectId}][%d] updateCrossConnectInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCrossConnectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCrossConnectDefault creates a UpdateCrossConnectDefault with default headers values
func NewUpdateCrossConnectDefault(code int) *UpdateCrossConnectDefault {
	return &UpdateCrossConnectDefault{
		_statusCode: code,
	}
}

/*UpdateCrossConnectDefault handles this case with default header values.

An error has occurred.
*/
type UpdateCrossConnectDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update cross connect default response
func (o *UpdateCrossConnectDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCrossConnectDefault) Error() string {
	return fmt.Sprintf("[PUT /crossConnects/{crossConnectId}][%d] UpdateCrossConnect default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCrossConnectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
