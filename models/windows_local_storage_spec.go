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

// WindowsLocalStorageSpec windows local storage spec
//
// swagger:model WindowsLocalStorageSpec
type WindowsLocalStorageSpec struct {
	RepositorySpec

	// ID of the server that is used as a backup repository.
	// Required: true
	// Format: uuid
	HostID *strfmt.UUID `json:"hostId"`

	// mount server
	// Required: true
	MountServer *MountServerSettingsModel `json:"mountServer"`

	// repository
	// Required: true
	Repository *WindowsLocalRepositorySettingsModel `json:"repository"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WindowsLocalStorageSpec) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RepositorySpec
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RepositorySpec = aO0

	// AO1
	var dataAO1 struct {
		HostID *strfmt.UUID `json:"hostId"`

		MountServer *MountServerSettingsModel `json:"mountServer"`

		Repository *WindowsLocalRepositorySettingsModel `json:"repository"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.HostID = dataAO1.HostID

	m.MountServer = dataAO1.MountServer

	m.Repository = dataAO1.Repository

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WindowsLocalStorageSpec) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.RepositorySpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		HostID *strfmt.UUID `json:"hostId"`

		MountServer *MountServerSettingsModel `json:"mountServer"`

		Repository *WindowsLocalRepositorySettingsModel `json:"repository"`
	}

	dataAO1.HostID = m.HostID

	dataAO1.MountServer = m.MountServer

	dataAO1.Repository = m.Repository

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this windows local storage spec
func (m *WindowsLocalStorageSpec) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RepositorySpec
	if err := m.RepositorySpec.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WindowsLocalStorageSpec) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("hostId", "body", m.HostID); err != nil {
		return err
	}

	if err := validate.FormatOf("hostId", "body", "uuid", m.HostID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WindowsLocalStorageSpec) validateMountServer(formats strfmt.Registry) error {

	if err := validate.Required("mountServer", "body", m.MountServer); err != nil {
		return err
	}

	if m.MountServer != nil {
		if err := m.MountServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mountServer")
			}
			return err
		}
	}

	return nil
}

func (m *WindowsLocalStorageSpec) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this windows local storage spec based on the context it is used
func (m *WindowsLocalStorageSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RepositorySpec
	if err := m.RepositorySpec.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMountServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WindowsLocalStorageSpec) contextValidateMountServer(ctx context.Context, formats strfmt.Registry) error {

	if m.MountServer != nil {
		if err := m.MountServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mountServer")
			}
			return err
		}
	}

	return nil
}

func (m *WindowsLocalStorageSpec) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.Repository != nil {
		if err := m.Repository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WindowsLocalStorageSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WindowsLocalStorageSpec) UnmarshalBinary(b []byte) error {
	var res WindowsLocalStorageSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
