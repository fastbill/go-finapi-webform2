# TranslationBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | Custom view title.&lt;br/&gt;&lt;br/&gt;If no value is provided, the default finAPI translation will be applied by when a web form is opened. &lt;br/&gt;&lt;br/&gt;If an empty string is provided as the value, the section will be hidden. | [optional] 
**Subtitle** | Pointer to **NullableString** | Custom view subtitle.&lt;br/&gt;&lt;br/&gt;If no value is provided, the default finAPI translation will be applied by when a web form is opened. &lt;br/&gt;&lt;br/&gt;If an empty string is provided as the value, the section will be hidden. | [optional] 

## Methods

### NewTranslationBlock

`func NewTranslationBlock() *TranslationBlock`

NewTranslationBlock instantiates a new TranslationBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslationBlockWithDefaults

`func NewTranslationBlockWithDefaults() *TranslationBlock`

NewTranslationBlockWithDefaults instantiates a new TranslationBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TranslationBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TranslationBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TranslationBlock) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TranslationBlock) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TranslationBlock) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TranslationBlock) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSubtitle

`func (o *TranslationBlock) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *TranslationBlock) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *TranslationBlock) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *TranslationBlock) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### SetSubtitleNil

`func (o *TranslationBlock) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *TranslationBlock) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


