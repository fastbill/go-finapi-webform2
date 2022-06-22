# PaymentPayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **string** | A normalized (without spaces) IBAN of the counterparty&#39;s account | 
**Name** | **string** | Name of the counterparty.&lt;br/&gt;Note that neither finAPI nor the involved bank are guaranteed to validate the counterparty name. Even if the name does not depict the actual registered account holder of the target account, the order might still be successful. | 
**Bic** | Pointer to **NullableString** | BIC of the counterparty&#39;s account | [optional] 
**Address** | Pointer to **NullableString** | The postal address of the counterparty. This should be defined for direct debits created for debtors outside of the european union. | [optional] 
**Country** | Pointer to [**NullableCountryCode**](CountryCode.md) |  | [optional] 

## Methods

### NewPaymentPayer

`func NewPaymentPayer(iban string, name string, ) *PaymentPayer`

NewPaymentPayer instantiates a new PaymentPayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayerWithDefaults

`func NewPaymentPayerWithDefaults() *PaymentPayer`

NewPaymentPayerWithDefaults instantiates a new PaymentPayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *PaymentPayer) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentPayer) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentPayer) SetIban(v string)`

SetIban sets Iban field to given value.


### GetName

`func (o *PaymentPayer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentPayer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentPayer) SetName(v string)`

SetName sets Name field to given value.


### GetBic

`func (o *PaymentPayer) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *PaymentPayer) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *PaymentPayer) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *PaymentPayer) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *PaymentPayer) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *PaymentPayer) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil
### GetAddress

`func (o *PaymentPayer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentPayer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentPayer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaymentPayer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *PaymentPayer) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *PaymentPayer) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCountry

`func (o *PaymentPayer) GetCountry() CountryCode`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentPayer) GetCountryOk() (*CountryCode, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentPayer) SetCountry(v CountryCode)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentPayer) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PaymentPayer) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentPayer) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


