# \CloudBrowserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseCloudEntity**](CloudBrowserApi.md#BrowseCloudEntity) | **Post** /api/v1/cloudBrowser | Get Cloud Hierarchy
[**CreateNewCloudFolder**](CloudBrowserApi.md#CreateNewCloudFolder) | **Post** /api/v1/cloudBrowser/newFolder | Creates new folder



## BrowseCloudEntity

> CloudBrowserModel BrowseCloudEntity(ctx).XApiVersion(xApiVersion).CloudBrowserSpec(cloudBrowserSpec).Execute()

Get Cloud Hierarchy



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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the *\\<version\\>-\\<revision\\>* format. (default to "1.1-rev0")
    cloudBrowserSpec := *openapiclient.NewCloudBrowserSpec("CredentialsId_example", openapiclient.ECloudServiceType("AzureBlob")) // CloudBrowserSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudBrowserApi.BrowseCloudEntity(context.Background()).XApiVersion(xApiVersion).CloudBrowserSpec(cloudBrowserSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBrowserApi.BrowseCloudEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseCloudEntity`: CloudBrowserModel
    fmt.Fprintf(os.Stdout, "Response from `CloudBrowserApi.BrowseCloudEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrowseCloudEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the *\\&lt;version\\&gt;-\\&lt;revision\\&gt;* format. | [default to &quot;1.1-rev0&quot;]
 **cloudBrowserSpec** | [**CloudBrowserSpec**](CloudBrowserSpec.md) |  | 

### Return type

[**CloudBrowserModel**](CloudBrowserModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewCloudFolder

> map[string]interface{} CreateNewCloudFolder(ctx).XApiVersion(xApiVersion).CloudBrowserNewFolderSpec(cloudBrowserNewFolderSpec).Execute()

Creates new folder



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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the *\\<version\\>-\\<revision\\>* format. (default to "1.1-rev0")
    cloudBrowserNewFolderSpec := *openapiclient.NewCloudBrowserNewFolderSpec("CredentialsId_example", openapiclient.ECloudServiceType("AzureBlob"), "NewFolderName_example") // CloudBrowserNewFolderSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudBrowserApi.CreateNewCloudFolder(context.Background()).XApiVersion(xApiVersion).CloudBrowserNewFolderSpec(cloudBrowserNewFolderSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBrowserApi.CreateNewCloudFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewCloudFolder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudBrowserApi.CreateNewCloudFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewCloudFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the *\\&lt;version\\&gt;-\\&lt;revision\\&gt;* format. | [default to &quot;1.1-rev0&quot;]
 **cloudBrowserNewFolderSpec** | [**CloudBrowserNewFolderSpec**](CloudBrowserNewFolderSpec.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
