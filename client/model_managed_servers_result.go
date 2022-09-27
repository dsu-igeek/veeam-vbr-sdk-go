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

// ManagedServersResult struct for ManagedServersResult
type ManagedServersResult struct {
	// Array of managed servers.
	Data []ManagedServerModel `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// NewManagedServersResult instantiates a new ManagedServersResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedServersResult(data []ManagedServerModel, pagination PaginationResult, ) *ManagedServersResult {
	this := ManagedServersResult{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewManagedServersResultWithDefaults instantiates a new ManagedServersResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedServersResultWithDefaults() *ManagedServersResult {
	this := ManagedServersResult{}
	return &this
}

// GetData returns the Data field value
func (o *ManagedServersResult) GetData() []ManagedServerModel {
	if o == nil  {
		var ret []ManagedServerModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManagedServersResult) GetDataOk() (*[]ManagedServerModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ManagedServersResult) SetData(v []ManagedServerModel) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *ManagedServersResult) GetPagination() PaginationResult {
	if o == nil  {
		var ret PaginationResult
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ManagedServersResult) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ManagedServersResult) SetPagination(v PaginationResult) {
	o.Pagination = v
}

func (o ManagedServersResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableManagedServersResult struct {
	value *ManagedServersResult
	isSet bool
}

func (v NullableManagedServersResult) Get() *ManagedServersResult {
	return v.value
}

func (v *NullableManagedServersResult) Set(val *ManagedServersResult) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedServersResult) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedServersResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedServersResult(val *ManagedServersResult) *NullableManagedServersResult {
	return &NullableManagedServersResult{value: val, isSet: true}
}

func (v NullableManagedServersResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedServersResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


