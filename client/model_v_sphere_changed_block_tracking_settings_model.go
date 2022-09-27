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

// VSphereChangedBlockTrackingSettingsModel CBT settings for the backup job.
type VSphereChangedBlockTrackingSettingsModel struct {
	// If *true*, CBT data is used.
	IsEnabled bool `json:"isEnabled"`
	// If *true*, CBT is enabled for all processed VMs even if CBT is disabled in VM configuration. CBT is used for VMs with virtual hardware version 7 or later. These VMs must not have existing snapshots.
	EnableCBTautomatically *bool `json:"enableCBTautomatically,omitempty"`
	// If *true*, CBT is reset before creating active full backups.
	ResetCBTonActiveFull *bool `json:"resetCBTonActiveFull,omitempty"`
}

// NewVSphereChangedBlockTrackingSettingsModel instantiates a new VSphereChangedBlockTrackingSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVSphereChangedBlockTrackingSettingsModel(isEnabled bool, ) *VSphereChangedBlockTrackingSettingsModel {
	this := VSphereChangedBlockTrackingSettingsModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewVSphereChangedBlockTrackingSettingsModelWithDefaults instantiates a new VSphereChangedBlockTrackingSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVSphereChangedBlockTrackingSettingsModelWithDefaults() *VSphereChangedBlockTrackingSettingsModel {
	this := VSphereChangedBlockTrackingSettingsModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *VSphereChangedBlockTrackingSettingsModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *VSphereChangedBlockTrackingSettingsModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *VSphereChangedBlockTrackingSettingsModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetEnableCBTautomatically returns the EnableCBTautomatically field value if set, zero value otherwise.
func (o *VSphereChangedBlockTrackingSettingsModel) GetEnableCBTautomatically() bool {
	if o == nil || o.EnableCBTautomatically == nil {
		var ret bool
		return ret
	}
	return *o.EnableCBTautomatically
}

// GetEnableCBTautomaticallyOk returns a tuple with the EnableCBTautomatically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VSphereChangedBlockTrackingSettingsModel) GetEnableCBTautomaticallyOk() (*bool, bool) {
	if o == nil || o.EnableCBTautomatically == nil {
		return nil, false
	}
	return o.EnableCBTautomatically, true
}

// HasEnableCBTautomatically returns a boolean if a field has been set.
func (o *VSphereChangedBlockTrackingSettingsModel) HasEnableCBTautomatically() bool {
	if o != nil && o.EnableCBTautomatically != nil {
		return true
	}

	return false
}

// SetEnableCBTautomatically gets a reference to the given bool and assigns it to the EnableCBTautomatically field.
func (o *VSphereChangedBlockTrackingSettingsModel) SetEnableCBTautomatically(v bool) {
	o.EnableCBTautomatically = &v
}

// GetResetCBTonActiveFull returns the ResetCBTonActiveFull field value if set, zero value otherwise.
func (o *VSphereChangedBlockTrackingSettingsModel) GetResetCBTonActiveFull() bool {
	if o == nil || o.ResetCBTonActiveFull == nil {
		var ret bool
		return ret
	}
	return *o.ResetCBTonActiveFull
}

// GetResetCBTonActiveFullOk returns a tuple with the ResetCBTonActiveFull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VSphereChangedBlockTrackingSettingsModel) GetResetCBTonActiveFullOk() (*bool, bool) {
	if o == nil || o.ResetCBTonActiveFull == nil {
		return nil, false
	}
	return o.ResetCBTonActiveFull, true
}

// HasResetCBTonActiveFull returns a boolean if a field has been set.
func (o *VSphereChangedBlockTrackingSettingsModel) HasResetCBTonActiveFull() bool {
	if o != nil && o.ResetCBTonActiveFull != nil {
		return true
	}

	return false
}

// SetResetCBTonActiveFull gets a reference to the given bool and assigns it to the ResetCBTonActiveFull field.
func (o *VSphereChangedBlockTrackingSettingsModel) SetResetCBTonActiveFull(v bool) {
	o.ResetCBTonActiveFull = &v
}

func (o VSphereChangedBlockTrackingSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.EnableCBTautomatically != nil {
		toSerialize["enableCBTautomatically"] = o.EnableCBTautomatically
	}
	if o.ResetCBTonActiveFull != nil {
		toSerialize["resetCBTonActiveFull"] = o.ResetCBTonActiveFull
	}
	return json.Marshal(toSerialize)
}

type NullableVSphereChangedBlockTrackingSettingsModel struct {
	value *VSphereChangedBlockTrackingSettingsModel
	isSet bool
}

func (v NullableVSphereChangedBlockTrackingSettingsModel) Get() *VSphereChangedBlockTrackingSettingsModel {
	return v.value
}

func (v *NullableVSphereChangedBlockTrackingSettingsModel) Set(val *VSphereChangedBlockTrackingSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVSphereChangedBlockTrackingSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVSphereChangedBlockTrackingSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVSphereChangedBlockTrackingSettingsModel(val *VSphereChangedBlockTrackingSettingsModel) *NullableVSphereChangedBlockTrackingSettingsModel {
	return &NullableVSphereChangedBlockTrackingSettingsModel{value: val, isSet: true}
}

func (v NullableVSphereChangedBlockTrackingSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVSphereChangedBlockTrackingSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


