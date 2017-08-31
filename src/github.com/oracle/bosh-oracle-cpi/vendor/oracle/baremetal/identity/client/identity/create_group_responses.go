package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/identity/models"
)

// CreateGroupReader is a Reader for the CreateGroup structure.
type CreateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateGroupOK creates a CreateGroupOK with default headers values
func NewCreateGroupOK() *CreateGroupOK {
	return &CreateGroupOK{}
}

/*CreateGroupOK handles this case with default header values.

The group is being created.
*/
type CreateGroupOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Group
}

func (o *CreateGroupOK) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] createGroupOK  %+v", 200, o.Payload)
}

func (o *CreateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupBadRequest creates a CreateGroupBadRequest with default headers values
func NewCreateGroupBadRequest() *CreateGroupBadRequest {
	return &CreateGroupBadRequest{}
}

/*CreateGroupBadRequest handles this case with default header values.

Bad Request
*/
type CreateGroupBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] createGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupForbidden creates a CreateGroupForbidden with default headers values
func NewCreateGroupForbidden() *CreateGroupForbidden {
	return &CreateGroupForbidden{}
}

/*CreateGroupForbidden handles this case with default header values.

Forbidden
*/
type CreateGroupForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] createGroupForbidden  %+v", 403, o.Payload)
}

func (o *CreateGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupNotFound creates a CreateGroupNotFound with default headers values
func NewCreateGroupNotFound() *CreateGroupNotFound {
	return &CreateGroupNotFound{}
}

/*CreateGroupNotFound handles this case with default header values.

Not Found
*/
type CreateGroupNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] createGroupNotFound  %+v", 404, o.Payload)
}

func (o *CreateGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupConflict creates a CreateGroupConflict with default headers values
func NewCreateGroupConflict() *CreateGroupConflict {
	return &CreateGroupConflict{}
}

/*CreateGroupConflict handles this case with default header values.

Conflict
*/
type CreateGroupConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateGroupConflict) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] createGroupConflict  %+v", 409, o.Payload)
}

func (o *CreateGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupInternalServerError creates a CreateGroupInternalServerError with default headers values
func NewCreateGroupInternalServerError() *CreateGroupInternalServerError {
	return &CreateGroupInternalServerError{}
}

/*CreateGroupInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateGroupInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] createGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupDefault creates a CreateGroupDefault with default headers values
func NewCreateGroupDefault(code int) *CreateGroupDefault {
	return &CreateGroupDefault{
		_statusCode: code,
	}
}

/*CreateGroupDefault handles this case with default header values.

An error has occurred.

*/
type CreateGroupDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create group default response
func (o *CreateGroupDefault) Code() int {
	return o._statusCode
}

func (o *CreateGroupDefault) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] CreateGroup default  %+v", o._statusCode, o.Payload)
}

func (o *CreateGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
