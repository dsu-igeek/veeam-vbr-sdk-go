/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CredentialsSpec struct for CredentialsSpec
type CredentialsSpec struct {
	// User name.
	Username string `json:"username"`
	// Password.
	Password *string `json:"password,omitempty"`
	// Description of the credentials record.
	Description *string `json:"description,omitempty"`
	Type ECredentialsType `json:"type"`
}

// NewCredentialsSpec instantiates a new CredentialsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsSpec(username string, type_ ECredentialsType) *CredentialsSpec {
	this := CredentialsSpec{}
	this.Username = username
	this.Type = type_
	return &this
}

// NewCredentialsSpecWithDefaults instantiates a new CredentialsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsSpecWithDefaults() *CredentialsSpec {
	this := CredentialsSpec{}
	return &this
}

// GetUsername returns the Username field value
func (o *CredentialsSpec) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpec) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CredentialsSpec) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CredentialsSpec) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsSpec) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CredentialsSpec) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CredentialsSpec) SetPassword(v string) {
	o.Password = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CredentialsSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CredentialsSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CredentialsSpec) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *CredentialsSpec) GetType() ECredentialsType {
	if o == nil {
		var ret ECredentialsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpec) GetTypeOk() (*ECredentialsType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CredentialsSpec) SetType(v ECredentialsType) {
	o.Type = v
}

func (o CredentialsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsSpec struct {
	value *CredentialsSpec
	isSet bool
}

func (v NullableCredentialsSpec) Get() *CredentialsSpec {
	return v.value
}

func (v *NullableCredentialsSpec) Set(val *CredentialsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsSpec(val *CredentialsSpec) *NullableCredentialsSpec {
	return &NullableCredentialsSpec{value: val, isSet: true}
}

func (v NullableCredentialsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

