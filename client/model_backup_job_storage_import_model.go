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

// BackupJobStorageImportModel Backup storage settings.
type BackupJobStorageImportModel struct {
	BackupRepository BackupRepositoryImportModel `json:"backupRepository"`
	BackupProxies BackupJobImportProxiesModel `json:"backupProxies"`
	RetentionPolicy BackupJobRetentionPolicySettingsModel `json:"retentionPolicy"`
	GfsPolicy *GFSPolicySettingsModel `json:"gfsPolicy,omitempty"`
	AdvancedSettings *BackupJobAdvancedSettingsModel `json:"advancedSettings,omitempty"`
}

// NewBackupJobStorageImportModel instantiates a new BackupJobStorageImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobStorageImportModel(backupRepository BackupRepositoryImportModel, backupProxies BackupJobImportProxiesModel, retentionPolicy BackupJobRetentionPolicySettingsModel) *BackupJobStorageImportModel {
	this := BackupJobStorageImportModel{}
	this.BackupRepository = backupRepository
	this.BackupProxies = backupProxies
	this.RetentionPolicy = retentionPolicy
	return &this
}

// NewBackupJobStorageImportModelWithDefaults instantiates a new BackupJobStorageImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobStorageImportModelWithDefaults() *BackupJobStorageImportModel {
	this := BackupJobStorageImportModel{}
	return &this
}

// GetBackupRepository returns the BackupRepository field value
func (o *BackupJobStorageImportModel) GetBackupRepository() BackupRepositoryImportModel {
	if o == nil {
		var ret BackupRepositoryImportModel
		return ret
	}

	return o.BackupRepository
}

// GetBackupRepositoryOk returns a tuple with the BackupRepository field value
// and a boolean to check if the value has been set.
func (o *BackupJobStorageImportModel) GetBackupRepositoryOk() (*BackupRepositoryImportModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupRepository, true
}

// SetBackupRepository sets field value
func (o *BackupJobStorageImportModel) SetBackupRepository(v BackupRepositoryImportModel) {
	o.BackupRepository = v
}

// GetBackupProxies returns the BackupProxies field value
func (o *BackupJobStorageImportModel) GetBackupProxies() BackupJobImportProxiesModel {
	if o == nil {
		var ret BackupJobImportProxiesModel
		return ret
	}

	return o.BackupProxies
}

// GetBackupProxiesOk returns a tuple with the BackupProxies field value
// and a boolean to check if the value has been set.
func (o *BackupJobStorageImportModel) GetBackupProxiesOk() (*BackupJobImportProxiesModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupProxies, true
}

// SetBackupProxies sets field value
func (o *BackupJobStorageImportModel) SetBackupProxies(v BackupJobImportProxiesModel) {
	o.BackupProxies = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *BackupJobStorageImportModel) GetRetentionPolicy() BackupJobRetentionPolicySettingsModel {
	if o == nil {
		var ret BackupJobRetentionPolicySettingsModel
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *BackupJobStorageImportModel) GetRetentionPolicyOk() (*BackupJobRetentionPolicySettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *BackupJobStorageImportModel) SetRetentionPolicy(v BackupJobRetentionPolicySettingsModel) {
	o.RetentionPolicy = v
}

// GetGfsPolicy returns the GfsPolicy field value if set, zero value otherwise.
func (o *BackupJobStorageImportModel) GetGfsPolicy() GFSPolicySettingsModel {
	if o == nil || o.GfsPolicy == nil {
		var ret GFSPolicySettingsModel
		return ret
	}
	return *o.GfsPolicy
}

// GetGfsPolicyOk returns a tuple with the GfsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobStorageImportModel) GetGfsPolicyOk() (*GFSPolicySettingsModel, bool) {
	if o == nil || o.GfsPolicy == nil {
		return nil, false
	}
	return o.GfsPolicy, true
}

// HasGfsPolicy returns a boolean if a field has been set.
func (o *BackupJobStorageImportModel) HasGfsPolicy() bool {
	if o != nil && o.GfsPolicy != nil {
		return true
	}

	return false
}

// SetGfsPolicy gets a reference to the given GFSPolicySettingsModel and assigns it to the GfsPolicy field.
func (o *BackupJobStorageImportModel) SetGfsPolicy(v GFSPolicySettingsModel) {
	o.GfsPolicy = &v
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *BackupJobStorageImportModel) GetAdvancedSettings() BackupJobAdvancedSettingsModel {
	if o == nil || o.AdvancedSettings == nil {
		var ret BackupJobAdvancedSettingsModel
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobStorageImportModel) GetAdvancedSettingsOk() (*BackupJobAdvancedSettingsModel, bool) {
	if o == nil || o.AdvancedSettings == nil {
		return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *BackupJobStorageImportModel) HasAdvancedSettings() bool {
	if o != nil && o.AdvancedSettings != nil {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given BackupJobAdvancedSettingsModel and assigns it to the AdvancedSettings field.
func (o *BackupJobStorageImportModel) SetAdvancedSettings(v BackupJobAdvancedSettingsModel) {
	o.AdvancedSettings = &v
}

func (o BackupJobStorageImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupRepository"] = o.BackupRepository
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

type NullableBackupJobStorageImportModel struct {
	value *BackupJobStorageImportModel
	isSet bool
}

func (v NullableBackupJobStorageImportModel) Get() *BackupJobStorageImportModel {
	return v.value
}

func (v *NullableBackupJobStorageImportModel) Set(val *BackupJobStorageImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobStorageImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobStorageImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobStorageImportModel(val *BackupJobStorageImportModel) *NullableBackupJobStorageImportModel {
	return &NullableBackupJobStorageImportModel{value: val, isSet: true}
}

func (v NullableBackupJobStorageImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobStorageImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


