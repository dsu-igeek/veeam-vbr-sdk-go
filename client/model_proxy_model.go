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
	"fmt"
)

// ProxyModel - struct for ProxyModel
type ProxyModel struct {
	ViProxyModel *ViProxyModel
}

// ViProxyModelAsProxyModel is a convenience function that returns ViProxyModel wrapped in ProxyModel
func ViProxyModelAsProxyModel(v *ViProxyModel) ProxyModel {
	return ProxyModel{
		ViProxyModel: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProxyModel) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ViProxyModel
	err = newStrictDecoder(data).Decode(&dst.ViProxyModel)
	if err == nil {
		jsonViProxyModel, _ := json.Marshal(dst.ViProxyModel)
		if string(jsonViProxyModel) == "{}" { // empty struct
			dst.ViProxyModel = nil
		} else {
			match++
		}
	} else {
		dst.ViProxyModel = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ViProxyModel = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ProxyModel)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ProxyModel)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProxyModel) MarshalJSON() ([]byte, error) {
	if src.ViProxyModel != nil {
		return json.Marshal(&src.ViProxyModel)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProxyModel) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ViProxyModel != nil {
		return obj.ViProxyModel
	}

	// all schemas are nil
	return nil
}

type NullableProxyModel struct {
	value *ProxyModel
	isSet bool
}

func (v NullableProxyModel) Get() *ProxyModel {
	return v.value
}

func (v *NullableProxyModel) Set(val *ProxyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyModel(val *ProxyModel) *NullableProxyModel {
	return &NullableProxyModel{value: val, isSet: true}
}

func (v NullableProxyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


