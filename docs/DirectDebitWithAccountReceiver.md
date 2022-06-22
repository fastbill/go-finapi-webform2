# DirectDebitWithAccountReceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | The identifier of an account (which is initiating the direct debit) in the Access API | 

## Methods

### NewDirectDebitWithAccountReceiver

`func NewDirectDebitWithAccountReceiver(accountId int64, ) *DirectDebitWithAccountReceiver`

NewDirectDebitWithAccountReceiver instantiates a new DirectDebitWithAccountReceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitWithAccountReceiverWithDefaults

`func NewDirectDebitWithAccountReceiverWithDefaults() *DirectDebitWithAccountReceiver`

NewDirectDebitWithAccountReceiverWithDefaults instantiates a new DirectDebitWithAccountReceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DirectDebitWithAccountReceiver) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DirectDebitWithAccountReceiver) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DirectDebitWithAccountReceiver) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


