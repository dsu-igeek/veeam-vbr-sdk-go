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

// ConfigBackupEncryptionModel Encryption settings.
//
// swagger:model ConfigBackupEncryptionModel
type ConfigBackupEncryptionModel struct {

	// If *true*, backup encryption is enabled.
	// Required: true
	IsEnabled *bool `json:"isEnabled"`

	// ID of the password used for encryption.
	// Required: true
	// Format: uuid
	PasswordID *strfmt.UUID `json:"passwordId"`
}

// Validate validates this config backup encryption model
func (m *ConfigBackupEncryptionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBackupEncryptionModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *ConfigBackupEncryptionModel) validatePasswordID(formats strfmt.Registry) error {

	if err := validate.Required("passwordId", "body", m.PasswordID); err != nil {
		return err
	}

	if err := validate.FormatOf("passwordId", "body", "uuid", m.PasswordID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this config backup encryption model based on context it is used
func (m *ConfigBackupEncryptionModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigBackupEncryptionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigBackupEncryptionModel) UnmarshalBinary(b []byte) error {
	var res ConfigBackupEncryptionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
