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

// EProxyType Type of the backup proxy.
type EProxyType string

// List of EProxyType
const (
	EPROXYTYPE_VI_PROXY EProxyType = "ViProxy"
)

func (v *EProxyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EProxyType(value)
	for _, existing := range []EProxyType{ "ViProxy",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EProxyType", value)
}

// Ptr returns reference to EProxyType value
func (v EProxyType) Ptr() *EProxyType {
	return &v
}

type NullableEProxyType struct {
	value *EProxyType
	isSet bool
}

func (v NullableEProxyType) Get() *EProxyType {
	return v.value
}

func (v *NullableEProxyType) Set(val *EProxyType) {
	v.value = val
	v.isSet = true
}

func (v NullableEProxyType) IsSet() bool {
	return v.isSet
}

func (v *NullableEProxyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEProxyType(val *EProxyType) *NullableEProxyType {
	return &NullableEProxyType{value: val, isSet: true}
}

func (v NullableEProxyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEProxyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

