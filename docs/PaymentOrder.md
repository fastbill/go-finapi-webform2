# PaymentOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**Purpose** | Pointer to **NullableString** | The purpose of the transfer transaction | [optional] 
**SepaPurposeCode** | Pointer to **NullableString** | SEPA purpose code, according to ISO 20022, external codes set.&lt;br/&gt;Note that the SEPA purpose code may be ignored by some banks. | [optional] 
**EndToEndId** | Pointer to **NullableString** | End-To-End ID for the transfer transaction | [optional] 
**Recipient** | [**PaymentRecipient**](PaymentRecipient.md) |  | 

## Methods

### NewPaymentOrder

`func NewPaymentOrder(amount Amount, recipient PaymentRecipient, ) *PaymentOrder`

NewPaymentOrder instantiates a new PaymentOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOrderWithDefaults

`func NewPaymentOrderWithDefaults() *PaymentOrder`

NewPaymentOrderWithDefaults instantiates a new PaymentOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentOrder) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentOrder) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentOrder) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *PaymentOrder) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *PaymentOrder) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *PaymentOrder) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *PaymentOrder) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *PaymentOrder) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *PaymentOrder) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetSepaPurposeCode

`func (o *PaymentOrder) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *PaymentOrder) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *PaymentOrder) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *PaymentOrder) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.

### SetSepaPurposeCodeNil

`func (o *PaymentOrder) SetSepaPurposeCodeNil(b bool)`

 SetSepaPurposeCodeNil sets the value for SepaPurposeCode to be an explicit nil

### UnsetSepaPurposeCode
`func (o *PaymentOrder) UnsetSepaPurposeCode()`

UnsetSepaPurposeCode ensures that no value is present for SepaPurposeCode, not even an explicit nil
### GetEndToEndId

`func (o *PaymentOrder) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *PaymentOrder) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *PaymentOrder) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *PaymentOrder) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.

### SetEndToEndIdNil

`func (o *PaymentOrder) SetEndToEndIdNil(b bool)`

 SetEndToEndIdNil sets the value for EndToEndId to be an explicit nil

### UnsetEndToEndId
`func (o *PaymentOrder) UnsetEndToEndId()`

UnsetEndToEndId ensures that no value is present for EndToEndId, not even an explicit nil
### GetRecipient

`func (o *PaymentOrder) GetRecipient() PaymentRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *PaymentOrder) GetRecipientOk() (*PaymentRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *PaymentOrder) SetRecipient(v PaymentRecipient)`

SetRecipient sets Recipient field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


