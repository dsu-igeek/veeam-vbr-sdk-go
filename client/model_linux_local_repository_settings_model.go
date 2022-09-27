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

// LinuxLocalRepositorySettingsModel Repository settings.
type LinuxLocalRepositorySettingsModel struct {
	// Path to the folder where backup files are stored.
	Path *string `json:"path,omitempty"`
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	EnableReadWriteLimit *bool `json:"enableReadWriteLimit,omitempty"`
	// Maximum rate that restricts the total speed of reading and writing data to the backup repository disk.
	ReadWriteRate *int32 `json:"readWriteRate,omitempty"`
	// (For Linux repository) If *true*, fast cloning on XFS volumes is used.
	UseFastCloningOnXFSVolumes *bool `json:"useFastCloningOnXFSVolumes,omitempty"`
	// If *true*, the Object Lock feature is used to protect recent backups.
	UseImmutableBackups *bool `json:"useImmutableBackups,omitempty"`
	// Number of days to keep immutable backups.
	MakeRecentBackupsImmutableDays *int32 `json:"makeRecentBackupsImmutableDays,omitempty"`
	AdvancedSettings *RepositoryAdvancedSettingsModel `json:"advancedSettings,omitempty"`
}

// NewLinuxLocalRepositorySettingsModel instantiates a new LinuxLocalRepositorySettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxLocalRepositorySettingsModel() *LinuxLocalRepositorySettingsModel {
	this := LinuxLocalRepositorySettingsModel{}
	return &this
}

// NewLinuxLocalRepositorySettingsModelWithDefaults instantiates a new LinuxLocalRepositorySettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxLocalRepositorySettingsModelWithDefaults() *LinuxLocalRepositorySettingsModel {
	this := LinuxLocalRepositorySettingsModel{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *LinuxLocalRepositorySettingsModel) SetPath(v string) {
	o.Path = &v
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetEnableTaskLimit() bool {
	if o == nil || o.EnableTaskLimit == nil {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || o.EnableTaskLimit == nil {
		return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasEnableTaskLimit() bool {
	if o != nil && o.EnableTaskLimit != nil {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *LinuxLocalRepositorySettingsModel) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetMaxTaskCount() int32 {
	if o == nil || o.MaxTaskCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || o.MaxTaskCount == nil {
		return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasMaxTaskCount() bool {
	if o != nil && o.MaxTaskCount != nil {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *LinuxLocalRepositorySettingsModel) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetEnableReadWriteLimit returns the EnableReadWriteLimit field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetEnableReadWriteLimit() bool {
	if o == nil || o.EnableReadWriteLimit == nil {
		var ret bool
		return ret
	}
	return *o.EnableReadWriteLimit
}

// GetEnableReadWriteLimitOk returns a tuple with the EnableReadWriteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetEnableReadWriteLimitOk() (*bool, bool) {
	if o == nil || o.EnableReadWriteLimit == nil {
		return nil, false
	}
	return o.EnableReadWriteLimit, true
}

// HasEnableReadWriteLimit returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasEnableReadWriteLimit() bool {
	if o != nil && o.EnableReadWriteLimit != nil {
		return true
	}

	return false
}

// SetEnableReadWriteLimit gets a reference to the given bool and assigns it to the EnableReadWriteLimit field.
func (o *LinuxLocalRepositorySettingsModel) SetEnableReadWriteLimit(v bool) {
	o.EnableReadWriteLimit = &v
}

// GetReadWriteRate returns the ReadWriteRate field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetReadWriteRate() int32 {
	if o == nil || o.ReadWriteRate == nil {
		var ret int32
		return ret
	}
	return *o.ReadWriteRate
}

// GetReadWriteRateOk returns a tuple with the ReadWriteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetReadWriteRateOk() (*int32, bool) {
	if o == nil || o.ReadWriteRate == nil {
		return nil, false
	}
	return o.ReadWriteRate, true
}

// HasReadWriteRate returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasReadWriteRate() bool {
	if o != nil && o.ReadWriteRate != nil {
		return true
	}

	return false
}

// SetReadWriteRate gets a reference to the given int32 and assigns it to the ReadWriteRate field.
func (o *LinuxLocalRepositorySettingsModel) SetReadWriteRate(v int32) {
	o.ReadWriteRate = &v
}

// GetUseFastCloningOnXFSVolumes returns the UseFastCloningOnXFSVolumes field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetUseFastCloningOnXFSVolumes() bool {
	if o == nil || o.UseFastCloningOnXFSVolumes == nil {
		var ret bool
		return ret
	}
	return *o.UseFastCloningOnXFSVolumes
}

// GetUseFastCloningOnXFSVolumesOk returns a tuple with the UseFastCloningOnXFSVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetUseFastCloningOnXFSVolumesOk() (*bool, bool) {
	if o == nil || o.UseFastCloningOnXFSVolumes == nil {
		return nil, false
	}
	return o.UseFastCloningOnXFSVolumes, true
}

// HasUseFastCloningOnXFSVolumes returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasUseFastCloningOnXFSVolumes() bool {
	if o != nil && o.UseFastCloningOnXFSVolumes != nil {
		return true
	}

	return false
}

// SetUseFastCloningOnXFSVolumes gets a reference to the given bool and assigns it to the UseFastCloningOnXFSVolumes field.
func (o *LinuxLocalRepositorySettingsModel) SetUseFastCloningOnXFSVolumes(v bool) {
	o.UseFastCloningOnXFSVolumes = &v
}

// GetUseImmutableBackups returns the UseImmutableBackups field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetUseImmutableBackups() bool {
	if o == nil || o.UseImmutableBackups == nil {
		var ret bool
		return ret
	}
	return *o.UseImmutableBackups
}

// GetUseImmutableBackupsOk returns a tuple with the UseImmutableBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetUseImmutableBackupsOk() (*bool, bool) {
	if o == nil || o.UseImmutableBackups == nil {
		return nil, false
	}
	return o.UseImmutableBackups, true
}

// HasUseImmutableBackups returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasUseImmutableBackups() bool {
	if o != nil && o.UseImmutableBackups != nil {
		return true
	}

	return false
}

// SetUseImmutableBackups gets a reference to the given bool and assigns it to the UseImmutableBackups field.
func (o *LinuxLocalRepositorySettingsModel) SetUseImmutableBackups(v bool) {
	o.UseImmutableBackups = &v
}

// GetMakeRecentBackupsImmutableDays returns the MakeRecentBackupsImmutableDays field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetMakeRecentBackupsImmutableDays() int32 {
	if o == nil || o.MakeRecentBackupsImmutableDays == nil {
		var ret int32
		return ret
	}
	return *o.MakeRecentBackupsImmutableDays
}

// GetMakeRecentBackupsImmutableDaysOk returns a tuple with the MakeRecentBackupsImmutableDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetMakeRecentBackupsImmutableDaysOk() (*int32, bool) {
	if o == nil || o.MakeRecentBackupsImmutableDays == nil {
		return nil, false
	}
	return o.MakeRecentBackupsImmutableDays, true
}

// HasMakeRecentBackupsImmutableDays returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasMakeRecentBackupsImmutableDays() bool {
	if o != nil && o.MakeRecentBackupsImmutableDays != nil {
		return true
	}

	return false
}

// SetMakeRecentBackupsImmutableDays gets a reference to the given int32 and assigns it to the MakeRecentBackupsImmutableDays field.
func (o *LinuxLocalRepositorySettingsModel) SetMakeRecentBackupsImmutableDays(v int32) {
	o.MakeRecentBackupsImmutableDays = &v
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *LinuxLocalRepositorySettingsModel) GetAdvancedSettings() RepositoryAdvancedSettingsModel {
	if o == nil || o.AdvancedSettings == nil {
		var ret RepositoryAdvancedSettingsModel
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxLocalRepositorySettingsModel) GetAdvancedSettingsOk() (*RepositoryAdvancedSettingsModel, bool) {
	if o == nil || o.AdvancedSettings == nil {
		return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *LinuxLocalRepositorySettingsModel) HasAdvancedSettings() bool {
	if o != nil && o.AdvancedSettings != nil {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given RepositoryAdvancedSettingsModel and assigns it to the AdvancedSettings field.
func (o *LinuxLocalRepositorySettingsModel) SetAdvancedSettings(v RepositoryAdvancedSettingsModel) {
	o.AdvancedSettings = &v
}

func (o LinuxLocalRepositorySettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.EnableTaskLimit != nil {
		toSerialize["enableTaskLimit"] = o.EnableTaskLimit
	}
	if o.MaxTaskCount != nil {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	if o.EnableReadWriteLimit != nil {
		toSerialize["enableReadWriteLimit"] = o.EnableReadWriteLimit
	}
	if o.ReadWriteRate != nil {
		toSerialize["readWriteRate"] = o.ReadWriteRate
	}
	if o.UseFastCloningOnXFSVolumes != nil {
		toSerialize["useFastCloningOnXFSVolumes"] = o.UseFastCloningOnXFSVolumes
	}
	if o.UseImmutableBackups != nil {
		toSerialize["useImmutableBackups"] = o.UseImmutableBackups
	}
	if o.MakeRecentBackupsImmutableDays != nil {
		toSerialize["makeRecentBackupsImmutableDays"] = o.MakeRecentBackupsImmutableDays
	}
	if o.AdvancedSettings != nil {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxLocalRepositorySettingsModel struct {
	value *LinuxLocalRepositorySettingsModel
	isSet bool
}

func (v NullableLinuxLocalRepositorySettingsModel) Get() *LinuxLocalRepositorySettingsModel {
	return v.value
}

func (v *NullableLinuxLocalRepositorySettingsModel) Set(val *LinuxLocalRepositorySettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxLocalRepositorySettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxLocalRepositorySettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxLocalRepositorySettingsModel(val *LinuxLocalRepositorySettingsModel) *NullableLinuxLocalRepositorySettingsModel {
	return &NullableLinuxLocalRepositorySettingsModel{value: val, isSet: true}
}

func (v NullableLinuxLocalRepositorySettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxLocalRepositorySettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


