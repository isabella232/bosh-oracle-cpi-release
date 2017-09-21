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

// GetPrivateIPReader is a Reader for the GetPrivateIP structure.
type GetPrivateIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPrivateIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetPrivateIPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPrivateIPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetPrivateIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPrivateIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPrivateIPOK creates a GetPrivateIPOK with default headers values
func NewGetPrivateIPOK() *GetPrivateIPOK {
	return &GetPrivateIPOK{}
}

/*GetPrivateIPOK handles this case with default header values.

The private IP.
*/
type GetPrivateIPOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.PrivateIP
}

func (o *GetPrivateIPOK) Error() string {
	return fmt.Sprintf("[GET /privateIps/{privateIpId}][%d] getPrivateIpOK  %+v", 200, o.Payload)
}

func (o *GetPrivateIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.PrivateIP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateIPBadRequest creates a GetPrivateIPBadRequest with default headers values
func NewGetPrivateIPBadRequest() *GetPrivateIPBadRequest {
	return &GetPrivateIPBadRequest{}
}

/*GetPrivateIPBadRequest handles this case with default header values.

Bad Request
*/
type GetPrivateIPBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPrivateIPBadRequest) Error() string {
	return fmt.Sprintf("[GET /privateIps/{privateIpId}][%d] getPrivateIpBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrivateIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateIPUnauthorized creates a GetPrivateIPUnauthorized with default headers values
func NewGetPrivateIPUnauthorized() *GetPrivateIPUnauthorized {
	return &GetPrivateIPUnauthorized{}
}

/*GetPrivateIPUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPrivateIPUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPrivateIPUnauthorized) Error() string {
	return fmt.Sprintf("[GET /privateIps/{privateIpId}][%d] getPrivateIpUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPrivateIPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateIPNotFound creates a GetPrivateIPNotFound with default headers values
func NewGetPrivateIPNotFound() *GetPrivateIPNotFound {
	return &GetPrivateIPNotFound{}
}

/*GetPrivateIPNotFound handles this case with default header values.

Not Found
*/
type GetPrivateIPNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPrivateIPNotFound) Error() string {
	return fmt.Sprintf("[GET /privateIps/{privateIpId}][%d] getPrivateIpNotFound  %+v", 404, o.Payload)
}

func (o *GetPrivateIPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateIPInternalServerError creates a GetPrivateIPInternalServerError with default headers values
func NewGetPrivateIPInternalServerError() *GetPrivateIPInternalServerError {
	return &GetPrivateIPInternalServerError{}
}

/*GetPrivateIPInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPrivateIPInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPrivateIPInternalServerError) Error() string {
	return fmt.Sprintf("[GET /privateIps/{privateIpId}][%d] getPrivateIpInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPrivateIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateIPDefault creates a GetPrivateIPDefault with default headers values
func NewGetPrivateIPDefault(code int) *GetPrivateIPDefault {
	return &GetPrivateIPDefault{
		_statusCode: code,
	}
}

/*GetPrivateIPDefault handles this case with default header values.

An error has occurred.
*/
type GetPrivateIPDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get private Ip default response
func (o *GetPrivateIPDefault) Code() int {
	return o._statusCode
}

func (o *GetPrivateIPDefault) Error() string {
	return fmt.Sprintf("[GET /privateIps/{privateIpId}][%d] GetPrivateIp default  %+v", o._statusCode, o.Payload)
}

func (o *GetPrivateIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
