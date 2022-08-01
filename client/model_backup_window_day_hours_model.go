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

// BackupWindowDayHoursModel Hourly scheme for a day.
type BackupWindowDayHoursModel struct {
	Day EDayOfWeek `json:"day"`
	// String of hours in the following format: *1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1* where *1* means enabled, *0* means disabled. 
	Hours string `json:"hours"`
}

// NewBackupWindowDayHoursModel instantiates a new BackupWindowDayHoursModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupWindowDayHoursModel(day EDayOfWeek, hours string, ) *BackupWindowDayHoursModel {
	this := BackupWindowDayHoursModel{}
	this.Day = day
	this.Hours = hours
	return &this
}

// NewBackupWindowDayHoursModelWithDefaults instantiates a new BackupWindowDayHoursModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupWindowDayHoursModelWithDefaults() *BackupWindowDayHoursModel {
	this := BackupWindowDayHoursModel{}
	return &this
}

// GetDay returns the Day field value
func (o *BackupWindowDayHoursModel) GetDay() EDayOfWeek {
	if o == nil  {
		var ret EDayOfWeek
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *BackupWindowDayHoursModel) GetDayOk() (*EDayOfWeek, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *BackupWindowDayHoursModel) SetDay(v EDayOfWeek) {
	o.Day = v
}

// GetHours returns the Hours field value
func (o *BackupWindowDayHoursModel) GetHours() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *BackupWindowDayHoursModel) GetHoursOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hours, true
}

// SetHours sets field value
func (o *BackupWindowDayHoursModel) SetHours(v string) {
	o.Hours = v
}

func (o BackupWindowDayHoursModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["day"] = o.Day
	}
	if true {
		toSerialize["hours"] = o.Hours
	}
	return json.Marshal(toSerialize)
}

type NullableBackupWindowDayHoursModel struct {
	value *BackupWindowDayHoursModel
	isSet bool
}

func (v NullableBackupWindowDayHoursModel) Get() *BackupWindowDayHoursModel {
	return v.value
}

func (v *NullableBackupWindowDayHoursModel) Set(val *BackupWindowDayHoursModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupWindowDayHoursModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupWindowDayHoursModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupWindowDayHoursModel(val *BackupWindowDayHoursModel) *NullableBackupWindowDayHoursModel {
	return &NullableBackupWindowDayHoursModel{value: val, isSet: true}
}

func (v NullableBackupWindowDayHoursModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupWindowDayHoursModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


