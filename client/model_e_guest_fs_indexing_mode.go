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

// EGuestFSIndexingMode Indexing mode.
type EGuestFSIndexingMode string

// List of EGuestFSIndexingMode
const (
	EGUESTFSINDEXINGMODE_DISABLE EGuestFSIndexingMode = "disable"
	EGUESTFSINDEXINGMODE_INDEX_ALL EGuestFSIndexingMode = "indexAll"
	EGUESTFSINDEXINGMODE_INDEX_ALL_EXCEPT EGuestFSIndexingMode = "indexAllExcept"
	EGUESTFSINDEXINGMODE_INDEX_ONLY EGuestFSIndexingMode = "indexOnly"
)

// All allowed values of EGuestFSIndexingMode enum
var AllowedEGuestFSIndexingModeEnumValues = []EGuestFSIndexingMode{
	"disable",
	"indexAll",
	"indexAllExcept",
	"indexOnly",
}

func (v *EGuestFSIndexingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EGuestFSIndexingMode(value)
	for _, existing := range AllowedEGuestFSIndexingModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EGuestFSIndexingMode", value)
}

// NewEGuestFSIndexingModeFromValue returns a pointer to a valid EGuestFSIndexingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEGuestFSIndexingModeFromValue(v string) (*EGuestFSIndexingMode, error) {
	ev := EGuestFSIndexingMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EGuestFSIndexingMode: valid values are %v", v, AllowedEGuestFSIndexingModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EGuestFSIndexingMode) IsValid() bool {
	for _, existing := range AllowedEGuestFSIndexingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EGuestFSIndexingMode value
func (v EGuestFSIndexingMode) Ptr() *EGuestFSIndexingMode {
	return &v
}

type NullableEGuestFSIndexingMode struct {
	value *EGuestFSIndexingMode
	isSet bool
}

func (v NullableEGuestFSIndexingMode) Get() *EGuestFSIndexingMode {
	return v.value
}

func (v *NullableEGuestFSIndexingMode) Set(val *EGuestFSIndexingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEGuestFSIndexingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEGuestFSIndexingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEGuestFSIndexingMode(val *EGuestFSIndexingMode) *NullableEGuestFSIndexingMode {
	return &NullableEGuestFSIndexingMode{value: val, isSet: true}
}

func (v NullableEGuestFSIndexingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEGuestFSIndexingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

