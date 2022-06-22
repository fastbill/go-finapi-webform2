# ImportBankDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Bank ID of the bank the user wants to connect to (and import account data) in the finAPI repo (section \&quot;Banks\&quot;). Mutually exclusive with &lt;code&gt;search&lt;/code&gt; parameter. | [optional] 
**Search** | Pointer to **NullableString** | \&quot;IBAN\&quot;, \&quot;BLZ\&quot;, \&quot;BIC\&quot;, \&quot;Bank Name\&quot;, \&quot;City\&quot; of the bank the user wants to connect to (and import account data). Mutually exclusive with &lt;code&gt;id&lt;/code&gt; parameter. | [optional] 

## Methods

### NewImportBankDetails

`func NewImportBankDetails() *ImportBankDetails`

NewImportBankDetails instantiates a new ImportBankDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportBankDetailsWithDefaults

`func NewImportBankDetailsWithDefaults() *ImportBankDetails`

NewImportBankDetailsWithDefaults instantiates a new ImportBankDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportBankDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportBankDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportBankDetails) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ImportBankDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ImportBankDetails) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ImportBankDetails) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSearch

`func (o *ImportBankDetails) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ImportBankDetails) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ImportBankDetails) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ImportBankDetails) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ImportBankDetails) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ImportBankDetails) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


