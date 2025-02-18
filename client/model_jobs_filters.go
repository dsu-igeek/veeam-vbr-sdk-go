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

// JobsFilters Filters jobs by the specified parameters.
type JobsFilters struct {
	// Skips the specified number of jobs.
	Skip *int32 `json:"skip,omitempty"`
	// Returns the specified number of jobs.
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EJobFiltersOrderColumn `json:"orderColumn,omitempty"`
	// Sorts jobs in the ascending order by the `orderColumn` parameter.
	OrderAsc *bool `json:"orderAsc,omitempty"`
	// Filters jobs by the `nameFilter` pattern. The pattern can match any job parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end.
	NameFilter *string `json:"nameFilter,omitempty"`
	TypeFilter *EJobType `json:"typeFilter,omitempty"`
}

// NewJobsFilters instantiates a new JobsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobsFilters() *JobsFilters {
	this := JobsFilters{}
	return &this
}

// NewJobsFiltersWithDefaults instantiates a new JobsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobsFiltersWithDefaults() *JobsFilters {
	this := JobsFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *JobsFilters) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobsFilters) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *JobsFilters) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *JobsFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *JobsFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *JobsFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *JobsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *JobsFilters) GetOrderColumn() EJobFiltersOrderColumn {
	if o == nil || o.OrderColumn == nil {
		var ret EJobFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobsFilters) GetOrderColumnOk() (*EJobFiltersOrderColumn, bool) {
	if o == nil || o.OrderColumn == nil {
		return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *JobsFilters) HasOrderColumn() bool {
	if o != nil && o.OrderColumn != nil {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EJobFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *JobsFilters) SetOrderColumn(v EJobFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *JobsFilters) GetOrderAsc() bool {
	if o == nil || o.OrderAsc == nil {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobsFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || o.OrderAsc == nil {
		return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *JobsFilters) HasOrderAsc() bool {
	if o != nil && o.OrderAsc != nil {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *JobsFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *JobsFilters) GetNameFilter() string {
	if o == nil || o.NameFilter == nil {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobsFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || o.NameFilter == nil {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *JobsFilters) HasNameFilter() bool {
	if o != nil && o.NameFilter != nil {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *JobsFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetTypeFilter returns the TypeFilter field value if set, zero value otherwise.
func (o *JobsFilters) GetTypeFilter() EJobType {
	if o == nil || o.TypeFilter == nil {
		var ret EJobType
		return ret
	}
	return *o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobsFilters) GetTypeFilterOk() (*EJobType, bool) {
	if o == nil || o.TypeFilter == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// HasTypeFilter returns a boolean if a field has been set.
func (o *JobsFilters) HasTypeFilter() bool {
	if o != nil && o.TypeFilter != nil {
		return true
	}

	return false
}

// SetTypeFilter gets a reference to the given EJobType and assigns it to the TypeFilter field.
func (o *JobsFilters) SetTypeFilter(v EJobType) {
	o.TypeFilter = &v
}

func (o JobsFilters) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableJobsFilters struct {
	value *JobsFilters
	isSet bool
}

func (v NullableJobsFilters) Get() *JobsFilters {
	return v.value
}

func (v *NullableJobsFilters) Set(val *JobsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableJobsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableJobsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobsFilters(val *JobsFilters) *NullableJobsFilters {
	return &NullableJobsFilters{value: val, isSet: true}
}

func (v NullableJobsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


