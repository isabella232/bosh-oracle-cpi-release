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

// CreateCrossConnectGroupReader is a Reader for the CreateCrossConnectGroup structure.
type CreateCrossConnectGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCrossConnectGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCrossConnectGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateCrossConnectGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateCrossConnectGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateCrossConnectGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateCrossConnectGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateCrossConnectGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCrossConnectGroupOK creates a CreateCrossConnectGroupOK with default headers values
func NewCreateCrossConnectGroupOK() *CreateCrossConnectGroupOK {
	return &CreateCrossConnectGroupOK{}
}

/*CreateCrossConnectGroupOK handles this case with default header values.

The cross-connect group was created.
*/
type CreateCrossConnectGroupOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.CrossConnectGroup
}

func (o *CreateCrossConnectGroupOK) Error() string {
	return fmt.Sprintf("[POST /crossConnectGroups][%d] createCrossConnectGroupOK  %+v", 200, o.Payload)
}

func (o *CreateCrossConnectGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.CrossConnectGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectGroupBadRequest creates a CreateCrossConnectGroupBadRequest with default headers values
func NewCreateCrossConnectGroupBadRequest() *CreateCrossConnectGroupBadRequest {
	return &CreateCrossConnectGroupBadRequest{}
}

/*CreateCrossConnectGroupBadRequest handles this case with default header values.

Bad Request
*/
type CreateCrossConnectGroupBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCrossConnectGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /crossConnectGroups][%d] createCrossConnectGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCrossConnectGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectGroupUnauthorized creates a CreateCrossConnectGroupUnauthorized with default headers values
func NewCreateCrossConnectGroupUnauthorized() *CreateCrossConnectGroupUnauthorized {
	return &CreateCrossConnectGroupUnauthorized{}
}

/*CreateCrossConnectGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCrossConnectGroupUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCrossConnectGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /crossConnectGroups][%d] createCrossConnectGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateCrossConnectGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectGroupConflict creates a CreateCrossConnectGroupConflict with default headers values
func NewCreateCrossConnectGroupConflict() *CreateCrossConnectGroupConflict {
	return &CreateCrossConnectGroupConflict{}
}

/*CreateCrossConnectGroupConflict handles this case with default header values.

Conflict
*/
type CreateCrossConnectGroupConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCrossConnectGroupConflict) Error() string {
	return fmt.Sprintf("[POST /crossConnectGroups][%d] createCrossConnectGroupConflict  %+v", 409, o.Payload)
}

func (o *CreateCrossConnectGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectGroupInternalServerError creates a CreateCrossConnectGroupInternalServerError with default headers values
func NewCreateCrossConnectGroupInternalServerError() *CreateCrossConnectGroupInternalServerError {
	return &CreateCrossConnectGroupInternalServerError{}
}

/*CreateCrossConnectGroupInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateCrossConnectGroupInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCrossConnectGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /crossConnectGroups][%d] createCrossConnectGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCrossConnectGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectGroupDefault creates a CreateCrossConnectGroupDefault with default headers values
func NewCreateCrossConnectGroupDefault(code int) *CreateCrossConnectGroupDefault {
	return &CreateCrossConnectGroupDefault{
		_statusCode: code,
	}
}

/*CreateCrossConnectGroupDefault handles this case with default header values.

An error has occurred.
*/
type CreateCrossConnectGroupDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create cross connect group default response
func (o *CreateCrossConnectGroupDefault) Code() int {
	return o._statusCode
}

func (o *CreateCrossConnectGroupDefault) Error() string {
	return fmt.Sprintf("[POST /crossConnectGroups][%d] CreateCrossConnectGroup default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCrossConnectGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
