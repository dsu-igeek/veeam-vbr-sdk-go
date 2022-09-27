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

// AzureVirtualNetworkBrowserModel struct for AzureVirtualNetworkBrowserModel
type AzureVirtualNetworkBrowserModel struct {
	VirtualNetworkName *string `json:"virtualNetworkName,omitempty"`
	Subnets *[]string `json:"subnets,omitempty"`
}

// NewAzureVirtualNetworkBrowserModel instantiates a new AzureVirtualNetworkBrowserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureVirtualNetworkBrowserModel() *AzureVirtualNetworkBrowserModel {
	this := AzureVirtualNetworkBrowserModel{}
	return &this
}

// NewAzureVirtualNetworkBrowserModelWithDefaults instantiates a new AzureVirtualNetworkBrowserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureVirtualNetworkBrowserModelWithDefaults() *AzureVirtualNetworkBrowserModel {
	this := AzureVirtualNetworkBrowserModel{}
	return &this
}

// GetVirtualNetworkName returns the VirtualNetworkName field value if set, zero value otherwise.
func (o *AzureVirtualNetworkBrowserModel) GetVirtualNetworkName() string {
	if o == nil || o.VirtualNetworkName == nil {
		var ret string
		return ret
	}
	return *o.VirtualNetworkName
}

// GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureVirtualNetworkBrowserModel) GetVirtualNetworkNameOk() (*string, bool) {
	if o == nil || o.VirtualNetworkName == nil {
		return nil, false
	}
	return o.VirtualNetworkName, true
}

// HasVirtualNetworkName returns a boolean if a field has been set.
func (o *AzureVirtualNetworkBrowserModel) HasVirtualNetworkName() bool {
	if o != nil && o.VirtualNetworkName != nil {
		return true
	}

	return false
}

// SetVirtualNetworkName gets a reference to the given string and assigns it to the VirtualNetworkName field.
func (o *AzureVirtualNetworkBrowserModel) SetVirtualNetworkName(v string) {
	o.VirtualNetworkName = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *AzureVirtualNetworkBrowserModel) GetSubnets() []string {
	if o == nil || o.Subnets == nil {
		var ret []string
		return ret
	}
	return *o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureVirtualNetworkBrowserModel) GetSubnetsOk() (*[]string, bool) {
	if o == nil || o.Subnets == nil {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *AzureVirtualNetworkBrowserModel) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *AzureVirtualNetworkBrowserModel) SetSubnets(v []string) {
	o.Subnets = &v
}

func (o AzureVirtualNetworkBrowserModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualNetworkName != nil {
		toSerialize["virtualNetworkName"] = o.VirtualNetworkName
	}
	if o.Subnets != nil {
		toSerialize["subnets"] = o.Subnets
	}
	return json.Marshal(toSerialize)
}

type NullableAzureVirtualNetworkBrowserModel struct {
	value *AzureVirtualNetworkBrowserModel
	isSet bool
}

func (v NullableAzureVirtualNetworkBrowserModel) Get() *AzureVirtualNetworkBrowserModel {
	return v.value
}

func (v *NullableAzureVirtualNetworkBrowserModel) Set(val *AzureVirtualNetworkBrowserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureVirtualNetworkBrowserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureVirtualNetworkBrowserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureVirtualNetworkBrowserModel(val *AzureVirtualNetworkBrowserModel) *NullableAzureVirtualNetworkBrowserModel {
	return &NullableAzureVirtualNetworkBrowserModel{value: val, isSet: true}
}

func (v NullableAzureVirtualNetworkBrowserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureVirtualNetworkBrowserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


