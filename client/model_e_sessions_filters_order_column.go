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

// ESessionsFiltersOrderColumn the model 'ESessionsFiltersOrderColumn'
type ESessionsFiltersOrderColumn string

// List of ESessionsFiltersOrderColumn
const (
	ESESSIONSFILTERSORDERCOLUMN_NAME ESessionsFiltersOrderColumn = "Name"
	ESESSIONSFILTERSORDERCOLUMN_SESSION_TYPE ESessionsFiltersOrderColumn = "SessionType"
	ESESSIONSFILTERSORDERCOLUMN_CREATION_TIME ESessionsFiltersOrderColumn = "CreationTime"
	ESESSIONSFILTERSORDERCOLUMN_END_TIME ESessionsFiltersOrderColumn = "EndTime"
	ESESSIONSFILTERSORDERCOLUMN_STATE ESessionsFiltersOrderColumn = "State"
	ESESSIONSFILTERSORDERCOLUMN_RESULT ESessionsFiltersOrderColumn = "Result"
)

func (v *ESessionsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ESessionsFiltersOrderColumn(value)
	for _, existing := range []ESessionsFiltersOrderColumn{ "Name", "SessionType", "CreationTime", "EndTime", "State", "Result",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ESessionsFiltersOrderColumn", value)
}

// Ptr returns reference to ESessionsFiltersOrderColumn value
func (v ESessionsFiltersOrderColumn) Ptr() *ESessionsFiltersOrderColumn {
	return &v
}

type NullableESessionsFiltersOrderColumn struct {
	value *ESessionsFiltersOrderColumn
	isSet bool
}

func (v NullableESessionsFiltersOrderColumn) Get() *ESessionsFiltersOrderColumn {
	return v.value
}

func (v *NullableESessionsFiltersOrderColumn) Set(val *ESessionsFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableESessionsFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableESessionsFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESessionsFiltersOrderColumn(val *ESessionsFiltersOrderColumn) *NullableESessionsFiltersOrderColumn {
	return &NullableESessionsFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableESessionsFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESessionsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

