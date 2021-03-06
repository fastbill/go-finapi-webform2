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

// EditProfileDetails struct for EditProfileDetails
type EditProfileDetails struct {
	// Label to profile.<br/>The label must be unique.
	Label NullableString `json:"label,omitempty"`
	// Whether the profile will be used by default for all web forms.<br/>We urge you to set one profile as default. This way, if you do not pass the profile in the API call, we will use the default profile from you for the web forms.<br/><br/>There can only be one default profile at one time.
	Default NullableBool `json:"default,omitempty"`
	Brand NullableBrand `json:"brand,omitempty"`
	Functionality NullableFunctionality `json:"functionality,omitempty"`
	Aspect NullableAspect `json:"aspect,omitempty"`
}

// NewEditProfileDetails instantiates a new EditProfileDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProfileDetails() *EditProfileDetails {
	this := EditProfileDetails{}
	return &this
}

// NewEditProfileDetailsWithDefaults instantiates a new EditProfileDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProfileDetailsWithDefaults() *EditProfileDetails {
	this := EditProfileDetails{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditProfileDetails) GetLabel() string {
	if o == nil || o.Label.Get() == nil {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditProfileDetails) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *EditProfileDetails) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *EditProfileDetails) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *EditProfileDetails) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *EditProfileDetails) UnsetLabel() {
	o.Label.Unset()
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditProfileDetails) GetDefault() bool {
	if o == nil || o.Default.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditProfileDetails) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *EditProfileDetails) HasDefault() bool {
	if o != nil && o.Default.IsSet() {
		return true
	}

	return false
}

// SetDefault gets a reference to the given NullableBool and assigns it to the Default field.
func (o *EditProfileDetails) SetDefault(v bool) {
	o.Default.Set(&v)
}
// SetDefaultNil sets the value for Default to be an explicit nil
func (o *EditProfileDetails) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil
func (o *EditProfileDetails) UnsetDefault() {
	o.Default.Unset()
}

// GetBrand returns the Brand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditProfileDetails) GetBrand() Brand {
	if o == nil || o.Brand.Get() == nil {
		var ret Brand
		return ret
	}
	return *o.Brand.Get()
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditProfileDetails) GetBrandOk() (*Brand, bool) {
	if o == nil {
		return nil, false
	}
	return o.Brand.Get(), o.Brand.IsSet()
}

// HasBrand returns a boolean if a field has been set.
func (o *EditProfileDetails) HasBrand() bool {
	if o != nil && o.Brand.IsSet() {
		return true
	}

	return false
}

// SetBrand gets a reference to the given NullableBrand and assigns it to the Brand field.
func (o *EditProfileDetails) SetBrand(v Brand) {
	o.Brand.Set(&v)
}
// SetBrandNil sets the value for Brand to be an explicit nil
func (o *EditProfileDetails) SetBrandNil() {
	o.Brand.Set(nil)
}

// UnsetBrand ensures that no value is present for Brand, not even an explicit nil
func (o *EditProfileDetails) UnsetBrand() {
	o.Brand.Unset()
}

// GetFunctionality returns the Functionality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditProfileDetails) GetFunctionality() Functionality {
	if o == nil || o.Functionality.Get() == nil {
		var ret Functionality
		return ret
	}
	return *o.Functionality.Get()
}

// GetFunctionalityOk returns a tuple with the Functionality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditProfileDetails) GetFunctionalityOk() (*Functionality, bool) {
	if o == nil {
		return nil, false
	}
	return o.Functionality.Get(), o.Functionality.IsSet()
}

// HasFunctionality returns a boolean if a field has been set.
func (o *EditProfileDetails) HasFunctionality() bool {
	if o != nil && o.Functionality.IsSet() {
		return true
	}

	return false
}

// SetFunctionality gets a reference to the given NullableFunctionality and assigns it to the Functionality field.
func (o *EditProfileDetails) SetFunctionality(v Functionality) {
	o.Functionality.Set(&v)
}
// SetFunctionalityNil sets the value for Functionality to be an explicit nil
func (o *EditProfileDetails) SetFunctionalityNil() {
	o.Functionality.Set(nil)
}

// UnsetFunctionality ensures that no value is present for Functionality, not even an explicit nil
func (o *EditProfileDetails) UnsetFunctionality() {
	o.Functionality.Unset()
}

// GetAspect returns the Aspect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditProfileDetails) GetAspect() Aspect {
	if o == nil || o.Aspect.Get() == nil {
		var ret Aspect
		return ret
	}
	return *o.Aspect.Get()
}

// GetAspectOk returns a tuple with the Aspect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditProfileDetails) GetAspectOk() (*Aspect, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aspect.Get(), o.Aspect.IsSet()
}

// HasAspect returns a boolean if a field has been set.
func (o *EditProfileDetails) HasAspect() bool {
	if o != nil && o.Aspect.IsSet() {
		return true
	}

	return false
}

// SetAspect gets a reference to the given NullableAspect and assigns it to the Aspect field.
func (o *EditProfileDetails) SetAspect(v Aspect) {
	o.Aspect.Set(&v)
}
// SetAspectNil sets the value for Aspect to be an explicit nil
func (o *EditProfileDetails) SetAspectNil() {
	o.Aspect.Set(nil)
}

// UnsetAspect ensures that no value is present for Aspect, not even an explicit nil
func (o *EditProfileDetails) UnsetAspect() {
	o.Aspect.Unset()
}

func (o EditProfileDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Brand.IsSet() {
		toSerialize["brand"] = o.Brand.Get()
	}
	if o.Functionality.IsSet() {
		toSerialize["functionality"] = o.Functionality.Get()
	}
	if o.Aspect.IsSet() {
		toSerialize["aspect"] = o.Aspect.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEditProfileDetails struct {
	value *EditProfileDetails
	isSet bool
}

func (v NullableEditProfileDetails) Get() *EditProfileDetails {
	return v.value
}

func (v *NullableEditProfileDetails) Set(val *EditProfileDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProfileDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProfileDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProfileDetails(val *EditProfileDetails) *NullableEditProfileDetails {
	return &NullableEditProfileDetails{value: val, isSet: true}
}

func (v NullableEditProfileDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProfileDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


