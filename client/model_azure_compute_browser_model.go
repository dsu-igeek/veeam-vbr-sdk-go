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

// AzureComputeBrowserModel struct for AzureComputeBrowserModel
type AzureComputeBrowserModel struct {
	CloudBrowserModel
	RegionType *string `json:"regionType,omitempty"`
	Subscriptions *[]AzureSubscriptionBrowserModel `json:"subscriptions,omitempty"`
}

// NewAzureComputeBrowserModel instantiates a new AzureComputeBrowserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeBrowserModel() *AzureComputeBrowserModel {
	this := AzureComputeBrowserModel{}
	return &this
}

// NewAzureComputeBrowserModelWithDefaults instantiates a new AzureComputeBrowserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeBrowserModelWithDefaults() *AzureComputeBrowserModel {
	this := AzureComputeBrowserModel{}
	return &this
}

// GetRegionType returns the RegionType field value if set, zero value otherwise.
func (o *AzureComputeBrowserModel) GetRegionType() string {
	if o == nil || o.RegionType == nil {
		var ret string
		return ret
	}
	return *o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeBrowserModel) GetRegionTypeOk() (*string, bool) {
	if o == nil || o.RegionType == nil {
		return nil, false
	}
	return o.RegionType, true
}

// HasRegionType returns a boolean if a field has been set.
func (o *AzureComputeBrowserModel) HasRegionType() bool {
	if o != nil && o.RegionType != nil {
		return true
	}

	return false
}

// SetRegionType gets a reference to the given string and assigns it to the RegionType field.
func (o *AzureComputeBrowserModel) SetRegionType(v string) {
	o.RegionType = &v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *AzureComputeBrowserModel) GetSubscriptions() []AzureSubscriptionBrowserModel {
	if o == nil || o.Subscriptions == nil {
		var ret []AzureSubscriptionBrowserModel
		return ret
	}
	return *o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeBrowserModel) GetSubscriptionsOk() (*[]AzureSubscriptionBrowserModel, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *AzureComputeBrowserModel) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []AzureSubscriptionBrowserModel and assigns it to the Subscriptions field.
func (o *AzureComputeBrowserModel) SetSubscriptions(v []AzureSubscriptionBrowserModel) {
	o.Subscriptions = &v
}

func (o AzureComputeBrowserModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserModel, errCloudBrowserModel := json.Marshal(o.CloudBrowserModel)
	if errCloudBrowserModel != nil {
		return []byte{}, errCloudBrowserModel
	}
	errCloudBrowserModel = json.Unmarshal([]byte(serializedCloudBrowserModel), &toSerialize)
	if errCloudBrowserModel != nil {
		return []byte{}, errCloudBrowserModel
	}
	if o.RegionType != nil {
		toSerialize["regionType"] = o.RegionType
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	return json.Marshal(toSerialize)
}

type NullableAzureComputeBrowserModel struct {
	value *AzureComputeBrowserModel
	isSet bool
}

func (v NullableAzureComputeBrowserModel) Get() *AzureComputeBrowserModel {
	return v.value
}

func (v *NullableAzureComputeBrowserModel) Set(val *AzureComputeBrowserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeBrowserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeBrowserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeBrowserModel(val *AzureComputeBrowserModel) *NullableAzureComputeBrowserModel {
	return &NullableAzureComputeBrowserModel{value: val, isSet: true}
}

func (v NullableAzureComputeBrowserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeBrowserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


