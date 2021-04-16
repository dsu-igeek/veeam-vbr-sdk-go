// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GuestOsCredentialsModel VM custom credentials.
//
// swagger:model GuestOsCredentialsModel
type GuestOsCredentialsModel struct {

	// Individual credentials for VMs.
	// Unique: true
	CredentialsPerMachine []*GuestOsCredentialsPerMachineModel `json:"credentialsPerMachine"`

	// Credentials ID for Microsoft Windows VMs.
	// Required: true
	// Format: uuid
	CredsID *strfmt.UUID `json:"credsId"`

	// creds type
	// Required: true
	CredsType *ECredentialsType `json:"credsType"`
}

// Validate validates this guest os credentials model
func (m *GuestOsCredentialsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsPerMachine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredsID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredsType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GuestOsCredentialsModel) validateCredentialsPerMachine(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsPerMachine) { // not required
		return nil
	}

	if err := validate.UniqueItems("credentialsPerMachine", "body", m.CredentialsPerMachine); err != nil {
		return err
	}

	for i := 0; i < len(m.CredentialsPerMachine); i++ {
		if swag.IsZero(m.CredentialsPerMachine[i]) { // not required
			continue
		}

		if m.CredentialsPerMachine[i] != nil {
			if err := m.CredentialsPerMachine[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentialsPerMachine" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GuestOsCredentialsModel) validateCredsID(formats strfmt.Registry) error {

	if err := validate.Required("credsId", "body", m.CredsID); err != nil {
		return err
	}

	if err := validate.FormatOf("credsId", "body", "uuid", m.CredsID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GuestOsCredentialsModel) validateCredsType(formats strfmt.Registry) error {

	if err := validate.Required("credsType", "body", m.CredsType); err != nil {
		return err
	}

	if err := validate.Required("credsType", "body", m.CredsType); err != nil {
		return err
	}

	if m.CredsType != nil {
		if err := m.CredsType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credsType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this guest os credentials model based on the context it is used
func (m *GuestOsCredentialsModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentialsPerMachine(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GuestOsCredentialsModel) contextValidateCredentialsPerMachine(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CredentialsPerMachine); i++ {

		if m.CredentialsPerMachine[i] != nil {
			if err := m.CredentialsPerMachine[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentialsPerMachine" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GuestOsCredentialsModel) contextValidateCredsType(ctx context.Context, formats strfmt.Registry) error {

	if m.CredsType != nil {
		if err := m.CredsType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credsType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GuestOsCredentialsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GuestOsCredentialsModel) UnmarshalBinary(b []byte) error {
	var res GuestOsCredentialsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
