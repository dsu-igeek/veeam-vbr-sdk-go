// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// StopSessionReader is a Reader for the StopSession structure.
type StopSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopSessionOK creates a StopSessionOK with default headers values
func NewStopSessionOK() *StopSessionOK {
	return &StopSessionOK{}
}

/* StopSessionOK describes a response with status code 200, with default header values.

OK
*/
type StopSessionOK struct {
	Payload models.EmptySuccessResponse
}

func (o *StopSessionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/sessions/{id}/stop][%d] stopSessionOK  %+v", 200, o.Payload)
}
func (o *StopSessionOK) GetPayload() models.EmptySuccessResponse {
	return o.Payload
}

func (o *StopSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopSessionUnauthorized creates a StopSessionUnauthorized with default headers values
func NewStopSessionUnauthorized() *StopSessionUnauthorized {
	return &StopSessionUnauthorized{}
}

/* StopSessionUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type StopSessionUnauthorized struct {
	Payload *models.Error
}

func (o *StopSessionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/sessions/{id}/stop][%d] stopSessionUnauthorized  %+v", 401, o.Payload)
}
func (o *StopSessionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopSessionForbidden creates a StopSessionForbidden with default headers values
func NewStopSessionForbidden() *StopSessionForbidden {
	return &StopSessionForbidden{}
}

/* StopSessionForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type StopSessionForbidden struct {
	Payload *models.Error
}

func (o *StopSessionForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/sessions/{id}/stop][%d] stopSessionForbidden  %+v", 403, o.Payload)
}
func (o *StopSessionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopSessionNotFound creates a StopSessionNotFound with default headers values
func NewStopSessionNotFound() *StopSessionNotFound {
	return &StopSessionNotFound{}
}

/* StopSessionNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type StopSessionNotFound struct {
	Payload *models.Error
}

func (o *StopSessionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/sessions/{id}/stop][%d] stopSessionNotFound  %+v", 404, o.Payload)
}
func (o *StopSessionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopSessionInternalServerError creates a StopSessionInternalServerError with default headers values
func NewStopSessionInternalServerError() *StopSessionInternalServerError {
	return &StopSessionInternalServerError{}
}

/* StopSessionInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type StopSessionInternalServerError struct {
	Payload *models.Error
}

func (o *StopSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/sessions/{id}/stop][%d] stopSessionInternalServerError  %+v", 500, o.Payload)
}
func (o *StopSessionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
