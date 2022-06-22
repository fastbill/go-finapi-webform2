# TasksPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Task**](Task.md) | Page of resources | 
**Paging** | [**Paging**](Paging.md) |  | 

## Methods

### NewTasksPage

`func NewTasksPage(items []Task, paging Paging, ) *TasksPage`

NewTasksPage instantiates a new TasksPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksPageWithDefaults

`func NewTasksPageWithDefaults() *TasksPage`

NewTasksPageWithDefaults instantiates a new TasksPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TasksPage) GetItems() []Task`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TasksPage) GetItemsOk() (*[]Task, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TasksPage) SetItems(v []Task)`

SetItems sets Items field to given value.


### GetPaging

`func (o *TasksPage) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *TasksPage) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *TasksPage) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


