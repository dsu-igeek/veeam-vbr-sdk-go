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

// PreferredNetworksModel Preferred networks used for backup and replication traffic.
type PreferredNetworksModel struct {
	// If *true*, prefered networks are enabled.
	IsEnabled bool `json:"isEnabled"`
	// Array of networks.
	Networks []PreferredNetworkModel `json:"networks,omitempty"`
}

// NewPreferredNetworksModel instantiates a new PreferredNetworksModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferredNetworksModel(isEnabled bool) *PreferredNetworksModel {
	this := PreferredNetworksModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewPreferredNetworksModelWithDefaults instantiates a new PreferredNetworksModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferredNetworksModelWithDefaults() *PreferredNetworksModel {
	this := PreferredNetworksModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *PreferredNetworksModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *PreferredNetworksModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *PreferredNetworksModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *PreferredNetworksModel) GetNetworks() []PreferredNetworkModel {
	if o == nil || o.Networks == nil {
		var ret []PreferredNetworkModel
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredNetworksModel) GetNetworksOk() ([]PreferredNetworkModel, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *PreferredNetworksModel) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []PreferredNetworkModel and assigns it to the Networks field.
func (o *PreferredNetworksModel) SetNetworks(v []PreferredNetworkModel) {
	o.Networks = v
}

func (o PreferredNetworksModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	return json.Marshal(toSerialize)
}

type NullablePreferredNetworksModel struct {
	value *PreferredNetworksModel
	isSet bool
}

func (v NullablePreferredNetworksModel) Get() *PreferredNetworksModel {
	return v.value
}

func (v *NullablePreferredNetworksModel) Set(val *PreferredNetworksModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferredNetworksModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferredNetworksModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferredNetworksModel(val *PreferredNetworksModel) *NullablePreferredNetworksModel {
	return &NullablePreferredNetworksModel{value: val, isSet: true}
}

func (v NullablePreferredNetworksModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferredNetworksModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


