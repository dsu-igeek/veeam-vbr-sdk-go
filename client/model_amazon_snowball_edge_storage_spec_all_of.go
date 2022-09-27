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

// AmazonSnowballEdgeStorageSpecAllOf struct for AmazonSnowballEdgeStorageSpecAllOf
type AmazonSnowballEdgeStorageSpecAllOf struct {
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account AmazonSnowballEdgeStorageAccountModel `json:"account"`
	Bucket AmazonSnowballEdgeStorageBucketModel `json:"bucket"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewAmazonSnowballEdgeStorageSpecAllOf instantiates a new AmazonSnowballEdgeStorageSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonSnowballEdgeStorageSpecAllOf(account AmazonSnowballEdgeStorageAccountModel, bucket AmazonSnowballEdgeStorageBucketModel, mountServer MountServerSettingsModel, ) *AmazonSnowballEdgeStorageSpecAllOf {
	this := AmazonSnowballEdgeStorageSpecAllOf{}
	this.Account = account
	this.Bucket = bucket
	this.MountServer = mountServer
	return &this
}

// NewAmazonSnowballEdgeStorageSpecAllOfWithDefaults instantiates a new AmazonSnowballEdgeStorageSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonSnowballEdgeStorageSpecAllOfWithDefaults() *AmazonSnowballEdgeStorageSpecAllOf {
	this := AmazonSnowballEdgeStorageSpecAllOf{}
	return &this
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetEnableTaskLimit() bool {
	if o == nil || o.EnableTaskLimit == nil {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || o.EnableTaskLimit == nil {
		return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeStorageSpecAllOf) HasEnableTaskLimit() bool {
	if o != nil && o.EnableTaskLimit != nil {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *AmazonSnowballEdgeStorageSpecAllOf) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetMaxTaskCount() int32 {
	if o == nil || o.MaxTaskCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || o.MaxTaskCount == nil {
		return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeStorageSpecAllOf) HasMaxTaskCount() bool {
	if o != nil && o.MaxTaskCount != nil {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *AmazonSnowballEdgeStorageSpecAllOf) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetAccount() AmazonSnowballEdgeStorageAccountModel {
	if o == nil  {
		var ret AmazonSnowballEdgeStorageAccountModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetAccountOk() (*AmazonSnowballEdgeStorageAccountModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AmazonSnowballEdgeStorageSpecAllOf) SetAccount(v AmazonSnowballEdgeStorageAccountModel) {
	o.Account = v
}

// GetBucket returns the Bucket field value
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetBucket() AmazonSnowballEdgeStorageBucketModel {
	if o == nil  {
		var ret AmazonSnowballEdgeStorageBucketModel
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetBucketOk() (*AmazonSnowballEdgeStorageBucketModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *AmazonSnowballEdgeStorageSpecAllOf) SetBucket(v AmazonSnowballEdgeStorageBucketModel) {
	o.Bucket = v
}

// GetMountServer returns the MountServer field value
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetMountServer() MountServerSettingsModel {
	if o == nil  {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageSpecAllOf) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *AmazonSnowballEdgeStorageSpecAllOf) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o AmazonSnowballEdgeStorageSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableAmazonSnowballEdgeStorageSpecAllOf struct {
	value *AmazonSnowballEdgeStorageSpecAllOf
	isSet bool
}

func (v NullableAmazonSnowballEdgeStorageSpecAllOf) Get() *AmazonSnowballEdgeStorageSpecAllOf {
	return v.value
}

func (v *NullableAmazonSnowballEdgeStorageSpecAllOf) Set(val *AmazonSnowballEdgeStorageSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonSnowballEdgeStorageSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonSnowballEdgeStorageSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonSnowballEdgeStorageSpecAllOf(val *AmazonSnowballEdgeStorageSpecAllOf) *NullableAmazonSnowballEdgeStorageSpecAllOf {
	return &NullableAmazonSnowballEdgeStorageSpecAllOf{value: val, isSet: true}
}

func (v NullableAmazonSnowballEdgeStorageSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonSnowballEdgeStorageSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

