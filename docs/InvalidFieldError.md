# InvalidFieldError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field that caused the error | 
**Message** | **string** | Error message | 

## Methods

### NewInvalidFieldError

`func NewInvalidFieldError(field string, message string, ) *InvalidFieldError`

NewInvalidFieldError instantiates a new InvalidFieldError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidFieldErrorWithDefaults

`func NewInvalidFieldErrorWithDefaults() *InvalidFieldError`

NewInvalidFieldErrorWithDefaults instantiates a new InvalidFieldError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *InvalidFieldError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *InvalidFieldError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *InvalidFieldError) SetField(v string)`

SetField sets Field field to given value.


### GetMessage

`func (o *InvalidFieldError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvalidFieldError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvalidFieldError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


