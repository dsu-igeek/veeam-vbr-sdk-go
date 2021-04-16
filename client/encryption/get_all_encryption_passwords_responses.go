// Code generated by go-swagger; DO NOT EDIT.

package encryption

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// GetAllEncryptionPasswordsReader is a Reader for the GetAllEncryptionPasswords structure.
type GetAllEncryptionPasswordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllEncryptionPasswordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllEncryptionPasswordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllEncryptionPasswordsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllEncryptionPasswordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllEncryptionPasswordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllEncryptionPasswordsOK creates a GetAllEncryptionPasswordsOK with default headers values
func NewGetAllEncryptionPasswordsOK() *GetAllEncryptionPasswordsOK {
	return &GetAllEncryptionPasswordsOK{}
}

/* GetAllEncryptionPasswordsOK describes a response with status code 200, with default header values.

OK
*/
type GetAllEncryptionPasswordsOK struct {
	Payload *models.EncryptionPasswordsResult
}

func (o *GetAllEncryptionPasswordsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/encryptionPasswords][%d] getAllEncryptionPasswordsOK  %+v", 200, o.Payload)
}
func (o *GetAllEncryptionPasswordsOK) GetPayload() *models.EncryptionPasswordsResult {
	return o.Payload
}

func (o *GetAllEncryptionPasswordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EncryptionPasswordsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllEncryptionPasswordsUnauthorized creates a GetAllEncryptionPasswordsUnauthorized with default headers values
func NewGetAllEncryptionPasswordsUnauthorized() *GetAllEncryptionPasswordsUnauthorized {
	return &GetAllEncryptionPasswordsUnauthorized{}
}

/* GetAllEncryptionPasswordsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type GetAllEncryptionPasswordsUnauthorized struct {
	Payload *models.Error
}

func (o *GetAllEncryptionPasswordsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/encryptionPasswords][%d] getAllEncryptionPasswordsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAllEncryptionPasswordsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllEncryptionPasswordsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllEncryptionPasswordsForbidden creates a GetAllEncryptionPasswordsForbidden with default headers values
func NewGetAllEncryptionPasswordsForbidden() *GetAllEncryptionPasswordsForbidden {
	return &GetAllEncryptionPasswordsForbidden{}
}

/* GetAllEncryptionPasswordsForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type GetAllEncryptionPasswordsForbidden struct {
	Payload *models.Error
}

func (o *GetAllEncryptionPasswordsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/encryptionPasswords][%d] getAllEncryptionPasswordsForbidden  %+v", 403, o.Payload)
}
func (o *GetAllEncryptionPasswordsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllEncryptionPasswordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllEncryptionPasswordsInternalServerError creates a GetAllEncryptionPasswordsInternalServerError with default headers values
func NewGetAllEncryptionPasswordsInternalServerError() *GetAllEncryptionPasswordsInternalServerError {
	return &GetAllEncryptionPasswordsInternalServerError{}
}

/* GetAllEncryptionPasswordsInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type GetAllEncryptionPasswordsInternalServerError struct {
	Payload *models.Error
}

func (o *GetAllEncryptionPasswordsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/encryptionPasswords][%d] getAllEncryptionPasswordsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAllEncryptionPasswordsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllEncryptionPasswordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
