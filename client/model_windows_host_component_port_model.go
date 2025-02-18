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

// WindowsHostComponentPortModel Ports used by Veeam Backup & Replication components.
type WindowsHostComponentPortModel struct {
	ComponentName EWindowsHostComponentType `json:"componentName"`
	// Port used by the component.
	Port int32 `json:"port"`
}

// NewWindowsHostComponentPortModel instantiates a new WindowsHostComponentPortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsHostComponentPortModel(componentName EWindowsHostComponentType, port int32) *WindowsHostComponentPortModel {
	this := WindowsHostComponentPortModel{}
	this.ComponentName = componentName
	this.Port = port
	return &this
}

// NewWindowsHostComponentPortModelWithDefaults instantiates a new WindowsHostComponentPortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsHostComponentPortModelWithDefaults() *WindowsHostComponentPortModel {
	this := WindowsHostComponentPortModel{}
	return &this
}

// GetComponentName returns the ComponentName field value
func (o *WindowsHostComponentPortModel) GetComponentName() EWindowsHostComponentType {
	if o == nil {
		var ret EWindowsHostComponentType
		return ret
	}

	return o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value
// and a boolean to check if the value has been set.
func (o *WindowsHostComponentPortModel) GetComponentNameOk() (*EWindowsHostComponentType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ComponentName, true
}

// SetComponentName sets field value
func (o *WindowsHostComponentPortModel) SetComponentName(v EWindowsHostComponentType) {
	o.ComponentName = v
}

// GetPort returns the Port field value
func (o *WindowsHostComponentPortModel) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *WindowsHostComponentPortModel) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *WindowsHostComponentPortModel) SetPort(v int32) {
	o.Port = v
}

func (o WindowsHostComponentPortModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["componentName"] = o.ComponentName
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableWindowsHostComponentPortModel struct {
	value *WindowsHostComponentPortModel
	isSet bool
}

func (v NullableWindowsHostComponentPortModel) Get() *WindowsHostComponentPortModel {
	return v.value
}

func (v *NullableWindowsHostComponentPortModel) Set(val *WindowsHostComponentPortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsHostComponentPortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsHostComponentPortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsHostComponentPortModel(val *WindowsHostComponentPortModel) *NullableWindowsHostComponentPortModel {
	return &NullableWindowsHostComponentPortModel{value: val, isSet: true}
}

func (v NullableWindowsHostComponentPortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsHostComponentPortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


