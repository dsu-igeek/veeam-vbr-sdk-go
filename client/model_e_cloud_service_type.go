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
	"fmt"
)

// ECloudServiceType the model 'ECloudServiceType'
type ECloudServiceType string

// List of ECloudServiceType
const (
	ECLOUDSERVICETYPE_AZURE_BLOB ECloudServiceType = "AzureBlob"
	ECLOUDSERVICETYPE_AZURE_DATA_BOX ECloudServiceType = "AzureDataBox"
	ECLOUDSERVICETYPE_AMAZON_S3 ECloudServiceType = "AmazonS3"
	ECLOUDSERVICETYPE_AMAZON_EC2 ECloudServiceType = "AmazonEC2"
	ECLOUDSERVICETYPE_S3_COMPATIBLE ECloudServiceType = "S3Compatible"
	ECLOUDSERVICETYPE_AMAZON_SNOWBALL_EDGE ECloudServiceType = "AmazonSnowballEdge"
	ECLOUDSERVICETYPE_GOOGLE_CLOUD ECloudServiceType = "GoogleCloud"
	ECLOUDSERVICETYPE_IBM_CLOUD ECloudServiceType = "IBMCloud"
	ECLOUDSERVICETYPE_AZURE_COMPUTE ECloudServiceType = "AzureCompute"
	ECLOUDSERVICETYPE_WASABI_CLOUD ECloudServiceType = "WasabiCloud"
)

func (v *ECloudServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ECloudServiceType(value)
	for _, existing := range []ECloudServiceType{ "AzureBlob", "AzureDataBox", "AmazonS3", "AmazonEC2", "S3Compatible", "AmazonSnowballEdge", "GoogleCloud", "IBMCloud", "AzureCompute", "WasabiCloud",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ECloudServiceType", value)
}

// Ptr returns reference to ECloudServiceType value
func (v ECloudServiceType) Ptr() *ECloudServiceType {
	return &v
}

type NullableECloudServiceType struct {
	value *ECloudServiceType
	isSet bool
}

func (v NullableECloudServiceType) Get() *ECloudServiceType {
	return v.value
}

func (v *NullableECloudServiceType) Set(val *ECloudServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableECloudServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableECloudServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECloudServiceType(val *ECloudServiceType) *NullableECloudServiceType {
	return &NullableECloudServiceType{value: val, isSet: true}
}

func (v NullableECloudServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECloudServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

