# \ProxiesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProxy**](ProxiesApi.md#CreateProxy) | **Post** /api/v1/backupInfrastructure/proxies | Add Proxy
[**DeleteProxy**](ProxiesApi.md#DeleteProxy) | **Delete** /api/v1/backupInfrastructure/proxies/{id} | Remove Proxy
[**GetAllProxies**](ProxiesApi.md#GetAllProxies) | **Get** /api/v1/backupInfrastructure/proxies | Get All Proxies
[**GetProxy**](ProxiesApi.md#GetProxy) | **Get** /api/v1/backupInfrastructure/proxies/{id} | Get Proxy
[**UpdateProxy**](ProxiesApi.md#UpdateProxy) | **Put** /api/v1/backupInfrastructure/proxies/{id} | Edit Proxy



## CreateProxy

> SessionModel CreateProxy(ctx).XApiVersion(xApiVersion).ProxySpec(proxySpec).Execute()

Add Proxy



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
    proxySpec := *openapiclient.NewProxySpec("Description_example", openapiclient.EProxyType("ViProxy")) // ProxySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProxiesApi.CreateProxy(context.Background()).XApiVersion(xApiVersion).ProxySpec(proxySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.CreateProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProxy`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.CreateProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **proxySpec** | [**ProxySpec**](ProxySpec.md) |  | 

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


## DeleteProxy

> map[string]interface{} DeleteProxy(ctx, id).XApiVersion(xApiVersion).Execute()

Remove Proxy



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
    id := TODO // string | ID of the backup proxy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProxiesApi.DeleteProxy(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.DeleteProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProxy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.DeleteProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the backup proxy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProxies

> ProxiesResult GetAllProxies(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).HostIdFilter(hostIdFilter).Execute()

Get All Proxies



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
    skip := int32(56) // int32 | Number of proxies to skip. (optional)
    limit := int32(56) // int32 | Maximum number of proxies to return. (optional)
    orderColumn := openapiclient.EProxiesFiltersOrderColumn("Name") // EProxiesFiltersOrderColumn | Sorts proxies by one of the proxy parameters. (optional)
    orderAsc := true // bool | Sorts proxies in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters proxies by the `nameFilter` pattern. The pattern can match any proxy parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    typeFilter := openapiclient.EProxyType("ViProxy") // EProxyType | Filters proxies by proxy type. (optional)
    hostIdFilter := TODO // string | Filters proxies by ID of the backup server. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProxiesApi.GetAllProxies(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).HostIdFilter(hostIdFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.GetAllProxies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProxies`: ProxiesResult
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.GetAllProxies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProxiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **skip** | **int32** | Number of proxies to skip. | 
 **limit** | **int32** | Maximum number of proxies to return. | 
 **orderColumn** | [**EProxiesFiltersOrderColumn**](EProxiesFiltersOrderColumn.md) | Sorts proxies by one of the proxy parameters. | 
 **orderAsc** | **bool** | Sorts proxies in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters proxies by the &#x60;nameFilter&#x60; pattern. The pattern can match any proxy parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **typeFilter** | [**EProxyType**](EProxyType.md) | Filters proxies by proxy type. | 
 **hostIdFilter** | [**string**](string.md) | Filters proxies by ID of the backup server. | 

### Return type

[**ProxiesResult**](ProxiesResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProxy

> ProxyModel GetProxy(ctx, id).XApiVersion(xApiVersion).Execute()

Get Proxy



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
    id := TODO // string | ID of the backup proxy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProxiesApi.GetProxy(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.GetProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProxy`: ProxyModel
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.GetProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the backup proxy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


### Return type

[**ProxyModel**](ProxyModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProxy

> SessionModel UpdateProxy(ctx, id).XApiVersion(xApiVersion).ProxyModel(proxyModel).Execute()

Edit Proxy



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
    id := TODO // string | ID of the backup proxy.
    proxyModel := *openapiclient.NewProxyModel("Id_example", "Name_example", "Description_example", openapiclient.EProxyType("ViProxy")) // ProxyModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProxiesApi.UpdateProxy(context.Background(), id).XApiVersion(xApiVersion).ProxyModel(proxyModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.UpdateProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProxy`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.UpdateProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the backup proxy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

 **proxyModel** | [**ProxyModel**](ProxyModel.md) |  | 

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

