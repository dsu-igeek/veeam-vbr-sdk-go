/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 

API version: 1.0-rev2
Contact: support@veeam.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GuestOsCredentialsImportModel VM custom credentials.
type GuestOsCredentialsImportModel struct {
	Creds *CredentialsImportModel `json:"creds,omitempty"`
	// Individual credentials for VMs.
	CredentialsPerMachine []GuestOsCredentialsPerMachineImportModel `json:"credentialsPerMachine,omitempty"`
}

// NewGuestOsCredentialsImportModel instantiates a new GuestOsCredentialsImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestOsCredentialsImportModel() *GuestOsCredentialsImportModel {
	this := GuestOsCredentialsImportModel{}
	return &this
}

// NewGuestOsCredentialsImportModelWithDefaults instantiates a new GuestOsCredentialsImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestOsCredentialsImportModelWithDefaults() *GuestOsCredentialsImportModel {
	this := GuestOsCredentialsImportModel{}
	return &this
}

// GetCreds returns the Creds field value if set, zero value otherwise.
func (o *GuestOsCredentialsImportModel) GetCreds() CredentialsImportModel {
	if o == nil || o.Creds == nil {
		var ret CredentialsImportModel
		return ret
	}
	return *o.Creds
}

// GetCredsOk returns a tuple with the Creds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestOsCredentialsImportModel) GetCredsOk() (*CredentialsImportModel, bool) {
	if o == nil || o.Creds == nil {
		return nil, false
	}
	return o.Creds, true
}

// HasCreds returns a boolean if a field has been set.
func (o *GuestOsCredentialsImportModel) HasCreds() bool {
	if o != nil && o.Creds != nil {
		return true
	}

	return false
}

// SetCreds gets a reference to the given CredentialsImportModel and assigns it to the Creds field.
func (o *GuestOsCredentialsImportModel) SetCreds(v CredentialsImportModel) {
	o.Creds = &v
}

// GetCredentialsPerMachine returns the CredentialsPerMachine field value if set, zero value otherwise.
func (o *GuestOsCredentialsImportModel) GetCredentialsPerMachine() []GuestOsCredentialsPerMachineImportModel {
	if o == nil || o.CredentialsPerMachine == nil {
		var ret []GuestOsCredentialsPerMachineImportModel
		return ret
	}
	return o.CredentialsPerMachine
}

// GetCredentialsPerMachineOk returns a tuple with the CredentialsPerMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestOsCredentialsImportModel) GetCredentialsPerMachineOk() ([]GuestOsCredentialsPerMachineImportModel, bool) {
	if o == nil || o.CredentialsPerMachine == nil {
		return nil, false
	}
	return o.CredentialsPerMachine, true
}

// HasCredentialsPerMachine returns a boolean if a field has been set.
func (o *GuestOsCredentialsImportModel) HasCredentialsPerMachine() bool {
	if o != nil && o.CredentialsPerMachine != nil {
		return true
	}

	return false
}

// SetCredentialsPerMachine gets a reference to the given []GuestOsCredentialsPerMachineImportModel and assigns it to the CredentialsPerMachine field.
func (o *GuestOsCredentialsImportModel) SetCredentialsPerMachine(v []GuestOsCredentialsPerMachineImportModel) {
	o.CredentialsPerMachine = v
}

func (o GuestOsCredentialsImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Creds != nil {
		toSerialize["creds"] = o.Creds
	}
	if o.CredentialsPerMachine != nil {
		toSerialize["credentialsPerMachine"] = o.CredentialsPerMachine
	}
	return json.Marshal(toSerialize)
}

type NullableGuestOsCredentialsImportModel struct {
	value *GuestOsCredentialsImportModel
	isSet bool
}

func (v NullableGuestOsCredentialsImportModel) Get() *GuestOsCredentialsImportModel {
	return v.value
}

func (v *NullableGuestOsCredentialsImportModel) Set(val *GuestOsCredentialsImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestOsCredentialsImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestOsCredentialsImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestOsCredentialsImportModel(val *GuestOsCredentialsImportModel) *NullableGuestOsCredentialsImportModel {
	return &NullableGuestOsCredentialsImportModel{value: val, isSet: true}
}

func (v NullableGuestOsCredentialsImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestOsCredentialsImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


