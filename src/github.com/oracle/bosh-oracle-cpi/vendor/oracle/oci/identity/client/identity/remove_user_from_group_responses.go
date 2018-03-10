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

// RemoveUserFromGroupReader is a Reader for the RemoveUserFromGroup structure.
type RemoveUserFromGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveUserFromGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRemoveUserFromGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveUserFromGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveUserFromGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveUserFromGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveUserFromGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewRemoveUserFromGroupPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRemoveUserFromGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRemoveUserFromGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveUserFromGroupNoContent creates a RemoveUserFromGroupNoContent with default headers values
func NewRemoveUserFromGroupNoContent() *RemoveUserFromGroupNoContent {
	return &RemoveUserFromGroupNoContent{}
}

/*RemoveUserFromGroupNoContent handles this case with default header values.

The user is being removed from the group.
*/
type RemoveUserFromGroupNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *RemoveUserFromGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /userGroupMemberships/{userGroupMembershipId}][%d] removeUserFromGroupNoContent ", 204)
}

func (o *RemoveUserFromGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewRemoveUserFromGroupBadRequest creates a RemoveUserFromGroupBadRequest with default headers values
func NewRemoveUserFromGroupBadRequest() *RemoveUserFromGroupBadRequest {
	return &RemoveUserFromGroupBadRequest{}
}

/*RemoveUserFromGroupBadRequest handles this case with default header values.

Bad Request
*/
type RemoveUserFromGroupBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveUserFromGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /userGroupMemberships/{userGroupMembershipId}][%d] removeUserFromGroupBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveUserFromGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserFromGroupForbidden creates a RemoveUserFromGroupForbidden with default headers values
func NewRemoveUserFromGroupForbidden() *RemoveUserFromGroupForbidden {
	return &RemoveUserFromGroupForbidden{}
}

/*RemoveUserFromGroupForbidden handles this case with default header values.

Forbidden
*/
type RemoveUserFromGroupForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveUserFromGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /userGroupMemberships/{userGroupMembershipId}][%d] removeUserFromGroupForbidden  %+v", 403, o.Payload)
}

func (o *RemoveUserFromGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserFromGroupNotFound creates a RemoveUserFromGroupNotFound with default headers values
func NewRemoveUserFromGroupNotFound() *RemoveUserFromGroupNotFound {
	return &RemoveUserFromGroupNotFound{}
}

/*RemoveUserFromGroupNotFound handles this case with default header values.

Not Found
*/
type RemoveUserFromGroupNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveUserFromGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /userGroupMemberships/{userGroupMembershipId}][%d] removeUserFromGroupNotFound  %+v", 404, o.Payload)
}

func (o *RemoveUserFromGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserFromGroupConflict creates a RemoveUserFromGroupConflict with default headers values
func NewRemoveUserFromGroupConflict() *RemoveUserFromGroupConflict {
	return &RemoveUserFromGroupConflict{}
}

/*RemoveUserFromGroupConflict handles this case with default header values.

Conflict
*/
type RemoveUserFromGroupConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveUserFromGroupConflict) Error() string {
	return fmt.Sprintf("[DELETE /userGroupMemberships/{userGroupMembershipId}][%d] removeUserFromGroupConflict  %+v", 409, o.Payload)
}

func (o *RemoveUserFromGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserFromGroupPreconditionFailed creates a RemoveUserFromGroupPreconditionFailed with default headers values
func NewRemoveUserFromGroupPreconditionFailed() *RemoveUserFromGroupPreconditionFailed {
	return &RemoveUserFromGroupPreconditionFailed{}
}

/*RemoveUserFromGroupPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type RemoveUserFromGroupPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveUserFromGroupPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /userGroupMemberships/{userGroupMembershipId}][%d] removeUserFromGroupPreconditionFailed  %+v", 412, o.Payload)
}

func (o *RemoveUserFromGroupPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserFromGroupInternalServerError creates a RemoveUserFromGroupInternalServerError with default headers values
func NewRemoveUserFromGroupInternalServerError() *RemoveUserFromGroupInternalServerError {
	return &RemoveUserFromGroupInternalServerError{}
}

/*RemoveUserFromGroupInternalServerError handles this case with default header values.

Internal Server Error
*/
type RemoveUserFromGroupInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveUserFromGroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /userGroupMemberships/{userGroupMembershipId}][%d] removeUserFromGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveUserFromGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserFromGroupDefault creates a RemoveUserFromGroupDefault with default headers values
func NewRemoveUserFromGroupDefault(code int) *RemoveUserFromGroupDefault {
	return &RemoveUserFromGroupDefault{
		_statusCode: code,
	}
}

/*RemoveUserFromGroupDefault handles this case with default header values.

An error has occurred.

*/
type RemoveUserFromGroupDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the remove user from group default response
func (o *RemoveUserFromGroupDefault) Code() int {
	return o._statusCode
}

func (o *RemoveUserFromGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /userGroupMemberships/{userGroupMembershipId}][%d] RemoveUserFromGroup default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveUserFromGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
