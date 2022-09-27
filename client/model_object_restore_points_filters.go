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

// ObjectRestorePointsFilters struct for ObjectRestorePointsFilters
type ObjectRestorePointsFilters struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EObjectRestorePointsFiltersOrderColumn `json:"orderColumn,omitempty"`
	OrderAsc *bool `json:"orderAsc,omitempty"`
	CreatedAfterFilter *time.Time `json:"createdAfterFilter,omitempty"`
	CreatedBeforeFilter *time.Time `json:"createdBeforeFilter,omitempty"`
	NameFilter *string `json:"nameFilter,omitempty"`
	PlatformNameFilter *EPlatformType `json:"platformNameFilter,omitempty"`
	PlatformIdFilter *string `json:"platformIdFilter,omitempty"`
	BackupIdFilter *string `json:"backupIdFilter,omitempty"`
	BackupObjectIdFilter *string `json:"backupObjectIdFilter,omitempty"`
}

// NewObjectRestorePointsFilters instantiates a new ObjectRestorePointsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectRestorePointsFilters() *ObjectRestorePointsFilters {
	this := ObjectRestorePointsFilters{}
	return &this
}

// NewObjectRestorePointsFiltersWithDefaults instantiates a new ObjectRestorePointsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectRestorePointsFiltersWithDefaults() *ObjectRestorePointsFilters {
	this := ObjectRestorePointsFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ObjectRestorePointsFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ObjectRestorePointsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetOrderColumn() EObjectRestorePointsFiltersOrderColumn {
	if o == nil || o.OrderColumn == nil {
		var ret EObjectRestorePointsFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetOrderColumnOk() (*EObjectRestorePointsFiltersOrderColumn, bool) {
	if o == nil || o.OrderColumn == nil {
		return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasOrderColumn() bool {
	if o != nil && o.OrderColumn != nil {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EObjectRestorePointsFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *ObjectRestorePointsFilters) SetOrderColumn(v EObjectRestorePointsFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetOrderAsc() bool {
	if o == nil || o.OrderAsc == nil {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || o.OrderAsc == nil {
		return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasOrderAsc() bool {
	if o != nil && o.OrderAsc != nil {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *ObjectRestorePointsFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetCreatedAfterFilter returns the CreatedAfterFilter field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetCreatedAfterFilter() time.Time {
	if o == nil || o.CreatedAfterFilter == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAfterFilter
}

// GetCreatedAfterFilterOk returns a tuple with the CreatedAfterFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetCreatedAfterFilterOk() (*time.Time, bool) {
	if o == nil || o.CreatedAfterFilter == nil {
		return nil, false
	}
	return o.CreatedAfterFilter, true
}

// HasCreatedAfterFilter returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasCreatedAfterFilter() bool {
	if o != nil && o.CreatedAfterFilter != nil {
		return true
	}

	return false
}

// SetCreatedAfterFilter gets a reference to the given time.Time and assigns it to the CreatedAfterFilter field.
func (o *ObjectRestorePointsFilters) SetCreatedAfterFilter(v time.Time) {
	o.CreatedAfterFilter = &v
}

// GetCreatedBeforeFilter returns the CreatedBeforeFilter field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetCreatedBeforeFilter() time.Time {
	if o == nil || o.CreatedBeforeFilter == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedBeforeFilter
}

// GetCreatedBeforeFilterOk returns a tuple with the CreatedBeforeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetCreatedBeforeFilterOk() (*time.Time, bool) {
	if o == nil || o.CreatedBeforeFilter == nil {
		return nil, false
	}
	return o.CreatedBeforeFilter, true
}

// HasCreatedBeforeFilter returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasCreatedBeforeFilter() bool {
	if o != nil && o.CreatedBeforeFilter != nil {
		return true
	}

	return false
}

// SetCreatedBeforeFilter gets a reference to the given time.Time and assigns it to the CreatedBeforeFilter field.
func (o *ObjectRestorePointsFilters) SetCreatedBeforeFilter(v time.Time) {
	o.CreatedBeforeFilter = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetNameFilter() string {
	if o == nil || o.NameFilter == nil {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || o.NameFilter == nil {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasNameFilter() bool {
	if o != nil && o.NameFilter != nil {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *ObjectRestorePointsFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetPlatformNameFilter returns the PlatformNameFilter field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetPlatformNameFilter() EPlatformType {
	if o == nil || o.PlatformNameFilter == nil {
		var ret EPlatformType
		return ret
	}
	return *o.PlatformNameFilter
}

// GetPlatformNameFilterOk returns a tuple with the PlatformNameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetPlatformNameFilterOk() (*EPlatformType, bool) {
	if o == nil || o.PlatformNameFilter == nil {
		return nil, false
	}
	return o.PlatformNameFilter, true
}

// HasPlatformNameFilter returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasPlatformNameFilter() bool {
	if o != nil && o.PlatformNameFilter != nil {
		return true
	}

	return false
}

// SetPlatformNameFilter gets a reference to the given EPlatformType and assigns it to the PlatformNameFilter field.
func (o *ObjectRestorePointsFilters) SetPlatformNameFilter(v EPlatformType) {
	o.PlatformNameFilter = &v
}

// GetPlatformIdFilter returns the PlatformIdFilter field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetPlatformIdFilter() string {
	if o == nil || o.PlatformIdFilter == nil {
		var ret string
		return ret
	}
	return *o.PlatformIdFilter
}

// GetPlatformIdFilterOk returns a tuple with the PlatformIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetPlatformIdFilterOk() (*string, bool) {
	if o == nil || o.PlatformIdFilter == nil {
		return nil, false
	}
	return o.PlatformIdFilter, true
}

// HasPlatformIdFilter returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasPlatformIdFilter() bool {
	if o != nil && o.PlatformIdFilter != nil {
		return true
	}

	return false
}

// SetPlatformIdFilter gets a reference to the given string and assigns it to the PlatformIdFilter field.
func (o *ObjectRestorePointsFilters) SetPlatformIdFilter(v string) {
	o.PlatformIdFilter = &v
}

// GetBackupIdFilter returns the BackupIdFilter field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetBackupIdFilter() string {
	if o == nil || o.BackupIdFilter == nil {
		var ret string
		return ret
	}
	return *o.BackupIdFilter
}

// GetBackupIdFilterOk returns a tuple with the BackupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetBackupIdFilterOk() (*string, bool) {
	if o == nil || o.BackupIdFilter == nil {
		return nil, false
	}
	return o.BackupIdFilter, true
}

// HasBackupIdFilter returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasBackupIdFilter() bool {
	if o != nil && o.BackupIdFilter != nil {
		return true
	}

	return false
}

// SetBackupIdFilter gets a reference to the given string and assigns it to the BackupIdFilter field.
func (o *ObjectRestorePointsFilters) SetBackupIdFilter(v string) {
	o.BackupIdFilter = &v
}

// GetBackupObjectIdFilter returns the BackupObjectIdFilter field value if set, zero value otherwise.
func (o *ObjectRestorePointsFilters) GetBackupObjectIdFilter() string {
	if o == nil || o.BackupObjectIdFilter == nil {
		var ret string
		return ret
	}
	return *o.BackupObjectIdFilter
}

// GetBackupObjectIdFilterOk returns a tuple with the BackupObjectIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsFilters) GetBackupObjectIdFilterOk() (*string, bool) {
	if o == nil || o.BackupObjectIdFilter == nil {
		return nil, false
	}
	return o.BackupObjectIdFilter, true
}

// HasBackupObjectIdFilter returns a boolean if a field has been set.
func (o *ObjectRestorePointsFilters) HasBackupObjectIdFilter() bool {
	if o != nil && o.BackupObjectIdFilter != nil {
		return true
	}

	return false
}

// SetBackupObjectIdFilter gets a reference to the given string and assigns it to the BackupObjectIdFilter field.
func (o *ObjectRestorePointsFilters) SetBackupObjectIdFilter(v string) {
	o.BackupObjectIdFilter = &v
}

func (o ObjectRestorePointsFilters) MarshalJSON() ([]byte, error) {
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
	if o.CreatedAfterFilter != nil {
		toSerialize["createdAfterFilter"] = o.CreatedAfterFilter
	}
	if o.CreatedBeforeFilter != nil {
		toSerialize["createdBeforeFilter"] = o.CreatedBeforeFilter
	}
	if o.NameFilter != nil {
		toSerialize["nameFilter"] = o.NameFilter
	}
	if o.PlatformNameFilter != nil {
		toSerialize["platformNameFilter"] = o.PlatformNameFilter
	}
	if o.PlatformIdFilter != nil {
		toSerialize["platformIdFilter"] = o.PlatformIdFilter
	}
	if o.BackupIdFilter != nil {
		toSerialize["backupIdFilter"] = o.BackupIdFilter
	}
	if o.BackupObjectIdFilter != nil {
		toSerialize["backupObjectIdFilter"] = o.BackupObjectIdFilter
	}
	return json.Marshal(toSerialize)
}

type NullableObjectRestorePointsFilters struct {
	value *ObjectRestorePointsFilters
	isSet bool
}

func (v NullableObjectRestorePointsFilters) Get() *ObjectRestorePointsFilters {
	return v.value
}

func (v *NullableObjectRestorePointsFilters) Set(val *ObjectRestorePointsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectRestorePointsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectRestorePointsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectRestorePointsFilters(val *ObjectRestorePointsFilters) *NullableObjectRestorePointsFilters {
	return &NullableObjectRestorePointsFilters{value: val, isSet: true}
}

func (v NullableObjectRestorePointsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectRestorePointsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


