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

// ESessionType Type of the session.
type ESessionType string

// List of ESessionType
const (
	ESESSIONTYPE_INFRASTRUCTURE ESessionType = "Infrastructure"
	ESESSIONTYPE_JOB ESessionType = "Job"
	ESESSIONTYPE_AUTOMATION ESessionType = "Automation"
	ESESSIONTYPE_CONFIGURATION_BACKUP ESessionType = "ConfigurationBackup"
	ESESSIONTYPE_REPOSITORY_MAINTENANCE ESessionType = "RepositoryMaintenance"
	ESESSIONTYPE_REPOSITORY_EVACUATE ESessionType = "RepositoryEvacuate"
	ESESSIONTYPE_INFRASTRUCTURE_ITEM_DELETION ESessionType = "InfrastructureItemDeletion"
)

func (v *ESessionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ESessionType(value)
	for _, existing := range []ESessionType{ "Infrastructure", "Job", "Automation", "ConfigurationBackup", "RepositoryMaintenance", "RepositoryEvacuate", "InfrastructureItemDeletion",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ESessionType", value)
}

// Ptr returns reference to ESessionType value
func (v ESessionType) Ptr() *ESessionType {
	return &v
}

type NullableESessionType struct {
	value *ESessionType
	isSet bool
}

func (v NullableESessionType) Get() *ESessionType {
	return v.value
}

func (v *NullableESessionType) Set(val *ESessionType) {
	v.value = val
	v.isSet = true
}

func (v NullableESessionType) IsSet() bool {
	return v.isSet
}

func (v *NullableESessionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESessionType(val *ESessionType) *NullableESessionType {
	return &NullableESessionType{value: val, isSet: true}
}

func (v NullableESessionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESessionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

