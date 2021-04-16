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

// GeneralOptionsModel general options model
//
// swagger:model GeneralOptionsModel
type GeneralOptionsModel struct {

	// email settings
	EmailSettings *GeneralOptionsEmailNotificationsModel `json:"emailSettings,omitempty"`

	// notifications
	Notifications *GeneralOptionsNotificationsModel `json:"notifications,omitempty"`
}

// Validate validates this general options model
func (m *GeneralOptionsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralOptionsModel) validateEmailSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailSettings) { // not required
		return nil
	}

	if m.EmailSettings != nil {
		if err := m.EmailSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailSettings")
			}
			return err
		}
	}

	return nil
}

func (m *GeneralOptionsModel) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this general options model based on the context it is used
func (m *GeneralOptionsModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmailSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralOptionsModel) contextValidateEmailSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.EmailSettings != nil {
		if err := m.EmailSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailSettings")
			}
			return err
		}
	}

	return nil
}

func (m *GeneralOptionsModel) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifications != nil {
		if err := m.Notifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeneralOptionsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralOptionsModel) UnmarshalBinary(b []byte) error {
	var res GeneralOptionsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
