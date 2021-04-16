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

// GeneralOptionsEmailNotificationsModel Global email notification settings.
//
// swagger:model GeneralOptionsEmailNotificationsModel
type GeneralOptionsEmailNotificationsModel struct {

	// advanced Smtp options
	// Required: true
	AdvancedSMTPOptions *AdvancedSMTPOptionsModel `json:"advancedSmtpOptions"`

	// Email address from which email notifications must be sent.
	// Required: true
	From *string `json:"from"`

	// If *true*, global email notification settings are enabled.
	// Required: true
	IsEnabled *bool `json:"isEnabled"`

	// If *true*, email notifications are sent when the job fails.
	// Required: true
	NotifyOnFailure *bool `json:"notifyOnFailure"`

	// If *true*, email notifications are sent about the final job status only (not per every job retry).
	// Required: true
	NotifyOnLastRetry *bool `json:"notifyOnLastRetry"`

	// If *true*, email notifications are sent when the job completes successfully.
	// Required: true
	NotifyOnSuccess *bool `json:"notifyOnSuccess"`

	// If *true*, email notifications are sent when the job completes with a warning.
	// Required: true
	NotifyOnWarning *bool `json:"notifyOnWarning"`

	// Time when Veeam Backup & Replication sends daily email reports.
	// Required: true
	// Format: date-time
	SendDailyReportsAt *strfmt.DateTime `json:"sendDailyReportsAt"`

	// Full DNS name or IP address of the SMTP server.
	// Required: true
	SMTPServerName *string `json:"smtpServerName"`

	// Notification subject. Use the following variables in the subject:<br> <ol>
	//   <li>%Time% — completion time</li>
	//   <li>%JobName%</li>
	//   <li>%JobResult%</li>
	//   <li>%ObjectCount% — number of VMs in the job</li>
	//   <li>%Issues% — number of VMs in the job that have been processed with the Warning or Failed status</li>
	// </ol>
	//
	// Required: true
	Subject *string `json:"subject"`

	// Recipient email addresses. Use a semicolon to separate multiple addresses.
	// Required: true
	To *string `json:"to"`
}

// Validate validates this general options email notifications model
func (m *GeneralOptionsEmailNotificationsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedSMTPOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyOnFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyOnLastRetry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyOnSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyOnWarning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSendDailyReportsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateAdvancedSMTPOptions(formats strfmt.Registry) error {

	if err := validate.Required("advancedSmtpOptions", "body", m.AdvancedSMTPOptions); err != nil {
		return err
	}

	if m.AdvancedSMTPOptions != nil {
		if err := m.AdvancedSMTPOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advancedSmtpOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateNotifyOnFailure(formats strfmt.Registry) error {

	if err := validate.Required("notifyOnFailure", "body", m.NotifyOnFailure); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateNotifyOnLastRetry(formats strfmt.Registry) error {

	if err := validate.Required("notifyOnLastRetry", "body", m.NotifyOnLastRetry); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateNotifyOnSuccess(formats strfmt.Registry) error {

	if err := validate.Required("notifyOnSuccess", "body", m.NotifyOnSuccess); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateNotifyOnWarning(formats strfmt.Registry) error {

	if err := validate.Required("notifyOnWarning", "body", m.NotifyOnWarning); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateSendDailyReportsAt(formats strfmt.Registry) error {

	if err := validate.Required("sendDailyReportsAt", "body", m.SendDailyReportsAt); err != nil {
		return err
	}

	if err := validate.FormatOf("sendDailyReportsAt", "body", "date-time", m.SendDailyReportsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateSMTPServerName(formats strfmt.Registry) error {

	if err := validate.Required("smtpServerName", "body", m.SMTPServerName); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this general options email notifications model based on the context it is used
func (m *GeneralOptionsEmailNotificationsModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvancedSMTPOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralOptionsEmailNotificationsModel) contextValidateAdvancedSMTPOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.AdvancedSMTPOptions != nil {
		if err := m.AdvancedSMTPOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advancedSmtpOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeneralOptionsEmailNotificationsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralOptionsEmailNotificationsModel) UnmarshalBinary(b []byte) error {
	var res GeneralOptionsEmailNotificationsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
