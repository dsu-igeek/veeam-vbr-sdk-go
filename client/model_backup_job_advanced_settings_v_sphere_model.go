/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 

API version: 1.0-rev2
Contact: support@veeam.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BackupJobAdvancedSettingsVSphereModel VMware vSphere settings for the backup job.
type BackupJobAdvancedSettingsVSphereModel struct {
	// If *true*, VMware Tools quiescence is enabled for freezing the VM file system and application data.
	EnableVMWareToolsQuiescence *bool `json:"enableVMWareToolsQuiescence,omitempty"`
	ChangedBlockTracking *VSphereChangedBlockTrackingSettingsModel `json:"changedBlockTracking,omitempty"`
}

// NewBackupJobAdvancedSettingsVSphereModel instantiates a new BackupJobAdvancedSettingsVSphereModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobAdvancedSettingsVSphereModel() *BackupJobAdvancedSettingsVSphereModel {
	this := BackupJobAdvancedSettingsVSphereModel{}
	return &this
}

// NewBackupJobAdvancedSettingsVSphereModelWithDefaults instantiates a new BackupJobAdvancedSettingsVSphereModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobAdvancedSettingsVSphereModelWithDefaults() *BackupJobAdvancedSettingsVSphereModel {
	this := BackupJobAdvancedSettingsVSphereModel{}
	return &this
}

// GetEnableVMWareToolsQuiescence returns the EnableVMWareToolsQuiescence field value if set, zero value otherwise.
func (o *BackupJobAdvancedSettingsVSphereModel) GetEnableVMWareToolsQuiescence() bool {
	if o == nil || o.EnableVMWareToolsQuiescence == nil {
		var ret bool
		return ret
	}
	return *o.EnableVMWareToolsQuiescence
}

// GetEnableVMWareToolsQuiescenceOk returns a tuple with the EnableVMWareToolsQuiescence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobAdvancedSettingsVSphereModel) GetEnableVMWareToolsQuiescenceOk() (*bool, bool) {
	if o == nil || o.EnableVMWareToolsQuiescence == nil {
		return nil, false
	}
	return o.EnableVMWareToolsQuiescence, true
}

// HasEnableVMWareToolsQuiescence returns a boolean if a field has been set.
func (o *BackupJobAdvancedSettingsVSphereModel) HasEnableVMWareToolsQuiescence() bool {
	if o != nil && o.EnableVMWareToolsQuiescence != nil {
		return true
	}

	return false
}

// SetEnableVMWareToolsQuiescence gets a reference to the given bool and assigns it to the EnableVMWareToolsQuiescence field.
func (o *BackupJobAdvancedSettingsVSphereModel) SetEnableVMWareToolsQuiescence(v bool) {
	o.EnableVMWareToolsQuiescence = &v
}

// GetChangedBlockTracking returns the ChangedBlockTracking field value if set, zero value otherwise.
func (o *BackupJobAdvancedSettingsVSphereModel) GetChangedBlockTracking() VSphereChangedBlockTrackingSettingsModel {
	if o == nil || o.ChangedBlockTracking == nil {
		var ret VSphereChangedBlockTrackingSettingsModel
		return ret
	}
	return *o.ChangedBlockTracking
}

// GetChangedBlockTrackingOk returns a tuple with the ChangedBlockTracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobAdvancedSettingsVSphereModel) GetChangedBlockTrackingOk() (*VSphereChangedBlockTrackingSettingsModel, bool) {
	if o == nil || o.ChangedBlockTracking == nil {
		return nil, false
	}
	return o.ChangedBlockTracking, true
}

// HasChangedBlockTracking returns a boolean if a field has been set.
func (o *BackupJobAdvancedSettingsVSphereModel) HasChangedBlockTracking() bool {
	if o != nil && o.ChangedBlockTracking != nil {
		return true
	}

	return false
}

// SetChangedBlockTracking gets a reference to the given VSphereChangedBlockTrackingSettingsModel and assigns it to the ChangedBlockTracking field.
func (o *BackupJobAdvancedSettingsVSphereModel) SetChangedBlockTracking(v VSphereChangedBlockTrackingSettingsModel) {
	o.ChangedBlockTracking = &v
}

func (o BackupJobAdvancedSettingsVSphereModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableVMWareToolsQuiescence != nil {
		toSerialize["enableVMWareToolsQuiescence"] = o.EnableVMWareToolsQuiescence
	}
	if o.ChangedBlockTracking != nil {
		toSerialize["changedBlockTracking"] = o.ChangedBlockTracking
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobAdvancedSettingsVSphereModel struct {
	value *BackupJobAdvancedSettingsVSphereModel
	isSet bool
}

func (v NullableBackupJobAdvancedSettingsVSphereModel) Get() *BackupJobAdvancedSettingsVSphereModel {
	return v.value
}

func (v *NullableBackupJobAdvancedSettingsVSphereModel) Set(val *BackupJobAdvancedSettingsVSphereModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobAdvancedSettingsVSphereModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobAdvancedSettingsVSphereModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobAdvancedSettingsVSphereModel(val *BackupJobAdvancedSettingsVSphereModel) *NullableBackupJobAdvancedSettingsVSphereModel {
	return &NullableBackupJobAdvancedSettingsVSphereModel{value: val, isSet: true}
}

func (v NullableBackupJobAdvancedSettingsVSphereModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobAdvancedSettingsVSphereModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


