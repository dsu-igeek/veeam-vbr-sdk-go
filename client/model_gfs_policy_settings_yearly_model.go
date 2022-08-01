/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev2
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GFSPolicySettingsYearlyModel Yearly GFS retention policy.
type GFSPolicySettingsYearlyModel struct {
	// If *true*, the yearly GFS retention policy is enabled.
	IsEnabled bool `json:"isEnabled"`
	// Number of years to keep full backups for archival purposes.
	KeepForNumberOfYears *int32 `json:"keepForNumberOfYears,omitempty"`
	DesiredTime *EMonth `json:"desiredTime,omitempty"`
}

// NewGFSPolicySettingsYearlyModel instantiates a new GFSPolicySettingsYearlyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGFSPolicySettingsYearlyModel(isEnabled bool, ) *GFSPolicySettingsYearlyModel {
	this := GFSPolicySettingsYearlyModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewGFSPolicySettingsYearlyModelWithDefaults instantiates a new GFSPolicySettingsYearlyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGFSPolicySettingsYearlyModelWithDefaults() *GFSPolicySettingsYearlyModel {
	this := GFSPolicySettingsYearlyModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *GFSPolicySettingsYearlyModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsYearlyModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *GFSPolicySettingsYearlyModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetKeepForNumberOfYears returns the KeepForNumberOfYears field value if set, zero value otherwise.
func (o *GFSPolicySettingsYearlyModel) GetKeepForNumberOfYears() int32 {
	if o == nil || o.KeepForNumberOfYears == nil {
		var ret int32
		return ret
	}
	return *o.KeepForNumberOfYears
}

// GetKeepForNumberOfYearsOk returns a tuple with the KeepForNumberOfYears field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsYearlyModel) GetKeepForNumberOfYearsOk() (*int32, bool) {
	if o == nil || o.KeepForNumberOfYears == nil {
		return nil, false
	}
	return o.KeepForNumberOfYears, true
}

// HasKeepForNumberOfYears returns a boolean if a field has been set.
func (o *GFSPolicySettingsYearlyModel) HasKeepForNumberOfYears() bool {
	if o != nil && o.KeepForNumberOfYears != nil {
		return true
	}

	return false
}

// SetKeepForNumberOfYears gets a reference to the given int32 and assigns it to the KeepForNumberOfYears field.
func (o *GFSPolicySettingsYearlyModel) SetKeepForNumberOfYears(v int32) {
	o.KeepForNumberOfYears = &v
}

// GetDesiredTime returns the DesiredTime field value if set, zero value otherwise.
func (o *GFSPolicySettingsYearlyModel) GetDesiredTime() EMonth {
	if o == nil || o.DesiredTime == nil {
		var ret EMonth
		return ret
	}
	return *o.DesiredTime
}

// GetDesiredTimeOk returns a tuple with the DesiredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsYearlyModel) GetDesiredTimeOk() (*EMonth, bool) {
	if o == nil || o.DesiredTime == nil {
		return nil, false
	}
	return o.DesiredTime, true
}

// HasDesiredTime returns a boolean if a field has been set.
func (o *GFSPolicySettingsYearlyModel) HasDesiredTime() bool {
	if o != nil && o.DesiredTime != nil {
		return true
	}

	return false
}

// SetDesiredTime gets a reference to the given EMonth and assigns it to the DesiredTime field.
func (o *GFSPolicySettingsYearlyModel) SetDesiredTime(v EMonth) {
	o.DesiredTime = &v
}

func (o GFSPolicySettingsYearlyModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.KeepForNumberOfYears != nil {
		toSerialize["keepForNumberOfYears"] = o.KeepForNumberOfYears
	}
	if o.DesiredTime != nil {
		toSerialize["desiredTime"] = o.DesiredTime
	}
	return json.Marshal(toSerialize)
}

type NullableGFSPolicySettingsYearlyModel struct {
	value *GFSPolicySettingsYearlyModel
	isSet bool
}

func (v NullableGFSPolicySettingsYearlyModel) Get() *GFSPolicySettingsYearlyModel {
	return v.value
}

func (v *NullableGFSPolicySettingsYearlyModel) Set(val *GFSPolicySettingsYearlyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGFSPolicySettingsYearlyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGFSPolicySettingsYearlyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGFSPolicySettingsYearlyModel(val *GFSPolicySettingsYearlyModel) *NullableGFSPolicySettingsYearlyModel {
	return &NullableGFSPolicySettingsYearlyModel{value: val, isSet: true}
}

func (v NullableGFSPolicySettingsYearlyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGFSPolicySettingsYearlyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


