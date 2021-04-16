// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupApplicationAwareProcessingModel Application-aware processing settings.
//
// swagger:model BackupApplicationAwareProcessingModel
type BackupApplicationAwareProcessingModel struct {

	// Array of VMware objects and their application settings.
	AppSettings []*BackupApplicationSettingsModel `json:"appSettings"`

	// If *true*, application-aware processing is enabled.
	// Required: true
	IsEnabled *bool `json:"isEnabled"`
}

// Validate validates this backup application aware processing model
func (m *BackupApplicationAwareProcessingModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppSettings(formats); err != nil {
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

func (m *BackupApplicationAwareProcessingModel) validateAppSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.AppSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.AppSettings); i++ {
		if swag.IsZero(m.AppSettings[i]) { // not required
			continue
		}

		if m.AppSettings[i] != nil {
			if err := m.AppSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupApplicationAwareProcessingModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this backup application aware processing model based on the context it is used
func (m *BackupApplicationAwareProcessingModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupApplicationAwareProcessingModel) contextValidateAppSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppSettings); i++ {

		if m.AppSettings[i] != nil {
			if err := m.AppSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupApplicationAwareProcessingModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupApplicationAwareProcessingModel) UnmarshalBinary(b []byte) error {
	var res BackupApplicationAwareProcessingModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
