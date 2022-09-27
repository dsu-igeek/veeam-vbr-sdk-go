# AmazonS3StorageSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** |  | [optional] 
**MaxTaskCount** | Pointer to **int32** |  | [optional] 
**Account** | [**AmazonS3StorageAccountModel**](AmazonS3StorageAccountModel.md) |  | 
**Bucket** | [**AmazonS3StorageBucketModel**](AmazonS3StorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 
**ProxyAppliance** | Pointer to [**AmazonS3StorageProxyApplianceModel**](AmazonS3StorageProxyApplianceModel.md) |  | [optional] 

## Methods

### NewAmazonS3StorageSpecAllOf

`func NewAmazonS3StorageSpecAllOf(account AmazonS3StorageAccountModel, bucket AmazonS3StorageBucketModel, mountServer MountServerSettingsModel, ) *AmazonS3StorageSpecAllOf`

NewAmazonS3StorageSpecAllOf instantiates a new AmazonS3StorageSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3StorageSpecAllOfWithDefaults

`func NewAmazonS3StorageSpecAllOfWithDefaults() *AmazonS3StorageSpecAllOf`

NewAmazonS3StorageSpecAllOfWithDefaults instantiates a new AmazonS3StorageSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *AmazonS3StorageSpecAllOf) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AmazonS3StorageSpecAllOf) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AmazonS3StorageSpecAllOf) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AmazonS3StorageSpecAllOf) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AmazonS3StorageSpecAllOf) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AmazonS3StorageSpecAllOf) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AmazonS3StorageSpecAllOf) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AmazonS3StorageSpecAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AmazonS3StorageSpecAllOf) GetAccount() AmazonS3StorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonS3StorageSpecAllOf) GetAccountOk() (*AmazonS3StorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonS3StorageSpecAllOf) SetAccount(v AmazonS3StorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonS3StorageSpecAllOf) GetBucket() AmazonS3StorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonS3StorageSpecAllOf) GetBucketOk() (*AmazonS3StorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonS3StorageSpecAllOf) SetBucket(v AmazonS3StorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *AmazonS3StorageSpecAllOf) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AmazonS3StorageSpecAllOf) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AmazonS3StorageSpecAllOf) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *AmazonS3StorageSpecAllOf) GetProxyAppliance() AmazonS3StorageProxyApplianceModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AmazonS3StorageSpecAllOf) GetProxyApplianceOk() (*AmazonS3StorageProxyApplianceModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AmazonS3StorageSpecAllOf) SetProxyAppliance(v AmazonS3StorageProxyApplianceModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AmazonS3StorageSpecAllOf) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

