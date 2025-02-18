/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BackupLinuxScriptModel Paths to pre-freeze and post-thaw scripts for Linux VMs.
type BackupLinuxScriptModel struct {
	// Path to a pre-freeze script.
	PreFreezeScript *string `json:"preFreezeScript,omitempty"`
	// Path to a post-thaw script.
	PostThawScript *string `json:"postThawScript,omitempty"`
}

// NewBackupLinuxScriptModel instantiates a new BackupLinuxScriptModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupLinuxScriptModel() *BackupLinuxScriptModel {
	this := BackupLinuxScriptModel{}
	return &this
}

// NewBackupLinuxScriptModelWithDefaults instantiates a new BackupLinuxScriptModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupLinuxScriptModelWithDefaults() *BackupLinuxScriptModel {
	this := BackupLinuxScriptModel{}
	return &this
}

// GetPreFreezeScript returns the PreFreezeScript field value if set, zero value otherwise.
func (o *BackupLinuxScriptModel) GetPreFreezeScript() string {
	if o == nil || o.PreFreezeScript == nil {
		var ret string
		return ret
	}
	return *o.PreFreezeScript
}

// GetPreFreezeScriptOk returns a tuple with the PreFreezeScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupLinuxScriptModel) GetPreFreezeScriptOk() (*string, bool) {
	if o == nil || o.PreFreezeScript == nil {
		return nil, false
	}
	return o.PreFreezeScript, true
}

// HasPreFreezeScript returns a boolean if a field has been set.
func (o *BackupLinuxScriptModel) HasPreFreezeScript() bool {
	if o != nil && o.PreFreezeScript != nil {
		return true
	}

	return false
}

// SetPreFreezeScript gets a reference to the given string and assigns it to the PreFreezeScript field.
func (o *BackupLinuxScriptModel) SetPreFreezeScript(v string) {
	o.PreFreezeScript = &v
}

// GetPostThawScript returns the PostThawScript field value if set, zero value otherwise.
func (o *BackupLinuxScriptModel) GetPostThawScript() string {
	if o == nil || o.PostThawScript == nil {
		var ret string
		return ret
	}
	return *o.PostThawScript
}

// GetPostThawScriptOk returns a tuple with the PostThawScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupLinuxScriptModel) GetPostThawScriptOk() (*string, bool) {
	if o == nil || o.PostThawScript == nil {
		return nil, false
	}
	return o.PostThawScript, true
}

// HasPostThawScript returns a boolean if a field has been set.
func (o *BackupLinuxScriptModel) HasPostThawScript() bool {
	if o != nil && o.PostThawScript != nil {
		return true
	}

	return false
}

// SetPostThawScript gets a reference to the given string and assigns it to the PostThawScript field.
func (o *BackupLinuxScriptModel) SetPostThawScript(v string) {
	o.PostThawScript = &v
}

func (o BackupLinuxScriptModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PreFreezeScript != nil {
		toSerialize["preFreezeScript"] = o.PreFreezeScript
	}
	if o.PostThawScript != nil {
		toSerialize["postThawScript"] = o.PostThawScript
	}
	return json.Marshal(toSerialize)
}

type NullableBackupLinuxScriptModel struct {
	value *BackupLinuxScriptModel
	isSet bool
}

func (v NullableBackupLinuxScriptModel) Get() *BackupLinuxScriptModel {
	return v.value
}

func (v *NullableBackupLinuxScriptModel) Set(val *BackupLinuxScriptModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupLinuxScriptModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupLinuxScriptModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupLinuxScriptModel(val *BackupLinuxScriptModel) *NullableBackupLinuxScriptModel {
	return &NullableBackupLinuxScriptModel{value: val, isSet: true}
}

func (v NullableBackupLinuxScriptModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupLinuxScriptModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


