# TranslationSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Globally unique translation set&#39;s identifier | 
**CreatedAt** | **string** | The timestamp when the translation set was created in the format &lt;code&gt;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt;. | 
**Cs** | Pointer to [**NullableTranslation**](Translation.md) |  | [optional] 
**De** | Pointer to [**NullableTranslation**](Translation.md) |  | [optional] 
**En** | Pointer to [**NullableTranslation**](Translation.md) |  | [optional] 
**Sk** | Pointer to [**NullableTranslation**](Translation.md) |  | [optional] 

## Methods

### NewTranslationSet

`func NewTranslationSet(id string, createdAt string, ) *TranslationSet`

NewTranslationSet instantiates a new TranslationSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslationSetWithDefaults

`func NewTranslationSetWithDefaults() *TranslationSet`

NewTranslationSetWithDefaults instantiates a new TranslationSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TranslationSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TranslationSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TranslationSet) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TranslationSet) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TranslationSet) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TranslationSet) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCs

`func (o *TranslationSet) GetCs() Translation`

GetCs returns the Cs field if non-nil, zero value otherwise.

### GetCsOk

`func (o *TranslationSet) GetCsOk() (*Translation, bool)`

GetCsOk returns a tuple with the Cs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCs

`func (o *TranslationSet) SetCs(v Translation)`

SetCs sets Cs field to given value.

### HasCs

`func (o *TranslationSet) HasCs() bool`

HasCs returns a boolean if a field has been set.

### SetCsNil

`func (o *TranslationSet) SetCsNil(b bool)`

 SetCsNil sets the value for Cs to be an explicit nil

### UnsetCs
`func (o *TranslationSet) UnsetCs()`

UnsetCs ensures that no value is present for Cs, not even an explicit nil
### GetDe

`func (o *TranslationSet) GetDe() Translation`

GetDe returns the De field if non-nil, zero value otherwise.

### GetDeOk

`func (o *TranslationSet) GetDeOk() (*Translation, bool)`

GetDeOk returns a tuple with the De field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDe

`func (o *TranslationSet) SetDe(v Translation)`

SetDe sets De field to given value.

### HasDe

`func (o *TranslationSet) HasDe() bool`

HasDe returns a boolean if a field has been set.

### SetDeNil

`func (o *TranslationSet) SetDeNil(b bool)`

 SetDeNil sets the value for De to be an explicit nil

### UnsetDe
`func (o *TranslationSet) UnsetDe()`

UnsetDe ensures that no value is present for De, not even an explicit nil
### GetEn

`func (o *TranslationSet) GetEn() Translation`

GetEn returns the En field if non-nil, zero value otherwise.

### GetEnOk

`func (o *TranslationSet) GetEnOk() (*Translation, bool)`

GetEnOk returns a tuple with the En field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEn

`func (o *TranslationSet) SetEn(v Translation)`

SetEn sets En field to given value.

### HasEn

`func (o *TranslationSet) HasEn() bool`

HasEn returns a boolean if a field has been set.

### SetEnNil

`func (o *TranslationSet) SetEnNil(b bool)`

 SetEnNil sets the value for En to be an explicit nil

### UnsetEn
`func (o *TranslationSet) UnsetEn()`

UnsetEn ensures that no value is present for En, not even an explicit nil
### GetSk

`func (o *TranslationSet) GetSk() Translation`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *TranslationSet) GetSkOk() (*Translation, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *TranslationSet) SetSk(v Translation)`

SetSk sets Sk field to given value.

### HasSk

`func (o *TranslationSet) HasSk() bool`

HasSk returns a boolean if a field has been set.

### SetSkNil

`func (o *TranslationSet) SetSkNil(b bool)`

 SetSkNil sets the value for Sk to be an explicit nil

### UnsetSk
`func (o *TranslationSet) UnsetSk()`

UnsetSk ensures that no value is present for Sk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


