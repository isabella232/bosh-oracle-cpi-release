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

// CreateCpeReader is a Reader for the CreateCpe structure.
type CreateCpeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCpeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCpeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateCpeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateCpeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateCpeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateCpeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateCpeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateCpeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCpeOK creates a CreateCpeOK with default headers values
func NewCreateCpeOK() *CreateCpeOK {
	return &CreateCpeOK{}
}

/*CreateCpeOK handles this case with default header values.

The CPE has been created.
*/
type CreateCpeOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Cpe
}

func (o *CreateCpeOK) Error() string {
	return fmt.Sprintf("[POST /cpes][%d] createCpeOK  %+v", 200, o.Payload)
}

func (o *CreateCpeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Cpe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCpeBadRequest creates a CreateCpeBadRequest with default headers values
func NewCreateCpeBadRequest() *CreateCpeBadRequest {
	return &CreateCpeBadRequest{}
}

/*CreateCpeBadRequest handles this case with default header values.

Bad Request
*/
type CreateCpeBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCpeBadRequest) Error() string {
	return fmt.Sprintf("[POST /cpes][%d] createCpeBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCpeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCpeUnauthorized creates a CreateCpeUnauthorized with default headers values
func NewCreateCpeUnauthorized() *CreateCpeUnauthorized {
	return &CreateCpeUnauthorized{}
}

/*CreateCpeUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCpeUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCpeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cpes][%d] createCpeUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateCpeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCpeNotFound creates a CreateCpeNotFound with default headers values
func NewCreateCpeNotFound() *CreateCpeNotFound {
	return &CreateCpeNotFound{}
}

/*CreateCpeNotFound handles this case with default header values.

Not Found
*/
type CreateCpeNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCpeNotFound) Error() string {
	return fmt.Sprintf("[POST /cpes][%d] createCpeNotFound  %+v", 404, o.Payload)
}

func (o *CreateCpeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCpeConflict creates a CreateCpeConflict with default headers values
func NewCreateCpeConflict() *CreateCpeConflict {
	return &CreateCpeConflict{}
}

/*CreateCpeConflict handles this case with default header values.

Conflict
*/
type CreateCpeConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCpeConflict) Error() string {
	return fmt.Sprintf("[POST /cpes][%d] createCpeConflict  %+v", 409, o.Payload)
}

func (o *CreateCpeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCpeInternalServerError creates a CreateCpeInternalServerError with default headers values
func NewCreateCpeInternalServerError() *CreateCpeInternalServerError {
	return &CreateCpeInternalServerError{}
}

/*CreateCpeInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateCpeInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCpeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cpes][%d] createCpeInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCpeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCpeDefault creates a CreateCpeDefault with default headers values
func NewCreateCpeDefault(code int) *CreateCpeDefault {
	return &CreateCpeDefault{
		_statusCode: code,
	}
}

/*CreateCpeDefault handles this case with default header values.

An error has occurred.
*/
type CreateCpeDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create cpe default response
func (o *CreateCpeDefault) Code() int {
	return o._statusCode
}

func (o *CreateCpeDefault) Error() string {
	return fmt.Sprintf("[POST /cpes][%d] CreateCpe default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCpeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
