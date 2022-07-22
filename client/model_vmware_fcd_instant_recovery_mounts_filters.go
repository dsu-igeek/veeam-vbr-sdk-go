/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 

API version: 1.0-rev2
Contact: support@veeam.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// VmwareFcdInstantRecoveryMountsFilters struct for VmwareFcdInstantRecoveryMountsFilters
type VmwareFcdInstantRecoveryMountsFilters struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EVmwareFcdInstantRecoveryMountsFiltersOrderColumn `json:"orderColumn,omitempty"`
	OrderAsc *bool `json:"orderAsc,omitempty"`
	NameFilter *string `json:"nameFilter,omitempty"`
	StateFilter *EInstantRecoveryMountState `json:"stateFilter,omitempty"`
}

// NewVmwareFcdInstantRecoveryMountsFilters instantiates a new VmwareFcdInstantRecoveryMountsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareFcdInstantRecoveryMountsFilters() *VmwareFcdInstantRecoveryMountsFilters {
	this := VmwareFcdInstantRecoveryMountsFilters{}
	return &this
}

// NewVmwareFcdInstantRecoveryMountsFiltersWithDefaults instantiates a new VmwareFcdInstantRecoveryMountsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareFcdInstantRecoveryMountsFiltersWithDefaults() *VmwareFcdInstantRecoveryMountsFilters {
	this := VmwareFcdInstantRecoveryMountsFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *VmwareFcdInstantRecoveryMountsFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *VmwareFcdInstantRecoveryMountsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetOrderColumn() EVmwareFcdInstantRecoveryMountsFiltersOrderColumn {
	if o == nil || o.OrderColumn == nil {
		var ret EVmwareFcdInstantRecoveryMountsFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetOrderColumnOk() (*EVmwareFcdInstantRecoveryMountsFiltersOrderColumn, bool) {
	if o == nil || o.OrderColumn == nil {
		return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) HasOrderColumn() bool {
	if o != nil && o.OrderColumn != nil {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EVmwareFcdInstantRecoveryMountsFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *VmwareFcdInstantRecoveryMountsFilters) SetOrderColumn(v EVmwareFcdInstantRecoveryMountsFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetOrderAsc() bool {
	if o == nil || o.OrderAsc == nil {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || o.OrderAsc == nil {
		return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) HasOrderAsc() bool {
	if o != nil && o.OrderAsc != nil {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *VmwareFcdInstantRecoveryMountsFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetNameFilter() string {
	if o == nil || o.NameFilter == nil {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || o.NameFilter == nil {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) HasNameFilter() bool {
	if o != nil && o.NameFilter != nil {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *VmwareFcdInstantRecoveryMountsFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetStateFilter() EInstantRecoveryMountState {
	if o == nil || o.StateFilter == nil {
		var ret EInstantRecoveryMountState
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) GetStateFilterOk() (*EInstantRecoveryMountState, bool) {
	if o == nil || o.StateFilter == nil {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoveryMountsFilters) HasStateFilter() bool {
	if o != nil && o.StateFilter != nil {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given EInstantRecoveryMountState and assigns it to the StateFilter field.
func (o *VmwareFcdInstantRecoveryMountsFilters) SetStateFilter(v EInstantRecoveryMountState) {
	o.StateFilter = &v
}

func (o VmwareFcdInstantRecoveryMountsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.OrderColumn != nil {
		toSerialize["orderColumn"] = o.OrderColumn
	}
	if o.OrderAsc != nil {
		toSerialize["orderAsc"] = o.OrderAsc
	}
	if o.NameFilter != nil {
		toSerialize["nameFilter"] = o.NameFilter
	}
	if o.StateFilter != nil {
		toSerialize["stateFilter"] = o.StateFilter
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareFcdInstantRecoveryMountsFilters struct {
	value *VmwareFcdInstantRecoveryMountsFilters
	isSet bool
}

func (v NullableVmwareFcdInstantRecoveryMountsFilters) Get() *VmwareFcdInstantRecoveryMountsFilters {
	return v.value
}

func (v *NullableVmwareFcdInstantRecoveryMountsFilters) Set(val *VmwareFcdInstantRecoveryMountsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareFcdInstantRecoveryMountsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareFcdInstantRecoveryMountsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareFcdInstantRecoveryMountsFilters(val *VmwareFcdInstantRecoveryMountsFilters) *NullableVmwareFcdInstantRecoveryMountsFilters {
	return &NullableVmwareFcdInstantRecoveryMountsFilters{value: val, isSet: true}
}

func (v NullableVmwareFcdInstantRecoveryMountsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareFcdInstantRecoveryMountsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


