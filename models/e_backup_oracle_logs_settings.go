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

// EBackupOracleLogsSettings Type of archived logs processing.
//
// swagger:model EBackupOracleLogsSettings
type EBackupOracleLogsSettings string

func NewEBackupOracleLogsSettings(value EBackupOracleLogsSettings) *EBackupOracleLogsSettings {
	v := value
	return &v
}

const (

	// EBackupOracleLogsSettingsPreserve captures enum value "preserve"
	EBackupOracleLogsSettingsPreserve EBackupOracleLogsSettings = "preserve"

	// EBackupOracleLogsSettingsDeleteExpiredHours captures enum value "deleteExpiredHours"
	EBackupOracleLogsSettingsDeleteExpiredHours EBackupOracleLogsSettings = "deleteExpiredHours"

	// EBackupOracleLogsSettingsDeleteExpiredGBs captures enum value "deleteExpiredGBs"
	EBackupOracleLogsSettingsDeleteExpiredGBs EBackupOracleLogsSettings = "deleteExpiredGBs"
)

// for schema
var eBackupOracleLogsSettingsEnum []interface{}

func init() {
	var res []EBackupOracleLogsSettings
	if err := json.Unmarshal([]byte(`["preserve","deleteExpiredHours","deleteExpiredGBs"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eBackupOracleLogsSettingsEnum = append(eBackupOracleLogsSettingsEnum, v)
	}
}

func (m EBackupOracleLogsSettings) validateEBackupOracleLogsSettingsEnum(path, location string, value EBackupOracleLogsSettings) error {
	if err := validate.EnumCase(path, location, value, eBackupOracleLogsSettingsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e backup oracle logs settings
func (m EBackupOracleLogsSettings) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEBackupOracleLogsSettingsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e backup oracle logs settings based on context it is used
func (m EBackupOracleLogsSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
