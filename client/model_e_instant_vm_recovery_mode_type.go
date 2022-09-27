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

// EInstantVMRecoveryModeType Restore mode.
type EInstantVMRecoveryModeType string

// List of EInstantVMRecoveryModeType
const (
	EINSTANTVMRECOVERYMODETYPE_ORIGINAL_LOCATION EInstantVMRecoveryModeType = "OriginalLocation"
	EINSTANTVMRECOVERYMODETYPE_CUSTOMIZED EInstantVMRecoveryModeType = "Customized"
)

func (v *EInstantVMRecoveryModeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EInstantVMRecoveryModeType(value)
	for _, existing := range []EInstantVMRecoveryModeType{ "OriginalLocation", "Customized",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EInstantVMRecoveryModeType", value)
}

// Ptr returns reference to EInstantVMRecoveryModeType value
func (v EInstantVMRecoveryModeType) Ptr() *EInstantVMRecoveryModeType {
	return &v
}

type NullableEInstantVMRecoveryModeType struct {
	value *EInstantVMRecoveryModeType
	isSet bool
}

func (v NullableEInstantVMRecoveryModeType) Get() *EInstantVMRecoveryModeType {
	return v.value
}

func (v *NullableEInstantVMRecoveryModeType) Set(val *EInstantVMRecoveryModeType) {
	v.value = val
	v.isSet = true
}

func (v NullableEInstantVMRecoveryModeType) IsSet() bool {
	return v.isSet
}

func (v *NullableEInstantVMRecoveryModeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEInstantVMRecoveryModeType(val *EInstantVMRecoveryModeType) *NullableEInstantVMRecoveryModeType {
	return &NullableEInstantVMRecoveryModeType{value: val, isSet: true}
}

func (v NullableEInstantVMRecoveryModeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEInstantVMRecoveryModeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
