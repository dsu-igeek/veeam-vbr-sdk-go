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
	"fmt"
)

// EObjectRestorePointOperation the model 'EObjectRestorePointOperation'
type EObjectRestorePointOperation string

// List of EObjectRestorePointOperation
const (
	EOBJECTRESTOREPOINTOPERATION_VMWARE_INSTANT_RECOVERY_FCD EObjectRestorePointOperation = "VmwareInstantRecoveryFcd"
)

func (v *EObjectRestorePointOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EObjectRestorePointOperation(value)
	for _, existing := range []EObjectRestorePointOperation{ "VmwareInstantRecoveryFcd",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EObjectRestorePointOperation", value)
}

// Ptr returns reference to EObjectRestorePointOperation value
func (v EObjectRestorePointOperation) Ptr() *EObjectRestorePointOperation {
	return &v
}

type NullableEObjectRestorePointOperation struct {
	value *EObjectRestorePointOperation
	isSet bool
}

func (v NullableEObjectRestorePointOperation) Get() *EObjectRestorePointOperation {
	return v.value
}

func (v *NullableEObjectRestorePointOperation) Set(val *EObjectRestorePointOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableEObjectRestorePointOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableEObjectRestorePointOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEObjectRestorePointOperation(val *EObjectRestorePointOperation) *NullableEObjectRestorePointOperation {
	return &NullableEObjectRestorePointOperation{value: val, isSet: true}
}

func (v NullableEObjectRestorePointOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEObjectRestorePointOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

