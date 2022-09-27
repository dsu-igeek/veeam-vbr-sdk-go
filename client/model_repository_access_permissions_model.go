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

// RepositoryAccessPermissionsModel Repository access permissions.
type RepositoryAccessPermissionsModel struct {
	AccessPolicy ERepositoryAccessType `json:"accessPolicy"`
	// Array of accounts that have access to the backup repository.
	Accounts *[]string `json:"accounts,omitempty"`
	// If *true*, Veeam Backup & Replication encrypts Veeam Agent backup files stored in the backup repository.
	EncryptBackups bool `json:"encryptBackups"`
	// ID of the password used for encryption.
	PasswordId *string `json:"passwordId,omitempty"`
}

// NewRepositoryAccessPermissionsModel instantiates a new RepositoryAccessPermissionsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryAccessPermissionsModel(accessPolicy ERepositoryAccessType, encryptBackups bool, ) *RepositoryAccessPermissionsModel {
	this := RepositoryAccessPermissionsModel{}
	this.AccessPolicy = accessPolicy
	this.EncryptBackups = encryptBackups
	return &this
}

// NewRepositoryAccessPermissionsModelWithDefaults instantiates a new RepositoryAccessPermissionsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryAccessPermissionsModelWithDefaults() *RepositoryAccessPermissionsModel {
	this := RepositoryAccessPermissionsModel{}
	return &this
}

// GetAccessPolicy returns the AccessPolicy field value
func (o *RepositoryAccessPermissionsModel) GetAccessPolicy() ERepositoryAccessType {
	if o == nil  {
		var ret ERepositoryAccessType
		return ret
	}

	return o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value
// and a boolean to check if the value has been set.
func (o *RepositoryAccessPermissionsModel) GetAccessPolicyOk() (*ERepositoryAccessType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessPolicy, true
}

// SetAccessPolicy sets field value
func (o *RepositoryAccessPermissionsModel) SetAccessPolicy(v ERepositoryAccessType) {
	o.AccessPolicy = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *RepositoryAccessPermissionsModel) GetAccounts() []string {
	if o == nil || o.Accounts == nil {
		var ret []string
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAccessPermissionsModel) GetAccountsOk() (*[]string, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *RepositoryAccessPermissionsModel) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []string and assigns it to the Accounts field.
func (o *RepositoryAccessPermissionsModel) SetAccounts(v []string) {
	o.Accounts = &v
}

// GetEncryptBackups returns the EncryptBackups field value
func (o *RepositoryAccessPermissionsModel) GetEncryptBackups() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.EncryptBackups
}

// GetEncryptBackupsOk returns a tuple with the EncryptBackups field value
// and a boolean to check if the value has been set.
func (o *RepositoryAccessPermissionsModel) GetEncryptBackupsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EncryptBackups, true
}

// SetEncryptBackups sets field value
func (o *RepositoryAccessPermissionsModel) SetEncryptBackups(v bool) {
	o.EncryptBackups = v
}

// GetPasswordId returns the PasswordId field value if set, zero value otherwise.
func (o *RepositoryAccessPermissionsModel) GetPasswordId() string {
	if o == nil || o.PasswordId == nil {
		var ret string
		return ret
	}
	return *o.PasswordId
}

// GetPasswordIdOk returns a tuple with the PasswordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAccessPermissionsModel) GetPasswordIdOk() (*string, bool) {
	if o == nil || o.PasswordId == nil {
		return nil, false
	}
	return o.PasswordId, true
}

// HasPasswordId returns a boolean if a field has been set.
func (o *RepositoryAccessPermissionsModel) HasPasswordId() bool {
	if o != nil && o.PasswordId != nil {
		return true
	}

	return false
}

// SetPasswordId gets a reference to the given string and assigns it to the PasswordId field.
func (o *RepositoryAccessPermissionsModel) SetPasswordId(v string) {
	o.PasswordId = &v
}

func (o RepositoryAccessPermissionsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessPolicy"] = o.AccessPolicy
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["encryptBackups"] = o.EncryptBackups
	}
	if o.PasswordId != nil {
		toSerialize["passwordId"] = o.PasswordId
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryAccessPermissionsModel struct {
	value *RepositoryAccessPermissionsModel
	isSet bool
}

func (v NullableRepositoryAccessPermissionsModel) Get() *RepositoryAccessPermissionsModel {
	return v.value
}

func (v *NullableRepositoryAccessPermissionsModel) Set(val *RepositoryAccessPermissionsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryAccessPermissionsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryAccessPermissionsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryAccessPermissionsModel(val *RepositoryAccessPermissionsModel) *NullableRepositoryAccessPermissionsModel {
	return &NullableRepositoryAccessPermissionsModel{value: val, isSet: true}
}

func (v NullableRepositoryAccessPermissionsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryAccessPermissionsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

