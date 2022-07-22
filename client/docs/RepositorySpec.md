# RepositorySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | Pointer to **string** | VMware vSphere tag assigned to the backup repository. | [optional] 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**HostId** | **string** | ID of the server that is used as a backup repository. | 
**Repository** | [**NetworkRepositorySettingsModel**](NetworkRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 
**Share** | [**SmbRepositoryShareSettingsModel**](SmbRepositoryShareSettingsModel.md) |  | 

## Methods

### NewRepositorySpec

`func NewRepositorySpec(name string, description string, type_ ERepositoryType, hostId string, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsModel, share SmbRepositoryShareSettingsModel, ) *RepositorySpec`

NewRepositorySpec instantiates a new RepositorySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositorySpecWithDefaults

`func NewRepositorySpecWithDefaults() *RepositorySpec`

NewRepositorySpecWithDefaults instantiates a new RepositorySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositorySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositorySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositorySpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RepositorySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositorySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositorySpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *RepositorySpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RepositorySpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RepositorySpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RepositorySpec) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *RepositorySpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositorySpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositorySpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetHostId

`func (o *RepositorySpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *RepositorySpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *RepositorySpec) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRepository

`func (o *RepositorySpec) GetRepository() NetworkRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepositorySpec) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepositorySpec) SetRepository(v NetworkRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *RepositorySpec) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *RepositorySpec) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *RepositorySpec) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.


### GetShare

`func (o *RepositorySpec) GetShare() SmbRepositoryShareSettingsModel`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *RepositorySpec) GetShareOk() (*SmbRepositoryShareSettingsModel, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *RepositorySpec) SetShare(v SmbRepositoryShareSettingsModel)`

SetShare sets Share field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


