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

// BackupStorageSettingsEncryptionModel Encription of backup files.
type BackupStorageSettingsEncryptionModel struct {
	// If *true*, the content of backup files is encrypted.
	IsEnabled bool `json:"isEnabled"`
	// ID of the password used for encryption. The value is *null* for exported objects.
	EncryptionPasswordIdOrNull *string `json:"encryptionPasswordIdOrNull,omitempty"`
	// Tag used to identify the password.
	EncryptionPasswordTag *string `json:"encryptionPasswordTag,omitempty"`
}

// NewBackupStorageSettingsEncryptionModel instantiates a new BackupStorageSettingsEncryptionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupStorageSettingsEncryptionModel(isEnabled bool) *BackupStorageSettingsEncryptionModel {
	this := BackupStorageSettingsEncryptionModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewBackupStorageSettingsEncryptionModelWithDefaults instantiates a new BackupStorageSettingsEncryptionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupStorageSettingsEncryptionModelWithDefaults() *BackupStorageSettingsEncryptionModel {
	this := BackupStorageSettingsEncryptionModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *BackupStorageSettingsEncryptionModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingsEncryptionModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *BackupStorageSettingsEncryptionModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetEncryptionPasswordIdOrNull returns the EncryptionPasswordIdOrNull field value if set, zero value otherwise.
func (o *BackupStorageSettingsEncryptionModel) GetEncryptionPasswordIdOrNull() string {
	if o == nil || o.EncryptionPasswordIdOrNull == nil {
		var ret string
		return ret
	}
	return *o.EncryptionPasswordIdOrNull
}

// GetEncryptionPasswordIdOrNullOk returns a tuple with the EncryptionPasswordIdOrNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingsEncryptionModel) GetEncryptionPasswordIdOrNullOk() (*string, bool) {
	if o == nil || o.EncryptionPasswordIdOrNull == nil {
		return nil, false
	}
	return o.EncryptionPasswordIdOrNull, true
}

// HasEncryptionPasswordIdOrNull returns a boolean if a field has been set.
func (o *BackupStorageSettingsEncryptionModel) HasEncryptionPasswordIdOrNull() bool {
	if o != nil && o.EncryptionPasswordIdOrNull != nil {
		return true
	}

	return false
}

// SetEncryptionPasswordIdOrNull gets a reference to the given string and assigns it to the EncryptionPasswordIdOrNull field.
func (o *BackupStorageSettingsEncryptionModel) SetEncryptionPasswordIdOrNull(v string) {
	o.EncryptionPasswordIdOrNull = &v
}

// GetEncryptionPasswordTag returns the EncryptionPasswordTag field value if set, zero value otherwise.
func (o *BackupStorageSettingsEncryptionModel) GetEncryptionPasswordTag() string {
	if o == nil || o.EncryptionPasswordTag == nil {
		var ret string
		return ret
	}
	return *o.EncryptionPasswordTag
}

// GetEncryptionPasswordTagOk returns a tuple with the EncryptionPasswordTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingsEncryptionModel) GetEncryptionPasswordTagOk() (*string, bool) {
	if o == nil || o.EncryptionPasswordTag == nil {
		return nil, false
	}
	return o.EncryptionPasswordTag, true
}

// HasEncryptionPasswordTag returns a boolean if a field has been set.
func (o *BackupStorageSettingsEncryptionModel) HasEncryptionPasswordTag() bool {
	if o != nil && o.EncryptionPasswordTag != nil {
		return true
	}

	return false
}

// SetEncryptionPasswordTag gets a reference to the given string and assigns it to the EncryptionPasswordTag field.
func (o *BackupStorageSettingsEncryptionModel) SetEncryptionPasswordTag(v string) {
	o.EncryptionPasswordTag = &v
}

func (o BackupStorageSettingsEncryptionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.EncryptionPasswordIdOrNull != nil {
		toSerialize["encryptionPasswordIdOrNull"] = o.EncryptionPasswordIdOrNull
	}
	if o.EncryptionPasswordTag != nil {
		toSerialize["encryptionPasswordTag"] = o.EncryptionPasswordTag
	}
	return json.Marshal(toSerialize)
}

type NullableBackupStorageSettingsEncryptionModel struct {
	value *BackupStorageSettingsEncryptionModel
	isSet bool
}

func (v NullableBackupStorageSettingsEncryptionModel) Get() *BackupStorageSettingsEncryptionModel {
	return v.value
}

func (v *NullableBackupStorageSettingsEncryptionModel) Set(val *BackupStorageSettingsEncryptionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupStorageSettingsEncryptionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupStorageSettingsEncryptionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupStorageSettingsEncryptionModel(val *BackupStorageSettingsEncryptionModel) *NullableBackupStorageSettingsEncryptionModel {
	return &NullableBackupStorageSettingsEncryptionModel{value: val, isSet: true}
}

func (v NullableBackupStorageSettingsEncryptionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupStorageSettingsEncryptionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


