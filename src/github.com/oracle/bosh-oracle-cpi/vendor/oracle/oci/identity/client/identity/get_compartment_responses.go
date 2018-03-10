package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/identity/models"
)

// GetCompartmentReader is a Reader for the GetCompartment structure.
type GetCompartmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompartmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCompartmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetCompartmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCompartmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCompartmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCompartmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCompartmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCompartmentOK creates a GetCompartmentOK with default headers values
func NewGetCompartmentOK() *GetCompartmentOK {
	return &GetCompartmentOK{}
}

/*GetCompartmentOK handles this case with default header values.

The compartment was retrieved.
*/
type GetCompartmentOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Compartment
}

func (o *GetCompartmentOK) Error() string {
	return fmt.Sprintf("[GET /compartments/{compartmentId}][%d] getCompartmentOK  %+v", 200, o.Payload)
}

func (o *GetCompartmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Compartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompartmentBadRequest creates a GetCompartmentBadRequest with default headers values
func NewGetCompartmentBadRequest() *GetCompartmentBadRequest {
	return &GetCompartmentBadRequest{}
}

/*GetCompartmentBadRequest handles this case with default header values.

Bad Request
*/
type GetCompartmentBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCompartmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /compartments/{compartmentId}][%d] getCompartmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetCompartmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompartmentForbidden creates a GetCompartmentForbidden with default headers values
func NewGetCompartmentForbidden() *GetCompartmentForbidden {
	return &GetCompartmentForbidden{}
}

/*GetCompartmentForbidden handles this case with default header values.

Forbidden
*/
type GetCompartmentForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCompartmentForbidden) Error() string {
	return fmt.Sprintf("[GET /compartments/{compartmentId}][%d] getCompartmentForbidden  %+v", 403, o.Payload)
}

func (o *GetCompartmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompartmentNotFound creates a GetCompartmentNotFound with default headers values
func NewGetCompartmentNotFound() *GetCompartmentNotFound {
	return &GetCompartmentNotFound{}
}

/*GetCompartmentNotFound handles this case with default header values.

Not Found
*/
type GetCompartmentNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCompartmentNotFound) Error() string {
	return fmt.Sprintf("[GET /compartments/{compartmentId}][%d] getCompartmentNotFound  %+v", 404, o.Payload)
}

func (o *GetCompartmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompartmentInternalServerError creates a GetCompartmentInternalServerError with default headers values
func NewGetCompartmentInternalServerError() *GetCompartmentInternalServerError {
	return &GetCompartmentInternalServerError{}
}

/*GetCompartmentInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetCompartmentInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCompartmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compartments/{compartmentId}][%d] getCompartmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCompartmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompartmentDefault creates a GetCompartmentDefault with default headers values
func NewGetCompartmentDefault(code int) *GetCompartmentDefault {
	return &GetCompartmentDefault{
		_statusCode: code,
	}
}

/*GetCompartmentDefault handles this case with default header values.

An error has occurred.

*/
type GetCompartmentDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get compartment default response
func (o *GetCompartmentDefault) Code() int {
	return o._statusCode
}

func (o *GetCompartmentDefault) Error() string {
	return fmt.Sprintf("[GET /compartments/{compartmentId}][%d] GetCompartment default  %+v", o._statusCode, o.Payload)
}

func (o *GetCompartmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
