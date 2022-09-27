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

// FullBackupMaintenanceModel Maintenance settings for full backup files.
type FullBackupMaintenanceModel struct {
	RemoveData *FullBackupMaintenanceRemoveDataModel `json:"RemoveData,omitempty"`
	DefragmentAndCompact *FullBackupMaintenanceDefragmentAndCompactModel `json:"defragmentAndCompact,omitempty"`
}

// NewFullBackupMaintenanceModel instantiates a new FullBackupMaintenanceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullBackupMaintenanceModel() *FullBackupMaintenanceModel {
	this := FullBackupMaintenanceModel{}
	return &this
}

// NewFullBackupMaintenanceModelWithDefaults instantiates a new FullBackupMaintenanceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullBackupMaintenanceModelWithDefaults() *FullBackupMaintenanceModel {
	this := FullBackupMaintenanceModel{}
	return &this
}

// GetRemoveData returns the RemoveData field value if set, zero value otherwise.
func (o *FullBackupMaintenanceModel) GetRemoveData() FullBackupMaintenanceRemoveDataModel {
	if o == nil || o.RemoveData == nil {
		var ret FullBackupMaintenanceRemoveDataModel
		return ret
	}
	return *o.RemoveData
}

// GetRemoveDataOk returns a tuple with the RemoveData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullBackupMaintenanceModel) GetRemoveDataOk() (*FullBackupMaintenanceRemoveDataModel, bool) {
	if o == nil || o.RemoveData == nil {
		return nil, false
	}
	return o.RemoveData, true
}

// HasRemoveData returns a boolean if a field has been set.
func (o *FullBackupMaintenanceModel) HasRemoveData() bool {
	if o != nil && o.RemoveData != nil {
		return true
	}

	return false
}

// SetRemoveData gets a reference to the given FullBackupMaintenanceRemoveDataModel and assigns it to the RemoveData field.
func (o *FullBackupMaintenanceModel) SetRemoveData(v FullBackupMaintenanceRemoveDataModel) {
	o.RemoveData = &v
}

// GetDefragmentAndCompact returns the DefragmentAndCompact field value if set, zero value otherwise.
func (o *FullBackupMaintenanceModel) GetDefragmentAndCompact() FullBackupMaintenanceDefragmentAndCompactModel {
	if o == nil || o.DefragmentAndCompact == nil {
		var ret FullBackupMaintenanceDefragmentAndCompactModel
		return ret
	}
	return *o.DefragmentAndCompact
}

// GetDefragmentAndCompactOk returns a tuple with the DefragmentAndCompact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullBackupMaintenanceModel) GetDefragmentAndCompactOk() (*FullBackupMaintenanceDefragmentAndCompactModel, bool) {
	if o == nil || o.DefragmentAndCompact == nil {
		return nil, false
	}
	return o.DefragmentAndCompact, true
}

// HasDefragmentAndCompact returns a boolean if a field has been set.
func (o *FullBackupMaintenanceModel) HasDefragmentAndCompact() bool {
	if o != nil && o.DefragmentAndCompact != nil {
		return true
	}

	return false
}

// SetDefragmentAndCompact gets a reference to the given FullBackupMaintenanceDefragmentAndCompactModel and assigns it to the DefragmentAndCompact field.
func (o *FullBackupMaintenanceModel) SetDefragmentAndCompact(v FullBackupMaintenanceDefragmentAndCompactModel) {
	o.DefragmentAndCompact = &v
}

func (o FullBackupMaintenanceModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoveData != nil {
		toSerialize["RemoveData"] = o.RemoveData
	}
	if o.DefragmentAndCompact != nil {
		toSerialize["defragmentAndCompact"] = o.DefragmentAndCompact
	}
	return json.Marshal(toSerialize)
}

type NullableFullBackupMaintenanceModel struct {
	value *FullBackupMaintenanceModel
	isSet bool
}

func (v NullableFullBackupMaintenanceModel) Get() *FullBackupMaintenanceModel {
	return v.value
}

func (v *NullableFullBackupMaintenanceModel) Set(val *FullBackupMaintenanceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFullBackupMaintenanceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFullBackupMaintenanceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullBackupMaintenanceModel(val *FullBackupMaintenanceModel) *NullableFullBackupMaintenanceModel {
	return &NullableFullBackupMaintenanceModel{value: val, isSet: true}
}

func (v NullableFullBackupMaintenanceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullBackupMaintenanceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


