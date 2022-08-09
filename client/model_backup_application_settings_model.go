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
)

// BackupApplicationSettingsModel struct for BackupApplicationSettingsModel
type BackupApplicationSettingsModel struct {
	VmObject VmwareObjectModel `json:"vmObject"`
	Vss EApplicationSettingsVSS `json:"vss"`
	// If *true*, persistent guest agent is used.
	UsePersistentGuestAgent *bool `json:"usePersistentGuestAgent,omitempty"`
	TransactionLogs *ETransactionLogsSettings `json:"transactionLogs,omitempty"`
	Sql *BackupSQLSettingsModel `json:"sql,omitempty"`
	Oracle *BackupOracleSettingsModel `json:"oracle,omitempty"`
	Exclusions *BackupFSExclusionsModel `json:"exclusions,omitempty"`
	Scripts *BackupScriptSettingsModel `json:"scripts,omitempty"`
}

// NewBackupApplicationSettingsModel instantiates a new BackupApplicationSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupApplicationSettingsModel(vmObject VmwareObjectModel, vss EApplicationSettingsVSS, ) *BackupApplicationSettingsModel {
	this := BackupApplicationSettingsModel{}
	this.VmObject = vmObject
	this.Vss = vss
	return &this
}

// NewBackupApplicationSettingsModelWithDefaults instantiates a new BackupApplicationSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupApplicationSettingsModelWithDefaults() *BackupApplicationSettingsModel {
	this := BackupApplicationSettingsModel{}
	return &this
}

// GetVmObject returns the VmObject field value
func (o *BackupApplicationSettingsModel) GetVmObject() VmwareObjectModel {
	if o == nil  {
		var ret VmwareObjectModel
		return ret
	}

	return o.VmObject
}

// GetVmObjectOk returns a tuple with the VmObject field value
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsModel) GetVmObjectOk() (*VmwareObjectModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VmObject, true
}

// SetVmObject sets field value
func (o *BackupApplicationSettingsModel) SetVmObject(v VmwareObjectModel) {
	o.VmObject = v
}

// GetVss returns the Vss field value
func (o *BackupApplicationSettingsModel) GetVss() EApplicationSettingsVSS {
	if o == nil  {
		var ret EApplicationSettingsVSS
		return ret
	}

	return o.Vss
}

// GetVssOk returns a tuple with the Vss field value
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsModel) GetVssOk() (*EApplicationSettingsVSS, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vss, true
}

// SetVss sets field value
func (o *BackupApplicationSettingsModel) SetVss(v EApplicationSettingsVSS) {
	o.Vss = v
}

// GetUsePersistentGuestAgent returns the UsePersistentGuestAgent field value if set, zero value otherwise.
func (o *BackupApplicationSettingsModel) GetUsePersistentGuestAgent() bool {
	if o == nil || o.UsePersistentGuestAgent == nil {
		var ret bool
		return ret
	}
	return *o.UsePersistentGuestAgent
}

// GetUsePersistentGuestAgentOk returns a tuple with the UsePersistentGuestAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsModel) GetUsePersistentGuestAgentOk() (*bool, bool) {
	if o == nil || o.UsePersistentGuestAgent == nil {
		return nil, false
	}
	return o.UsePersistentGuestAgent, true
}

// HasUsePersistentGuestAgent returns a boolean if a field has been set.
func (o *BackupApplicationSettingsModel) HasUsePersistentGuestAgent() bool {
	if o != nil && o.UsePersistentGuestAgent != nil {
		return true
	}

	return false
}

// SetUsePersistentGuestAgent gets a reference to the given bool and assigns it to the UsePersistentGuestAgent field.
func (o *BackupApplicationSettingsModel) SetUsePersistentGuestAgent(v bool) {
	o.UsePersistentGuestAgent = &v
}

// GetTransactionLogs returns the TransactionLogs field value if set, zero value otherwise.
func (o *BackupApplicationSettingsModel) GetTransactionLogs() ETransactionLogsSettings {
	if o == nil || o.TransactionLogs == nil {
		var ret ETransactionLogsSettings
		return ret
	}
	return *o.TransactionLogs
}

// GetTransactionLogsOk returns a tuple with the TransactionLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsModel) GetTransactionLogsOk() (*ETransactionLogsSettings, bool) {
	if o == nil || o.TransactionLogs == nil {
		return nil, false
	}
	return o.TransactionLogs, true
}

// HasTransactionLogs returns a boolean if a field has been set.
func (o *BackupApplicationSettingsModel) HasTransactionLogs() bool {
	if o != nil && o.TransactionLogs != nil {
		return true
	}

	return false
}

// SetTransactionLogs gets a reference to the given ETransactionLogsSettings and assigns it to the TransactionLogs field.
func (o *BackupApplicationSettingsModel) SetTransactionLogs(v ETransactionLogsSettings) {
	o.TransactionLogs = &v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *BackupApplicationSettingsModel) GetSql() BackupSQLSettingsModel {
	if o == nil || o.Sql == nil {
		var ret BackupSQLSettingsModel
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsModel) GetSqlOk() (*BackupSQLSettingsModel, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *BackupApplicationSettingsModel) HasSql() bool {
	if o != nil && o.Sql != nil {
		return true
	}

	return false
}

// SetSql gets a reference to the given BackupSQLSettingsModel and assigns it to the Sql field.
func (o *BackupApplicationSettingsModel) SetSql(v BackupSQLSettingsModel) {
	o.Sql = &v
}

// GetOracle returns the Oracle field value if set, zero value otherwise.
func (o *BackupApplicationSettingsModel) GetOracle() BackupOracleSettingsModel {
	if o == nil || o.Oracle == nil {
		var ret BackupOracleSettingsModel
		return ret
	}
	return *o.Oracle
}

// GetOracleOk returns a tuple with the Oracle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsModel) GetOracleOk() (*BackupOracleSettingsModel, bool) {
	if o == nil || o.Oracle == nil {
		return nil, false
	}
	return o.Oracle, true
}

// HasOracle returns a boolean if a field has been set.
func (o *BackupApplicationSettingsModel) HasOracle() bool {
	if o != nil && o.Oracle != nil {
		return true
	}

	return false
}

// SetOracle gets a reference to the given BackupOracleSettingsModel and assigns it to the Oracle field.
func (o *BackupApplicationSettingsModel) SetOracle(v BackupOracleSettingsModel) {
	o.Oracle = &v
}

// GetExclusions returns the Exclusions field value if set, zero value otherwise.
func (o *BackupApplicationSettingsModel) GetExclusions() BackupFSExclusionsModel {
	if o == nil || o.Exclusions == nil {
		var ret BackupFSExclusionsModel
		return ret
	}
	return *o.Exclusions
}

// GetExclusionsOk returns a tuple with the Exclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsModel) GetExclusionsOk() (*BackupFSExclusionsModel, bool) {
	if o == nil || o.Exclusions == nil {
		return nil, false
	}
	return o.Exclusions, true
}

// HasExclusions returns a boolean if a field has been set.
func (o *BackupApplicationSettingsModel) HasExclusions() bool {
	if o != nil && o.Exclusions != nil {
		return true
	}

	return false
}

// SetExclusions gets a reference to the given BackupFSExclusionsModel and assigns it to the Exclusions field.
func (o *BackupApplicationSettingsModel) SetExclusions(v BackupFSExclusionsModel) {
	o.Exclusions = &v
}

// GetScripts returns the Scripts field value if set, zero value otherwise.
func (o *BackupApplicationSettingsModel) GetScripts() BackupScriptSettingsModel {
	if o == nil || o.Scripts == nil {
		var ret BackupScriptSettingsModel
		return ret
	}
	return *o.Scripts
}

// GetScriptsOk returns a tuple with the Scripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsModel) GetScriptsOk() (*BackupScriptSettingsModel, bool) {
	if o == nil || o.Scripts == nil {
		return nil, false
	}
	return o.Scripts, true
}

// HasScripts returns a boolean if a field has been set.
func (o *BackupApplicationSettingsModel) HasScripts() bool {
	if o != nil && o.Scripts != nil {
		return true
	}

	return false
}

// SetScripts gets a reference to the given BackupScriptSettingsModel and assigns it to the Scripts field.
func (o *BackupApplicationSettingsModel) SetScripts(v BackupScriptSettingsModel) {
	o.Scripts = &v
}

func (o BackupApplicationSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vmObject"] = o.VmObject
	}
	if true {
		toSerialize["vss"] = o.Vss
	}
	if o.UsePersistentGuestAgent != nil {
		toSerialize["usePersistentGuestAgent"] = o.UsePersistentGuestAgent
	}
	if o.TransactionLogs != nil {
		toSerialize["transactionLogs"] = o.TransactionLogs
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.Oracle != nil {
		toSerialize["oracle"] = o.Oracle
	}
	if o.Exclusions != nil {
		toSerialize["exclusions"] = o.Exclusions
	}
	if o.Scripts != nil {
		toSerialize["scripts"] = o.Scripts
	}
	return json.Marshal(toSerialize)
}

type NullableBackupApplicationSettingsModel struct {
	value *BackupApplicationSettingsModel
	isSet bool
}

func (v NullableBackupApplicationSettingsModel) Get() *BackupApplicationSettingsModel {
	return v.value
}

func (v *NullableBackupApplicationSettingsModel) Set(val *BackupApplicationSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupApplicationSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupApplicationSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupApplicationSettingsModel(val *BackupApplicationSettingsModel) *NullableBackupApplicationSettingsModel {
	return &NullableBackupApplicationSettingsModel{value: val, isSet: true}
}

func (v NullableBackupApplicationSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupApplicationSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


