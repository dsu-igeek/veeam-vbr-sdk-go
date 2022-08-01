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

// SessionLogResult Log of the session.
type SessionLogResult struct {
	// Total number of records.
	TotalRecords *int32 `json:"totalRecords,omitempty"`
	// Array of log records.
	Records *[]SessionLogRecordModel `json:"records,omitempty"`
}

// NewSessionLogResult instantiates a new SessionLogResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionLogResult() *SessionLogResult {
	this := SessionLogResult{}
	return &this
}

// NewSessionLogResultWithDefaults instantiates a new SessionLogResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionLogResultWithDefaults() *SessionLogResult {
	this := SessionLogResult{}
	return &this
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *SessionLogResult) GetTotalRecords() int32 {
	if o == nil || o.TotalRecords == nil {
		var ret int32
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLogResult) GetTotalRecordsOk() (*int32, bool) {
	if o == nil || o.TotalRecords == nil {
		return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *SessionLogResult) HasTotalRecords() bool {
	if o != nil && o.TotalRecords != nil {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given int32 and assigns it to the TotalRecords field.
func (o *SessionLogResult) SetTotalRecords(v int32) {
	o.TotalRecords = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *SessionLogResult) GetRecords() []SessionLogRecordModel {
	if o == nil || o.Records == nil {
		var ret []SessionLogRecordModel
		return ret
	}
	return *o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLogResult) GetRecordsOk() (*[]SessionLogRecordModel, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *SessionLogResult) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []SessionLogRecordModel and assigns it to the Records field.
func (o *SessionLogResult) SetRecords(v []SessionLogRecordModel) {
	o.Records = &v
}

func (o SessionLogResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalRecords != nil {
		toSerialize["totalRecords"] = o.TotalRecords
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableSessionLogResult struct {
	value *SessionLogResult
	isSet bool
}

func (v NullableSessionLogResult) Get() *SessionLogResult {
	return v.value
}

func (v *NullableSessionLogResult) Set(val *SessionLogResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionLogResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionLogResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionLogResult(val *SessionLogResult) *NullableSessionLogResult {
	return &NullableSessionLogResult{value: val, isSet: true}
}

func (v NullableSessionLogResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionLogResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


