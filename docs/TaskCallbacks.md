# TaskCallbacks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finalised** | Pointer to **NullableString** | URL to which finAPI sends the notification message after the task is finalised (either completed successfully or with an error).&lt;br/&gt;finAPI will call this URL with HTTP method POST ignoring the response of the target endpoint. Only HTTPS URLs are allowed. | [optional] 
**WebFormRequired** | Pointer to **NullableString** | URL to which finAPI sends the notification message when the update can longer run in the background.&lt;br/&gt;finAPI will call this URL with HTTP method POST ignoring the response of the target endpoint. Only HTTPS URLs are allowed. | [optional] 

## Methods

### NewTaskCallbacks

`func NewTaskCallbacks() *TaskCallbacks`

NewTaskCallbacks instantiates a new TaskCallbacks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCallbacksWithDefaults

`func NewTaskCallbacksWithDefaults() *TaskCallbacks`

NewTaskCallbacksWithDefaults instantiates a new TaskCallbacks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalised

`func (o *TaskCallbacks) GetFinalised() string`

GetFinalised returns the Finalised field if non-nil, zero value otherwise.

### GetFinalisedOk

`func (o *TaskCallbacks) GetFinalisedOk() (*string, bool)`

GetFinalisedOk returns a tuple with the Finalised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalised

`func (o *TaskCallbacks) SetFinalised(v string)`

SetFinalised sets Finalised field to given value.

### HasFinalised

`func (o *TaskCallbacks) HasFinalised() bool`

HasFinalised returns a boolean if a field has been set.

### SetFinalisedNil

`func (o *TaskCallbacks) SetFinalisedNil(b bool)`

 SetFinalisedNil sets the value for Finalised to be an explicit nil

### UnsetFinalised
`func (o *TaskCallbacks) UnsetFinalised()`

UnsetFinalised ensures that no value is present for Finalised, not even an explicit nil
### GetWebFormRequired

`func (o *TaskCallbacks) GetWebFormRequired() string`

GetWebFormRequired returns the WebFormRequired field if non-nil, zero value otherwise.

### GetWebFormRequiredOk

`func (o *TaskCallbacks) GetWebFormRequiredOk() (*string, bool)`

GetWebFormRequiredOk returns a tuple with the WebFormRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebFormRequired

`func (o *TaskCallbacks) SetWebFormRequired(v string)`

SetWebFormRequired sets WebFormRequired field to given value.

### HasWebFormRequired

`func (o *TaskCallbacks) HasWebFormRequired() bool`

HasWebFormRequired returns a boolean if a field has been set.

### SetWebFormRequiredNil

`func (o *TaskCallbacks) SetWebFormRequiredNil(b bool)`

 SetWebFormRequiredNil sets the value for WebFormRequired to be an explicit nil

### UnsetWebFormRequired
`func (o *TaskCallbacks) UnsetWebFormRequired()`

UnsetWebFormRequired ensures that no value is present for WebFormRequired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


