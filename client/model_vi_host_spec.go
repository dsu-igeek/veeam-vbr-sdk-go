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

// ViHostSpec struct for ViHostSpec
type ViHostSpec struct {
	ManagedServerSpec
	// Port used to communicate with the server.
	Port *int32 `json:"port,omitempty"`
	// [Optional] Certificate thumbprint used to verify the server identity. For details on how to get the thumbprint, see [Get TLS Certificate or SSH Fingerprint](#operation/GetConnectionCertificate). 
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
}

// NewViHostSpec instantiates a new ViHostSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViHostSpec() *ViHostSpec {
	this := ViHostSpec{}
	return &this
}

// NewViHostSpecWithDefaults instantiates a new ViHostSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViHostSpecWithDefaults() *ViHostSpec {
	this := ViHostSpec{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ViHostSpec) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViHostSpec) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ViHostSpec) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ViHostSpec) SetPort(v int32) {
	o.Port = &v
}

// GetCertificateThumbprint returns the CertificateThumbprint field value if set, zero value otherwise.
func (o *ViHostSpec) GetCertificateThumbprint() string {
	if o == nil || o.CertificateThumbprint == nil {
		var ret string
		return ret
	}
	return *o.CertificateThumbprint
}

// GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViHostSpec) GetCertificateThumbprintOk() (*string, bool) {
	if o == nil || o.CertificateThumbprint == nil {
		return nil, false
	}
	return o.CertificateThumbprint, true
}

// HasCertificateThumbprint returns a boolean if a field has been set.
func (o *ViHostSpec) HasCertificateThumbprint() bool {
	if o != nil && o.CertificateThumbprint != nil {
		return true
	}

	return false
}

// SetCertificateThumbprint gets a reference to the given string and assigns it to the CertificateThumbprint field.
func (o *ViHostSpec) SetCertificateThumbprint(v string) {
	o.CertificateThumbprint = &v
}

func (o ViHostSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedServerSpec, errManagedServerSpec := json.Marshal(o.ManagedServerSpec)
	if errManagedServerSpec != nil {
		return []byte{}, errManagedServerSpec
	}
	errManagedServerSpec = json.Unmarshal([]byte(serializedManagedServerSpec), &toSerialize)
	if errManagedServerSpec != nil {
		return []byte{}, errManagedServerSpec
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.CertificateThumbprint != nil {
		toSerialize["certificateThumbprint"] = o.CertificateThumbprint
	}
	return json.Marshal(toSerialize)
}

type NullableViHostSpec struct {
	value *ViHostSpec
	isSet bool
}

func (v NullableViHostSpec) Get() *ViHostSpec {
	return v.value
}

func (v *NullableViHostSpec) Set(val *ViHostSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableViHostSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableViHostSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViHostSpec(val *ViHostSpec) *NullableViHostSpec {
	return &NullableViHostSpec{value: val, isSet: true}
}

func (v NullableViHostSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViHostSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


