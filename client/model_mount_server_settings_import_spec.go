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
)

// MountServerSettingsImportSpec struct for MountServerSettingsImportSpec
type MountServerSettingsImportSpec struct {
	// Name of the mount server.
	MountServerName string `json:"mountServerName"`
	// Path to the folder used for writing cache during mount operations.
	WriteCacheFolder string `json:"writeCacheFolder"`
	// If *true*, the vPower NFS Service is enabled on the mount server.
	VPowerNFSEnabled bool `json:"vPowerNFSEnabled"`
	VPowerNFSPortSettings *VPowerNFSPortSettingsModel `json:"vPowerNFSPortSettings,omitempty"`
}

// NewMountServerSettingsImportSpec instantiates a new MountServerSettingsImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMountServerSettingsImportSpec(mountServerName string, writeCacheFolder string, vPowerNFSEnabled bool) *MountServerSettingsImportSpec {
	this := MountServerSettingsImportSpec{}
	this.MountServerName = mountServerName
	this.WriteCacheFolder = writeCacheFolder
	this.VPowerNFSEnabled = vPowerNFSEnabled
	return &this
}

// NewMountServerSettingsImportSpecWithDefaults instantiates a new MountServerSettingsImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountServerSettingsImportSpecWithDefaults() *MountServerSettingsImportSpec {
	this := MountServerSettingsImportSpec{}
	return &this
}

// GetMountServerName returns the MountServerName field value
func (o *MountServerSettingsImportSpec) GetMountServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountServerName
}

// GetMountServerNameOk returns a tuple with the MountServerName field value
// and a boolean to check if the value has been set.
func (o *MountServerSettingsImportSpec) GetMountServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountServerName, true
}

// SetMountServerName sets field value
func (o *MountServerSettingsImportSpec) SetMountServerName(v string) {
	o.MountServerName = v
}

// GetWriteCacheFolder returns the WriteCacheFolder field value
func (o *MountServerSettingsImportSpec) GetWriteCacheFolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WriteCacheFolder
}

// GetWriteCacheFolderOk returns a tuple with the WriteCacheFolder field value
// and a boolean to check if the value has been set.
func (o *MountServerSettingsImportSpec) GetWriteCacheFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteCacheFolder, true
}

// SetWriteCacheFolder sets field value
func (o *MountServerSettingsImportSpec) SetWriteCacheFolder(v string) {
	o.WriteCacheFolder = v
}

// GetVPowerNFSEnabled returns the VPowerNFSEnabled field value
func (o *MountServerSettingsImportSpec) GetVPowerNFSEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VPowerNFSEnabled
}

// GetVPowerNFSEnabledOk returns a tuple with the VPowerNFSEnabled field value
// and a boolean to check if the value has been set.
func (o *MountServerSettingsImportSpec) GetVPowerNFSEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VPowerNFSEnabled, true
}

// SetVPowerNFSEnabled sets field value
func (o *MountServerSettingsImportSpec) SetVPowerNFSEnabled(v bool) {
	o.VPowerNFSEnabled = v
}

// GetVPowerNFSPortSettings returns the VPowerNFSPortSettings field value if set, zero value otherwise.
func (o *MountServerSettingsImportSpec) GetVPowerNFSPortSettings() VPowerNFSPortSettingsModel {
	if o == nil || o.VPowerNFSPortSettings == nil {
		var ret VPowerNFSPortSettingsModel
		return ret
	}
	return *o.VPowerNFSPortSettings
}

// GetVPowerNFSPortSettingsOk returns a tuple with the VPowerNFSPortSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MountServerSettingsImportSpec) GetVPowerNFSPortSettingsOk() (*VPowerNFSPortSettingsModel, bool) {
	if o == nil || o.VPowerNFSPortSettings == nil {
		return nil, false
	}
	return o.VPowerNFSPortSettings, true
}

// HasVPowerNFSPortSettings returns a boolean if a field has been set.
func (o *MountServerSettingsImportSpec) HasVPowerNFSPortSettings() bool {
	if o != nil && o.VPowerNFSPortSettings != nil {
		return true
	}

	return false
}

// SetVPowerNFSPortSettings gets a reference to the given VPowerNFSPortSettingsModel and assigns it to the VPowerNFSPortSettings field.
func (o *MountServerSettingsImportSpec) SetVPowerNFSPortSettings(v VPowerNFSPortSettingsModel) {
	o.VPowerNFSPortSettings = &v
}

func (o MountServerSettingsImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mountServerName"] = o.MountServerName
	}
	if true {
		toSerialize["writeCacheFolder"] = o.WriteCacheFolder
	}
	if true {
		toSerialize["vPowerNFSEnabled"] = o.VPowerNFSEnabled
	}
	if o.VPowerNFSPortSettings != nil {
		toSerialize["vPowerNFSPortSettings"] = o.VPowerNFSPortSettings
	}
	return json.Marshal(toSerialize)
}

type NullableMountServerSettingsImportSpec struct {
	value *MountServerSettingsImportSpec
	isSet bool
}

func (v NullableMountServerSettingsImportSpec) Get() *MountServerSettingsImportSpec {
	return v.value
}

func (v *NullableMountServerSettingsImportSpec) Set(val *MountServerSettingsImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableMountServerSettingsImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableMountServerSettingsImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMountServerSettingsImportSpec(val *MountServerSettingsImportSpec) *NullableMountServerSettingsImportSpec {
	return &NullableMountServerSettingsImportSpec{value: val, isSet: true}
}

func (v NullableMountServerSettingsImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMountServerSettingsImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


