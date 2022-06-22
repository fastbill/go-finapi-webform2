# ProfilesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Profile**](Profile.md) | Page of resources | 
**Paging** | [**Paging**](Paging.md) |  | 

## Methods

### NewProfilesPage

`func NewProfilesPage(items []Profile, paging Paging, ) *ProfilesPage`

NewProfilesPage instantiates a new ProfilesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilesPageWithDefaults

`func NewProfilesPageWithDefaults() *ProfilesPage`

NewProfilesPageWithDefaults instantiates a new ProfilesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ProfilesPage) GetItems() []Profile`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProfilesPage) GetItemsOk() (*[]Profile, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProfilesPage) SetItems(v []Profile)`

SetItems sets Items field to given value.


### GetPaging

`func (o *ProfilesPage) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ProfilesPage) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ProfilesPage) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


