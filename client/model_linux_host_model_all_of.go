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

// LinuxHostModelAllOf struct for LinuxHostModelAllOf
type LinuxHostModelAllOf struct {
	SshSettings *LinuxHostSSHSettingsModel `json:"sshSettings,omitempty"`
}

// NewLinuxHostModelAllOf instantiates a new LinuxHostModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxHostModelAllOf() *LinuxHostModelAllOf {
	this := LinuxHostModelAllOf{}
	return &this
}

// NewLinuxHostModelAllOfWithDefaults instantiates a new LinuxHostModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxHostModelAllOfWithDefaults() *LinuxHostModelAllOf {
	this := LinuxHostModelAllOf{}
	return &this
}

// GetSshSettings returns the SshSettings field value if set, zero value otherwise.
func (o *LinuxHostModelAllOf) GetSshSettings() LinuxHostSSHSettingsModel {
	if o == nil || o.SshSettings == nil {
		var ret LinuxHostSSHSettingsModel
		return ret
	}
	return *o.SshSettings
}

// GetSshSettingsOk returns a tuple with the SshSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHostModelAllOf) GetSshSettingsOk() (*LinuxHostSSHSettingsModel, bool) {
	if o == nil || o.SshSettings == nil {
		return nil, false
	}
	return o.SshSettings, true
}

// HasSshSettings returns a boolean if a field has been set.
func (o *LinuxHostModelAllOf) HasSshSettings() bool {
	if o != nil && o.SshSettings != nil {
		return true
	}

	return false
}

// SetSshSettings gets a reference to the given LinuxHostSSHSettingsModel and assigns it to the SshSettings field.
func (o *LinuxHostModelAllOf) SetSshSettings(v LinuxHostSSHSettingsModel) {
	o.SshSettings = &v
}

func (o LinuxHostModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SshSettings != nil {
		toSerialize["sshSettings"] = o.SshSettings
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxHostModelAllOf struct {
	value *LinuxHostModelAllOf
	isSet bool
}

func (v NullableLinuxHostModelAllOf) Get() *LinuxHostModelAllOf {
	return v.value
}

func (v *NullableLinuxHostModelAllOf) Set(val *LinuxHostModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxHostModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxHostModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxHostModelAllOf(val *LinuxHostModelAllOf) *NullableLinuxHostModelAllOf {
	return &NullableLinuxHostModelAllOf{value: val, isSet: true}
}

func (v NullableLinuxHostModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxHostModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


