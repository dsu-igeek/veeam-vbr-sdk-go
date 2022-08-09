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
)

// VmwareFcdInstantRecoverySpec struct for VmwareFcdInstantRecoverySpec
type VmwareFcdInstantRecoverySpec struct {
	// ID of the restore point.
	ObjectRestorePointId string `json:"objectRestorePointId"`
	DestinationCluster VmwareObjectModel `json:"destinationCluster"`
	// Array of disks that will be restored.
	DisksMapping []VmwareFcdInstantRecoveryDiskSpec `json:"disksMapping"`
	WriteCache *VmwareFcdWriteCacheSpec `json:"writeCache,omitempty"`
}

// NewVmwareFcdInstantRecoverySpec instantiates a new VmwareFcdInstantRecoverySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareFcdInstantRecoverySpec(objectRestorePointId string, destinationCluster VmwareObjectModel, disksMapping []VmwareFcdInstantRecoveryDiskSpec, ) *VmwareFcdInstantRecoverySpec {
	this := VmwareFcdInstantRecoverySpec{}
	this.ObjectRestorePointId = objectRestorePointId
	this.DestinationCluster = destinationCluster
	this.DisksMapping = disksMapping
	return &this
}

// NewVmwareFcdInstantRecoverySpecWithDefaults instantiates a new VmwareFcdInstantRecoverySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareFcdInstantRecoverySpecWithDefaults() *VmwareFcdInstantRecoverySpec {
	this := VmwareFcdInstantRecoverySpec{}
	return &this
}

// GetObjectRestorePointId returns the ObjectRestorePointId field value
func (o *VmwareFcdInstantRecoverySpec) GetObjectRestorePointId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ObjectRestorePointId
}

// GetObjectRestorePointIdOk returns a tuple with the ObjectRestorePointId field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoverySpec) GetObjectRestorePointIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectRestorePointId, true
}

// SetObjectRestorePointId sets field value
func (o *VmwareFcdInstantRecoverySpec) SetObjectRestorePointId(v string) {
	o.ObjectRestorePointId = v
}

// GetDestinationCluster returns the DestinationCluster field value
func (o *VmwareFcdInstantRecoverySpec) GetDestinationCluster() VmwareObjectModel {
	if o == nil  {
		var ret VmwareObjectModel
		return ret
	}

	return o.DestinationCluster
}

// GetDestinationClusterOk returns a tuple with the DestinationCluster field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoverySpec) GetDestinationClusterOk() (*VmwareObjectModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationCluster, true
}

// SetDestinationCluster sets field value
func (o *VmwareFcdInstantRecoverySpec) SetDestinationCluster(v VmwareObjectModel) {
	o.DestinationCluster = v
}

// GetDisksMapping returns the DisksMapping field value
func (o *VmwareFcdInstantRecoverySpec) GetDisksMapping() []VmwareFcdInstantRecoveryDiskSpec {
	if o == nil  {
		var ret []VmwareFcdInstantRecoveryDiskSpec
		return ret
	}

	return o.DisksMapping
}

// GetDisksMappingOk returns a tuple with the DisksMapping field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoverySpec) GetDisksMappingOk() (*[]VmwareFcdInstantRecoveryDiskSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisksMapping, true
}

// SetDisksMapping sets field value
func (o *VmwareFcdInstantRecoverySpec) SetDisksMapping(v []VmwareFcdInstantRecoveryDiskSpec) {
	o.DisksMapping = v
}

// GetWriteCache returns the WriteCache field value if set, zero value otherwise.
func (o *VmwareFcdInstantRecoverySpec) GetWriteCache() VmwareFcdWriteCacheSpec {
	if o == nil || o.WriteCache == nil {
		var ret VmwareFcdWriteCacheSpec
		return ret
	}
	return *o.WriteCache
}

// GetWriteCacheOk returns a tuple with the WriteCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdInstantRecoverySpec) GetWriteCacheOk() (*VmwareFcdWriteCacheSpec, bool) {
	if o == nil || o.WriteCache == nil {
		return nil, false
	}
	return o.WriteCache, true
}

// HasWriteCache returns a boolean if a field has been set.
func (o *VmwareFcdInstantRecoverySpec) HasWriteCache() bool {
	if o != nil && o.WriteCache != nil {
		return true
	}

	return false
}

// SetWriteCache gets a reference to the given VmwareFcdWriteCacheSpec and assigns it to the WriteCache field.
func (o *VmwareFcdInstantRecoverySpec) SetWriteCache(v VmwareFcdWriteCacheSpec) {
	o.WriteCache = &v
}

func (o VmwareFcdInstantRecoverySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objectRestorePointId"] = o.ObjectRestorePointId
	}
	if true {
		toSerialize["destinationCluster"] = o.DestinationCluster
	}
	if true {
		toSerialize["disksMapping"] = o.DisksMapping
	}
	if o.WriteCache != nil {
		toSerialize["writeCache"] = o.WriteCache
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareFcdInstantRecoverySpec struct {
	value *VmwareFcdInstantRecoverySpec
	isSet bool
}

func (v NullableVmwareFcdInstantRecoverySpec) Get() *VmwareFcdInstantRecoverySpec {
	return v.value
}

func (v *NullableVmwareFcdInstantRecoverySpec) Set(val *VmwareFcdInstantRecoverySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareFcdInstantRecoverySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareFcdInstantRecoverySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareFcdInstantRecoverySpec(val *VmwareFcdInstantRecoverySpec) *NullableVmwareFcdInstantRecoverySpec {
	return &NullableVmwareFcdInstantRecoverySpec{value: val, isSet: true}
}

func (v NullableVmwareFcdInstantRecoverySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareFcdInstantRecoverySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


