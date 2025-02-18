# \ServicesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllServices**](ServicesApi.md#GetAllServices) | **Get** /api/v1/services | Get Associated Services



## GetAllServices

> ServicesResult GetAllServices(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()

Get Associated Services



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
    skip := int32(56) // int32 | Number of services to skip. (optional)
    limit := int32(56) // int32 | Maximum number of services to return. (optional)
    orderColumn := openapiclient.EServicesFiltersOrderColumn("Name") // EServicesFiltersOrderColumn | Sorts services by one of the service parameters. (optional)
    orderAsc := true // bool | Sorts services in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters services by the `nameFilter` pattern. The pattern can match any service parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.GetAllServices(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetAllServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllServices`: ServicesResult
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetAllServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **skip** | **int32** | Number of services to skip. | 
 **limit** | **int32** | Maximum number of services to return. | 
 **orderColumn** | [**EServicesFiltersOrderColumn**](EServicesFiltersOrderColumn.md) | Sorts services by one of the service parameters. | 
 **orderAsc** | **bool** | Sorts services in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters services by the &#x60;nameFilter&#x60; pattern. The pattern can match any service parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 

### Return type

[**ServicesResult**](ServicesResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

