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

// EProxiesFiltersOrderColumn the model 'EProxiesFiltersOrderColumn'
type EProxiesFiltersOrderColumn string

// List of EProxiesFiltersOrderColumn
const (
	EPROXIESFILTERSORDERCOLUMN_NAME EProxiesFiltersOrderColumn = "Name"
	EPROXIESFILTERSORDERCOLUMN_TYPE EProxiesFiltersOrderColumn = "Type"
	EPROXIESFILTERSORDERCOLUMN_DESCRIPTION EProxiesFiltersOrderColumn = "Description"
)

func (v *EProxiesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EProxiesFiltersOrderColumn(value)
	for _, existing := range []EProxiesFiltersOrderColumn{ "Name", "Type", "Description",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EProxiesFiltersOrderColumn", value)
}

// Ptr returns reference to EProxiesFiltersOrderColumn value
func (v EProxiesFiltersOrderColumn) Ptr() *EProxiesFiltersOrderColumn {
	return &v
}

type NullableEProxiesFiltersOrderColumn struct {
	value *EProxiesFiltersOrderColumn
	isSet bool
}

func (v NullableEProxiesFiltersOrderColumn) Get() *EProxiesFiltersOrderColumn {
	return v.value
}

func (v *NullableEProxiesFiltersOrderColumn) Set(val *EProxiesFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEProxiesFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEProxiesFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEProxiesFiltersOrderColumn(val *EProxiesFiltersOrderColumn) *NullableEProxiesFiltersOrderColumn {
	return &NullableEProxiesFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEProxiesFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEProxiesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

