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
	"time"
)

// ServerTimeModel struct for ServerTimeModel
type ServerTimeModel struct {
	// Current date and time on the server.
	ServerTime time.Time `json:"serverTime"`
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewServerTimeModel instantiates a new ServerTimeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerTimeModel(serverTime time.Time, ) *ServerTimeModel {
	this := ServerTimeModel{}
	this.ServerTime = serverTime
	return &this
}

// NewServerTimeModelWithDefaults instantiates a new ServerTimeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerTimeModelWithDefaults() *ServerTimeModel {
	this := ServerTimeModel{}
	return &this
}

// GetServerTime returns the ServerTime field value
func (o *ServerTimeModel) GetServerTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value
// and a boolean to check if the value has been set.
func (o *ServerTimeModel) GetServerTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServerTime, true
}

// SetServerTime sets field value
func (o *ServerTimeModel) SetServerTime(v time.Time) {
	o.ServerTime = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ServerTimeModel) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerTimeModel) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ServerTimeModel) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ServerTimeModel) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o ServerTimeModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serverTime"] = o.ServerTime
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	return json.Marshal(toSerialize)
}

type NullableServerTimeModel struct {
	value *ServerTimeModel
	isSet bool
}

func (v NullableServerTimeModel) Get() *ServerTimeModel {
	return v.value
}

func (v *NullableServerTimeModel) Set(val *ServerTimeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableServerTimeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableServerTimeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerTimeModel(val *ServerTimeModel) *NullableServerTimeModel {
	return &NullableServerTimeModel{value: val, isSet: true}
}

func (v NullableServerTimeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerTimeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


