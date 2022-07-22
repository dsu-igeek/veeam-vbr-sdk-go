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
	"time"
)

// JobStateModel struct for JobStateModel
type JobStateModel struct {
	// ID of the job.
	Id string `json:"id"`
	// Name of the job.
	Name string `json:"name"`
	Type EJobType `json:"type"`
	// Description of the job.
	Description string `json:"description"`
	Status EJobStatus `json:"status"`
	// Last run of the job.
	LastRun *time.Time `json:"lastRun,omitempty"`
	LastResult ESessionResult `json:"lastResult"`
	// Next run of the job.
	NextRun *time.Time `json:"nextRun,omitempty"`
	Workload EJobWorkload `json:"workload"`
	// ID of the backup repository.
	RepositoryId *string `json:"repositoryId,omitempty"`
	// Name of the backup repository.
	RepositoryName *string `json:"repositoryName,omitempty"`
	// Number of objects processed by the job.
	ObjectsCount int32 `json:"objectsCount"`
	// ID of the last job session.
	SessionId *string `json:"sessionId,omitempty"`
}

// NewJobStateModel instantiates a new JobStateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobStateModel(id string, name string, type_ EJobType, description string, status EJobStatus, lastResult ESessionResult, workload EJobWorkload, objectsCount int32) *JobStateModel {
	this := JobStateModel{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Description = description
	this.Status = status
	this.LastResult = lastResult
	this.Workload = workload
	this.ObjectsCount = objectsCount
	return &this
}

// NewJobStateModelWithDefaults instantiates a new JobStateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobStateModelWithDefaults() *JobStateModel {
	this := JobStateModel{}
	return &this
}

// GetId returns the Id field value
func (o *JobStateModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobStateModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *JobStateModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobStateModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *JobStateModel) GetType() EJobType {
	if o == nil {
		var ret EJobType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetTypeOk() (*EJobType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobStateModel) SetType(v EJobType) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *JobStateModel) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *JobStateModel) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value
func (o *JobStateModel) GetStatus() EJobStatus {
	if o == nil {
		var ret EJobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetStatusOk() (*EJobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *JobStateModel) SetStatus(v EJobStatus) {
	o.Status = v
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *JobStateModel) GetLastRun() time.Time {
	if o == nil || o.LastRun == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetLastRunOk() (*time.Time, bool) {
	if o == nil || o.LastRun == nil {
		return nil, false
	}
	return o.LastRun, true
}

// HasLastRun returns a boolean if a field has been set.
func (o *JobStateModel) HasLastRun() bool {
	if o != nil && o.LastRun != nil {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given time.Time and assigns it to the LastRun field.
func (o *JobStateModel) SetLastRun(v time.Time) {
	o.LastRun = &v
}

// GetLastResult returns the LastResult field value
func (o *JobStateModel) GetLastResult() ESessionResult {
	if o == nil {
		var ret ESessionResult
		return ret
	}

	return o.LastResult
}

// GetLastResultOk returns a tuple with the LastResult field value
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetLastResultOk() (*ESessionResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastResult, true
}

// SetLastResult sets field value
func (o *JobStateModel) SetLastResult(v ESessionResult) {
	o.LastResult = v
}

// GetNextRun returns the NextRun field value if set, zero value otherwise.
func (o *JobStateModel) GetNextRun() time.Time {
	if o == nil || o.NextRun == nil {
		var ret time.Time
		return ret
	}
	return *o.NextRun
}

// GetNextRunOk returns a tuple with the NextRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetNextRunOk() (*time.Time, bool) {
	if o == nil || o.NextRun == nil {
		return nil, false
	}
	return o.NextRun, true
}

// HasNextRun returns a boolean if a field has been set.
func (o *JobStateModel) HasNextRun() bool {
	if o != nil && o.NextRun != nil {
		return true
	}

	return false
}

// SetNextRun gets a reference to the given time.Time and assigns it to the NextRun field.
func (o *JobStateModel) SetNextRun(v time.Time) {
	o.NextRun = &v
}

// GetWorkload returns the Workload field value
func (o *JobStateModel) GetWorkload() EJobWorkload {
	if o == nil {
		var ret EJobWorkload
		return ret
	}

	return o.Workload
}

// GetWorkloadOk returns a tuple with the Workload field value
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetWorkloadOk() (*EJobWorkload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workload, true
}

// SetWorkload sets field value
func (o *JobStateModel) SetWorkload(v EJobWorkload) {
	o.Workload = v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *JobStateModel) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *JobStateModel) HasRepositoryId() bool {
	if o != nil && o.RepositoryId != nil {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *JobStateModel) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetRepositoryName returns the RepositoryName field value if set, zero value otherwise.
func (o *JobStateModel) GetRepositoryName() string {
	if o == nil || o.RepositoryName == nil {
		var ret string
		return ret
	}
	return *o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetRepositoryNameOk() (*string, bool) {
	if o == nil || o.RepositoryName == nil {
		return nil, false
	}
	return o.RepositoryName, true
}

// HasRepositoryName returns a boolean if a field has been set.
func (o *JobStateModel) HasRepositoryName() bool {
	if o != nil && o.RepositoryName != nil {
		return true
	}

	return false
}

// SetRepositoryName gets a reference to the given string and assigns it to the RepositoryName field.
func (o *JobStateModel) SetRepositoryName(v string) {
	o.RepositoryName = &v
}

// GetObjectsCount returns the ObjectsCount field value
func (o *JobStateModel) GetObjectsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ObjectsCount
}

// GetObjectsCountOk returns a tuple with the ObjectsCount field value
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetObjectsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectsCount, true
}

// SetObjectsCount sets field value
func (o *JobStateModel) SetObjectsCount(v int32) {
	o.ObjectsCount = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *JobStateModel) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStateModel) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *JobStateModel) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *JobStateModel) SetSessionId(v string) {
	o.SessionId = &v
}

func (o JobStateModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.LastRun != nil {
		toSerialize["lastRun"] = o.LastRun
	}
	if true {
		toSerialize["lastResult"] = o.LastResult
	}
	if o.NextRun != nil {
		toSerialize["nextRun"] = o.NextRun
	}
	if true {
		toSerialize["workload"] = o.Workload
	}
	if o.RepositoryId != nil {
		toSerialize["repositoryId"] = o.RepositoryId
	}
	if o.RepositoryName != nil {
		toSerialize["repositoryName"] = o.RepositoryName
	}
	if true {
		toSerialize["objectsCount"] = o.ObjectsCount
	}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	return json.Marshal(toSerialize)
}

type NullableJobStateModel struct {
	value *JobStateModel
	isSet bool
}

func (v NullableJobStateModel) Get() *JobStateModel {
	return v.value
}

func (v *NullableJobStateModel) Set(val *JobStateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableJobStateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableJobStateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobStateModel(val *JobStateModel) *NullableJobStateModel {
	return &NullableJobStateModel{value: val, isSet: true}
}

func (v NullableJobStateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobStateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


