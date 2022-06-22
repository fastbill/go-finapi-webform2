# DirectDebitOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**Purpose** | Pointer to **NullableString** | The purpose of the transfer transaction | [optional] 
**SepaPurposeCode** | Pointer to **NullableString** | SEPA purpose code, according to ISO 20022, external codes set.&lt;br/&gt;Note that the SEPA purpose code may be ignored by some banks. | [optional] 
**EndToEndId** | Pointer to **NullableString** | End-To-End ID for the transfer transaction | [optional] 
**Payer** | [**PaymentPayer**](PaymentPayer.md) |  | 
**MandateId** | **string** | Mandate ID that this direct debit order is based on. | 
**MandateDate** | **string** | Date of the mandate that this direct debit order is based on, in the format &#39;YYYY-MM-DD&#39; | 
**CreditorId** | **string** | Creditor ID of the source account&#39;s holder | 

## Methods

### NewDirectDebitOrder

`func NewDirectDebitOrder(amount Amount, payer PaymentPayer, mandateId string, mandateDate string, creditorId string, ) *DirectDebitOrder`

NewDirectDebitOrder instantiates a new DirectDebitOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitOrderWithDefaults

`func NewDirectDebitOrderWithDefaults() *DirectDebitOrder`

NewDirectDebitOrderWithDefaults instantiates a new DirectDebitOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DirectDebitOrder) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DirectDebitOrder) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DirectDebitOrder) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *DirectDebitOrder) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DirectDebitOrder) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DirectDebitOrder) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *DirectDebitOrder) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *DirectDebitOrder) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *DirectDebitOrder) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetSepaPurposeCode

`func (o *DirectDebitOrder) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *DirectDebitOrder) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *DirectDebitOrder) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *DirectDebitOrder) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.

### SetSepaPurposeCodeNil

`func (o *DirectDebitOrder) SetSepaPurposeCodeNil(b bool)`

 SetSepaPurposeCodeNil sets the value for SepaPurposeCode to be an explicit nil

### UnsetSepaPurposeCode
`func (o *DirectDebitOrder) UnsetSepaPurposeCode()`

UnsetSepaPurposeCode ensures that no value is present for SepaPurposeCode, not even an explicit nil
### GetEndToEndId

`func (o *DirectDebitOrder) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *DirectDebitOrder) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *DirectDebitOrder) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *DirectDebitOrder) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.

### SetEndToEndIdNil

`func (o *DirectDebitOrder) SetEndToEndIdNil(b bool)`

 SetEndToEndIdNil sets the value for EndToEndId to be an explicit nil

### UnsetEndToEndId
`func (o *DirectDebitOrder) UnsetEndToEndId()`

UnsetEndToEndId ensures that no value is present for EndToEndId, not even an explicit nil
### GetPayer

`func (o *DirectDebitOrder) GetPayer() PaymentPayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *DirectDebitOrder) GetPayerOk() (*PaymentPayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *DirectDebitOrder) SetPayer(v PaymentPayer)`

SetPayer sets Payer field to given value.


### GetMandateId

`func (o *DirectDebitOrder) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *DirectDebitOrder) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *DirectDebitOrder) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### GetMandateDate

`func (o *DirectDebitOrder) GetMandateDate() string`

GetMandateDate returns the MandateDate field if non-nil, zero value otherwise.

### GetMandateDateOk

`func (o *DirectDebitOrder) GetMandateDateOk() (*string, bool)`

GetMandateDateOk returns a tuple with the MandateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateDate

`func (o *DirectDebitOrder) SetMandateDate(v string)`

SetMandateDate sets MandateDate field to given value.


### GetCreditorId

`func (o *DirectDebitOrder) GetCreditorId() string`

GetCreditorId returns the CreditorId field if non-nil, zero value otherwise.

### GetCreditorIdOk

`func (o *DirectDebitOrder) GetCreditorIdOk() (*string, bool)`

GetCreditorIdOk returns a tuple with the CreditorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorId

`func (o *DirectDebitOrder) SetCreditorId(v string)`

SetCreditorId sets CreditorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


