// Code generated by go-swagger; DO NOT EDIT.

package automation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// GetAutomationSessionReader is a Reader for the GetAutomationSession structure.
type GetAutomationSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAutomationSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAutomationSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAutomationSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAutomationSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAutomationSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAutomationSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAutomationSessionOK creates a GetAutomationSessionOK with default headers values
func NewGetAutomationSessionOK() *GetAutomationSessionOK {
	return &GetAutomationSessionOK{}
}

/* GetAutomationSessionOK describes a response with status code 200, with default header values.

OK
*/
type GetAutomationSessionOK struct {
	Payload *models.SessionModel
}

func (o *GetAutomationSessionOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/automation/sessions/{id}][%d] getAutomationSessionOK  %+v", 200, o.Payload)
}
func (o *GetAutomationSessionOK) GetPayload() *models.SessionModel {
	return o.Payload
}

func (o *GetAutomationSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutomationSessionUnauthorized creates a GetAutomationSessionUnauthorized with default headers values
func NewGetAutomationSessionUnauthorized() *GetAutomationSessionUnauthorized {
	return &GetAutomationSessionUnauthorized{}
}

/* GetAutomationSessionUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type GetAutomationSessionUnauthorized struct {
	Payload *models.Error
}

func (o *GetAutomationSessionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/automation/sessions/{id}][%d] getAutomationSessionUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAutomationSessionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAutomationSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutomationSessionForbidden creates a GetAutomationSessionForbidden with default headers values
func NewGetAutomationSessionForbidden() *GetAutomationSessionForbidden {
	return &GetAutomationSessionForbidden{}
}

/* GetAutomationSessionForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type GetAutomationSessionForbidden struct {
	Payload *models.Error
}

func (o *GetAutomationSessionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/automation/sessions/{id}][%d] getAutomationSessionForbidden  %+v", 403, o.Payload)
}
func (o *GetAutomationSessionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAutomationSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutomationSessionNotFound creates a GetAutomationSessionNotFound with default headers values
func NewGetAutomationSessionNotFound() *GetAutomationSessionNotFound {
	return &GetAutomationSessionNotFound{}
}

/* GetAutomationSessionNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type GetAutomationSessionNotFound struct {
	Payload *models.Error
}

func (o *GetAutomationSessionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/automation/sessions/{id}][%d] getAutomationSessionNotFound  %+v", 404, o.Payload)
}
func (o *GetAutomationSessionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAutomationSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutomationSessionInternalServerError creates a GetAutomationSessionInternalServerError with default headers values
func NewGetAutomationSessionInternalServerError() *GetAutomationSessionInternalServerError {
	return &GetAutomationSessionInternalServerError{}
}

/* GetAutomationSessionInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type GetAutomationSessionInternalServerError struct {
	Payload *models.Error
}

func (o *GetAutomationSessionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/automation/sessions/{id}][%d] getAutomationSessionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAutomationSessionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAutomationSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
