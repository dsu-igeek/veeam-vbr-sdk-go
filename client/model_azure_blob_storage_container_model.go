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

// AzureBlobStorageContainerModel Azure Blob storage container.
type AzureBlobStorageContainerModel struct {
	// Name of the container that contains backups created with Veeam Backup for Microsoft Azure.
	ContainerName string `json:"containerName"`
	// Folder that contains backups created with Veeam Backup for Microsoft Azure.
	FolderName string `json:"folderName"`
	StorageConsumptionLimit *ObjectStorageConsumptionLimitModel `json:"storageConsumptionLimit,omitempty"`
}

// NewAzureBlobStorageContainerModel instantiates a new AzureBlobStorageContainerModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStorageContainerModel(containerName string, folderName string, ) *AzureBlobStorageContainerModel {
	this := AzureBlobStorageContainerModel{}
	this.ContainerName = containerName
	this.FolderName = folderName
	return &this
}

// NewAzureBlobStorageContainerModelWithDefaults instantiates a new AzureBlobStorageContainerModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStorageContainerModelWithDefaults() *AzureBlobStorageContainerModel {
	this := AzureBlobStorageContainerModel{}
	return &this
}

// GetContainerName returns the ContainerName field value
func (o *AzureBlobStorageContainerModel) GetContainerName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageContainerModel) GetContainerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value
func (o *AzureBlobStorageContainerModel) SetContainerName(v string) {
	o.ContainerName = v
}

// GetFolderName returns the FolderName field value
func (o *AzureBlobStorageContainerModel) GetFolderName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageContainerModel) GetFolderNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FolderName, true
}

// SetFolderName sets field value
func (o *AzureBlobStorageContainerModel) SetFolderName(v string) {
	o.FolderName = v
}

// GetStorageConsumptionLimit returns the StorageConsumptionLimit field value if set, zero value otherwise.
func (o *AzureBlobStorageContainerModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel {
	if o == nil || o.StorageConsumptionLimit == nil {
		var ret ObjectStorageConsumptionLimitModel
		return ret
	}
	return *o.StorageConsumptionLimit
}

// GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageContainerModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool) {
	if o == nil || o.StorageConsumptionLimit == nil {
		return nil, false
	}
	return o.StorageConsumptionLimit, true
}

// HasStorageConsumptionLimit returns a boolean if a field has been set.
func (o *AzureBlobStorageContainerModel) HasStorageConsumptionLimit() bool {
	if o != nil && o.StorageConsumptionLimit != nil {
		return true
	}

	return false
}

// SetStorageConsumptionLimit gets a reference to the given ObjectStorageConsumptionLimitModel and assigns it to the StorageConsumptionLimit field.
func (o *AzureBlobStorageContainerModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel) {
	o.StorageConsumptionLimit = &v
}

func (o AzureBlobStorageContainerModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["containerName"] = o.ContainerName
	}
	if true {
		toSerialize["folderName"] = o.FolderName
	}
	if o.StorageConsumptionLimit != nil {
		toSerialize["storageConsumptionLimit"] = o.StorageConsumptionLimit
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobStorageContainerModel struct {
	value *AzureBlobStorageContainerModel
	isSet bool
}

func (v NullableAzureBlobStorageContainerModel) Get() *AzureBlobStorageContainerModel {
	return v.value
}

func (v *NullableAzureBlobStorageContainerModel) Set(val *AzureBlobStorageContainerModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStorageContainerModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStorageContainerModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStorageContainerModel(val *AzureBlobStorageContainerModel) *NullableAzureBlobStorageContainerModel {
	return &NullableAzureBlobStorageContainerModel{value: val, isSet: true}
}

func (v NullableAzureBlobStorageContainerModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStorageContainerModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


