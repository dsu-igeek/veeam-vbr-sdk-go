/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev2
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CredentialsImportSpecCollection struct for CredentialsImportSpecCollection
type CredentialsImportSpecCollection struct {
	// Array of credentials.
	Credentials []CredentialsImportSpec `json:"credentials"`
}

// NewCredentialsImportSpecCollection instantiates a new CredentialsImportSpecCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsImportSpecCollection(credentials []CredentialsImportSpec, ) *CredentialsImportSpecCollection {
	this := CredentialsImportSpecCollection{}
	this.Credentials = credentials
	return &this
}

// NewCredentialsImportSpecCollectionWithDefaults instantiates a new CredentialsImportSpecCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsImportSpecCollectionWithDefaults() *CredentialsImportSpecCollection {
	this := CredentialsImportSpecCollection{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *CredentialsImportSpecCollection) GetCredentials() []CredentialsImportSpec {
	if o == nil  {
		var ret []CredentialsImportSpec
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *CredentialsImportSpecCollection) GetCredentialsOk() (*[]CredentialsImportSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *CredentialsImportSpecCollection) SetCredentials(v []CredentialsImportSpec) {
	o.Credentials = v
}

func (o CredentialsImportSpecCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsImportSpecCollection struct {
	value *CredentialsImportSpecCollection
	isSet bool
}

func (v NullableCredentialsImportSpecCollection) Get() *CredentialsImportSpecCollection {
	return v.value
}

func (v *NullableCredentialsImportSpecCollection) Set(val *CredentialsImportSpecCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsImportSpecCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsImportSpecCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsImportSpecCollection(val *CredentialsImportSpecCollection) *NullableCredentialsImportSpecCollection {
	return &NullableCredentialsImportSpecCollection{value: val, isSet: true}
}

func (v NullableCredentialsImportSpecCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsImportSpecCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


