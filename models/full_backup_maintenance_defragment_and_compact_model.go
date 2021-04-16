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

// FullBackupMaintenanceDefragmentAndCompactModel Compact operation settings.
//
// swagger:model FullBackupMaintenanceDefragmentAndCompactModel
type FullBackupMaintenanceDefragmentAndCompactModel struct {

	// If *true*, Veeam Backup & Replication compacts full backup files.
	// Required: true
	IsEnabled *bool `json:"isEnabled"`

	// monthly
	Monthly *AdvancedStorageScheduleMonthlyModel `json:"monthly,omitempty"`

	// weekly
	Weekly *AdvancedStorageScheduleWeeklyModel `json:"weekly,omitempty"`
}

// Validate validates this full backup maintenance defragment and compact model
func (m *FullBackupMaintenanceDefragmentAndCompactModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonthly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekly(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullBackupMaintenanceDefragmentAndCompactModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *FullBackupMaintenanceDefragmentAndCompactModel) validateMonthly(formats strfmt.Registry) error {
	if swag.IsZero(m.Monthly) { // not required
		return nil
	}

	if m.Monthly != nil {
		if err := m.Monthly.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monthly")
			}
			return err
		}
	}

	return nil
}

func (m *FullBackupMaintenanceDefragmentAndCompactModel) validateWeekly(formats strfmt.Registry) error {
	if swag.IsZero(m.Weekly) { // not required
		return nil
	}

	if m.Weekly != nil {
		if err := m.Weekly.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weekly")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this full backup maintenance defragment and compact model based on the context it is used
func (m *FullBackupMaintenanceDefragmentAndCompactModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMonthly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeekly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullBackupMaintenanceDefragmentAndCompactModel) contextValidateMonthly(ctx context.Context, formats strfmt.Registry) error {

	if m.Monthly != nil {
		if err := m.Monthly.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monthly")
			}
			return err
		}
	}

	return nil
}

func (m *FullBackupMaintenanceDefragmentAndCompactModel) contextValidateWeekly(ctx context.Context, formats strfmt.Registry) error {

	if m.Weekly != nil {
		if err := m.Weekly.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weekly")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FullBackupMaintenanceDefragmentAndCompactModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullBackupMaintenanceDefragmentAndCompactModel) UnmarshalBinary(b []byte) error {
	var res FullBackupMaintenanceDefragmentAndCompactModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
