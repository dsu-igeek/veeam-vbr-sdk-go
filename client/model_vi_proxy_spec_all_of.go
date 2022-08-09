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

// ViProxySpecAllOf struct for ViProxySpecAllOf
type ViProxySpecAllOf struct {
	Server ProxyServerSettingsModel `json:"server"`
}

// NewViProxySpecAllOf instantiates a new ViProxySpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViProxySpecAllOf(server ProxyServerSettingsModel, ) *ViProxySpecAllOf {
	this := ViProxySpecAllOf{}
	this.Server = server
	return &this
}

// NewViProxySpecAllOfWithDefaults instantiates a new ViProxySpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViProxySpecAllOfWithDefaults() *ViProxySpecAllOf {
	this := ViProxySpecAllOf{}
	return &this
}

// GetServer returns the Server field value
func (o *ViProxySpecAllOf) GetServer() ProxyServerSettingsModel {
	if o == nil  {
		var ret ProxyServerSettingsModel
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *ViProxySpecAllOf) GetServerOk() (*ProxyServerSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *ViProxySpecAllOf) SetServer(v ProxyServerSettingsModel) {
	o.Server = v
}

func (o ViProxySpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableViProxySpecAllOf struct {
	value *ViProxySpecAllOf
	isSet bool
}

func (v NullableViProxySpecAllOf) Get() *ViProxySpecAllOf {
	return v.value
}

func (v *NullableViProxySpecAllOf) Set(val *ViProxySpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableViProxySpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableViProxySpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViProxySpecAllOf(val *ViProxySpecAllOf) *NullableViProxySpecAllOf {
	return &NullableViProxySpecAllOf{value: val, isSet: true}
}

func (v NullableViProxySpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViProxySpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


