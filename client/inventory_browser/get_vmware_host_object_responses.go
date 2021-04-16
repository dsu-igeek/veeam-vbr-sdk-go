// Code generated by go-swagger; DO NOT EDIT.

package inventory_browser

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// GetVmwareHostObjectReader is a Reader for the GetVmwareHostObject structure.
type GetVmwareHostObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVmwareHostObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVmwareHostObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVmwareHostObjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVmwareHostObjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVmwareHostObjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVmwareHostObjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVmwareHostObjectOK creates a GetVmwareHostObjectOK with default headers values
func NewGetVmwareHostObjectOK() *GetVmwareHostObjectOK {
	return &GetVmwareHostObjectOK{}
}

/* GetVmwareHostObjectOK describes a response with status code 200, with default header values.

OK
*/
type GetVmwareHostObjectOK struct {
	Payload *models.VCenterInventoryResult
}

func (o *GetVmwareHostObjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts/{name}][%d] getVmwareHostObjectOK  %+v", 200, o.Payload)
}
func (o *GetVmwareHostObjectOK) GetPayload() *models.VCenterInventoryResult {
	return o.Payload
}

func (o *GetVmwareHostObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VCenterInventoryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmwareHostObjectUnauthorized creates a GetVmwareHostObjectUnauthorized with default headers values
func NewGetVmwareHostObjectUnauthorized() *GetVmwareHostObjectUnauthorized {
	return &GetVmwareHostObjectUnauthorized{}
}

/* GetVmwareHostObjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type GetVmwareHostObjectUnauthorized struct {
	Payload *models.Error
}

func (o *GetVmwareHostObjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts/{name}][%d] getVmwareHostObjectUnauthorized  %+v", 401, o.Payload)
}
func (o *GetVmwareHostObjectUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVmwareHostObjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmwareHostObjectForbidden creates a GetVmwareHostObjectForbidden with default headers values
func NewGetVmwareHostObjectForbidden() *GetVmwareHostObjectForbidden {
	return &GetVmwareHostObjectForbidden{}
}

/* GetVmwareHostObjectForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type GetVmwareHostObjectForbidden struct {
	Payload *models.Error
}

func (o *GetVmwareHostObjectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts/{name}][%d] getVmwareHostObjectForbidden  %+v", 403, o.Payload)
}
func (o *GetVmwareHostObjectForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVmwareHostObjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmwareHostObjectNotFound creates a GetVmwareHostObjectNotFound with default headers values
func NewGetVmwareHostObjectNotFound() *GetVmwareHostObjectNotFound {
	return &GetVmwareHostObjectNotFound{}
}

/* GetVmwareHostObjectNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type GetVmwareHostObjectNotFound struct {
	Payload *models.Error
}

func (o *GetVmwareHostObjectNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts/{name}][%d] getVmwareHostObjectNotFound  %+v", 404, o.Payload)
}
func (o *GetVmwareHostObjectNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVmwareHostObjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmwareHostObjectInternalServerError creates a GetVmwareHostObjectInternalServerError with default headers values
func NewGetVmwareHostObjectInternalServerError() *GetVmwareHostObjectInternalServerError {
	return &GetVmwareHostObjectInternalServerError{}
}

/* GetVmwareHostObjectInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type GetVmwareHostObjectInternalServerError struct {
	Payload *models.Error
}

func (o *GetVmwareHostObjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts/{name}][%d] getVmwareHostObjectInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVmwareHostObjectInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVmwareHostObjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
