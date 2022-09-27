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

// ERepositoryFiltersOrderColumn Sorts repositories by one of the repository parameters.
type ERepositoryFiltersOrderColumn string

// List of ERepositoryFiltersOrderColumn
const (
	EREPOSITORYFILTERSORDERCOLUMN_NAME ERepositoryFiltersOrderColumn = "Name"
	EREPOSITORYFILTERSORDERCOLUMN_DESCRIPTION ERepositoryFiltersOrderColumn = "Description"
	EREPOSITORYFILTERSORDERCOLUMN_TYPE ERepositoryFiltersOrderColumn = "Type"
	EREPOSITORYFILTERSORDERCOLUMN_HOST ERepositoryFiltersOrderColumn = "Host"
	EREPOSITORYFILTERSORDERCOLUMN_PATH ERepositoryFiltersOrderColumn = "Path"
)

func (v *ERepositoryFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ERepositoryFiltersOrderColumn(value)
	for _, existing := range []ERepositoryFiltersOrderColumn{ "Name", "Description", "Type", "Host", "Path",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ERepositoryFiltersOrderColumn", value)
}

// Ptr returns reference to ERepositoryFiltersOrderColumn value
func (v ERepositoryFiltersOrderColumn) Ptr() *ERepositoryFiltersOrderColumn {
	return &v
}

type NullableERepositoryFiltersOrderColumn struct {
	value *ERepositoryFiltersOrderColumn
	isSet bool
}

func (v NullableERepositoryFiltersOrderColumn) Get() *ERepositoryFiltersOrderColumn {
	return v.value
}

func (v *NullableERepositoryFiltersOrderColumn) Set(val *ERepositoryFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableERepositoryFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableERepositoryFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableERepositoryFiltersOrderColumn(val *ERepositoryFiltersOrderColumn) *NullableERepositoryFiltersOrderColumn {
	return &NullableERepositoryFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableERepositoryFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableERepositoryFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

