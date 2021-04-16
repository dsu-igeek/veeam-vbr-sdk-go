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

// DeleteJobReader is a Reader for the DeleteJob structure.
type DeleteJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteJobNoContent creates a DeleteJobNoContent with default headers values
func NewDeleteJobNoContent() *DeleteJobNoContent {
	return &DeleteJobNoContent{}
}

/* DeleteJobNoContent describes a response with status code 204, with default header values.

Removed.
*/
type DeleteJobNoContent struct {
	Payload models.EmptySuccessResponse
}

func (o *DeleteJobNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/jobs/{id}][%d] deleteJobNoContent  %+v", 204, o.Payload)
}
func (o *DeleteJobNoContent) GetPayload() models.EmptySuccessResponse {
	return o.Payload
}

func (o *DeleteJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJobUnauthorized creates a DeleteJobUnauthorized with default headers values
func NewDeleteJobUnauthorized() *DeleteJobUnauthorized {
	return &DeleteJobUnauthorized{}
}

/* DeleteJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type DeleteJobUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteJobUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/jobs/{id}][%d] deleteJobUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteJobUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJobForbidden creates a DeleteJobForbidden with default headers values
func NewDeleteJobForbidden() *DeleteJobForbidden {
	return &DeleteJobForbidden{}
}

/* DeleteJobForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type DeleteJobForbidden struct {
	Payload *models.Error
}

func (o *DeleteJobForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/jobs/{id}][%d] deleteJobForbidden  %+v", 403, o.Payload)
}
func (o *DeleteJobForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJobNotFound creates a DeleteJobNotFound with default headers values
func NewDeleteJobNotFound() *DeleteJobNotFound {
	return &DeleteJobNotFound{}
}

/* DeleteJobNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type DeleteJobNotFound struct {
	Payload *models.Error
}

func (o *DeleteJobNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/jobs/{id}][%d] deleteJobNotFound  %+v", 404, o.Payload)
}
func (o *DeleteJobNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJobInternalServerError creates a DeleteJobInternalServerError with default headers values
func NewDeleteJobInternalServerError() *DeleteJobInternalServerError {
	return &DeleteJobInternalServerError{}
}

/* DeleteJobInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type DeleteJobInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteJobInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/jobs/{id}][%d] deleteJobInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteJobInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}