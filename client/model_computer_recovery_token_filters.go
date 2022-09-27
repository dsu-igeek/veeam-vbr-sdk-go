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
	"time"
)

// ComputerRecoveryTokenFilters struct for ComputerRecoveryTokenFilters
type ComputerRecoveryTokenFilters struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EComputerRecoveryTokenFiltersOrderColumn `json:"orderColumn,omitempty"`
	OrderAsc *bool `json:"orderAsc,omitempty"`
	NameFilter *string `json:"nameFilter,omitempty"`
	ExpirationDateBefore *time.Time `json:"expirationDateBefore,omitempty"`
	ExpirationDateAfter *time.Time `json:"expirationDateAfter,omitempty"`
}

// NewComputerRecoveryTokenFilters instantiates a new ComputerRecoveryTokenFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerRecoveryTokenFilters() *ComputerRecoveryTokenFilters {
	this := ComputerRecoveryTokenFilters{}
	return &this
}

// NewComputerRecoveryTokenFiltersWithDefaults instantiates a new ComputerRecoveryTokenFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerRecoveryTokenFiltersWithDefaults() *ComputerRecoveryTokenFilters {
	this := ComputerRecoveryTokenFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ComputerRecoveryTokenFilters) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenFilters) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ComputerRecoveryTokenFilters) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ComputerRecoveryTokenFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ComputerRecoveryTokenFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ComputerRecoveryTokenFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ComputerRecoveryTokenFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *ComputerRecoveryTokenFilters) GetOrderColumn() EComputerRecoveryTokenFiltersOrderColumn {
	if o == nil || o.OrderColumn == nil {
		var ret EComputerRecoveryTokenFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenFilters) GetOrderColumnOk() (*EComputerRecoveryTokenFiltersOrderColumn, bool) {
	if o == nil || o.OrderColumn == nil {
		return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *ComputerRecoveryTokenFilters) HasOrderColumn() bool {
	if o != nil && o.OrderColumn != nil {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EComputerRecoveryTokenFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *ComputerRecoveryTokenFilters) SetOrderColumn(v EComputerRecoveryTokenFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *ComputerRecoveryTokenFilters) GetOrderAsc() bool {
	if o == nil || o.OrderAsc == nil {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || o.OrderAsc == nil {
		return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *ComputerRecoveryTokenFilters) HasOrderAsc() bool {
	if o != nil && o.OrderAsc != nil {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *ComputerRecoveryTokenFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *ComputerRecoveryTokenFilters) GetNameFilter() string {
	if o == nil || o.NameFilter == nil {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || o.NameFilter == nil {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *ComputerRecoveryTokenFilters) HasNameFilter() bool {
	if o != nil && o.NameFilter != nil {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *ComputerRecoveryTokenFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetExpirationDateBefore returns the ExpirationDateBefore field value if set, zero value otherwise.
func (o *ComputerRecoveryTokenFilters) GetExpirationDateBefore() time.Time {
	if o == nil || o.ExpirationDateBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateBefore
}

// GetExpirationDateBeforeOk returns a tuple with the ExpirationDateBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenFilters) GetExpirationDateBeforeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDateBefore == nil {
		return nil, false
	}
	return o.ExpirationDateBefore, true
}

// HasExpirationDateBefore returns a boolean if a field has been set.
func (o *ComputerRecoveryTokenFilters) HasExpirationDateBefore() bool {
	if o != nil && o.ExpirationDateBefore != nil {
		return true
	}

	return false
}

// SetExpirationDateBefore gets a reference to the given time.Time and assigns it to the ExpirationDateBefore field.
func (o *ComputerRecoveryTokenFilters) SetExpirationDateBefore(v time.Time) {
	o.ExpirationDateBefore = &v
}

// GetExpirationDateAfter returns the ExpirationDateAfter field value if set, zero value otherwise.
func (o *ComputerRecoveryTokenFilters) GetExpirationDateAfter() time.Time {
	if o == nil || o.ExpirationDateAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateAfter
}

// GetExpirationDateAfterOk returns a tuple with the ExpirationDateAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenFilters) GetExpirationDateAfterOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDateAfter == nil {
		return nil, false
	}
	return o.ExpirationDateAfter, true
}

// HasExpirationDateAfter returns a boolean if a field has been set.
func (o *ComputerRecoveryTokenFilters) HasExpirationDateAfter() bool {
	if o != nil && o.ExpirationDateAfter != nil {
		return true
	}

	return false
}

// SetExpirationDateAfter gets a reference to the given time.Time and assigns it to the ExpirationDateAfter field.
func (o *ComputerRecoveryTokenFilters) SetExpirationDateAfter(v time.Time) {
	o.ExpirationDateAfter = &v
}

func (o ComputerRecoveryTokenFilters) MarshalJSON() ([]byte, error) {
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
	if o.ExpirationDateBefore != nil {
		toSerialize["expirationDateBefore"] = o.ExpirationDateBefore
	}
	if o.ExpirationDateAfter != nil {
		toSerialize["expirationDateAfter"] = o.ExpirationDateAfter
	}
	return json.Marshal(toSerialize)
}

type NullableComputerRecoveryTokenFilters struct {
	value *ComputerRecoveryTokenFilters
	isSet bool
}

func (v NullableComputerRecoveryTokenFilters) Get() *ComputerRecoveryTokenFilters {
	return v.value
}

func (v *NullableComputerRecoveryTokenFilters) Set(val *ComputerRecoveryTokenFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerRecoveryTokenFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerRecoveryTokenFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerRecoveryTokenFilters(val *ComputerRecoveryTokenFilters) *NullableComputerRecoveryTokenFilters {
	return &NullableComputerRecoveryTokenFilters{value: val, isSet: true}
}

func (v NullableComputerRecoveryTokenFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerRecoveryTokenFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


