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

// GoogleCloudCredentialsModel struct for GoogleCloudCredentialsModel
type GoogleCloudCredentialsModel struct {
	CloudCredentialsModel
	// Tag used to identify the credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewGoogleCloudCredentialsModel instantiates a new GoogleCloudCredentialsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudCredentialsModel() *GoogleCloudCredentialsModel {
	this := GoogleCloudCredentialsModel{}
	return &this
}

// NewGoogleCloudCredentialsModelWithDefaults instantiates a new GoogleCloudCredentialsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudCredentialsModelWithDefaults() *GoogleCloudCredentialsModel {
	this := GoogleCloudCredentialsModel{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GoogleCloudCredentialsModel) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudCredentialsModel) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GoogleCloudCredentialsModel) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GoogleCloudCredentialsModel) SetTag(v string) {
	o.Tag = &v
}

func (o GoogleCloudCredentialsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudCredentialsModel, errCloudCredentialsModel := json.Marshal(o.CloudCredentialsModel)
	if errCloudCredentialsModel != nil {
		return []byte{}, errCloudCredentialsModel
	}
	errCloudCredentialsModel = json.Unmarshal([]byte(serializedCloudCredentialsModel), &toSerialize)
	if errCloudCredentialsModel != nil {
		return []byte{}, errCloudCredentialsModel
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleCloudCredentialsModel struct {
	value *GoogleCloudCredentialsModel
	isSet bool
}

func (v NullableGoogleCloudCredentialsModel) Get() *GoogleCloudCredentialsModel {
	return v.value
}

func (v *NullableGoogleCloudCredentialsModel) Set(val *GoogleCloudCredentialsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudCredentialsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudCredentialsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudCredentialsModel(val *GoogleCloudCredentialsModel) *NullableGoogleCloudCredentialsModel {
	return &NullableGoogleCloudCredentialsModel{value: val, isSet: true}
}

func (v NullableGoogleCloudCredentialsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudCredentialsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

