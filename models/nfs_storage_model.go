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

// NfsStorageModel Network attached storage.
//
// swagger:model NfsStorageModel
type NfsStorageModel struct {
	RepositoryModel

	// mount server
	// Required: true
	MountServer *MountServerSettingsModel `json:"mountServer"`

	// repository
	// Required: true
	Repository *NetworkRepositorySettingsModel `json:"repository"`

	// share
	// Required: true
	Share *NfsRepositoryShareSettingsModel `json:"share"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NfsStorageModel) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RepositoryModel
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RepositoryModel = aO0

	// AO1
	var dataAO1 struct {
		MountServer *MountServerSettingsModel `json:"mountServer"`

		Repository *NetworkRepositorySettingsModel `json:"repository"`

		Share *NfsRepositoryShareSettingsModel `json:"share"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.MountServer = dataAO1.MountServer

	m.Repository = dataAO1.Repository

	m.Share = dataAO1.Share

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NfsStorageModel) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.RepositoryModel)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		MountServer *MountServerSettingsModel `json:"mountServer"`

		Repository *NetworkRepositorySettingsModel `json:"repository"`

		Share *NfsRepositoryShareSettingsModel `json:"share"`
	}

	dataAO1.MountServer = m.MountServer

	dataAO1.Repository = m.Repository

	dataAO1.Share = m.Share

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this nfs storage model
func (m *NfsStorageModel) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RepositoryModel
	if err := m.RepositoryModel.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShare(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsStorageModel) validateMountServer(formats strfmt.Registry) error {

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

func (m *NfsStorageModel) validateRepository(formats strfmt.Registry) error {

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

func (m *NfsStorageModel) validateShare(formats strfmt.Registry) error {

	if err := validate.Required("share", "body", m.Share); err != nil {
		return err
	}

	if m.Share != nil {
		if err := m.Share.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("share")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nfs storage model based on the context it is used
func (m *NfsStorageModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RepositoryModel
	if err := m.RepositoryModel.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMountServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShare(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsStorageModel) contextValidateMountServer(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NfsStorageModel) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NfsStorageModel) contextValidateShare(ctx context.Context, formats strfmt.Registry) error {

	if m.Share != nil {
		if err := m.Share.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("share")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NfsStorageModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NfsStorageModel) UnmarshalBinary(b []byte) error {
	var res NfsStorageModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
