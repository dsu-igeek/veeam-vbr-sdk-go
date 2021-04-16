// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationSettingsModel Notification settings for the backup job.
//
// swagger:model NotificationSettingsModel
type NotificationSettingsModel struct {

	// email notifications
	EmailNotifications *EmailNotificationSettingsModel `json:"emailNotifications,omitempty"`

	// If *true*, SNMP notifications are enabled for this job.
	SendSNMPNotifications bool `json:"sendSNMPNotifications,omitempty"`

	// vm attribute
	VMAttribute *NotificationVMAttributeSettingsModel `json:"vmAttribute,omitempty"`
}

// Validate validates this notification settings model
func (m *NotificationSettingsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMAttribute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationSettingsModel) validateEmailNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailNotifications) { // not required
		return nil
	}

	if m.EmailNotifications != nil {
		if err := m.EmailNotifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailNotifications")
			}
			return err
		}
	}

	return nil
}

func (m *NotificationSettingsModel) validateVMAttribute(formats strfmt.Registry) error {
	if swag.IsZero(m.VMAttribute) { // not required
		return nil
	}

	if m.VMAttribute != nil {
		if err := m.VMAttribute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmAttribute")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this notification settings model based on the context it is used
func (m *NotificationSettingsModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmailNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMAttribute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationSettingsModel) contextValidateEmailNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.EmailNotifications != nil {
		if err := m.EmailNotifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailNotifications")
			}
			return err
		}
	}

	return nil
}

func (m *NotificationSettingsModel) contextValidateVMAttribute(ctx context.Context, formats strfmt.Registry) error {

	if m.VMAttribute != nil {
		if err := m.VMAttribute.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmAttribute")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationSettingsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationSettingsModel) UnmarshalBinary(b []byte) error {
	var res NotificationSettingsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
