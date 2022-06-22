# BankConnectionUpdateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | [**UpdateBankDetails**](UpdateBankDetails.md) |  | 
**ImportNewAccounts** | Pointer to **NullableBool** | Whether new accounts that have not yet been imported will be imported or not.&lt;br/&gt;&lt;br/&gt;By setting this parameter to true, we will try to get new accounts the user might have at the bank. The user will have a possibility to stop the process once he finds the accounts he is interested in have been imported. The set of newly imported accounts can be limited by using \&quot;accountTypes\&quot; parameter. | [optional] [default to false]
**AccountTypes** | Pointer to [**[]AccountType**](AccountType.md) | A set of account types that are considered for the adding. If no values is given, then all accounts will be imported. Only applied if \&quot;importNewAccounts\&quot; is set to \&quot;true\&quot; and ignored otherwise. &lt;br/&gt;This parameter refers to the same parameters of \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/update&#39;&gt;Update a bank connection&lt;/a&gt;\&quot; and \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/connectInterface&#39;&gt;Connect a new interface&lt;/a&gt;\&quot; services in the finAPI Access API. | [optional] 
**SkipPositionsDownload** | Pointer to **NullableBool** | Whether to skip the download of transactions and securities or not.&lt;br/&gt;This parameter refers to the same parameters of \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/update&#39;&gt;Update a bank connection&lt;/a&gt;\&quot; and \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/connectInterface&#39;&gt;Connect a new interface&lt;/a&gt;\&quot; services in the finAPI Access API. | [optional] [default to false]
**LoadOwnerData** | Pointer to **NullableBool** | Whether to load information about the bank connection owner(s).&lt;br/&gt;This parameter refers to the same parameters of \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/update&#39;&gt;Update a bank connection&lt;/a&gt;\&quot; and \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/connectInterface&#39;&gt;Connect a new interface&lt;/a&gt;\&quot; services in the finAPI Access API. | [optional] [default to false]
**EditSavedSettings** | Pointer to **[]string** | If the user has stored credentials in finAPI for the selected bank connection, then the finAPI web form will automatically use those to login to the bank. If a previous update failed because of invalid credentials or you want to allow the user to change his stored data, you can set this field. It will force the web form flow for the user and allow him to make changes to the data that he has stored in finAPI.&lt;br/&gt;&amp;bull; &lt;code&gt;CREDENTIALS&lt;/code&gt; - the end user can edit his stored credentials. | [optional] 
**Callbacks** | Pointer to [**NullableCallbacks**](Callbacks.md) |  | [optional] 
**ProfileId** | Pointer to **NullableString** | The profile to be applied to the web form.&lt;br/&gt;This will overwrite the default profile, if such a profile exists. | [optional] 

## Methods

### NewBankConnectionUpdateDetails

`func NewBankConnectionUpdateDetails(bank UpdateBankDetails, ) *BankConnectionUpdateDetails`

NewBankConnectionUpdateDetails instantiates a new BankConnectionUpdateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankConnectionUpdateDetailsWithDefaults

`func NewBankConnectionUpdateDetailsWithDefaults() *BankConnectionUpdateDetails`

NewBankConnectionUpdateDetailsWithDefaults instantiates a new BankConnectionUpdateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *BankConnectionUpdateDetails) GetBank() UpdateBankDetails`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *BankConnectionUpdateDetails) GetBankOk() (*UpdateBankDetails, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *BankConnectionUpdateDetails) SetBank(v UpdateBankDetails)`

SetBank sets Bank field to given value.


### GetImportNewAccounts

`func (o *BankConnectionUpdateDetails) GetImportNewAccounts() bool`

GetImportNewAccounts returns the ImportNewAccounts field if non-nil, zero value otherwise.

### GetImportNewAccountsOk

`func (o *BankConnectionUpdateDetails) GetImportNewAccountsOk() (*bool, bool)`

GetImportNewAccountsOk returns a tuple with the ImportNewAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNewAccounts

`func (o *BankConnectionUpdateDetails) SetImportNewAccounts(v bool)`

SetImportNewAccounts sets ImportNewAccounts field to given value.

### HasImportNewAccounts

`func (o *BankConnectionUpdateDetails) HasImportNewAccounts() bool`

HasImportNewAccounts returns a boolean if a field has been set.

### SetImportNewAccountsNil

`func (o *BankConnectionUpdateDetails) SetImportNewAccountsNil(b bool)`

 SetImportNewAccountsNil sets the value for ImportNewAccounts to be an explicit nil

### UnsetImportNewAccounts
`func (o *BankConnectionUpdateDetails) UnsetImportNewAccounts()`

UnsetImportNewAccounts ensures that no value is present for ImportNewAccounts, not even an explicit nil
### GetAccountTypes

`func (o *BankConnectionUpdateDetails) GetAccountTypes() []AccountType`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *BankConnectionUpdateDetails) GetAccountTypesOk() (*[]AccountType, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *BankConnectionUpdateDetails) SetAccountTypes(v []AccountType)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *BankConnectionUpdateDetails) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### SetAccountTypesNil

`func (o *BankConnectionUpdateDetails) SetAccountTypesNil(b bool)`

 SetAccountTypesNil sets the value for AccountTypes to be an explicit nil

### UnsetAccountTypes
`func (o *BankConnectionUpdateDetails) UnsetAccountTypes()`

UnsetAccountTypes ensures that no value is present for AccountTypes, not even an explicit nil
### GetSkipPositionsDownload

`func (o *BankConnectionUpdateDetails) GetSkipPositionsDownload() bool`

GetSkipPositionsDownload returns the SkipPositionsDownload field if non-nil, zero value otherwise.

### GetSkipPositionsDownloadOk

`func (o *BankConnectionUpdateDetails) GetSkipPositionsDownloadOk() (*bool, bool)`

GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPositionsDownload

`func (o *BankConnectionUpdateDetails) SetSkipPositionsDownload(v bool)`

SetSkipPositionsDownload sets SkipPositionsDownload field to given value.

### HasSkipPositionsDownload

`func (o *BankConnectionUpdateDetails) HasSkipPositionsDownload() bool`

HasSkipPositionsDownload returns a boolean if a field has been set.

### SetSkipPositionsDownloadNil

`func (o *BankConnectionUpdateDetails) SetSkipPositionsDownloadNil(b bool)`

 SetSkipPositionsDownloadNil sets the value for SkipPositionsDownload to be an explicit nil

### UnsetSkipPositionsDownload
`func (o *BankConnectionUpdateDetails) UnsetSkipPositionsDownload()`

UnsetSkipPositionsDownload ensures that no value is present for SkipPositionsDownload, not even an explicit nil
### GetLoadOwnerData

`func (o *BankConnectionUpdateDetails) GetLoadOwnerData() bool`

GetLoadOwnerData returns the LoadOwnerData field if non-nil, zero value otherwise.

### GetLoadOwnerDataOk

`func (o *BankConnectionUpdateDetails) GetLoadOwnerDataOk() (*bool, bool)`

GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOwnerData

`func (o *BankConnectionUpdateDetails) SetLoadOwnerData(v bool)`

SetLoadOwnerData sets LoadOwnerData field to given value.

### HasLoadOwnerData

`func (o *BankConnectionUpdateDetails) HasLoadOwnerData() bool`

HasLoadOwnerData returns a boolean if a field has been set.

### SetLoadOwnerDataNil

`func (o *BankConnectionUpdateDetails) SetLoadOwnerDataNil(b bool)`

 SetLoadOwnerDataNil sets the value for LoadOwnerData to be an explicit nil

### UnsetLoadOwnerData
`func (o *BankConnectionUpdateDetails) UnsetLoadOwnerData()`

UnsetLoadOwnerData ensures that no value is present for LoadOwnerData, not even an explicit nil
### GetEditSavedSettings

`func (o *BankConnectionUpdateDetails) GetEditSavedSettings() []string`

GetEditSavedSettings returns the EditSavedSettings field if non-nil, zero value otherwise.

### GetEditSavedSettingsOk

`func (o *BankConnectionUpdateDetails) GetEditSavedSettingsOk() (*[]string, bool)`

GetEditSavedSettingsOk returns a tuple with the EditSavedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditSavedSettings

`func (o *BankConnectionUpdateDetails) SetEditSavedSettings(v []string)`

SetEditSavedSettings sets EditSavedSettings field to given value.

### HasEditSavedSettings

`func (o *BankConnectionUpdateDetails) HasEditSavedSettings() bool`

HasEditSavedSettings returns a boolean if a field has been set.

### SetEditSavedSettingsNil

`func (o *BankConnectionUpdateDetails) SetEditSavedSettingsNil(b bool)`

 SetEditSavedSettingsNil sets the value for EditSavedSettings to be an explicit nil

### UnsetEditSavedSettings
`func (o *BankConnectionUpdateDetails) UnsetEditSavedSettings()`

UnsetEditSavedSettings ensures that no value is present for EditSavedSettings, not even an explicit nil
### GetCallbacks

`func (o *BankConnectionUpdateDetails) GetCallbacks() Callbacks`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *BankConnectionUpdateDetails) GetCallbacksOk() (*Callbacks, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *BankConnectionUpdateDetails) SetCallbacks(v Callbacks)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *BankConnectionUpdateDetails) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### SetCallbacksNil

`func (o *BankConnectionUpdateDetails) SetCallbacksNil(b bool)`

 SetCallbacksNil sets the value for Callbacks to be an explicit nil

### UnsetCallbacks
`func (o *BankConnectionUpdateDetails) UnsetCallbacks()`

UnsetCallbacks ensures that no value is present for Callbacks, not even an explicit nil
### GetProfileId

`func (o *BankConnectionUpdateDetails) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *BankConnectionUpdateDetails) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *BankConnectionUpdateDetails) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *BankConnectionUpdateDetails) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *BankConnectionUpdateDetails) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *BankConnectionUpdateDetails) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


