# PaymentWithAccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionDate** | Pointer to **NullableString** | Execution date for the money transfer(s), in the format &lt;code&gt;YYYY-MM-DD&lt;/code&gt;. May not be in the past. If not specified, then the current date will be used. | [optional] 
**BatchBookingPreferred** | Pointer to **NullableBool** | This field is only relevant when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of &lt;code&gt;true&lt;/code&gt;), or as separate bookings (in case of &lt;code&gt;false&lt;/code&gt;). Note that it is subject to the bank whether it will regard the field. | [optional] [default to true]
**ProfileId** | Pointer to **NullableString** | The profile to be applied to the web form.&lt;br/&gt;This will overwrite the default profile, if such a profile exists. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The URL where the end-user will be redirected to after completing the bank login and (possibly) the SCA on the bank&#39;s website. Must always be provided by mandators with &lt;code&gt;FULLY_LICENSED&lt;/code&gt; or &lt;code&gt;PISP&lt;/code&gt; license type, and may not be provided by mandators with other license types. Find more info in the &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://documentation.finapi.io/webform/Licensed-customers-using-the-Web-Form.2832302195.html&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | [optional] 
**Callbacks** | Pointer to [**NullableCallbacks**](Callbacks.md) |  | [optional] 
**Sender** | [**PaymentWithAccountSender**](PaymentWithAccountSender.md) |  | 
**Orders** | [**[]PaymentOrder**](PaymentOrder.md) | Payment orders | 

## Methods

### NewPaymentWithAccountDetails

`func NewPaymentWithAccountDetails(sender PaymentWithAccountSender, orders []PaymentOrder, ) *PaymentWithAccountDetails`

NewPaymentWithAccountDetails instantiates a new PaymentWithAccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithAccountDetailsWithDefaults

`func NewPaymentWithAccountDetailsWithDefaults() *PaymentWithAccountDetails`

NewPaymentWithAccountDetailsWithDefaults instantiates a new PaymentWithAccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionDate

`func (o *PaymentWithAccountDetails) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *PaymentWithAccountDetails) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *PaymentWithAccountDetails) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *PaymentWithAccountDetails) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### SetExecutionDateNil

`func (o *PaymentWithAccountDetails) SetExecutionDateNil(b bool)`

 SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil

### UnsetExecutionDate
`func (o *PaymentWithAccountDetails) UnsetExecutionDate()`

UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil
### GetBatchBookingPreferred

`func (o *PaymentWithAccountDetails) GetBatchBookingPreferred() bool`

GetBatchBookingPreferred returns the BatchBookingPreferred field if non-nil, zero value otherwise.

### GetBatchBookingPreferredOk

`func (o *PaymentWithAccountDetails) GetBatchBookingPreferredOk() (*bool, bool)`

GetBatchBookingPreferredOk returns a tuple with the BatchBookingPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchBookingPreferred

`func (o *PaymentWithAccountDetails) SetBatchBookingPreferred(v bool)`

SetBatchBookingPreferred sets BatchBookingPreferred field to given value.

### HasBatchBookingPreferred

`func (o *PaymentWithAccountDetails) HasBatchBookingPreferred() bool`

HasBatchBookingPreferred returns a boolean if a field has been set.

### SetBatchBookingPreferredNil

`func (o *PaymentWithAccountDetails) SetBatchBookingPreferredNil(b bool)`

 SetBatchBookingPreferredNil sets the value for BatchBookingPreferred to be an explicit nil

### UnsetBatchBookingPreferred
`func (o *PaymentWithAccountDetails) UnsetBatchBookingPreferred()`

UnsetBatchBookingPreferred ensures that no value is present for BatchBookingPreferred, not even an explicit nil
### GetProfileId

`func (o *PaymentWithAccountDetails) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *PaymentWithAccountDetails) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *PaymentWithAccountDetails) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *PaymentWithAccountDetails) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *PaymentWithAccountDetails) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *PaymentWithAccountDetails) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetRedirectUrl

`func (o *PaymentWithAccountDetails) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PaymentWithAccountDetails) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PaymentWithAccountDetails) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *PaymentWithAccountDetails) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *PaymentWithAccountDetails) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *PaymentWithAccountDetails) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetCallbacks

`func (o *PaymentWithAccountDetails) GetCallbacks() Callbacks`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *PaymentWithAccountDetails) GetCallbacksOk() (*Callbacks, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *PaymentWithAccountDetails) SetCallbacks(v Callbacks)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *PaymentWithAccountDetails) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### SetCallbacksNil

`func (o *PaymentWithAccountDetails) SetCallbacksNil(b bool)`

 SetCallbacksNil sets the value for Callbacks to be an explicit nil

### UnsetCallbacks
`func (o *PaymentWithAccountDetails) UnsetCallbacks()`

UnsetCallbacks ensures that no value is present for Callbacks, not even an explicit nil
### GetSender

`func (o *PaymentWithAccountDetails) GetSender() PaymentWithAccountSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *PaymentWithAccountDetails) GetSenderOk() (*PaymentWithAccountSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *PaymentWithAccountDetails) SetSender(v PaymentWithAccountSender)`

SetSender sets Sender field to given value.


### GetOrders

`func (o *PaymentWithAccountDetails) GetOrders() []PaymentOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PaymentWithAccountDetails) GetOrdersOk() (*[]PaymentOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PaymentWithAccountDetails) SetOrders(v []PaymentOrder)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


