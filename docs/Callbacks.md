# Callbacks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finalised** | Pointer to **NullableString** | URL to which finAPI sends the notification message after the web form is finalised (either completed successfully or with an error, or aborted by the user).&lt;br/&gt;finAPI will call this URL with HTTP method POST ignoring the response of the target endpoint. Only HTTPS URLs are allowed. | [optional] 

## Methods

### NewCallbacks

`func NewCallbacks() *Callbacks`

NewCallbacks instantiates a new Callbacks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbacksWithDefaults

`func NewCallbacksWithDefaults() *Callbacks`

NewCallbacksWithDefaults instantiates a new Callbacks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalised

`func (o *Callbacks) GetFinalised() string`

GetFinalised returns the Finalised field if non-nil, zero value otherwise.

### GetFinalisedOk

`func (o *Callbacks) GetFinalisedOk() (*string, bool)`

GetFinalisedOk returns a tuple with the Finalised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalised

`func (o *Callbacks) SetFinalised(v string)`

SetFinalised sets Finalised field to given value.

### HasFinalised

`func (o *Callbacks) HasFinalised() bool`

HasFinalised returns a boolean if a field has been set.

### SetFinalisedNil

`func (o *Callbacks) SetFinalisedNil(b bool)`

 SetFinalisedNil sets the value for Finalised to be an explicit nil

### UnsetFinalised
`func (o *Callbacks) UnsetFinalised()`

UnsetFinalised ensures that no value is present for Finalised, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


