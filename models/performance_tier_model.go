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
)

// PerformanceTierModel Performance tier.
//
// swagger:model PerformanceTierModel
type PerformanceTierModel struct {

	// advanced settings
	AdvancedSettings *PerformanceTierAdvancedSettingsModel `json:"advancedSettings,omitempty"`

	// Array of performance extents.
	PerformanceExtents []*PerformanceExtentModel `json:"performanceExtents"`
}

// Validate validates this performance tier model
func (m *PerformanceTierModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceExtents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceTierModel) validateAdvancedSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvancedSettings) { // not required
		return nil
	}

	if m.AdvancedSettings != nil {
		if err := m.AdvancedSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advancedSettings")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceTierModel) validatePerformanceExtents(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceExtents) { // not required
		return nil
	}

	for i := 0; i < len(m.PerformanceExtents); i++ {
		if swag.IsZero(m.PerformanceExtents[i]) { // not required
			continue
		}

		if m.PerformanceExtents[i] != nil {
			if err := m.PerformanceExtents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("performanceExtents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance tier model based on the context it is used
func (m *PerformanceTierModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvancedSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformanceExtents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceTierModel) contextValidateAdvancedSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.AdvancedSettings != nil {
		if err := m.AdvancedSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advancedSettings")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceTierModel) contextValidatePerformanceExtents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PerformanceExtents); i++ {

		if m.PerformanceExtents[i] != nil {
			if err := m.PerformanceExtents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("performanceExtents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceTierModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceTierModel) UnmarshalBinary(b []byte) error {
	var res PerformanceTierModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
