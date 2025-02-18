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

// JobImportSpecCollection struct for JobImportSpecCollection
type JobImportSpecCollection struct {
	// Array of jobs.
	Jobs []JobImportSpec `json:"jobs"`
}

// NewJobImportSpecCollection instantiates a new JobImportSpecCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobImportSpecCollection(jobs []JobImportSpec) *JobImportSpecCollection {
	this := JobImportSpecCollection{}
	this.Jobs = jobs
	return &this
}

// NewJobImportSpecCollectionWithDefaults instantiates a new JobImportSpecCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobImportSpecCollectionWithDefaults() *JobImportSpecCollection {
	this := JobImportSpecCollection{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *JobImportSpecCollection) GetJobs() []JobImportSpec {
	if o == nil {
		var ret []JobImportSpec
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *JobImportSpecCollection) GetJobsOk() (*[]JobImportSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Jobs, true
}

// SetJobs sets field value
func (o *JobImportSpecCollection) SetJobs(v []JobImportSpec) {
	o.Jobs = v
}

func (o JobImportSpecCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobs"] = o.Jobs
	}
	return json.Marshal(toSerialize)
}

type NullableJobImportSpecCollection struct {
	value *JobImportSpecCollection
	isSet bool
}

func (v NullableJobImportSpecCollection) Get() *JobImportSpecCollection {
	return v.value
}

func (v *NullableJobImportSpecCollection) Set(val *JobImportSpecCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableJobImportSpecCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableJobImportSpecCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobImportSpecCollection(val *JobImportSpecCollection) *NullableJobImportSpecCollection {
	return &NullableJobImportSpecCollection{value: val, isSet: true}
}

func (v NullableJobImportSpecCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobImportSpecCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


