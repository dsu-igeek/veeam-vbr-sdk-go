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

// ECredentialsType Credentials type.
type ECredentialsType string

// List of ECredentialsType
const (
	ECREDENTIALSTYPE_STANDARD ECredentialsType = "Standard"
	ECREDENTIALSTYPE_LINUX ECredentialsType = "Linux"
)

func (v *ECredentialsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ECredentialsType(value)
	for _, existing := range []ECredentialsType{ "Standard", "Linux",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ECredentialsType", value)
}

// Ptr returns reference to ECredentialsType value
func (v ECredentialsType) Ptr() *ECredentialsType {
	return &v
}

type NullableECredentialsType struct {
	value *ECredentialsType
	isSet bool
}

func (v NullableECredentialsType) Get() *ECredentialsType {
	return v.value
}

func (v *NullableECredentialsType) Set(val *ECredentialsType) {
	v.value = val
	v.isSet = true
}

func (v NullableECredentialsType) IsSet() bool {
	return v.isSet
}

func (v *NullableECredentialsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECredentialsType(val *ECredentialsType) *NullableECredentialsType {
	return &NullableECredentialsType{value: val, isSet: true}
}

func (v NullableECredentialsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECredentialsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

