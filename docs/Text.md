# Text

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FontFamily** | Pointer to **NullableString** | List of comma-separated fonts. The fonts will be chosen in the order they are given. If the font cannot be rendered on the end-user OS, the next font will be chosen.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;Calibri,Roboto,\&quot;Segoe UI\&quot;,\&quot;Helvetica Neue\&quot;&lt;/code&gt; | [optional] 

## Methods

### NewText

`func NewText() *Text`

NewText instantiates a new Text object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextWithDefaults

`func NewTextWithDefaults() *Text`

NewTextWithDefaults instantiates a new Text object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFontFamily

`func (o *Text) GetFontFamily() string`

GetFontFamily returns the FontFamily field if non-nil, zero value otherwise.

### GetFontFamilyOk

`func (o *Text) GetFontFamilyOk() (*string, bool)`

GetFontFamilyOk returns a tuple with the FontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontFamily

`func (o *Text) SetFontFamily(v string)`

SetFontFamily sets FontFamily field to given value.

### HasFontFamily

`func (o *Text) HasFontFamily() bool`

HasFontFamily returns a boolean if a field has been set.

### SetFontFamilyNil

`func (o *Text) SetFontFamilyNil(b bool)`

 SetFontFamilyNil sets the value for FontFamily to be an explicit nil

### UnsetFontFamily
`func (o *Text) UnsetFontFamily()`

UnsetFontFamily ensures that no value is present for FontFamily, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


