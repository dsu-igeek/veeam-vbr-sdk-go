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

// ProxyDatastoreModel struct for ProxyDatastoreModel
type ProxyDatastoreModel struct {
	Datastore *VmwareObjectModel `json:"datastore,omitempty"`
	// Number of VMs.
	VmCount *int32 `json:"vmCount,omitempty"`
}

// NewProxyDatastoreModel instantiates a new ProxyDatastoreModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyDatastoreModel() *ProxyDatastoreModel {
	this := ProxyDatastoreModel{}
	return &this
}

// NewProxyDatastoreModelWithDefaults instantiates a new ProxyDatastoreModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyDatastoreModelWithDefaults() *ProxyDatastoreModel {
	this := ProxyDatastoreModel{}
	return &this
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *ProxyDatastoreModel) GetDatastore() VmwareObjectModel {
	if o == nil || o.Datastore == nil {
		var ret VmwareObjectModel
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyDatastoreModel) GetDatastoreOk() (*VmwareObjectModel, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *ProxyDatastoreModel) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given VmwareObjectModel and assigns it to the Datastore field.
func (o *ProxyDatastoreModel) SetDatastore(v VmwareObjectModel) {
	o.Datastore = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *ProxyDatastoreModel) GetVmCount() int32 {
	if o == nil || o.VmCount == nil {
		var ret int32
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyDatastoreModel) GetVmCountOk() (*int32, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *ProxyDatastoreModel) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int32 and assigns it to the VmCount field.
func (o *ProxyDatastoreModel) SetVmCount(v int32) {
	o.VmCount = &v
}

func (o ProxyDatastoreModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datastore != nil {
		toSerialize["datastore"] = o.Datastore
	}
	if o.VmCount != nil {
		toSerialize["vmCount"] = o.VmCount
	}
	return json.Marshal(toSerialize)
}

type NullableProxyDatastoreModel struct {
	value *ProxyDatastoreModel
	isSet bool
}

func (v NullableProxyDatastoreModel) Get() *ProxyDatastoreModel {
	return v.value
}

func (v *NullableProxyDatastoreModel) Set(val *ProxyDatastoreModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyDatastoreModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyDatastoreModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyDatastoreModel(val *ProxyDatastoreModel) *NullableProxyDatastoreModel {
	return &NullableProxyDatastoreModel{value: val, isSet: true}
}

func (v NullableProxyDatastoreModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyDatastoreModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


