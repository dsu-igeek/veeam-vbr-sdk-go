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

// BackupApplicationSettingsImportModel struct for BackupApplicationSettingsImportModel
type BackupApplicationSettingsImportModel struct {
	VmObject VmwareObjectModel `json:"vmObject"`
	Vss EApplicationSettingsVSS `json:"vss"`
	// If *true*, persistent guest agent is used.
	UsePersistentGuestAgent *bool `json:"usePersistentGuestAgent,omitempty"`
	TransactionLogs *ETransactionLogsSettings `json:"transactionLogs,omitempty"`
	Sql *BackupSQLSettingsImportModel `json:"sql,omitempty"`
	Oracle *BackupOracleSettingsImportModel `json:"oracle,omitempty"`
	Exclusions *BackupFSExclusionsModel `json:"exclusions,omitempty"`
	Scripts *BackupScriptSettingsModel `json:"scripts,omitempty"`
}

// NewBackupApplicationSettingsImportModel instantiates a new BackupApplicationSettingsImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupApplicationSettingsImportModel(vmObject VmwareObjectModel, vss EApplicationSettingsVSS, ) *BackupApplicationSettingsImportModel {
	this := BackupApplicationSettingsImportModel{}
	this.VmObject = vmObject
	this.Vss = vss
	return &this
}

// NewBackupApplicationSettingsImportModelWithDefaults instantiates a new BackupApplicationSettingsImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupApplicationSettingsImportModelWithDefaults() *BackupApplicationSettingsImportModel {
	this := BackupApplicationSettingsImportModel{}
	return &this
}

// GetVmObject returns the VmObject field value
func (o *BackupApplicationSettingsImportModel) GetVmObject() VmwareObjectModel {
	if o == nil  {
		var ret VmwareObjectModel
		return ret
	}

	return o.VmObject
}

// GetVmObjectOk returns a tuple with the VmObject field value
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsImportModel) GetVmObjectOk() (*VmwareObjectModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VmObject, true
}

// SetVmObject sets field value
func (o *BackupApplicationSettingsImportModel) SetVmObject(v VmwareObjectModel) {
	o.VmObject = v
}

// GetVss returns the Vss field value
func (o *BackupApplicationSettingsImportModel) GetVss() EApplicationSettingsVSS {
	if o == nil  {
		var ret EApplicationSettingsVSS
		return ret
	}

	return o.Vss
}

// GetVssOk returns a tuple with the Vss field value
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsImportModel) GetVssOk() (*EApplicationSettingsVSS, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vss, true
}

// SetVss sets field value
func (o *BackupApplicationSettingsImportModel) SetVss(v EApplicationSettingsVSS) {
	o.Vss = v
}

// GetUsePersistentGuestAgent returns the UsePersistentGuestAgent field value if set, zero value otherwise.
func (o *BackupApplicationSettingsImportModel) GetUsePersistentGuestAgent() bool {
	if o == nil || o.UsePersistentGuestAgent == nil {
		var ret bool
		return ret
	}
	return *o.UsePersistentGuestAgent
}

// GetUsePersistentGuestAgentOk returns a tuple with the UsePersistentGuestAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsImportModel) GetUsePersistentGuestAgentOk() (*bool, bool) {
	if o == nil || o.UsePersistentGuestAgent == nil {
		return nil, false
	}
	return o.UsePersistentGuestAgent, true
}

// HasUsePersistentGuestAgent returns a boolean if a field has been set.
func (o *BackupApplicationSettingsImportModel) HasUsePersistentGuestAgent() bool {
	if o != nil && o.UsePersistentGuestAgent != nil {
		return true
	}

	return false
}

// SetUsePersistentGuestAgent gets a reference to the given bool and assigns it to the UsePersistentGuestAgent field.
func (o *BackupApplicationSettingsImportModel) SetUsePersistentGuestAgent(v bool) {
	o.UsePersistentGuestAgent = &v
}

// GetTransactionLogs returns the TransactionLogs field value if set, zero value otherwise.
func (o *BackupApplicationSettingsImportModel) GetTransactionLogs() ETransactionLogsSettings {
	if o == nil || o.TransactionLogs == nil {
		var ret ETransactionLogsSettings
		return ret
	}
	return *o.TransactionLogs
}

// GetTransactionLogsOk returns a tuple with the TransactionLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsImportModel) GetTransactionLogsOk() (*ETransactionLogsSettings, bool) {
	if o == nil || o.TransactionLogs == nil {
		return nil, false
	}
	return o.TransactionLogs, true
}

// HasTransactionLogs returns a boolean if a field has been set.
func (o *BackupApplicationSettingsImportModel) HasTransactionLogs() bool {
	if o != nil && o.TransactionLogs != nil {
		return true
	}

	return false
}

// SetTransactionLogs gets a reference to the given ETransactionLogsSettings and assigns it to the TransactionLogs field.
func (o *BackupApplicationSettingsImportModel) SetTransactionLogs(v ETransactionLogsSettings) {
	o.TransactionLogs = &v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *BackupApplicationSettingsImportModel) GetSql() BackupSQLSettingsImportModel {
	if o == nil || o.Sql == nil {
		var ret BackupSQLSettingsImportModel
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsImportModel) GetSqlOk() (*BackupSQLSettingsImportModel, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *BackupApplicationSettingsImportModel) HasSql() bool {
	if o != nil && o.Sql != nil {
		return true
	}

	return false
}

// SetSql gets a reference to the given BackupSQLSettingsImportModel and assigns it to the Sql field.
func (o *BackupApplicationSettingsImportModel) SetSql(v BackupSQLSettingsImportModel) {
	o.Sql = &v
}

// GetOracle returns the Oracle field value if set, zero value otherwise.
func (o *BackupApplicationSettingsImportModel) GetOracle() BackupOracleSettingsImportModel {
	if o == nil || o.Oracle == nil {
		var ret BackupOracleSettingsImportModel
		return ret
	}
	return *o.Oracle
}

// GetOracleOk returns a tuple with the Oracle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsImportModel) GetOracleOk() (*BackupOracleSettingsImportModel, bool) {
	if o == nil || o.Oracle == nil {
		return nil, false
	}
	return o.Oracle, true
}

// HasOracle returns a boolean if a field has been set.
func (o *BackupApplicationSettingsImportModel) HasOracle() bool {
	if o != nil && o.Oracle != nil {
		return true
	}

	return false
}

// SetOracle gets a reference to the given BackupOracleSettingsImportModel and assigns it to the Oracle field.
func (o *BackupApplicationSettingsImportModel) SetOracle(v BackupOracleSettingsImportModel) {
	o.Oracle = &v
}

// GetExclusions returns the Exclusions field value if set, zero value otherwise.
func (o *BackupApplicationSettingsImportModel) GetExclusions() BackupFSExclusionsModel {
	if o == nil || o.Exclusions == nil {
		var ret BackupFSExclusionsModel
		return ret
	}
	return *o.Exclusions
}

// GetExclusionsOk returns a tuple with the Exclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsImportModel) GetExclusionsOk() (*BackupFSExclusionsModel, bool) {
	if o == nil || o.Exclusions == nil {
		return nil, false
	}
	return o.Exclusions, true
}

// HasExclusions returns a boolean if a field has been set.
func (o *BackupApplicationSettingsImportModel) HasExclusions() bool {
	if o != nil && o.Exclusions != nil {
		return true
	}

	return false
}

// SetExclusions gets a reference to the given BackupFSExclusionsModel and assigns it to the Exclusions field.
func (o *BackupApplicationSettingsImportModel) SetExclusions(v BackupFSExclusionsModel) {
	o.Exclusions = &v
}

// GetScripts returns the Scripts field value if set, zero value otherwise.
func (o *BackupApplicationSettingsImportModel) GetScripts() BackupScriptSettingsModel {
	if o == nil || o.Scripts == nil {
		var ret BackupScriptSettingsModel
		return ret
	}
	return *o.Scripts
}

// GetScriptsOk returns a tuple with the Scripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationSettingsImportModel) GetScriptsOk() (*BackupScriptSettingsModel, bool) {
	if o == nil || o.Scripts == nil {
		return nil, false
	}
	return o.Scripts, true
}

// HasScripts returns a boolean if a field has been set.
func (o *BackupApplicationSettingsImportModel) HasScripts() bool {
	if o != nil && o.Scripts != nil {
		return true
	}

	return false
}

// SetScripts gets a reference to the given BackupScriptSettingsModel and assigns it to the Scripts field.
func (o *BackupApplicationSettingsImportModel) SetScripts(v BackupScriptSettingsModel) {
	o.Scripts = &v
}

func (o BackupApplicationSettingsImportModel) MarshalJSON() ([]byte, error) {
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

type NullableBackupApplicationSettingsImportModel struct {
	value *BackupApplicationSettingsImportModel
	isSet bool
}

func (v NullableBackupApplicationSettingsImportModel) Get() *BackupApplicationSettingsImportModel {
	return v.value
}

func (v *NullableBackupApplicationSettingsImportModel) Set(val *BackupApplicationSettingsImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupApplicationSettingsImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupApplicationSettingsImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupApplicationSettingsImportModel(val *BackupApplicationSettingsImportModel) *NullableBackupApplicationSettingsImportModel {
	return &NullableBackupApplicationSettingsImportModel{value: val, isSet: true}
}

func (v NullableBackupApplicationSettingsImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupApplicationSettingsImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


