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

// GoogleCloudStorageBucketBrowserModel struct for GoogleCloudStorageBucketBrowserModel
type GoogleCloudStorageBucketBrowserModel struct {
	Name *string `json:"name,omitempty"`
	Folders *[]string `json:"folders,omitempty"`
}

// NewGoogleCloudStorageBucketBrowserModel instantiates a new GoogleCloudStorageBucketBrowserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudStorageBucketBrowserModel() *GoogleCloudStorageBucketBrowserModel {
	this := GoogleCloudStorageBucketBrowserModel{}
	return &this
}

// NewGoogleCloudStorageBucketBrowserModelWithDefaults instantiates a new GoogleCloudStorageBucketBrowserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudStorageBucketBrowserModelWithDefaults() *GoogleCloudStorageBucketBrowserModel {
	this := GoogleCloudStorageBucketBrowserModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GoogleCloudStorageBucketBrowserModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBucketBrowserModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GoogleCloudStorageBucketBrowserModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GoogleCloudStorageBucketBrowserModel) SetName(v string) {
	o.Name = &v
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *GoogleCloudStorageBucketBrowserModel) GetFolders() []string {
	if o == nil || o.Folders == nil {
		var ret []string
		return ret
	}
	return *o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBucketBrowserModel) GetFoldersOk() (*[]string, bool) {
	if o == nil || o.Folders == nil {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *GoogleCloudStorageBucketBrowserModel) HasFolders() bool {
	if o != nil && o.Folders != nil {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []string and assigns it to the Folders field.
func (o *GoogleCloudStorageBucketBrowserModel) SetFolders(v []string) {
	o.Folders = &v
}

func (o GoogleCloudStorageBucketBrowserModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Folders != nil {
		toSerialize["folders"] = o.Folders
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleCloudStorageBucketBrowserModel struct {
	value *GoogleCloudStorageBucketBrowserModel
	isSet bool
}

func (v NullableGoogleCloudStorageBucketBrowserModel) Get() *GoogleCloudStorageBucketBrowserModel {
	return v.value
}

func (v *NullableGoogleCloudStorageBucketBrowserModel) Set(val *GoogleCloudStorageBucketBrowserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudStorageBucketBrowserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudStorageBucketBrowserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudStorageBucketBrowserModel(val *GoogleCloudStorageBucketBrowserModel) *NullableGoogleCloudStorageBucketBrowserModel {
	return &NullableGoogleCloudStorageBucketBrowserModel{value: val, isSet: true}
}

func (v NullableGoogleCloudStorageBucketBrowserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudStorageBucketBrowserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


