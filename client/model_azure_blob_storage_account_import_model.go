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

// AzureBlobStorageAccountImportModel Account used to access Azure blob storage.
type AzureBlobStorageAccountImportModel struct {
	Credentials CredentialsImportModel `json:"credentials"`
	RegionType EAzureRegionType `json:"regionType"`
	GatewayServer RepositoryShareGatewayImportSpec `json:"gatewayServer"`
}

// NewAzureBlobStorageAccountImportModel instantiates a new AzureBlobStorageAccountImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStorageAccountImportModel(credentials CredentialsImportModel, regionType EAzureRegionType, gatewayServer RepositoryShareGatewayImportSpec, ) *AzureBlobStorageAccountImportModel {
	this := AzureBlobStorageAccountImportModel{}
	this.Credentials = credentials
	this.RegionType = regionType
	this.GatewayServer = gatewayServer
	return &this
}

// NewAzureBlobStorageAccountImportModelWithDefaults instantiates a new AzureBlobStorageAccountImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStorageAccountImportModelWithDefaults() *AzureBlobStorageAccountImportModel {
	this := AzureBlobStorageAccountImportModel{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *AzureBlobStorageAccountImportModel) GetCredentials() CredentialsImportModel {
	if o == nil  {
		var ret CredentialsImportModel
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageAccountImportModel) GetCredentialsOk() (*CredentialsImportModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *AzureBlobStorageAccountImportModel) SetCredentials(v CredentialsImportModel) {
	o.Credentials = v
}

// GetRegionType returns the RegionType field value
func (o *AzureBlobStorageAccountImportModel) GetRegionType() EAzureRegionType {
	if o == nil  {
		var ret EAzureRegionType
		return ret
	}

	return o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageAccountImportModel) GetRegionTypeOk() (*EAzureRegionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionType, true
}

// SetRegionType sets field value
func (o *AzureBlobStorageAccountImportModel) SetRegionType(v EAzureRegionType) {
	o.RegionType = v
}

// GetGatewayServer returns the GatewayServer field value
func (o *AzureBlobStorageAccountImportModel) GetGatewayServer() RepositoryShareGatewayImportSpec {
	if o == nil  {
		var ret RepositoryShareGatewayImportSpec
		return ret
	}

	return o.GatewayServer
}

// GetGatewayServerOk returns a tuple with the GatewayServer field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageAccountImportModel) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GatewayServer, true
}

// SetGatewayServer sets field value
func (o *AzureBlobStorageAccountImportModel) SetGatewayServer(v RepositoryShareGatewayImportSpec) {
	o.GatewayServer = v
}

func (o AzureBlobStorageAccountImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if true {
		toSerialize["regionType"] = o.RegionType
	}
	if true {
		toSerialize["gatewayServer"] = o.GatewayServer
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobStorageAccountImportModel struct {
	value *AzureBlobStorageAccountImportModel
	isSet bool
}

func (v NullableAzureBlobStorageAccountImportModel) Get() *AzureBlobStorageAccountImportModel {
	return v.value
}

func (v *NullableAzureBlobStorageAccountImportModel) Set(val *AzureBlobStorageAccountImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStorageAccountImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStorageAccountImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStorageAccountImportModel(val *AzureBlobStorageAccountImportModel) *NullableAzureBlobStorageAccountImportModel {
	return &NullableAzureBlobStorageAccountImportModel{value: val, isSet: true}
}

func (v NullableAzureBlobStorageAccountImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStorageAccountImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


