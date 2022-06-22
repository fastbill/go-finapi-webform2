# Brand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | Pointer to **NullableString** | Logo of minimum height 40px that will be shown in the web form header.Bigger images are also accepted, but will be scaled down.&lt;br/&gt;The image must follow the DATA URI scheme: &lt;code&gt;data:[\\&lt;media type\\&gt;][;base64],\\&lt;data\\&gt;&lt;/code&gt;&lt;br/&gt;&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; For &#x60;UNLICENSED&#x60; mandators, if no value is provided, then the following value will be applied by default when web form is opened: &lt;br/&gt;&lt;code&gt;finAPI GmbH Logo&lt;/code&gt; | [optional] 
**Favicon** | Pointer to **NullableString** | Logo of minimum size 32x32px that will be shown in the browser tab favicon.&lt;br/&gt;The image must follow the DATA URI scheme: &lt;code&gt;data:[\\&lt;media type\\&gt;][;base64],\\&lt;data\\&gt;&lt;/code&gt;&lt;br/&gt;&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; For &#x60;UNLICENSED&#x60; mandators, if no value is provided, then the following value will be applied by default when web form is opened: &lt;br/&gt;&lt;code&gt;finAPI GmbH favicon&lt;/code&gt; | [optional] 
**Icon** | Pointer to [**NullableIcon**](Icon.md) |  | [optional] 
**IntroText** | Pointer to **NullableString** | An introduction text that will be shown on the first web form view for bank selection.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;null&lt;/code&gt;. | [optional] 

## Methods

### NewBrand

`func NewBrand() *Brand`

NewBrand instantiates a new Brand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandWithDefaults

`func NewBrandWithDefaults() *Brand`

NewBrandWithDefaults instantiates a new Brand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *Brand) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Brand) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Brand) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Brand) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *Brand) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *Brand) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetFavicon

`func (o *Brand) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *Brand) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *Brand) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *Brand) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### SetFaviconNil

`func (o *Brand) SetFaviconNil(b bool)`

 SetFaviconNil sets the value for Favicon to be an explicit nil

### UnsetFavicon
`func (o *Brand) UnsetFavicon()`

UnsetFavicon ensures that no value is present for Favicon, not even an explicit nil
### GetIcon

`func (o *Brand) GetIcon() Icon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Brand) GetIconOk() (*Icon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Brand) SetIcon(v Icon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Brand) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *Brand) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *Brand) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetIntroText

`func (o *Brand) GetIntroText() string`

GetIntroText returns the IntroText field if non-nil, zero value otherwise.

### GetIntroTextOk

`func (o *Brand) GetIntroTextOk() (*string, bool)`

GetIntroTextOk returns a tuple with the IntroText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroText

`func (o *Brand) SetIntroText(v string)`

SetIntroText sets IntroText field to given value.

### HasIntroText

`func (o *Brand) HasIntroText() bool`

HasIntroText returns a boolean if a field has been set.

### SetIntroTextNil

`func (o *Brand) SetIntroTextNil(b bool)`

 SetIntroTextNil sets the value for IntroText to be an explicit nil

### UnsetIntroText
`func (o *Brand) UnsetIntroText()`

UnsetIntroText ensures that no value is present for IntroText, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


