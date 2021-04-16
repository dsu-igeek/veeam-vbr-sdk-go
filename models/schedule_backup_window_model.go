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

// ScheduleBackupWindowModel schedule backup window model
//
// swagger:model ScheduleBackupWindowModel
type ScheduleBackupWindowModel struct {

	// window setting
	WindowSetting *BackupWindowSettingModel `json:"WindowSetting,omitempty"`

	// If *true*, periodic schedule is enabled.
	// Required: true
	IsEnabled bool `json:"isEnabled"`
}

// Validate validates this schedule backup window model
func (m *ScheduleBackupWindowModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWindowSetting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleBackupWindowModel) validateWindowSetting(formats strfmt.Registry) error {
	if swag.IsZero(m.WindowSetting) { // not required
		return nil
	}

	if m.WindowSetting != nil {
		if err := m.WindowSetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WindowSetting")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleBackupWindowModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", bool(m.IsEnabled)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this schedule backup window model based on the context it is used
func (m *ScheduleBackupWindowModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWindowSetting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleBackupWindowModel) contextValidateWindowSetting(ctx context.Context, formats strfmt.Registry) error {

	if m.WindowSetting != nil {
		if err := m.WindowSetting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WindowSetting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleBackupWindowModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleBackupWindowModel) UnmarshalBinary(b []byte) error {
	var res ScheduleBackupWindowModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
