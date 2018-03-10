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

// GetIPSecConnectionDeviceConfigReader is a Reader for the GetIPSecConnectionDeviceConfig structure.
type GetIPSecConnectionDeviceConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPSecConnectionDeviceConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPSecConnectionDeviceConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetIPSecConnectionDeviceConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetIPSecConnectionDeviceConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetIPSecConnectionDeviceConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetIPSecConnectionDeviceConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPSecConnectionDeviceConfigOK creates a GetIPSecConnectionDeviceConfigOK with default headers values
func NewGetIPSecConnectionDeviceConfigOK() *GetIPSecConnectionDeviceConfigOK {
	return &GetIPSecConnectionDeviceConfigOK{}
}

/*GetIPSecConnectionDeviceConfigOK handles this case with default header values.

The configuration information was retrieved.
*/
type GetIPSecConnectionDeviceConfigOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.IPSecConnectionDeviceConfig
}

func (o *GetIPSecConnectionDeviceConfigOK) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceConfig][%d] getIpSecConnectionDeviceConfigOK  %+v", 200, o.Payload)
}

func (o *GetIPSecConnectionDeviceConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.IPSecConnectionDeviceConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDeviceConfigUnauthorized creates a GetIPSecConnectionDeviceConfigUnauthorized with default headers values
func NewGetIPSecConnectionDeviceConfigUnauthorized() *GetIPSecConnectionDeviceConfigUnauthorized {
	return &GetIPSecConnectionDeviceConfigUnauthorized{}
}

/*GetIPSecConnectionDeviceConfigUnauthorized handles this case with default header values.

Unauthorized
*/
type GetIPSecConnectionDeviceConfigUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionDeviceConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceConfig][%d] getIpSecConnectionDeviceConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIPSecConnectionDeviceConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDeviceConfigNotFound creates a GetIPSecConnectionDeviceConfigNotFound with default headers values
func NewGetIPSecConnectionDeviceConfigNotFound() *GetIPSecConnectionDeviceConfigNotFound {
	return &GetIPSecConnectionDeviceConfigNotFound{}
}

/*GetIPSecConnectionDeviceConfigNotFound handles this case with default header values.

Not Found
*/
type GetIPSecConnectionDeviceConfigNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionDeviceConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceConfig][%d] getIpSecConnectionDeviceConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetIPSecConnectionDeviceConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDeviceConfigInternalServerError creates a GetIPSecConnectionDeviceConfigInternalServerError with default headers values
func NewGetIPSecConnectionDeviceConfigInternalServerError() *GetIPSecConnectionDeviceConfigInternalServerError {
	return &GetIPSecConnectionDeviceConfigInternalServerError{}
}

/*GetIPSecConnectionDeviceConfigInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetIPSecConnectionDeviceConfigInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionDeviceConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceConfig][%d] getIpSecConnectionDeviceConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIPSecConnectionDeviceConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDeviceConfigDefault creates a GetIPSecConnectionDeviceConfigDefault with default headers values
func NewGetIPSecConnectionDeviceConfigDefault(code int) *GetIPSecConnectionDeviceConfigDefault {
	return &GetIPSecConnectionDeviceConfigDefault{
		_statusCode: code,
	}
}

/*GetIPSecConnectionDeviceConfigDefault handles this case with default header values.

An error has occurred.
*/
type GetIPSecConnectionDeviceConfigDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get IP sec connection device config default response
func (o *GetIPSecConnectionDeviceConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetIPSecConnectionDeviceConfigDefault) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceConfig][%d] GetIPSecConnectionDeviceConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPSecConnectionDeviceConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
