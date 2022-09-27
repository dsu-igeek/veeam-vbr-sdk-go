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
)

// VmwareFcdInstantRecoveryMount struct for VmwareFcdInstantRecoveryMount
type VmwareFcdInstantRecoveryMount struct {
	// Mount ID.
	Id string `json:"id"`
	// ID of the restore session. Use the ID to track the progress. For details, see [Get Session](#operation/GetSession).
	SessionId string `json:"sessionId"`
	State EInstantRecoveryMountState `json:"state"`
	Spec VmwareFcdInstantRecoverySpec `json:"spec"`
	// Error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Array of mounted disks.
	MountedDisks *[]VmwareFcdInstantRecoveryDiskInfo `json:"mountedDisks,omitempty"`
}

// NewVmwareFcdInstantRecoveryMount instantiates a new VmwareFcdInstantRecoveryMount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareFcdInstantRecoveryMount(id string, sessionId string, state EInstantRecoveryMountState, spec VmwareFcdInstantRecoverySpec, ) *VmwareFcdInstantRecoveryMount {
	this := VmwareFcdInstantRecoveryMount{}
	this.Id = id
	this.SessionId = sessionId
	this.State = state
	this.Spec = spec
	return &this
}

// NewVmwareFcdInstantRecoveryMountWithDefaults instantiates a new VmwareFcdInstantRecoveryMount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareFcdInstantRecoveryMountWithDefaults() *VmwareFcdInstantRecoveryMount {
	this := VmwareFcdInstantRecoveryMount{}
	return &this
}

// GetId returns the Id field value
func (o *VmwareFcdInstantRecoveryMount) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMount) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VmwareFcdInstantRecoveryMount) SetId(v string) {
	o.Id = v
}

// GetSessionId returns the SessionId field value
func (o *VmwareFcdInstantRecoveryMount) GetSessionId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMount) GetSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *VmwareFcdInstantRecoveryMount) SetSessionId(v string) {
	o.SessionId = v
}

// GetState returns the State field value
func (o *VmwareFcdInstantRecoveryMount) GetState() EInstantRecoveryMountState {
	if o == nil  {
		var ret EInstantRecoveryMountState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMount) GetStateOk() (*EInstantRecoveryMountState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VmwareFcdInstantRecoveryMount) SetState(v EInstantRecoveryMountState) {
	o.State = v
}

// GetSpec returns the Spec field value
func (o *VmwareFcdInstantRecoveryMount) GetSpec() VmwareFcdInstantRecoverySpec {
	if o == nil  {
		var ret VmwareFcdInstantRecoverySpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMount) GetSpecOk() (*VmwareFcdInstantRecoverySpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *VmwareFcdInstantRecoveryMount) SetSpec(v VmwareFcdInstantRecoverySpec) {
	o.Spec = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoveryMount) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMount) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoveryMount) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *VmwareFcdInstantRecoveryMount) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetMountedDisks returns the MountedDisks field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoveryMount) GetMountedDisks() []VmwareFcdInstantRecoveryDiskInfo {
	if o == nil || o.MountedDisks == nil {
		var ret []VmwareFcdInstantRecoveryDiskInfo
		return ret
	}
	return *o.MountedDisks
}

// GetMountedDisksOk returns a tuple with the MountedDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryMount) GetMountedDisksOk() (*[]VmwareFcdInstantRecoveryDiskInfo, bool) {
	if o == nil || o.MountedDisks == nil {
		return nil, false
	}
	return o.MountedDisks, true
}

// HasMountedDisks returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoveryMount) HasMountedDisks() bool {
	if o != nil && o.MountedDisks != nil {
		return true
	}

	return false
}

// SetMountedDisks gets a reference to the given []VmwareFcdInstantRecoveryDiskInfo and assigns it to the MountedDisks field.
func (o *VmwareFcdInstantRecoveryMount) SetMountedDisks(v []VmwareFcdInstantRecoveryDiskInfo) {
	o.MountedDisks = &v
}

func (o VmwareFcdInstantRecoveryMount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["sessionId"] = o.SessionId
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["spec"] = o.Spec
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.MountedDisks != nil {
		toSerialize["mountedDisks"] = o.MountedDisks
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareFcdInstantRecoveryMount struct {
	value *VmwareFcdInstantRecoveryMount
	isSet bool
}

func (v NullableVmwareFcdInstantRecoveryMount) Get() *VmwareFcdInstantRecoveryMount {
	return v.value
}

func (v *NullableVmwareFcdInstantRecoveryMount) Set(val *VmwareFcdInstantRecoveryMount) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareFcdInstantRecoveryMount) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareFcdInstantRecoveryMount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareFcdInstantRecoveryMount(val *VmwareFcdInstantRecoveryMount) *NullableVmwareFcdInstantRecoveryMount {
	return &NullableVmwareFcdInstantRecoveryMount{value: val, isSet: true}
}

func (v NullableVmwareFcdInstantRecoveryMount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareFcdInstantRecoveryMount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


