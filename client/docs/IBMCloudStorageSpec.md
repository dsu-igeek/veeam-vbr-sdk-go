# IBMCloudStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** |  | [optional] 
**MaxTaskCount** | Pointer to **int32** |  | [optional] 
**Account** | [**IBMCloudStorageAccountModel**](IBMCloudStorageAccountModel.md) |  | 
**Bucket** | [**IBMCloudStorageBucketModel**](IBMCloudStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 
**ProxyAppliance** | Pointer to [**S3CompatibleProxyModel**](S3CompatibleProxyModel.md) |  | [optional] 

## Methods

### NewIBMCloudStorageSpec

`func NewIBMCloudStorageSpec(account IBMCloudStorageAccountModel, bucket IBMCloudStorageBucketModel, mountServer MountServerSettingsModel, ) *IBMCloudStorageSpec`

NewIBMCloudStorageSpec instantiates a new IBMCloudStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBMCloudStorageSpecWithDefaults

`func NewIBMCloudStorageSpecWithDefaults() *IBMCloudStorageSpec`

NewIBMCloudStorageSpecWithDefaults instantiates a new IBMCloudStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *IBMCloudStorageSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *IBMCloudStorageSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *IBMCloudStorageSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *IBMCloudStorageSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *IBMCloudStorageSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *IBMCloudStorageSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *IBMCloudStorageSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *IBMCloudStorageSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *IBMCloudStorageSpec) GetAccount() IBMCloudStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IBMCloudStorageSpec) GetAccountOk() (*IBMCloudStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IBMCloudStorageSpec) SetAccount(v IBMCloudStorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *IBMCloudStorageSpec) GetBucket() IBMCloudStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IBMCloudStorageSpec) GetBucketOk() (*IBMCloudStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IBMCloudStorageSpec) SetBucket(v IBMCloudStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *IBMCloudStorageSpec) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *IBMCloudStorageSpec) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *IBMCloudStorageSpec) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *IBMCloudStorageSpec) GetProxyAppliance() S3CompatibleProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *IBMCloudStorageSpec) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *IBMCloudStorageSpec) SetProxyAppliance(v S3CompatibleProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *IBMCloudStorageSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

