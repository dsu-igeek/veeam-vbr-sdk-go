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

// EBackupsFiltersOrderColumn the model 'EBackupsFiltersOrderColumn'
type EBackupsFiltersOrderColumn string

// List of EBackupsFiltersOrderColumn
const (
	EBACKUPSFILTERSORDERCOLUMN_NAME EBackupsFiltersOrderColumn = "Name"
	EBACKUPSFILTERSORDERCOLUMN_CREATION_TIME EBackupsFiltersOrderColumn = "CreationTime"
	EBACKUPSFILTERSORDERCOLUMN_PLATFORM_ID EBackupsFiltersOrderColumn = "PlatformId"
	EBACKUPSFILTERSORDERCOLUMN_JOB_ID EBackupsFiltersOrderColumn = "JobId"
	EBACKUPSFILTERSORDERCOLUMN_POLICY_TAG EBackupsFiltersOrderColumn = "PolicyTag"
)

func (v *EBackupsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EBackupsFiltersOrderColumn(value)
	for _, existing := range []EBackupsFiltersOrderColumn{ "Name", "CreationTime", "PlatformId", "JobId", "PolicyTag",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EBackupsFiltersOrderColumn", value)
}

// Ptr returns reference to EBackupsFiltersOrderColumn value
func (v EBackupsFiltersOrderColumn) Ptr() *EBackupsFiltersOrderColumn {
	return &v
}

type NullableEBackupsFiltersOrderColumn struct {
	value *EBackupsFiltersOrderColumn
	isSet bool
}

func (v NullableEBackupsFiltersOrderColumn) Get() *EBackupsFiltersOrderColumn {
	return v.value
}

func (v *NullableEBackupsFiltersOrderColumn) Set(val *EBackupsFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEBackupsFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEBackupsFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEBackupsFiltersOrderColumn(val *EBackupsFiltersOrderColumn) *NullableEBackupsFiltersOrderColumn {
	return &NullableEBackupsFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEBackupsFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEBackupsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

