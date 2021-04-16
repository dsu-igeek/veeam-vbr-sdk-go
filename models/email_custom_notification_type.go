// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailCustomNotificationType email custom notification type
//
// swagger:model EmailCustomNotificationType
type EmailCustomNotificationType struct {

	// If *true*, email notifications are sent about the final job status only (not per every job retry).
	SuppressNotificationUntilLastRetry bool `json:"SuppressNotificationUntilLastRetry,omitempty"`

	// If *true*, email notifications are sent when the job fails.
	NotifyOnError bool `json:"notifyOnError,omitempty"`

	// If *true*, email notifications are sent when the job completes successfully.
	NotifyOnSuccess bool `json:"notifyOnSuccess,omitempty"`

	// If *true*, email notifications are sent when the job completes with a warning.
	NotifyOnWarning bool `json:"notifyOnWarning,omitempty"`

	// Notification subject. Use the following variables in the subject: *%Time%* (completion time), *%JobName%*, *%JobResult%*, *%ObjectCount%* (number of VMs in the job) and *%Issues%* (number of VMs in the job that have finished with the Warning or Failed status).
	//
	Subject string `json:"subject,omitempty"`
}

// Validate validates this email custom notification type
func (m *EmailCustomNotificationType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this email custom notification type based on context it is used
func (m *EmailCustomNotificationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailCustomNotificationType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailCustomNotificationType) UnmarshalBinary(b []byte) error {
	var res EmailCustomNotificationType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
