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

// BackupPlacementSettingsModel Settings of the performance placement policy.
//
// swagger:model BackupPlacementSettingsModel
type BackupPlacementSettingsModel struct {

	// allowed backups
	// Required: true
	AllowedBackups *EAllowedBackupsType `json:"allowedBackups"`

	// Name of a performance extent.
	// Required: true
	ExtentName *string `json:"extentName"`
}

// Validate validates this backup placement settings model
func (m *BackupPlacementSettingsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedBackups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlacementSettingsModel) validateAllowedBackups(formats strfmt.Registry) error {

	if err := validate.Required("allowedBackups", "body", m.AllowedBackups); err != nil {
		return err
	}

	if err := validate.Required("allowedBackups", "body", m.AllowedBackups); err != nil {
		return err
	}

	if m.AllowedBackups != nil {
		if err := m.AllowedBackups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowedBackups")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlacementSettingsModel) validateExtentName(formats strfmt.Registry) error {

	if err := validate.Required("extentName", "body", m.ExtentName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this backup placement settings model based on the context it is used
func (m *BackupPlacementSettingsModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowedBackups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlacementSettingsModel) contextValidateAllowedBackups(ctx context.Context, formats strfmt.Registry) error {

	if m.AllowedBackups != nil {
		if err := m.AllowedBackups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowedBackups")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupPlacementSettingsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPlacementSettingsModel) UnmarshalBinary(b []byte) error {
	var res BackupPlacementSettingsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
