# CreateProfileDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label to profile.&lt;br/&gt;The label must be unique. | 
**Default** | Pointer to **NullableBool** | Whether the profile will be used by default for all web forms. Default value is &lt;code&gt;false&lt;/code&gt;.&lt;br/&gt;We urge you to set one profile as default. This way, if you do not pass the profile in the API call, we will use the default profile from you for the web forms.&lt;br/&gt;&lt;br/&gt;There can only be one default profile at one time. | [optional] [default to false]
**Brand** | Pointer to [**NullableBrand**](Brand.md) |  | [optional] 
**Functionality** | Pointer to [**NullableFunctionality**](Functionality.md) |  | [optional] 
**Aspect** | Pointer to [**NullableAspect**](Aspect.md) |  | [optional] 

## Methods

### NewCreateProfileDetails

`func NewCreateProfileDetails(label string, ) *CreateProfileDetails`

NewCreateProfileDetails instantiates a new CreateProfileDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileDetailsWithDefaults

`func NewCreateProfileDetailsWithDefaults() *CreateProfileDetails`

NewCreateProfileDetailsWithDefaults instantiates a new CreateProfileDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateProfileDetails) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateProfileDetails) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateProfileDetails) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDefault

`func (o *CreateProfileDetails) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CreateProfileDetails) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CreateProfileDetails) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CreateProfileDetails) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *CreateProfileDetails) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *CreateProfileDetails) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetBrand

`func (o *CreateProfileDetails) GetBrand() Brand`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreateProfileDetails) GetBrandOk() (*Brand, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreateProfileDetails) SetBrand(v Brand)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreateProfileDetails) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CreateProfileDetails) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CreateProfileDetails) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetFunctionality

`func (o *CreateProfileDetails) GetFunctionality() Functionality`

GetFunctionality returns the Functionality field if non-nil, zero value otherwise.

### GetFunctionalityOk

`func (o *CreateProfileDetails) GetFunctionalityOk() (*Functionality, bool)`

GetFunctionalityOk returns a tuple with the Functionality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionality

`func (o *CreateProfileDetails) SetFunctionality(v Functionality)`

SetFunctionality sets Functionality field to given value.

### HasFunctionality

`func (o *CreateProfileDetails) HasFunctionality() bool`

HasFunctionality returns a boolean if a field has been set.

### SetFunctionalityNil

`func (o *CreateProfileDetails) SetFunctionalityNil(b bool)`

 SetFunctionalityNil sets the value for Functionality to be an explicit nil

### UnsetFunctionality
`func (o *CreateProfileDetails) UnsetFunctionality()`

UnsetFunctionality ensures that no value is present for Functionality, not even an explicit nil
### GetAspect

`func (o *CreateProfileDetails) GetAspect() Aspect`

GetAspect returns the Aspect field if non-nil, zero value otherwise.

### GetAspectOk

`func (o *CreateProfileDetails) GetAspectOk() (*Aspect, bool)`

GetAspectOk returns a tuple with the Aspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspect

`func (o *CreateProfileDetails) SetAspect(v Aspect)`

SetAspect sets Aspect field to given value.

### HasAspect

`func (o *CreateProfileDetails) HasAspect() bool`

HasAspect returns a boolean if a field has been set.

### SetAspectNil

`func (o *CreateProfileDetails) SetAspectNil(b bool)`

 SetAspectNil sets the value for Aspect to be an explicit nil

### UnsetAspect
`func (o *CreateProfileDetails) UnsetAspect()`

UnsetAspect ensures that no value is present for Aspect, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


