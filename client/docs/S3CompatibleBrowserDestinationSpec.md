# S3CompatibleBrowserDestinationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** |  | [optional] 
**ConnectionPoint** | **string** |  | 
**RegionId** | **string** |  | 
**BucketName** | **string** |  | 

## Methods

### NewS3CompatibleBrowserDestinationSpec

`func NewS3CompatibleBrowserDestinationSpec(connectionPoint string, regionId string, bucketName string, ) *S3CompatibleBrowserDestinationSpec`

NewS3CompatibleBrowserDestinationSpec instantiates a new S3CompatibleBrowserDestinationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleBrowserDestinationSpecWithDefaults

`func NewS3CompatibleBrowserDestinationSpecWithDefaults() *S3CompatibleBrowserDestinationSpec`

NewS3CompatibleBrowserDestinationSpecWithDefaults instantiates a new S3CompatibleBrowserDestinationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *S3CompatibleBrowserDestinationSpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *S3CompatibleBrowserDestinationSpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *S3CompatibleBrowserDestinationSpec) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *S3CompatibleBrowserDestinationSpec) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetConnectionPoint

`func (o *S3CompatibleBrowserDestinationSpec) GetConnectionPoint() string`

GetConnectionPoint returns the ConnectionPoint field if non-nil, zero value otherwise.

### GetConnectionPointOk

`func (o *S3CompatibleBrowserDestinationSpec) GetConnectionPointOk() (*string, bool)`

GetConnectionPointOk returns a tuple with the ConnectionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoint

`func (o *S3CompatibleBrowserDestinationSpec) SetConnectionPoint(v string)`

SetConnectionPoint sets ConnectionPoint field to given value.


### GetRegionId

`func (o *S3CompatibleBrowserDestinationSpec) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *S3CompatibleBrowserDestinationSpec) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *S3CompatibleBrowserDestinationSpec) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *S3CompatibleBrowserDestinationSpec) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3CompatibleBrowserDestinationSpec) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3CompatibleBrowserDestinationSpec) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

