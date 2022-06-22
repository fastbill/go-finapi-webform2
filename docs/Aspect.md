# Aspect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorScheme** | Pointer to [**NullableColor**](Color.md) |  | [optional] 
**Text** | Pointer to [**NullableText**](Text.md) |  | [optional] 
**Theme** | Pointer to **NullableString** | &lt;strong&gt;Allowed:&lt;/strong&gt; DEFAULT | MOBILE_MINIMALISTIC&lt;br/&gt;&lt;strong&gt;BETA&lt;/strong&gt;: This service is currently in BETA phase. Provided themes can be changed and/or deleted before the service goes live.&lt;br/&gt;Theme will provide further granularity with the styling possibilities for customers. This includes styling of buttons, text fields, text spacing, sizing of titles, and subtitles amongst others.&lt;br/&gt;Find more info in the &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/themes&#39; target&#x3D;&#39;_blank&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | [optional] 

## Methods

### NewAspect

`func NewAspect() *Aspect`

NewAspect instantiates a new Aspect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectWithDefaults

`func NewAspectWithDefaults() *Aspect`

NewAspectWithDefaults instantiates a new Aspect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColorScheme

`func (o *Aspect) GetColorScheme() Color`

GetColorScheme returns the ColorScheme field if non-nil, zero value otherwise.

### GetColorSchemeOk

`func (o *Aspect) GetColorSchemeOk() (*Color, bool)`

GetColorSchemeOk returns a tuple with the ColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorScheme

`func (o *Aspect) SetColorScheme(v Color)`

SetColorScheme sets ColorScheme field to given value.

### HasColorScheme

`func (o *Aspect) HasColorScheme() bool`

HasColorScheme returns a boolean if a field has been set.

### SetColorSchemeNil

`func (o *Aspect) SetColorSchemeNil(b bool)`

 SetColorSchemeNil sets the value for ColorScheme to be an explicit nil

### UnsetColorScheme
`func (o *Aspect) UnsetColorScheme()`

UnsetColorScheme ensures that no value is present for ColorScheme, not even an explicit nil
### GetText

`func (o *Aspect) GetText() Text`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Aspect) GetTextOk() (*Text, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Aspect) SetText(v Text)`

SetText sets Text field to given value.

### HasText

`func (o *Aspect) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *Aspect) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *Aspect) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetTheme

`func (o *Aspect) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Aspect) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Aspect) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Aspect) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *Aspect) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *Aspect) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


