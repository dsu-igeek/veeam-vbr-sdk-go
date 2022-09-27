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

// WindowsHostModel struct for WindowsHostModel
type WindowsHostModel struct {
	ManagedServerModel
	NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
}

// NewWindowsHostModel instantiates a new WindowsHostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsHostModel() *WindowsHostModel {
	this := WindowsHostModel{}
	return &this
}

// NewWindowsHostModelWithDefaults instantiates a new WindowsHostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsHostModelWithDefaults() *WindowsHostModel {
	this := WindowsHostModel{}
	return &this
}

// GetNetworkSettings returns the NetworkSettings field value if set, zero value otherwise.
func (o *WindowsHostModel) GetNetworkSettings() WindowsHostPortsModel {
	if o == nil || o.NetworkSettings == nil {
		var ret WindowsHostPortsModel
		return ret
	}
	return *o.NetworkSettings
}

// GetNetworkSettingsOk returns a tuple with the NetworkSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsHostModel) GetNetworkSettingsOk() (*WindowsHostPortsModel, bool) {
	if o == nil || o.NetworkSettings == nil {
		return nil, false
	}
	return o.NetworkSettings, true
}

// HasNetworkSettings returns a boolean if a field has been set.
func (o *WindowsHostModel) HasNetworkSettings() bool {
	if o != nil && o.NetworkSettings != nil {
		return true
	}

	return false
}

// SetNetworkSettings gets a reference to the given WindowsHostPortsModel and assigns it to the NetworkSettings field.
func (o *WindowsHostModel) SetNetworkSettings(v WindowsHostPortsModel) {
	o.NetworkSettings = &v
}

func (o WindowsHostModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedServerModel, errManagedServerModel := json.Marshal(o.ManagedServerModel)
	if errManagedServerModel != nil {
		return []byte{}, errManagedServerModel
	}
	errManagedServerModel = json.Unmarshal([]byte(serializedManagedServerModel), &toSerialize)
	if errManagedServerModel != nil {
		return []byte{}, errManagedServerModel
	}
	if o.NetworkSettings != nil {
		toSerialize["networkSettings"] = o.NetworkSettings
	}
	return json.Marshal(toSerialize)
}

type NullableWindowsHostModel struct {
	value *WindowsHostModel
	isSet bool
}

func (v NullableWindowsHostModel) Get() *WindowsHostModel {
	return v.value
}

func (v *NullableWindowsHostModel) Set(val *WindowsHostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsHostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsHostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsHostModel(val *WindowsHostModel) *NullableWindowsHostModel {
	return &NullableWindowsHostModel{value: val, isSet: true}
}

func (v NullableWindowsHostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsHostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


