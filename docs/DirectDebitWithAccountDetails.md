# DirectDebitWithAccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionDate** | **string** | Execution date for the direct debit(s), in the format &lt;code&gt;YYYY-MM-DD&lt;/code&gt;. May not be in the past. | 
**BatchBookingPreferred** | Pointer to **NullableBool** | This field is only relevant when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of &lt;code&gt;true&lt;/code&gt;), or as separate bookings (in case of &lt;code&gt;false&lt;/code&gt;). Note that it is subject to the bank whether it will regard the field. | [optional] [default to true]
**ProfileId** | Pointer to **NullableString** | The profile to be applied to the web form.&lt;br/&gt;This will overwrite the default profile, if such a profile exists. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The URL where the end-user will be redirected to after completing the bank login and (possibly) the SCA on the bank&#39;s website. Must always be provided by mandators with &lt;code&gt;FULLY_LICENSED&lt;/code&gt; or &lt;code&gt;PISP&lt;/code&gt; license type, and may not be provided by mandators with other license types. Find more info in the &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://documentation.finapi.io/webform/Licensed-customers-using-the-Web-Form.2832302195.html&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | [optional] 
**Callbacks** | Pointer to [**NullableCallbacks**](Callbacks.md) |  | [optional] 
**Orders** | [**[]DirectDebitOrder**](DirectDebitOrder.md) | Direct debit orders | 
**Receiver** | [**DirectDebitWithAccountReceiver**](DirectDebitWithAccountReceiver.md) |  | 
**DirectDebitType** | **string** | Type of the direct debit; either &lt;code&gt;BASIC&lt;/code&gt; or &lt;code&gt;B2B&lt;/code&gt; (Business-To-Business). | 
**SequenceType** | **string** | Sequence type of the direct debit. Possible values:&lt;br/&gt;&amp;bull; &lt;code&gt;OOFF&lt;/code&gt; - means that this is a one-time direct debit order;&lt;br/&gt;&amp;bull; &lt;code&gt;FRST&lt;/code&gt; - means that this is the first in a row of multiple direct debit orders;&lt;br/&gt;&amp;bull; &lt;code&gt;RCUR&lt;/code&gt; - means that this is one (but not the first or final) within a row of multiple direct debit orders;&lt;br/&gt;&amp;bull; &lt;code&gt;FNAL&lt;/code&gt; - means that this is the final in a row of multiple direct debit orders. | 

## Methods

### NewDirectDebitWithAccountDetails

`func NewDirectDebitWithAccountDetails(executionDate string, orders []DirectDebitOrder, receiver DirectDebitWithAccountReceiver, directDebitType string, sequenceType string, ) *DirectDebitWithAccountDetails`

NewDirectDebitWithAccountDetails instantiates a new DirectDebitWithAccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitWithAccountDetailsWithDefaults

`func NewDirectDebitWithAccountDetailsWithDefaults() *DirectDebitWithAccountDetails`

NewDirectDebitWithAccountDetailsWithDefaults instantiates a new DirectDebitWithAccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionDate

`func (o *DirectDebitWithAccountDetails) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *DirectDebitWithAccountDetails) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *DirectDebitWithAccountDetails) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.


### GetBatchBookingPreferred

`func (o *DirectDebitWithAccountDetails) GetBatchBookingPreferred() bool`

GetBatchBookingPreferred returns the BatchBookingPreferred field if non-nil, zero value otherwise.

### GetBatchBookingPreferredOk

`func (o *DirectDebitWithAccountDetails) GetBatchBookingPreferredOk() (*bool, bool)`

GetBatchBookingPreferredOk returns a tuple with the BatchBookingPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchBookingPreferred

`func (o *DirectDebitWithAccountDetails) SetBatchBookingPreferred(v bool)`

SetBatchBookingPreferred sets BatchBookingPreferred field to given value.

### HasBatchBookingPreferred

`func (o *DirectDebitWithAccountDetails) HasBatchBookingPreferred() bool`

HasBatchBookingPreferred returns a boolean if a field has been set.

### SetBatchBookingPreferredNil

`func (o *DirectDebitWithAccountDetails) SetBatchBookingPreferredNil(b bool)`

 SetBatchBookingPreferredNil sets the value for BatchBookingPreferred to be an explicit nil

### UnsetBatchBookingPreferred
`func (o *DirectDebitWithAccountDetails) UnsetBatchBookingPreferred()`

UnsetBatchBookingPreferred ensures that no value is present for BatchBookingPreferred, not even an explicit nil
### GetProfileId

`func (o *DirectDebitWithAccountDetails) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *DirectDebitWithAccountDetails) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *DirectDebitWithAccountDetails) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *DirectDebitWithAccountDetails) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *DirectDebitWithAccountDetails) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *DirectDebitWithAccountDetails) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetRedirectUrl

`func (o *DirectDebitWithAccountDetails) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *DirectDebitWithAccountDetails) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *DirectDebitWithAccountDetails) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *DirectDebitWithAccountDetails) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *DirectDebitWithAccountDetails) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *DirectDebitWithAccountDetails) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetCallbacks

`func (o *DirectDebitWithAccountDetails) GetCallbacks() Callbacks`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *DirectDebitWithAccountDetails) GetCallbacksOk() (*Callbacks, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *DirectDebitWithAccountDetails) SetCallbacks(v Callbacks)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *DirectDebitWithAccountDetails) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### SetCallbacksNil

`func (o *DirectDebitWithAccountDetails) SetCallbacksNil(b bool)`

 SetCallbacksNil sets the value for Callbacks to be an explicit nil

### UnsetCallbacks
`func (o *DirectDebitWithAccountDetails) UnsetCallbacks()`

UnsetCallbacks ensures that no value is present for Callbacks, not even an explicit nil
### GetOrders

`func (o *DirectDebitWithAccountDetails) GetOrders() []DirectDebitOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *DirectDebitWithAccountDetails) GetOrdersOk() (*[]DirectDebitOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *DirectDebitWithAccountDetails) SetOrders(v []DirectDebitOrder)`

SetOrders sets Orders field to given value.


### GetReceiver

`func (o *DirectDebitWithAccountDetails) GetReceiver() DirectDebitWithAccountReceiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *DirectDebitWithAccountDetails) GetReceiverOk() (*DirectDebitWithAccountReceiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *DirectDebitWithAccountDetails) SetReceiver(v DirectDebitWithAccountReceiver)`

SetReceiver sets Receiver field to given value.


### GetDirectDebitType

`func (o *DirectDebitWithAccountDetails) GetDirectDebitType() string`

GetDirectDebitType returns the DirectDebitType field if non-nil, zero value otherwise.

### GetDirectDebitTypeOk

`func (o *DirectDebitWithAccountDetails) GetDirectDebitTypeOk() (*string, bool)`

GetDirectDebitTypeOk returns a tuple with the DirectDebitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebitType

`func (o *DirectDebitWithAccountDetails) SetDirectDebitType(v string)`

SetDirectDebitType sets DirectDebitType field to given value.


### GetSequenceType

`func (o *DirectDebitWithAccountDetails) GetSequenceType() string`

GetSequenceType returns the SequenceType field if non-nil, zero value otherwise.

### GetSequenceTypeOk

`func (o *DirectDebitWithAccountDetails) GetSequenceTypeOk() (*string, bool)`

GetSequenceTypeOk returns a tuple with the SequenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceType

`func (o *DirectDebitWithAccountDetails) SetSequenceType(v string)`

SetSequenceType sets SequenceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


