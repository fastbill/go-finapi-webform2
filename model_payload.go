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

// Payload Payload of the web form
type Payload struct {
	// Identifier of the bank connection in the Access API. Initialized as soon as the Web Form reaches a final status and a bank connection exists in Access. Use this ID to gather Bank Connection data from Access endpoints like, \"<a target='_blank' href='https://docs.finapi.io/?product=access#get-/api/v1/bankConnections/-id-'>Get a bank connection</a>\" or \"<a target='_blank' href='https://docs.finapi.io/?product=access#get-/api/v1/accounts'>Get and search all accounts</a>\".<br/>This field is mutually exclusive with <code>paymentId</code> and <code>standingOrderId</code>.
	BankConnectionId NullableInt64 `json:"bankConnectionId,omitempty"`
	// Identifier of the payment in the Access API. Initialized as soon as the Web Form reaches a final status and a payment exists in Access. Use this ID to get Payment initialization data from the Access endpoint, \"<a target='_blank' href='https://docs.finapi.io/?product=access#get-/api/v1/payments'>Get payments</a>\".<br/>This field is mutually exclusive with <code>bankConnectionId</code> and <code>standingOrderId</code>.
	PaymentId NullableInt64 `json:"paymentId,omitempty"`
	// Identifier of the standing order in the Access API. Initialized as soon as the Web Form reaches a final status and a standing order exists in Access. Use this ID to get Standing Order initialization data from the Access endpoint, \"<a target='_blank' href='https://docs.finapi.io/?product=access#get-/api/v1/standingOrders'>Get standing orders</a>\".<br/>This field is mutually exclusive with <code>bankConnectionId</code> and <code>paymentId</code>.
	StandingOrderId NullableInt64 `json:"standingOrderId,omitempty"`
	// Reason of the web form failure.<br/><strong>NOTE:</strong> This enum can be extended in the future as new cases arise!<br/><br/>Codes can be interpreted as follows:<br/>&bull; <code>ENTITY_EXISTS</code> - Access API rejected the import because of detected bank connection duplication;<br/>&bull; <code>BANK_SERVER_REJECTION</code> - the flow has been terminated on the bank side, e.g., in case of incorrect credentials;<br/>&bull; <code>INTERRUPTED</code> - web form has been reloaded or re-opened on a step where it's not supported;<br/>&bull; <code>INVALID_TOKEN</code> - the given access token expired or became invalid during the flow; <br/>&bull; <code>MANDATOR_MISCONFIGURATION</code> - the mandator is not properly configured on the Access side, e.g., it is still configured to use the old web form, or the scraper interface has been selected by the Web Form 2.0 but using it is not allowed on the mandator level. Please contact our support team (<a href='mailto:support@finapi.io'>support@finapi.io</a>) for troubleshooting;<br/>&bull; <code>NO_ACCOUNTS_FOR_TYPE_LIST</code> - in the end of the web form flow there were no accounts of type requested in the API call; <br/>&bull; <code>UNDETERMINED_BANK</code> - the given search criteria resulted in either zero or multiple bank entries;<br/>&bull; <code>UNEXPECTED_ACCESS_RESPONSE</code> - an unexpected response has been received from the Access API - similarly to the <code>INTERNAL_ERROR</code> code, please forward all details to our Customer Support team; <br/>&bull; <code>UNSUPPORTED_FEATURE</code> - Access API rejected the request because the requested feature is not supported, e.g., a payment with the execution date in the future was requested for a bank that does not support it;<br/>&bull; <code>UNSUPPORTED_ORDER</code> - Access API rejected the payment request because the associated account does not have the required capabilities;<br/>&bull; <code>INTERNAL_ERROR</code> - the reason of the failure can not be identified - please forward all the details to our Customer Support team in order to get more info and also help us to eliminate the issue.
	ErrorCode NullableString `json:"errorCode,omitempty"`
}

// NewPayload instantiates a new Payload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayload() *Payload {
	this := Payload{}
	return &this
}

// NewPayloadWithDefaults instantiates a new Payload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadWithDefaults() *Payload {
	this := Payload{}
	return &this
}

// GetBankConnectionId returns the BankConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payload) GetBankConnectionId() int64 {
	if o == nil || o.BankConnectionId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BankConnectionId.Get()
}

// GetBankConnectionIdOk returns a tuple with the BankConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payload) GetBankConnectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankConnectionId.Get(), o.BankConnectionId.IsSet()
}

// HasBankConnectionId returns a boolean if a field has been set.
func (o *Payload) HasBankConnectionId() bool {
	if o != nil && o.BankConnectionId.IsSet() {
		return true
	}

	return false
}

// SetBankConnectionId gets a reference to the given NullableInt64 and assigns it to the BankConnectionId field.
func (o *Payload) SetBankConnectionId(v int64) {
	o.BankConnectionId.Set(&v)
}
// SetBankConnectionIdNil sets the value for BankConnectionId to be an explicit nil
func (o *Payload) SetBankConnectionIdNil() {
	o.BankConnectionId.Set(nil)
}

// UnsetBankConnectionId ensures that no value is present for BankConnectionId, not even an explicit nil
func (o *Payload) UnsetBankConnectionId() {
	o.BankConnectionId.Unset()
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payload) GetPaymentId() int64 {
	if o == nil || o.PaymentId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payload) GetPaymentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *Payload) HasPaymentId() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableInt64 and assigns it to the PaymentId field.
func (o *Payload) SetPaymentId(v int64) {
	o.PaymentId.Set(&v)
}
// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *Payload) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *Payload) UnsetPaymentId() {
	o.PaymentId.Unset()
}

// GetStandingOrderId returns the StandingOrderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payload) GetStandingOrderId() int64 {
	if o == nil || o.StandingOrderId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StandingOrderId.Get()
}

// GetStandingOrderIdOk returns a tuple with the StandingOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payload) GetStandingOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StandingOrderId.Get(), o.StandingOrderId.IsSet()
}

// HasStandingOrderId returns a boolean if a field has been set.
func (o *Payload) HasStandingOrderId() bool {
	if o != nil && o.StandingOrderId.IsSet() {
		return true
	}

	return false
}

// SetStandingOrderId gets a reference to the given NullableInt64 and assigns it to the StandingOrderId field.
func (o *Payload) SetStandingOrderId(v int64) {
	o.StandingOrderId.Set(&v)
}
// SetStandingOrderIdNil sets the value for StandingOrderId to be an explicit nil
func (o *Payload) SetStandingOrderIdNil() {
	o.StandingOrderId.Set(nil)
}

// UnsetStandingOrderId ensures that no value is present for StandingOrderId, not even an explicit nil
func (o *Payload) UnsetStandingOrderId() {
	o.StandingOrderId.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payload) GetErrorCode() string {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payload) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *Payload) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *Payload) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *Payload) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *Payload) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

func (o Payload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BankConnectionId.IsSet() {
		toSerialize["bankConnectionId"] = o.BankConnectionId.Get()
	}
	if o.PaymentId.IsSet() {
		toSerialize["paymentId"] = o.PaymentId.Get()
	}
	if o.StandingOrderId.IsSet() {
		toSerialize["standingOrderId"] = o.StandingOrderId.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePayload struct {
	value *Payload
	isSet bool
}

func (v NullablePayload) Get() *Payload {
	return v.value
}

func (v *NullablePayload) Set(val *Payload) {
	v.value = val
	v.isSet = true
}

func (v NullablePayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayload(val *Payload) *NullablePayload {
	return &NullablePayload{value: val, isSet: true}
}

func (v NullablePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


