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

// AmazonS3BucketBrowserModel struct for AmazonS3BucketBrowserModel
type AmazonS3BucketBrowserModel struct {
	Name *string `json:"name,omitempty"`
	Folders *[]string `json:"folders,omitempty"`
}

// NewAmazonS3BucketBrowserModel instantiates a new AmazonS3BucketBrowserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonS3BucketBrowserModel() *AmazonS3BucketBrowserModel {
	this := AmazonS3BucketBrowserModel{}
	return &this
}

// NewAmazonS3BucketBrowserModelWithDefaults instantiates a new AmazonS3BucketBrowserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonS3BucketBrowserModelWithDefaults() *AmazonS3BucketBrowserModel {
	this := AmazonS3BucketBrowserModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AmazonS3BucketBrowserModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3BucketBrowserModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AmazonS3BucketBrowserModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AmazonS3BucketBrowserModel) SetName(v string) {
	o.Name = &v
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *AmazonS3BucketBrowserModel) GetFolders() []string {
	if o == nil || o.Folders == nil {
		var ret []string
		return ret
	}
	return *o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3BucketBrowserModel) GetFoldersOk() (*[]string, bool) {
	if o == nil || o.Folders == nil {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *AmazonS3BucketBrowserModel) HasFolders() bool {
	if o != nil && o.Folders != nil {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []string and assigns it to the Folders field.
func (o *AmazonS3BucketBrowserModel) SetFolders(v []string) {
	o.Folders = &v
}

func (o AmazonS3BucketBrowserModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Folders != nil {
		toSerialize["folders"] = o.Folders
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonS3BucketBrowserModel struct {
	value *AmazonS3BucketBrowserModel
	isSet bool
}

func (v NullableAmazonS3BucketBrowserModel) Get() *AmazonS3BucketBrowserModel {
	return v.value
}

func (v *NullableAmazonS3BucketBrowserModel) Set(val *AmazonS3BucketBrowserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonS3BucketBrowserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonS3BucketBrowserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonS3BucketBrowserModel(val *AmazonS3BucketBrowserModel) *NullableAmazonS3BucketBrowserModel {
	return &NullableAmazonS3BucketBrowserModel{value: val, isSet: true}
}

func (v NullableAmazonS3BucketBrowserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonS3BucketBrowserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


