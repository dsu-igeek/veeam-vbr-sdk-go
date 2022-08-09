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

// EInstantRecoveryMountState Mount state.
type EInstantRecoveryMountState string

// List of EInstantRecoveryMountState
const (
	EINSTANTRECOVERYMOUNTSTATE_FAILED EInstantRecoveryMountState = "Failed"
	EINSTANTRECOVERYMOUNTSTATE_MOUNTING EInstantRecoveryMountState = "Mounting"
	EINSTANTRECOVERYMOUNTSTATE_MOUNTED EInstantRecoveryMountState = "Mounted"
	EINSTANTRECOVERYMOUNTSTATE_DISMOUNTING EInstantRecoveryMountState = "Dismounting"
)

func (v *EInstantRecoveryMountState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EInstantRecoveryMountState(value)
	for _, existing := range []EInstantRecoveryMountState{ "Failed", "Mounting", "Mounted", "Dismounting",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EInstantRecoveryMountState", value)
}

// Ptr returns reference to EInstantRecoveryMountState value
func (v EInstantRecoveryMountState) Ptr() *EInstantRecoveryMountState {
	return &v
}

type NullableEInstantRecoveryMountState struct {
	value *EInstantRecoveryMountState
	isSet bool
}

func (v NullableEInstantRecoveryMountState) Get() *EInstantRecoveryMountState {
	return v.value
}

func (v *NullableEInstantRecoveryMountState) Set(val *EInstantRecoveryMountState) {
	v.value = val
	v.isSet = true
}

func (v NullableEInstantRecoveryMountState) IsSet() bool {
	return v.isSet
}

func (v *NullableEInstantRecoveryMountState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEInstantRecoveryMountState(val *EInstantRecoveryMountState) *NullableEInstantRecoveryMountState {
	return &NullableEInstantRecoveryMountState{value: val, isSet: true}
}

func (v NullableEInstantRecoveryMountState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEInstantRecoveryMountState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

