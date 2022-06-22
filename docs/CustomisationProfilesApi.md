# \CustomisationProfilesApi

All URIs are relative to *https://webform-sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfile**](CustomisationProfilesApi.md#CreateProfile) | **Post** /api/profiles | Create a profile
[**DeleteProfile**](CustomisationProfilesApi.md#DeleteProfile) | **Delete** /api/profiles/{id} | Delete a profile
[**EditProfile**](CustomisationProfilesApi.md#EditProfile) | **Patch** /api/profiles/{id} | Edit a profile
[**GetProfile**](CustomisationProfilesApi.md#GetProfile) | **Get** /api/profiles/{id} | Get a profile
[**GetProfiles**](CustomisationProfilesApi.md#GetProfiles) | **Get** /api/profiles | Get profiles



## CreateProfile

> Profile CreateProfile(ctx).CreateProfileDetails(createProfileDetails).Execute()

Create a profile



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
    createProfileDetails := *openapiclient.NewCreateProfileDetails("Mobile application label") // CreateProfileDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomisationProfilesApi.CreateProfile(context.Background()).CreateProfileDetails(createProfileDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationProfilesApi.CreateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfile`: Profile
    fmt.Fprintf(os.Stdout, "Response from `CustomisationProfilesApi.CreateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfileDetails** | [**CreateProfileDetails**](CreateProfileDetails.md) |  | 

### Return type

[**Profile**](Profile.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfile

> DeleteProfile(ctx, id).Execute()

Delete a profile



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
    id := "id_example" // string | Identifier of the profile

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomisationProfilesApi.DeleteProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationProfilesApi.DeleteProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


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


## EditProfile

> Profile EditProfile(ctx, id).EditProfileDetails(editProfileDetails).Execute()

Edit a profile



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
    id := "id_example" // string | Identifier of the profile
    editProfileDetails := *openapiclient.NewEditProfileDetails() // EditProfileDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomisationProfilesApi.EditProfile(context.Background(), id).EditProfileDetails(editProfileDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationProfilesApi.EditProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProfile`: Profile
    fmt.Fprintf(os.Stdout, "Response from `CustomisationProfilesApi.EditProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editProfileDetails** | [**EditProfileDetails**](EditProfileDetails.md) |  | 

### Return type

[**Profile**](Profile.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> Profile GetProfile(ctx, id).Execute()

Get a profile



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
    id := "id_example" // string | Identifier of the profile

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomisationProfilesApi.GetProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationProfilesApi.GetProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfile`: Profile
    fmt.Fprintf(os.Stdout, "Response from `CustomisationProfilesApi.GetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Profile**](Profile.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfiles

> ProfilesPage GetProfiles(ctx).Order(order).Page(page).PerPage(perPage).Execute()

Get profiles



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
    resp, r, err := apiClient.CustomisationProfilesApi.GetProfiles(context.Background()).Order(order).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomisationProfilesApi.GetProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfiles`: ProfilesPage
    fmt.Fprintf(os.Stdout, "Response from `CustomisationProfilesApi.GetProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Determines the order of the results. You can order by &lt;code&gt;createdAt&lt;/code&gt; field ascending or descending. The general format is &lt;code&gt;property[,asc|desc]&lt;/code&gt;, with &lt;code&gt;asc&lt;/code&gt; being the default value.The default order is &lt;code&gt;createdAt,asc&lt;/code&gt;. | 
 **page** | **int32** | Page to load | [default to 1]
 **perPage** | **int32** | The number of items on the page | [default to 20]

### Return type

[**ProfilesPage**](ProfilesPage.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

