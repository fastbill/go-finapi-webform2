# StandalonePaymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | [**[]PaymentOrder**](PaymentOrder.md) | Payment orders | 
**ExecutionDate** | Pointer to **NullableString** | Execution date for the money transfer(s), in the format &lt;code&gt;YYYY-MM-DD&lt;/code&gt;. May not be in the past. If not specified, then the current date will be used. | [optional] 
**BatchBookingPreferred** | Pointer to **NullableBool** | This field is only relevant when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of &lt;code&gt;true&lt;/code&gt;), or as separate bookings (in case of &lt;code&gt;false&lt;/code&gt;). Note that it is subject to the bank whether it will regard the field. | [optional] [default to true]
**InstantPayment** | Pointer to **bool** | Whether the order should be submitted to the bank as an instant SEPA order. &lt;br/&gt;&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt;&lt;br/&gt;&amp;bull; Submitting an instant payment will work only with interfaces that support it.&lt;br/&gt;&amp;bull; Instant payments work only for a single order, not for collective orders.&lt;br/&gt;&amp;bull; The bank may charge a fee for instant payments, depending on the agreement between the user and the bank.&lt;br/&gt;&amp;bull; The payment might get rejected if the source and/or target account doesn&#39;t support instant payments.&lt;br/&gt;&lt;br/&gt;If you are interested in additional banks where we can support Instant Payments, please contact our Sales or support team | [optional] [default to false]
**ProfileId** | Pointer to **NullableString** | The profile to be applied to the web form.&lt;br/&gt;This will overwrite the default profile, if such a profile exists. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The URL where the end-user will be redirected to after completing the bank login and (possibly) the SCA on the bank&#39;s website. Must always be provided by mandators with &lt;code&gt;FULLY_LICENSED&lt;/code&gt; or &lt;code&gt;PISP&lt;/code&gt; license type, and may not be provided by mandators with other license types. Find more info in the &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://documentation.finapi.io/webform/Licensed-customers-using-the-Web-Form.2832302195.html&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | [optional] 
**Callbacks** | Pointer to [**NullableCallbacks**](Callbacks.md) |  | [optional] 
**Sender** | Pointer to [**NullableSender**](Sender.md) |  | [optional] 

## Methods

### NewStandalonePaymentDetails

`func NewStandalonePaymentDetails(orders []PaymentOrder, ) *StandalonePaymentDetails`

NewStandalonePaymentDetails instantiates a new StandalonePaymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandalonePaymentDetailsWithDefaults

`func NewStandalonePaymentDetailsWithDefaults() *StandalonePaymentDetails`

NewStandalonePaymentDetailsWithDefaults instantiates a new StandalonePaymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *StandalonePaymentDetails) GetOrders() []PaymentOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *StandalonePaymentDetails) GetOrdersOk() (*[]PaymentOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *StandalonePaymentDetails) SetOrders(v []PaymentOrder)`

SetOrders sets Orders field to given value.


### GetExecutionDate

`func (o *StandalonePaymentDetails) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *StandalonePaymentDetails) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *StandalonePaymentDetails) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *StandalonePaymentDetails) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### SetExecutionDateNil

`func (o *StandalonePaymentDetails) SetExecutionDateNil(b bool)`

 SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil

### UnsetExecutionDate
`func (o *StandalonePaymentDetails) UnsetExecutionDate()`

UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil
### GetBatchBookingPreferred

`func (o *StandalonePaymentDetails) GetBatchBookingPreferred() bool`

GetBatchBookingPreferred returns the BatchBookingPreferred field if non-nil, zero value otherwise.

### GetBatchBookingPreferredOk

`func (o *StandalonePaymentDetails) GetBatchBookingPreferredOk() (*bool, bool)`

GetBatchBookingPreferredOk returns a tuple with the BatchBookingPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchBookingPreferred

`func (o *StandalonePaymentDetails) SetBatchBookingPreferred(v bool)`

SetBatchBookingPreferred sets BatchBookingPreferred field to given value.

### HasBatchBookingPreferred

`func (o *StandalonePaymentDetails) HasBatchBookingPreferred() bool`

HasBatchBookingPreferred returns a boolean if a field has been set.

### SetBatchBookingPreferredNil

`func (o *StandalonePaymentDetails) SetBatchBookingPreferredNil(b bool)`

 SetBatchBookingPreferredNil sets the value for BatchBookingPreferred to be an explicit nil

### UnsetBatchBookingPreferred
`func (o *StandalonePaymentDetails) UnsetBatchBookingPreferred()`

UnsetBatchBookingPreferred ensures that no value is present for BatchBookingPreferred, not even an explicit nil
### GetInstantPayment

`func (o *StandalonePaymentDetails) GetInstantPayment() bool`

GetInstantPayment returns the InstantPayment field if non-nil, zero value otherwise.

### GetInstantPaymentOk

`func (o *StandalonePaymentDetails) GetInstantPaymentOk() (*bool, bool)`

GetInstantPaymentOk returns a tuple with the InstantPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantPayment

`func (o *StandalonePaymentDetails) SetInstantPayment(v bool)`

SetInstantPayment sets InstantPayment field to given value.

### HasInstantPayment

`func (o *StandalonePaymentDetails) HasInstantPayment() bool`

HasInstantPayment returns a boolean if a field has been set.

### GetProfileId

`func (o *StandalonePaymentDetails) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *StandalonePaymentDetails) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *StandalonePaymentDetails) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *StandalonePaymentDetails) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *StandalonePaymentDetails) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *StandalonePaymentDetails) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetRedirectUrl

`func (o *StandalonePaymentDetails) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *StandalonePaymentDetails) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *StandalonePaymentDetails) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *StandalonePaymentDetails) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *StandalonePaymentDetails) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *StandalonePaymentDetails) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetCallbacks

`func (o *StandalonePaymentDetails) GetCallbacks() Callbacks`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *StandalonePaymentDetails) GetCallbacksOk() (*Callbacks, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *StandalonePaymentDetails) SetCallbacks(v Callbacks)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *StandalonePaymentDetails) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### SetCallbacksNil

`func (o *StandalonePaymentDetails) SetCallbacksNil(b bool)`

 SetCallbacksNil sets the value for Callbacks to be an explicit nil

### UnsetCallbacks
`func (o *StandalonePaymentDetails) UnsetCallbacks()`

UnsetCallbacks ensures that no value is present for Callbacks, not even an explicit nil
### GetSender

`func (o *StandalonePaymentDetails) GetSender() Sender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *StandalonePaymentDetails) GetSenderOk() (*Sender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *StandalonePaymentDetails) SetSender(v Sender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *StandalonePaymentDetails) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *StandalonePaymentDetails) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *StandalonePaymentDetails) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


