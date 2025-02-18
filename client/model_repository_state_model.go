/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RepositoryStateModel struct for RepositoryStateModel
type RepositoryStateModel struct {
	// ID of the backup repository.
	Id string `json:"id"`
	// Name of the backup repository.
	Name string `json:"name"`
	Type ERepositoryType `json:"type"`
	Description string `json:"description"`
	// ID of the server that is used as a backup repository.
	HostId *string `json:"hostId,omitempty"`
	// Name of the server that is used as a backup repository.
	HostName *string `json:"hostName,omitempty"`
	// Path to the folder where backup files are stored.
	Path *string `json:"path,omitempty"`
	// Repository capacity in GB.
	CapacityGB float64 `json:"capacityGB"`
	// Repository free space in GB.
	FreeGB float64 `json:"freeGB"`
	// Repository used space in GB.
	UsedSpaceGB float64 `json:"usedSpaceGB"`
}

// NewRepositoryStateModel instantiates a new RepositoryStateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryStateModel(id string, name string, type_ ERepositoryType, description string, capacityGB float64, freeGB float64, usedSpaceGB float64) *RepositoryStateModel {
	this := RepositoryStateModel{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Description = description
	this.CapacityGB = capacityGB
	this.FreeGB = freeGB
	this.UsedSpaceGB = usedSpaceGB
	return &this
}

// NewRepositoryStateModelWithDefaults instantiates a new RepositoryStateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryStateModelWithDefaults() *RepositoryStateModel {
	this := RepositoryStateModel{}
	return &this
}

// GetId returns the Id field value
func (o *RepositoryStateModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RepositoryStateModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RepositoryStateModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RepositoryStateModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *RepositoryStateModel) GetType() ERepositoryType {
	if o == nil {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RepositoryStateModel) SetType(v ERepositoryType) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *RepositoryStateModel) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RepositoryStateModel) SetDescription(v string) {
	o.Description = v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *RepositoryStateModel) GetHostId() string {
	if o == nil || o.HostId == nil {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetHostIdOk() (*string, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *RepositoryStateModel) HasHostId() bool {
	if o != nil && o.HostId != nil {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *RepositoryStateModel) SetHostId(v string) {
	o.HostId = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *RepositoryStateModel) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *RepositoryStateModel) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *RepositoryStateModel) SetHostName(v string) {
	o.HostName = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RepositoryStateModel) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RepositoryStateModel) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RepositoryStateModel) SetPath(v string) {
	o.Path = &v
}

// GetCapacityGB returns the CapacityGB field value
func (o *RepositoryStateModel) GetCapacityGB() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CapacityGB
}

// GetCapacityGBOk returns a tuple with the CapacityGB field value
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetCapacityGBOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CapacityGB, true
}

// SetCapacityGB sets field value
func (o *RepositoryStateModel) SetCapacityGB(v float64) {
	o.CapacityGB = v
}

// GetFreeGB returns the FreeGB field value
func (o *RepositoryStateModel) GetFreeGB() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FreeGB
}

// GetFreeGBOk returns a tuple with the FreeGB field value
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetFreeGBOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FreeGB, true
}

// SetFreeGB sets field value
func (o *RepositoryStateModel) SetFreeGB(v float64) {
	o.FreeGB = v
}

// GetUsedSpaceGB returns the UsedSpaceGB field value
func (o *RepositoryStateModel) GetUsedSpaceGB() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UsedSpaceGB
}

// GetUsedSpaceGBOk returns a tuple with the UsedSpaceGB field value
// and a boolean to check if the value has been set.
func (o *RepositoryStateModel) GetUsedSpaceGBOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UsedSpaceGB, true
}

// SetUsedSpaceGB sets field value
func (o *RepositoryStateModel) SetUsedSpaceGB(v float64) {
	o.UsedSpaceGB = v
}

func (o RepositoryStateModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.HostId != nil {
		toSerialize["hostId"] = o.HostId
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["capacityGB"] = o.CapacityGB
	}
	if true {
		toSerialize["freeGB"] = o.FreeGB
	}
	if true {
		toSerialize["usedSpaceGB"] = o.UsedSpaceGB
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryStateModel struct {
	value *RepositoryStateModel
	isSet bool
}

func (v NullableRepositoryStateModel) Get() *RepositoryStateModel {
	return v.value
}

func (v *NullableRepositoryStateModel) Set(val *RepositoryStateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryStateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryStateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryStateModel(val *RepositoryStateModel) *NullableRepositoryStateModel {
	return &NullableRepositoryStateModel{value: val, isSet: true}
}

func (v NullableRepositoryStateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryStateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


