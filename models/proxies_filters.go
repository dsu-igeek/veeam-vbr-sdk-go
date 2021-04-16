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

// ProxiesFilters proxies filters
//
// swagger:model ProxiesFilters
type ProxiesFilters struct {

	// host Id filter
	// Format: uuid
	HostIDFilter strfmt.UUID `json:"hostIdFilter,omitempty"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// name filter
	NameFilter string `json:"nameFilter,omitempty"`

	// order asc
	OrderAsc bool `json:"orderAsc,omitempty"`

	// order column
	OrderColumn EProxiesFiltersOrderColumn `json:"orderColumn,omitempty"`

	// skip
	Skip int32 `json:"skip,omitempty"`

	// type filter
	TypeFilter EProxyType `json:"typeFilter,omitempty"`
}

// Validate validates this proxies filters
func (m *ProxiesFilters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostIDFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderColumn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxiesFilters) validateHostIDFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.HostIDFilter) { // not required
		return nil
	}

	if err := validate.FormatOf("hostIdFilter", "body", "uuid", m.HostIDFilter.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesFilters) validateOrderColumn(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderColumn) { // not required
		return nil
	}

	if err := m.OrderColumn.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderColumn")
		}
		return err
	}

	return nil
}

func (m *ProxiesFilters) validateTypeFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeFilter) { // not required
		return nil
	}

	if err := m.TypeFilter.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeFilter")
		}
		return err
	}

	return nil
}

// ContextValidate validate this proxies filters based on the context it is used
func (m *ProxiesFilters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderColumn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxiesFilters) contextValidateOrderColumn(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderColumn.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderColumn")
		}
		return err
	}

	return nil
}

func (m *ProxiesFilters) contextValidateTypeFilter(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TypeFilter.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeFilter")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxiesFilters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxiesFilters) UnmarshalBinary(b []byte) error {
	var res ProxiesFilters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
