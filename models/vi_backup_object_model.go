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

// ViBackupObjectModel vi backup object model
//
// swagger:model ViBackupObjectModel
type ViBackupObjectModel struct {
	BackupObjectModel

	// ID of the virtual infrastructure object: mo-ref or ID, depending on the virtualization platform.
	//
	ObjectID string `json:"objectId,omitempty"`

	// Path to the object.
	Path string `json:"path,omitempty"`

	// vi type
	ViType EVmwareInventoryType `json:"viType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ViBackupObjectModel) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BackupObjectModel
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BackupObjectModel = aO0

	// AO1
	var dataAO1 struct {
		ObjectID string `json:"objectId,omitempty"`

		Path string `json:"path,omitempty"`

		ViType EVmwareInventoryType `json:"viType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ObjectID = dataAO1.ObjectID

	m.Path = dataAO1.Path

	m.ViType = dataAO1.ViType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ViBackupObjectModel) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BackupObjectModel)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ObjectID string `json:"objectId,omitempty"`

		Path string `json:"path,omitempty"`

		ViType EVmwareInventoryType `json:"viType,omitempty"`
	}

	dataAO1.ObjectID = m.ObjectID

	dataAO1.Path = m.Path

	dataAO1.ViType = m.ViType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vi backup object model
func (m *ViBackupObjectModel) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BackupObjectModel
	if err := m.BackupObjectModel.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViBackupObjectModel) validateViType(formats strfmt.Registry) error {

	if swag.IsZero(m.ViType) { // not required
		return nil
	}

	if err := m.ViType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("viType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this vi backup object model based on the context it is used
func (m *ViBackupObjectModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BackupObjectModel
	if err := m.BackupObjectModel.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViBackupObjectModel) contextValidateViType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ViType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("viType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ViBackupObjectModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViBackupObjectModel) UnmarshalBinary(b []byte) error {
	var res ViBackupObjectModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}