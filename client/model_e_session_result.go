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

// ESessionResult Result of the session.
type ESessionResult string

// List of ESessionResult
const (
	ESESSIONRESULT_NONE ESessionResult = "None"
	ESESSIONRESULT_SUCCESS ESessionResult = "Success"
	ESESSIONRESULT_WARNING ESessionResult = "Warning"
	ESESSIONRESULT_FAILED ESessionResult = "Failed"
)

// All allowed values of ESessionResult enum
var AllowedESessionResultEnumValues = []ESessionResult{
	"None",
	"Success",
	"Warning",
	"Failed",
}

func (v *ESessionResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ESessionResult(value)
	for _, existing := range AllowedESessionResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ESessionResult", value)
}

// NewESessionResultFromValue returns a pointer to a valid ESessionResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewESessionResultFromValue(v string) (*ESessionResult, error) {
	ev := ESessionResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ESessionResult: valid values are %v", v, AllowedESessionResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ESessionResult) IsValid() bool {
	for _, existing := range AllowedESessionResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ESessionResult value
func (v ESessionResult) Ptr() *ESessionResult {
	return &v
}

type NullableESessionResult struct {
	value *ESessionResult
	isSet bool
}

func (v NullableESessionResult) Get() *ESessionResult {
	return v.value
}

func (v *NullableESessionResult) Set(val *ESessionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableESessionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableESessionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESessionResult(val *ESessionResult) *NullableESessionResult {
	return &NullableESessionResult{value: val, isSet: true}
}

func (v NullableESessionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESessionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

