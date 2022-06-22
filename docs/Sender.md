# Sender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **string** | A normalized (without spaces) IBAN of the sender&#39;s account. If the value is provided, the web form will be prefilled with that value. Otherwise, it will ask the end user to provide a valid IBAN. | 

## Methods

### NewSender

`func NewSender(iban string, ) *Sender`

NewSender instantiates a new Sender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSenderWithDefaults

`func NewSenderWithDefaults() *Sender`

NewSenderWithDefaults instantiates a new Sender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *Sender) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Sender) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Sender) SetIban(v string)`

SetIban sets Iban field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


