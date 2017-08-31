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

// GetCrossConnectLetterOfAuthorityReader is a Reader for the GetCrossConnectLetterOfAuthority structure.
type GetCrossConnectLetterOfAuthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCrossConnectLetterOfAuthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCrossConnectLetterOfAuthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetCrossConnectLetterOfAuthorityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCrossConnectLetterOfAuthorityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCrossConnectLetterOfAuthorityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCrossConnectLetterOfAuthorityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCrossConnectLetterOfAuthorityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCrossConnectLetterOfAuthorityOK creates a GetCrossConnectLetterOfAuthorityOK with default headers values
func NewGetCrossConnectLetterOfAuthorityOK() *GetCrossConnectLetterOfAuthorityOK {
	return &GetCrossConnectLetterOfAuthorityOK{}
}

/*GetCrossConnectLetterOfAuthorityOK handles this case with default header values.

The Letter of Authority was retrieved.
*/
type GetCrossConnectLetterOfAuthorityOK struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.LetterOfAuthority
}

func (o *GetCrossConnectLetterOfAuthorityOK) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/letterOfAuthority][%d] getCrossConnectLetterOfAuthorityOK  %+v", 200, o.Payload)
}

func (o *GetCrossConnectLetterOfAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.LetterOfAuthority)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectLetterOfAuthorityBadRequest creates a GetCrossConnectLetterOfAuthorityBadRequest with default headers values
func NewGetCrossConnectLetterOfAuthorityBadRequest() *GetCrossConnectLetterOfAuthorityBadRequest {
	return &GetCrossConnectLetterOfAuthorityBadRequest{}
}

/*GetCrossConnectLetterOfAuthorityBadRequest handles this case with default header values.

Bad Request
*/
type GetCrossConnectLetterOfAuthorityBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCrossConnectLetterOfAuthorityBadRequest) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/letterOfAuthority][%d] getCrossConnectLetterOfAuthorityBadRequest  %+v", 400, o.Payload)
}

func (o *GetCrossConnectLetterOfAuthorityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectLetterOfAuthorityUnauthorized creates a GetCrossConnectLetterOfAuthorityUnauthorized with default headers values
func NewGetCrossConnectLetterOfAuthorityUnauthorized() *GetCrossConnectLetterOfAuthorityUnauthorized {
	return &GetCrossConnectLetterOfAuthorityUnauthorized{}
}

/*GetCrossConnectLetterOfAuthorityUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCrossConnectLetterOfAuthorityUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCrossConnectLetterOfAuthorityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/letterOfAuthority][%d] getCrossConnectLetterOfAuthorityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCrossConnectLetterOfAuthorityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectLetterOfAuthorityNotFound creates a GetCrossConnectLetterOfAuthorityNotFound with default headers values
func NewGetCrossConnectLetterOfAuthorityNotFound() *GetCrossConnectLetterOfAuthorityNotFound {
	return &GetCrossConnectLetterOfAuthorityNotFound{}
}

/*GetCrossConnectLetterOfAuthorityNotFound handles this case with default header values.

Not Found
*/
type GetCrossConnectLetterOfAuthorityNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCrossConnectLetterOfAuthorityNotFound) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/letterOfAuthority][%d] getCrossConnectLetterOfAuthorityNotFound  %+v", 404, o.Payload)
}

func (o *GetCrossConnectLetterOfAuthorityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectLetterOfAuthorityInternalServerError creates a GetCrossConnectLetterOfAuthorityInternalServerError with default headers values
func NewGetCrossConnectLetterOfAuthorityInternalServerError() *GetCrossConnectLetterOfAuthorityInternalServerError {
	return &GetCrossConnectLetterOfAuthorityInternalServerError{}
}

/*GetCrossConnectLetterOfAuthorityInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetCrossConnectLetterOfAuthorityInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCrossConnectLetterOfAuthorityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/letterOfAuthority][%d] getCrossConnectLetterOfAuthorityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCrossConnectLetterOfAuthorityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectLetterOfAuthorityDefault creates a GetCrossConnectLetterOfAuthorityDefault with default headers values
func NewGetCrossConnectLetterOfAuthorityDefault(code int) *GetCrossConnectLetterOfAuthorityDefault {
	return &GetCrossConnectLetterOfAuthorityDefault{
		_statusCode: code,
	}
}

/*GetCrossConnectLetterOfAuthorityDefault handles this case with default header values.

An error has occurred.
*/
type GetCrossConnectLetterOfAuthorityDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get cross connect letter of authority default response
func (o *GetCrossConnectLetterOfAuthorityDefault) Code() int {
	return o._statusCode
}

func (o *GetCrossConnectLetterOfAuthorityDefault) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/letterOfAuthority][%d] GetCrossConnectLetterOfAuthority default  %+v", o._statusCode, o.Payload)
}

func (o *GetCrossConnectLetterOfAuthorityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
