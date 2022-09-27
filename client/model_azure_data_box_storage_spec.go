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

// AzureDataBoxStorageSpec struct for AzureDataBoxStorageSpec
type AzureDataBoxStorageSpec struct {
	RepositorySpec
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account AzureDataBoxStorageAccountModel `json:"account"`
	Container AzureDataBoxStorageContainerModel `json:"container"`
}

// NewAzureDataBoxStorageSpec instantiates a new AzureDataBoxStorageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureDataBoxStorageSpec(account AzureDataBoxStorageAccountModel, container AzureDataBoxStorageContainerModel, ) *AzureDataBoxStorageSpec {
	this := AzureDataBoxStorageSpec{}
	this.Account = account
	this.Container = container
	return &this
}

// NewAzureDataBoxStorageSpecWithDefaults instantiates a new AzureDataBoxStorageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureDataBoxStorageSpecWithDefaults() *AzureDataBoxStorageSpec {
	this := AzureDataBoxStorageSpec{}
	return &this
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *AzureDataBoxStorageSpec) GetEnableTaskLimit() bool {
	if o == nil || o.EnableTaskLimit == nil {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureDataBoxStorageSpec) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || o.EnableTaskLimit == nil {
		return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *AzureDataBoxStorageSpec) HasEnableTaskLimit() bool {
	if o != nil && o.EnableTaskLimit != nil {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *AzureDataBoxStorageSpec) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *AzureDataBoxStorageSpec) GetMaxTaskCount() int32 {
	if o == nil || o.MaxTaskCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureDataBoxStorageSpec) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || o.MaxTaskCount == nil {
		return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *AzureDataBoxStorageSpec) HasMaxTaskCount() bool {
	if o != nil && o.MaxTaskCount != nil {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *AzureDataBoxStorageSpec) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *AzureDataBoxStorageSpec) GetAccount() AzureDataBoxStorageAccountModel {
	if o == nil  {
		var ret AzureDataBoxStorageAccountModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureDataBoxStorageSpec) GetAccountOk() (*AzureDataBoxStorageAccountModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureDataBoxStorageSpec) SetAccount(v AzureDataBoxStorageAccountModel) {
	o.Account = v
}

// GetContainer returns the Container field value
func (o *AzureDataBoxStorageSpec) GetContainer() AzureDataBoxStorageContainerModel {
	if o == nil  {
		var ret AzureDataBoxStorageContainerModel
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureDataBoxStorageSpec) GetContainerOk() (*AzureDataBoxStorageContainerModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *AzureDataBoxStorageSpec) SetContainer(v AzureDataBoxStorageContainerModel) {
	o.Container = v
}

func (o AzureDataBoxStorageSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRepositorySpec, errRepositorySpec := json.Marshal(o.RepositorySpec)
	if errRepositorySpec != nil {
		return []byte{}, errRepositorySpec
	}
	errRepositorySpec = json.Unmarshal([]byte(serializedRepositorySpec), &toSerialize)
	if errRepositorySpec != nil {
		return []byte{}, errRepositorySpec
	}
	if o.EnableTaskLimit != nil {
		toSerialize["enableTaskLimit"] = o.EnableTaskLimit
	}
	if o.MaxTaskCount != nil {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["container"] = o.Container
	}
	return json.Marshal(toSerialize)
}

type NullableAzureDataBoxStorageSpec struct {
	value *AzureDataBoxStorageSpec
	isSet bool
}

func (v NullableAzureDataBoxStorageSpec) Get() *AzureDataBoxStorageSpec {
	return v.value
}

func (v *NullableAzureDataBoxStorageSpec) Set(val *AzureDataBoxStorageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureDataBoxStorageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureDataBoxStorageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureDataBoxStorageSpec(val *AzureDataBoxStorageSpec) *NullableAzureDataBoxStorageSpec {
	return &NullableAzureDataBoxStorageSpec{value: val, isSet: true}
}

func (v NullableAzureDataBoxStorageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureDataBoxStorageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


