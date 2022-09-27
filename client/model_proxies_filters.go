/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br>Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic.
 *
 * API version: 1.1-rev0
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProxiesFilters struct for ProxiesFilters
type ProxiesFilters struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EProxiesFiltersOrderColumn `json:"orderColumn,omitempty"`
	OrderAsc *bool `json:"orderAsc,omitempty"`
	NameFilter *string `json:"nameFilter,omitempty"`
	TypeFilter *EProxyType `json:"typeFilter,omitempty"`
	HostIdFilter *string `json:"hostIdFilter,omitempty"`
}

// NewProxiesFilters instantiates a new ProxiesFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxiesFilters() *ProxiesFilters {
	this := ProxiesFilters{}
	return &this
}

// NewProxiesFiltersWithDefaults instantiates a new ProxiesFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxiesFiltersWithDefaults() *ProxiesFilters {
	this := ProxiesFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ProxiesFilters) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxiesFilters) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ProxiesFilters) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ProxiesFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProxiesFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxiesFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProxiesFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ProxiesFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *ProxiesFilters) GetOrderColumn() EProxiesFiltersOrderColumn {
	if o == nil || o.OrderColumn == nil {
		var ret EProxiesFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxiesFilters) GetOrderColumnOk() (*EProxiesFiltersOrderColumn, bool) {
	if o == nil || o.OrderColumn == nil {
		return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *ProxiesFilters) HasOrderColumn() bool {
	if o != nil && o.OrderColumn != nil {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EProxiesFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *ProxiesFilters) SetOrderColumn(v EProxiesFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *ProxiesFilters) GetOrderAsc() bool {
	if o == nil || o.OrderAsc == nil {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxiesFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || o.OrderAsc == nil {
		return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *ProxiesFilters) HasOrderAsc() bool {
	if o != nil && o.OrderAsc != nil {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *ProxiesFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *ProxiesFilters) GetNameFilter() string {
	if o == nil || o.NameFilter == nil {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxiesFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || o.NameFilter == nil {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *ProxiesFilters) HasNameFilter() bool {
	if o != nil && o.NameFilter != nil {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *ProxiesFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetTypeFilter returns the TypeFilter field value if set, zero value otherwise.
func (o *ProxiesFilters) GetTypeFilter() EProxyType {
	if o == nil || o.TypeFilter == nil {
		var ret EProxyType
		return ret
	}
	return *o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxiesFilters) GetTypeFilterOk() (*EProxyType, bool) {
	if o == nil || o.TypeFilter == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// HasTypeFilter returns a boolean if a field has been set.
func (o *ProxiesFilters) HasTypeFilter() bool {
	if o != nil && o.TypeFilter != nil {
		return true
	}

	return false
}

// SetTypeFilter gets a reference to the given EProxyType and assigns it to the TypeFilter field.
func (o *ProxiesFilters) SetTypeFilter(v EProxyType) {
	o.TypeFilter = &v
}

// GetHostIdFilter returns the HostIdFilter field value if set, zero value otherwise.
func (o *ProxiesFilters) GetHostIdFilter() string {
	if o == nil || o.HostIdFilter == nil {
		var ret string
		return ret
	}
	return *o.HostIdFilter
}

// GetHostIdFilterOk returns a tuple with the HostIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxiesFilters) GetHostIdFilterOk() (*string, bool) {
	if o == nil || o.HostIdFilter == nil {
		return nil, false
	}
	return o.HostIdFilter, true
}

// HasHostIdFilter returns a boolean if a field has been set.
func (o *ProxiesFilters) HasHostIdFilter() bool {
	if o != nil && o.HostIdFilter != nil {
		return true
	}

	return false
}

// SetHostIdFilter gets a reference to the given string and assigns it to the HostIdFilter field.
func (o *ProxiesFilters) SetHostIdFilter(v string) {
	o.HostIdFilter = &v
}

func (o ProxiesFilters) MarshalJSON() ([]byte, error) {
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
	if o.HostIdFilter != nil {
		toSerialize["hostIdFilter"] = o.HostIdFilter
	}
	return json.Marshal(toSerialize)
}

type NullableProxiesFilters struct {
	value *ProxiesFilters
	isSet bool
}

func (v NullableProxiesFilters) Get() *ProxiesFilters {
	return v.value
}

func (v *NullableProxiesFilters) Set(val *ProxiesFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableProxiesFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableProxiesFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxiesFilters(val *ProxiesFilters) *NullableProxiesFilters {
	return &NullableProxiesFilters{value: val, isSet: true}
}

func (v NullableProxiesFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxiesFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


