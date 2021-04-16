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

// VCenterInventoryFilters v center inventory filters
//
// swagger:model VCenterInventoryFilters
type VCenterInventoryFilters struct {

	// hierarchy type filter
	HierarchyTypeFilter EHierarchyType `json:"hierarchyTypeFilter,omitempty"`

	// Maximum number of objects to return.
	Limit int32 `json:"limit,omitempty"`

	// Filters objects by the `nameFilter` pattern. The pattern can match any object parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end.
	NameFilter string `json:"nameFilter,omitempty"`

	// Filters objects by object ID.
	ObjectIDFilter string `json:"objectIdFilter,omitempty"`

	// Sorts objects in the ascending order by the `orderColumn` parameter.
	OrderAsc bool `json:"orderAsc,omitempty"`

	// order column
	OrderColumn EvCentersInventoryFiltersOrderColumn `json:"orderColumn,omitempty"`

	// Filters objects by name of the parent container.
	ParentContainerNameFilter string `json:"parentContainerNameFilter,omitempty"`

	// Number of objects to skip.
	Skip int32 `json:"skip,omitempty"`

	// type filter
	TypeFilter EVmwareInventoryType `json:"typeFilter,omitempty"`
}

// Validate validates this v center inventory filters
func (m *VCenterInventoryFilters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHierarchyTypeFilter(formats); err != nil {
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

func (m *VCenterInventoryFilters) validateHierarchyTypeFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.HierarchyTypeFilter) { // not required
		return nil
	}

	if err := m.HierarchyTypeFilter.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hierarchyTypeFilter")
		}
		return err
	}

	return nil
}

func (m *VCenterInventoryFilters) validateOrderColumn(formats strfmt.Registry) error {
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

func (m *VCenterInventoryFilters) validateTypeFilter(formats strfmt.Registry) error {
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

// ContextValidate validate this v center inventory filters based on the context it is used
func (m *VCenterInventoryFilters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHierarchyTypeFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *VCenterInventoryFilters) contextValidateHierarchyTypeFilter(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HierarchyTypeFilter.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hierarchyTypeFilter")
		}
		return err
	}

	return nil
}

func (m *VCenterInventoryFilters) contextValidateOrderColumn(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderColumn.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderColumn")
		}
		return err
	}

	return nil
}

func (m *VCenterInventoryFilters) contextValidateTypeFilter(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TypeFilter.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeFilter")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VCenterInventoryFilters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VCenterInventoryFilters) UnmarshalBinary(b []byte) error {
	var res VCenterInventoryFilters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
