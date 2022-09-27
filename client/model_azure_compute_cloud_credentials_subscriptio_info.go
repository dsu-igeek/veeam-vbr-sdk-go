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

// AzureComputeCloudCredentialsSubscriptioInfo struct for AzureComputeCloudCredentialsSubscriptioInfo
type AzureComputeCloudCredentialsSubscriptioInfo struct {
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// subscription as it is in Azure
	OriginalSubscriptionId *string `json:"originalSubscriptionId,omitempty"`
}

// NewAzureComputeCloudCredentialsSubscriptioInfo instantiates a new AzureComputeCloudCredentialsSubscriptioInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeCloudCredentialsSubscriptioInfo() *AzureComputeCloudCredentialsSubscriptioInfo {
	this := AzureComputeCloudCredentialsSubscriptioInfo{}
	return &this
}

// NewAzureComputeCloudCredentialsSubscriptioInfoWithDefaults instantiates a new AzureComputeCloudCredentialsSubscriptioInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeCloudCredentialsSubscriptioInfoWithDefaults() *AzureComputeCloudCredentialsSubscriptioInfo {
	this := AzureComputeCloudCredentialsSubscriptioInfo{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *AzureComputeCloudCredentialsSubscriptioInfo) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsSubscriptioInfo) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AzureComputeCloudCredentialsSubscriptioInfo) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *AzureComputeCloudCredentialsSubscriptioInfo) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetOriginalSubscriptionId returns the OriginalSubscriptionId field value if set, zero value otherwise.
func (o *AzureComputeCloudCredentialsSubscriptioInfo) GetOriginalSubscriptionId() string {
	if o == nil || o.OriginalSubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.OriginalSubscriptionId
}

// GetOriginalSubscriptionIdOk returns a tuple with the OriginalSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsSubscriptioInfo) GetOriginalSubscriptionIdOk() (*string, bool) {
	if o == nil || o.OriginalSubscriptionId == nil {
		return nil, false
	}
	return o.OriginalSubscriptionId, true
}

// HasOriginalSubscriptionId returns a boolean if a field has been set.
func (o *AzureComputeCloudCredentialsSubscriptioInfo) HasOriginalSubscriptionId() bool {
	if o != nil && o.OriginalSubscriptionId != nil {
		return true
	}

	return false
}

// SetOriginalSubscriptionId gets a reference to the given string and assigns it to the OriginalSubscriptionId field.
func (o *AzureComputeCloudCredentialsSubscriptioInfo) SetOriginalSubscriptionId(v string) {
	o.OriginalSubscriptionId = &v
}

func (o AzureComputeCloudCredentialsSubscriptioInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.OriginalSubscriptionId != nil {
		toSerialize["originalSubscriptionId"] = o.OriginalSubscriptionId
	}
	return json.Marshal(toSerialize)
}

type NullableAzureComputeCloudCredentialsSubscriptioInfo struct {
	value *AzureComputeCloudCredentialsSubscriptioInfo
	isSet bool
}

func (v NullableAzureComputeCloudCredentialsSubscriptioInfo) Get() *AzureComputeCloudCredentialsSubscriptioInfo {
	return v.value
}

func (v *NullableAzureComputeCloudCredentialsSubscriptioInfo) Set(val *AzureComputeCloudCredentialsSubscriptioInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeCloudCredentialsSubscriptioInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeCloudCredentialsSubscriptioInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeCloudCredentialsSubscriptioInfo(val *AzureComputeCloudCredentialsSubscriptioInfo) *NullableAzureComputeCloudCredentialsSubscriptioInfo {
	return &NullableAzureComputeCloudCredentialsSubscriptioInfo{value: val, isSet: true}
}

func (v NullableAzureComputeCloudCredentialsSubscriptioInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeCloudCredentialsSubscriptioInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


