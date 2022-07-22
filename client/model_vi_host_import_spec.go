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

// ViHostImportSpec struct for ViHostImportSpec
type ViHostImportSpec struct {
	// Full DNS name or IP address of the server.
	Name string `json:"name"`
	// Description of the server.
	Description string `json:"description"`
	Type EManagedServerType `json:"type"`
	ViHostType EViHostType `json:"viHostType"`
	Credentials CredentialsImportModel `json:"credentials"`
	// Port used to communicate with the server.
	Port *int32 `json:"port,omitempty"`
	// [Optional] Certificate thumbprint used to verify the server identity. 
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
}

// NewViHostImportSpec instantiates a new ViHostImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViHostImportSpec(name string, description string, type_ EManagedServerType, viHostType EViHostType, credentials CredentialsImportModel) *ViHostImportSpec {
	this := ViHostImportSpec{}
	this.Name = name
	this.Description = description
	this.Type = type_
	this.ViHostType = viHostType
	this.Credentials = credentials
	return &this
}

// NewViHostImportSpecWithDefaults instantiates a new ViHostImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViHostImportSpecWithDefaults() *ViHostImportSpec {
	this := ViHostImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *ViHostImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ViHostImportSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ViHostImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ViHostImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ViHostImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ViHostImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *ViHostImportSpec) GetType() EManagedServerType {
	if o == nil {
		var ret EManagedServerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ViHostImportSpec) GetTypeOk() (*EManagedServerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ViHostImportSpec) SetType(v EManagedServerType) {
	o.Type = v
}

// GetViHostType returns the ViHostType field value
func (o *ViHostImportSpec) GetViHostType() EViHostType {
	if o == nil {
		var ret EViHostType
		return ret
	}

	return o.ViHostType
}

// GetViHostTypeOk returns a tuple with the ViHostType field value
// and a boolean to check if the value has been set.
func (o *ViHostImportSpec) GetViHostTypeOk() (*EViHostType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViHostType, true
}

// SetViHostType sets field value
func (o *ViHostImportSpec) SetViHostType(v EViHostType) {
	o.ViHostType = v
}

// GetCredentials returns the Credentials field value
func (o *ViHostImportSpec) GetCredentials() CredentialsImportModel {
	if o == nil {
		var ret CredentialsImportModel
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *ViHostImportSpec) GetCredentialsOk() (*CredentialsImportModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *ViHostImportSpec) SetCredentials(v CredentialsImportModel) {
	o.Credentials = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ViHostImportSpec) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViHostImportSpec) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ViHostImportSpec) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ViHostImportSpec) SetPort(v int32) {
	o.Port = &v
}

// GetCertificateThumbprint returns the CertificateThumbprint field value if set, zero value otherwise.
func (o *ViHostImportSpec) GetCertificateThumbprint() string {
	if o == nil || o.CertificateThumbprint == nil {
		var ret string
		return ret
	}
	return *o.CertificateThumbprint
}

// GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViHostImportSpec) GetCertificateThumbprintOk() (*string, bool) {
	if o == nil || o.CertificateThumbprint == nil {
		return nil, false
	}
	return o.CertificateThumbprint, true
}

// HasCertificateThumbprint returns a boolean if a field has been set.
func (o *ViHostImportSpec) HasCertificateThumbprint() bool {
	if o != nil && o.CertificateThumbprint != nil {
		return true
	}

	return false
}

// SetCertificateThumbprint gets a reference to the given string and assigns it to the CertificateThumbprint field.
func (o *ViHostImportSpec) SetCertificateThumbprint(v string) {
	o.CertificateThumbprint = &v
}

func (o ViHostImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["viHostType"] = o.ViHostType
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.CertificateThumbprint != nil {
		toSerialize["certificateThumbprint"] = o.CertificateThumbprint
	}
	return json.Marshal(toSerialize)
}

type NullableViHostImportSpec struct {
	value *ViHostImportSpec
	isSet bool
}

func (v NullableViHostImportSpec) Get() *ViHostImportSpec {
	return v.value
}

func (v *NullableViHostImportSpec) Set(val *ViHostImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableViHostImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableViHostImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViHostImportSpec(val *ViHostImportSpec) *NullableViHostImportSpec {
	return &NullableViHostImportSpec{value: val, isSet: true}
}

func (v NullableViHostImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViHostImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


