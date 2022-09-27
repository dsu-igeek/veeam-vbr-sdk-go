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

// AmazonEC2BrowserModel struct for AmazonEC2BrowserModel
type AmazonEC2BrowserModel struct {
	CloudBrowserModel
	// ID of the server that is used as a backup repository.
	HostId string `json:"hostId"`
	RegionType EAmazonRegionType `json:"regionType"`
	// Array of regions.
	Regions []AmazonEC2RegionBrowserModel `json:"regions"`
}

// NewAmazonEC2BrowserModel instantiates a new AmazonEC2BrowserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonEC2BrowserModel(hostId string, regionType EAmazonRegionType, regions []AmazonEC2RegionBrowserModel, ) *AmazonEC2BrowserModel {
	this := AmazonEC2BrowserModel{}
	this.HostId = hostId
	this.RegionType = regionType
	this.Regions = regions
	return &this
}

// NewAmazonEC2BrowserModelWithDefaults instantiates a new AmazonEC2BrowserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonEC2BrowserModelWithDefaults() *AmazonEC2BrowserModel {
	this := AmazonEC2BrowserModel{}
	return &this
}

// GetHostId returns the HostId field value
func (o *AmazonEC2BrowserModel) GetHostId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *AmazonEC2BrowserModel) GetHostIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *AmazonEC2BrowserModel) SetHostId(v string) {
	o.HostId = v
}

// GetRegionType returns the RegionType field value
func (o *AmazonEC2BrowserModel) GetRegionType() EAmazonRegionType {
	if o == nil  {
		var ret EAmazonRegionType
		return ret
	}

	return o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value
// and a boolean to check if the value has been set.
func (o *AmazonEC2BrowserModel) GetRegionTypeOk() (*EAmazonRegionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionType, true
}

// SetRegionType sets field value
func (o *AmazonEC2BrowserModel) SetRegionType(v EAmazonRegionType) {
	o.RegionType = v
}

// GetRegions returns the Regions field value
func (o *AmazonEC2BrowserModel) GetRegions() []AmazonEC2RegionBrowserModel {
	if o == nil  {
		var ret []AmazonEC2RegionBrowserModel
		return ret
	}

	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value
// and a boolean to check if the value has been set.
func (o *AmazonEC2BrowserModel) GetRegionsOk() (*[]AmazonEC2RegionBrowserModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Regions, true
}

// SetRegions sets field value
func (o *AmazonEC2BrowserModel) SetRegions(v []AmazonEC2RegionBrowserModel) {
	o.Regions = v
}

func (o AmazonEC2BrowserModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserModel, errCloudBrowserModel := json.Marshal(o.CloudBrowserModel)
	if errCloudBrowserModel != nil {
		return []byte{}, errCloudBrowserModel
	}
	errCloudBrowserModel = json.Unmarshal([]byte(serializedCloudBrowserModel), &toSerialize)
	if errCloudBrowserModel != nil {
		return []byte{}, errCloudBrowserModel
	}
	if true {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["regionType"] = o.RegionType
	}
	if true {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonEC2BrowserModel struct {
	value *AmazonEC2BrowserModel
	isSet bool
}

func (v NullableAmazonEC2BrowserModel) Get() *AmazonEC2BrowserModel {
	return v.value
}

func (v *NullableAmazonEC2BrowserModel) Set(val *AmazonEC2BrowserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonEC2BrowserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonEC2BrowserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonEC2BrowserModel(val *AmazonEC2BrowserModel) *NullableAmazonEC2BrowserModel {
	return &NullableAmazonEC2BrowserModel{value: val, isSet: true}
}

func (v NullableAmazonEC2BrowserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonEC2BrowserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


