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

// BackupJobSpec struct for BackupJobSpec
type BackupJobSpec struct {
	JobSpec
	// If *true*, the job has a high priority in getting backup infrastructure resources.
	IsHighPriority bool `json:"isHighPriority"`
	VirtualMachines BackupJobVirtualMachinesSpec `json:"virtualMachines"`
	Storage BackupJobStorageModel `json:"storage"`
	GuestProcessing BackupJobGuestProcessingModel `json:"guestProcessing"`
	Schedule BackupScheduleModel `json:"schedule"`
}

// NewBackupJobSpec instantiates a new BackupJobSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobSpec(isHighPriority bool, virtualMachines BackupJobVirtualMachinesSpec, storage BackupJobStorageModel, guestProcessing BackupJobGuestProcessingModel, schedule BackupScheduleModel, ) *BackupJobSpec {
	this := BackupJobSpec{}
	this.IsHighPriority = isHighPriority
	this.VirtualMachines = virtualMachines
	this.Storage = storage
	this.GuestProcessing = guestProcessing
	this.Schedule = schedule
	return &this
}

// NewBackupJobSpecWithDefaults instantiates a new BackupJobSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobSpecWithDefaults() *BackupJobSpec {
	this := BackupJobSpec{}
	var isHighPriority bool = false
	this.IsHighPriority = isHighPriority
	return &this
}

// GetIsHighPriority returns the IsHighPriority field value
func (o *BackupJobSpec) GetIsHighPriority() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsHighPriority
}

// GetIsHighPriorityOk returns a tuple with the IsHighPriority field value
// and a boolean to check if the value has been set.
func (o *BackupJobSpec) GetIsHighPriorityOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsHighPriority, true
}

// SetIsHighPriority sets field value
func (o *BackupJobSpec) SetIsHighPriority(v bool) {
	o.IsHighPriority = v
}

// GetVirtualMachines returns the VirtualMachines field value
func (o *BackupJobSpec) GetVirtualMachines() BackupJobVirtualMachinesSpec {
	if o == nil  {
		var ret BackupJobVirtualMachinesSpec
		return ret
	}

	return o.VirtualMachines
}

// GetVirtualMachinesOk returns a tuple with the VirtualMachines field value
// and a boolean to check if the value has been set.
func (o *BackupJobSpec) GetVirtualMachinesOk() (*BackupJobVirtualMachinesSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualMachines, true
}

// SetVirtualMachines sets field value
func (o *BackupJobSpec) SetVirtualMachines(v BackupJobVirtualMachinesSpec) {
	o.VirtualMachines = v
}

// GetStorage returns the Storage field value
func (o *BackupJobSpec) GetStorage() BackupJobStorageModel {
	if o == nil  {
		var ret BackupJobStorageModel
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *BackupJobSpec) GetStorageOk() (*BackupJobStorageModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *BackupJobSpec) SetStorage(v BackupJobStorageModel) {
	o.Storage = v
}

// GetGuestProcessing returns the GuestProcessing field value
func (o *BackupJobSpec) GetGuestProcessing() BackupJobGuestProcessingModel {
	if o == nil  {
		var ret BackupJobGuestProcessingModel
		return ret
	}

	return o.GuestProcessing
}

// GetGuestProcessingOk returns a tuple with the GuestProcessing field value
// and a boolean to check if the value has been set.
func (o *BackupJobSpec) GetGuestProcessingOk() (*BackupJobGuestProcessingModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GuestProcessing, true
}

// SetGuestProcessing sets field value
func (o *BackupJobSpec) SetGuestProcessing(v BackupJobGuestProcessingModel) {
	o.GuestProcessing = v
}

// GetSchedule returns the Schedule field value
func (o *BackupJobSpec) GetSchedule() BackupScheduleModel {
	if o == nil  {
		var ret BackupScheduleModel
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *BackupJobSpec) GetScheduleOk() (*BackupScheduleModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *BackupJobSpec) SetSchedule(v BackupScheduleModel) {
	o.Schedule = v
}

func (o BackupJobSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedJobSpec, errJobSpec := json.Marshal(o.JobSpec)
	if errJobSpec != nil {
		return []byte{}, errJobSpec
	}
	errJobSpec = json.Unmarshal([]byte(serializedJobSpec), &toSerialize)
	if errJobSpec != nil {
		return []byte{}, errJobSpec
	}
	if true {
		toSerialize["isHighPriority"] = o.IsHighPriority
	}
	if true {
		toSerialize["virtualMachines"] = o.VirtualMachines
	}
	if true {
		toSerialize["storage"] = o.Storage
	}
	if true {
		toSerialize["guestProcessing"] = o.GuestProcessing
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobSpec struct {
	value *BackupJobSpec
	isSet bool
}

func (v NullableBackupJobSpec) Get() *BackupJobSpec {
	return v.value
}

func (v *NullableBackupJobSpec) Set(val *BackupJobSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobSpec(val *BackupJobSpec) *NullableBackupJobSpec {
	return &NullableBackupJobSpec{value: val, isSet: true}
}

func (v NullableBackupJobSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


