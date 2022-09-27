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

// SessionResultModel struct for SessionResultModel
type SessionResultModel struct {
	Result ESessionResult `json:"result"`
	// Message that explains the session result.
	Message *string `json:"message,omitempty"`
	// If *true*, the session has been canceled.
	IsCanceled *bool `json:"isCanceled,omitempty"`
}

// NewSessionResultModel instantiates a new SessionResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionResultModel(result ESessionResult, ) *SessionResultModel {
	this := SessionResultModel{}
	this.Result = result
	return &this
}

// NewSessionResultModelWithDefaults instantiates a new SessionResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionResultModelWithDefaults() *SessionResultModel {
	this := SessionResultModel{}
	return &this
}

// GetResult returns the Result field value
func (o *SessionResultModel) GetResult() ESessionResult {
	if o == nil  {
		var ret ESessionResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *SessionResultModel) GetResultOk() (*ESessionResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *SessionResultModel) SetResult(v ESessionResult) {
	o.Result = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SessionResultModel) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResultModel) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SessionResultModel) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SessionResultModel) SetMessage(v string) {
	o.Message = &v
}

// GetIsCanceled returns the IsCanceled field value if set, zero value otherwise.
func (o *SessionResultModel) GetIsCanceled() bool {
	if o == nil || o.IsCanceled == nil {
		var ret bool
		return ret
	}
	return *o.IsCanceled
}

// GetIsCanceledOk returns a tuple with the IsCanceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResultModel) GetIsCanceledOk() (*bool, bool) {
	if o == nil || o.IsCanceled == nil {
		return nil, false
	}
	return o.IsCanceled, true
}

// HasIsCanceled returns a boolean if a field has been set.
func (o *SessionResultModel) HasIsCanceled() bool {
	if o != nil && o.IsCanceled != nil {
		return true
	}

	return false
}

// SetIsCanceled gets a reference to the given bool and assigns it to the IsCanceled field.
func (o *SessionResultModel) SetIsCanceled(v bool) {
	o.IsCanceled = &v
}

func (o SessionResultModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["result"] = o.Result
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.IsCanceled != nil {
		toSerialize["isCanceled"] = o.IsCanceled
	}
	return json.Marshal(toSerialize)
}

type NullableSessionResultModel struct {
	value *SessionResultModel
	isSet bool
}

func (v NullableSessionResultModel) Get() *SessionResultModel {
	return v.value
}

func (v *NullableSessionResultModel) Set(val *SessionResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionResultModel(val *SessionResultModel) *NullableSessionResultModel {
	return &NullableSessionResultModel{value: val, isSet: true}
}

func (v NullableSessionResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


