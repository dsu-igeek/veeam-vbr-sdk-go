// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EConfigBackupSMTPSettingsType Type of notification settings.
//
// swagger:model EConfigBackupSMTPSettingsType
type EConfigBackupSMTPSettingsType string

func NewEConfigBackupSMTPSettingsType(value EConfigBackupSMTPSettingsType) *EConfigBackupSMTPSettingsType {
	v := value
	return &v
}

const (

	// EConfigBackupSMTPSettingsTypeGlobal captures enum value "Global"
	EConfigBackupSMTPSettingsTypeGlobal EConfigBackupSMTPSettingsType = "Global"

	// EConfigBackupSMTPSettingsTypeCustom captures enum value "Custom"
	EConfigBackupSMTPSettingsTypeCustom EConfigBackupSMTPSettingsType = "Custom"
)

// for schema
var eConfigBackupSmtpSettingsTypeEnum []interface{}

func init() {
	var res []EConfigBackupSMTPSettingsType
	if err := json.Unmarshal([]byte(`["Global","Custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eConfigBackupSmtpSettingsTypeEnum = append(eConfigBackupSmtpSettingsTypeEnum, v)
	}
}

func (m EConfigBackupSMTPSettingsType) validateEConfigBackupSMTPSettingsTypeEnum(path, location string, value EConfigBackupSMTPSettingsType) error {
	if err := validate.EnumCase(path, location, value, eConfigBackupSmtpSettingsTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e config backup SMTP settings type
func (m EConfigBackupSMTPSettingsType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEConfigBackupSMTPSettingsTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e config backup SMTP settings type based on context it is used
func (m EConfigBackupSMTPSettingsType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
