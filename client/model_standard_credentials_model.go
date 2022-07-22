/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 

API version: 1.0-rev2
Contact: support@veeam.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// StandardCredentialsModel struct for StandardCredentialsModel
type StandardCredentialsModel struct {
	CredentialsModel
	// Tag used to identify the credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewStandardCredentialsModel instantiates a new StandardCredentialsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardCredentialsModel() *StandardCredentialsModel {
	this := StandardCredentialsModel{}
	return &this
}

// NewStandardCredentialsModelWithDefaults instantiates a new StandardCredentialsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardCredentialsModelWithDefaults() *StandardCredentialsModel {
	this := StandardCredentialsModel{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *StandardCredentialsModel) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardCredentialsModel) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *StandardCredentialsModel) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *StandardCredentialsModel) SetTag(v string) {
	o.Tag = &v
}

func (o StandardCredentialsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCredentialsModel, errCredentialsModel := json.Marshal(o.CredentialsModel)
	if errCredentialsModel != nil {
		return []byte{}, errCredentialsModel
	}
	errCredentialsModel = json.Unmarshal([]byte(serializedCredentialsModel), &toSerialize)
	if errCredentialsModel != nil {
		return []byte{}, errCredentialsModel
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableStandardCredentialsModel struct {
	value *StandardCredentialsModel
	isSet bool
}

func (v NullableStandardCredentialsModel) Get() *StandardCredentialsModel {
	return v.value
}

func (v *NullableStandardCredentialsModel) Set(val *StandardCredentialsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardCredentialsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardCredentialsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardCredentialsModel(val *StandardCredentialsModel) *NullableStandardCredentialsModel {
	return &NullableStandardCredentialsModel{value: val, isSet: true}
}

func (v NullableStandardCredentialsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardCredentialsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


