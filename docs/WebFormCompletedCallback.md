# WebFormCompletedCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebFormId** | **string** | Web form&#39;s identifier | 
**Status** | [**WebFormStatus**](WebFormStatus.md) |  | 

## Methods

### NewWebFormCompletedCallback

`func NewWebFormCompletedCallback(webFormId string, status WebFormStatus, ) *WebFormCompletedCallback`

NewWebFormCompletedCallback instantiates a new WebFormCompletedCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebFormCompletedCallbackWithDefaults

`func NewWebFormCompletedCallbackWithDefaults() *WebFormCompletedCallback`

NewWebFormCompletedCallbackWithDefaults instantiates a new WebFormCompletedCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebFormId

`func (o *WebFormCompletedCallback) GetWebFormId() string`

GetWebFormId returns the WebFormId field if non-nil, zero value otherwise.

### GetWebFormIdOk

`func (o *WebFormCompletedCallback) GetWebFormIdOk() (*string, bool)`

GetWebFormIdOk returns a tuple with the WebFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebFormId

`func (o *WebFormCompletedCallback) SetWebFormId(v string)`

SetWebFormId sets WebFormId field to given value.


### GetStatus

`func (o *WebFormCompletedCallback) GetStatus() WebFormStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebFormCompletedCallback) GetStatusOk() (*WebFormStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebFormCompletedCallback) SetStatus(v WebFormStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


