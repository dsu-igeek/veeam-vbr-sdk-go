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

// BackupJobStorageModel Backup storage settings.
type BackupJobStorageModel struct {
	// ID of the backup repository.
	BackupRepositoryId string `json:"backupRepositoryId"`
	BackupProxies BackupProxiesSettingsModel `json:"backupProxies"`
	RetentionPolicy BackupJobRetentionPolicySettingsModel `json:"retentionPolicy"`
	GfsPolicy *GFSPolicySettingsModel `json:"gfsPolicy,omitempty"`
	AdvancedSettings *BackupJobAdvancedSettingsModel `json:"advancedSettings,omitempty"`
}

// NewBackupJobStorageModel instantiates a new BackupJobStorageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobStorageModel(backupRepositoryId string, backupProxies BackupProxiesSettingsModel, retentionPolicy BackupJobRetentionPolicySettingsModel, ) *BackupJobStorageModel {
	this := BackupJobStorageModel{}
	this.BackupRepositoryId = backupRepositoryId
	this.BackupProxies = backupProxies
	this.RetentionPolicy = retentionPolicy
	return &this
}

// NewBackupJobStorageModelWithDefaults instantiates a new BackupJobStorageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobStorageModelWithDefaults() *BackupJobStorageModel {
	this := BackupJobStorageModel{}
	return &this
}

// GetBackupRepositoryId returns the BackupRepositoryId field value
func (o *BackupJobStorageModel) GetBackupRepositoryId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BackupRepositoryId
}

// GetBackupRepositoryIdOk returns a tuple with the BackupRepositoryId field value
// and a boolean to check if the value has been set.
func (o *BackupJobStorageModel) GetBackupRepositoryIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupRepositoryId, true
}

// SetBackupRepositoryId sets field value
func (o *BackupJobStorageModel) SetBackupRepositoryId(v string) {
	o.BackupRepositoryId = v
}

// GetBackupProxies returns the BackupProxies field value
func (o *BackupJobStorageModel) GetBackupProxies() BackupProxiesSettingsModel {
	if o == nil  {
		var ret BackupProxiesSettingsModel
		return ret
	}

	return o.BackupProxies
}

// GetBackupProxiesOk returns a tuple with the BackupProxies field value
// and a boolean to check if the value has been set.
func (o *BackupJobStorageModel) GetBackupProxiesOk() (*BackupProxiesSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupProxies, true
}

// SetBackupProxies sets field value
func (o *BackupJobStorageModel) SetBackupProxies(v BackupProxiesSettingsModel) {
	o.BackupProxies = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *BackupJobStorageModel) GetRetentionPolicy() BackupJobRetentionPolicySettingsModel {
	if o == nil  {
		var ret BackupJobRetentionPolicySettingsModel
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *BackupJobStorageModel) GetRetentionPolicyOk() (*BackupJobRetentionPolicySettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *BackupJobStorageModel) SetRetentionPolicy(v BackupJobRetentionPolicySettingsModel) {
	o.RetentionPolicy = v
}

// GetGfsPolicy returns the GfsPolicy field value if set, zero value otherwise.
func (o *BackupJobStorageModel) GetGfsPolicy() GFSPolicySettingsModel {
	if o == nil || o.GfsPolicy == nil {
		var ret GFSPolicySettingsModel
		return ret
	}
	return *o.GfsPolicy
}

// GetGfsPolicyOk returns a tuple with the GfsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobStorageModel) GetGfsPolicyOk() (*GFSPolicySettingsModel, bool) {
	if o == nil || o.GfsPolicy == nil {
		return nil, false
	}
	return o.GfsPolicy, true
}

// HasGfsPolicy returns a boolean if a field has been set.
func (o *BackupJobStorageModel) HasGfsPolicy() bool {
	if o != nil && o.GfsPolicy != nil {
		return true
	}

	return false
}

// SetGfsPolicy gets a reference to the given GFSPolicySettingsModel and assigns it to the GfsPolicy field.
func (o *BackupJobStorageModel) SetGfsPolicy(v GFSPolicySettingsModel) {
	o.GfsPolicy = &v
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *BackupJobStorageModel) GetAdvancedSettings() BackupJobAdvancedSettingsModel {
	if o == nil || o.AdvancedSettings == nil {
		var ret BackupJobAdvancedSettingsModel
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobStorageModel) GetAdvancedSettingsOk() (*BackupJobAdvancedSettingsModel, bool) {
	if o == nil || o.AdvancedSettings == nil {
		return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *BackupJobStorageModel) HasAdvancedSettings() bool {
	if o != nil && o.AdvancedSettings != nil {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given BackupJobAdvancedSettingsModel and assigns it to the AdvancedSettings field.
func (o *BackupJobStorageModel) SetAdvancedSettings(v BackupJobAdvancedSettingsModel) {
	o.AdvancedSettings = &v
}

func (o BackupJobStorageModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupRepositoryId"] = o.BackupRepositoryId
	}
	if true {
		toSerialize["backupProxies"] = o.BackupProxies
	}
	if true {
		toSerialize["retentionPolicy"] = o.RetentionPolicy
	}
	if o.GfsPolicy != nil {
		toSerialize["gfsPolicy"] = o.GfsPolicy
	}
	if o.AdvancedSettings != nil {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobStorageModel struct {
	value *BackupJobStorageModel
	isSet bool
}

func (v NullableBackupJobStorageModel) Get() *BackupJobStorageModel {
	return v.value
}

func (v *NullableBackupJobStorageModel) Set(val *BackupJobStorageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobStorageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobStorageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobStorageModel(val *BackupJobStorageModel) *NullableBackupJobStorageModel {
	return &NullableBackupJobStorageModel{value: val, isSet: true}
}

func (v NullableBackupJobStorageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobStorageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


