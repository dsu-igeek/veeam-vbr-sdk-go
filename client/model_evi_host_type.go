/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EViHostType Type of the VMware vSphere server.
type EViHostType string

// List of EViHostType
const (
	EVIHOSTTYPE_ESX EViHostType = "ESX"
	EVIHOSTTYPE_ESXI EViHostType = "ESXi"
	EVIHOSTTYPE_VC EViHostType = "VC"
)

func (v *EViHostType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EViHostType(value)
	for _, existing := range []EViHostType{ "ESX", "ESXi", "VC",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EViHostType", value)
}

// Ptr returns reference to EViHostType value
func (v EViHostType) Ptr() *EViHostType {
	return &v
}

type NullableEViHostType struct {
	value *EViHostType
	isSet bool
}

func (v NullableEViHostType) Get() *EViHostType {
	return v.value
}

func (v *NullableEViHostType) Set(val *EViHostType) {
	v.value = val
	v.isSet = true
}

func (v NullableEViHostType) IsSet() bool {
	return v.isSet
}

func (v *NullableEViHostType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEViHostType(val *EViHostType) *NullableEViHostType {
	return &NullableEViHostType{value: val, isSet: true}
}

func (v NullableEViHostType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEViHostType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

