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

// Functionality Customization of web form functionality
type Functionality struct {
	// Whether a progress bar is shown on the web form, letting the user know on which step he is.<br/>&bull; <code>RENDER</code> - the progress bar will be shown;<br/>&bull; <code>HIDDEN</code> - the progress bar will NOT be shown.<br/><strong>NOTE:</strong> If no value is provided, then the following value will be applied by default when web form is opened: <code>RENDER</code>.
	ProgressBar NullableString `json:"progressBar,omitempty"`
	// How the bank login hint will be shown to the end user<br/>&bull; <code>EXPANDED</code> - the user will see the login hint and will have the option to collapse it;<br/>&bull; <code>COLLAPSED</code> - the login hint will be collapsed and the user can see it if he expands the field;<br/>&bull; <code>HIDDEN</code> - the login hint is hidden and the user cannot see it.<br/><strong>NOTE:</strong> If no value is provided, then the following value will be applied by default when web form is opened: <code>EXPANDED</code>.
	BankLoginHint NullableString `json:"bankLoginHint,omitempty"`
	// Whether the user will have a checkbox to ask for storing login secrets (like a PIN) in finAPI or not.<br/>&bull; <code>RENDER</code> - the checkbox will be shown;<br/>&bull; <code>HIDDEN</code> - the checkbox will NOT be shown;<br/>&bull; <code>MANDATORY</code> - the checkbox will be shown and it will be mandatory for the end user to check it in order to continue.<br/>&bull; <code>IMPLICIT_APPROVAL</code> - the checkbox will NOT be shown but login secrets will nevertheless be stored;<br/>&nbsp;&nbsp;&nbsp;&nbsp;<strong>NOTE:</strong> This value will also automatically store the TAN method. This value can be applied ONLY by our support team. Please contact <a href='mailto:support@finapi.io'>support@finapi.io</a> with the <code>profile.id</code> as soon as you have finalized the customization for other parameters.<br/><strong>NOTE:</strong> If no value is provided, then the following value will be applied by default when web form is opened: <code>RENDER</code>.
	StoreSecrets NullableString `json:"storeSecrets,omitempty"`
	// Whether the user will be allowed to change the selected bank, in case a BLZ/BIC/IBAN was sent in the API request by the client.<br/>&bull; <code>LOCKED</code> - the user will be directly routed to login to the pre-selected bank;<br/>&bull; <code>EDITABLE</code> - the user will see the pre-selected bank and have the option to change it.<br/><strong>NOTE:</strong> If no value is provided, then the following value will be applied by default when web form is opened: <code>LOCKED</code>.
	BankDetails NullableString `json:"bankDetails,omitempty"`
	// Whether the header will be displayed on the web form.<br/>&bull; <code>RENDER</code> - the header will be shown;<br/>&bull; <code>HIDDEN</code> - the header will NOT be shown.<br/><strong>NOTE:</strong> If no value is provided, then the following value will be applied by default when web form is opened: <code>RENDER</code>.
	Header NullableString `json:"header,omitempty"`
	Language NullableLanguage `json:"language,omitempty"`
	// When the web form is completed successfully, it determines whether the last view will be rendered. It applies to embedded and standalone web forms. It also applies to all endpoints in the \"Account Information Services\" and \"Payment Initiation Services\".<br/>If you are embedding the web form in your application, please set up appropriate handling for the 'onComplete' method to take advantage of the feature.<br/><strong>NOTE:</strong> If no value is provided, then the following value will be applied by default when web form is opened: <code>false</code>
	SkipConfirmationView NullableBool `json:"skipConfirmationView,omitempty"`
	// Whether or not the Web Form will render the \"Account Selection View\" for the end-user to choose which of the imported accounts should be saved to the database and available on the customer application.<br/><strong>NOTE:</strong> If no value is provided, then the following value will be applied by default when web form is opened: <code>false</code>
	RenderAccountSelectionView NullableBool `json:"renderAccountSelectionView,omitempty"`
	// Whether the entire payment summary is rendered on the Web Form. When set to TRUE, the counterpart data is not rendered on the Payment Summary of the Web Form.<br/><strong>NOTE:</strong> This value can be applied ONLY by our support team. Please contact <a href='mailto:support@finapi.io'>support@finapi.io</a> with the <code>profile.id</code> as soon as you have finalized the customization for other parameters.<br/><strong>NOTE:</strong> If no value is provided, then the following value will be applied by default when web form is opened: <code>false</code>
	HidePaymentSummary NullableBool `json:"hidePaymentSummary,omitempty"`
}

// NewFunctionality instantiates a new Functionality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionality() *Functionality {
	this := Functionality{}
	return &this
}

// NewFunctionalityWithDefaults instantiates a new Functionality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionalityWithDefaults() *Functionality {
	this := Functionality{}
	return &this
}

// GetProgressBar returns the ProgressBar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetProgressBar() string {
	if o == nil || o.ProgressBar.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProgressBar.Get()
}

// GetProgressBarOk returns a tuple with the ProgressBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetProgressBarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProgressBar.Get(), o.ProgressBar.IsSet()
}

// HasProgressBar returns a boolean if a field has been set.
func (o *Functionality) HasProgressBar() bool {
	if o != nil && o.ProgressBar.IsSet() {
		return true
	}

	return false
}

// SetProgressBar gets a reference to the given NullableString and assigns it to the ProgressBar field.
func (o *Functionality) SetProgressBar(v string) {
	o.ProgressBar.Set(&v)
}
// SetProgressBarNil sets the value for ProgressBar to be an explicit nil
func (o *Functionality) SetProgressBarNil() {
	o.ProgressBar.Set(nil)
}

// UnsetProgressBar ensures that no value is present for ProgressBar, not even an explicit nil
func (o *Functionality) UnsetProgressBar() {
	o.ProgressBar.Unset()
}

// GetBankLoginHint returns the BankLoginHint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetBankLoginHint() string {
	if o == nil || o.BankLoginHint.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankLoginHint.Get()
}

// GetBankLoginHintOk returns a tuple with the BankLoginHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetBankLoginHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankLoginHint.Get(), o.BankLoginHint.IsSet()
}

// HasBankLoginHint returns a boolean if a field has been set.
func (o *Functionality) HasBankLoginHint() bool {
	if o != nil && o.BankLoginHint.IsSet() {
		return true
	}

	return false
}

// SetBankLoginHint gets a reference to the given NullableString and assigns it to the BankLoginHint field.
func (o *Functionality) SetBankLoginHint(v string) {
	o.BankLoginHint.Set(&v)
}
// SetBankLoginHintNil sets the value for BankLoginHint to be an explicit nil
func (o *Functionality) SetBankLoginHintNil() {
	o.BankLoginHint.Set(nil)
}

// UnsetBankLoginHint ensures that no value is present for BankLoginHint, not even an explicit nil
func (o *Functionality) UnsetBankLoginHint() {
	o.BankLoginHint.Unset()
}

// GetStoreSecrets returns the StoreSecrets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetStoreSecrets() string {
	if o == nil || o.StoreSecrets.Get() == nil {
		var ret string
		return ret
	}
	return *o.StoreSecrets.Get()
}

// GetStoreSecretsOk returns a tuple with the StoreSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetStoreSecretsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoreSecrets.Get(), o.StoreSecrets.IsSet()
}

// HasStoreSecrets returns a boolean if a field has been set.
func (o *Functionality) HasStoreSecrets() bool {
	if o != nil && o.StoreSecrets.IsSet() {
		return true
	}

	return false
}

// SetStoreSecrets gets a reference to the given NullableString and assigns it to the StoreSecrets field.
func (o *Functionality) SetStoreSecrets(v string) {
	o.StoreSecrets.Set(&v)
}
// SetStoreSecretsNil sets the value for StoreSecrets to be an explicit nil
func (o *Functionality) SetStoreSecretsNil() {
	o.StoreSecrets.Set(nil)
}

// UnsetStoreSecrets ensures that no value is present for StoreSecrets, not even an explicit nil
func (o *Functionality) UnsetStoreSecrets() {
	o.StoreSecrets.Unset()
}

// GetBankDetails returns the BankDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetBankDetails() string {
	if o == nil || o.BankDetails.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankDetails.Get()
}

// GetBankDetailsOk returns a tuple with the BankDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetBankDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankDetails.Get(), o.BankDetails.IsSet()
}

// HasBankDetails returns a boolean if a field has been set.
func (o *Functionality) HasBankDetails() bool {
	if o != nil && o.BankDetails.IsSet() {
		return true
	}

	return false
}

// SetBankDetails gets a reference to the given NullableString and assigns it to the BankDetails field.
func (o *Functionality) SetBankDetails(v string) {
	o.BankDetails.Set(&v)
}
// SetBankDetailsNil sets the value for BankDetails to be an explicit nil
func (o *Functionality) SetBankDetailsNil() {
	o.BankDetails.Set(nil)
}

// UnsetBankDetails ensures that no value is present for BankDetails, not even an explicit nil
func (o *Functionality) UnsetBankDetails() {
	o.BankDetails.Unset()
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetHeader() string {
	if o == nil || o.Header.Get() == nil {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *Functionality) HasHeader() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
func (o *Functionality) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *Functionality) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *Functionality) UnsetHeader() {
	o.Header.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetLanguage() Language {
	if o == nil || o.Language.Get() == nil {
		var ret Language
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetLanguageOk() (*Language, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *Functionality) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableLanguage and assigns it to the Language field.
func (o *Functionality) SetLanguage(v Language) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *Functionality) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *Functionality) UnsetLanguage() {
	o.Language.Unset()
}

// GetSkipConfirmationView returns the SkipConfirmationView field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetSkipConfirmationView() bool {
	if o == nil || o.SkipConfirmationView.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SkipConfirmationView.Get()
}

// GetSkipConfirmationViewOk returns a tuple with the SkipConfirmationView field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetSkipConfirmationViewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipConfirmationView.Get(), o.SkipConfirmationView.IsSet()
}

// HasSkipConfirmationView returns a boolean if a field has been set.
func (o *Functionality) HasSkipConfirmationView() bool {
	if o != nil && o.SkipConfirmationView.IsSet() {
		return true
	}

	return false
}

// SetSkipConfirmationView gets a reference to the given NullableBool and assigns it to the SkipConfirmationView field.
func (o *Functionality) SetSkipConfirmationView(v bool) {
	o.SkipConfirmationView.Set(&v)
}
// SetSkipConfirmationViewNil sets the value for SkipConfirmationView to be an explicit nil
func (o *Functionality) SetSkipConfirmationViewNil() {
	o.SkipConfirmationView.Set(nil)
}

// UnsetSkipConfirmationView ensures that no value is present for SkipConfirmationView, not even an explicit nil
func (o *Functionality) UnsetSkipConfirmationView() {
	o.SkipConfirmationView.Unset()
}

// GetRenderAccountSelectionView returns the RenderAccountSelectionView field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetRenderAccountSelectionView() bool {
	if o == nil || o.RenderAccountSelectionView.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RenderAccountSelectionView.Get()
}

// GetRenderAccountSelectionViewOk returns a tuple with the RenderAccountSelectionView field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetRenderAccountSelectionViewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenderAccountSelectionView.Get(), o.RenderAccountSelectionView.IsSet()
}

// HasRenderAccountSelectionView returns a boolean if a field has been set.
func (o *Functionality) HasRenderAccountSelectionView() bool {
	if o != nil && o.RenderAccountSelectionView.IsSet() {
		return true
	}

	return false
}

// SetRenderAccountSelectionView gets a reference to the given NullableBool and assigns it to the RenderAccountSelectionView field.
func (o *Functionality) SetRenderAccountSelectionView(v bool) {
	o.RenderAccountSelectionView.Set(&v)
}
// SetRenderAccountSelectionViewNil sets the value for RenderAccountSelectionView to be an explicit nil
func (o *Functionality) SetRenderAccountSelectionViewNil() {
	o.RenderAccountSelectionView.Set(nil)
}

// UnsetRenderAccountSelectionView ensures that no value is present for RenderAccountSelectionView, not even an explicit nil
func (o *Functionality) UnsetRenderAccountSelectionView() {
	o.RenderAccountSelectionView.Unset()
}

// GetHidePaymentSummary returns the HidePaymentSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Functionality) GetHidePaymentSummary() bool {
	if o == nil || o.HidePaymentSummary.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HidePaymentSummary.Get()
}

// GetHidePaymentSummaryOk returns a tuple with the HidePaymentSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Functionality) GetHidePaymentSummaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HidePaymentSummary.Get(), o.HidePaymentSummary.IsSet()
}

// HasHidePaymentSummary returns a boolean if a field has been set.
func (o *Functionality) HasHidePaymentSummary() bool {
	if o != nil && o.HidePaymentSummary.IsSet() {
		return true
	}

	return false
}

// SetHidePaymentSummary gets a reference to the given NullableBool and assigns it to the HidePaymentSummary field.
func (o *Functionality) SetHidePaymentSummary(v bool) {
	o.HidePaymentSummary.Set(&v)
}
// SetHidePaymentSummaryNil sets the value for HidePaymentSummary to be an explicit nil
func (o *Functionality) SetHidePaymentSummaryNil() {
	o.HidePaymentSummary.Set(nil)
}

// UnsetHidePaymentSummary ensures that no value is present for HidePaymentSummary, not even an explicit nil
func (o *Functionality) UnsetHidePaymentSummary() {
	o.HidePaymentSummary.Unset()
}

func (o Functionality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProgressBar.IsSet() {
		toSerialize["progressBar"] = o.ProgressBar.Get()
	}
	if o.BankLoginHint.IsSet() {
		toSerialize["bankLoginHint"] = o.BankLoginHint.Get()
	}
	if o.StoreSecrets.IsSet() {
		toSerialize["storeSecrets"] = o.StoreSecrets.Get()
	}
	if o.BankDetails.IsSet() {
		toSerialize["bankDetails"] = o.BankDetails.Get()
	}
	if o.Header.IsSet() {
		toSerialize["header"] = o.Header.Get()
	}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.SkipConfirmationView.IsSet() {
		toSerialize["skipConfirmationView"] = o.SkipConfirmationView.Get()
	}
	if o.RenderAccountSelectionView.IsSet() {
		toSerialize["renderAccountSelectionView"] = o.RenderAccountSelectionView.Get()
	}
	if o.HidePaymentSummary.IsSet() {
		toSerialize["hidePaymentSummary"] = o.HidePaymentSummary.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionality struct {
	value *Functionality
	isSet bool
}

func (v NullableFunctionality) Get() *Functionality {
	return v.value
}

func (v *NullableFunctionality) Set(val *Functionality) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionality) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionality(val *Functionality) *NullableFunctionality {
	return &NullableFunctionality{value: val, isSet: true}
}

func (v NullableFunctionality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


