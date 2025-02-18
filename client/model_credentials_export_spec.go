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

// CredentialsExportSpec struct for CredentialsExportSpec
type CredentialsExportSpec struct {
	// If *true*, service credentials are exported.
	ShowHiddenCreds *bool `json:"showHiddenCreds,omitempty"`
	// Array of credentials IDs.
	Ids *[]string `json:"ids,omitempty"`
	// Array of credentials types.
	Types *[]ECredentialsType `json:"types,omitempty"`
	// Array of credentials user names. Wildcard characters are supported.
	Names *[]string `json:"names,omitempty"`
}

// NewCredentialsExportSpec instantiates a new CredentialsExportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsExportSpec() *CredentialsExportSpec {
	this := CredentialsExportSpec{}
	return &this
}

// NewCredentialsExportSpecWithDefaults instantiates a new CredentialsExportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsExportSpecWithDefaults() *CredentialsExportSpec {
	this := CredentialsExportSpec{}
	return &this
}

// GetShowHiddenCreds returns the ShowHiddenCreds field value if set, zero value otherwise.
func (o *CredentialsExportSpec) GetShowHiddenCreds() bool {
	if o == nil || o.ShowHiddenCreds == nil {
		var ret bool
		return ret
	}
	return *o.ShowHiddenCreds
}

// GetShowHiddenCredsOk returns a tuple with the ShowHiddenCreds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsExportSpec) GetShowHiddenCredsOk() (*bool, bool) {
	if o == nil || o.ShowHiddenCreds == nil {
		return nil, false
	}
	return o.ShowHiddenCreds, true
}

// HasShowHiddenCreds returns a boolean if a field has been set.
func (o *CredentialsExportSpec) HasShowHiddenCreds() bool {
	if o != nil && o.ShowHiddenCreds != nil {
		return true
	}

	return false
}

// SetShowHiddenCreds gets a reference to the given bool and assigns it to the ShowHiddenCreds field.
func (o *CredentialsExportSpec) SetShowHiddenCreds(v bool) {
	o.ShowHiddenCreds = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CredentialsExportSpec) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsExportSpec) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CredentialsExportSpec) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CredentialsExportSpec) SetIds(v []string) {
	o.Ids = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *CredentialsExportSpec) GetTypes() []ECredentialsType {
	if o == nil || o.Types == nil {
		var ret []ECredentialsType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsExportSpec) GetTypesOk() (*[]ECredentialsType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *CredentialsExportSpec) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []ECredentialsType and assigns it to the Types field.
func (o *CredentialsExportSpec) SetTypes(v []ECredentialsType) {
	o.Types = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *CredentialsExportSpec) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsExportSpec) GetNamesOk() (*[]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *CredentialsExportSpec) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *CredentialsExportSpec) SetNames(v []string) {
	o.Names = &v
}

func (o CredentialsExportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShowHiddenCreds != nil {
		toSerialize["showHiddenCreds"] = o.ShowHiddenCreds
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsExportSpec struct {
	value *CredentialsExportSpec
	isSet bool
}

func (v NullableCredentialsExportSpec) Get() *CredentialsExportSpec {
	return v.value
}

func (v *NullableCredentialsExportSpec) Set(val *CredentialsExportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsExportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsExportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsExportSpec(val *CredentialsExportSpec) *NullableCredentialsExportSpec {
	return &NullableCredentialsExportSpec{value: val, isSet: true}
}

func (v NullableCredentialsExportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsExportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


