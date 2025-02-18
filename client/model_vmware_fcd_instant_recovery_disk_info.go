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
)

// VmwareFcdInstantRecoveryDiskInfo struct for VmwareFcdInstantRecoveryDiskInfo
type VmwareFcdInstantRecoveryDiskInfo struct {
	// Disk name displayed in the backup.
	NameInBackup string `json:"nameInBackup"`
	// Name of the VMDK file that is stored in the datastore.
	MountedDiskName string `json:"mountedDiskName"`
	// Name under which the disk is registered as an FCD in the vCenter Managed Object Browser.
	RegisteredFcdName string `json:"registeredFcdName"`
	// FCD ID.
	ObjectId string `json:"objectId"`
}

// NewVmwareFcdInstantRecoveryDiskInfo instantiates a new VmwareFcdInstantRecoveryDiskInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareFcdInstantRecoveryDiskInfo(nameInBackup string, mountedDiskName string, registeredFcdName string, objectId string) *VmwareFcdInstantRecoveryDiskInfo {
	this := VmwareFcdInstantRecoveryDiskInfo{}
	this.NameInBackup = nameInBackup
	this.MountedDiskName = mountedDiskName
	this.RegisteredFcdName = registeredFcdName
	this.ObjectId = objectId
	return &this
}

// NewVmwareFcdInstantRecoveryDiskInfoWithDefaults instantiates a new VmwareFcdInstantRecoveryDiskInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareFcdInstantRecoveryDiskInfoWithDefaults() *VmwareFcdInstantRecoveryDiskInfo {
	this := VmwareFcdInstantRecoveryDiskInfo{}
	return &this
}

// GetNameInBackup returns the NameInBackup field value
func (o *VmwareFcdInstantRecoveryDiskInfo) GetNameInBackup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameInBackup
}

// GetNameInBackupOk returns a tuple with the NameInBackup field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryDiskInfo) GetNameInBackupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NameInBackup, true
}

// SetNameInBackup sets field value
func (o *VmwareFcdInstantRecoveryDiskInfo) SetNameInBackup(v string) {
	o.NameInBackup = v
}

// GetMountedDiskName returns the MountedDiskName field value
func (o *VmwareFcdInstantRecoveryDiskInfo) GetMountedDiskName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountedDiskName
}

// GetMountedDiskNameOk returns a tuple with the MountedDiskName field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryDiskInfo) GetMountedDiskNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountedDiskName, true
}

// SetMountedDiskName sets field value
func (o *VmwareFcdInstantRecoveryDiskInfo) SetMountedDiskName(v string) {
	o.MountedDiskName = v
}

// GetRegisteredFcdName returns the RegisteredFcdName field value
func (o *VmwareFcdInstantRecoveryDiskInfo) GetRegisteredFcdName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegisteredFcdName
}

// GetRegisteredFcdNameOk returns a tuple with the RegisteredFcdName field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryDiskInfo) GetRegisteredFcdNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegisteredFcdName, true
}

// SetRegisteredFcdName sets field value
func (o *VmwareFcdInstantRecoveryDiskInfo) SetRegisteredFcdName(v string) {
	o.RegisteredFcdName = v
}

// GetObjectId returns the ObjectId field value
func (o *VmwareFcdInstantRecoveryDiskInfo) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoveryDiskInfo) GetObjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *VmwareFcdInstantRecoveryDiskInfo) SetObjectId(v string) {
	o.ObjectId = v
}

func (o VmwareFcdInstantRecoveryDiskInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nameInBackup"] = o.NameInBackup
	}
	if true {
		toSerialize["mountedDiskName"] = o.MountedDiskName
	}
	if true {
		toSerialize["registeredFcdName"] = o.RegisteredFcdName
	}
	if true {
		toSerialize["objectId"] = o.ObjectId
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareFcdInstantRecoveryDiskInfo struct {
	value *VmwareFcdInstantRecoveryDiskInfo
	isSet bool
}

func (v NullableVmwareFcdInstantRecoveryDiskInfo) Get() *VmwareFcdInstantRecoveryDiskInfo {
	return v.value
}

func (v *NullableVmwareFcdInstantRecoveryDiskInfo) Set(val *VmwareFcdInstantRecoveryDiskInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareFcdInstantRecoveryDiskInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareFcdInstantRecoveryDiskInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareFcdInstantRecoveryDiskInfo(val *VmwareFcdInstantRecoveryDiskInfo) *NullableVmwareFcdInstantRecoveryDiskInfo {
	return &NullableVmwareFcdInstantRecoveryDiskInfo{value: val, isSet: true}
}

func (v NullableVmwareFcdInstantRecoveryDiskInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareFcdInstantRecoveryDiskInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


