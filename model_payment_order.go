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

// PaymentOrder Payment order
type PaymentOrder struct {
	Amount Amount `json:"amount"`
	// The purpose of the transfer transaction
	Purpose NullableString `json:"purpose,omitempty"`
	// SEPA purpose code, according to ISO 20022, external codes set.<br/>Note that the SEPA purpose code may be ignored by some banks.
	SepaPurposeCode NullableString `json:"sepaPurposeCode,omitempty"`
	// End-To-End ID for the transfer transaction
	EndToEndId NullableString `json:"endToEndId,omitempty"`
	Recipient PaymentRecipient `json:"recipient"`
}

// NewPaymentOrder instantiates a new PaymentOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentOrder(amount Amount, recipient PaymentRecipient) *PaymentOrder {
	this := PaymentOrder{}
	this.Amount = amount
	this.Recipient = recipient
	return &this
}

// NewPaymentOrderWithDefaults instantiates a new PaymentOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentOrderWithDefaults() *PaymentOrder {
	this := PaymentOrder{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentOrder) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentOrder) SetAmount(v Amount) {
	o.Amount = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentOrder) GetPurpose() string {
	if o == nil || o.Purpose.Get() == nil {
		var ret string
		return ret
	}
	return *o.Purpose.Get()
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentOrder) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Purpose.Get(), o.Purpose.IsSet()
}

// HasPurpose returns a boolean if a field has been set.
func (o *PaymentOrder) HasPurpose() bool {
	if o != nil && o.Purpose.IsSet() {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given NullableString and assigns it to the Purpose field.
func (o *PaymentOrder) SetPurpose(v string) {
	o.Purpose.Set(&v)
}
// SetPurposeNil sets the value for Purpose to be an explicit nil
func (o *PaymentOrder) SetPurposeNil() {
	o.Purpose.Set(nil)
}

// UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
func (o *PaymentOrder) UnsetPurpose() {
	o.Purpose.Unset()
}

// GetSepaPurposeCode returns the SepaPurposeCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentOrder) GetSepaPurposeCode() string {
	if o == nil || o.SepaPurposeCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.SepaPurposeCode.Get()
}

// GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentOrder) GetSepaPurposeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SepaPurposeCode.Get(), o.SepaPurposeCode.IsSet()
}

// HasSepaPurposeCode returns a boolean if a field has been set.
func (o *PaymentOrder) HasSepaPurposeCode() bool {
	if o != nil && o.SepaPurposeCode.IsSet() {
		return true
	}

	return false
}

// SetSepaPurposeCode gets a reference to the given NullableString and assigns it to the SepaPurposeCode field.
func (o *PaymentOrder) SetSepaPurposeCode(v string) {
	o.SepaPurposeCode.Set(&v)
}
// SetSepaPurposeCodeNil sets the value for SepaPurposeCode to be an explicit nil
func (o *PaymentOrder) SetSepaPurposeCodeNil() {
	o.SepaPurposeCode.Set(nil)
}

// UnsetSepaPurposeCode ensures that no value is present for SepaPurposeCode, not even an explicit nil
func (o *PaymentOrder) UnsetSepaPurposeCode() {
	o.SepaPurposeCode.Unset()
}

// GetEndToEndId returns the EndToEndId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentOrder) GetEndToEndId() string {
	if o == nil || o.EndToEndId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndToEndId.Get()
}

// GetEndToEndIdOk returns a tuple with the EndToEndId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentOrder) GetEndToEndIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndToEndId.Get(), o.EndToEndId.IsSet()
}

// HasEndToEndId returns a boolean if a field has been set.
func (o *PaymentOrder) HasEndToEndId() bool {
	if o != nil && o.EndToEndId.IsSet() {
		return true
	}

	return false
}

// SetEndToEndId gets a reference to the given NullableString and assigns it to the EndToEndId field.
func (o *PaymentOrder) SetEndToEndId(v string) {
	o.EndToEndId.Set(&v)
}
// SetEndToEndIdNil sets the value for EndToEndId to be an explicit nil
func (o *PaymentOrder) SetEndToEndIdNil() {
	o.EndToEndId.Set(nil)
}

// UnsetEndToEndId ensures that no value is present for EndToEndId, not even an explicit nil
func (o *PaymentOrder) UnsetEndToEndId() {
	o.EndToEndId.Unset()
}

// GetRecipient returns the Recipient field value
func (o *PaymentOrder) GetRecipient() PaymentRecipient {
	if o == nil {
		var ret PaymentRecipient
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetRecipientOk() (*PaymentRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *PaymentOrder) SetRecipient(v PaymentRecipient) {
	o.Recipient = v
}

func (o PaymentOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Purpose.IsSet() {
		toSerialize["purpose"] = o.Purpose.Get()
	}
	if o.SepaPurposeCode.IsSet() {
		toSerialize["sepaPurposeCode"] = o.SepaPurposeCode.Get()
	}
	if o.EndToEndId.IsSet() {
		toSerialize["endToEndId"] = o.EndToEndId.Get()
	}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentOrder struct {
	value *PaymentOrder
	isSet bool
}

func (v NullablePaymentOrder) Get() *PaymentOrder {
	return v.value
}

func (v *NullablePaymentOrder) Set(val *PaymentOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentOrder(val *PaymentOrder) *NullablePaymentOrder {
	return &NullablePaymentOrder{value: val, isSet: true}
}

func (v NullablePaymentOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


