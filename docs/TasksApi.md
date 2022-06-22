# \TasksApi

All URIs are relative to *https://webform-sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTask**](TasksApi.md#GetTask) | **Get** /api/tasks/{id} | Get a task
[**GetTasks**](TasksApi.md#GetTasks) | **Get** /api/tasks | Get tasks



## GetTask

> Task GetTask(ctx, id).Execute()

Get a task



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
    id := "id_example" // string | Identifier of the task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.GetTask(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> TasksPage GetTasks(ctx).Order(order).Page(page).PerPage(perPage).Execute()

Get tasks



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
    order := "createdAt,desc" // string | Determines the order of the results. You can order by <code>createdAt</code> field ascending or descending. The general format is <code>property[,asc|desc]</code>, with <code>asc</code> being the default value.The default order is <code>createdAt,asc</code>. (optional)
    page := int32(1) // int32 | Page to load (optional) (default to 1)
    perPage := int32(20) // int32 | The number of items on the page (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.GetTasks(context.Background()).Order(order).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasks`: TasksPage
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Determines the order of the results. You can order by &lt;code&gt;createdAt&lt;/code&gt; field ascending or descending. The general format is &lt;code&gt;property[,asc|desc]&lt;/code&gt;, with &lt;code&gt;asc&lt;/code&gt; being the default value.The default order is &lt;code&gt;createdAt,asc&lt;/code&gt;. | 
 **page** | **int32** | Page to load | [default to 1]
 **perPage** | **int32** | The number of items on the page | [default to 20]

### Return type

[**TasksPage**](TasksPage.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

