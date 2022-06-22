# Functionality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgressBar** | Pointer to **NullableString** | Whether a progress bar is shown on the web form, letting the user know on which step he is.&lt;br/&gt;&amp;bull; &lt;code&gt;RENDER&lt;/code&gt; - the progress bar will be shown;&lt;br/&gt;&amp;bull; &lt;code&gt;HIDDEN&lt;/code&gt; - the progress bar will NOT be shown.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;RENDER&lt;/code&gt;. | [optional] 
**BankLoginHint** | Pointer to **NullableString** | How the bank login hint will be shown to the end user&lt;br/&gt;&amp;bull; &lt;code&gt;EXPANDED&lt;/code&gt; - the user will see the login hint and will have the option to collapse it;&lt;br/&gt;&amp;bull; &lt;code&gt;COLLAPSED&lt;/code&gt; - the login hint will be collapsed and the user can see it if he expands the field;&lt;br/&gt;&amp;bull; &lt;code&gt;HIDDEN&lt;/code&gt; - the login hint is hidden and the user cannot see it.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;EXPANDED&lt;/code&gt;. | [optional] 
**StoreSecrets** | Pointer to **NullableString** | Whether the user will have a checkbox to ask for storing login secrets (like a PIN) in finAPI or not.&lt;br/&gt;&amp;bull; &lt;code&gt;RENDER&lt;/code&gt; - the checkbox will be shown;&lt;br/&gt;&amp;bull; &lt;code&gt;HIDDEN&lt;/code&gt; - the checkbox will NOT be shown;&lt;br/&gt;&amp;bull; &lt;code&gt;MANDATORY&lt;/code&gt; - the checkbox will be shown and it will be mandatory for the end user to check it in order to continue.&lt;br/&gt;&amp;bull; &lt;code&gt;IMPLICIT_APPROVAL&lt;/code&gt; - the checkbox will NOT be shown but login secrets will nevertheless be stored;&lt;br/&gt;&amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;strong&gt;NOTE:&lt;/strong&gt; This value will also automatically store the TAN method. This value can be applied ONLY by our support team. Please contact &lt;a href&#x3D;&#39;mailto:support@finapi.io&#39;&gt;support@finapi.io&lt;/a&gt; with the &lt;code&gt;profile.id&lt;/code&gt; as soon as you have finalized the customization for other parameters.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;RENDER&lt;/code&gt;. | [optional] 
**BankDetails** | Pointer to **NullableString** | Whether the user will be allowed to change the selected bank, in case a BLZ/BIC/IBAN was sent in the API request by the client.&lt;br/&gt;&amp;bull; &lt;code&gt;LOCKED&lt;/code&gt; - the user will be directly routed to login to the pre-selected bank;&lt;br/&gt;&amp;bull; &lt;code&gt;EDITABLE&lt;/code&gt; - the user will see the pre-selected bank and have the option to change it.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;LOCKED&lt;/code&gt;. | [optional] 
**Header** | Pointer to **NullableString** | Whether the header will be displayed on the web form.&lt;br/&gt;&amp;bull; &lt;code&gt;RENDER&lt;/code&gt; - the header will be shown;&lt;br/&gt;&amp;bull; &lt;code&gt;HIDDEN&lt;/code&gt; - the header will NOT be shown.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;RENDER&lt;/code&gt;. | [optional] 
**Language** | Pointer to [**NullableLanguage**](Language.md) |  | [optional] 
**SkipConfirmationView** | Pointer to **NullableBool** | When the web form is completed successfully, it determines whether the last view will be rendered. It applies to embedded and standalone web forms. It also applies to all endpoints in the \&quot;Account Information Services\&quot; and \&quot;Payment Initiation Services\&quot;.&lt;br/&gt;If you are embedding the web form in your application, please set up appropriate handling for the &#39;onComplete&#39; method to take advantage of the feature.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;false&lt;/code&gt; | [optional] 
**RenderAccountSelectionView** | Pointer to **NullableBool** | Whether or not the Web Form will render the \&quot;Account Selection View\&quot; for the end-user to choose which of the imported accounts should be saved to the database and available on the customer application.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;false&lt;/code&gt; | [optional] 
**HidePaymentSummary** | Pointer to **NullableBool** | Whether the entire payment summary is rendered on the Web Form. When set to TRUE, the counterpart data is not rendered on the Payment Summary of the Web Form.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; This value can be applied ONLY by our support team. Please contact &lt;a href&#x3D;&#39;mailto:support@finapi.io&#39;&gt;support@finapi.io&lt;/a&gt; with the &lt;code&gt;profile.id&lt;/code&gt; as soon as you have finalized the customization for other parameters.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If no value is provided, then the following value will be applied by default when web form is opened: &lt;code&gt;false&lt;/code&gt; | [optional] 

## Methods

### NewFunctionality

`func NewFunctionality() *Functionality`

NewFunctionality instantiates a new Functionality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionalityWithDefaults

`func NewFunctionalityWithDefaults() *Functionality`

NewFunctionalityWithDefaults instantiates a new Functionality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgressBar

`func (o *Functionality) GetProgressBar() string`

GetProgressBar returns the ProgressBar field if non-nil, zero value otherwise.

### GetProgressBarOk

`func (o *Functionality) GetProgressBarOk() (*string, bool)`

GetProgressBarOk returns a tuple with the ProgressBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressBar

`func (o *Functionality) SetProgressBar(v string)`

SetProgressBar sets ProgressBar field to given value.

### HasProgressBar

`func (o *Functionality) HasProgressBar() bool`

HasProgressBar returns a boolean if a field has been set.

### SetProgressBarNil

`func (o *Functionality) SetProgressBarNil(b bool)`

 SetProgressBarNil sets the value for ProgressBar to be an explicit nil

### UnsetProgressBar
`func (o *Functionality) UnsetProgressBar()`

UnsetProgressBar ensures that no value is present for ProgressBar, not even an explicit nil
### GetBankLoginHint

`func (o *Functionality) GetBankLoginHint() string`

GetBankLoginHint returns the BankLoginHint field if non-nil, zero value otherwise.

### GetBankLoginHintOk

`func (o *Functionality) GetBankLoginHintOk() (*string, bool)`

GetBankLoginHintOk returns a tuple with the BankLoginHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankLoginHint

`func (o *Functionality) SetBankLoginHint(v string)`

SetBankLoginHint sets BankLoginHint field to given value.

### HasBankLoginHint

`func (o *Functionality) HasBankLoginHint() bool`

HasBankLoginHint returns a boolean if a field has been set.

### SetBankLoginHintNil

`func (o *Functionality) SetBankLoginHintNil(b bool)`

 SetBankLoginHintNil sets the value for BankLoginHint to be an explicit nil

### UnsetBankLoginHint
`func (o *Functionality) UnsetBankLoginHint()`

UnsetBankLoginHint ensures that no value is present for BankLoginHint, not even an explicit nil
### GetStoreSecrets

`func (o *Functionality) GetStoreSecrets() string`

GetStoreSecrets returns the StoreSecrets field if non-nil, zero value otherwise.

### GetStoreSecretsOk

`func (o *Functionality) GetStoreSecretsOk() (*string, bool)`

GetStoreSecretsOk returns a tuple with the StoreSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSecrets

`func (o *Functionality) SetStoreSecrets(v string)`

SetStoreSecrets sets StoreSecrets field to given value.

### HasStoreSecrets

`func (o *Functionality) HasStoreSecrets() bool`

HasStoreSecrets returns a boolean if a field has been set.

### SetStoreSecretsNil

`func (o *Functionality) SetStoreSecretsNil(b bool)`

 SetStoreSecretsNil sets the value for StoreSecrets to be an explicit nil

### UnsetStoreSecrets
`func (o *Functionality) UnsetStoreSecrets()`

UnsetStoreSecrets ensures that no value is present for StoreSecrets, not even an explicit nil
### GetBankDetails

`func (o *Functionality) GetBankDetails() string`

GetBankDetails returns the BankDetails field if non-nil, zero value otherwise.

### GetBankDetailsOk

`func (o *Functionality) GetBankDetailsOk() (*string, bool)`

GetBankDetailsOk returns a tuple with the BankDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankDetails

`func (o *Functionality) SetBankDetails(v string)`

SetBankDetails sets BankDetails field to given value.

### HasBankDetails

`func (o *Functionality) HasBankDetails() bool`

HasBankDetails returns a boolean if a field has been set.

### SetBankDetailsNil

`func (o *Functionality) SetBankDetailsNil(b bool)`

 SetBankDetailsNil sets the value for BankDetails to be an explicit nil

### UnsetBankDetails
`func (o *Functionality) UnsetBankDetails()`

UnsetBankDetails ensures that no value is present for BankDetails, not even an explicit nil
### GetHeader

`func (o *Functionality) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *Functionality) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *Functionality) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *Functionality) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *Functionality) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *Functionality) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetLanguage

`func (o *Functionality) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Functionality) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Functionality) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Functionality) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *Functionality) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *Functionality) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetSkipConfirmationView

`func (o *Functionality) GetSkipConfirmationView() bool`

GetSkipConfirmationView returns the SkipConfirmationView field if non-nil, zero value otherwise.

### GetSkipConfirmationViewOk

`func (o *Functionality) GetSkipConfirmationViewOk() (*bool, bool)`

GetSkipConfirmationViewOk returns a tuple with the SkipConfirmationView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirmationView

`func (o *Functionality) SetSkipConfirmationView(v bool)`

SetSkipConfirmationView sets SkipConfirmationView field to given value.

### HasSkipConfirmationView

`func (o *Functionality) HasSkipConfirmationView() bool`

HasSkipConfirmationView returns a boolean if a field has been set.

### SetSkipConfirmationViewNil

`func (o *Functionality) SetSkipConfirmationViewNil(b bool)`

 SetSkipConfirmationViewNil sets the value for SkipConfirmationView to be an explicit nil

### UnsetSkipConfirmationView
`func (o *Functionality) UnsetSkipConfirmationView()`

UnsetSkipConfirmationView ensures that no value is present for SkipConfirmationView, not even an explicit nil
### GetRenderAccountSelectionView

`func (o *Functionality) GetRenderAccountSelectionView() bool`

GetRenderAccountSelectionView returns the RenderAccountSelectionView field if non-nil, zero value otherwise.

### GetRenderAccountSelectionViewOk

`func (o *Functionality) GetRenderAccountSelectionViewOk() (*bool, bool)`

GetRenderAccountSelectionViewOk returns a tuple with the RenderAccountSelectionView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderAccountSelectionView

`func (o *Functionality) SetRenderAccountSelectionView(v bool)`

SetRenderAccountSelectionView sets RenderAccountSelectionView field to given value.

### HasRenderAccountSelectionView

`func (o *Functionality) HasRenderAccountSelectionView() bool`

HasRenderAccountSelectionView returns a boolean if a field has been set.

### SetRenderAccountSelectionViewNil

`func (o *Functionality) SetRenderAccountSelectionViewNil(b bool)`

 SetRenderAccountSelectionViewNil sets the value for RenderAccountSelectionView to be an explicit nil

### UnsetRenderAccountSelectionView
`func (o *Functionality) UnsetRenderAccountSelectionView()`

UnsetRenderAccountSelectionView ensures that no value is present for RenderAccountSelectionView, not even an explicit nil
### GetHidePaymentSummary

`func (o *Functionality) GetHidePaymentSummary() bool`

GetHidePaymentSummary returns the HidePaymentSummary field if non-nil, zero value otherwise.

### GetHidePaymentSummaryOk

`func (o *Functionality) GetHidePaymentSummaryOk() (*bool, bool)`

GetHidePaymentSummaryOk returns a tuple with the HidePaymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePaymentSummary

`func (o *Functionality) SetHidePaymentSummary(v bool)`

SetHidePaymentSummary sets HidePaymentSummary field to given value.

### HasHidePaymentSummary

`func (o *Functionality) HasHidePaymentSummary() bool`

HasHidePaymentSummary returns a boolean if a field has been set.

### SetHidePaymentSummaryNil

`func (o *Functionality) SetHidePaymentSummaryNil(b bool)`

 SetHidePaymentSummaryNil sets the value for HidePaymentSummary to be an explicit nil

### UnsetHidePaymentSummary
`func (o *Functionality) UnsetHidePaymentSummary()`

UnsetHidePaymentSummary ensures that no value is present for HidePaymentSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


