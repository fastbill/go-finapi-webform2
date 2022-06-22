/*
finAPI Web Form 2.0

The following pages give you some general information on how to use our APIs.<br/>The actual API services documentation then follows further below. You can use the menu to jump between API sections.<br/><br/>This page has a built-in HTTP(S) client, so you can test the services directly from within this page, by filling in the request parameters and/or body in the respective services, and then hitting the TRY button. Note that you need to be authorized to make a successful API call. To authorize, refer to the '<a target='_blank' href='https://docs.finapi.io/?product=access#tag--Authorization'>Authorization</a>' section of Access, or in case you already have a valid user token, just use the QUICK AUTH on the left.<br/>Please also remember that all user management functions should be looked up in <a target='_blank' href='https://docs.finapi.io/?product=access'>Access</a>.<br/><br/>You should also check out the <a target='_blank' href='https://documentation.finapi.io/webform/'>Web Form 2.0 Public Documentation</a> as well as <a target='_blank' href='https://documentation.finapi.io/access/'>Access Public Documentation</a> for more information. If you need any help with the API, contact <a href='mailto:support@finapi.io'>support@finapi.io</a>.<br/><h2 id=\"general-information\">General information</h2><h3 id=\"general-request-ids\"><strong>Request IDs</strong></h3>With any API call, you can pass a request ID via a header with name \"X-Request-Id\". The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error.<br/><br/>If you don't pass a request ID for a call, finAPI will generate a random ID internally.<br/><br/>The request ID is always returned back in the response of a service, as a header with name \"X-Request-Id\".<br/><br/>We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response(especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster.<h3 id=\"type-coercion\"><strong>Type Coercion</strong></h3>In order to ease the integration for some languages, which do not natively support high precision number representations, Web Form 2.0 API supports relax type binding for the openAPI type <code>number</code>, which is used for money amount fields. If you use one of those languages, to avoid precision errors that can appear from <code>float</code> values, you can pass the amount as a <code>string</code>.<h3 id=\"general-faq\"><strong>FAQ</strong></h3><strong>Is there a finAPI SDK?</strong><br/>Currently we do not offer a native SDK, but there is the option to generate an SDKfor almost any target language via OpenAPI. Use the 'Download SDK' button on this page for SDK generation.<br/><br/><strong>Why do I need to keep authorizing when calling services on this page?</strong><br/>This page is a \"one-page-app\". Reloading the page resets the OAuth authorization context. There is generally no need to reload the page, so just don't do it and your authorization will persist.

API version: 2.430.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webform2

import (
	"encoding/json"
	"time"
)

// TranslationSet struct for TranslationSet
type TranslationSet struct {
	// Globally unique translation set's identifier
	Id string `json:"id"`
	// The timestamp when the translation set was created in the format <code>yyyy-MM-dd'T'HH:mm:ss.SSSZ</code>.
	CreatedAt time.Time `json:"createdAt"`
	Cs NullableTranslation `json:"cs,omitempty"`
	De NullableTranslation `json:"de,omitempty"`
	En NullableTranslation `json:"en,omitempty"`
	Sk NullableTranslation `json:"sk,omitempty"`
}

// NewTranslationSet instantiates a new TranslationSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationSet(id string, createdAt time.Time) *TranslationSet {
	this := TranslationSet{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewTranslationSetWithDefaults instantiates a new TranslationSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationSetWithDefaults() *TranslationSet {
	this := TranslationSet{}
	return &this
}

// GetId returns the Id field value
func (o *TranslationSet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TranslationSet) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TranslationSet) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TranslationSet) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TranslationSet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TranslationSet) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCs returns the Cs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranslationSet) GetCs() Translation {
	if o == nil || o.Cs.Get() == nil {
		var ret Translation
		return ret
	}
	return *o.Cs.Get()
}

// GetCsOk returns a tuple with the Cs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranslationSet) GetCsOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cs.Get(), o.Cs.IsSet()
}

// HasCs returns a boolean if a field has been set.
func (o *TranslationSet) HasCs() bool {
	if o != nil && o.Cs.IsSet() {
		return true
	}

	return false
}

// SetCs gets a reference to the given NullableTranslation and assigns it to the Cs field.
func (o *TranslationSet) SetCs(v Translation) {
	o.Cs.Set(&v)
}
// SetCsNil sets the value for Cs to be an explicit nil
func (o *TranslationSet) SetCsNil() {
	o.Cs.Set(nil)
}

// UnsetCs ensures that no value is present for Cs, not even an explicit nil
func (o *TranslationSet) UnsetCs() {
	o.Cs.Unset()
}

// GetDe returns the De field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranslationSet) GetDe() Translation {
	if o == nil || o.De.Get() == nil {
		var ret Translation
		return ret
	}
	return *o.De.Get()
}

// GetDeOk returns a tuple with the De field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranslationSet) GetDeOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return o.De.Get(), o.De.IsSet()
}

// HasDe returns a boolean if a field has been set.
func (o *TranslationSet) HasDe() bool {
	if o != nil && o.De.IsSet() {
		return true
	}

	return false
}

// SetDe gets a reference to the given NullableTranslation and assigns it to the De field.
func (o *TranslationSet) SetDe(v Translation) {
	o.De.Set(&v)
}
// SetDeNil sets the value for De to be an explicit nil
func (o *TranslationSet) SetDeNil() {
	o.De.Set(nil)
}

// UnsetDe ensures that no value is present for De, not even an explicit nil
func (o *TranslationSet) UnsetDe() {
	o.De.Unset()
}

// GetEn returns the En field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranslationSet) GetEn() Translation {
	if o == nil || o.En.Get() == nil {
		var ret Translation
		return ret
	}
	return *o.En.Get()
}

// GetEnOk returns a tuple with the En field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranslationSet) GetEnOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return o.En.Get(), o.En.IsSet()
}

// HasEn returns a boolean if a field has been set.
func (o *TranslationSet) HasEn() bool {
	if o != nil && o.En.IsSet() {
		return true
	}

	return false
}

// SetEn gets a reference to the given NullableTranslation and assigns it to the En field.
func (o *TranslationSet) SetEn(v Translation) {
	o.En.Set(&v)
}
// SetEnNil sets the value for En to be an explicit nil
func (o *TranslationSet) SetEnNil() {
	o.En.Set(nil)
}

// UnsetEn ensures that no value is present for En, not even an explicit nil
func (o *TranslationSet) UnsetEn() {
	o.En.Unset()
}

// GetSk returns the Sk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranslationSet) GetSk() Translation {
	if o == nil || o.Sk.Get() == nil {
		var ret Translation
		return ret
	}
	return *o.Sk.Get()
}

// GetSkOk returns a tuple with the Sk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranslationSet) GetSkOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sk.Get(), o.Sk.IsSet()
}

// HasSk returns a boolean if a field has been set.
func (o *TranslationSet) HasSk() bool {
	if o != nil && o.Sk.IsSet() {
		return true
	}

	return false
}

// SetSk gets a reference to the given NullableTranslation and assigns it to the Sk field.
func (o *TranslationSet) SetSk(v Translation) {
	o.Sk.Set(&v)
}
// SetSkNil sets the value for Sk to be an explicit nil
func (o *TranslationSet) SetSkNil() {
	o.Sk.Set(nil)
}

// UnsetSk ensures that no value is present for Sk, not even an explicit nil
func (o *TranslationSet) UnsetSk() {
	o.Sk.Unset()
}

func (o TranslationSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Cs.IsSet() {
		toSerialize["cs"] = o.Cs.Get()
	}
	if o.De.IsSet() {
		toSerialize["de"] = o.De.Get()
	}
	if o.En.IsSet() {
		toSerialize["en"] = o.En.Get()
	}
	if o.Sk.IsSet() {
		toSerialize["sk"] = o.Sk.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTranslationSet struct {
	value *TranslationSet
	isSet bool
}

func (v NullableTranslationSet) Get() *TranslationSet {
	return v.value
}

func (v *NullableTranslationSet) Set(val *TranslationSet) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationSet) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationSet(val *TranslationSet) *NullableTranslationSet {
	return &NullableTranslationSet{value: val, isSet: true}
}

func (v NullableTranslationSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


