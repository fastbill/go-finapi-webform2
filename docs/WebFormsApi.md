# \WebFormsApi

All URIs are relative to *https://webform-sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebForm**](WebFormsApi.md#DeleteWebForm) | **Delete** /api/webForms/{id} | Delete a web form
[**GetWebForm**](WebFormsApi.md#GetWebForm) | **Get** /api/webForms/{id} | Get a web form
[**GetWebForms**](WebFormsApi.md#GetWebForms) | **Get** /api/webForms | Get web forms



## DeleteWebForm

> DeleteWebForm(ctx, id).Execute()

Delete a web form



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
    id := "id_example" // string | Identifier of the web form

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebFormsApi.DeleteWebForm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebFormsApi.DeleteWebForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the web form | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebForm

> WebForm GetWebForm(ctx, id).Execute()

Get a web form



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
    id := "id_example" // string | Identifier of the web form

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebFormsApi.GetWebForm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebFormsApi.GetWebForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebForm`: WebForm
    fmt.Fprintf(os.Stdout, "Response from `WebFormsApi.GetWebForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the web form | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebForm**](WebForm.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebForms

> WebFormsPage GetWebForms(ctx).Order(order).Page(page).PerPage(perPage).Execute()

Get web forms



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
    resp, r, err := apiClient.WebFormsApi.GetWebForms(context.Background()).Order(order).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebFormsApi.GetWebForms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebForms`: WebFormsPage
    fmt.Fprintf(os.Stdout, "Response from `WebFormsApi.GetWebForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Determines the order of the results. You can order by &lt;code&gt;createdAt&lt;/code&gt; field ascending or descending. The general format is &lt;code&gt;property[,asc|desc]&lt;/code&gt;, with &lt;code&gt;asc&lt;/code&gt; being the default value.The default order is &lt;code&gt;createdAt,asc&lt;/code&gt;. | 
 **page** | **int32** | Page to load | [default to 1]
 **perPage** | **int32** | The number of items on the page | [default to 20]

### Return type

[**WebFormsPage**](WebFormsPage.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

