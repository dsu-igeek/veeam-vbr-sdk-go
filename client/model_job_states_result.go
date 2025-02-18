/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// JobStatesResult struct for JobStatesResult
type JobStatesResult struct {
	// Array of job states.
	Data []JobStateModel `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// NewJobStatesResult instantiates a new JobStatesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobStatesResult(data []JobStateModel, pagination PaginationResult) *JobStatesResult {
	this := JobStatesResult{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewJobStatesResultWithDefaults instantiates a new JobStatesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobStatesResultWithDefaults() *JobStatesResult {
	this := JobStatesResult{}
	return &this
}

// GetData returns the Data field value
func (o *JobStatesResult) GetData() []JobStateModel {
	if o == nil {
		var ret []JobStateModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *JobStatesResult) GetDataOk() (*[]JobStateModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *JobStatesResult) SetData(v []JobStateModel) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *JobStatesResult) GetPagination() PaginationResult {
	if o == nil {
		var ret PaginationResult
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *JobStatesResult) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *JobStatesResult) SetPagination(v PaginationResult) {
	o.Pagination = v
}

func (o JobStatesResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableJobStatesResult struct {
	value *JobStatesResult
	isSet bool
}

func (v NullableJobStatesResult) Get() *JobStatesResult {
	return v.value
}

func (v *NullableJobStatesResult) Set(val *JobStatesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableJobStatesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableJobStatesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobStatesResult(val *JobStatesResult) *NullableJobStatesResult {
	return &NullableJobStatesResult{value: val, isSet: true}
}

func (v NullableJobStatesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobStatesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


