// Code generated by go-swagger; DO NOT EDIT.

package managed_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// CreateManagedServerReader is a Reader for the CreateManagedServer structure.
type CreateManagedServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateManagedServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateManagedServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateManagedServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateManagedServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateManagedServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateManagedServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateManagedServerCreated creates a CreateManagedServerCreated with default headers values
func NewCreateManagedServerCreated() *CreateManagedServerCreated {
	return &CreateManagedServerCreated{}
}

/* CreateManagedServerCreated describes a response with status code 201, with default header values.

Infrastructure session has been created to add the server. To check the progress, track the session `state`.
*/
type CreateManagedServerCreated struct {
	Payload *models.SessionModel
}

func (o *CreateManagedServerCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/managedServers][%d] createManagedServerCreated  %+v", 201, o.Payload)
}
func (o *CreateManagedServerCreated) GetPayload() *models.SessionModel {
	return o.Payload
}

func (o *CreateManagedServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateManagedServerBadRequest creates a CreateManagedServerBadRequest with default headers values
func NewCreateManagedServerBadRequest() *CreateManagedServerBadRequest {
	return &CreateManagedServerBadRequest{}
}

/* CreateManagedServerBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type CreateManagedServerBadRequest struct {
	Payload *models.Error
}

func (o *CreateManagedServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/managedServers][%d] createManagedServerBadRequest  %+v", 400, o.Payload)
}
func (o *CreateManagedServerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateManagedServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateManagedServerUnauthorized creates a CreateManagedServerUnauthorized with default headers values
func NewCreateManagedServerUnauthorized() *CreateManagedServerUnauthorized {
	return &CreateManagedServerUnauthorized{}
}

/* CreateManagedServerUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type CreateManagedServerUnauthorized struct {
	Payload *models.Error
}

func (o *CreateManagedServerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/managedServers][%d] createManagedServerUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateManagedServerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateManagedServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateManagedServerForbidden creates a CreateManagedServerForbidden with default headers values
func NewCreateManagedServerForbidden() *CreateManagedServerForbidden {
	return &CreateManagedServerForbidden{}
}

/* CreateManagedServerForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type CreateManagedServerForbidden struct {
	Payload *models.Error
}

func (o *CreateManagedServerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/managedServers][%d] createManagedServerForbidden  %+v", 403, o.Payload)
}
func (o *CreateManagedServerForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateManagedServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateManagedServerInternalServerError creates a CreateManagedServerInternalServerError with default headers values
func NewCreateManagedServerInternalServerError() *CreateManagedServerInternalServerError {
	return &CreateManagedServerInternalServerError{}
}

/* CreateManagedServerInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type CreateManagedServerInternalServerError struct {
	Payload *models.Error
}

func (o *CreateManagedServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/managedServers][%d] createManagedServerInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateManagedServerInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateManagedServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}