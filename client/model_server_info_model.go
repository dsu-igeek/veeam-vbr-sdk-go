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

// ServerInfoModel struct for ServerInfoModel
type ServerInfoModel struct {
	VbrId *string `json:"vbrId,omitempty"`
	Name string `json:"name"`
	BuildVersion string `json:"buildVersion"`
	Patches []string `json:"patches"`
}

// NewServerInfoModel instantiates a new ServerInfoModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInfoModel(name string, buildVersion string, patches []string, ) *ServerInfoModel {
	this := ServerInfoModel{}
	this.Name = name
	this.BuildVersion = buildVersion
	this.Patches = patches
	return &this
}

// NewServerInfoModelWithDefaults instantiates a new ServerInfoModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInfoModelWithDefaults() *ServerInfoModel {
	this := ServerInfoModel{}
	return &this
}

// GetVbrId returns the VbrId field value if set, zero value otherwise.
func (o *ServerInfoModel) GetVbrId() string {
	if o == nil || o.VbrId == nil {
		var ret string
		return ret
	}
	return *o.VbrId
}

// GetVbrIdOk returns a tuple with the VbrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInfoModel) GetVbrIdOk() (*string, bool) {
	if o == nil || o.VbrId == nil {
		return nil, false
	}
	return o.VbrId, true
}

// HasVbrId returns a boolean if a field has been set.
func (o *ServerInfoModel) HasVbrId() bool {
	if o != nil && o.VbrId != nil {
		return true
	}

	return false
}

// SetVbrId gets a reference to the given string and assigns it to the VbrId field.
func (o *ServerInfoModel) SetVbrId(v string) {
	o.VbrId = &v
}

// GetName returns the Name field value
func (o *ServerInfoModel) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerInfoModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerInfoModel) SetName(v string) {
	o.Name = v
}

// GetBuildVersion returns the BuildVersion field value
func (o *ServerInfoModel) GetBuildVersion() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BuildVersion
}

// GetBuildVersionOk returns a tuple with the BuildVersion field value
// and a boolean to check if the value has been set.
func (o *ServerInfoModel) GetBuildVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BuildVersion, true
}

// SetBuildVersion sets field value
func (o *ServerInfoModel) SetBuildVersion(v string) {
	o.BuildVersion = v
}

// GetPatches returns the Patches field value
func (o *ServerInfoModel) GetPatches() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Patches
}

// GetPatchesOk returns a tuple with the Patches field value
// and a boolean to check if the value has been set.
func (o *ServerInfoModel) GetPatchesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Patches, true
}

// SetPatches sets field value
func (o *ServerInfoModel) SetPatches(v []string) {
	o.Patches = v
}

func (o ServerInfoModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VbrId != nil {
		toSerialize["vbrId"] = o.VbrId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["buildVersion"] = o.BuildVersion
	}
	if true {
		toSerialize["patches"] = o.Patches
	}
	return json.Marshal(toSerialize)
}

type NullableServerInfoModel struct {
	value *ServerInfoModel
	isSet bool
}

func (v NullableServerInfoModel) Get() *ServerInfoModel {
	return v.value
}

func (v *NullableServerInfoModel) Set(val *ServerInfoModel) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInfoModel) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInfoModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInfoModel(val *ServerInfoModel) *NullableServerInfoModel {
	return &NullableServerInfoModel{value: val, isSet: true}
}

func (v NullableServerInfoModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInfoModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


