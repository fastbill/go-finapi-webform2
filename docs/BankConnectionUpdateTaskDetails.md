# BankConnectionUpdateTaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionId** | **int64** | Bank connection identifier. | 
**ImportNewAccounts** | Pointer to **NullableBool** | Whether new accounts that have not yet been imported will be imported or not.&lt;br/&gt;&lt;br/&gt;By setting this parameter to true, we will try to get new accounts the user might have at the bank. The user will have a possibility to stop the process once he finds the accounts he is interested in have been imported. The set of newly imported accounts can be limited by using \&quot;accountTypes\&quot; parameter. | [optional] [default to false]
**AccountTypes** | Pointer to [**[]AccountType**](AccountType.md) | It defines the account types to be added on the bank connection. If no value is given, then all new available accounts will be imported.&lt;br/&gt;Only applied if \&quot;importNewAccounts\&quot; is set to \&quot;true\&quot; and ignored otherwise.&lt;br/&gt;This parameter refers to the same parameters of \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/update&#39;&gt;Update a bank connection&lt;/a&gt;\&quot; and \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/connectInterface&#39;&gt;Connect a new interface&lt;/a&gt;\&quot; services in the finAPI Access API. | [optional] 
**SkipPositionsDownload** | Pointer to **NullableBool** | Whether to skip the download of transactions and securities or not.&lt;br/&gt;This parameter refers to the same parameters of \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/update&#39;&gt;Update a bank connection&lt;/a&gt;\&quot; and \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/connectInterface&#39;&gt;Connect a new interface&lt;/a&gt;\&quot; services in the finAPI Access API. | [optional] [default to false]
**LoadOwnerData** | Pointer to **NullableBool** | Whether to load information about the bank connection owner(s).&lt;br/&gt;This parameter refers to the same parameters of \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/update&#39;&gt;Update a bank connection&lt;/a&gt;\&quot; and \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/connectInterface&#39;&gt;Connect a new interface&lt;/a&gt;\&quot; services in the finAPI Access API.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; This feature is supported only by the WEB_SCRAPER interface. | [optional] [default to false]
**ManageSavedSettings** | Pointer to [**[]ManageSavedSettings**](ManageSavedSettings.md) | If users have stored bank credentials or their preferred SCA procedure in finAPI, then each time it is necessary, the web form will automatically apply the saved settings. This field will force the web form to be presented. Depending on the value you provide in the API, the end-user will have the possibility to provide new credentials, decide if he wants to store the new credentials in finAPI or delete the stored SCA procedure. He might also be able to save a different SCA procedure as preference if the workflow requires it. Use this parameter, for example, if a previous update failed because of invalid credentials, or you want to allow the end-user to reset his preferred SCA method.&lt;br/&gt;&lt;strong&gt;Manage saved settings:&lt;/strong&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;CREDENTIALS&lt;/code&gt; - the end user can edit his stored credentials;&lt;br/&gt;&amp;bull; &lt;code&gt;DEFAULT_TWO_STEP_PROCEDURE&lt;/code&gt; - the end user can edit his stored default two step procedure. | [optional] 
**Callbacks** | Pointer to [**NullableTaskCallbacks**](TaskCallbacks.md) |  | [optional] 
**ProfileId** | Pointer to **NullableString** | The profile to be applied to the web form.&lt;br/&gt;This will overwrite the default profile, if such a profile exists. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The URL where the end-user will be redirected to after completing the bank login and (possibly) the SCA on the bank&#39;s website. Must always be provided by mandators with &lt;code&gt;FULLY_LICENSED&lt;/code&gt; or &lt;code&gt;AISP&lt;/code&gt; license type, and may not be provided by mandators with other license types. Find more info in the &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://documentation.finapi.io/webform/Licensed-customers-using-the-Web-Form.2832302195.html&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | [optional] 
**UserMetadata** | Pointer to [**NullableUserMetadata**](UserMetadata.md) |  | [optional] 

## Methods

### NewBankConnectionUpdateTaskDetails

`func NewBankConnectionUpdateTaskDetails(bankConnectionId int64, ) *BankConnectionUpdateTaskDetails`

NewBankConnectionUpdateTaskDetails instantiates a new BankConnectionUpdateTaskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankConnectionUpdateTaskDetailsWithDefaults

`func NewBankConnectionUpdateTaskDetailsWithDefaults() *BankConnectionUpdateTaskDetails`

NewBankConnectionUpdateTaskDetailsWithDefaults instantiates a new BankConnectionUpdateTaskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankConnectionId

`func (o *BankConnectionUpdateTaskDetails) GetBankConnectionId() int64`

GetBankConnectionId returns the BankConnectionId field if non-nil, zero value otherwise.

### GetBankConnectionIdOk

`func (o *BankConnectionUpdateTaskDetails) GetBankConnectionIdOk() (*int64, bool)`

GetBankConnectionIdOk returns a tuple with the BankConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionId

`func (o *BankConnectionUpdateTaskDetails) SetBankConnectionId(v int64)`

SetBankConnectionId sets BankConnectionId field to given value.


### GetImportNewAccounts

`func (o *BankConnectionUpdateTaskDetails) GetImportNewAccounts() bool`

GetImportNewAccounts returns the ImportNewAccounts field if non-nil, zero value otherwise.

### GetImportNewAccountsOk

`func (o *BankConnectionUpdateTaskDetails) GetImportNewAccountsOk() (*bool, bool)`

GetImportNewAccountsOk returns a tuple with the ImportNewAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNewAccounts

`func (o *BankConnectionUpdateTaskDetails) SetImportNewAccounts(v bool)`

SetImportNewAccounts sets ImportNewAccounts field to given value.

### HasImportNewAccounts

`func (o *BankConnectionUpdateTaskDetails) HasImportNewAccounts() bool`

HasImportNewAccounts returns a boolean if a field has been set.

### SetImportNewAccountsNil

`func (o *BankConnectionUpdateTaskDetails) SetImportNewAccountsNil(b bool)`

 SetImportNewAccountsNil sets the value for ImportNewAccounts to be an explicit nil

### UnsetImportNewAccounts
`func (o *BankConnectionUpdateTaskDetails) UnsetImportNewAccounts()`

UnsetImportNewAccounts ensures that no value is present for ImportNewAccounts, not even an explicit nil
### GetAccountTypes

`func (o *BankConnectionUpdateTaskDetails) GetAccountTypes() []AccountType`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *BankConnectionUpdateTaskDetails) GetAccountTypesOk() (*[]AccountType, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *BankConnectionUpdateTaskDetails) SetAccountTypes(v []AccountType)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *BankConnectionUpdateTaskDetails) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### SetAccountTypesNil

`func (o *BankConnectionUpdateTaskDetails) SetAccountTypesNil(b bool)`

 SetAccountTypesNil sets the value for AccountTypes to be an explicit nil

### UnsetAccountTypes
`func (o *BankConnectionUpdateTaskDetails) UnsetAccountTypes()`

UnsetAccountTypes ensures that no value is present for AccountTypes, not even an explicit nil
### GetSkipPositionsDownload

`func (o *BankConnectionUpdateTaskDetails) GetSkipPositionsDownload() bool`

GetSkipPositionsDownload returns the SkipPositionsDownload field if non-nil, zero value otherwise.

### GetSkipPositionsDownloadOk

`func (o *BankConnectionUpdateTaskDetails) GetSkipPositionsDownloadOk() (*bool, bool)`

GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPositionsDownload

`func (o *BankConnectionUpdateTaskDetails) SetSkipPositionsDownload(v bool)`

SetSkipPositionsDownload sets SkipPositionsDownload field to given value.

### HasSkipPositionsDownload

`func (o *BankConnectionUpdateTaskDetails) HasSkipPositionsDownload() bool`

HasSkipPositionsDownload returns a boolean if a field has been set.

### SetSkipPositionsDownloadNil

`func (o *BankConnectionUpdateTaskDetails) SetSkipPositionsDownloadNil(b bool)`

 SetSkipPositionsDownloadNil sets the value for SkipPositionsDownload to be an explicit nil

### UnsetSkipPositionsDownload
`func (o *BankConnectionUpdateTaskDetails) UnsetSkipPositionsDownload()`

UnsetSkipPositionsDownload ensures that no value is present for SkipPositionsDownload, not even an explicit nil
### GetLoadOwnerData

`func (o *BankConnectionUpdateTaskDetails) GetLoadOwnerData() bool`

GetLoadOwnerData returns the LoadOwnerData field if non-nil, zero value otherwise.

### GetLoadOwnerDataOk

`func (o *BankConnectionUpdateTaskDetails) GetLoadOwnerDataOk() (*bool, bool)`

GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOwnerData

`func (o *BankConnectionUpdateTaskDetails) SetLoadOwnerData(v bool)`

SetLoadOwnerData sets LoadOwnerData field to given value.

### HasLoadOwnerData

`func (o *BankConnectionUpdateTaskDetails) HasLoadOwnerData() bool`

HasLoadOwnerData returns a boolean if a field has been set.

### SetLoadOwnerDataNil

`func (o *BankConnectionUpdateTaskDetails) SetLoadOwnerDataNil(b bool)`

 SetLoadOwnerDataNil sets the value for LoadOwnerData to be an explicit nil

### UnsetLoadOwnerData
`func (o *BankConnectionUpdateTaskDetails) UnsetLoadOwnerData()`

UnsetLoadOwnerData ensures that no value is present for LoadOwnerData, not even an explicit nil
### GetManageSavedSettings

`func (o *BankConnectionUpdateTaskDetails) GetManageSavedSettings() []ManageSavedSettings`

GetManageSavedSettings returns the ManageSavedSettings field if non-nil, zero value otherwise.

### GetManageSavedSettingsOk

`func (o *BankConnectionUpdateTaskDetails) GetManageSavedSettingsOk() (*[]ManageSavedSettings, bool)`

GetManageSavedSettingsOk returns a tuple with the ManageSavedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageSavedSettings

`func (o *BankConnectionUpdateTaskDetails) SetManageSavedSettings(v []ManageSavedSettings)`

SetManageSavedSettings sets ManageSavedSettings field to given value.

### HasManageSavedSettings

`func (o *BankConnectionUpdateTaskDetails) HasManageSavedSettings() bool`

HasManageSavedSettings returns a boolean if a field has been set.

### SetManageSavedSettingsNil

`func (o *BankConnectionUpdateTaskDetails) SetManageSavedSettingsNil(b bool)`

 SetManageSavedSettingsNil sets the value for ManageSavedSettings to be an explicit nil

### UnsetManageSavedSettings
`func (o *BankConnectionUpdateTaskDetails) UnsetManageSavedSettings()`

UnsetManageSavedSettings ensures that no value is present for ManageSavedSettings, not even an explicit nil
### GetCallbacks

`func (o *BankConnectionUpdateTaskDetails) GetCallbacks() TaskCallbacks`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *BankConnectionUpdateTaskDetails) GetCallbacksOk() (*TaskCallbacks, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *BankConnectionUpdateTaskDetails) SetCallbacks(v TaskCallbacks)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *BankConnectionUpdateTaskDetails) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### SetCallbacksNil

`func (o *BankConnectionUpdateTaskDetails) SetCallbacksNil(b bool)`

 SetCallbacksNil sets the value for Callbacks to be an explicit nil

### UnsetCallbacks
`func (o *BankConnectionUpdateTaskDetails) UnsetCallbacks()`

UnsetCallbacks ensures that no value is present for Callbacks, not even an explicit nil
### GetProfileId

`func (o *BankConnectionUpdateTaskDetails) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *BankConnectionUpdateTaskDetails) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *BankConnectionUpdateTaskDetails) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *BankConnectionUpdateTaskDetails) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *BankConnectionUpdateTaskDetails) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *BankConnectionUpdateTaskDetails) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetRedirectUrl

`func (o *BankConnectionUpdateTaskDetails) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *BankConnectionUpdateTaskDetails) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *BankConnectionUpdateTaskDetails) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *BankConnectionUpdateTaskDetails) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *BankConnectionUpdateTaskDetails) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *BankConnectionUpdateTaskDetails) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetUserMetadata

`func (o *BankConnectionUpdateTaskDetails) GetUserMetadata() UserMetadata`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *BankConnectionUpdateTaskDetails) GetUserMetadataOk() (*UserMetadata, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *BankConnectionUpdateTaskDetails) SetUserMetadata(v UserMetadata)`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *BankConnectionUpdateTaskDetails) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.

### SetUserMetadataNil

`func (o *BankConnectionUpdateTaskDetails) SetUserMetadataNil(b bool)`

 SetUserMetadataNil sets the value for UserMetadata to be an explicit nil

### UnsetUserMetadata
`func (o *BankConnectionUpdateTaskDetails) UnsetUserMetadata()`

UnsetUserMetadata ensures that no value is present for UserMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


