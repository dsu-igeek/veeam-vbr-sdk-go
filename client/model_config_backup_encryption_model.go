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

// ConfigBackupEncryptionModel Encryption settings.
type ConfigBackupEncryptionModel struct {
	// If *true*, backup encryption is enabled.
	IsEnabled bool `json:"isEnabled"`
	// ID of the password used for encryption.
	PasswordId string `json:"passwordId"`
}

// NewConfigBackupEncryptionModel instantiates a new ConfigBackupEncryptionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigBackupEncryptionModel(isEnabled bool, passwordId string, ) *ConfigBackupEncryptionModel {
	this := ConfigBackupEncryptionModel{}
	this.IsEnabled = isEnabled
	this.PasswordId = passwordId
	return &this
}

// NewConfigBackupEncryptionModelWithDefaults instantiates a new ConfigBackupEncryptionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigBackupEncryptionModelWithDefaults() *ConfigBackupEncryptionModel {
	this := ConfigBackupEncryptionModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *ConfigBackupEncryptionModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigBackupEncryptionModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *ConfigBackupEncryptionModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetPasswordId returns the PasswordId field value
func (o *ConfigBackupEncryptionModel) GetPasswordId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PasswordId
}

// GetPasswordIdOk returns a tuple with the PasswordId field value
// and a boolean to check if the value has been set.
func (o *ConfigBackupEncryptionModel) GetPasswordIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PasswordId, true
}

// SetPasswordId sets field value
func (o *ConfigBackupEncryptionModel) SetPasswordId(v string) {
	o.PasswordId = v
}

func (o ConfigBackupEncryptionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if true {
		toSerialize["passwordId"] = o.PasswordId
	}
	return json.Marshal(toSerialize)
}

type NullableConfigBackupEncryptionModel struct {
	value *ConfigBackupEncryptionModel
	isSet bool
}

func (v NullableConfigBackupEncryptionModel) Get() *ConfigBackupEncryptionModel {
	return v.value
}

func (v *NullableConfigBackupEncryptionModel) Set(val *ConfigBackupEncryptionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigBackupEncryptionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigBackupEncryptionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigBackupEncryptionModel(val *ConfigBackupEncryptionModel) *NullableConfigBackupEncryptionModel {
	return &NullableConfigBackupEncryptionModel{value: val, isSet: true}
}

func (v NullableConfigBackupEncryptionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigBackupEncryptionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


