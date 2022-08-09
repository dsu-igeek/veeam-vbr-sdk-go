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

// WindowsHostSpec struct for WindowsHostSpec
type WindowsHostSpec struct {
	ManagedServerSpec
	NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
}

// NewWindowsHostSpec instantiates a new WindowsHostSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsHostSpec() *WindowsHostSpec {
	this := WindowsHostSpec{}
	return &this
}

// NewWindowsHostSpecWithDefaults instantiates a new WindowsHostSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsHostSpecWithDefaults() *WindowsHostSpec {
	this := WindowsHostSpec{}
	return &this
}

// GetNetworkSettings returns the NetworkSettings field value if set, zero value otherwise.
func (o *WindowsHostSpec) GetNetworkSettings() WindowsHostPortsModel {
	if o == nil || o.NetworkSettings == nil {
		var ret WindowsHostPortsModel
		return ret
	}
	return *o.NetworkSettings
}

// GetNetworkSettingsOk returns a tuple with the NetworkSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsHostSpec) GetNetworkSettingsOk() (*WindowsHostPortsModel, bool) {
	if o == nil || o.NetworkSettings == nil {
		return nil, false
	}
	return o.NetworkSettings, true
}

// HasNetworkSettings returns a boolean if a field has been set.
func (o *WindowsHostSpec) HasNetworkSettings() bool {
	if o != nil && o.NetworkSettings != nil {
		return true
	}

	return false
}

// SetNetworkSettings gets a reference to the given WindowsHostPortsModel and assigns it to the NetworkSettings field.
func (o *WindowsHostSpec) SetNetworkSettings(v WindowsHostPortsModel) {
	o.NetworkSettings = &v
}

func (o WindowsHostSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedServerSpec, errManagedServerSpec := json.Marshal(o.ManagedServerSpec)
	if errManagedServerSpec != nil {
		return []byte{}, errManagedServerSpec
	}
	errManagedServerSpec = json.Unmarshal([]byte(serializedManagedServerSpec), &toSerialize)
	if errManagedServerSpec != nil {
		return []byte{}, errManagedServerSpec
	}
	if o.NetworkSettings != nil {
		toSerialize["networkSettings"] = o.NetworkSettings
	}
	return json.Marshal(toSerialize)
}

type NullableWindowsHostSpec struct {
	value *WindowsHostSpec
	isSet bool
}

func (v NullableWindowsHostSpec) Get() *WindowsHostSpec {
	return v.value
}

func (v *NullableWindowsHostSpec) Set(val *WindowsHostSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsHostSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsHostSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsHostSpec(val *WindowsHostSpec) *NullableWindowsHostSpec {
	return &NullableWindowsHostSpec{value: val, isSet: true}
}

func (v NullableWindowsHostSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsHostSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


