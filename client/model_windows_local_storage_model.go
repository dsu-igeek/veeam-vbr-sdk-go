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

// WindowsLocalStorageModel Direct attached storage.
type WindowsLocalStorageModel struct {
	RepositoryModel
	// ID of the server that is used as a backup repository.
	HostId string `json:"hostId"`
	Repository WindowsLocalRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewWindowsLocalStorageModel instantiates a new WindowsLocalStorageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsLocalStorageModel(hostId string, repository WindowsLocalRepositorySettingsModel, mountServer MountServerSettingsModel) *WindowsLocalStorageModel {
	this := WindowsLocalStorageModel{}
	this.HostId = hostId
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewWindowsLocalStorageModelWithDefaults instantiates a new WindowsLocalStorageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsLocalStorageModelWithDefaults() *WindowsLocalStorageModel {
	this := WindowsLocalStorageModel{}
	return &this
}

// GetHostId returns the HostId field value
func (o *WindowsLocalStorageModel) GetHostId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *WindowsLocalStorageModel) GetHostIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *WindowsLocalStorageModel) SetHostId(v string) {
	o.HostId = v
}

// GetRepository returns the Repository field value
func (o *WindowsLocalStorageModel) GetRepository() WindowsLocalRepositorySettingsModel {
	if o == nil {
		var ret WindowsLocalRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *WindowsLocalStorageModel) GetRepositoryOk() (*WindowsLocalRepositorySettingsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *WindowsLocalStorageModel) SetRepository(v WindowsLocalRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *WindowsLocalStorageModel) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *WindowsLocalStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *WindowsLocalStorageModel) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o WindowsLocalStorageModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRepositoryModel, errRepositoryModel := json.Marshal(o.RepositoryModel)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	errRepositoryModel = json.Unmarshal([]byte(serializedRepositoryModel), &toSerialize)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	if true {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableWindowsLocalStorageModel struct {
	value *WindowsLocalStorageModel
	isSet bool
}

func (v NullableWindowsLocalStorageModel) Get() *WindowsLocalStorageModel {
	return v.value
}

func (v *NullableWindowsLocalStorageModel) Set(val *WindowsLocalStorageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsLocalStorageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsLocalStorageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsLocalStorageModel(val *WindowsLocalStorageModel) *NullableWindowsLocalStorageModel {
	return &NullableWindowsLocalStorageModel{value: val, isSet: true}
}

func (v NullableWindowsLocalStorageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsLocalStorageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


