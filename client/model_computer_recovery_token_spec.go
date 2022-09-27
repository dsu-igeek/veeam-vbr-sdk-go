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
	"time"
)

// ComputerRecoveryTokenSpec struct for ComputerRecoveryTokenSpec
type ComputerRecoveryTokenSpec struct {
	BackupIds []string `json:"backupIds"`
	// Date and time the access token expires.
	ExpirationDate time.Time `json:"expirationDate"`
}

// NewComputerRecoveryTokenSpec instantiates a new ComputerRecoveryTokenSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerRecoveryTokenSpec(backupIds []string, expirationDate time.Time, ) *ComputerRecoveryTokenSpec {
	this := ComputerRecoveryTokenSpec{}
	this.BackupIds = backupIds
	this.ExpirationDate = expirationDate
	return &this
}

// NewComputerRecoveryTokenSpecWithDefaults instantiates a new ComputerRecoveryTokenSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerRecoveryTokenSpecWithDefaults() *ComputerRecoveryTokenSpec {
	this := ComputerRecoveryTokenSpec{}
	return &this
}

// GetBackupIds returns the BackupIds field value
func (o *ComputerRecoveryTokenSpec) GetBackupIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.BackupIds
}

// GetBackupIdsOk returns a tuple with the BackupIds field value
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenSpec) GetBackupIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupIds, true
}

// SetBackupIds sets field value
func (o *ComputerRecoveryTokenSpec) SetBackupIds(v []string) {
	o.BackupIds = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *ComputerRecoveryTokenSpec) GetExpirationDate() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *ComputerRecoveryTokenSpec) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *ComputerRecoveryTokenSpec) SetExpirationDate(v time.Time) {
	o.ExpirationDate = v
}

func (o ComputerRecoveryTokenSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupIds"] = o.BackupIds
	}
	if true {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	return json.Marshal(toSerialize)
}

type NullableComputerRecoveryTokenSpec struct {
	value *ComputerRecoveryTokenSpec
	isSet bool
}

func (v NullableComputerRecoveryTokenSpec) Get() *ComputerRecoveryTokenSpec {
	return v.value
}

func (v *NullableComputerRecoveryTokenSpec) Set(val *ComputerRecoveryTokenSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerRecoveryTokenSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerRecoveryTokenSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerRecoveryTokenSpec(val *ComputerRecoveryTokenSpec) *NullableComputerRecoveryTokenSpec {
	return &NullableComputerRecoveryTokenSpec{value: val, isSet: true}
}

func (v NullableComputerRecoveryTokenSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerRecoveryTokenSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


