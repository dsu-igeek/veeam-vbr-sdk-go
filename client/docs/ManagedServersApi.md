# \ManagedServersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagedServer**](ManagedServersApi.md#CreateManagedServer) | **Post** /api/v1/backupInfrastructure/managedServers | Add Server
[**DeleteManagedServer**](ManagedServersApi.md#DeleteManagedServer) | **Delete** /api/v1/backupInfrastructure/managedServers/{id} | Remove Server
[**GetAllManagedServers**](ManagedServersApi.md#GetAllManagedServers) | **Get** /api/v1/backupInfrastructure/managedServers | Get All Servers
[**GetManagedServer**](ManagedServersApi.md#GetManagedServer) | **Get** /api/v1/backupInfrastructure/managedServers/{id} | Get Server
[**UpdateManagedServer**](ManagedServersApi.md#UpdateManagedServer) | **Put** /api/v1/backupInfrastructure/managedServers/{id} | Edit Server



## CreateManagedServer

> SessionModel CreateManagedServer(ctx).XApiVersion(xApiVersion).ManagedServerSpec(managedServerSpec).Execute()

Add Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    managedServerSpec := *openapiclient.NewManagedServerSpec("Name_example", "Description_example", openapiclient.EManagedServerType("WindowsHost"), "CredentialsId_example") // ManagedServerSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedServersApi.CreateManagedServer(context.Background()).XApiVersion(xApiVersion).ManagedServerSpec(managedServerSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServersApi.CreateManagedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedServer`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `ManagedServersApi.CreateManagedServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **managedServerSpec** | [**ManagedServerSpec**](ManagedServerSpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManagedServer

> SessionModel DeleteManagedServer(ctx, id).XApiVersion(xApiVersion).Execute()

Remove Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the managed server.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedServersApi.DeleteManagedServer(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServersApi.DeleteManagedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteManagedServer`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `ManagedServersApi.DeleteManagedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the managed server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllManagedServers

> ManagedServersResult GetAllManagedServers(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).ViTypeFilter(viTypeFilter).Execute()

Get All Servers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    skip := int32(56) // int32 | Number of servers to skip. (optional)
    limit := int32(56) // int32 | Maximum number of servers to return. (optional)
    orderColumn := openapiclient.EManagedServersFiltersOrderColumn("Name") // EManagedServersFiltersOrderColumn | Sorts servers by one of the server parameters. (optional)
    orderAsc := true // bool | Sorts servers in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters servers by the `nameFilter` pattern. The pattern can match any server parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    typeFilter := openapiclient.EManagedServerType("WindowsHost") // EManagedServerType | Filters servers by server type. (optional)
    viTypeFilter := openapiclient.EViHostType("ESX") // EViHostType | Filters servers by the type of VMware vSphere server. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedServersApi.GetAllManagedServers(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).ViTypeFilter(viTypeFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServersApi.GetAllManagedServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllManagedServers`: ManagedServersResult
    fmt.Fprintf(os.Stdout, "Response from `ManagedServersApi.GetAllManagedServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllManagedServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **skip** | **int32** | Number of servers to skip. | 
 **limit** | **int32** | Maximum number of servers to return. | 
 **orderColumn** | [**EManagedServersFiltersOrderColumn**](EManagedServersFiltersOrderColumn.md) | Sorts servers by one of the server parameters. | 
 **orderAsc** | **bool** | Sorts servers in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters servers by the &#x60;nameFilter&#x60; pattern. The pattern can match any server parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **typeFilter** | [**EManagedServerType**](EManagedServerType.md) | Filters servers by server type. | 
 **viTypeFilter** | [**EViHostType**](EViHostType.md) | Filters servers by the type of VMware vSphere server. | 

### Return type

[**ManagedServersResult**](ManagedServersResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedServer

> ManagedServerModel GetManagedServer(ctx, id).XApiVersion(xApiVersion).Execute()

Get Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the managed server.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedServersApi.GetManagedServer(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServersApi.GetManagedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedServer`: ManagedServerModel
    fmt.Fprintf(os.Stdout, "Response from `ManagedServersApi.GetManagedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the managed server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


### Return type

[**ManagedServerModel**](ManagedServerModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedServer

> SessionModel UpdateManagedServer(ctx, id).XApiVersion(xApiVersion).ManagedServerModel(managedServerModel).Execute()

Edit Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the managed server.
    managedServerModel := *openapiclient.NewManagedServerModel("Id_example", "Name_example", "Description_example", openapiclient.EManagedServerType("WindowsHost"), "CredentialsId_example") // ManagedServerModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedServersApi.UpdateManagedServer(context.Background(), id).XApiVersion(xApiVersion).ManagedServerModel(managedServerModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServersApi.UpdateManagedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedServer`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `ManagedServersApi.UpdateManagedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the managed server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

 **managedServerModel** | [**ManagedServerModel**](ManagedServerModel.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

