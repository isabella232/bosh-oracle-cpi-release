package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// ListImagesReader is a Reader for the ListImages structure.
type ListImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListImagesOK creates a ListImagesOK with default headers values
func NewListImagesOK() *ListImagesOK {
	return &ListImagesOK{}
}

/*ListImagesOK handles this case with default header values.

The list is being retrieved.
*/
type ListImagesOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.Image
}

func (o *ListImagesOK) Error() string {
	return fmt.Sprintf("[GET /images/][%d] listImagesOK  %+v", 200, o.Payload)
}

func (o *ListImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListImagesBadRequest creates a ListImagesBadRequest with default headers values
func NewListImagesBadRequest() *ListImagesBadRequest {
	return &ListImagesBadRequest{}
}

/*ListImagesBadRequest handles this case with default header values.

Bad Request
*/
type ListImagesBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /images/][%d] listImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ListImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesUnauthorized creates a ListImagesUnauthorized with default headers values
func NewListImagesUnauthorized() *ListImagesUnauthorized {
	return &ListImagesUnauthorized{}
}

/*ListImagesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListImagesUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /images/][%d] listImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesNotFound creates a ListImagesNotFound with default headers values
func NewListImagesNotFound() *ListImagesNotFound {
	return &ListImagesNotFound{}
}

/*ListImagesNotFound handles this case with default header values.

Not Found
*/
type ListImagesNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /images/][%d] listImagesNotFound  %+v", 404, o.Payload)
}

func (o *ListImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesInternalServerError creates a ListImagesInternalServerError with default headers values
func NewListImagesInternalServerError() *ListImagesInternalServerError {
	return &ListImagesInternalServerError{}
}

/*ListImagesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListImagesInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/][%d] listImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesDefault creates a ListImagesDefault with default headers values
func NewListImagesDefault(code int) *ListImagesDefault {
	return &ListImagesDefault{
		_statusCode: code,
	}
}

/*ListImagesDefault handles this case with default header values.

An error has occurred.
*/
type ListImagesDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list images default response
func (o *ListImagesDefault) Code() int {
	return o._statusCode
}

func (o *ListImagesDefault) Error() string {
	return fmt.Sprintf("[GET /images/][%d] ListImages default  %+v", o._statusCode, o.Payload)
}

func (o *ListImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
