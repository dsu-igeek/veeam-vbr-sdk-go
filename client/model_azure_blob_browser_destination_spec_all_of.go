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

// AzureBlobBrowserDestinationSpecAllOf struct for AzureBlobBrowserDestinationSpecAllOf
type AzureBlobBrowserDestinationSpecAllOf struct {
	HostId *string `json:"hostId,omitempty"`
	RegionType EAzureRegionType `json:"regionType"`
	ContainerName string `json:"containerName"`
	FolderType *ECloudBrowserFolderType `json:"folderType,omitempty"`
}

// NewAzureBlobBrowserDestinationSpecAllOf instantiates a new AzureBlobBrowserDestinationSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobBrowserDestinationSpecAllOf(regionType EAzureRegionType, containerName string, ) *AzureBlobBrowserDestinationSpecAllOf {
	this := AzureBlobBrowserDestinationSpecAllOf{}
	this.RegionType = regionType
	this.ContainerName = containerName
	return &this
}

// NewAzureBlobBrowserDestinationSpecAllOfWithDefaults instantiates a new AzureBlobBrowserDestinationSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobBrowserDestinationSpecAllOfWithDefaults() *AzureBlobBrowserDestinationSpecAllOf {
	this := AzureBlobBrowserDestinationSpecAllOf{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *AzureBlobBrowserDestinationSpecAllOf) GetHostId() string {
	if o == nil || o.HostId == nil {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobBrowserDestinationSpecAllOf) GetHostIdOk() (*string, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *AzureBlobBrowserDestinationSpecAllOf) HasHostId() bool {
	if o != nil && o.HostId != nil {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *AzureBlobBrowserDestinationSpecAllOf) SetHostId(v string) {
	o.HostId = &v
}

// GetRegionType returns the RegionType field value
func (o *AzureBlobBrowserDestinationSpecAllOf) GetRegionType() EAzureRegionType {
	if o == nil  {
		var ret EAzureRegionType
		return ret
	}

	return o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value
// and a boolean to check if the value has been set.
func (o *AzureBlobBrowserDestinationSpecAllOf) GetRegionTypeOk() (*EAzureRegionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionType, true
}

// SetRegionType sets field value
func (o *AzureBlobBrowserDestinationSpecAllOf) SetRegionType(v EAzureRegionType) {
	o.RegionType = v
}

// GetContainerName returns the ContainerName field value
func (o *AzureBlobBrowserDestinationSpecAllOf) GetContainerName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *AzureBlobBrowserDestinationSpecAllOf) GetContainerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value
func (o *AzureBlobBrowserDestinationSpecAllOf) SetContainerName(v string) {
	o.ContainerName = v
}

// GetFolderType returns the FolderType field value if set, zero value otherwise.
func (o *AzureBlobBrowserDestinationSpecAllOf) GetFolderType() ECloudBrowserFolderType {
	if o == nil || o.FolderType == nil {
		var ret ECloudBrowserFolderType
		return ret
	}
	return *o.FolderType
}

// GetFolderTypeOk returns a tuple with the FolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobBrowserDestinationSpecAllOf) GetFolderTypeOk() (*ECloudBrowserFolderType, bool) {
	if o == nil || o.FolderType == nil {
		return nil, false
	}
	return o.FolderType, true
}

// HasFolderType returns a boolean if a field has been set.
func (o *AzureBlobBrowserDestinationSpecAllOf) HasFolderType() bool {
	if o != nil && o.FolderType != nil {
		return true
	}

	return false
}

// SetFolderType gets a reference to the given ECloudBrowserFolderType and assigns it to the FolderType field.
func (o *AzureBlobBrowserDestinationSpecAllOf) SetFolderType(v ECloudBrowserFolderType) {
	o.FolderType = &v
}

func (o AzureBlobBrowserDestinationSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HostId != nil {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["regionType"] = o.RegionType
	}
	if true {
		toSerialize["containerName"] = o.ContainerName
	}
	if o.FolderType != nil {
		toSerialize["folderType"] = o.FolderType
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobBrowserDestinationSpecAllOf struct {
	value *AzureBlobBrowserDestinationSpecAllOf
	isSet bool
}

func (v NullableAzureBlobBrowserDestinationSpecAllOf) Get() *AzureBlobBrowserDestinationSpecAllOf {
	return v.value
}

func (v *NullableAzureBlobBrowserDestinationSpecAllOf) Set(val *AzureBlobBrowserDestinationSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobBrowserDestinationSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobBrowserDestinationSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobBrowserDestinationSpecAllOf(val *AzureBlobBrowserDestinationSpecAllOf) *NullableAzureBlobBrowserDestinationSpecAllOf {
	return &NullableAzureBlobBrowserDestinationSpecAllOf{value: val, isSet: true}
}

func (v NullableAzureBlobBrowserDestinationSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobBrowserDestinationSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


