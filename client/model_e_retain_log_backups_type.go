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

// ERetainLogBackupsType Retention policy for the logs stored in the backup repository.
type ERetainLogBackupsType string

// List of ERetainLogBackupsType
const (
	ERETAINLOGBACKUPSTYPE_UNTIL_BACKUP_DELETED ERetainLogBackupsType = "untilBackupDeleted"
	ERETAINLOGBACKUPSTYPE_KEEP_ONLY_DAYS ERetainLogBackupsType = "KeepOnlyDays"
)

func (v *ERetainLogBackupsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ERetainLogBackupsType(value)
	for _, existing := range []ERetainLogBackupsType{ "untilBackupDeleted", "KeepOnlyDays",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ERetainLogBackupsType", value)
}

// Ptr returns reference to ERetainLogBackupsType value
func (v ERetainLogBackupsType) Ptr() *ERetainLogBackupsType {
	return &v
}

type NullableERetainLogBackupsType struct {
	value *ERetainLogBackupsType
	isSet bool
}

func (v NullableERetainLogBackupsType) Get() *ERetainLogBackupsType {
	return v.value
}

func (v *NullableERetainLogBackupsType) Set(val *ERetainLogBackupsType) {
	v.value = val
	v.isSet = true
}

func (v NullableERetainLogBackupsType) IsSet() bool {
	return v.isSet
}

func (v *NullableERetainLogBackupsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableERetainLogBackupsType(val *ERetainLogBackupsType) *NullableERetainLogBackupsType {
	return &NullableERetainLogBackupsType{value: val, isSet: true}
}

func (v NullableERetainLogBackupsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableERetainLogBackupsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

