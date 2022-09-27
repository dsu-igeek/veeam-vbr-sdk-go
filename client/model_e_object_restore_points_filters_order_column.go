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

// EObjectRestorePointsFiltersOrderColumn the model 'EObjectRestorePointsFiltersOrderColumn'
type EObjectRestorePointsFiltersOrderColumn string

// List of EObjectRestorePointsFiltersOrderColumn
const (
	EOBJECTRESTOREPOINTSFILTERSORDERCOLUMN_CREATION_TIME EObjectRestorePointsFiltersOrderColumn = "CreationTime"
	EOBJECTRESTOREPOINTSFILTERSORDERCOLUMN_PLATFORM_ID EObjectRestorePointsFiltersOrderColumn = "PlatformId"
	EOBJECTRESTOREPOINTSFILTERSORDERCOLUMN_BACKUP_ID EObjectRestorePointsFiltersOrderColumn = "BackupId"
)

func (v *EObjectRestorePointsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EObjectRestorePointsFiltersOrderColumn(value)
	for _, existing := range []EObjectRestorePointsFiltersOrderColumn{ "CreationTime", "PlatformId", "BackupId",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EObjectRestorePointsFiltersOrderColumn", value)
}

// Ptr returns reference to EObjectRestorePointsFiltersOrderColumn value
func (v EObjectRestorePointsFiltersOrderColumn) Ptr() *EObjectRestorePointsFiltersOrderColumn {
	return &v
}

type NullableEObjectRestorePointsFiltersOrderColumn struct {
	value *EObjectRestorePointsFiltersOrderColumn
	isSet bool
}

func (v NullableEObjectRestorePointsFiltersOrderColumn) Get() *EObjectRestorePointsFiltersOrderColumn {
	return v.value
}

func (v *NullableEObjectRestorePointsFiltersOrderColumn) Set(val *EObjectRestorePointsFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEObjectRestorePointsFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEObjectRestorePointsFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEObjectRestorePointsFiltersOrderColumn(val *EObjectRestorePointsFiltersOrderColumn) *NullableEObjectRestorePointsFiltersOrderColumn {
	return &NullableEObjectRestorePointsFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEObjectRestorePointsFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEObjectRestorePointsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

