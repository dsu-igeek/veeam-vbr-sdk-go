// Code generated by go-swagger; DO NOT EDIT.

package object_restore_points

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// GetObjectRestorePointReader is a Reader for the GetObjectRestorePoint structure.
type GetObjectRestorePointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetObjectRestorePointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetObjectRestorePointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetObjectRestorePointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetObjectRestorePointForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetObjectRestorePointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetObjectRestorePointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetObjectRestorePointOK creates a GetObjectRestorePointOK with default headers values
func NewGetObjectRestorePointOK() *GetObjectRestorePointOK {
	return &GetObjectRestorePointOK{}
}

/* GetObjectRestorePointOK describes a response with status code 200, with default header values.

OK
*/
type GetObjectRestorePointOK struct {
	Payload *models.ObjectRestorePointModel
}

func (o *GetObjectRestorePointOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/objectRestorePoints/{id}][%d] getObjectRestorePointOK  %+v", 200, o.Payload)
}
func (o *GetObjectRestorePointOK) GetPayload() *models.ObjectRestorePointModel {
	return o.Payload
}

func (o *GetObjectRestorePointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectRestorePointModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObjectRestorePointUnauthorized creates a GetObjectRestorePointUnauthorized with default headers values
func NewGetObjectRestorePointUnauthorized() *GetObjectRestorePointUnauthorized {
	return &GetObjectRestorePointUnauthorized{}
}

/* GetObjectRestorePointUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type GetObjectRestorePointUnauthorized struct {
	Payload *models.Error
}

func (o *GetObjectRestorePointUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/objectRestorePoints/{id}][%d] getObjectRestorePointUnauthorized  %+v", 401, o.Payload)
}
func (o *GetObjectRestorePointUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetObjectRestorePointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObjectRestorePointForbidden creates a GetObjectRestorePointForbidden with default headers values
func NewGetObjectRestorePointForbidden() *GetObjectRestorePointForbidden {
	return &GetObjectRestorePointForbidden{}
}

/* GetObjectRestorePointForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type GetObjectRestorePointForbidden struct {
	Payload *models.Error
}

func (o *GetObjectRestorePointForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/objectRestorePoints/{id}][%d] getObjectRestorePointForbidden  %+v", 403, o.Payload)
}
func (o *GetObjectRestorePointForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetObjectRestorePointForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObjectRestorePointNotFound creates a GetObjectRestorePointNotFound with default headers values
func NewGetObjectRestorePointNotFound() *GetObjectRestorePointNotFound {
	return &GetObjectRestorePointNotFound{}
}

/* GetObjectRestorePointNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type GetObjectRestorePointNotFound struct {
	Payload *models.Error
}

func (o *GetObjectRestorePointNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/objectRestorePoints/{id}][%d] getObjectRestorePointNotFound  %+v", 404, o.Payload)
}
func (o *GetObjectRestorePointNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetObjectRestorePointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObjectRestorePointInternalServerError creates a GetObjectRestorePointInternalServerError with default headers values
func NewGetObjectRestorePointInternalServerError() *GetObjectRestorePointInternalServerError {
	return &GetObjectRestorePointInternalServerError{}
}

/* GetObjectRestorePointInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type GetObjectRestorePointInternalServerError struct {
	Payload *models.Error
}

func (o *GetObjectRestorePointInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/objectRestorePoints/{id}][%d] getObjectRestorePointInternalServerError  %+v", 500, o.Payload)
}
func (o *GetObjectRestorePointInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetObjectRestorePointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}