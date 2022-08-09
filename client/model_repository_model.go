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

// RepositoryModel struct for RepositoryModel
type RepositoryModel struct {
	// ID of the backup repository.
	Id string `json:"id"`
	// Name of the backup repository.
	Name string `json:"name"`
	// Description of the backup repository.
	Description string `json:"description"`
	// VMware vSphere tag assigned to the backup repository.
	Tag *string `json:"tag,omitempty"`
	Type ERepositoryType `json:"type"`
}

// NewRepositoryModel instantiates a new RepositoryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryModel(id string, name string, description string, type_ ERepositoryType, ) *RepositoryModel {
	this := RepositoryModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	return &this
}

// NewRepositoryModelWithDefaults instantiates a new RepositoryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryModelWithDefaults() *RepositoryModel {
	this := RepositoryModel{}
	return &this
}

// GetId returns the Id field value
func (o *RepositoryModel) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RepositoryModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RepositoryModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RepositoryModel) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RepositoryModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RepositoryModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *RepositoryModel) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RepositoryModel) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RepositoryModel) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *RepositoryModel) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryModel) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *RepositoryModel) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *RepositoryModel) SetTag(v string) {
	o.Tag = &v
}

// GetType returns the Type field value
func (o *RepositoryModel) GetType() ERepositoryType {
	if o == nil  {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RepositoryModel) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RepositoryModel) SetType(v ERepositoryType) {
	o.Type = v
}

func (o RepositoryModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryModel struct {
	value *RepositoryModel
	isSet bool
}

func (v NullableRepositoryModel) Get() *RepositoryModel {
	return v.value
}

func (v *NullableRepositoryModel) Set(val *RepositoryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryModel(val *RepositoryModel) *NullableRepositoryModel {
	return &NullableRepositoryModel{value: val, isSet: true}
}

func (v NullableRepositoryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


