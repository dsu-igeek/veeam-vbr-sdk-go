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

// IBMCloudStorageAccountModel Account used to access IBM Cloud storage.
type IBMCloudStorageAccountModel struct {
	// Endpoint address of the storage.
	ServicePoint string `json:"servicePoint"`
	// ID of a region where the storage is located.
	RegionId string `json:"regionId"`
	// ID of a credentials record used to access the storage.
	CredentialsId string `json:"credentialsId"`
	ConnectionSettings *ObjectStorageConnectionModel `json:"connectionSettings,omitempty"`
}

// NewIBMCloudStorageAccountModel instantiates a new IBMCloudStorageAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIBMCloudStorageAccountModel(servicePoint string, regionId string, credentialsId string, ) *IBMCloudStorageAccountModel {
	this := IBMCloudStorageAccountModel{}
	this.ServicePoint = servicePoint
	this.RegionId = regionId
	this.CredentialsId = credentialsId
	return &this
}

// NewIBMCloudStorageAccountModelWithDefaults instantiates a new IBMCloudStorageAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIBMCloudStorageAccountModelWithDefaults() *IBMCloudStorageAccountModel {
	this := IBMCloudStorageAccountModel{}
	return &this
}

// GetServicePoint returns the ServicePoint field value
func (o *IBMCloudStorageAccountModel) GetServicePoint() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServicePoint
}

// GetServicePointOk returns a tuple with the ServicePoint field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageAccountModel) GetServicePointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServicePoint, true
}

// SetServicePoint sets field value
func (o *IBMCloudStorageAccountModel) SetServicePoint(v string) {
	o.ServicePoint = v
}

// GetRegionId returns the RegionId field value
func (o *IBMCloudStorageAccountModel) GetRegionId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageAccountModel) GetRegionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *IBMCloudStorageAccountModel) SetRegionId(v string) {
	o.RegionId = v
}

// GetCredentialsId returns the CredentialsId field value
func (o *IBMCloudStorageAccountModel) GetCredentialsId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageAccountModel) GetCredentialsIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CredentialsId, true
}

// SetCredentialsId sets field value
func (o *IBMCloudStorageAccountModel) SetCredentialsId(v string) {
	o.CredentialsId = v
}

// GetConnectionSettings returns the ConnectionSettings field value if set, zero value otherwise.
func (o *IBMCloudStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel {
	if o == nil || o.ConnectionSettings == nil {
		var ret ObjectStorageConnectionModel
		return ret
	}
	return *o.ConnectionSettings
}

// GetConnectionSettingsOk returns a tuple with the ConnectionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool) {
	if o == nil || o.ConnectionSettings == nil {
		return nil, false
	}
	return o.ConnectionSettings, true
}

// HasConnectionSettings returns a boolean if a field has been set.
func (o *IBMCloudStorageAccountModel) HasConnectionSettings() bool {
	if o != nil && o.ConnectionSettings != nil {
		return true
	}

	return false
}

// SetConnectionSettings gets a reference to the given ObjectStorageConnectionModel and assigns it to the ConnectionSettings field.
func (o *IBMCloudStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel) {
	o.ConnectionSettings = &v
}

func (o IBMCloudStorageAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servicePoint"] = o.ServicePoint
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	if o.ConnectionSettings != nil {
		toSerialize["connectionSettings"] = o.ConnectionSettings
	}
	return json.Marshal(toSerialize)
}

type NullableIBMCloudStorageAccountModel struct {
	value *IBMCloudStorageAccountModel
	isSet bool
}

func (v NullableIBMCloudStorageAccountModel) Get() *IBMCloudStorageAccountModel {
	return v.value
}

func (v *NullableIBMCloudStorageAccountModel) Set(val *IBMCloudStorageAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIBMCloudStorageAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIBMCloudStorageAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIBMCloudStorageAccountModel(val *IBMCloudStorageAccountModel) *NullableIBMCloudStorageAccountModel {
	return &NullableIBMCloudStorageAccountModel{value: val, isSet: true}
}

func (v NullableIBMCloudStorageAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIBMCloudStorageAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


