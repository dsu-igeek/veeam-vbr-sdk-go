# \SessionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSessions**](SessionsApi.md#GetAllSessions) | **Get** /api/v1/sessions | Get All Sessions
[**GetSession**](SessionsApi.md#GetSession) | **Get** /api/v1/sessions/{id} | Get Session
[**GetSessionLogs**](SessionsApi.md#GetSessionLogs) | **Get** /api/v1/sessions/{id}/logs | Get Session Logs
[**StopSession**](SessionsApi.md#StopSession) | **Post** /api/v1/sessions/{id}/stop | Stop Session



## GetAllSessions

> SessionsResult GetAllSessions(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).CreatedAfterFilter(createdAfterFilter).CreatedBeforeFilter(createdBeforeFilter).EndedAfterFilter(endedAfterFilter).EndedBeforeFilter(endedBeforeFilter).TypeFilter(typeFilter).StateFilter(stateFilter).ResultFilter(resultFilter).JobIdFilter(jobIdFilter).Execute()

Get All Sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    skip := int32(56) // int32 | Number of sessions to skip. (optional)
    limit := int32(56) // int32 | Maximum number of sessions to return. (optional)
    orderColumn := openapiclient.ESessionsFiltersOrderColumn("Name") // ESessionsFiltersOrderColumn | Sorts sessions by one of the session parameters. (optional)
    orderAsc := true // bool | Sorts sessions in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters sessions by the `nameFilter` pattern. The pattern can match any session parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    createdAfterFilter := time.Now() // time.Time | Returns sessions that are created after the specified date and time. (optional)
    createdBeforeFilter := time.Now() // time.Time | Returns sessions that are created before the specified date and time. (optional)
    endedAfterFilter := time.Now() // time.Time | Returns sessions that are finished after the specified date and time. (optional)
    endedBeforeFilter := time.Now() // time.Time | Returns sessions that are finished before the specified date and time. (optional)
    typeFilter := openapiclient.ESessionType("Infrastructure") // ESessionType | Filters sessions by session type. (optional)
    stateFilter := openapiclient.ESessionState("Stopped") // ESessionState | Filters sessions by session state. (optional)
    resultFilter := openapiclient.ESessionResult("None") // ESessionResult | Filters sessions by session result. (optional)
    jobIdFilter := TODO // string | Filters sessions by job ID. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetAllSessions(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).CreatedAfterFilter(createdAfterFilter).CreatedBeforeFilter(createdBeforeFilter).EndedAfterFilter(endedAfterFilter).EndedBeforeFilter(endedBeforeFilter).TypeFilter(typeFilter).StateFilter(stateFilter).ResultFilter(resultFilter).JobIdFilter(jobIdFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetAllSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSessions`: SessionsResult
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetAllSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **skip** | **int32** | Number of sessions to skip. | 
 **limit** | **int32** | Maximum number of sessions to return. | 
 **orderColumn** | [**ESessionsFiltersOrderColumn**](ESessionsFiltersOrderColumn.md) | Sorts sessions by one of the session parameters. | 
 **orderAsc** | **bool** | Sorts sessions in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters sessions by the &#x60;nameFilter&#x60; pattern. The pattern can match any session parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **createdAfterFilter** | **time.Time** | Returns sessions that are created after the specified date and time. | 
 **createdBeforeFilter** | **time.Time** | Returns sessions that are created before the specified date and time. | 
 **endedAfterFilter** | **time.Time** | Returns sessions that are finished after the specified date and time. | 
 **endedBeforeFilter** | **time.Time** | Returns sessions that are finished before the specified date and time. | 
 **typeFilter** | [**ESessionType**](ESessionType.md) | Filters sessions by session type. | 
 **stateFilter** | [**ESessionState**](ESessionState.md) | Filters sessions by session state. | 
 **resultFilter** | [**ESessionResult**](ESessionResult.md) | Filters sessions by session result. | 
 **jobIdFilter** | [**string**](string.md) | Filters sessions by job ID. | 

### Return type

[**SessionsResult**](SessionsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSession

> SessionModel GetSession(ctx, id).XApiVersion(xApiVersion).Execute()

Get Session



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
    id := TODO // string | ID of the session.
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetSession(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSession`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRequest struct via the builder pattern


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


## GetSessionLogs

> SessionLogResult GetSessionLogs(ctx, id).XApiVersion(xApiVersion).Execute()

Get Session Logs



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
    id := TODO // string | ID of the session.
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetSessionLogs(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessionLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionLogs`: SessionLogResult
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

### Return type

[**SessionLogResult**](SessionLogResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopSession

> map[string]interface{} StopSession(ctx, id).XApiVersion(xApiVersion).Execute()

Stop Session



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
    id := TODO // string | ID of the session.
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.StopSession(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.StopSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopSession`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.StopSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSessionRequest struct via the builder pattern


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

