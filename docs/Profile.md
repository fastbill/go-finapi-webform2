# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Globally unique profile&#39;s identifier | 
**Label** | **string** | Label to profile | 
**CreatedAt** | **time.Time** | The timestamp when the profile was created in the format &lt;code&gt;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt;. | 
**Default** | **bool** | Whether the profile will be used by default for all web forms.&lt;br/&gt;We urge you to set one profile as default. This way, if you do not pass the profile in the API call, we will use the default profile from you for the web forms. | 
**Brand** | Pointer to [**NullableBrand**](Brand.md) |  | [optional] 
**Functionality** | Pointer to [**NullableFunctionality**](Functionality.md) |  | [optional] 
**Aspect** | Pointer to [**NullableAspect**](Aspect.md) |  | [optional] 

## Methods

### NewProfile

`func NewProfile(id string, label string, createdAt time.Time, default_ bool, ) *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Profile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Profile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Profile) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *Profile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Profile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Profile) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedAt

`func (o *Profile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Profile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Profile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefault

`func (o *Profile) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Profile) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Profile) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetBrand

`func (o *Profile) GetBrand() Brand`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Profile) GetBrandOk() (*Brand, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Profile) SetBrand(v Brand)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Profile) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *Profile) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *Profile) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetFunctionality

`func (o *Profile) GetFunctionality() Functionality`

GetFunctionality returns the Functionality field if non-nil, zero value otherwise.

### GetFunctionalityOk

`func (o *Profile) GetFunctionalityOk() (*Functionality, bool)`

GetFunctionalityOk returns a tuple with the Functionality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionality

`func (o *Profile) SetFunctionality(v Functionality)`

SetFunctionality sets Functionality field to given value.

### HasFunctionality

`func (o *Profile) HasFunctionality() bool`

HasFunctionality returns a boolean if a field has been set.

### SetFunctionalityNil

`func (o *Profile) SetFunctionalityNil(b bool)`

 SetFunctionalityNil sets the value for Functionality to be an explicit nil

### UnsetFunctionality
`func (o *Profile) UnsetFunctionality()`

UnsetFunctionality ensures that no value is present for Functionality, not even an explicit nil
### GetAspect

`func (o *Profile) GetAspect() Aspect`

GetAspect returns the Aspect field if non-nil, zero value otherwise.

### GetAspectOk

`func (o *Profile) GetAspectOk() (*Aspect, bool)`

GetAspectOk returns a tuple with the Aspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspect

`func (o *Profile) SetAspect(v Aspect)`

SetAspect sets Aspect field to given value.

### HasAspect

`func (o *Profile) HasAspect() bool`

HasAspect returns a boolean if a field has been set.

### SetAspectNil

`func (o *Profile) SetAspectNil(b bool)`

 SetAspectNil sets the value for Aspect to be an explicit nil

### UnsetAspect
`func (o *Profile) UnsetAspect()`

UnsetAspect ensures that no value is present for Aspect, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


