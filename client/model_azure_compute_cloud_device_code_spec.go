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

// AzureComputeCloudDeviceCodeSpec struct for AzureComputeCloudDeviceCodeSpec
type AzureComputeCloudDeviceCodeSpec struct {
	CloudDeviceCodeSpec
	Region EAzureRegionType `json:"region"`
}

// NewAzureComputeCloudDeviceCodeSpec instantiates a new AzureComputeCloudDeviceCodeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeCloudDeviceCodeSpec(region EAzureRegionType, ) *AzureComputeCloudDeviceCodeSpec {
	this := AzureComputeCloudDeviceCodeSpec{}
	this.Region = region
	return &this
}

// NewAzureComputeCloudDeviceCodeSpecWithDefaults instantiates a new AzureComputeCloudDeviceCodeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeCloudDeviceCodeSpecWithDefaults() *AzureComputeCloudDeviceCodeSpec {
	this := AzureComputeCloudDeviceCodeSpec{}
	return &this
}

// GetRegion returns the Region field value
func (o *AzureComputeCloudDeviceCodeSpec) GetRegion() EAzureRegionType {
	if o == nil  {
		var ret EAzureRegionType
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudDeviceCodeSpec) GetRegionOk() (*EAzureRegionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AzureComputeCloudDeviceCodeSpec) SetRegion(v EAzureRegionType) {
	o.Region = v
}

func (o AzureComputeCloudDeviceCodeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudDeviceCodeSpec, errCloudDeviceCodeSpec := json.Marshal(o.CloudDeviceCodeSpec)
	if errCloudDeviceCodeSpec != nil {
		return []byte{}, errCloudDeviceCodeSpec
	}
	errCloudDeviceCodeSpec = json.Unmarshal([]byte(serializedCloudDeviceCodeSpec), &toSerialize)
	if errCloudDeviceCodeSpec != nil {
		return []byte{}, errCloudDeviceCodeSpec
	}
	if true {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableAzureComputeCloudDeviceCodeSpec struct {
	value *AzureComputeCloudDeviceCodeSpec
	isSet bool
}

func (v NullableAzureComputeCloudDeviceCodeSpec) Get() *AzureComputeCloudDeviceCodeSpec {
	return v.value
}

func (v *NullableAzureComputeCloudDeviceCodeSpec) Set(val *AzureComputeCloudDeviceCodeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeCloudDeviceCodeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeCloudDeviceCodeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeCloudDeviceCodeSpec(val *AzureComputeCloudDeviceCodeSpec) *NullableAzureComputeCloudDeviceCodeSpec {
	return &NullableAzureComputeCloudDeviceCodeSpec{value: val, isSet: true}
}

func (v NullableAzureComputeCloudDeviceCodeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeCloudDeviceCodeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

