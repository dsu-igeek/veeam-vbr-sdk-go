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

// AmazonS3BrowserModel struct for AmazonS3BrowserModel
type AmazonS3BrowserModel struct {
	CloudBrowserModel
	HostId string `json:"hostId"`
	RegionType EAmazonRegionType `json:"regionType"`
	Regions *[]AmazonS3RegionBrowserModel `json:"regions,omitempty"`
}

// NewAmazonS3BrowserModel instantiates a new AmazonS3BrowserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonS3BrowserModel(hostId string, regionType EAmazonRegionType, ) *AmazonS3BrowserModel {
	this := AmazonS3BrowserModel{}
	this.HostId = hostId
	this.RegionType = regionType
	return &this
}

// NewAmazonS3BrowserModelWithDefaults instantiates a new AmazonS3BrowserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonS3BrowserModelWithDefaults() *AmazonS3BrowserModel {
	this := AmazonS3BrowserModel{}
	return &this
}

// GetHostId returns the HostId field value
func (o *AmazonS3BrowserModel) GetHostId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *AmazonS3BrowserModel) GetHostIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *AmazonS3BrowserModel) SetHostId(v string) {
	o.HostId = v
}

// GetRegionType returns the RegionType field value
func (o *AmazonS3BrowserModel) GetRegionType() EAmazonRegionType {
	if o == nil  {
		var ret EAmazonRegionType
		return ret
	}

	return o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value
// and a boolean to check if the value has been set.
func (o *AmazonS3BrowserModel) GetRegionTypeOk() (*EAmazonRegionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionType, true
}

// SetRegionType sets field value
func (o *AmazonS3BrowserModel) SetRegionType(v EAmazonRegionType) {
	o.RegionType = v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *AmazonS3BrowserModel) GetRegions() []AmazonS3RegionBrowserModel {
	if o == nil || o.Regions == nil {
		var ret []AmazonS3RegionBrowserModel
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3BrowserModel) GetRegionsOk() (*[]AmazonS3RegionBrowserModel, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *AmazonS3BrowserModel) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []AmazonS3RegionBrowserModel and assigns it to the Regions field.
func (o *AmazonS3BrowserModel) SetRegions(v []AmazonS3RegionBrowserModel) {
	o.Regions = &v
}

func (o AmazonS3BrowserModel) MarshalJSON() ([]byte, error) {
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
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonS3BrowserModel struct {
	value *AmazonS3BrowserModel
	isSet bool
}

func (v NullableAmazonS3BrowserModel) Get() *AmazonS3BrowserModel {
	return v.value
}

func (v *NullableAmazonS3BrowserModel) Set(val *AmazonS3BrowserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonS3BrowserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonS3BrowserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonS3BrowserModel(val *AmazonS3BrowserModel) *NullableAmazonS3BrowserModel {
	return &NullableAmazonS3BrowserModel{value: val, isSet: true}
}

func (v NullableAmazonS3BrowserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonS3BrowserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


