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

// ETransactionLogsSettings Transaction logs settings that define whether copy-only backups must be created, or transaction logs for Microsoft Exchange, Microsoft SQL and Oracle VMs must be processed.<br> If transaction log processing is selected, specify the following parameters:<ul> <li>[For Microsoft SQL Server VMs] Microsoft SQL Server transaction log settings</li> <li>[For Oracle VMs] Oracle archived log settings</li></ul>
type ETransactionLogsSettings string

// List of ETransactionLogsSettings
const (
	ETRANSACTIONLOGSSETTINGS_PROCESS ETransactionLogsSettings = "process"
	ETRANSACTIONLOGSSETTINGS_COPY_ONLY ETransactionLogsSettings = "copyOnly"
)

func (v *ETransactionLogsSettings) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ETransactionLogsSettings(value)
	for _, existing := range []ETransactionLogsSettings{ "process", "copyOnly",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ETransactionLogsSettings", value)
}

// Ptr returns reference to ETransactionLogsSettings value
func (v ETransactionLogsSettings) Ptr() *ETransactionLogsSettings {
	return &v
}

type NullableETransactionLogsSettings struct {
	value *ETransactionLogsSettings
	isSet bool
}

func (v NullableETransactionLogsSettings) Get() *ETransactionLogsSettings {
	return v.value
}

func (v *NullableETransactionLogsSettings) Set(val *ETransactionLogsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableETransactionLogsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableETransactionLogsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableETransactionLogsSettings(val *ETransactionLogsSettings) *NullableETransactionLogsSettings {
	return &NullableETransactionLogsSettings{value: val, isSet: true}
}

func (v NullableETransactionLogsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableETransactionLogsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

