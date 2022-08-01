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

// ScheduleAfterThisJobModel Job chaining options.
type ScheduleAfterThisJobModel struct {
	// If *true*, job chaining is enabled.
	IsEnabled bool `json:"isEnabled"`
	// Name of the preceding job.
	JobName *string `json:"jobName,omitempty"`
}

// NewScheduleAfterThisJobModel instantiates a new ScheduleAfterThisJobModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleAfterThisJobModel(isEnabled bool, ) *ScheduleAfterThisJobModel {
	this := ScheduleAfterThisJobModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewScheduleAfterThisJobModelWithDefaults instantiates a new ScheduleAfterThisJobModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleAfterThisJobModelWithDefaults() *ScheduleAfterThisJobModel {
	this := ScheduleAfterThisJobModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *ScheduleAfterThisJobModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *ScheduleAfterThisJobModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *ScheduleAfterThisJobModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetJobName returns the JobName field value if set, zero value otherwise.
func (o *ScheduleAfterThisJobModel) GetJobName() string {
	if o == nil || o.JobName == nil {
		var ret string
		return ret
	}
	return *o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleAfterThisJobModel) GetJobNameOk() (*string, bool) {
	if o == nil || o.JobName == nil {
		return nil, false
	}
	return o.JobName, true
}

// HasJobName returns a boolean if a field has been set.
func (o *ScheduleAfterThisJobModel) HasJobName() bool {
	if o != nil && o.JobName != nil {
		return true
	}

	return false
}

// SetJobName gets a reference to the given string and assigns it to the JobName field.
func (o *ScheduleAfterThisJobModel) SetJobName(v string) {
	o.JobName = &v
}

func (o ScheduleAfterThisJobModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.JobName != nil {
		toSerialize["jobName"] = o.JobName
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleAfterThisJobModel struct {
	value *ScheduleAfterThisJobModel
	isSet bool
}

func (v NullableScheduleAfterThisJobModel) Get() *ScheduleAfterThisJobModel {
	return v.value
}

func (v *NullableScheduleAfterThisJobModel) Set(val *ScheduleAfterThisJobModel) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleAfterThisJobModel) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleAfterThisJobModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleAfterThisJobModel(val *ScheduleAfterThisJobModel) *NullableScheduleAfterThisJobModel {
	return &NullableScheduleAfterThisJobModel{value: val, isSet: true}
}

func (v NullableScheduleAfterThisJobModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleAfterThisJobModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


