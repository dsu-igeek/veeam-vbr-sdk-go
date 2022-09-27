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

// S3CompatibleStorageAccountImportModel struct for S3CompatibleStorageAccountImportModel
type S3CompatibleStorageAccountImportModel struct {
	ServicePoint string `json:"servicePoint"`
	RegionId string `json:"regionId"`
	Credentials CredentialsImportModel `json:"credentials"`
	GatewayServer RepositoryShareGatewayImportSpec `json:"gatewayServer"`
}

// NewS3CompatibleStorageAccountImportModel instantiates a new S3CompatibleStorageAccountImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3CompatibleStorageAccountImportModel(servicePoint string, regionId string, credentials CredentialsImportModel, gatewayServer RepositoryShareGatewayImportSpec, ) *S3CompatibleStorageAccountImportModel {
	this := S3CompatibleStorageAccountImportModel{}
	this.ServicePoint = servicePoint
	this.RegionId = regionId
	this.Credentials = credentials
	this.GatewayServer = gatewayServer
	return &this
}

// NewS3CompatibleStorageAccountImportModelWithDefaults instantiates a new S3CompatibleStorageAccountImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3CompatibleStorageAccountImportModelWithDefaults() *S3CompatibleStorageAccountImportModel {
	this := S3CompatibleStorageAccountImportModel{}
	return &this
}

// GetServicePoint returns the ServicePoint field value
func (o *S3CompatibleStorageAccountImportModel) GetServicePoint() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServicePoint
}

// GetServicePointOk returns a tuple with the ServicePoint field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageAccountImportModel) GetServicePointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServicePoint, true
}

// SetServicePoint sets field value
func (o *S3CompatibleStorageAccountImportModel) SetServicePoint(v string) {
	o.ServicePoint = v
}

// GetRegionId returns the RegionId field value
func (o *S3CompatibleStorageAccountImportModel) GetRegionId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageAccountImportModel) GetRegionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *S3CompatibleStorageAccountImportModel) SetRegionId(v string) {
	o.RegionId = v
}

// GetCredentials returns the Credentials field value
func (o *S3CompatibleStorageAccountImportModel) GetCredentials() CredentialsImportModel {
	if o == nil  {
		var ret CredentialsImportModel
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageAccountImportModel) GetCredentialsOk() (*CredentialsImportModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *S3CompatibleStorageAccountImportModel) SetCredentials(v CredentialsImportModel) {
	o.Credentials = v
}

// GetGatewayServer returns the GatewayServer field value
func (o *S3CompatibleStorageAccountImportModel) GetGatewayServer() RepositoryShareGatewayImportSpec {
	if o == nil  {
		var ret RepositoryShareGatewayImportSpec
		return ret
	}

	return o.GatewayServer
}

// GetGatewayServerOk returns a tuple with the GatewayServer field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageAccountImportModel) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GatewayServer, true
}

// SetGatewayServer sets field value
func (o *S3CompatibleStorageAccountImportModel) SetGatewayServer(v RepositoryShareGatewayImportSpec) {
	o.GatewayServer = v
}

func (o S3CompatibleStorageAccountImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servicePoint"] = o.ServicePoint
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if true {
		toSerialize["gatewayServer"] = o.GatewayServer
	}
	return json.Marshal(toSerialize)
}

type NullableS3CompatibleStorageAccountImportModel struct {
	value *S3CompatibleStorageAccountImportModel
	isSet bool
}

func (v NullableS3CompatibleStorageAccountImportModel) Get() *S3CompatibleStorageAccountImportModel {
	return v.value
}

func (v *NullableS3CompatibleStorageAccountImportModel) Set(val *S3CompatibleStorageAccountImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableS3CompatibleStorageAccountImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableS3CompatibleStorageAccountImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3CompatibleStorageAccountImportModel(val *S3CompatibleStorageAccountImportModel) *NullableS3CompatibleStorageAccountImportModel {
	return &NullableS3CompatibleStorageAccountImportModel{value: val, isSet: true}
}

func (v NullableS3CompatibleStorageAccountImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3CompatibleStorageAccountImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


