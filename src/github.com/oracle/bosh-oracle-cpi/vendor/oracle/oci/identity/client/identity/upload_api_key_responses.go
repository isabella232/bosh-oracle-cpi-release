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

// UploadAPIKeyReader is a Reader for the UploadAPIKey structure.
type UploadAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUploadAPIKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUploadAPIKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUploadAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUploadAPIKeyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUploadAPIKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUploadAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadAPIKeyOK creates a UploadAPIKeyOK with default headers values
func NewUploadAPIKeyOK() *UploadAPIKeyOK {
	return &UploadAPIKeyOK{}
}

/*UploadAPIKeyOK handles this case with default header values.

The key is being uploaded.
*/
type UploadAPIKeyOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.APIKey
}

func (o *UploadAPIKeyOK) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/apiKeys/][%d] uploadApiKeyOK  %+v", 200, o.Payload)
}

func (o *UploadAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.APIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAPIKeyBadRequest creates a UploadAPIKeyBadRequest with default headers values
func NewUploadAPIKeyBadRequest() *UploadAPIKeyBadRequest {
	return &UploadAPIKeyBadRequest{}
}

/*UploadAPIKeyBadRequest handles this case with default header values.

Bad Request
*/
type UploadAPIKeyBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UploadAPIKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/apiKeys/][%d] uploadApiKeyBadRequest  %+v", 400, o.Payload)
}

func (o *UploadAPIKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAPIKeyForbidden creates a UploadAPIKeyForbidden with default headers values
func NewUploadAPIKeyForbidden() *UploadAPIKeyForbidden {
	return &UploadAPIKeyForbidden{}
}

/*UploadAPIKeyForbidden handles this case with default header values.

Forbidden
*/
type UploadAPIKeyForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UploadAPIKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/apiKeys/][%d] uploadApiKeyForbidden  %+v", 403, o.Payload)
}

func (o *UploadAPIKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAPIKeyNotFound creates a UploadAPIKeyNotFound with default headers values
func NewUploadAPIKeyNotFound() *UploadAPIKeyNotFound {
	return &UploadAPIKeyNotFound{}
}

/*UploadAPIKeyNotFound handles this case with default header values.

Not Found
*/
type UploadAPIKeyNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UploadAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/apiKeys/][%d] uploadApiKeyNotFound  %+v", 404, o.Payload)
}

func (o *UploadAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAPIKeyConflict creates a UploadAPIKeyConflict with default headers values
func NewUploadAPIKeyConflict() *UploadAPIKeyConflict {
	return &UploadAPIKeyConflict{}
}

/*UploadAPIKeyConflict handles this case with default header values.

Conflict
*/
type UploadAPIKeyConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UploadAPIKeyConflict) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/apiKeys/][%d] uploadApiKeyConflict  %+v", 409, o.Payload)
}

func (o *UploadAPIKeyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAPIKeyInternalServerError creates a UploadAPIKeyInternalServerError with default headers values
func NewUploadAPIKeyInternalServerError() *UploadAPIKeyInternalServerError {
	return &UploadAPIKeyInternalServerError{}
}

/*UploadAPIKeyInternalServerError handles this case with default header values.

Internal Server Error
*/
type UploadAPIKeyInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UploadAPIKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/apiKeys/][%d] uploadApiKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadAPIKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAPIKeyDefault creates a UploadAPIKeyDefault with default headers values
func NewUploadAPIKeyDefault(code int) *UploadAPIKeyDefault {
	return &UploadAPIKeyDefault{
		_statusCode: code,
	}
}

/*UploadAPIKeyDefault handles this case with default header values.

An error has occurred.

*/
type UploadAPIKeyDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the upload Api key default response
func (o *UploadAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *UploadAPIKeyDefault) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/apiKeys/][%d] UploadApiKey default  %+v", o._statusCode, o.Payload)
}

func (o *UploadAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
