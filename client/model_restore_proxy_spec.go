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

// RestoreProxySpec Backup proxies for VM data transport.
type RestoreProxySpec struct {
	// If *true*, Veeam Backup & Replication will detect backup proxies that are connected to the source datastore and will automatically assign optimal proxy resources for processing VM data.
	AutoSelection bool `json:"autoSelection"`
	// Array of backup proxy IDs.
	ProxyIds *[]string `json:"proxyIds,omitempty"`
}

// NewRestoreProxySpec instantiates a new RestoreProxySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreProxySpec(autoSelection bool, ) *RestoreProxySpec {
	this := RestoreProxySpec{}
	this.AutoSelection = autoSelection
	return &this
}

// NewRestoreProxySpecWithDefaults instantiates a new RestoreProxySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreProxySpecWithDefaults() *RestoreProxySpec {
	this := RestoreProxySpec{}
	return &this
}

// GetAutoSelection returns the AutoSelection field value
func (o *RestoreProxySpec) GetAutoSelection() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AutoSelection
}

// GetAutoSelectionOk returns a tuple with the AutoSelection field value
// and a boolean to check if the value has been set.
func (o *RestoreProxySpec) GetAutoSelectionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AutoSelection, true
}

// SetAutoSelection sets field value
func (o *RestoreProxySpec) SetAutoSelection(v bool) {
	o.AutoSelection = v
}

// GetProxyIds returns the ProxyIds field value if set, zero value otherwise.
func (o *RestoreProxySpec) GetProxyIds() []string {
	if o == nil || o.ProxyIds == nil {
		var ret []string
		return ret
	}
	return *o.ProxyIds
}

// GetProxyIdsOk returns a tuple with the ProxyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreProxySpec) GetProxyIdsOk() (*[]string, bool) {
	if o == nil || o.ProxyIds == nil {
		return nil, false
	}
	return o.ProxyIds, true
}

// HasProxyIds returns a boolean if a field has been set.
func (o *RestoreProxySpec) HasProxyIds() bool {
	if o != nil && o.ProxyIds != nil {
		return true
	}

	return false
}

// SetProxyIds gets a reference to the given []string and assigns it to the ProxyIds field.
func (o *RestoreProxySpec) SetProxyIds(v []string) {
	o.ProxyIds = &v
}

func (o RestoreProxySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["autoSelection"] = o.AutoSelection
	}
	if o.ProxyIds != nil {
		toSerialize["proxyIds"] = o.ProxyIds
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreProxySpec struct {
	value *RestoreProxySpec
	isSet bool
}

func (v NullableRestoreProxySpec) Get() *RestoreProxySpec {
	return v.value
}

func (v *NullableRestoreProxySpec) Set(val *RestoreProxySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreProxySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreProxySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreProxySpec(val *RestoreProxySpec) *NullableRestoreProxySpec {
	return &NullableRestoreProxySpec{value: val, isSet: true}
}

func (v NullableRestoreProxySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreProxySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

