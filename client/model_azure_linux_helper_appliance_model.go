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

// AzureLinuxHelperApplianceModel struct for AzureLinuxHelperApplianceModel
type AzureLinuxHelperApplianceModel struct {
	CloudHelperApplianceModel
	Id string `json:"id"`
	// ID of a Microsoft Azure subscription credentials.
	SubscriptionId string `json:"subscriptionId"`
	VmName *string `json:"vmName,omitempty"`
	Location *string `json:"location,omitempty"`
	StorageAccount *string `json:"storageAccount,omitempty"`
	// Resource group associated with the proxy appliance.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// Network to which the helper appliance is connected.
	VirtualNetwork *string `json:"virtualNetwork,omitempty"`
	// Subnet for the proxy appliance.
	Subnet *string `json:"subnet,omitempty"`
	// TCP port used to route requests between the proxy appliance and backup infrastructure components.
	SSHPort *int32 `json:"SSHPort,omitempty"`
}

// NewAzureLinuxHelperApplianceModel instantiates a new AzureLinuxHelperApplianceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureLinuxHelperApplianceModel(id string, subscriptionId string, ) *AzureLinuxHelperApplianceModel {
	this := AzureLinuxHelperApplianceModel{}
	this.Id = id
	this.SubscriptionId = subscriptionId
	return &this
}

// NewAzureLinuxHelperApplianceModelWithDefaults instantiates a new AzureLinuxHelperApplianceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureLinuxHelperApplianceModelWithDefaults() *AzureLinuxHelperApplianceModel {
	this := AzureLinuxHelperApplianceModel{}
	return &this
}

// GetId returns the Id field value
func (o *AzureLinuxHelperApplianceModel) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureLinuxHelperApplianceModel) SetId(v string) {
	o.Id = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *AzureLinuxHelperApplianceModel) GetSubscriptionId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *AzureLinuxHelperApplianceModel) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetVmName returns the VmName field value if set, zero value otherwise.
func (o *AzureLinuxHelperApplianceModel) GetVmName() string {
	if o == nil || o.VmName == nil {
		var ret string
		return ret
	}
	return *o.VmName
}

// GetVmNameOk returns a tuple with the VmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetVmNameOk() (*string, bool) {
	if o == nil || o.VmName == nil {
		return nil, false
	}
	return o.VmName, true
}

// HasVmName returns a boolean if a field has been set.
func (o *AzureLinuxHelperApplianceModel) HasVmName() bool {
	if o != nil && o.VmName != nil {
		return true
	}

	return false
}

// SetVmName gets a reference to the given string and assigns it to the VmName field.
func (o *AzureLinuxHelperApplianceModel) SetVmName(v string) {
	o.VmName = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AzureLinuxHelperApplianceModel) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AzureLinuxHelperApplianceModel) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AzureLinuxHelperApplianceModel) SetLocation(v string) {
	o.Location = &v
}

// GetStorageAccount returns the StorageAccount field value if set, zero value otherwise.
func (o *AzureLinuxHelperApplianceModel) GetStorageAccount() string {
	if o == nil || o.StorageAccount == nil {
		var ret string
		return ret
	}
	return *o.StorageAccount
}

// GetStorageAccountOk returns a tuple with the StorageAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetStorageAccountOk() (*string, bool) {
	if o == nil || o.StorageAccount == nil {
		return nil, false
	}
	return o.StorageAccount, true
}

// HasStorageAccount returns a boolean if a field has been set.
func (o *AzureLinuxHelperApplianceModel) HasStorageAccount() bool {
	if o != nil && o.StorageAccount != nil {
		return true
	}

	return false
}

// SetStorageAccount gets a reference to the given string and assigns it to the StorageAccount field.
func (o *AzureLinuxHelperApplianceModel) SetStorageAccount(v string) {
	o.StorageAccount = &v
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise.
func (o *AzureLinuxHelperApplianceModel) GetResourceGroup() string {
	if o == nil || o.ResourceGroup == nil {
		var ret string
		return ret
	}
	return *o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetResourceGroupOk() (*string, bool) {
	if o == nil || o.ResourceGroup == nil {
		return nil, false
	}
	return o.ResourceGroup, true
}

// HasResourceGroup returns a boolean if a field has been set.
func (o *AzureLinuxHelperApplianceModel) HasResourceGroup() bool {
	if o != nil && o.ResourceGroup != nil {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given string and assigns it to the ResourceGroup field.
func (o *AzureLinuxHelperApplianceModel) SetResourceGroup(v string) {
	o.ResourceGroup = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *AzureLinuxHelperApplianceModel) GetVirtualNetwork() string {
	if o == nil || o.VirtualNetwork == nil {
		var ret string
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetVirtualNetworkOk() (*string, bool) {
	if o == nil || o.VirtualNetwork == nil {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *AzureLinuxHelperApplianceModel) HasVirtualNetwork() bool {
	if o != nil && o.VirtualNetwork != nil {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given string and assigns it to the VirtualNetwork field.
func (o *AzureLinuxHelperApplianceModel) SetVirtualNetwork(v string) {
	o.VirtualNetwork = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *AzureLinuxHelperApplianceModel) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *AzureLinuxHelperApplianceModel) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *AzureLinuxHelperApplianceModel) SetSubnet(v string) {
	o.Subnet = &v
}

// GetSSHPort returns the SSHPort field value if set, zero value otherwise.
func (o *AzureLinuxHelperApplianceModel) GetSSHPort() int32 {
	if o == nil || o.SSHPort == nil {
		var ret int32
		return ret
	}
	return *o.SSHPort
}

// GetSSHPortOk returns a tuple with the SSHPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLinuxHelperApplianceModel) GetSSHPortOk() (*int32, bool) {
	if o == nil || o.SSHPort == nil {
		return nil, false
	}
	return o.SSHPort, true
}

// HasSSHPort returns a boolean if a field has been set.
func (o *AzureLinuxHelperApplianceModel) HasSSHPort() bool {
	if o != nil && o.SSHPort != nil {
		return true
	}

	return false
}

// SetSSHPort gets a reference to the given int32 and assigns it to the SSHPort field.
func (o *AzureLinuxHelperApplianceModel) SetSSHPort(v int32) {
	o.SSHPort = &v
}

func (o AzureLinuxHelperApplianceModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudHelperApplianceModel, errCloudHelperApplianceModel := json.Marshal(o.CloudHelperApplianceModel)
	if errCloudHelperApplianceModel != nil {
		return []byte{}, errCloudHelperApplianceModel
	}
	errCloudHelperApplianceModel = json.Unmarshal([]byte(serializedCloudHelperApplianceModel), &toSerialize)
	if errCloudHelperApplianceModel != nil {
		return []byte{}, errCloudHelperApplianceModel
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.VmName != nil {
		toSerialize["vmName"] = o.VmName
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.StorageAccount != nil {
		toSerialize["storageAccount"] = o.StorageAccount
	}
	if o.ResourceGroup != nil {
		toSerialize["resourceGroup"] = o.ResourceGroup
	}
	if o.VirtualNetwork != nil {
		toSerialize["virtualNetwork"] = o.VirtualNetwork
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.SSHPort != nil {
		toSerialize["SSHPort"] = o.SSHPort
	}
	return json.Marshal(toSerialize)
}

type NullableAzureLinuxHelperApplianceModel struct {
	value *AzureLinuxHelperApplianceModel
	isSet bool
}

func (v NullableAzureLinuxHelperApplianceModel) Get() *AzureLinuxHelperApplianceModel {
	return v.value
}

func (v *NullableAzureLinuxHelperApplianceModel) Set(val *AzureLinuxHelperApplianceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureLinuxHelperApplianceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureLinuxHelperApplianceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureLinuxHelperApplianceModel(val *AzureLinuxHelperApplianceModel) *NullableAzureLinuxHelperApplianceModel {
	return &NullableAzureLinuxHelperApplianceModel{value: val, isSet: true}
}

func (v NullableAzureLinuxHelperApplianceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureLinuxHelperApplianceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

