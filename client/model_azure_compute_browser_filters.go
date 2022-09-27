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

// AzureComputeBrowserFilters struct for AzureComputeBrowserFilters
type AzureComputeBrowserFilters struct {
	ShowAllStorageAccounts *bool `json:"showAllStorageAccounts,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	Location *string `json:"location,omitempty"`
	HasNetworks *bool `json:"hasNetworks,omitempty"`
}

// NewAzureComputeBrowserFilters instantiates a new AzureComputeBrowserFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeBrowserFilters() *AzureComputeBrowserFilters {
	this := AzureComputeBrowserFilters{}
	return &this
}

// NewAzureComputeBrowserFiltersWithDefaults instantiates a new AzureComputeBrowserFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeBrowserFiltersWithDefaults() *AzureComputeBrowserFilters {
	this := AzureComputeBrowserFilters{}
	return &this
}

// GetShowAllStorageAccounts returns the ShowAllStorageAccounts field value if set, zero value otherwise.
func (o *AzureComputeBrowserFilters) GetShowAllStorageAccounts() bool {
	if o == nil || o.ShowAllStorageAccounts == nil {
		var ret bool
		return ret
	}
	return *o.ShowAllStorageAccounts
}

// GetShowAllStorageAccountsOk returns a tuple with the ShowAllStorageAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeBrowserFilters) GetShowAllStorageAccountsOk() (*bool, bool) {
	if o == nil || o.ShowAllStorageAccounts == nil {
		return nil, false
	}
	return o.ShowAllStorageAccounts, true
}

// HasShowAllStorageAccounts returns a boolean if a field has been set.
func (o *AzureComputeBrowserFilters) HasShowAllStorageAccounts() bool {
	if o != nil && o.ShowAllStorageAccounts != nil {
		return true
	}

	return false
}

// SetShowAllStorageAccounts gets a reference to the given bool and assigns it to the ShowAllStorageAccounts field.
func (o *AzureComputeBrowserFilters) SetShowAllStorageAccounts(v bool) {
	o.ShowAllStorageAccounts = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *AzureComputeBrowserFilters) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeBrowserFilters) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AzureComputeBrowserFilters) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *AzureComputeBrowserFilters) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AzureComputeBrowserFilters) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeBrowserFilters) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AzureComputeBrowserFilters) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AzureComputeBrowserFilters) SetLocation(v string) {
	o.Location = &v
}

// GetHasNetworks returns the HasNetworks field value if set, zero value otherwise.
func (o *AzureComputeBrowserFilters) GetHasNetworks() bool {
	if o == nil || o.HasNetworks == nil {
		var ret bool
		return ret
	}
	return *o.HasNetworks
}

// GetHasNetworksOk returns a tuple with the HasNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeBrowserFilters) GetHasNetworksOk() (*bool, bool) {
	if o == nil || o.HasNetworks == nil {
		return nil, false
	}
	return o.HasNetworks, true
}

// HasHasNetworks returns a boolean if a field has been set.
func (o *AzureComputeBrowserFilters) HasHasNetworks() bool {
	if o != nil && o.HasNetworks != nil {
		return true
	}

	return false
}

// SetHasNetworks gets a reference to the given bool and assigns it to the HasNetworks field.
func (o *AzureComputeBrowserFilters) SetHasNetworks(v bool) {
	o.HasNetworks = &v
}

func (o AzureComputeBrowserFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShowAllStorageAccounts != nil {
		toSerialize["showAllStorageAccounts"] = o.ShowAllStorageAccounts
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.HasNetworks != nil {
		toSerialize["hasNetworks"] = o.HasNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableAzureComputeBrowserFilters struct {
	value *AzureComputeBrowserFilters
	isSet bool
}

func (v NullableAzureComputeBrowserFilters) Get() *AzureComputeBrowserFilters {
	return v.value
}

func (v *NullableAzureComputeBrowserFilters) Set(val *AzureComputeBrowserFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeBrowserFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeBrowserFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeBrowserFilters(val *AzureComputeBrowserFilters) *NullableAzureComputeBrowserFilters {
	return &NullableAzureComputeBrowserFilters{value: val, isSet: true}
}

func (v NullableAzureComputeBrowserFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeBrowserFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

