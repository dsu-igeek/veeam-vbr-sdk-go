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

// LinuxLocalStorageImportSpec struct for LinuxLocalStorageImportSpec
type LinuxLocalStorageImportSpec struct {
	// Name of the backup repository.
	Name string `json:"name"`
	// Description of the backup repository.
	Description string `json:"description"`
	// VMware vSphere tag assigned to the backup repository.
	Tag string `json:"tag"`
	// ID of the server that is used as a backup repository.
	HostName string `json:"hostName"`
	Type ERepositoryType `json:"type"`
	Repository LinuxLocalRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsImportSpec `json:"mountServer"`
}

// NewLinuxLocalStorageImportSpec instantiates a new LinuxLocalStorageImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxLocalStorageImportSpec(name string, description string, tag string, hostName string, type_ ERepositoryType, repository LinuxLocalRepositorySettingsModel, mountServer MountServerSettingsImportSpec) *LinuxLocalStorageImportSpec {
	this := LinuxLocalStorageImportSpec{}
	this.Name = name
	this.Description = description
	this.Tag = tag
	this.HostName = hostName
	this.Type = type_
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewLinuxLocalStorageImportSpecWithDefaults instantiates a new LinuxLocalStorageImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxLocalStorageImportSpecWithDefaults() *LinuxLocalStorageImportSpec {
	this := LinuxLocalStorageImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *LinuxLocalStorageImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageImportSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LinuxLocalStorageImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *LinuxLocalStorageImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LinuxLocalStorageImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value
func (o *LinuxLocalStorageImportSpec) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageImportSpec) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *LinuxLocalStorageImportSpec) SetTag(v string) {
	o.Tag = v
}

// GetHostName returns the HostName field value
func (o *LinuxLocalStorageImportSpec) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageImportSpec) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *LinuxLocalStorageImportSpec) SetHostName(v string) {
	o.HostName = v
}

// GetType returns the Type field value
func (o *LinuxLocalStorageImportSpec) GetType() ERepositoryType {
	if o == nil {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageImportSpec) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LinuxLocalStorageImportSpec) SetType(v ERepositoryType) {
	o.Type = v
}

// GetRepository returns the Repository field value
func (o *LinuxLocalStorageImportSpec) GetRepository() LinuxLocalRepositorySettingsModel {
	if o == nil {
		var ret LinuxLocalRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageImportSpec) GetRepositoryOk() (*LinuxLocalRepositorySettingsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *LinuxLocalStorageImportSpec) SetRepository(v LinuxLocalRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *LinuxLocalStorageImportSpec) GetMountServer() MountServerSettingsImportSpec {
	if o == nil {
		var ret MountServerSettingsImportSpec
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *LinuxLocalStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec) {
	o.MountServer = v
}

func (o LinuxLocalStorageImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["hostName"] = o.HostName
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxLocalStorageImportSpec struct {
	value *LinuxLocalStorageImportSpec
	isSet bool
}

func (v NullableLinuxLocalStorageImportSpec) Get() *LinuxLocalStorageImportSpec {
	return v.value
}

func (v *NullableLinuxLocalStorageImportSpec) Set(val *LinuxLocalStorageImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxLocalStorageImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxLocalStorageImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxLocalStorageImportSpec(val *LinuxLocalStorageImportSpec) *NullableLinuxLocalStorageImportSpec {
	return &NullableLinuxLocalStorageImportSpec{value: val, isSet: true}
}

func (v NullableLinuxLocalStorageImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxLocalStorageImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


