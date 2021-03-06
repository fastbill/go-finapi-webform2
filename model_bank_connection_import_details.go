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

// BankConnectionImportDetails struct for BankConnectionImportDetails
type BankConnectionImportDetails struct {
	Bank NullableImportBankDetails `json:"bank,omitempty"`
	// Custom name for the bank connection
	BankConnectionName NullableString `json:"bankConnectionName,omitempty"`
	// Whether to skip the download of transactions and securities or not.<br/>This parameter refers to the same parameter of the \"<a target='_blank' href='https://docs.finapi.io/?product=access#post-/api/v1/bankConnections/import'>Import a new bank connection</a>\" service in the finAPI Access API.
	SkipPositionsDownload NullableBool `json:"skipPositionsDownload,omitempty"`
	// Whether to load information about the bank connection owner(s).<br/>This parameter refers to the same parameter of the \"<a target='_blank' href='https://docs.finapi.io/?product=access#post-/api/v1/bankConnections/import'>Import a new bank connection</a>\" service in the finAPI Access API.
	LoadOwnerData NullableBool `json:"loadOwnerData,omitempty"`
	// Defines the limit of how many days of the transaction history will be requested from the bank.<br/>This parameter refers to the same parameter of the \"<a target='_blank' href='https://docs.finapi.io/?product=access#post-/api/v1/bankConnections/import'>Import a new bank connection</a>\" service in the finAPI Access API.
	MaxDaysForDownload NullableInt32 `json:"maxDaysForDownload,omitempty"`
	// A set of account types that are considered for the import. If no values is given, then all accounts will be imported.<br/>This parameter refers to the same parameter of the \"<a target='_blank' href='https://docs.finapi.io/?product=access#post-/api/v1/bankConnections/import'>Import a new bank connection</a>\" service in the finAPI Access API.
	AccountTypes []AccountType `json:"accountTypes,omitempty"`
	Callbacks NullableCallbacks `json:"callbacks,omitempty"`
	// The profile to be applied to the web form.<br/>This will overwrite the default profile, if such a profile exists.
	ProfileId NullableString `json:"profileId,omitempty"`
	// The URL where the end-user will be redirected to after completing the bank login and (possibly) the SCA on the bank's website. Must always be provided by mandators with <code>FULLY_LICENSED</code> or <code>AISP</code> license type, and may not be provided by mandators with other license types. Find more info in the <a target='_blank' href='https://documentation.finapi.io/webform/Licensed-customers-using-the-Web-Form.2832302195.html'>Web Form 2.0 Public Documentation</a>.
	RedirectUrl NullableString `json:"redirectUrl,omitempty"`
}

// NewBankConnectionImportDetails instantiates a new BankConnectionImportDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankConnectionImportDetails() *BankConnectionImportDetails {
	this := BankConnectionImportDetails{}
	var skipPositionsDownload bool = false
	this.SkipPositionsDownload = *NewNullableBool(&skipPositionsDownload)
	var loadOwnerData bool = false
	this.LoadOwnerData = *NewNullableBool(&loadOwnerData)
	var maxDaysForDownload int32 = 0
	this.MaxDaysForDownload = *NewNullableInt32(&maxDaysForDownload)
	return &this
}

// NewBankConnectionImportDetailsWithDefaults instantiates a new BankConnectionImportDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankConnectionImportDetailsWithDefaults() *BankConnectionImportDetails {
	this := BankConnectionImportDetails{}
	var skipPositionsDownload bool = false
	this.SkipPositionsDownload = *NewNullableBool(&skipPositionsDownload)
	var loadOwnerData bool = false
	this.LoadOwnerData = *NewNullableBool(&loadOwnerData)
	var maxDaysForDownload int32 = 0
	this.MaxDaysForDownload = *NewNullableInt32(&maxDaysForDownload)
	return &this
}

// GetBank returns the Bank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetBank() ImportBankDetails {
	if o == nil || o.Bank.Get() == nil {
		var ret ImportBankDetails
		return ret
	}
	return *o.Bank.Get()
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetBankOk() (*ImportBankDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bank.Get(), o.Bank.IsSet()
}

// HasBank returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasBank() bool {
	if o != nil && o.Bank.IsSet() {
		return true
	}

	return false
}

// SetBank gets a reference to the given NullableImportBankDetails and assigns it to the Bank field.
func (o *BankConnectionImportDetails) SetBank(v ImportBankDetails) {
	o.Bank.Set(&v)
}
// SetBankNil sets the value for Bank to be an explicit nil
func (o *BankConnectionImportDetails) SetBankNil() {
	o.Bank.Set(nil)
}

// UnsetBank ensures that no value is present for Bank, not even an explicit nil
func (o *BankConnectionImportDetails) UnsetBank() {
	o.Bank.Unset()
}

// GetBankConnectionName returns the BankConnectionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetBankConnectionName() string {
	if o == nil || o.BankConnectionName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankConnectionName.Get()
}

// GetBankConnectionNameOk returns a tuple with the BankConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetBankConnectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankConnectionName.Get(), o.BankConnectionName.IsSet()
}

// HasBankConnectionName returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasBankConnectionName() bool {
	if o != nil && o.BankConnectionName.IsSet() {
		return true
	}

	return false
}

// SetBankConnectionName gets a reference to the given NullableString and assigns it to the BankConnectionName field.
func (o *BankConnectionImportDetails) SetBankConnectionName(v string) {
	o.BankConnectionName.Set(&v)
}
// SetBankConnectionNameNil sets the value for BankConnectionName to be an explicit nil
func (o *BankConnectionImportDetails) SetBankConnectionNameNil() {
	o.BankConnectionName.Set(nil)
}

// UnsetBankConnectionName ensures that no value is present for BankConnectionName, not even an explicit nil
func (o *BankConnectionImportDetails) UnsetBankConnectionName() {
	o.BankConnectionName.Unset()
}

// GetSkipPositionsDownload returns the SkipPositionsDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetSkipPositionsDownload() bool {
	if o == nil || o.SkipPositionsDownload.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SkipPositionsDownload.Get()
}

// GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetSkipPositionsDownloadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipPositionsDownload.Get(), o.SkipPositionsDownload.IsSet()
}

// HasSkipPositionsDownload returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasSkipPositionsDownload() bool {
	if o != nil && o.SkipPositionsDownload.IsSet() {
		return true
	}

	return false
}

// SetSkipPositionsDownload gets a reference to the given NullableBool and assigns it to the SkipPositionsDownload field.
func (o *BankConnectionImportDetails) SetSkipPositionsDownload(v bool) {
	o.SkipPositionsDownload.Set(&v)
}
// SetSkipPositionsDownloadNil sets the value for SkipPositionsDownload to be an explicit nil
func (o *BankConnectionImportDetails) SetSkipPositionsDownloadNil() {
	o.SkipPositionsDownload.Set(nil)
}

// UnsetSkipPositionsDownload ensures that no value is present for SkipPositionsDownload, not even an explicit nil
func (o *BankConnectionImportDetails) UnsetSkipPositionsDownload() {
	o.SkipPositionsDownload.Unset()
}

// GetLoadOwnerData returns the LoadOwnerData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetLoadOwnerData() bool {
	if o == nil || o.LoadOwnerData.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LoadOwnerData.Get()
}

// GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetLoadOwnerDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoadOwnerData.Get(), o.LoadOwnerData.IsSet()
}

// HasLoadOwnerData returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasLoadOwnerData() bool {
	if o != nil && o.LoadOwnerData.IsSet() {
		return true
	}

	return false
}

// SetLoadOwnerData gets a reference to the given NullableBool and assigns it to the LoadOwnerData field.
func (o *BankConnectionImportDetails) SetLoadOwnerData(v bool) {
	o.LoadOwnerData.Set(&v)
}
// SetLoadOwnerDataNil sets the value for LoadOwnerData to be an explicit nil
func (o *BankConnectionImportDetails) SetLoadOwnerDataNil() {
	o.LoadOwnerData.Set(nil)
}

// UnsetLoadOwnerData ensures that no value is present for LoadOwnerData, not even an explicit nil
func (o *BankConnectionImportDetails) UnsetLoadOwnerData() {
	o.LoadOwnerData.Unset()
}

// GetMaxDaysForDownload returns the MaxDaysForDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetMaxDaysForDownload() int32 {
	if o == nil || o.MaxDaysForDownload.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxDaysForDownload.Get()
}

// GetMaxDaysForDownloadOk returns a tuple with the MaxDaysForDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetMaxDaysForDownloadOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxDaysForDownload.Get(), o.MaxDaysForDownload.IsSet()
}

// HasMaxDaysForDownload returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasMaxDaysForDownload() bool {
	if o != nil && o.MaxDaysForDownload.IsSet() {
		return true
	}

	return false
}

// SetMaxDaysForDownload gets a reference to the given NullableInt32 and assigns it to the MaxDaysForDownload field.
func (o *BankConnectionImportDetails) SetMaxDaysForDownload(v int32) {
	o.MaxDaysForDownload.Set(&v)
}
// SetMaxDaysForDownloadNil sets the value for MaxDaysForDownload to be an explicit nil
func (o *BankConnectionImportDetails) SetMaxDaysForDownloadNil() {
	o.MaxDaysForDownload.Set(nil)
}

// UnsetMaxDaysForDownload ensures that no value is present for MaxDaysForDownload, not even an explicit nil
func (o *BankConnectionImportDetails) UnsetMaxDaysForDownload() {
	o.MaxDaysForDownload.Unset()
}

// GetAccountTypes returns the AccountTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetAccountTypes() []AccountType {
	if o == nil {
		var ret []AccountType
		return ret
	}
	return o.AccountTypes
}

// GetAccountTypesOk returns a tuple with the AccountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetAccountTypesOk() ([]AccountType, bool) {
	if o == nil || o.AccountTypes == nil {
		return nil, false
	}
	return o.AccountTypes, true
}

// HasAccountTypes returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasAccountTypes() bool {
	if o != nil && o.AccountTypes != nil {
		return true
	}

	return false
}

// SetAccountTypes gets a reference to the given []AccountType and assigns it to the AccountTypes field.
func (o *BankConnectionImportDetails) SetAccountTypes(v []AccountType) {
	o.AccountTypes = v
}

// GetCallbacks returns the Callbacks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetCallbacks() Callbacks {
	if o == nil || o.Callbacks.Get() == nil {
		var ret Callbacks
		return ret
	}
	return *o.Callbacks.Get()
}

// GetCallbacksOk returns a tuple with the Callbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetCallbacksOk() (*Callbacks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Callbacks.Get(), o.Callbacks.IsSet()
}

// HasCallbacks returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasCallbacks() bool {
	if o != nil && o.Callbacks.IsSet() {
		return true
	}

	return false
}

// SetCallbacks gets a reference to the given NullableCallbacks and assigns it to the Callbacks field.
func (o *BankConnectionImportDetails) SetCallbacks(v Callbacks) {
	o.Callbacks.Set(&v)
}
// SetCallbacksNil sets the value for Callbacks to be an explicit nil
func (o *BankConnectionImportDetails) SetCallbacksNil() {
	o.Callbacks.Set(nil)
}

// UnsetCallbacks ensures that no value is present for Callbacks, not even an explicit nil
func (o *BankConnectionImportDetails) UnsetCallbacks() {
	o.Callbacks.Unset()
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetProfileId() string {
	if o == nil || o.ProfileId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProfileId.Get()
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileId.Get(), o.ProfileId.IsSet()
}

// HasProfileId returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasProfileId() bool {
	if o != nil && o.ProfileId.IsSet() {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given NullableString and assigns it to the ProfileId field.
func (o *BankConnectionImportDetails) SetProfileId(v string) {
	o.ProfileId.Set(&v)
}
// SetProfileIdNil sets the value for ProfileId to be an explicit nil
func (o *BankConnectionImportDetails) SetProfileIdNil() {
	o.ProfileId.Set(nil)
}

// UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
func (o *BankConnectionImportDetails) UnsetProfileId() {
	o.ProfileId.Unset()
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankConnectionImportDetails) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionImportDetails) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *BankConnectionImportDetails) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given NullableString and assigns it to the RedirectUrl field.
func (o *BankConnectionImportDetails) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *BankConnectionImportDetails) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *BankConnectionImportDetails) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

func (o BankConnectionImportDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bank.IsSet() {
		toSerialize["bank"] = o.Bank.Get()
	}
	if o.BankConnectionName.IsSet() {
		toSerialize["bankConnectionName"] = o.BankConnectionName.Get()
	}
	if o.SkipPositionsDownload.IsSet() {
		toSerialize["skipPositionsDownload"] = o.SkipPositionsDownload.Get()
	}
	if o.LoadOwnerData.IsSet() {
		toSerialize["loadOwnerData"] = o.LoadOwnerData.Get()
	}
	if o.MaxDaysForDownload.IsSet() {
		toSerialize["maxDaysForDownload"] = o.MaxDaysForDownload.Get()
	}
	if o.AccountTypes != nil {
		toSerialize["accountTypes"] = o.AccountTypes
	}
	if o.Callbacks.IsSet() {
		toSerialize["callbacks"] = o.Callbacks.Get()
	}
	if o.ProfileId.IsSet() {
		toSerialize["profileId"] = o.ProfileId.Get()
	}
	if o.RedirectUrl.IsSet() {
		toSerialize["redirectUrl"] = o.RedirectUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBankConnectionImportDetails struct {
	value *BankConnectionImportDetails
	isSet bool
}

func (v NullableBankConnectionImportDetails) Get() *BankConnectionImportDetails {
	return v.value
}

func (v *NullableBankConnectionImportDetails) Set(val *BankConnectionImportDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBankConnectionImportDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBankConnectionImportDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankConnectionImportDetails(val *BankConnectionImportDetails) *NullableBankConnectionImportDetails {
	return &NullableBankConnectionImportDetails{value: val, isSet: true}
}

func (v NullableBankConnectionImportDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankConnectionImportDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


