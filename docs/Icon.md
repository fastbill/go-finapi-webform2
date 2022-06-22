# Icon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to **NullableString** | The image that will be rendered on the Web Form when an info icon needs to be rendered. When it is provided, this image will be used instead of the default info icon.&lt;br/&gt;The image must follow the DATA URI scheme: &lt;code&gt;data:[\\&lt;media type\\&gt;][;base64],\\&lt;data\\&gt;&lt;/code&gt; | [optional] 
**Loading** | Pointer to **NullableString** | The loading icon which will be rendered on the web form while it is importing data in the background. When it is provided, this GIF will be used instead of the default loading animation.&lt;br/&gt;The image must follow the DATA URI scheme: &lt;code&gt;data:[\\&lt;media type\\&gt;][;base64],\\&lt;data\\&gt;&lt;/code&gt; | [optional] 

## Methods

### NewIcon

`func NewIcon() *Icon`

NewIcon instantiates a new Icon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIconWithDefaults

`func NewIconWithDefaults() *Icon`

NewIconWithDefaults instantiates a new Icon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *Icon) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Icon) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Icon) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Icon) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *Icon) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *Icon) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetLoading

`func (o *Icon) GetLoading() string`

GetLoading returns the Loading field if non-nil, zero value otherwise.

### GetLoadingOk

`func (o *Icon) GetLoadingOk() (*string, bool)`

GetLoadingOk returns a tuple with the Loading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoading

`func (o *Icon) SetLoading(v string)`

SetLoading sets Loading field to given value.

### HasLoading

`func (o *Icon) HasLoading() bool`

HasLoading returns a boolean if a field has been set.

### SetLoadingNil

`func (o *Icon) SetLoadingNil(b bool)`

 SetLoadingNil sets the value for Loading to be an explicit nil

### UnsetLoading
`func (o *Icon) UnsetLoading()`

UnsetLoading ensures that no value is present for Loading, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


