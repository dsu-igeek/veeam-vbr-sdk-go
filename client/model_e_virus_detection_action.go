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

// EVirusDetectionAction Action that Veeam Backup & Replication will take if the antivirus software finds a threat.
type EVirusDetectionAction string

// List of EVirusDetectionAction
const (
	EVIRUSDETECTIONACTION_DISABLE_NETWORK EVirusDetectionAction = "DisableNetwork"
	EVIRUSDETECTIONACTION_ABORT_RECOVERY EVirusDetectionAction = "AbortRecovery"
	EVIRUSDETECTIONACTION_IGNORE EVirusDetectionAction = "Ignore"
)

func (v *EVirusDetectionAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EVirusDetectionAction(value)
	for _, existing := range []EVirusDetectionAction{ "DisableNetwork", "AbortRecovery", "Ignore",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EVirusDetectionAction", value)
}

// Ptr returns reference to EVirusDetectionAction value
func (v EVirusDetectionAction) Ptr() *EVirusDetectionAction {
	return &v
}

type NullableEVirusDetectionAction struct {
	value *EVirusDetectionAction
	isSet bool
}

func (v NullableEVirusDetectionAction) Get() *EVirusDetectionAction {
	return v.value
}

func (v *NullableEVirusDetectionAction) Set(val *EVirusDetectionAction) {
	v.value = val
	v.isSet = true
}

func (v NullableEVirusDetectionAction) IsSet() bool {
	return v.isSet
}

func (v *NullableEVirusDetectionAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEVirusDetectionAction(val *EVirusDetectionAction) *NullableEVirusDetectionAction {
	return &NullableEVirusDetectionAction{value: val, isSet: true}
}

func (v NullableEVirusDetectionAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEVirusDetectionAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
