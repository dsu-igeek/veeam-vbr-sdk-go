/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev2
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EMonth the model 'EMonth'
type EMonth string

// List of EMonth
const (
	EMONTH_JANUARY EMonth = "January"
	EMONTH_FEBRUARY EMonth = "February"
	EMONTH_MARCH EMonth = "March"
	EMONTH_APRIL EMonth = "April"
	EMONTH_MAY EMonth = "May"
	EMONTH_JUNE EMonth = "June"
	EMONTH_JULY EMonth = "July"
	EMONTH_AUGUST EMonth = "August"
	EMONTH_SEPTEMBER EMonth = "September"
	EMONTH_OCTOBER EMonth = "October"
	EMONTH_NOVEMBER EMonth = "November"
	EMONTH_DECEMBER EMonth = "December"
)

func (v *EMonth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EMonth(value)
	for _, existing := range []EMonth{ "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EMonth", value)
}

// Ptr returns reference to EMonth value
func (v EMonth) Ptr() *EMonth {
	return &v
}

type NullableEMonth struct {
	value *EMonth
	isSet bool
}

func (v NullableEMonth) Get() *EMonth {
	return v.value
}

func (v *NullableEMonth) Set(val *EMonth) {
	v.value = val
	v.isSet = true
}

func (v NullableEMonth) IsSet() bool {
	return v.isSet
}

func (v *NullableEMonth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEMonth(val *EMonth) *NullableEMonth {
	return &NullableEMonth{value: val, isSet: true}
}

func (v NullableEMonth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEMonth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

