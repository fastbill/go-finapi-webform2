# BankConnectionImportDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to [**NullableImportBankDetails**](ImportBankDetails.md) |  | [optional] 
**BankConnectionName** | Pointer to **NullableString** | Custom name for the bank connection | [optional] 
**SkipPositionsDownload** | Pointer to **NullableBool** | Whether to skip the download of transactions and securities or not.&lt;br/&gt;This parameter refers to the same parameter of the \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/import&#39;&gt;Import a new bank connection&lt;/a&gt;\&quot; service in the finAPI Access API. | [optional] [default to false]
**LoadOwnerData** | Pointer to **NullableBool** | Whether to load information about the bank connection owner(s).&lt;br/&gt;This parameter refers to the same parameter of the \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/import&#39;&gt;Import a new bank connection&lt;/a&gt;\&quot; service in the finAPI Access API. | [optional] [default to false]
**MaxDaysForDownload** | Pointer to **NullableInt32** | Defines the limit of how many days of the transaction history will be requested from the bank.&lt;br/&gt;This parameter refers to the same parameter of the \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/import&#39;&gt;Import a new bank connection&lt;/a&gt;\&quot; service in the finAPI Access API. | [optional] [default to 0]
**AccountTypes** | Pointer to [**[]AccountType**](AccountType.md) | A set of account types that are considered for the import. If no values is given, then all accounts will be imported.&lt;br/&gt;This parameter refers to the same parameter of the \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#post-/api/v1/bankConnections/import&#39;&gt;Import a new bank connection&lt;/a&gt;\&quot; service in the finAPI Access API. | [optional] 
**Callbacks** | Pointer to [**NullableCallbacks**](Callbacks.md) |  | [optional] 
**ProfileId** | Pointer to **NullableString** | The profile to be applied to the web form.&lt;br/&gt;This will overwrite the default profile, if such a profile exists. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The URL where the end-user will be redirected to after completing the bank login and (possibly) the SCA on the bank&#39;s website. Must always be provided by mandators with &lt;code&gt;FULLY_LICENSED&lt;/code&gt; or &lt;code&gt;AISP&lt;/code&gt; license type, and may not be provided by mandators with other license types. Find more info in the &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://documentation.finapi.io/webform/Licensed-customers-using-the-Web-Form.2832302195.html&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | [optional] 

## Methods

### NewBankConnectionImportDetails

`func NewBankConnectionImportDetails() *BankConnectionImportDetails`

NewBankConnectionImportDetails instantiates a new BankConnectionImportDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankConnectionImportDetailsWithDefaults

`func NewBankConnectionImportDetailsWithDefaults() *BankConnectionImportDetails`

NewBankConnectionImportDetailsWithDefaults instantiates a new BankConnectionImportDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *BankConnectionImportDetails) GetBank() ImportBankDetails`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *BankConnectionImportDetails) GetBankOk() (*ImportBankDetails, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *BankConnectionImportDetails) SetBank(v ImportBankDetails)`

SetBank sets Bank field to given value.

### HasBank

`func (o *BankConnectionImportDetails) HasBank() bool`

HasBank returns a boolean if a field has been set.

### SetBankNil

`func (o *BankConnectionImportDetails) SetBankNil(b bool)`

 SetBankNil sets the value for Bank to be an explicit nil

### UnsetBank
`func (o *BankConnectionImportDetails) UnsetBank()`

UnsetBank ensures that no value is present for Bank, not even an explicit nil
### GetBankConnectionName

`func (o *BankConnectionImportDetails) GetBankConnectionName() string`

GetBankConnectionName returns the BankConnectionName field if non-nil, zero value otherwise.

### GetBankConnectionNameOk

`func (o *BankConnectionImportDetails) GetBankConnectionNameOk() (*string, bool)`

GetBankConnectionNameOk returns a tuple with the BankConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionName

`func (o *BankConnectionImportDetails) SetBankConnectionName(v string)`

SetBankConnectionName sets BankConnectionName field to given value.

### HasBankConnectionName

`func (o *BankConnectionImportDetails) HasBankConnectionName() bool`

HasBankConnectionName returns a boolean if a field has been set.

### SetBankConnectionNameNil

`func (o *BankConnectionImportDetails) SetBankConnectionNameNil(b bool)`

 SetBankConnectionNameNil sets the value for BankConnectionName to be an explicit nil

### UnsetBankConnectionName
`func (o *BankConnectionImportDetails) UnsetBankConnectionName()`

UnsetBankConnectionName ensures that no value is present for BankConnectionName, not even an explicit nil
### GetSkipPositionsDownload

`func (o *BankConnectionImportDetails) GetSkipPositionsDownload() bool`

GetSkipPositionsDownload returns the SkipPositionsDownload field if non-nil, zero value otherwise.

### GetSkipPositionsDownloadOk

`func (o *BankConnectionImportDetails) GetSkipPositionsDownloadOk() (*bool, bool)`

GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPositionsDownload

`func (o *BankConnectionImportDetails) SetSkipPositionsDownload(v bool)`

SetSkipPositionsDownload sets SkipPositionsDownload field to given value.

### HasSkipPositionsDownload

`func (o *BankConnectionImportDetails) HasSkipPositionsDownload() bool`

HasSkipPositionsDownload returns a boolean if a field has been set.

### SetSkipPositionsDownloadNil

`func (o *BankConnectionImportDetails) SetSkipPositionsDownloadNil(b bool)`

 SetSkipPositionsDownloadNil sets the value for SkipPositionsDownload to be an explicit nil

### UnsetSkipPositionsDownload
`func (o *BankConnectionImportDetails) UnsetSkipPositionsDownload()`

UnsetSkipPositionsDownload ensures that no value is present for SkipPositionsDownload, not even an explicit nil
### GetLoadOwnerData

`func (o *BankConnectionImportDetails) GetLoadOwnerData() bool`

GetLoadOwnerData returns the LoadOwnerData field if non-nil, zero value otherwise.

### GetLoadOwnerDataOk

`func (o *BankConnectionImportDetails) GetLoadOwnerDataOk() (*bool, bool)`

GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOwnerData

`func (o *BankConnectionImportDetails) SetLoadOwnerData(v bool)`

SetLoadOwnerData sets LoadOwnerData field to given value.

### HasLoadOwnerData

`func (o *BankConnectionImportDetails) HasLoadOwnerData() bool`

HasLoadOwnerData returns a boolean if a field has been set.

### SetLoadOwnerDataNil

`func (o *BankConnectionImportDetails) SetLoadOwnerDataNil(b bool)`

 SetLoadOwnerDataNil sets the value for LoadOwnerData to be an explicit nil

### UnsetLoadOwnerData
`func (o *BankConnectionImportDetails) UnsetLoadOwnerData()`

UnsetLoadOwnerData ensures that no value is present for LoadOwnerData, not even an explicit nil
### GetMaxDaysForDownload

`func (o *BankConnectionImportDetails) GetMaxDaysForDownload() int32`

GetMaxDaysForDownload returns the MaxDaysForDownload field if non-nil, zero value otherwise.

### GetMaxDaysForDownloadOk

`func (o *BankConnectionImportDetails) GetMaxDaysForDownloadOk() (*int32, bool)`

GetMaxDaysForDownloadOk returns a tuple with the MaxDaysForDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDaysForDownload

`func (o *BankConnectionImportDetails) SetMaxDaysForDownload(v int32)`

SetMaxDaysForDownload sets MaxDaysForDownload field to given value.

### HasMaxDaysForDownload

`func (o *BankConnectionImportDetails) HasMaxDaysForDownload() bool`

HasMaxDaysForDownload returns a boolean if a field has been set.

### SetMaxDaysForDownloadNil

`func (o *BankConnectionImportDetails) SetMaxDaysForDownloadNil(b bool)`

 SetMaxDaysForDownloadNil sets the value for MaxDaysForDownload to be an explicit nil

### UnsetMaxDaysForDownload
`func (o *BankConnectionImportDetails) UnsetMaxDaysForDownload()`

UnsetMaxDaysForDownload ensures that no value is present for MaxDaysForDownload, not even an explicit nil
### GetAccountTypes

`func (o *BankConnectionImportDetails) GetAccountTypes() []AccountType`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *BankConnectionImportDetails) GetAccountTypesOk() (*[]AccountType, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *BankConnectionImportDetails) SetAccountTypes(v []AccountType)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *BankConnectionImportDetails) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### SetAccountTypesNil

`func (o *BankConnectionImportDetails) SetAccountTypesNil(b bool)`

 SetAccountTypesNil sets the value for AccountTypes to be an explicit nil

### UnsetAccountTypes
`func (o *BankConnectionImportDetails) UnsetAccountTypes()`

UnsetAccountTypes ensures that no value is present for AccountTypes, not even an explicit nil
### GetCallbacks

`func (o *BankConnectionImportDetails) GetCallbacks() Callbacks`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *BankConnectionImportDetails) GetCallbacksOk() (*Callbacks, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *BankConnectionImportDetails) SetCallbacks(v Callbacks)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *BankConnectionImportDetails) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### SetCallbacksNil

`func (o *BankConnectionImportDetails) SetCallbacksNil(b bool)`

 SetCallbacksNil sets the value for Callbacks to be an explicit nil

### UnsetCallbacks
`func (o *BankConnectionImportDetails) UnsetCallbacks()`

UnsetCallbacks ensures that no value is present for Callbacks, not even an explicit nil
### GetProfileId

`func (o *BankConnectionImportDetails) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *BankConnectionImportDetails) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *BankConnectionImportDetails) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *BankConnectionImportDetails) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *BankConnectionImportDetails) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *BankConnectionImportDetails) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetRedirectUrl

`func (o *BankConnectionImportDetails) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *BankConnectionImportDetails) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *BankConnectionImportDetails) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *BankConnectionImportDetails) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *BankConnectionImportDetails) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *BankConnectionImportDetails) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


