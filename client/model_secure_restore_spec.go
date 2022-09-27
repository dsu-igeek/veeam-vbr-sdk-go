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

// SecureRestoreSpec Secure restore settings.
type SecureRestoreSpec struct {
	// If *true*, Veeam Backup & Replication will scan machine data with antivirus software before restoring the machine to the production environment.
	EnableAntivirusScan bool `json:"enableAntivirusScan"`
	VirusDetectionAction *EVirusDetectionAction `json:"virusDetectionAction,omitempty"`
	// If *true*, the antivirus will continue machine scan after the first malware is found.
	EnableEntireVolumeScan *bool `json:"enableEntireVolumeScan,omitempty"`
}

// NewSecureRestoreSpec instantiates a new SecureRestoreSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecureRestoreSpec(enableAntivirusScan bool, ) *SecureRestoreSpec {
	this := SecureRestoreSpec{}
	this.EnableAntivirusScan = enableAntivirusScan
	return &this
}

// NewSecureRestoreSpecWithDefaults instantiates a new SecureRestoreSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecureRestoreSpecWithDefaults() *SecureRestoreSpec {
	this := SecureRestoreSpec{}
	return &this
}

// GetEnableAntivirusScan returns the EnableAntivirusScan field value
func (o *SecureRestoreSpec) GetEnableAntivirusScan() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.EnableAntivirusScan
}

// GetEnableAntivirusScanOk returns a tuple with the EnableAntivirusScan field value
// and a boolean to check if the value has been set.
func (o *SecureRestoreSpec) GetEnableAntivirusScanOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnableAntivirusScan, true
}

// SetEnableAntivirusScan sets field value
func (o *SecureRestoreSpec) SetEnableAntivirusScan(v bool) {
	o.EnableAntivirusScan = v
}

// GetVirusDetectionAction returns the VirusDetectionAction field value if set, zero value otherwise.
func (o *SecureRestoreSpec) GetVirusDetectionAction() EVirusDetectionAction {
	if o == nil || o.VirusDetectionAction == nil {
		var ret EVirusDetectionAction
		return ret
	}
	return *o.VirusDetectionAction
}

// GetVirusDetectionActionOk returns a tuple with the VirusDetectionAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRestoreSpec) GetVirusDetectionActionOk() (*EVirusDetectionAction, bool) {
	if o == nil || o.VirusDetectionAction == nil {
		return nil, false
	}
	return o.VirusDetectionAction, true
}

// HasVirusDetectionAction returns a boolean if a field has been set.
func (o *SecureRestoreSpec) HasVirusDetectionAction() bool {
	if o != nil && o.VirusDetectionAction != nil {
		return true
	}

	return false
}

// SetVirusDetectionAction gets a reference to the given EVirusDetectionAction and assigns it to the VirusDetectionAction field.
func (o *SecureRestoreSpec) SetVirusDetectionAction(v EVirusDetectionAction) {
	o.VirusDetectionAction = &v
}

// GetEnableEntireVolumeScan returns the EnableEntireVolumeScan field value if set, zero value otherwise.
func (o *SecureRestoreSpec) GetEnableEntireVolumeScan() bool {
	if o == nil || o.EnableEntireVolumeScan == nil {
		var ret bool
		return ret
	}
	return *o.EnableEntireVolumeScan
}

// GetEnableEntireVolumeScanOk returns a tuple with the EnableEntireVolumeScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRestoreSpec) GetEnableEntireVolumeScanOk() (*bool, bool) {
	if o == nil || o.EnableEntireVolumeScan == nil {
		return nil, false
	}
	return o.EnableEntireVolumeScan, true
}

// HasEnableEntireVolumeScan returns a boolean if a field has been set.
func (o *SecureRestoreSpec) HasEnableEntireVolumeScan() bool {
	if o != nil && o.EnableEntireVolumeScan != nil {
		return true
	}

	return false
}

// SetEnableEntireVolumeScan gets a reference to the given bool and assigns it to the EnableEntireVolumeScan field.
func (o *SecureRestoreSpec) SetEnableEntireVolumeScan(v bool) {
	o.EnableEntireVolumeScan = &v
}

func (o SecureRestoreSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enableAntivirusScan"] = o.EnableAntivirusScan
	}
	if o.VirusDetectionAction != nil {
		toSerialize["virusDetectionAction"] = o.VirusDetectionAction
	}
	if o.EnableEntireVolumeScan != nil {
		toSerialize["enableEntireVolumeScan"] = o.EnableEntireVolumeScan
	}
	return json.Marshal(toSerialize)
}

type NullableSecureRestoreSpec struct {
	value *SecureRestoreSpec
	isSet bool
}

func (v NullableSecureRestoreSpec) Get() *SecureRestoreSpec {
	return v.value
}

func (v *NullableSecureRestoreSpec) Set(val *SecureRestoreSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSecureRestoreSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSecureRestoreSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecureRestoreSpec(val *SecureRestoreSpec) *NullableSecureRestoreSpec {
	return &NullableSecureRestoreSpec{value: val, isSet: true}
}

func (v NullableSecureRestoreSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecureRestoreSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


