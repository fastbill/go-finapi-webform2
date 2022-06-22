# TaskCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Task&#39;s identifier | 
**Status** | [**TaskStatus**](TaskStatus.md) |  | 
**WebForm** | Pointer to [**NullableWebFormInfo**](WebFormInfo.md) |  | [optional] 

## Methods

### NewTaskCallback

`func NewTaskCallback(taskId string, status TaskStatus, ) *TaskCallback`

NewTaskCallback instantiates a new TaskCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCallbackWithDefaults

`func NewTaskCallbackWithDefaults() *TaskCallback`

NewTaskCallbackWithDefaults instantiates a new TaskCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskCallback) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskCallback) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskCallback) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *TaskCallback) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskCallback) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskCallback) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.


### GetWebForm

`func (o *TaskCallback) GetWebForm() WebFormInfo`

GetWebForm returns the WebForm field if non-nil, zero value otherwise.

### GetWebFormOk

`func (o *TaskCallback) GetWebFormOk() (*WebFormInfo, bool)`

GetWebFormOk returns a tuple with the WebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebForm

`func (o *TaskCallback) SetWebForm(v WebFormInfo)`

SetWebForm sets WebForm field to given value.

### HasWebForm

`func (o *TaskCallback) HasWebForm() bool`

HasWebForm returns a boolean if a field has been set.

### SetWebFormNil

`func (o *TaskCallback) SetWebFormNil(b bool)`

 SetWebFormNil sets the value for WebForm to be an explicit nil

### UnsetWebForm
`func (o *TaskCallback) UnsetWebForm()`

UnsetWebForm ensures that no value is present for WebForm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


