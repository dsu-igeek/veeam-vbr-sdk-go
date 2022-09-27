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

// AzureComputeCloudCredentialsSpecAllOf struct for AzureComputeCloudCredentialsSpecAllOf
type AzureComputeCloudCredentialsSpecAllOf struct {
	ConnectionName string `json:"connectionName"`
	CreationMode EAzureComputeCredentialsCreationMode `json:"creationMode"`
	ExistingAccount *AzureComputeCredentialsExistingAccountSpec `json:"existingAccount,omitempty"`
	NewAccount *AzureComputeCredentialsNewAccountSpec `json:"newAccount,omitempty"`
	// Tag used to identify the credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewAzureComputeCloudCredentialsSpecAllOf instantiates a new AzureComputeCloudCredentialsSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeCloudCredentialsSpecAllOf(connectionName string, creationMode EAzureComputeCredentialsCreationMode, ) *AzureComputeCloudCredentialsSpecAllOf {
	this := AzureComputeCloudCredentialsSpecAllOf{}
	this.ConnectionName = connectionName
	this.CreationMode = creationMode
	return &this
}

// NewAzureComputeCloudCredentialsSpecAllOfWithDefaults instantiates a new AzureComputeCloudCredentialsSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeCloudCredentialsSpecAllOfWithDefaults() *AzureComputeCloudCredentialsSpecAllOf {
	this := AzureComputeCloudCredentialsSpecAllOf{}
	return &this
}

// GetConnectionName returns the ConnectionName field value
func (o *AzureComputeCloudCredentialsSpecAllOf) GetConnectionName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsSpecAllOf) GetConnectionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectionName, true
}

// SetConnectionName sets field value
func (o *AzureComputeCloudCredentialsSpecAllOf) SetConnectionName(v string) {
	o.ConnectionName = v
}

// GetCreationMode returns the CreationMode field value
func (o *AzureComputeCloudCredentialsSpecAllOf) GetCreationMode() EAzureComputeCredentialsCreationMode {
	if o == nil  {
		var ret EAzureComputeCredentialsCreationMode
		return ret
	}

	return o.CreationMode
}

// GetCreationModeOk returns a tuple with the CreationMode field value
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsSpecAllOf) GetCreationModeOk() (*EAzureComputeCredentialsCreationMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreationMode, true
}

// SetCreationMode sets field value
func (o *AzureComputeCloudCredentialsSpecAllOf) SetCreationMode(v EAzureComputeCredentialsCreationMode) {
	o.CreationMode = v
}

// GetExistingAccount returns the ExistingAccount field value if set, zero value otherwise.
func (o *AzureComputeCloudCredentialsSpecAllOf) GetExistingAccount() AzureComputeCredentialsExistingAccountSpec {
	if o == nil || o.ExistingAccount == nil {
		var ret AzureComputeCredentialsExistingAccountSpec
		return ret
	}
	return *o.ExistingAccount
}

// GetExistingAccountOk returns a tuple with the ExistingAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsSpecAllOf) GetExistingAccountOk() (*AzureComputeCredentialsExistingAccountSpec, bool) {
	if o == nil || o.ExistingAccount == nil {
		return nil, false
	}
	return o.ExistingAccount, true
}

// HasExistingAccount returns a boolean if a field has been set.
func (o *AzureComputeCloudCredentialsSpecAllOf) HasExistingAccount() bool {
	if o != nil && o.ExistingAccount != nil {
		return true
	}

	return false
}

// SetExistingAccount gets a reference to the given AzureComputeCredentialsExistingAccountSpec and assigns it to the ExistingAccount field.
func (o *AzureComputeCloudCredentialsSpecAllOf) SetExistingAccount(v AzureComputeCredentialsExistingAccountSpec) {
	o.ExistingAccount = &v
}

// GetNewAccount returns the NewAccount field value if set, zero value otherwise.
func (o *AzureComputeCloudCredentialsSpecAllOf) GetNewAccount() AzureComputeCredentialsNewAccountSpec {
	if o == nil || o.NewAccount == nil {
		var ret AzureComputeCredentialsNewAccountSpec
		return ret
	}
	return *o.NewAccount
}

// GetNewAccountOk returns a tuple with the NewAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsSpecAllOf) GetNewAccountOk() (*AzureComputeCredentialsNewAccountSpec, bool) {
	if o == nil || o.NewAccount == nil {
		return nil, false
	}
	return o.NewAccount, true
}

// HasNewAccount returns a boolean if a field has been set.
func (o *AzureComputeCloudCredentialsSpecAllOf) HasNewAccount() bool {
	if o != nil && o.NewAccount != nil {
		return true
	}

	return false
}

// SetNewAccount gets a reference to the given AzureComputeCredentialsNewAccountSpec and assigns it to the NewAccount field.
func (o *AzureComputeCloudCredentialsSpecAllOf) SetNewAccount(v AzureComputeCredentialsNewAccountSpec) {
	o.NewAccount = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AzureComputeCloudCredentialsSpecAllOf) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsSpecAllOf) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AzureComputeCloudCredentialsSpecAllOf) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AzureComputeCloudCredentialsSpecAllOf) SetTag(v string) {
	o.Tag = &v
}

func (o AzureComputeCloudCredentialsSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connectionName"] = o.ConnectionName
	}
	if true {
		toSerialize["creationMode"] = o.CreationMode
	}
	if o.ExistingAccount != nil {
		toSerialize["existingAccount"] = o.ExistingAccount
	}
	if o.NewAccount != nil {
		toSerialize["newAccount"] = o.NewAccount
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableAzureComputeCloudCredentialsSpecAllOf struct {
	value *AzureComputeCloudCredentialsSpecAllOf
	isSet bool
}

func (v NullableAzureComputeCloudCredentialsSpecAllOf) Get() *AzureComputeCloudCredentialsSpecAllOf {
	return v.value
}

func (v *NullableAzureComputeCloudCredentialsSpecAllOf) Set(val *AzureComputeCloudCredentialsSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeCloudCredentialsSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeCloudCredentialsSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeCloudCredentialsSpecAllOf(val *AzureComputeCloudCredentialsSpecAllOf) *NullableAzureComputeCloudCredentialsSpecAllOf {
	return &NullableAzureComputeCloudCredentialsSpecAllOf{value: val, isSet: true}
}

func (v NullableAzureComputeCloudCredentialsSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeCloudCredentialsSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

