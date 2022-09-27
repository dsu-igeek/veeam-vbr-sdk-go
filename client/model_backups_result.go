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

// BackupsResult struct for BackupsResult
type BackupsResult struct {
	// Array of backups.
	Data []BackupModel `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// NewBackupsResult instantiates a new BackupsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupsResult(data []BackupModel, pagination PaginationResult, ) *BackupsResult {
	this := BackupsResult{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewBackupsResultWithDefaults instantiates a new BackupsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupsResultWithDefaults() *BackupsResult {
	this := BackupsResult{}
	return &this
}

// GetData returns the Data field value
func (o *BackupsResult) GetData() []BackupModel {
	if o == nil  {
		var ret []BackupModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BackupsResult) GetDataOk() (*[]BackupModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BackupsResult) SetData(v []BackupModel) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *BackupsResult) GetPagination() PaginationResult {
	if o == nil  {
		var ret PaginationResult
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *BackupsResult) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *BackupsResult) SetPagination(v PaginationResult) {
	o.Pagination = v
}

func (o BackupsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableBackupsResult struct {
	value *BackupsResult
	isSet bool
}

func (v NullableBackupsResult) Get() *BackupsResult {
	return v.value
}

func (v *NullableBackupsResult) Set(val *BackupsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupsResult(val *BackupsResult) *NullableBackupsResult {
	return &NullableBackupsResult{value: val, isSet: true}
}

func (v NullableBackupsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


