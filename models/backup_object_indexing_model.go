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

// BackupObjectIndexingModel Guest OS indexing options for the VM.
//
// swagger:model BackupObjectIndexingModel
type BackupObjectIndexingModel struct {

	// guest f s indexing mode
	// Required: true
	GuestFSIndexingMode *EGuestFSIndexingMode `json:"guestFSIndexingMode"`

	// Array of folders. Environmental variables and full paths to folders can be used.
	IndexingList []string `json:"indexingList"`
}

// Validate validates this backup object indexing model
func (m *BackupObjectIndexingModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGuestFSIndexingMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupObjectIndexingModel) validateGuestFSIndexingMode(formats strfmt.Registry) error {

	if err := validate.Required("guestFSIndexingMode", "body", m.GuestFSIndexingMode); err != nil {
		return err
	}

	if err := validate.Required("guestFSIndexingMode", "body", m.GuestFSIndexingMode); err != nil {
		return err
	}

	if m.GuestFSIndexingMode != nil {
		if err := m.GuestFSIndexingMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guestFSIndexingMode")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup object indexing model based on the context it is used
func (m *BackupObjectIndexingModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGuestFSIndexingMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupObjectIndexingModel) contextValidateGuestFSIndexingMode(ctx context.Context, formats strfmt.Registry) error {

	if m.GuestFSIndexingMode != nil {
		if err := m.GuestFSIndexingMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guestFSIndexingMode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupObjectIndexingModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupObjectIndexingModel) UnmarshalBinary(b []byte) error {
	var res BackupObjectIndexingModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
