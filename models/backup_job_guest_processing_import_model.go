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

// BackupJobGuestProcessingImportModel Guest processing settings.
//
// swagger:model BackupJobGuestProcessingImportModel
type BackupJobGuestProcessingImportModel struct {

	// application aware processing
	// Required: true
	ApplicationAwareProcessing *BackupApplicationAwareProcessingImportModel `json:"applicationAwareProcessing"`

	// guest credentials
	GuestCredentials *GuestOsCredentialsImportModel `json:"guestCredentials,omitempty"`

	// guest file system indexing
	// Required: true
	GuestFileSystemIndexing *GuestFileSystemIndexingModel `json:"guestFileSystemIndexing"`

	// guest interaction proxies
	GuestInteractionProxies *GuestInteractionProxiesSettingsImportModel `json:"guestInteractionProxies,omitempty"`
}

// Validate validates this backup job guest processing import model
func (m *BackupJobGuestProcessingImportModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationAwareProcessing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestFileSystemIndexing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestInteractionProxies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupJobGuestProcessingImportModel) validateApplicationAwareProcessing(formats strfmt.Registry) error {

	if err := validate.Required("applicationAwareProcessing", "body", m.ApplicationAwareProcessing); err != nil {
		return err
	}

	if m.ApplicationAwareProcessing != nil {
		if err := m.ApplicationAwareProcessing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationAwareProcessing")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobGuestProcessingImportModel) validateGuestCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.GuestCredentials) { // not required
		return nil
	}

	if m.GuestCredentials != nil {
		if err := m.GuestCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guestCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobGuestProcessingImportModel) validateGuestFileSystemIndexing(formats strfmt.Registry) error {

	if err := validate.Required("guestFileSystemIndexing", "body", m.GuestFileSystemIndexing); err != nil {
		return err
	}

	if m.GuestFileSystemIndexing != nil {
		if err := m.GuestFileSystemIndexing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guestFileSystemIndexing")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobGuestProcessingImportModel) validateGuestInteractionProxies(formats strfmt.Registry) error {
	if swag.IsZero(m.GuestInteractionProxies) { // not required
		return nil
	}

	if m.GuestInteractionProxies != nil {
		if err := m.GuestInteractionProxies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guestInteractionProxies")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup job guest processing import model based on the context it is used
func (m *BackupJobGuestProcessingImportModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationAwareProcessing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuestCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuestFileSystemIndexing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuestInteractionProxies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupJobGuestProcessingImportModel) contextValidateApplicationAwareProcessing(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationAwareProcessing != nil {
		if err := m.ApplicationAwareProcessing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationAwareProcessing")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobGuestProcessingImportModel) contextValidateGuestCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.GuestCredentials != nil {
		if err := m.GuestCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guestCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobGuestProcessingImportModel) contextValidateGuestFileSystemIndexing(ctx context.Context, formats strfmt.Registry) error {

	if m.GuestFileSystemIndexing != nil {
		if err := m.GuestFileSystemIndexing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guestFileSystemIndexing")
			}
			return err
		}
	}

	return nil
}

func (m *BackupJobGuestProcessingImportModel) contextValidateGuestInteractionProxies(ctx context.Context, formats strfmt.Registry) error {

	if m.GuestInteractionProxies != nil {
		if err := m.GuestInteractionProxies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guestInteractionProxies")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupJobGuestProcessingImportModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupJobGuestProcessingImportModel) UnmarshalBinary(b []byte) error {
	var res BackupJobGuestProcessingImportModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}