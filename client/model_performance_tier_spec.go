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

// PerformanceTierSpec Performance tier.
type PerformanceTierSpec struct {
	// Array of performance extents.
	PerformanceExtents []PerformanceExtentSpec `json:"performanceExtents"`
	AdvancedSettings *PerformanceTierAdvancedSettingsModel `json:"advancedSettings,omitempty"`
}

// NewPerformanceTierSpec instantiates a new PerformanceTierSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceTierSpec(performanceExtents []PerformanceExtentSpec, ) *PerformanceTierSpec {
	this := PerformanceTierSpec{}
	this.PerformanceExtents = performanceExtents
	return &this
}

// NewPerformanceTierSpecWithDefaults instantiates a new PerformanceTierSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceTierSpecWithDefaults() *PerformanceTierSpec {
	this := PerformanceTierSpec{}
	return &this
}

// GetPerformanceExtents returns the PerformanceExtents field value
func (o *PerformanceTierSpec) GetPerformanceExtents() []PerformanceExtentSpec {
	if o == nil  {
		var ret []PerformanceExtentSpec
		return ret
	}

	return o.PerformanceExtents
}

// GetPerformanceExtentsOk returns a tuple with the PerformanceExtents field value
// and a boolean to check if the value has been set.
func (o *PerformanceTierSpec) GetPerformanceExtentsOk() (*[]PerformanceExtentSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PerformanceExtents, true
}

// SetPerformanceExtents sets field value
func (o *PerformanceTierSpec) SetPerformanceExtents(v []PerformanceExtentSpec) {
	o.PerformanceExtents = v
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *PerformanceTierSpec) GetAdvancedSettings() PerformanceTierAdvancedSettingsModel {
	if o == nil || o.AdvancedSettings == nil {
		var ret PerformanceTierAdvancedSettingsModel
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceTierSpec) GetAdvancedSettingsOk() (*PerformanceTierAdvancedSettingsModel, bool) {
	if o == nil || o.AdvancedSettings == nil {
		return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *PerformanceTierSpec) HasAdvancedSettings() bool {
	if o != nil && o.AdvancedSettings != nil {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given PerformanceTierAdvancedSettingsModel and assigns it to the AdvancedSettings field.
func (o *PerformanceTierSpec) SetAdvancedSettings(v PerformanceTierAdvancedSettingsModel) {
	o.AdvancedSettings = &v
}

func (o PerformanceTierSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["performanceExtents"] = o.PerformanceExtents
	}
	if o.AdvancedSettings != nil {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	return json.Marshal(toSerialize)
}

type NullablePerformanceTierSpec struct {
	value *PerformanceTierSpec
	isSet bool
}

func (v NullablePerformanceTierSpec) Get() *PerformanceTierSpec {
	return v.value
}

func (v *NullablePerformanceTierSpec) Set(val *PerformanceTierSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceTierSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceTierSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceTierSpec(val *PerformanceTierSpec) *NullablePerformanceTierSpec {
	return &NullablePerformanceTierSpec{value: val, isSet: true}
}

func (v NullablePerformanceTierSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceTierSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


