// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// StopJobReader is a Reader for the StopJob structure.
type StopJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewStopJobCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStopJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopJobCreated creates a StopJobCreated with default headers values
func NewStopJobCreated() *StopJobCreated {
	return &StopJobCreated{}
}

/* StopJobCreated describes a response with status code 201, with default header values.

Job session has been stopped.
*/
type StopJobCreated struct {
	Payload *models.SessionModel
}

func (o *StopJobCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/jobs/{id}/stop][%d] stopJobCreated  %+v", 201, o.Payload)
}
func (o *StopJobCreated) GetPayload() *models.SessionModel {
	return o.Payload
}

func (o *StopJobCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopJobBadRequest creates a StopJobBadRequest with default headers values
func NewStopJobBadRequest() *StopJobBadRequest {
	return &StopJobBadRequest{}
}

/* StopJobBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type StopJobBadRequest struct {
	Payload *models.Error
}

func (o *StopJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/jobs/{id}/stop][%d] stopJobBadRequest  %+v", 400, o.Payload)
}
func (o *StopJobBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopJobUnauthorized creates a StopJobUnauthorized with default headers values
func NewStopJobUnauthorized() *StopJobUnauthorized {
	return &StopJobUnauthorized{}
}

/* StopJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type StopJobUnauthorized struct {
	Payload *models.Error
}

func (o *StopJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/jobs/{id}/stop][%d] stopJobUnauthorized  %+v", 401, o.Payload)
}
func (o *StopJobUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopJobForbidden creates a StopJobForbidden with default headers values
func NewStopJobForbidden() *StopJobForbidden {
	return &StopJobForbidden{}
}

/* StopJobForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type StopJobForbidden struct {
	Payload *models.Error
}

func (o *StopJobForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/jobs/{id}/stop][%d] stopJobForbidden  %+v", 403, o.Payload)
}
func (o *StopJobForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopJobNotFound creates a StopJobNotFound with default headers values
func NewStopJobNotFound() *StopJobNotFound {
	return &StopJobNotFound{}
}

/* StopJobNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type StopJobNotFound struct {
	Payload *models.Error
}

func (o *StopJobNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/jobs/{id}/stop][%d] stopJobNotFound  %+v", 404, o.Payload)
}
func (o *StopJobNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopJobInternalServerError creates a StopJobInternalServerError with default headers values
func NewStopJobInternalServerError() *StopJobInternalServerError {
	return &StopJobInternalServerError{}
}

/* StopJobInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type StopJobInternalServerError struct {
	Payload *models.Error
}

func (o *StopJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/jobs/{id}/stop][%d] stopJobInternalServerError  %+v", 500, o.Payload)
}
func (o *StopJobInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
