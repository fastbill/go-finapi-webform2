# \PaymentInitiationServicesApi

All URIs are relative to *https://webform-sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForDirectDebitWithAccountId**](PaymentInitiationServicesApi.md#CreateForDirectDebitWithAccountId) | **Post** /api/webForms/directDebitWithAccountId | Create a direct debit with account ID
[**CreateForPaymentWithAccountId**](PaymentInitiationServicesApi.md#CreateForPaymentWithAccountId) | **Post** /api/webForms/paymentWithAccountId | Create a payment with account ID
[**CreateForStandalonePayment**](PaymentInitiationServicesApi.md#CreateForStandalonePayment) | **Post** /api/webForms/standalonePayment | Create a standalone payment
[**CreateForStandingOrder**](PaymentInitiationServicesApi.md#CreateForStandingOrder) | **Post** /api/webForms/standingOrder | Create a standing order



## CreateForDirectDebitWithAccountId

> WebForm CreateForDirectDebitWithAccountId(ctx).DirectDebitWithAccountDetails(directDebitWithAccountDetails).Execute()

Create a direct debit with account ID



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
    directDebitWithAccountDetails := *openapiclient.NewDirectDebitWithAccountDetails(time.Now(), []openapiclient.DirectDebitOrder{*openapiclient.NewDirectDebitOrder(*openapiclient.NewAmount(float32(0.04), openapiclient.Currency("AED")), *openapiclient.NewPaymentPayer("DE77533700080111111100", "Max Mustermann"), "1", time.Now(), "DE98ZZZ09999999999")}, *openapiclient.NewDirectDebitWithAccountReceiver(int64(42)), "B2B", "OOFF") // DirectDebitWithAccountDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentInitiationServicesApi.CreateForDirectDebitWithAccountId(context.Background()).DirectDebitWithAccountDetails(directDebitWithAccountDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentInitiationServicesApi.CreateForDirectDebitWithAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateForDirectDebitWithAccountId`: WebForm
    fmt.Fprintf(os.Stdout, "Response from `PaymentInitiationServicesApi.CreateForDirectDebitWithAccountId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForDirectDebitWithAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directDebitWithAccountDetails** | [**DirectDebitWithAccountDetails**](DirectDebitWithAccountDetails.md) |  | 

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


## CreateForPaymentWithAccountId

> WebForm CreateForPaymentWithAccountId(ctx).PaymentWithAccountDetails(paymentWithAccountDetails).Execute()

Create a payment with account ID



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
    paymentWithAccountDetails := *openapiclient.NewPaymentWithAccountDetails(*openapiclient.NewPaymentWithAccountSender(int64(42)), []openapiclient.PaymentOrder{*openapiclient.NewPaymentOrder(*openapiclient.NewAmount(float32(0.04), openapiclient.Currency("AED")), *openapiclient.NewPaymentRecipient("DE77533700080111111100", "Max Mustermann"))}) // PaymentWithAccountDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentInitiationServicesApi.CreateForPaymentWithAccountId(context.Background()).PaymentWithAccountDetails(paymentWithAccountDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentInitiationServicesApi.CreateForPaymentWithAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateForPaymentWithAccountId`: WebForm
    fmt.Fprintf(os.Stdout, "Response from `PaymentInitiationServicesApi.CreateForPaymentWithAccountId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForPaymentWithAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentWithAccountDetails** | [**PaymentWithAccountDetails**](PaymentWithAccountDetails.md) |  | 

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


## CreateForStandalonePayment

> WebForm CreateForStandalonePayment(ctx).StandalonePaymentDetails(standalonePaymentDetails).Execute()

Create a standalone payment



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
    standalonePaymentDetails := *openapiclient.NewStandalonePaymentDetails([]openapiclient.PaymentOrder{*openapiclient.NewPaymentOrder(*openapiclient.NewAmount(float32(0.04), openapiclient.Currency("AED")), *openapiclient.NewPaymentRecipient("DE77533700080111111100", "Max Mustermann"))}) // StandalonePaymentDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentInitiationServicesApi.CreateForStandalonePayment(context.Background()).StandalonePaymentDetails(standalonePaymentDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentInitiationServicesApi.CreateForStandalonePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateForStandalonePayment`: WebForm
    fmt.Fprintf(os.Stdout, "Response from `PaymentInitiationServicesApi.CreateForStandalonePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForStandalonePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standalonePaymentDetails** | [**StandalonePaymentDetails**](StandalonePaymentDetails.md) |  | 

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


## CreateForStandingOrder

> WebForm CreateForStandingOrder(ctx).StandingOrderDetails(standingOrderDetails).Execute()

Create a standing order



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
    standingOrderDetails := *openapiclient.NewStandingOrderDetails(*openapiclient.NewAmount(float32(0.04), openapiclient.Currency("AED")), *openapiclient.NewRecipient("DE77533700080111111100", "Max Mustermann"), time.Now(), "MONTHLY") // StandingOrderDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentInitiationServicesApi.CreateForStandingOrder(context.Background()).StandingOrderDetails(standingOrderDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentInitiationServicesApi.CreateForStandingOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateForStandingOrder`: WebForm
    fmt.Fprintf(os.Stdout, "Response from `PaymentInitiationServicesApi.CreateForStandingOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForStandingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standingOrderDetails** | [**StandingOrderDetails**](StandingOrderDetails.md) |  | 

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

