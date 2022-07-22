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

// SmbStorageModel Network attached storage.
type SmbStorageModel struct {
	RepositoryModel
	Share SmbRepositoryShareSettingsModel `json:"share"`
	Repository NetworkRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewSmbStorageModel instantiates a new SmbStorageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbStorageModel(share SmbRepositoryShareSettingsModel, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsModel) *SmbStorageModel {
	this := SmbStorageModel{}
	this.Repository = repository
	this.MountServer = mountServer
	this.Share = share
	return &this
}

// NewSmbStorageModelWithDefaults instantiates a new SmbStorageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbStorageModelWithDefaults() *SmbStorageModel {
	this := SmbStorageModel{}
	return &this
}

// GetShare returns the Share field value
func (o *SmbStorageModel) GetShare() SmbRepositoryShareSettingsModel {
	if o == nil {
		var ret SmbRepositoryShareSettingsModel
		return ret
	}

	return o.Share
}

// GetShareOk returns a tuple with the Share field value
// and a boolean to check if the value has been set.
func (o *SmbStorageModel) GetShareOk() (*SmbRepositoryShareSettingsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Share, true
}

// SetShare sets field value
func (o *SmbStorageModel) SetShare(v SmbRepositoryShareSettingsModel) {
	o.Share = v
}

// GetRepository returns the Repository field value
func (o *SmbStorageModel) GetRepository() NetworkRepositorySettingsModel {
	if o == nil {
		var ret NetworkRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *SmbStorageModel) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *SmbStorageModel) SetRepository(v NetworkRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *SmbStorageModel) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *SmbStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *SmbStorageModel) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o SmbStorageModel) MarshalJSON() ([]byte, error) {
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
		toSerialize["share"] = o.Share
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableSmbStorageModel struct {
	value *SmbStorageModel
	isSet bool
}

func (v NullableSmbStorageModel) Get() *SmbStorageModel {
	return v.value
}

func (v *NullableSmbStorageModel) Set(val *SmbStorageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbStorageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbStorageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbStorageModel(val *SmbStorageModel) *NullableSmbStorageModel {
	return &NullableSmbStorageModel{value: val, isSet: true}
}

func (v NullableSmbStorageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbStorageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


