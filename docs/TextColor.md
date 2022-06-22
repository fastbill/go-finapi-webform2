# TextColor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to **NullableString** | The primary color of the text.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;#232323&lt;/code&gt; | [optional] 
**Secondary** | Pointer to **NullableString** | The secondary color of the text (e.g. login hint, summaries, etc). Can be the same as primary color.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;#848484&lt;/code&gt; | [optional] 

## Methods

### NewTextColor

`func NewTextColor() *TextColor`

NewTextColor instantiates a new TextColor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextColorWithDefaults

`func NewTextColorWithDefaults() *TextColor`

NewTextColorWithDefaults instantiates a new TextColor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *TextColor) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *TextColor) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *TextColor) SetPrimary(v string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *TextColor) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *TextColor) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *TextColor) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetSecondary

`func (o *TextColor) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *TextColor) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *TextColor) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *TextColor) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### SetSecondaryNil

`func (o *TextColor) SetSecondaryNil(b bool)`

 SetSecondaryNil sets the value for Secondary to be an explicit nil

### UnsetSecondary
`func (o *TextColor) UnsetSecondary()`

UnsetSecondary ensures that no value is present for Secondary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


