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

// EGuestOSType Type of the guest OS.
type EGuestOSType string

// List of EGuestOSType
const (
	EGUESTOSTYPE_WINDOWS EGuestOSType = "Windows"
	EGUESTOSTYPE_LINUX EGuestOSType = "Linux"
)

func (v *EGuestOSType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EGuestOSType(value)
	for _, existing := range []EGuestOSType{ "Windows", "Linux",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EGuestOSType", value)
}

// Ptr returns reference to EGuestOSType value
func (v EGuestOSType) Ptr() *EGuestOSType {
	return &v
}

type NullableEGuestOSType struct {
	value *EGuestOSType
	isSet bool
}

func (v NullableEGuestOSType) Get() *EGuestOSType {
	return v.value
}

func (v *NullableEGuestOSType) Set(val *EGuestOSType) {
	v.value = val
	v.isSet = true
}

func (v NullableEGuestOSType) IsSet() bool {
	return v.isSet
}

func (v *NullableEGuestOSType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEGuestOSType(val *EGuestOSType) *NullableEGuestOSType {
	return &NullableEGuestOSType{value: val, isSet: true}
}

func (v NullableEGuestOSType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEGuestOSType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

