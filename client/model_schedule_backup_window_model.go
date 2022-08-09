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

// ScheduleBackupWindowModel struct for ScheduleBackupWindowModel
type ScheduleBackupWindowModel struct {
	// If *true*, periodic schedule is enabled.
	IsEnabled bool `json:"isEnabled"`
	BackupWindow *BackupWindowSettingModel `json:"backupWindow,omitempty"`
}

// NewScheduleBackupWindowModel instantiates a new ScheduleBackupWindowModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleBackupWindowModel(isEnabled bool, ) *ScheduleBackupWindowModel {
	this := ScheduleBackupWindowModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewScheduleBackupWindowModelWithDefaults instantiates a new ScheduleBackupWindowModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleBackupWindowModelWithDefaults() *ScheduleBackupWindowModel {
	this := ScheduleBackupWindowModel{}
	var isEnabled bool = false
	this.IsEnabled = isEnabled
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *ScheduleBackupWindowModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *ScheduleBackupWindowModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *ScheduleBackupWindowModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetBackupWindow returns the BackupWindow field value if set, zero value otherwise.
func (o *ScheduleBackupWindowModel) GetBackupWindow() BackupWindowSettingModel {
	if o == nil || o.BackupWindow == nil {
		var ret BackupWindowSettingModel
		return ret
	}
	return *o.BackupWindow
}

// GetBackupWindowOk returns a tuple with the BackupWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleBackupWindowModel) GetBackupWindowOk() (*BackupWindowSettingModel, bool) {
	if o == nil || o.BackupWindow == nil {
		return nil, false
	}
	return o.BackupWindow, true
}

// HasBackupWindow returns a boolean if a field has been set.
func (o *ScheduleBackupWindowModel) HasBackupWindow() bool {
	if o != nil && o.BackupWindow != nil {
		return true
	}

	return false
}

// SetBackupWindow gets a reference to the given BackupWindowSettingModel and assigns it to the BackupWindow field.
func (o *ScheduleBackupWindowModel) SetBackupWindow(v BackupWindowSettingModel) {
	o.BackupWindow = &v
}

func (o ScheduleBackupWindowModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.BackupWindow != nil {
		toSerialize["backupWindow"] = o.BackupWindow
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleBackupWindowModel struct {
	value *ScheduleBackupWindowModel
	isSet bool
}

func (v NullableScheduleBackupWindowModel) Get() *ScheduleBackupWindowModel {
	return v.value
}

func (v *NullableScheduleBackupWindowModel) Set(val *ScheduleBackupWindowModel) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleBackupWindowModel) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleBackupWindowModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleBackupWindowModel(val *ScheduleBackupWindowModel) *NullableScheduleBackupWindowModel {
	return &NullableScheduleBackupWindowModel{value: val, isSet: true}
}

func (v NullableScheduleBackupWindowModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleBackupWindowModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


