// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// GetBackupReader is a Reader for the GetBackup structure.
type GetBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBackupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBackupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBackupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBackupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupOK creates a GetBackupOK with default headers values
func NewGetBackupOK() *GetBackupOK {
	return &GetBackupOK{}
}

/* GetBackupOK describes a response with status code 200, with default header values.

OK
*/
type GetBackupOK struct {
	Payload *models.BackupModel
}

func (o *GetBackupOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/backups/{id}][%d] getBackupOK  %+v", 200, o.Payload)
}
func (o *GetBackupOK) GetPayload() *models.BackupModel {
	return o.Payload
}

func (o *GetBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupUnauthorized creates a GetBackupUnauthorized with default headers values
func NewGetBackupUnauthorized() *GetBackupUnauthorized {
	return &GetBackupUnauthorized{}
}

/* GetBackupUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type GetBackupUnauthorized struct {
	Payload *models.Error
}

func (o *GetBackupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/backups/{id}][%d] getBackupUnauthorized  %+v", 401, o.Payload)
}
func (o *GetBackupUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupForbidden creates a GetBackupForbidden with default headers values
func NewGetBackupForbidden() *GetBackupForbidden {
	return &GetBackupForbidden{}
}

/* GetBackupForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type GetBackupForbidden struct {
	Payload *models.Error
}

func (o *GetBackupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/backups/{id}][%d] getBackupForbidden  %+v", 403, o.Payload)
}
func (o *GetBackupForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupNotFound creates a GetBackupNotFound with default headers values
func NewGetBackupNotFound() *GetBackupNotFound {
	return &GetBackupNotFound{}
}

/* GetBackupNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type GetBackupNotFound struct {
	Payload *models.Error
}

func (o *GetBackupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/backups/{id}][%d] getBackupNotFound  %+v", 404, o.Payload)
}
func (o *GetBackupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupInternalServerError creates a GetBackupInternalServerError with default headers values
func NewGetBackupInternalServerError() *GetBackupInternalServerError {
	return &GetBackupInternalServerError{}
}

/* GetBackupInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type GetBackupInternalServerError struct {
	Payload *models.Error
}

func (o *GetBackupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/backups/{id}][%d] getBackupInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBackupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}