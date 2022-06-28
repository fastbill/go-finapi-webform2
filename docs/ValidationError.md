# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]InvalidFieldError**](InvalidFieldError.md) | Fields that caused the error | 
**Timestamp** | **string** | Timestamp when the error occurred in the format &lt;code&gt;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt;. | 
**Endpoint** | **string** | Endpoint caused the error | 

## Methods

### NewValidationError

`func NewValidationError(errors []InvalidFieldError, timestamp string, endpoint string, ) *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ValidationError) GetErrors() []InvalidFieldError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidationError) GetErrorsOk() (*[]InvalidFieldError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidationError) SetErrors(v []InvalidFieldError)`

SetErrors sets Errors field to given value.


### GetTimestamp

`func (o *ValidationError) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ValidationError) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ValidationError) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetEndpoint

`func (o *ValidationError) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ValidationError) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ValidationError) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


