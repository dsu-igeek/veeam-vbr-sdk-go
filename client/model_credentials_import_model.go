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

// CredentialsImportModel struct for CredentialsImportModel
type CredentialsImportModel struct {
	// User name, account name or access key.
	CredentialsName string `json:"credentialsName"`
	// Tag used to identify the credentials record.
	CredentialsTag *string `json:"credentialsTag,omitempty"`
}

// NewCredentialsImportModel instantiates a new CredentialsImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsImportModel(credentialsName string, ) *CredentialsImportModel {
	this := CredentialsImportModel{}
	this.CredentialsName = credentialsName
	return &this
}

// NewCredentialsImportModelWithDefaults instantiates a new CredentialsImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsImportModelWithDefaults() *CredentialsImportModel {
	this := CredentialsImportModel{}
	return &this
}

// GetCredentialsName returns the CredentialsName field value
func (o *CredentialsImportModel) GetCredentialsName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CredentialsName
}

// GetCredentialsNameOk returns a tuple with the CredentialsName field value
// and a boolean to check if the value has been set.
func (o *CredentialsImportModel) GetCredentialsNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CredentialsName, true
}

// SetCredentialsName sets field value
func (o *CredentialsImportModel) SetCredentialsName(v string) {
	o.CredentialsName = v
}

// GetCredentialsTag returns the CredentialsTag field value if set, zero value otherwise.
func (o *CredentialsImportModel) GetCredentialsTag() string {
	if o == nil || o.CredentialsTag == nil {
		var ret string
		return ret
	}
	return *o.CredentialsTag
}

// GetCredentialsTagOk returns a tuple with the CredentialsTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsImportModel) GetCredentialsTagOk() (*string, bool) {
	if o == nil || o.CredentialsTag == nil {
		return nil, false
	}
	return o.CredentialsTag, true
}

// HasCredentialsTag returns a boolean if a field has been set.
func (o *CredentialsImportModel) HasCredentialsTag() bool {
	if o != nil && o.CredentialsTag != nil {
		return true
	}

	return false
}

// SetCredentialsTag gets a reference to the given string and assigns it to the CredentialsTag field.
func (o *CredentialsImportModel) SetCredentialsTag(v string) {
	o.CredentialsTag = &v
}

func (o CredentialsImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credentialsName"] = o.CredentialsName
	}
	if o.CredentialsTag != nil {
		toSerialize["credentialsTag"] = o.CredentialsTag
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsImportModel struct {
	value *CredentialsImportModel
	isSet bool
}

func (v NullableCredentialsImportModel) Get() *CredentialsImportModel {
	return v.value
}

func (v *NullableCredentialsImportModel) Set(val *CredentialsImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsImportModel(val *CredentialsImportModel) *NullableCredentialsImportModel {
	return &NullableCredentialsImportModel{value: val, isSet: true}
}

func (v NullableCredentialsImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


