/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 

API version: 1.0-rev2
Contact: support@veeam.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EVmwareInventoryType Type of the VMware vSphere object.
type EVmwareInventoryType string

// List of EVmwareInventoryType
const (
	EVMWAREINVENTORYTYPE_UNKNOWN EVmwareInventoryType = "Unknown"
	EVMWAREINVENTORYTYPE_VIRTUAL_MACHINE EVmwareInventoryType = "VirtualMachine"
	EVMWAREINVENTORYTYPE_V_CENTER_SERVER EVmwareInventoryType = "vCenterServer"
	EVMWAREINVENTORYTYPE_DATACENTER EVmwareInventoryType = "Datacenter"
	EVMWAREINVENTORYTYPE_CLUSTER EVmwareInventoryType = "Cluster"
	EVMWAREINVENTORYTYPE_HOST EVmwareInventoryType = "Host"
	EVMWAREINVENTORYTYPE_RESOURCE_POOL EVmwareInventoryType = "ResourcePool"
	EVMWAREINVENTORYTYPE_FOLDER EVmwareInventoryType = "Folder"
	EVMWAREINVENTORYTYPE_DATASTORE EVmwareInventoryType = "Datastore"
	EVMWAREINVENTORYTYPE_DATASTORE_CLUSTER EVmwareInventoryType = "DatastoreCluster"
	EVMWAREINVENTORYTYPE_STORAGE_POLICY EVmwareInventoryType = "StoragePolicy"
	EVMWAREINVENTORYTYPE_TEMPLATE EVmwareInventoryType = "Template"
	EVMWAREINVENTORYTYPE_COMPUTE_RESOURCE EVmwareInventoryType = "ComputeResource"
	EVMWAREINVENTORYTYPE_VIRTUAL_APP EVmwareInventoryType = "VirtualApp"
	EVMWAREINVENTORYTYPE_TAG EVmwareInventoryType = "Tag"
	EVMWAREINVENTORYTYPE_CATEGORY EVmwareInventoryType = "Category"
	EVMWAREINVENTORYTYPE_MULTITAG EVmwareInventoryType = "Multitag"
)

// All allowed values of EVmwareInventoryType enum
var AllowedEVmwareInventoryTypeEnumValues = []EVmwareInventoryType{
	"Unknown",
	"VirtualMachine",
	"vCenterServer",
	"Datacenter",
	"Cluster",
	"Host",
	"ResourcePool",
	"Folder",
	"Datastore",
	"DatastoreCluster",
	"StoragePolicy",
	"Template",
	"ComputeResource",
	"VirtualApp",
	"Tag",
	"Category",
	"Multitag",
}

func (v *EVmwareInventoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EVmwareInventoryType(value)
	for _, existing := range AllowedEVmwareInventoryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EVmwareInventoryType", value)
}

// NewEVmwareInventoryTypeFromValue returns a pointer to a valid EVmwareInventoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEVmwareInventoryTypeFromValue(v string) (*EVmwareInventoryType, error) {
	ev := EVmwareInventoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EVmwareInventoryType: valid values are %v", v, AllowedEVmwareInventoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EVmwareInventoryType) IsValid() bool {
	for _, existing := range AllowedEVmwareInventoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EVmwareInventoryType value
func (v EVmwareInventoryType) Ptr() *EVmwareInventoryType {
	return &v
}

type NullableEVmwareInventoryType struct {
	value *EVmwareInventoryType
	isSet bool
}

func (v NullableEVmwareInventoryType) Get() *EVmwareInventoryType {
	return v.value
}

func (v *NullableEVmwareInventoryType) Set(val *EVmwareInventoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableEVmwareInventoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableEVmwareInventoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEVmwareInventoryType(val *EVmwareInventoryType) *NullableEVmwareInventoryType {
	return &NullableEVmwareInventoryType{value: val, isSet: true}
}

func (v NullableEVmwareInventoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEVmwareInventoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

