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

// AzureDataBoxBrowserModelAllOf struct for AzureDataBoxBrowserModelAllOf
type AzureDataBoxBrowserModelAllOf struct {
	HostId *string `json:"hostId,omitempty"`
	Containers *[]AzureDataBoxContainerBrowserModel `json:"containers,omitempty"`
}

// NewAzureDataBoxBrowserModelAllOf instantiates a new AzureDataBoxBrowserModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureDataBoxBrowserModelAllOf() *AzureDataBoxBrowserModelAllOf {
	this := AzureDataBoxBrowserModelAllOf{}
	return &this
}

// NewAzureDataBoxBrowserModelAllOfWithDefaults instantiates a new AzureDataBoxBrowserModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureDataBoxBrowserModelAllOfWithDefaults() *AzureDataBoxBrowserModelAllOf {
	this := AzureDataBoxBrowserModelAllOf{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *AzureDataBoxBrowserModelAllOf) GetHostId() string {
	if o == nil || o.HostId == nil {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureDataBoxBrowserModelAllOf) GetHostIdOk() (*string, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *AzureDataBoxBrowserModelAllOf) HasHostId() bool {
	if o != nil && o.HostId != nil {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *AzureDataBoxBrowserModelAllOf) SetHostId(v string) {
	o.HostId = &v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *AzureDataBoxBrowserModelAllOf) GetContainers() []AzureDataBoxContainerBrowserModel {
	if o == nil || o.Containers == nil {
		var ret []AzureDataBoxContainerBrowserModel
		return ret
	}
	return *o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureDataBoxBrowserModelAllOf) GetContainersOk() (*[]AzureDataBoxContainerBrowserModel, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *AzureDataBoxBrowserModelAllOf) HasContainers() bool {
	if o != nil && o.Containers != nil {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []AzureDataBoxContainerBrowserModel and assigns it to the Containers field.
func (o *AzureDataBoxBrowserModelAllOf) SetContainers(v []AzureDataBoxContainerBrowserModel) {
	o.Containers = &v
}

func (o AzureDataBoxBrowserModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HostId != nil {
		toSerialize["hostId"] = o.HostId
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	return json.Marshal(toSerialize)
}

type NullableAzureDataBoxBrowserModelAllOf struct {
	value *AzureDataBoxBrowserModelAllOf
	isSet bool
}

func (v NullableAzureDataBoxBrowserModelAllOf) Get() *AzureDataBoxBrowserModelAllOf {
	return v.value
}

func (v *NullableAzureDataBoxBrowserModelAllOf) Set(val *AzureDataBoxBrowserModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureDataBoxBrowserModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureDataBoxBrowserModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureDataBoxBrowserModelAllOf(val *AzureDataBoxBrowserModelAllOf) *NullableAzureDataBoxBrowserModelAllOf {
	return &NullableAzureDataBoxBrowserModelAllOf{value: val, isSet: true}
}

func (v NullableAzureDataBoxBrowserModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureDataBoxBrowserModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


