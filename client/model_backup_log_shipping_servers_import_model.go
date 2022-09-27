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

// BackupLogShippingServersImportModel Log shipping server used to transport logs.
type BackupLogShippingServersImportModel struct {
	// If *true*, Veeam Backup & Replication chooses an optimal log shipping server automatically.
	AutoSelection bool `json:"autoSelection"`
	// Array of servers that are explicitly selected for log shipping.
	ShippingServerNames *[]string `json:"shippingServerNames,omitempty"`
}

// NewBackupLogShippingServersImportModel instantiates a new BackupLogShippingServersImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupLogShippingServersImportModel(autoSelection bool, ) *BackupLogShippingServersImportModel {
	this := BackupLogShippingServersImportModel{}
	this.AutoSelection = autoSelection
	return &this
}

// NewBackupLogShippingServersImportModelWithDefaults instantiates a new BackupLogShippingServersImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupLogShippingServersImportModelWithDefaults() *BackupLogShippingServersImportModel {
	this := BackupLogShippingServersImportModel{}
	return &this
}

// GetAutoSelection returns the AutoSelection field value
func (o *BackupLogShippingServersImportModel) GetAutoSelection() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AutoSelection
}

// GetAutoSelectionOk returns a tuple with the AutoSelection field value
// and a boolean to check if the value has been set.
func (o *BackupLogShippingServersImportModel) GetAutoSelectionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AutoSelection, true
}

// SetAutoSelection sets field value
func (o *BackupLogShippingServersImportModel) SetAutoSelection(v bool) {
	o.AutoSelection = v
}

// GetShippingServerNames returns the ShippingServerNames field value if set, zero value otherwise.
func (o *BackupLogShippingServersImportModel) GetShippingServerNames() []string {
	if o == nil || o.ShippingServerNames == nil {
		var ret []string
		return ret
	}
	return *o.ShippingServerNames
}

// GetShippingServerNamesOk returns a tuple with the ShippingServerNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupLogShippingServersImportModel) GetShippingServerNamesOk() (*[]string, bool) {
	if o == nil || o.ShippingServerNames == nil {
		return nil, false
	}
	return o.ShippingServerNames, true
}

// HasShippingServerNames returns a boolean if a field has been set.
func (o *BackupLogShippingServersImportModel) HasShippingServerNames() bool {
	if o != nil && o.ShippingServerNames != nil {
		return true
	}

	return false
}

// SetShippingServerNames gets a reference to the given []string and assigns it to the ShippingServerNames field.
func (o *BackupLogShippingServersImportModel) SetShippingServerNames(v []string) {
	o.ShippingServerNames = &v
}

func (o BackupLogShippingServersImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["autoSelection"] = o.AutoSelection
	}
	if o.ShippingServerNames != nil {
		toSerialize["shippingServerNames"] = o.ShippingServerNames
	}
	return json.Marshal(toSerialize)
}

type NullableBackupLogShippingServersImportModel struct {
	value *BackupLogShippingServersImportModel
	isSet bool
}

func (v NullableBackupLogShippingServersImportModel) Get() *BackupLogShippingServersImportModel {
	return v.value
}

func (v *NullableBackupLogShippingServersImportModel) Set(val *BackupLogShippingServersImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupLogShippingServersImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupLogShippingServersImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupLogShippingServersImportModel(val *BackupLogShippingServersImportModel) *NullableBackupLogShippingServersImportModel {
	return &NullableBackupLogShippingServersImportModel{value: val, isSet: true}
}

func (v NullableBackupLogShippingServersImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupLogShippingServersImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


