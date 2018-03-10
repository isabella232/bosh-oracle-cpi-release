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

// CreateSecurityListReader is a Reader for the CreateSecurityList structure.
type CreateSecurityListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSecurityListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSecurityListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSecurityListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateSecurityListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateSecurityListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSecurityListConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateSecurityListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSecurityListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSecurityListOK creates a CreateSecurityListOK with default headers values
func NewCreateSecurityListOK() *CreateSecurityListOK {
	return &CreateSecurityListOK{}
}

/*CreateSecurityListOK handles this case with default header values.

The security list has been created.
*/
type CreateSecurityListOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.SecurityList
}

func (o *CreateSecurityListOK) Error() string {
	return fmt.Sprintf("[POST /securityLists][%d] createSecurityListOK  %+v", 200, o.Payload)
}

func (o *CreateSecurityListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.SecurityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityListBadRequest creates a CreateSecurityListBadRequest with default headers values
func NewCreateSecurityListBadRequest() *CreateSecurityListBadRequest {
	return &CreateSecurityListBadRequest{}
}

/*CreateSecurityListBadRequest handles this case with default header values.

Bad Request
*/
type CreateSecurityListBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSecurityListBadRequest) Error() string {
	return fmt.Sprintf("[POST /securityLists][%d] createSecurityListBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSecurityListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityListUnauthorized creates a CreateSecurityListUnauthorized with default headers values
func NewCreateSecurityListUnauthorized() *CreateSecurityListUnauthorized {
	return &CreateSecurityListUnauthorized{}
}

/*CreateSecurityListUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateSecurityListUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSecurityListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /securityLists][%d] createSecurityListUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateSecurityListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityListNotFound creates a CreateSecurityListNotFound with default headers values
func NewCreateSecurityListNotFound() *CreateSecurityListNotFound {
	return &CreateSecurityListNotFound{}
}

/*CreateSecurityListNotFound handles this case with default header values.

Not Found
*/
type CreateSecurityListNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSecurityListNotFound) Error() string {
	return fmt.Sprintf("[POST /securityLists][%d] createSecurityListNotFound  %+v", 404, o.Payload)
}

func (o *CreateSecurityListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityListConflict creates a CreateSecurityListConflict with default headers values
func NewCreateSecurityListConflict() *CreateSecurityListConflict {
	return &CreateSecurityListConflict{}
}

/*CreateSecurityListConflict handles this case with default header values.

Conflict
*/
type CreateSecurityListConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSecurityListConflict) Error() string {
	return fmt.Sprintf("[POST /securityLists][%d] createSecurityListConflict  %+v", 409, o.Payload)
}

func (o *CreateSecurityListConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityListInternalServerError creates a CreateSecurityListInternalServerError with default headers values
func NewCreateSecurityListInternalServerError() *CreateSecurityListInternalServerError {
	return &CreateSecurityListInternalServerError{}
}

/*CreateSecurityListInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateSecurityListInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSecurityListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /securityLists][%d] createSecurityListInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSecurityListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityListDefault creates a CreateSecurityListDefault with default headers values
func NewCreateSecurityListDefault(code int) *CreateSecurityListDefault {
	return &CreateSecurityListDefault{
		_statusCode: code,
	}
}

/*CreateSecurityListDefault handles this case with default header values.

An error has occurred.
*/
type CreateSecurityListDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create security list default response
func (o *CreateSecurityListDefault) Code() int {
	return o._statusCode
}

func (o *CreateSecurityListDefault) Error() string {
	return fmt.Sprintf("[POST /securityLists][%d] CreateSecurityList default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSecurityListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
