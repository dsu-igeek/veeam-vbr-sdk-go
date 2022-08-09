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

// ESessionState State of the session.
type ESessionState string

// List of ESessionState
const (
	ESESSIONSTATE_STOPPED ESessionState = "Stopped"
	ESESSIONSTATE_STARTING ESessionState = "Starting"
	ESESSIONSTATE_STOPPING ESessionState = "Stopping"
	ESESSIONSTATE_WORKING ESessionState = "Working"
	ESESSIONSTATE_PAUSING ESessionState = "Pausing"
	ESESSIONSTATE_RESUMING ESessionState = "Resuming"
	ESESSIONSTATE_WAITING_TAPE ESessionState = "WaitingTape"
	ESESSIONSTATE_IDLE ESessionState = "Idle"
	ESESSIONSTATE_POSTPROCESSING ESessionState = "Postprocessing"
	ESESSIONSTATE_WAITING_REPOSITORY ESessionState = "WaitingRepository"
	ESESSIONSTATE_WAITING_SLOT ESessionState = "WaitingSlot"
)

func (v *ESessionState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ESessionState(value)
	for _, existing := range []ESessionState{ "Stopped", "Starting", "Stopping", "Working", "Pausing", "Resuming", "WaitingTape", "Idle", "Postprocessing", "WaitingRepository", "WaitingSlot",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ESessionState", value)
}

// Ptr returns reference to ESessionState value
func (v ESessionState) Ptr() *ESessionState {
	return &v
}

type NullableESessionState struct {
	value *ESessionState
	isSet bool
}

func (v NullableESessionState) Get() *ESessionState {
	return v.value
}

func (v *NullableESessionState) Set(val *ESessionState) {
	v.value = val
	v.isSet = true
}

func (v NullableESessionState) IsSet() bool {
	return v.isSet
}

func (v *NullableESessionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESessionState(val *ESessionState) *NullableESessionState {
	return &NullableESessionState{value: val, isSet: true}
}

func (v NullableESessionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESessionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

