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

// GFSPolicySettingsMonthlyModel Monthly GFS retention policy.
type GFSPolicySettingsMonthlyModel struct {
	// If *true*, the monthly GFS retention policy is enabled.
	IsEnabled bool `json:"isEnabled"`
	// Number of months to keep full backups for archival purposes.
	KeepForNumberOfMonths *int32 `json:"keepForNumberOfMonths,omitempty"`
	DesiredTime *ESennightOfMonth `json:"desiredTime,omitempty"`
}

// NewGFSPolicySettingsMonthlyModel instantiates a new GFSPolicySettingsMonthlyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGFSPolicySettingsMonthlyModel(isEnabled bool, ) *GFSPolicySettingsMonthlyModel {
	this := GFSPolicySettingsMonthlyModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewGFSPolicySettingsMonthlyModelWithDefaults instantiates a new GFSPolicySettingsMonthlyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGFSPolicySettingsMonthlyModelWithDefaults() *GFSPolicySettingsMonthlyModel {
	this := GFSPolicySettingsMonthlyModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *GFSPolicySettingsMonthlyModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsMonthlyModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *GFSPolicySettingsMonthlyModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetKeepForNumberOfMonths returns the KeepForNumberOfMonths field value if set, zero value otherwise.
func (o *GFSPolicySettingsMonthlyModel) GetKeepForNumberOfMonths() int32 {
	if o == nil || o.KeepForNumberOfMonths == nil {
		var ret int32
		return ret
	}
	return *o.KeepForNumberOfMonths
}

// GetKeepForNumberOfMonthsOk returns a tuple with the KeepForNumberOfMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsMonthlyModel) GetKeepForNumberOfMonthsOk() (*int32, bool) {
	if o == nil || o.KeepForNumberOfMonths == nil {
		return nil, false
	}
	return o.KeepForNumberOfMonths, true
}

// HasKeepForNumberOfMonths returns a boolean if a field has been set.
func (o *GFSPolicySettingsMonthlyModel) HasKeepForNumberOfMonths() bool {
	if o != nil && o.KeepForNumberOfMonths != nil {
		return true
	}

	return false
}

// SetKeepForNumberOfMonths gets a reference to the given int32 and assigns it to the KeepForNumberOfMonths field.
func (o *GFSPolicySettingsMonthlyModel) SetKeepForNumberOfMonths(v int32) {
	o.KeepForNumberOfMonths = &v
}

// GetDesiredTime returns the DesiredTime field value if set, zero value otherwise.
func (o *GFSPolicySettingsMonthlyModel) GetDesiredTime() ESennightOfMonth {
	if o == nil || o.DesiredTime == nil {
		var ret ESennightOfMonth
		return ret
	}
	return *o.DesiredTime
}

// GetDesiredTimeOk returns a tuple with the DesiredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsMonthlyModel) GetDesiredTimeOk() (*ESennightOfMonth, bool) {
	if o == nil || o.DesiredTime == nil {
		return nil, false
	}
	return o.DesiredTime, true
}

// HasDesiredTime returns a boolean if a field has been set.
func (o *GFSPolicySettingsMonthlyModel) HasDesiredTime() bool {
	if o != nil && o.DesiredTime != nil {
		return true
	}

	return false
}

// SetDesiredTime gets a reference to the given ESennightOfMonth and assigns it to the DesiredTime field.
func (o *GFSPolicySettingsMonthlyModel) SetDesiredTime(v ESennightOfMonth) {
	o.DesiredTime = &v
}

func (o GFSPolicySettingsMonthlyModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.KeepForNumberOfMonths != nil {
		toSerialize["keepForNumberOfMonths"] = o.KeepForNumberOfMonths
	}
	if o.DesiredTime != nil {
		toSerialize["desiredTime"] = o.DesiredTime
	}
	return json.Marshal(toSerialize)
}

type NullableGFSPolicySettingsMonthlyModel struct {
	value *GFSPolicySettingsMonthlyModel
	isSet bool
}

func (v NullableGFSPolicySettingsMonthlyModel) Get() *GFSPolicySettingsMonthlyModel {
	return v.value
}

func (v *NullableGFSPolicySettingsMonthlyModel) Set(val *GFSPolicySettingsMonthlyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGFSPolicySettingsMonthlyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGFSPolicySettingsMonthlyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGFSPolicySettingsMonthlyModel(val *GFSPolicySettingsMonthlyModel) *NullableGFSPolicySettingsMonthlyModel {
	return &NullableGFSPolicySettingsMonthlyModel{value: val, isSet: true}
}

func (v NullableGFSPolicySettingsMonthlyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGFSPolicySettingsMonthlyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


