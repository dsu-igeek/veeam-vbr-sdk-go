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

// GoogleCloudStorageBrowserFilters struct for GoogleCloudStorageBrowserFilters
type GoogleCloudStorageBrowserFilters struct {
	RegionId string `json:"regionId"`
	BucketName *string `json:"bucketName,omitempty"`
}

// NewGoogleCloudStorageBrowserFilters instantiates a new GoogleCloudStorageBrowserFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudStorageBrowserFilters(regionId string, ) *GoogleCloudStorageBrowserFilters {
	this := GoogleCloudStorageBrowserFilters{}
	this.RegionId = regionId
	return &this
}

// NewGoogleCloudStorageBrowserFiltersWithDefaults instantiates a new GoogleCloudStorageBrowserFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudStorageBrowserFiltersWithDefaults() *GoogleCloudStorageBrowserFilters {
	this := GoogleCloudStorageBrowserFilters{}
	return &this
}

// GetRegionId returns the RegionId field value
func (o *GoogleCloudStorageBrowserFilters) GetRegionId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBrowserFilters) GetRegionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *GoogleCloudStorageBrowserFilters) SetRegionId(v string) {
	o.RegionId = v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *GoogleCloudStorageBrowserFilters) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBrowserFilters) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *GoogleCloudStorageBrowserFilters) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *GoogleCloudStorageBrowserFilters) SetBucketName(v string) {
	o.BucketName = &v
}

func (o GoogleCloudStorageBrowserFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if o.BucketName != nil {
		toSerialize["bucketName"] = o.BucketName
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleCloudStorageBrowserFilters struct {
	value *GoogleCloudStorageBrowserFilters
	isSet bool
}

func (v NullableGoogleCloudStorageBrowserFilters) Get() *GoogleCloudStorageBrowserFilters {
	return v.value
}

func (v *NullableGoogleCloudStorageBrowserFilters) Set(val *GoogleCloudStorageBrowserFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudStorageBrowserFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudStorageBrowserFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudStorageBrowserFilters(val *GoogleCloudStorageBrowserFilters) *NullableGoogleCloudStorageBrowserFilters {
	return &NullableGoogleCloudStorageBrowserFilters{value: val, isSet: true}
}

func (v NullableGoogleCloudStorageBrowserFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudStorageBrowserFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

