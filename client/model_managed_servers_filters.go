/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ManagedServersFilters struct for ManagedServersFilters
type ManagedServersFilters struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EManagedServersFiltersOrderColumn `json:"orderColumn,omitempty"`
	OrderAsc *bool `json:"orderAsc,omitempty"`
	NameFilter *string `json:"nameFilter,omitempty"`
	TypeFilter *EManagedServerType `json:"typeFilter,omitempty"`
	ViTypeFilter *EViHostType `json:"viTypeFilter,omitempty"`
}

// NewManagedServersFilters instantiates a new ManagedServersFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedServersFilters() *ManagedServersFilters {
	this := ManagedServersFilters{}
	return &this
}

// NewManagedServersFiltersWithDefaults instantiates a new ManagedServersFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedServersFiltersWithDefaults() *ManagedServersFilters {
	this := ManagedServersFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ManagedServersFilters) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServersFilters) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ManagedServersFilters) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ManagedServersFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ManagedServersFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServersFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ManagedServersFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ManagedServersFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *ManagedServersFilters) GetOrderColumn() EManagedServersFiltersOrderColumn {
	if o == nil || o.OrderColumn == nil {
		var ret EManagedServersFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServersFilters) GetOrderColumnOk() (*EManagedServersFiltersOrderColumn, bool) {
	if o == nil || o.OrderColumn == nil {
		return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *ManagedServersFilters) HasOrderColumn() bool {
	if o != nil && o.OrderColumn != nil {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EManagedServersFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *ManagedServersFilters) SetOrderColumn(v EManagedServersFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *ManagedServersFilters) GetOrderAsc() bool {
	if o == nil || o.OrderAsc == nil {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServersFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || o.OrderAsc == nil {
		return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *ManagedServersFilters) HasOrderAsc() bool {
	if o != nil && o.OrderAsc != nil {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *ManagedServersFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *ManagedServersFilters) GetNameFilter() string {
	if o == nil || o.NameFilter == nil {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServersFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || o.NameFilter == nil {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *ManagedServersFilters) HasNameFilter() bool {
	if o != nil && o.NameFilter != nil {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *ManagedServersFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetTypeFilter returns the TypeFilter field value if set, zero value otherwise.
func (o *ManagedServersFilters) GetTypeFilter() EManagedServerType {
	if o == nil || o.TypeFilter == nil {
		var ret EManagedServerType
		return ret
	}
	return *o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServersFilters) GetTypeFilterOk() (*EManagedServerType, bool) {
	if o == nil || o.TypeFilter == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// HasTypeFilter returns a boolean if a field has been set.
func (o *ManagedServersFilters) HasTypeFilter() bool {
	if o != nil && o.TypeFilter != nil {
		return true
	}

	return false
}

// SetTypeFilter gets a reference to the given EManagedServerType and assigns it to the TypeFilter field.
func (o *ManagedServersFilters) SetTypeFilter(v EManagedServerType) {
	o.TypeFilter = &v
}

// GetViTypeFilter returns the ViTypeFilter field value if set, zero value otherwise.
func (o *ManagedServersFilters) GetViTypeFilter() EViHostType {
	if o == nil || o.ViTypeFilter == nil {
		var ret EViHostType
		return ret
	}
	return *o.ViTypeFilter
}

// GetViTypeFilterOk returns a tuple with the ViTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedServersFilters) GetViTypeFilterOk() (*EViHostType, bool) {
	if o == nil || o.ViTypeFilter == nil {
		return nil, false
	}
	return o.ViTypeFilter, true
}

// HasViTypeFilter returns a boolean if a field has been set.
func (o *ManagedServersFilters) HasViTypeFilter() bool {
	if o != nil && o.ViTypeFilter != nil {
		return true
	}

	return false
}

// SetViTypeFilter gets a reference to the given EViHostType and assigns it to the ViTypeFilter field.
func (o *ManagedServersFilters) SetViTypeFilter(v EViHostType) {
	o.ViTypeFilter = &v
}

func (o ManagedServersFilters) MarshalJSON() ([]byte, error) {
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
	if o.TypeFilter != nil {
		toSerialize["typeFilter"] = o.TypeFilter
	}
	if o.ViTypeFilter != nil {
		toSerialize["viTypeFilter"] = o.ViTypeFilter
	}
	return json.Marshal(toSerialize)
}

type NullableManagedServersFilters struct {
	value *ManagedServersFilters
	isSet bool
}

func (v NullableManagedServersFilters) Get() *ManagedServersFilters {
	return v.value
}

func (v *NullableManagedServersFilters) Set(val *ManagedServersFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedServersFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedServersFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedServersFilters(val *ManagedServersFilters) *NullableManagedServersFilters {
	return &NullableManagedServersFilters{value: val, isSet: true}
}

func (v NullableManagedServersFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedServersFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

