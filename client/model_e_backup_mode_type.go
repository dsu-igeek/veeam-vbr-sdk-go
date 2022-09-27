/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br>Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic.
 *
 * API version: 1.1-rev0
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EBackupModeType Type of the backup method used to create a restore point.
type EBackupModeType string

// List of EBackupModeType
const (
	EBACKUPMODETYPE_FULL EBackupModeType = "Full"
	EBACKUPMODETYPE_INCREMENTAL EBackupModeType = "Incremental"
	EBACKUPMODETYPE_REVERSE_INCREMENTAL EBackupModeType = "ReverseIncremental"
	EBACKUPMODETYPE_TRANSOFRM EBackupModeType = "Transofrm"
	EBACKUPMODETYPE_TRANSFORM_FOREVER_INCREMENTAL EBackupModeType = "TransformForeverIncremental"
)

func (v *EBackupModeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EBackupModeType(value)
	for _, existing := range []EBackupModeType{ "Full", "Incremental", "ReverseIncremental", "Transofrm", "TransformForeverIncremental",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EBackupModeType", value)
}

// Ptr returns reference to EBackupModeType value
func (v EBackupModeType) Ptr() *EBackupModeType {
	return &v
}

type NullableEBackupModeType struct {
	value *EBackupModeType
	isSet bool
}

func (v NullableEBackupModeType) Get() *EBackupModeType {
	return v.value
}

func (v *NullableEBackupModeType) Set(val *EBackupModeType) {
	v.value = val
	v.isSet = true
}

func (v NullableEBackupModeType) IsSet() bool {
	return v.isSet
}

func (v *NullableEBackupModeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEBackupModeType(val *EBackupModeType) *NullableEBackupModeType {
	return &NullableEBackupModeType{value: val, isSet: true}
}

func (v NullableEBackupModeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEBackupModeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

