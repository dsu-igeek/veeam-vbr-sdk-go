// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CredentialsImportModel credentials import model
//
// swagger:model CredentialsImportModel
type CredentialsImportModel struct {

	// User name, account name or access key.
	// Required: true
	CredentialsName *string `json:"credentialsName"`

	// Tag used to identify the credentials record.
	CredentialsTag string `json:"credentialsTag,omitempty"`
}

// Validate validates this credentials import model
func (m *CredentialsImportModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsImportModel) validateCredentialsName(formats strfmt.Registry) error {

	if err := validate.Required("credentialsName", "body", m.CredentialsName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this credentials import model based on context it is used
func (m *CredentialsImportModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CredentialsImportModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialsImportModel) UnmarshalBinary(b []byte) error {
	var res CredentialsImportModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}