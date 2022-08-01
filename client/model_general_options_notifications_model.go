/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev2
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GeneralOptionsNotificationsModel struct for GeneralOptionsNotificationsModel
type GeneralOptionsNotificationsModel struct {
	// If *true*, notifications about critical amount of free space in backup storage are enabled.
	StorageSpaceThresholdEnabled bool `json:"storageSpaceThresholdEnabled"`
	// Space threshold of backup storage, in percent.
	StorageSpaceThreshold int32 `json:"storageSpaceThreshold"`
	// If *true*, notifications about critical amount of free space in production datastore are enabled.
	DatastoreSpaceThresholdEnabled bool `json:"datastoreSpaceThresholdEnabled"`
	// Space threshold of production datastore, in percent.
	DatastoreSpaceThreshold int32 `json:"datastoreSpaceThreshold"`
	// If *true* and the `skipVMSpaceThreshold` threshold is reached, Veeam Backup & Replication terminates backup and replication jobs working with production datastores before VM snapshots are taken.
	SkipVMSpaceThresholdEnabled bool `json:"skipVMSpaceThresholdEnabled"`
	// Space threshold of production datastore, in percent.
	SkipVMSpaceThreshold int32 `json:"skipVMSpaceThreshold"`
	// If *true*, notifications about support contract expiration are enabled.
	NotifyOnSupportExpiration bool `json:"notifyOnSupportExpiration"`
	// If *true*, notifications about updates are enabled.
	NotifyOnUpdates bool `json:"notifyOnUpdates"`
}

// NewGeneralOptionsNotificationsModel instantiates a new GeneralOptionsNotificationsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralOptionsNotificationsModel(storageSpaceThresholdEnabled bool, storageSpaceThreshold int32, datastoreSpaceThresholdEnabled bool, datastoreSpaceThreshold int32, skipVMSpaceThresholdEnabled bool, skipVMSpaceThreshold int32, notifyOnSupportExpiration bool, notifyOnUpdates bool, ) *GeneralOptionsNotificationsModel {
	this := GeneralOptionsNotificationsModel{}
	this.StorageSpaceThresholdEnabled = storageSpaceThresholdEnabled
	this.StorageSpaceThreshold = storageSpaceThreshold
	this.DatastoreSpaceThresholdEnabled = datastoreSpaceThresholdEnabled
	this.DatastoreSpaceThreshold = datastoreSpaceThreshold
	this.SkipVMSpaceThresholdEnabled = skipVMSpaceThresholdEnabled
	this.SkipVMSpaceThreshold = skipVMSpaceThreshold
	this.NotifyOnSupportExpiration = notifyOnSupportExpiration
	this.NotifyOnUpdates = notifyOnUpdates
	return &this
}

// NewGeneralOptionsNotificationsModelWithDefaults instantiates a new GeneralOptionsNotificationsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralOptionsNotificationsModelWithDefaults() *GeneralOptionsNotificationsModel {
	this := GeneralOptionsNotificationsModel{}
	return &this
}

// GetStorageSpaceThresholdEnabled returns the StorageSpaceThresholdEnabled field value
func (o *GeneralOptionsNotificationsModel) GetStorageSpaceThresholdEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.StorageSpaceThresholdEnabled
}

// GetStorageSpaceThresholdEnabledOk returns a tuple with the StorageSpaceThresholdEnabled field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsNotificationsModel) GetStorageSpaceThresholdEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageSpaceThresholdEnabled, true
}

// SetStorageSpaceThresholdEnabled sets field value
func (o *GeneralOptionsNotificationsModel) SetStorageSpaceThresholdEnabled(v bool) {
	o.StorageSpaceThresholdEnabled = v
}

// GetStorageSpaceThreshold returns the StorageSpaceThreshold field value
func (o *GeneralOptionsNotificationsModel) GetStorageSpaceThreshold() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.StorageSpaceThreshold
}

// GetStorageSpaceThresholdOk returns a tuple with the StorageSpaceThreshold field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsNotificationsModel) GetStorageSpaceThresholdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageSpaceThreshold, true
}

// SetStorageSpaceThreshold sets field value
func (o *GeneralOptionsNotificationsModel) SetStorageSpaceThreshold(v int32) {
	o.StorageSpaceThreshold = v
}

// GetDatastoreSpaceThresholdEnabled returns the DatastoreSpaceThresholdEnabled field value
func (o *GeneralOptionsNotificationsModel) GetDatastoreSpaceThresholdEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.DatastoreSpaceThresholdEnabled
}

// GetDatastoreSpaceThresholdEnabledOk returns a tuple with the DatastoreSpaceThresholdEnabled field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsNotificationsModel) GetDatastoreSpaceThresholdEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DatastoreSpaceThresholdEnabled, true
}

// SetDatastoreSpaceThresholdEnabled sets field value
func (o *GeneralOptionsNotificationsModel) SetDatastoreSpaceThresholdEnabled(v bool) {
	o.DatastoreSpaceThresholdEnabled = v
}

// GetDatastoreSpaceThreshold returns the DatastoreSpaceThreshold field value
func (o *GeneralOptionsNotificationsModel) GetDatastoreSpaceThreshold() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.DatastoreSpaceThreshold
}

// GetDatastoreSpaceThresholdOk returns a tuple with the DatastoreSpaceThreshold field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsNotificationsModel) GetDatastoreSpaceThresholdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DatastoreSpaceThreshold, true
}

// SetDatastoreSpaceThreshold sets field value
func (o *GeneralOptionsNotificationsModel) SetDatastoreSpaceThreshold(v int32) {
	o.DatastoreSpaceThreshold = v
}

// GetSkipVMSpaceThresholdEnabled returns the SkipVMSpaceThresholdEnabled field value
func (o *GeneralOptionsNotificationsModel) GetSkipVMSpaceThresholdEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.SkipVMSpaceThresholdEnabled
}

// GetSkipVMSpaceThresholdEnabledOk returns a tuple with the SkipVMSpaceThresholdEnabled field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsNotificationsModel) GetSkipVMSpaceThresholdEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SkipVMSpaceThresholdEnabled, true
}

// SetSkipVMSpaceThresholdEnabled sets field value
func (o *GeneralOptionsNotificationsModel) SetSkipVMSpaceThresholdEnabled(v bool) {
	o.SkipVMSpaceThresholdEnabled = v
}

// GetSkipVMSpaceThreshold returns the SkipVMSpaceThreshold field value
func (o *GeneralOptionsNotificationsModel) GetSkipVMSpaceThreshold() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.SkipVMSpaceThreshold
}

// GetSkipVMSpaceThresholdOk returns a tuple with the SkipVMSpaceThreshold field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsNotificationsModel) GetSkipVMSpaceThresholdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SkipVMSpaceThreshold, true
}

// SetSkipVMSpaceThreshold sets field value
func (o *GeneralOptionsNotificationsModel) SetSkipVMSpaceThreshold(v int32) {
	o.SkipVMSpaceThreshold = v
}

// GetNotifyOnSupportExpiration returns the NotifyOnSupportExpiration field value
func (o *GeneralOptionsNotificationsModel) GetNotifyOnSupportExpiration() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.NotifyOnSupportExpiration
}

// GetNotifyOnSupportExpirationOk returns a tuple with the NotifyOnSupportExpiration field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsNotificationsModel) GetNotifyOnSupportExpirationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotifyOnSupportExpiration, true
}

// SetNotifyOnSupportExpiration sets field value
func (o *GeneralOptionsNotificationsModel) SetNotifyOnSupportExpiration(v bool) {
	o.NotifyOnSupportExpiration = v
}

// GetNotifyOnUpdates returns the NotifyOnUpdates field value
func (o *GeneralOptionsNotificationsModel) GetNotifyOnUpdates() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.NotifyOnUpdates
}

// GetNotifyOnUpdatesOk returns a tuple with the NotifyOnUpdates field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsNotificationsModel) GetNotifyOnUpdatesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotifyOnUpdates, true
}

// SetNotifyOnUpdates sets field value
func (o *GeneralOptionsNotificationsModel) SetNotifyOnUpdates(v bool) {
	o.NotifyOnUpdates = v
}

func (o GeneralOptionsNotificationsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["storageSpaceThresholdEnabled"] = o.StorageSpaceThresholdEnabled
	}
	if true {
		toSerialize["storageSpaceThreshold"] = o.StorageSpaceThreshold
	}
	if true {
		toSerialize["datastoreSpaceThresholdEnabled"] = o.DatastoreSpaceThresholdEnabled
	}
	if true {
		toSerialize["datastoreSpaceThreshold"] = o.DatastoreSpaceThreshold
	}
	if true {
		toSerialize["skipVMSpaceThresholdEnabled"] = o.SkipVMSpaceThresholdEnabled
	}
	if true {
		toSerialize["skipVMSpaceThreshold"] = o.SkipVMSpaceThreshold
	}
	if true {
		toSerialize["notifyOnSupportExpiration"] = o.NotifyOnSupportExpiration
	}
	if true {
		toSerialize["notifyOnUpdates"] = o.NotifyOnUpdates
	}
	return json.Marshal(toSerialize)
}

type NullableGeneralOptionsNotificationsModel struct {
	value *GeneralOptionsNotificationsModel
	isSet bool
}

func (v NullableGeneralOptionsNotificationsModel) Get() *GeneralOptionsNotificationsModel {
	return v.value
}

func (v *NullableGeneralOptionsNotificationsModel) Set(val *GeneralOptionsNotificationsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralOptionsNotificationsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralOptionsNotificationsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralOptionsNotificationsModel(val *GeneralOptionsNotificationsModel) *NullableGeneralOptionsNotificationsModel {
	return &NullableGeneralOptionsNotificationsModel{value: val, isSet: true}
}

func (v NullableGeneralOptionsNotificationsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralOptionsNotificationsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


