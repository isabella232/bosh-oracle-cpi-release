// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// ListBootVolumeAttachmentsReader is a Reader for the ListBootVolumeAttachments structure.
type ListBootVolumeAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBootVolumeAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListBootVolumeAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListBootVolumeAttachmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListBootVolumeAttachmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListBootVolumeAttachmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListBootVolumeAttachmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListBootVolumeAttachmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBootVolumeAttachmentsOK creates a ListBootVolumeAttachmentsOK with default headers values
func NewListBootVolumeAttachmentsOK() *ListBootVolumeAttachmentsOK {
	return &ListBootVolumeAttachmentsOK{}
}

/*ListBootVolumeAttachmentsOK handles this case with default header values.

The list is being retrieved.
*/
type ListBootVolumeAttachmentsOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.BootVolumeAttachment
}

func (o *ListBootVolumeAttachmentsOK) Error() string {
	return fmt.Sprintf("[GET /bootVolumeAttachments/][%d] listBootVolumeAttachmentsOK  %+v", 200, o.Payload)
}

func (o *ListBootVolumeAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-next-page
	o.OpcNextPage = response.GetHeader("opc-next-page")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBootVolumeAttachmentsBadRequest creates a ListBootVolumeAttachmentsBadRequest with default headers values
func NewListBootVolumeAttachmentsBadRequest() *ListBootVolumeAttachmentsBadRequest {
	return &ListBootVolumeAttachmentsBadRequest{}
}

/*ListBootVolumeAttachmentsBadRequest handles this case with default header values.

Bad Request
*/
type ListBootVolumeAttachmentsBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListBootVolumeAttachmentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /bootVolumeAttachments/][%d] listBootVolumeAttachmentsBadRequest  %+v", 400, o.Payload)
}

func (o *ListBootVolumeAttachmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBootVolumeAttachmentsUnauthorized creates a ListBootVolumeAttachmentsUnauthorized with default headers values
func NewListBootVolumeAttachmentsUnauthorized() *ListBootVolumeAttachmentsUnauthorized {
	return &ListBootVolumeAttachmentsUnauthorized{}
}

/*ListBootVolumeAttachmentsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListBootVolumeAttachmentsUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListBootVolumeAttachmentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bootVolumeAttachments/][%d] listBootVolumeAttachmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListBootVolumeAttachmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBootVolumeAttachmentsNotFound creates a ListBootVolumeAttachmentsNotFound with default headers values
func NewListBootVolumeAttachmentsNotFound() *ListBootVolumeAttachmentsNotFound {
	return &ListBootVolumeAttachmentsNotFound{}
}

/*ListBootVolumeAttachmentsNotFound handles this case with default header values.

Not Found
*/
type ListBootVolumeAttachmentsNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListBootVolumeAttachmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /bootVolumeAttachments/][%d] listBootVolumeAttachmentsNotFound  %+v", 404, o.Payload)
}

func (o *ListBootVolumeAttachmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBootVolumeAttachmentsInternalServerError creates a ListBootVolumeAttachmentsInternalServerError with default headers values
func NewListBootVolumeAttachmentsInternalServerError() *ListBootVolumeAttachmentsInternalServerError {
	return &ListBootVolumeAttachmentsInternalServerError{}
}

/*ListBootVolumeAttachmentsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListBootVolumeAttachmentsInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListBootVolumeAttachmentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bootVolumeAttachments/][%d] listBootVolumeAttachmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListBootVolumeAttachmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBootVolumeAttachmentsDefault creates a ListBootVolumeAttachmentsDefault with default headers values
func NewListBootVolumeAttachmentsDefault(code int) *ListBootVolumeAttachmentsDefault {
	return &ListBootVolumeAttachmentsDefault{
		_statusCode: code,
	}
}

/*ListBootVolumeAttachmentsDefault handles this case with default header values.

An error has occurred.
*/
type ListBootVolumeAttachmentsDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list boot volume attachments default response
func (o *ListBootVolumeAttachmentsDefault) Code() int {
	return o._statusCode
}

func (o *ListBootVolumeAttachmentsDefault) Error() string {
	return fmt.Sprintf("[GET /bootVolumeAttachments/][%d] ListBootVolumeAttachments default  %+v", o._statusCode, o.Payload)
}

func (o *ListBootVolumeAttachmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
