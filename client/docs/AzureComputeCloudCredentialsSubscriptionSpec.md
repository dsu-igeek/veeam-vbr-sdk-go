# AzureComputeCloudCredentialsSubscriptionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**ApplicationId** | **string** |  | 
**Secret** | Pointer to **string** | either secret or certificate + passphrase | [optional] 
**Certificate** | Pointer to **string** | in pfx format, either secret or certificate + password | [optional] 
**Password** | Pointer to **string** | either secret or certificate + password | [optional] 

## Methods

### NewAzureComputeCloudCredentialsSubscriptionSpec

`func NewAzureComputeCloudCredentialsSubscriptionSpec(tenantId string, applicationId string, ) *AzureComputeCloudCredentialsSubscriptionSpec`

NewAzureComputeCloudCredentialsSubscriptionSpec instantiates a new AzureComputeCloudCredentialsSubscriptionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudCredentialsSubscriptionSpecWithDefaults

`func NewAzureComputeCloudCredentialsSubscriptionSpecWithDefaults() *AzureComputeCloudCredentialsSubscriptionSpec`

NewAzureComputeCloudCredentialsSubscriptionSpecWithDefaults instantiates a new AzureComputeCloudCredentialsSubscriptionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetApplicationId

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetSecret

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCertificate

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPassword

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AzureComputeCloudCredentialsSubscriptionSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

