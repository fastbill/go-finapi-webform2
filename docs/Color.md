# Color

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **NullableString** | The color of your brand. It will be applied to main page header/footer, input field borders and main button.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;#00ADDF&lt;/code | [optional] 
**Text** | Pointer to [**NullableTextColor**](TextColor.md) |  | [optional] 

## Methods

### NewColor

`func NewColor() *Color`

NewColor instantiates a new Color object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColorWithDefaults

`func NewColorWithDefaults() *Color`

NewColorWithDefaults instantiates a new Color object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *Color) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Color) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Color) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Color) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *Color) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *Color) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetText

`func (o *Color) GetText() TextColor`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Color) GetTextOk() (*TextColor, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Color) SetText(v TextColor)`

SetText sets Text field to given value.

### HasText

`func (o *Color) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *Color) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *Color) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


