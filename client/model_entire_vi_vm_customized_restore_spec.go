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

// EntireViVMCustomizedRestoreSpec Restore to a new location or with different settings.
type EntireViVMCustomizedRestoreSpec struct {
	EntireViVMRestoreSpec
	DestinationHost *VmwareObjectModel `json:"destinationHost,omitempty"`
	ResourcePool *VmwareObjectModel `json:"resourcePool,omitempty"`
	Datastore *RestoreTargetDatastoreSpec `json:"datastore,omitempty"`
	Folder *RestoreTargetFolderSpec `json:"folder,omitempty"`
	Network *RestoreTargetNetworkSpec `json:"network,omitempty"`
}

// NewEntireViVMCustomizedRestoreSpec instantiates a new EntireViVMCustomizedRestoreSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntireViVMCustomizedRestoreSpec() *EntireViVMCustomizedRestoreSpec {
	this := EntireViVMCustomizedRestoreSpec{}
	return &this
}

// NewEntireViVMCustomizedRestoreSpecWithDefaults instantiates a new EntireViVMCustomizedRestoreSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntireViVMCustomizedRestoreSpecWithDefaults() *EntireViVMCustomizedRestoreSpec {
	this := EntireViVMCustomizedRestoreSpec{}
	return &this
}

// GetDestinationHost returns the DestinationHost field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpec) GetDestinationHost() VmwareObjectModel {
	if o == nil || o.DestinationHost == nil {
		var ret VmwareObjectModel
		return ret
	}
	return *o.DestinationHost
}

// GetDestinationHostOk returns a tuple with the DestinationHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpec) GetDestinationHostOk() (*VmwareObjectModel, bool) {
	if o == nil || o.DestinationHost == nil {
		return nil, false
	}
	return o.DestinationHost, true
}

// HasDestinationHost returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpec) HasDestinationHost() bool {
	if o != nil && o.DestinationHost != nil {
		return true
	}

	return false
}

// SetDestinationHost gets a reference to the given VmwareObjectModel and assigns it to the DestinationHost field.
func (o *EntireViVMCustomizedRestoreSpec) SetDestinationHost(v VmwareObjectModel) {
	o.DestinationHost = &v
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpec) GetResourcePool() VmwareObjectModel {
	if o == nil || o.ResourcePool == nil {
		var ret VmwareObjectModel
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpec) GetResourcePoolOk() (*VmwareObjectModel, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpec) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given VmwareObjectModel and assigns it to the ResourcePool field.
func (o *EntireViVMCustomizedRestoreSpec) SetResourcePool(v VmwareObjectModel) {
	o.ResourcePool = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpec) GetDatastore() RestoreTargetDatastoreSpec {
	if o == nil || o.Datastore == nil {
		var ret RestoreTargetDatastoreSpec
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpec) GetDatastoreOk() (*RestoreTargetDatastoreSpec, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpec) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given RestoreTargetDatastoreSpec and assigns it to the Datastore field.
func (o *EntireViVMCustomizedRestoreSpec) SetDatastore(v RestoreTargetDatastoreSpec) {
	o.Datastore = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpec) GetFolder() RestoreTargetFolderSpec {
	if o == nil || o.Folder == nil {
		var ret RestoreTargetFolderSpec
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpec) GetFolderOk() (*RestoreTargetFolderSpec, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpec) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given RestoreTargetFolderSpec and assigns it to the Folder field.
func (o *EntireViVMCustomizedRestoreSpec) SetFolder(v RestoreTargetFolderSpec) {
	o.Folder = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpec) GetNetwork() RestoreTargetNetworkSpec {
	if o == nil || o.Network == nil {
		var ret RestoreTargetNetworkSpec
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpec) GetNetworkOk() (*RestoreTargetNetworkSpec, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given RestoreTargetNetworkSpec and assigns it to the Network field.
func (o *EntireViVMCustomizedRestoreSpec) SetNetwork(v RestoreTargetNetworkSpec) {
	o.Network = &v
}

func (o EntireViVMCustomizedRestoreSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEntireViVMRestoreSpec, errEntireViVMRestoreSpec := json.Marshal(o.EntireViVMRestoreSpec)
	if errEntireViVMRestoreSpec != nil {
		return []byte{}, errEntireViVMRestoreSpec
	}
	errEntireViVMRestoreSpec = json.Unmarshal([]byte(serializedEntireViVMRestoreSpec), &toSerialize)
	if errEntireViVMRestoreSpec != nil {
		return []byte{}, errEntireViVMRestoreSpec
	}
	if o.DestinationHost != nil {
		toSerialize["destinationHost"] = o.DestinationHost
	}
	if o.ResourcePool != nil {
		toSerialize["resourcePool"] = o.ResourcePool
	}
	if o.Datastore != nil {
		toSerialize["datastore"] = o.Datastore
	}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableEntireViVMCustomizedRestoreSpec struct {
	value *EntireViVMCustomizedRestoreSpec
	isSet bool
}

func (v NullableEntireViVMCustomizedRestoreSpec) Get() *EntireViVMCustomizedRestoreSpec {
	return v.value
}

func (v *NullableEntireViVMCustomizedRestoreSpec) Set(val *EntireViVMCustomizedRestoreSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableEntireViVMCustomizedRestoreSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableEntireViVMCustomizedRestoreSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntireViVMCustomizedRestoreSpec(val *EntireViVMCustomizedRestoreSpec) *NullableEntireViVMCustomizedRestoreSpec {
	return &NullableEntireViVMCustomizedRestoreSpec{value: val, isSet: true}
}

func (v NullableEntireViVMCustomizedRestoreSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntireViVMCustomizedRestoreSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

