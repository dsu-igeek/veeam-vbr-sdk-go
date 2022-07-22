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

// EJobWorkload Workload which the job must process.
type EJobWorkload string

// List of EJobWorkload
const (
	EJOBWORKLOAD_APPLICATION EJobWorkload = "application"
	EJOBWORKLOAD_CLOUD_VM EJobWorkload = "cloudVm"
	EJOBWORKLOAD_FILE EJobWorkload = "file"
	EJOBWORKLOAD_SERVER EJobWorkload = "server"
	EJOBWORKLOAD_WORKSTATION EJobWorkload = "workstation"
	EJOBWORKLOAD_VM EJobWorkload = "vm"
)

// All allowed values of EJobWorkload enum
var AllowedEJobWorkloadEnumValues = []EJobWorkload{
	"application",
	"cloudVm",
	"file",
	"server",
	"workstation",
	"vm",
}

func (v *EJobWorkload) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EJobWorkload(value)
	for _, existing := range AllowedEJobWorkloadEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EJobWorkload", value)
}

// NewEJobWorkloadFromValue returns a pointer to a valid EJobWorkload
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEJobWorkloadFromValue(v string) (*EJobWorkload, error) {
	ev := EJobWorkload(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EJobWorkload: valid values are %v", v, AllowedEJobWorkloadEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EJobWorkload) IsValid() bool {
	for _, existing := range AllowedEJobWorkloadEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EJobWorkload value
func (v EJobWorkload) Ptr() *EJobWorkload {
	return &v
}

type NullableEJobWorkload struct {
	value *EJobWorkload
	isSet bool
}

func (v NullableEJobWorkload) Get() *EJobWorkload {
	return v.value
}

func (v *NullableEJobWorkload) Set(val *EJobWorkload) {
	v.value = val
	v.isSet = true
}

func (v NullableEJobWorkload) IsSet() bool {
	return v.isSet
}

func (v *NullableEJobWorkload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEJobWorkload(val *EJobWorkload) *NullableEJobWorkload {
	return &NullableEJobWorkload{value: val, isSet: true}
}

func (v NullableEJobWorkload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEJobWorkload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

