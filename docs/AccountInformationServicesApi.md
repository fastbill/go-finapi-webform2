# \AccountInformationServicesApi

All URIs are relative to *https://webform-sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBankConnectionUpdateTask**](AccountInformationServicesApi.md#CreateBankConnectionUpdateTask) | **Post** /api/tasks/backgroundUpdate | Update a bank connection
[**CreateForBankConnectionImport**](AccountInformationServicesApi.md#CreateForBankConnectionImport) | **Post** /api/webForms/bankConnectionImport | Import a bank connection
[**CreateForBankConnectionUpdate**](AccountInformationServicesApi.md#CreateForBankConnectionUpdate) | **Post** /api/webForms/bankConnectionUpdate | Update a bank connection (BETA - DEPRECATED)



## CreateBankConnectionUpdateTask

> Task CreateBankConnectionUpdateTask(ctx).BankConnectionUpdateTaskDetails(bankConnectionUpdateTaskDetails).Execute()

Update a bank connection



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
    bankConnectionUpdateTaskDetails := *openapiclient.NewBankConnectionUpdateTaskDetails(int64(101)) // BankConnectionUpdateTaskDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountInformationServicesApi.CreateBankConnectionUpdateTask(context.Background()).BankConnectionUpdateTaskDetails(bankConnectionUpdateTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountInformationServicesApi.CreateBankConnectionUpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBankConnectionUpdateTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `AccountInformationServicesApi.CreateBankConnectionUpdateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankConnectionUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankConnectionUpdateTaskDetails** | [**BankConnectionUpdateTaskDetails**](BankConnectionUpdateTaskDetails.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateForBankConnectionImport

> WebForm CreateForBankConnectionImport(ctx).BankConnectionImportDetails(bankConnectionImportDetails).Execute()

Import a bank connection



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
    bankConnectionImportDetails := *openapiclient.NewBankConnectionImportDetails() // BankConnectionImportDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountInformationServicesApi.CreateForBankConnectionImport(context.Background()).BankConnectionImportDetails(bankConnectionImportDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountInformationServicesApi.CreateForBankConnectionImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateForBankConnectionImport`: WebForm
    fmt.Fprintf(os.Stdout, "Response from `AccountInformationServicesApi.CreateForBankConnectionImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForBankConnectionImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankConnectionImportDetails** | [**BankConnectionImportDetails**](BankConnectionImportDetails.md) |  | 

### Return type

[**WebForm**](WebForm.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateForBankConnectionUpdate

> WebForm CreateForBankConnectionUpdate(ctx).BankConnectionUpdateDetails(bankConnectionUpdateDetails).Execute()

Update a bank connection (BETA - DEPRECATED)



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
    bankConnectionUpdateDetails := *openapiclient.NewBankConnectionUpdateDetails(*openapiclient.NewUpdateBankDetails(int64(101))) // BankConnectionUpdateDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountInformationServicesApi.CreateForBankConnectionUpdate(context.Background()).BankConnectionUpdateDetails(bankConnectionUpdateDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountInformationServicesApi.CreateForBankConnectionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateForBankConnectionUpdate`: WebForm
    fmt.Fprintf(os.Stdout, "Response from `AccountInformationServicesApi.CreateForBankConnectionUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForBankConnectionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankConnectionUpdateDetails** | [**BankConnectionUpdateDetails**](BankConnectionUpdateDetails.md) |  | 

### Return type

[**WebForm**](WebForm.md)

### Authorization

[BearerAccessToken](../README.md#BearerAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

