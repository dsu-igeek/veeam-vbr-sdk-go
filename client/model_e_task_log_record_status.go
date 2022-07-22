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
	"fmt"
)

// ETaskLogRecordStatus the model 'ETaskLogRecordStatus'
type ETaskLogRecordStatus string

// List of ETaskLogRecordStatus
const (
	ETASKLOGRECORDSTATUS_NONE ETaskLogRecordStatus = "None"
	ETASKLOGRECORDSTATUS_SUCCEEDED ETaskLogRecordStatus = "Succeeded"
	ETASKLOGRECORDSTATUS_WARNING ETaskLogRecordStatus = "Warning"
	ETASKLOGRECORDSTATUS_FAILED ETaskLogRecordStatus = "Failed"
)

// All allowed values of ETaskLogRecordStatus enum
var AllowedETaskLogRecordStatusEnumValues = []ETaskLogRecordStatus{
	"None",
	"Succeeded",
	"Warning",
	"Failed",
}

func (v *ETaskLogRecordStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ETaskLogRecordStatus(value)
	for _, existing := range AllowedETaskLogRecordStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ETaskLogRecordStatus", value)
}

// NewETaskLogRecordStatusFromValue returns a pointer to a valid ETaskLogRecordStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewETaskLogRecordStatusFromValue(v string) (*ETaskLogRecordStatus, error) {
	ev := ETaskLogRecordStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ETaskLogRecordStatus: valid values are %v", v, AllowedETaskLogRecordStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ETaskLogRecordStatus) IsValid() bool {
	for _, existing := range AllowedETaskLogRecordStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ETaskLogRecordStatus value
func (v ETaskLogRecordStatus) Ptr() *ETaskLogRecordStatus {
	return &v
}

type NullableETaskLogRecordStatus struct {
	value *ETaskLogRecordStatus
	isSet bool
}

func (v NullableETaskLogRecordStatus) Get() *ETaskLogRecordStatus {
	return v.value
}

func (v *NullableETaskLogRecordStatus) Set(val *ETaskLogRecordStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableETaskLogRecordStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableETaskLogRecordStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableETaskLogRecordStatus(val *ETaskLogRecordStatus) *NullableETaskLogRecordStatus {
	return &NullableETaskLogRecordStatus{value: val, isSet: true}
}

func (v NullableETaskLogRecordStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableETaskLogRecordStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

