# \CustomisationTranslationsBETAApi

All URIs are relative to *https://webform-sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranslation**](CustomisationTranslationsBETAApi.md#CreateTranslation) | **Post** /api/translations | Create a translation (BETA)
[**DeleteTranslation**](CustomisationTranslationsBETAApi.md#DeleteTranslation) | **Delete** /api/translations/{id} | Delete a translation (BETA)
[**GetTranslation**](CustomisationTranslationsBETAApi.md#GetTranslation) | **Get** /api/translations/{id} | Get a translation (BETA)
[**GetTranslations**](CustomisationTranslationsBETAApi.md#GetTranslations) | **Get** /api/translations | Get translations (BETA)



## CreateTranslation

> TranslationSet CreateTranslation(ctx).CreateTranslationDetails(createTranslationDetails).Execute()

Create a translation (BETA)



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
    createTranslationDetails := *openapiclient.NewCreateTranslationDetails() // CreateTranslationDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomisationTranslationsBETAApi.CreateTranslation(context.Background()).CreateTranslationDetails(createTranslationDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationTranslationsBETAApi.CreateTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTranslation`: TranslationSet
    fmt.Fprintf(os.Stdout, "Response from `CustomisationTranslationsBETAApi.CreateTranslation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTranslationDetails** | [**CreateTranslationDetails**](CreateTranslationDetails.md) |  | 

### Return type

[**TranslationSet**](TranslationSet.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranslation

> DeleteTranslation(ctx, id).Execute()

Delete a translation (BETA)



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
    id := "id_example" // string | Identifier of the translation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomisationTranslationsBETAApi.DeleteTranslation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationTranslationsBETAApi.DeleteTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the translation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTranslationRequest struct via the builder pattern


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


## GetTranslation

> TranslationSet GetTranslation(ctx, id).Execute()

Get a translation (BETA)



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
    id := "id_example" // string | Identifier of the translation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomisationTranslationsBETAApi.GetTranslation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationTranslationsBETAApi.GetTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTranslation`: TranslationSet
    fmt.Fprintf(os.Stdout, "Response from `CustomisationTranslationsBETAApi.GetTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the translation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TranslationSet**](TranslationSet.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTranslations

> TranslationSetsPage GetTranslations(ctx).Page(page).PerPage(perPage).Execute()

Get translations (BETA)



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
    page := int32(1) // int32 | Page to load (optional) (default to 1)
    perPage := int32(20) // int32 | The number of items on the page (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomisationTranslationsBETAApi.GetTranslations(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationTranslationsBETAApi.GetTranslations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTranslations`: TranslationSetsPage
    fmt.Fprintf(os.Stdout, "Response from `CustomisationTranslationsBETAApi.GetTranslations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page to load | [default to 1]
 **perPage** | **int32** | The number of items on the page | [default to 20]

### Return type

[**TranslationSetsPage**](TranslationSetsPage.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

