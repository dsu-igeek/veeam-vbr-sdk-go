// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// EnableScaleOutExtentMaintenanceModeReader is a Reader for the EnableScaleOutExtentMaintenanceMode structure.
type EnableScaleOutExtentMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableScaleOutExtentMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEnableScaleOutExtentMaintenanceModeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnableScaleOutExtentMaintenanceModeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEnableScaleOutExtentMaintenanceModeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnableScaleOutExtentMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnableScaleOutExtentMaintenanceModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableScaleOutExtentMaintenanceModeCreated creates a EnableScaleOutExtentMaintenanceModeCreated with default headers values
func NewEnableScaleOutExtentMaintenanceModeCreated() *EnableScaleOutExtentMaintenanceModeCreated {
	return &EnableScaleOutExtentMaintenanceModeCreated{}
}

/* EnableScaleOutExtentMaintenanceModeCreated describes a response with status code 201, with default header values.

RepositoryMaintenance session has been created.
*/
type EnableScaleOutExtentMaintenanceModeCreated struct {
	Payload *models.SessionModel
}

func (o *EnableScaleOutExtentMaintenanceModeCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableMaintenanceMode][%d] enableScaleOutExtentMaintenanceModeCreated  %+v", 201, o.Payload)
}
func (o *EnableScaleOutExtentMaintenanceModeCreated) GetPayload() *models.SessionModel {
	return o.Payload
}

func (o *EnableScaleOutExtentMaintenanceModeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableScaleOutExtentMaintenanceModeBadRequest creates a EnableScaleOutExtentMaintenanceModeBadRequest with default headers values
func NewEnableScaleOutExtentMaintenanceModeBadRequest() *EnableScaleOutExtentMaintenanceModeBadRequest {
	return &EnableScaleOutExtentMaintenanceModeBadRequest{}
}

/* EnableScaleOutExtentMaintenanceModeBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type EnableScaleOutExtentMaintenanceModeBadRequest struct {
	Payload *models.Error
}

func (o *EnableScaleOutExtentMaintenanceModeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableMaintenanceMode][%d] enableScaleOutExtentMaintenanceModeBadRequest  %+v", 400, o.Payload)
}
func (o *EnableScaleOutExtentMaintenanceModeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnableScaleOutExtentMaintenanceModeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableScaleOutExtentMaintenanceModeUnauthorized creates a EnableScaleOutExtentMaintenanceModeUnauthorized with default headers values
func NewEnableScaleOutExtentMaintenanceModeUnauthorized() *EnableScaleOutExtentMaintenanceModeUnauthorized {
	return &EnableScaleOutExtentMaintenanceModeUnauthorized{}
}

/* EnableScaleOutExtentMaintenanceModeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type EnableScaleOutExtentMaintenanceModeUnauthorized struct {
	Payload *models.Error
}

func (o *EnableScaleOutExtentMaintenanceModeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableMaintenanceMode][%d] enableScaleOutExtentMaintenanceModeUnauthorized  %+v", 401, o.Payload)
}
func (o *EnableScaleOutExtentMaintenanceModeUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnableScaleOutExtentMaintenanceModeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableScaleOutExtentMaintenanceModeForbidden creates a EnableScaleOutExtentMaintenanceModeForbidden with default headers values
func NewEnableScaleOutExtentMaintenanceModeForbidden() *EnableScaleOutExtentMaintenanceModeForbidden {
	return &EnableScaleOutExtentMaintenanceModeForbidden{}
}

/* EnableScaleOutExtentMaintenanceModeForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type EnableScaleOutExtentMaintenanceModeForbidden struct {
	Payload *models.Error
}

func (o *EnableScaleOutExtentMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableMaintenanceMode][%d] enableScaleOutExtentMaintenanceModeForbidden  %+v", 403, o.Payload)
}
func (o *EnableScaleOutExtentMaintenanceModeForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnableScaleOutExtentMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableScaleOutExtentMaintenanceModeInternalServerError creates a EnableScaleOutExtentMaintenanceModeInternalServerError with default headers values
func NewEnableScaleOutExtentMaintenanceModeInternalServerError() *EnableScaleOutExtentMaintenanceModeInternalServerError {
	return &EnableScaleOutExtentMaintenanceModeInternalServerError{}
}

/* EnableScaleOutExtentMaintenanceModeInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type EnableScaleOutExtentMaintenanceModeInternalServerError struct {
	Payload *models.Error
}

func (o *EnableScaleOutExtentMaintenanceModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableMaintenanceMode][%d] enableScaleOutExtentMaintenanceModeInternalServerError  %+v", 500, o.Payload)
}
func (o *EnableScaleOutExtentMaintenanceModeInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnableScaleOutExtentMaintenanceModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}