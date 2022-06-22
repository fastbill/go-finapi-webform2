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

// UserMetadata With the migration to PSD2 APIs, a new term called \"User metadata\" (also known as \"PSU metadata\") has been introduced to the API. This user metadata aims to inform the banking API if there was a real end-user behind an HTTP request or if the request was triggered by a system (e.g. by an automatic batch update). In the latter case, the bank may apply some restrictions such as limiting the number of HTTP requests for a single consent. Also, some operations may be forbidden entirely by the banking API. For example, some banks do not allow issuing a new consent without the end-user being involved. Therefore, the PSU metadata must always be provided for such operations.<br/><br/>As finAPI does not have direct interaction with the end-user for the background update, it is the client application's responsibility to provide all the necessary information about the end-user.
type UserMetadata struct {
	// The IP address of the user's device
	IpAddress string `json:"ipAddress"`
	// The user's device and/or operating system identification
	DeviceOs string `json:"deviceOs"`
	// The user's web browser or other client device identification
	UserAgent string `json:"userAgent"`
}

// NewUserMetadata instantiates a new UserMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadata(ipAddress string, deviceOs string, userAgent string) *UserMetadata {
	this := UserMetadata{}
	this.IpAddress = ipAddress
	this.DeviceOs = deviceOs
	this.UserAgent = userAgent
	return &this
}

// NewUserMetadataWithDefaults instantiates a new UserMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataWithDefaults() *UserMetadata {
	this := UserMetadata{}
	return &this
}

// GetIpAddress returns the IpAddress field value
func (o *UserMetadata) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *UserMetadata) SetIpAddress(v string) {
	o.IpAddress = v
}

// GetDeviceOs returns the DeviceOs field value
func (o *UserMetadata) GetDeviceOs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceOs
}

// GetDeviceOsOk returns a tuple with the DeviceOs field value
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetDeviceOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceOs, true
}

// SetDeviceOs sets field value
func (o *UserMetadata) SetDeviceOs(v string) {
	o.DeviceOs = v
}

// GetUserAgent returns the UserAgent field value
func (o *UserMetadata) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *UserMetadata) SetUserAgent(v string) {
	o.UserAgent = v
}

func (o UserMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if true {
		toSerialize["deviceOs"] = o.DeviceOs
	}
	if true {
		toSerialize["userAgent"] = o.UserAgent
	}
	return json.Marshal(toSerialize)
}

type NullableUserMetadata struct {
	value *UserMetadata
	isSet bool
}

func (v NullableUserMetadata) Get() *UserMetadata {
	return v.value
}

func (v *NullableUserMetadata) Set(val *UserMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadata(val *UserMetadata) *NullableUserMetadata {
	return &NullableUserMetadata{value: val, isSet: true}
}

func (v NullableUserMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

