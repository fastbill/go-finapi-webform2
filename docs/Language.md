# Language

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to **NullableString** | Whether the language selector in the web form header will be displayed.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; in case &lt;code&gt;functionality.header&lt;/code&gt; is marked as &lt;code&gt;HIDDEN&lt;/code&gt;, then this property will be ignored.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;RENDER&lt;/code&gt;. | [optional] 
**Locked** | Pointer to **NullableString** | The language in which the web form will be shown. The language must be given in the &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://www.iso.org/iso-639-language-codes.html&#39;&gt;ISO-639&lt;/a&gt; 2 letter code format.&lt;br/&gt;&amp;bull; This must be initialized if either the &lt;code&gt;functionality.header&lt;/code&gt; or &lt;code&gt;functionality.language.selector&lt;/code&gt; are &lt;code&gt;HIDDEN&lt;/code&gt;;&lt;br/&gt;&amp;bull; This must be null if both &lt;code&gt;functionality.header&lt;/code&gt; and &lt;code&gt;functionality.language.selector&lt;/code&gt; are set to &lt;code&gt;RENDER&lt;/code&gt;.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;null&lt;/code&gt;. | [optional] 

## Methods

### NewLanguage

`func NewLanguage() *Language`

NewLanguage instantiates a new Language object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageWithDefaults

`func NewLanguageWithDefaults() *Language`

NewLanguageWithDefaults instantiates a new Language object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *Language) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *Language) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *Language) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *Language) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *Language) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *Language) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetLocked

`func (o *Language) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Language) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Language) SetLocked(v string)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Language) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *Language) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *Language) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


