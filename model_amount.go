/*
finAPI Web Form 2.0

The following pages give you some general information on how to use our APIs.<br/>The actual API services documentation then follows further below. You can use the menu to jump between API sections.<br/><br/>This page has a built-in HTTP(S) client, so you can test the services directly from within this page, by filling in the request parameters and/or body in the respective services, and then hitting the TRY button. Note that you need to be authorized to make a successful API call. To authorize, refer to the '<a target='_blank' href='https://docs.finapi.io/?product=access#tag--Authorization'>Authorization</a>' section of Access, or in case you already have a valid user token, just use the QUICK AUTH on the left.<br/>Please also remember that all user management functions should be looked up in <a target='_blank' href='https://docs.finapi.io/?product=access'>Access</a>.<br/><br/>You should also check out the <a target='_blank' href='https://documentation.finapi.io/webform/'>Web Form 2.0 Public Documentation</a> as well as <a target='_blank' href='https://documentation.finapi.io/access/'>Access Public Documentation</a> for more information. If you need any help with the API, contact <a href='mailto:support@finapi.io'>support@finapi.io</a>.<br/><h2 id=\"general-information\">General information</h2><h3 id=\"general-request-ids\"><strong>Request IDs</strong></h3>With any API call, you can pass a request ID via a header with name \"X-Request-Id\". The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error.<br/><br/>If you don't pass a request ID for a call, finAPI will generate a random ID internally.<br/><br/>The request ID is always returned back in the response of a service, as a header with name \"X-Request-Id\".<br/><br/>We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response(especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster.<h3 id=\"type-coercion\"><strong>Type Coercion</strong></h3>In order to ease the integration for some languages, which do not natively support high precision number representations, Web Form 2.0 API supports relax type binding for the openAPI type <code>number</code>, which is used for money amount fields. If you use one of those languages, to avoid precision errors that can appear from <code>float</code> values, you can pass the amount as a <code>string</code>.<h3 id=\"general-faq\"><strong>FAQ</strong></h3><strong>Is there a finAPI SDK?</strong><br/>Currently we do not offer a native SDK, but there is the option to generate an SDKfor almost any target language via OpenAPI. Use the 'Download SDK' button on this page for SDK generation.<br/><br/><strong>Why do I need to keep authorizing when calling services on this page?</strong><br/>This page is a \"one-page-app\". Reloading the page resets the OAuth authorization context. There is generally no need to reload the page, so just don't do it and your authorization will persist.

API version: 2.430.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webform2

import (
	"encoding/json"
)

// Amount Amount details
type Amount struct {
	// The amount of the payment. Must be a positive decimal number with at most two decimal places.
	Value float32 `json:"value"`
	Currency Currency `json:"currency"`
}

// NewAmount instantiates a new Amount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmount(value float32, currency Currency) *Amount {
	this := Amount{}
	this.Value = value
	this.Currency = currency
	return &this
}

// NewAmountWithDefaults instantiates a new Amount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountWithDefaults() *Amount {
	this := Amount{}
	return &this
}

// GetValue returns the Value field value
func (o *Amount) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Amount) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Amount) SetValue(v float32) {
	o.Value = v
}

// GetCurrency returns the Currency field value
func (o *Amount) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Amount) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Amount) SetCurrency(v Currency) {
	o.Currency = v
}

func (o Amount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableAmount struct {
	value *Amount
	isSet bool
}

func (v NullableAmount) Get() *Amount {
	return v.value
}

func (v *NullableAmount) Set(val *Amount) {
	v.value = val
	v.isSet = true
}

func (v NullableAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmount(val *Amount) *NullableAmount {
	return &NullableAmount{value: val, isSet: true}
}

func (v NullableAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


