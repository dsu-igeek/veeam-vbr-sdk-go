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

// NfsStorageSpecAllOf struct for NfsStorageSpecAllOf
type NfsStorageSpecAllOf struct {
	Share NfsRepositoryShareSettingsModel `json:"share"`
	Repository NetworkRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewNfsStorageSpecAllOf instantiates a new NfsStorageSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsStorageSpecAllOf(share NfsRepositoryShareSettingsModel, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsModel, ) *NfsStorageSpecAllOf {
	this := NfsStorageSpecAllOf{}
	this.Share = share
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewNfsStorageSpecAllOfWithDefaults instantiates a new NfsStorageSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsStorageSpecAllOfWithDefaults() *NfsStorageSpecAllOf {
	this := NfsStorageSpecAllOf{}
	return &this
}

// GetShare returns the Share field value
func (o *NfsStorageSpecAllOf) GetShare() NfsRepositoryShareSettingsModel {
	if o == nil  {
		var ret NfsRepositoryShareSettingsModel
		return ret
	}

	return o.Share
}

// GetShareOk returns a tuple with the Share field value
// and a boolean to check if the value has been set.
func (o *NfsStorageSpecAllOf) GetShareOk() (*NfsRepositoryShareSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Share, true
}

// SetShare sets field value
func (o *NfsStorageSpecAllOf) SetShare(v NfsRepositoryShareSettingsModel) {
	o.Share = v
}

// GetRepository returns the Repository field value
func (o *NfsStorageSpecAllOf) GetRepository() NetworkRepositorySettingsModel {
	if o == nil  {
		var ret NetworkRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *NfsStorageSpecAllOf) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *NfsStorageSpecAllOf) SetRepository(v NetworkRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *NfsStorageSpecAllOf) GetMountServer() MountServerSettingsModel {
	if o == nil  {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *NfsStorageSpecAllOf) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *NfsStorageSpecAllOf) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o NfsStorageSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableNfsStorageSpecAllOf struct {
	value *NfsStorageSpecAllOf
	isSet bool
}

func (v NullableNfsStorageSpecAllOf) Get() *NfsStorageSpecAllOf {
	return v.value
}

func (v *NullableNfsStorageSpecAllOf) Set(val *NfsStorageSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsStorageSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsStorageSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsStorageSpecAllOf(val *NfsStorageSpecAllOf) *NullableNfsStorageSpecAllOf {
	return &NullableNfsStorageSpecAllOf{value: val, isSet: true}
}

func (v NullableNfsStorageSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsStorageSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


