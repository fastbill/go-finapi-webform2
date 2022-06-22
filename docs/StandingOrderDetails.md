# StandingOrderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**Purpose** | Pointer to **NullableString** | The purpose of the transfer transaction | [optional] 
**SepaPurposeCode** | Pointer to **NullableString** | SEPA purpose code, according to ISO 20022, external codes set.&lt;br/&gt;Note that the SEPA purpose code may be ignored by some banks. | [optional] 
**EndToEndId** | Pointer to **NullableString** | End-To-End ID for the transfer transaction | [optional] 
**Recipient** | [**Recipient**](Recipient.md) |  | 
**Sender** | Pointer to [**NullableSender**](Sender.md) |  | [optional] 
**StartDate** | **string** | Date when the 1st of the standing orders should be executed, in the format &lt;code&gt;YYYY-MM-DD&lt;/code&gt;. The date may not be in the past. | 
**EndDate** | Pointer to **NullableString** | Date by when the last standing order in the request should be executed, in the format &lt;code&gt;YYYY-MM-DD&lt;/code&gt;. If is not provided, it is an infinite standing order. This date must be after the start date. | [optional] 
**Frequency** | **string** | The frequency of the recurring payment resulting from the standing order. | 
**ProfileId** | Pointer to **NullableString** | The profile to be applied to the web form.&lt;br/&gt;This will overwrite the default profile, if such a profile exists. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The URL where the end-user will be redirected to after completing the bank login and (possibly) the SCA on the bank&#39;s website. Must always be provided by mandators with &lt;code&gt;FULLY_LICENSED&lt;/code&gt; or &lt;code&gt;PISP&lt;/code&gt; license type, and may not be provided by mandators with other license types. Find more info in the &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://documentation.finapi.io/webform/Licensed-customers-using-the-Web-Form.2832302195.html&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | [optional] 
**Callbacks** | Pointer to [**NullableCallbacks**](Callbacks.md) |  | [optional] 

## Methods

### NewStandingOrderDetails

`func NewStandingOrderDetails(amount Amount, recipient Recipient, startDate string, frequency string, ) *StandingOrderDetails`

NewStandingOrderDetails instantiates a new StandingOrderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandingOrderDetailsWithDefaults

`func NewStandingOrderDetailsWithDefaults() *StandingOrderDetails`

NewStandingOrderDetailsWithDefaults instantiates a new StandingOrderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *StandingOrderDetails) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StandingOrderDetails) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StandingOrderDetails) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *StandingOrderDetails) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *StandingOrderDetails) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *StandingOrderDetails) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *StandingOrderDetails) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *StandingOrderDetails) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *StandingOrderDetails) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetSepaPurposeCode

`func (o *StandingOrderDetails) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *StandingOrderDetails) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *StandingOrderDetails) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *StandingOrderDetails) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.

### SetSepaPurposeCodeNil

`func (o *StandingOrderDetails) SetSepaPurposeCodeNil(b bool)`

 SetSepaPurposeCodeNil sets the value for SepaPurposeCode to be an explicit nil

### UnsetSepaPurposeCode
`func (o *StandingOrderDetails) UnsetSepaPurposeCode()`

UnsetSepaPurposeCode ensures that no value is present for SepaPurposeCode, not even an explicit nil
### GetEndToEndId

`func (o *StandingOrderDetails) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *StandingOrderDetails) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *StandingOrderDetails) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *StandingOrderDetails) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.

### SetEndToEndIdNil

`func (o *StandingOrderDetails) SetEndToEndIdNil(b bool)`

 SetEndToEndIdNil sets the value for EndToEndId to be an explicit nil

### UnsetEndToEndId
`func (o *StandingOrderDetails) UnsetEndToEndId()`

UnsetEndToEndId ensures that no value is present for EndToEndId, not even an explicit nil
### GetRecipient

`func (o *StandingOrderDetails) GetRecipient() Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *StandingOrderDetails) GetRecipientOk() (*Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *StandingOrderDetails) SetRecipient(v Recipient)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *StandingOrderDetails) GetSender() Sender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *StandingOrderDetails) GetSenderOk() (*Sender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *StandingOrderDetails) SetSender(v Sender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *StandingOrderDetails) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *StandingOrderDetails) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *StandingOrderDetails) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetStartDate

`func (o *StandingOrderDetails) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StandingOrderDetails) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StandingOrderDetails) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *StandingOrderDetails) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StandingOrderDetails) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StandingOrderDetails) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *StandingOrderDetails) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *StandingOrderDetails) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *StandingOrderDetails) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetFrequency

`func (o *StandingOrderDetails) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *StandingOrderDetails) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *StandingOrderDetails) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetProfileId

`func (o *StandingOrderDetails) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *StandingOrderDetails) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *StandingOrderDetails) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *StandingOrderDetails) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *StandingOrderDetails) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *StandingOrderDetails) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetRedirectUrl

`func (o *StandingOrderDetails) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *StandingOrderDetails) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *StandingOrderDetails) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *StandingOrderDetails) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *StandingOrderDetails) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *StandingOrderDetails) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetCallbacks

`func (o *StandingOrderDetails) GetCallbacks() Callbacks`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *StandingOrderDetails) GetCallbacksOk() (*Callbacks, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *StandingOrderDetails) SetCallbacks(v Callbacks)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *StandingOrderDetails) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### SetCallbacksNil

`func (o *StandingOrderDetails) SetCallbacksNil(b bool)`

 SetCallbacksNil sets the value for Callbacks to be an explicit nil

### UnsetCallbacks
`func (o *StandingOrderDetails) UnsetCallbacks()`

UnsetCallbacks ensures that no value is present for Callbacks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


