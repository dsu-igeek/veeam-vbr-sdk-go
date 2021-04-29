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

// GFSPolicySettingsModel GFS retention policy settings.
type GFSPolicySettingsModel struct {
	// If *true*, the long-term (GFS) retention policy is enabled.
	IsEnabled bool `json:"isEnabled"`
	Weekly *GFSPolicySettingsWeeklyModel `json:"weekly,omitempty"`
	Monthly *GFSPolicySettingsMonthlyModel `json:"monthly,omitempty"`
	Yearly *GFSPolicySettingsYearlyModel `json:"yearly,omitempty"`
}

// NewGFSPolicySettingsModel instantiates a new GFSPolicySettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGFSPolicySettingsModel(isEnabled bool) *GFSPolicySettingsModel {
	this := GFSPolicySettingsModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewGFSPolicySettingsModelWithDefaults instantiates a new GFSPolicySettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGFSPolicySettingsModelWithDefaults() *GFSPolicySettingsModel {
	this := GFSPolicySettingsModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *GFSPolicySettingsModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *GFSPolicySettingsModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetWeekly returns the Weekly field value if set, zero value otherwise.
func (o *GFSPolicySettingsModel) GetWeekly() GFSPolicySettingsWeeklyModel {
	if o == nil || o.Weekly == nil {
		var ret GFSPolicySettingsWeeklyModel
		return ret
	}
	return *o.Weekly
}

// GetWeeklyOk returns a tuple with the Weekly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsModel) GetWeeklyOk() (*GFSPolicySettingsWeeklyModel, bool) {
	if o == nil || o.Weekly == nil {
		return nil, false
	}
	return o.Weekly, true
}

// HasWeekly returns a boolean if a field has been set.
func (o *GFSPolicySettingsModel) HasWeekly() bool {
	if o != nil && o.Weekly != nil {
		return true
	}

	return false
}

// SetWeekly gets a reference to the given GFSPolicySettingsWeeklyModel and assigns it to the Weekly field.
func (o *GFSPolicySettingsModel) SetWeekly(v GFSPolicySettingsWeeklyModel) {
	o.Weekly = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *GFSPolicySettingsModel) GetMonthly() GFSPolicySettingsMonthlyModel {
	if o == nil || o.Monthly == nil {
		var ret GFSPolicySettingsMonthlyModel
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsModel) GetMonthlyOk() (*GFSPolicySettingsMonthlyModel, bool) {
	if o == nil || o.Monthly == nil {
		return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *GFSPolicySettingsModel) HasMonthly() bool {
	if o != nil && o.Monthly != nil {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given GFSPolicySettingsMonthlyModel and assigns it to the Monthly field.
func (o *GFSPolicySettingsModel) SetMonthly(v GFSPolicySettingsMonthlyModel) {
	o.Monthly = &v
}

// GetYearly returns the Yearly field value if set, zero value otherwise.
func (o *GFSPolicySettingsModel) GetYearly() GFSPolicySettingsYearlyModel {
	if o == nil || o.Yearly == nil {
		var ret GFSPolicySettingsYearlyModel
		return ret
	}
	return *o.Yearly
}

// GetYearlyOk returns a tuple with the Yearly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GFSPolicySettingsModel) GetYearlyOk() (*GFSPolicySettingsYearlyModel, bool) {
	if o == nil || o.Yearly == nil {
		return nil, false
	}
	return o.Yearly, true
}

// HasYearly returns a boolean if a field has been set.
func (o *GFSPolicySettingsModel) HasYearly() bool {
	if o != nil && o.Yearly != nil {
		return true
	}

	return false
}

// SetYearly gets a reference to the given GFSPolicySettingsYearlyModel and assigns it to the Yearly field.
func (o *GFSPolicySettingsModel) SetYearly(v GFSPolicySettingsYearlyModel) {
	o.Yearly = &v
}

func (o GFSPolicySettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Weekly != nil {
		toSerialize["weekly"] = o.Weekly
	}
	if o.Monthly != nil {
		toSerialize["monthly"] = o.Monthly
	}
	if o.Yearly != nil {
		toSerialize["yearly"] = o.Yearly
	}
	return json.Marshal(toSerialize)
}

type NullableGFSPolicySettingsModel struct {
	value *GFSPolicySettingsModel
	isSet bool
}

func (v NullableGFSPolicySettingsModel) Get() *GFSPolicySettingsModel {
	return v.value
}

func (v *NullableGFSPolicySettingsModel) Set(val *GFSPolicySettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGFSPolicySettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGFSPolicySettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGFSPolicySettingsModel(val *GFSPolicySettingsModel) *NullableGFSPolicySettingsModel {
	return &NullableGFSPolicySettingsModel{value: val, isSet: true}
}

func (v NullableGFSPolicySettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGFSPolicySettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

