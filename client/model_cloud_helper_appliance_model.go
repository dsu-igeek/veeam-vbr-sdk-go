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

// CloudHelperApplianceModel struct for CloudHelperApplianceModel
type CloudHelperApplianceModel struct {
	Type ECloudCredentialsType `json:"type"`
}

// NewCloudHelperApplianceModel instantiates a new CloudHelperApplianceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudHelperApplianceModel(type_ ECloudCredentialsType, ) *CloudHelperApplianceModel {
	this := CloudHelperApplianceModel{}
	this.Type = type_
	return &this
}

// NewCloudHelperApplianceModelWithDefaults instantiates a new CloudHelperApplianceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudHelperApplianceModelWithDefaults() *CloudHelperApplianceModel {
	this := CloudHelperApplianceModel{}
	return &this
}

// GetType returns the Type field value
func (o *CloudHelperApplianceModel) GetType() ECloudCredentialsType {
	if o == nil  {
		var ret ECloudCredentialsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudHelperApplianceModel) GetTypeOk() (*ECloudCredentialsType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudHelperApplianceModel) SetType(v ECloudCredentialsType) {
	o.Type = v
}

func (o CloudHelperApplianceModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCloudHelperApplianceModel struct {
	value *CloudHelperApplianceModel
	isSet bool
}

func (v NullableCloudHelperApplianceModel) Get() *CloudHelperApplianceModel {
	return v.value
}

func (v *NullableCloudHelperApplianceModel) Set(val *CloudHelperApplianceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudHelperApplianceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudHelperApplianceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudHelperApplianceModel(val *CloudHelperApplianceModel) *NullableCloudHelperApplianceModel {
	return &NullableCloudHelperApplianceModel{value: val, isSet: true}
}

func (v NullableCloudHelperApplianceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudHelperApplianceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


