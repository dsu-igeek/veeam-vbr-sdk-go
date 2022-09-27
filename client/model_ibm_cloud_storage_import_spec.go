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

// IBMCloudStorageImportSpec struct for IBMCloudStorageImportSpec
type IBMCloudStorageImportSpec struct {
	// Name of the backup repository.
	Name string `json:"name"`
	// Description of the backup repository.
	Description string `json:"description"`
	Tag string `json:"tag"`
	Type ERepositoryType `json:"type"`
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account IBMCloudStorageAccountImportModel `json:"account"`
	Bucket IBMCloudStorageBucketModel `json:"bucket"`
	MountServer MountServerSettingsImportSpec `json:"mountServer"`
}

// NewIBMCloudStorageImportSpec instantiates a new IBMCloudStorageImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIBMCloudStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account IBMCloudStorageAccountImportModel, bucket IBMCloudStorageBucketModel, mountServer MountServerSettingsImportSpec, ) *IBMCloudStorageImportSpec {
	this := IBMCloudStorageImportSpec{}
	this.Name = name
	this.Description = description
	this.Tag = tag
	this.Type = type_
	this.Account = account
	this.Bucket = bucket
	this.MountServer = mountServer
	return &this
}

// NewIBMCloudStorageImportSpecWithDefaults instantiates a new IBMCloudStorageImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIBMCloudStorageImportSpecWithDefaults() *IBMCloudStorageImportSpec {
	this := IBMCloudStorageImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *IBMCloudStorageImportSpec) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IBMCloudStorageImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *IBMCloudStorageImportSpec) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *IBMCloudStorageImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value
func (o *IBMCloudStorageImportSpec) GetTag() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *IBMCloudStorageImportSpec) SetTag(v string) {
	o.Tag = v
}

// GetType returns the Type field value
func (o *IBMCloudStorageImportSpec) GetType() ERepositoryType {
	if o == nil  {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IBMCloudStorageImportSpec) SetType(v ERepositoryType) {
	o.Type = v
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *IBMCloudStorageImportSpec) GetEnableTaskLimit() bool {
	if o == nil || o.EnableTaskLimit == nil {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || o.EnableTaskLimit == nil {
		return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *IBMCloudStorageImportSpec) HasEnableTaskLimit() bool {
	if o != nil && o.EnableTaskLimit != nil {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *IBMCloudStorageImportSpec) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *IBMCloudStorageImportSpec) GetMaxTaskCount() int32 {
	if o == nil || o.MaxTaskCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || o.MaxTaskCount == nil {
		return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *IBMCloudStorageImportSpec) HasMaxTaskCount() bool {
	if o != nil && o.MaxTaskCount != nil {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *IBMCloudStorageImportSpec) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *IBMCloudStorageImportSpec) GetAccount() IBMCloudStorageAccountImportModel {
	if o == nil  {
		var ret IBMCloudStorageAccountImportModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetAccountOk() (*IBMCloudStorageAccountImportModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *IBMCloudStorageImportSpec) SetAccount(v IBMCloudStorageAccountImportModel) {
	o.Account = v
}

// GetBucket returns the Bucket field value
func (o *IBMCloudStorageImportSpec) GetBucket() IBMCloudStorageBucketModel {
	if o == nil  {
		var ret IBMCloudStorageBucketModel
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetBucketOk() (*IBMCloudStorageBucketModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *IBMCloudStorageImportSpec) SetBucket(v IBMCloudStorageBucketModel) {
	o.Bucket = v
}

// GetMountServer returns the MountServer field value
func (o *IBMCloudStorageImportSpec) GetMountServer() MountServerSettingsImportSpec {
	if o == nil  {
		var ret MountServerSettingsImportSpec
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *IBMCloudStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec) {
	o.MountServer = v
}

func (o IBMCloudStorageImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["type"] = o.Type
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
		toSerialize["bucket"] = o.Bucket
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableIBMCloudStorageImportSpec struct {
	value *IBMCloudStorageImportSpec
	isSet bool
}

func (v NullableIBMCloudStorageImportSpec) Get() *IBMCloudStorageImportSpec {
	return v.value
}

func (v *NullableIBMCloudStorageImportSpec) Set(val *IBMCloudStorageImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableIBMCloudStorageImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableIBMCloudStorageImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIBMCloudStorageImportSpec(val *IBMCloudStorageImportSpec) *NullableIBMCloudStorageImportSpec {
	return &NullableIBMCloudStorageImportSpec{value: val, isSet: true}
}

func (v NullableIBMCloudStorageImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIBMCloudStorageImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

