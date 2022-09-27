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

// AzureStorageCloudCredentialsSpec struct for AzureStorageCloudCredentialsSpec
type AzureStorageCloudCredentialsSpec struct {
	CloudCredentialsSpec
	// Storage account type.
	Account *string `json:"account,omitempty"`
	// Storage account shared key.
	SharedKey *string `json:"sharedKey,omitempty"`
	// Tag used to identify the credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewAzureStorageCloudCredentialsSpec instantiates a new AzureStorageCloudCredentialsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureStorageCloudCredentialsSpec() *AzureStorageCloudCredentialsSpec {
	this := AzureStorageCloudCredentialsSpec{}
	return &this
}

// NewAzureStorageCloudCredentialsSpecWithDefaults instantiates a new AzureStorageCloudCredentialsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureStorageCloudCredentialsSpecWithDefaults() *AzureStorageCloudCredentialsSpec {
	this := AzureStorageCloudCredentialsSpec{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AzureStorageCloudCredentialsSpec) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageCloudCredentialsSpec) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AzureStorageCloudCredentialsSpec) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *AzureStorageCloudCredentialsSpec) SetAccount(v string) {
	o.Account = &v
}

// GetSharedKey returns the SharedKey field value if set, zero value otherwise.
func (o *AzureStorageCloudCredentialsSpec) GetSharedKey() string {
	if o == nil || o.SharedKey == nil {
		var ret string
		return ret
	}
	return *o.SharedKey
}

// GetSharedKeyOk returns a tuple with the SharedKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageCloudCredentialsSpec) GetSharedKeyOk() (*string, bool) {
	if o == nil || o.SharedKey == nil {
		return nil, false
	}
	return o.SharedKey, true
}

// HasSharedKey returns a boolean if a field has been set.
func (o *AzureStorageCloudCredentialsSpec) HasSharedKey() bool {
	if o != nil && o.SharedKey != nil {
		return true
	}

	return false
}

// SetSharedKey gets a reference to the given string and assigns it to the SharedKey field.
func (o *AzureStorageCloudCredentialsSpec) SetSharedKey(v string) {
	o.SharedKey = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AzureStorageCloudCredentialsSpec) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageCloudCredentialsSpec) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AzureStorageCloudCredentialsSpec) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AzureStorageCloudCredentialsSpec) SetTag(v string) {
	o.Tag = &v
}

func (o AzureStorageCloudCredentialsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudCredentialsSpec, errCloudCredentialsSpec := json.Marshal(o.CloudCredentialsSpec)
	if errCloudCredentialsSpec != nil {
		return []byte{}, errCloudCredentialsSpec
	}
	errCloudCredentialsSpec = json.Unmarshal([]byte(serializedCloudCredentialsSpec), &toSerialize)
	if errCloudCredentialsSpec != nil {
		return []byte{}, errCloudCredentialsSpec
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.SharedKey != nil {
		toSerialize["sharedKey"] = o.SharedKey
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableAzureStorageCloudCredentialsSpec struct {
	value *AzureStorageCloudCredentialsSpec
	isSet bool
}

func (v NullableAzureStorageCloudCredentialsSpec) Get() *AzureStorageCloudCredentialsSpec {
	return v.value
}

func (v *NullableAzureStorageCloudCredentialsSpec) Set(val *AzureStorageCloudCredentialsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureStorageCloudCredentialsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureStorageCloudCredentialsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureStorageCloudCredentialsSpec(val *AzureStorageCloudCredentialsSpec) *NullableAzureStorageCloudCredentialsSpec {
	return &NullableAzureStorageCloudCredentialsSpec{value: val, isSet: true}
}

func (v NullableAzureStorageCloudCredentialsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureStorageCloudCredentialsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

