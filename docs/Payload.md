# Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionId** | Pointer to **NullableInt64** | Identifier of the bank connection in the Access API. Initialized as soon as the Web Form reaches a final status and a bank connection exists in Access. Use this ID to gather Bank Connection data from Access endpoints like, \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#get-/api/v1/bankConnections/-id-&#39;&gt;Get a bank connection&lt;/a&gt;\&quot; or \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#get-/api/v1/accounts&#39;&gt;Get and search all accounts&lt;/a&gt;\&quot;.&lt;br/&gt;This field is mutually exclusive with &lt;code&gt;paymentId&lt;/code&gt; and &lt;code&gt;standingOrderId&lt;/code&gt;. | [optional] 
**PaymentId** | Pointer to **NullableInt64** | Identifier of the payment in the Access API. Initialized as soon as the Web Form reaches a final status and a payment exists in Access. Use this ID to get Payment initialization data from the Access endpoint, \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#get-/api/v1/payments&#39;&gt;Get payments&lt;/a&gt;\&quot;.&lt;br/&gt;This field is mutually exclusive with &lt;code&gt;bankConnectionId&lt;/code&gt; and &lt;code&gt;standingOrderId&lt;/code&gt;. | [optional] 
**StandingOrderId** | Pointer to **NullableInt64** | Identifier of the standing order in the Access API. Initialized as soon as the Web Form reaches a final status and a standing order exists in Access. Use this ID to get Standing Order initialization data from the Access endpoint, \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#get-/api/v1/standingOrders&#39;&gt;Get standing orders&lt;/a&gt;\&quot;.&lt;br/&gt;This field is mutually exclusive with &lt;code&gt;bankConnectionId&lt;/code&gt; and &lt;code&gt;paymentId&lt;/code&gt;. | [optional] 
**ErrorCode** | Pointer to **NullableString** | Reason of the web form failure.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; This enum can be extended in the future as new cases arise!&lt;br/&gt;&lt;br/&gt;Codes can be interpreted as follows:&lt;br/&gt;&amp;bull; &lt;code&gt;ENTITY_EXISTS&lt;/code&gt; - Access API rejected the import because of detected bank connection duplication;&lt;br/&gt;&amp;bull; &lt;code&gt;BANK_SERVER_REJECTION&lt;/code&gt; - the flow has been terminated on the bank side, e.g., in case of incorrect credentials;&lt;br/&gt;&amp;bull; &lt;code&gt;INTERRUPTED&lt;/code&gt; - web form has been reloaded or re-opened on a step where it&#39;s not supported;&lt;br/&gt;&amp;bull; &lt;code&gt;INVALID_TOKEN&lt;/code&gt; - the given access token expired or became invalid during the flow; &lt;br/&gt;&amp;bull; &lt;code&gt;MANDATOR_MISCONFIGURATION&lt;/code&gt; - the mandator is not properly configured on the Access side, e.g., it is still configured to use the old web form, or the scraper interface has been selected by the Web Form 2.0 but using it is not allowed on the mandator level. Please contact our support team (&lt;a href&#x3D;&#39;mailto:support@finapi.io&#39;&gt;support@finapi.io&lt;/a&gt;) for troubleshooting;&lt;br/&gt;&amp;bull; &lt;code&gt;NO_ACCOUNTS_FOR_TYPE_LIST&lt;/code&gt; - in the end of the web form flow there were no accounts of type requested in the API call; &lt;br/&gt;&amp;bull; &lt;code&gt;UNDETERMINED_BANK&lt;/code&gt; - the given search criteria resulted in either zero or multiple bank entries;&lt;br/&gt;&amp;bull; &lt;code&gt;UNEXPECTED_ACCESS_RESPONSE&lt;/code&gt; - an unexpected response has been received from the Access API - similarly to the &lt;code&gt;INTERNAL_ERROR&lt;/code&gt; code, please forward all details to our Customer Support team; &lt;br/&gt;&amp;bull; &lt;code&gt;UNSUPPORTED_FEATURE&lt;/code&gt; - Access API rejected the request because the requested feature is not supported, e.g., a payment with the execution date in the future was requested for a bank that does not support it;&lt;br/&gt;&amp;bull; &lt;code&gt;UNSUPPORTED_ORDER&lt;/code&gt; - Access API rejected the payment request because the associated account does not have the required capabilities;&lt;br/&gt;&amp;bull; &lt;code&gt;INTERNAL_ERROR&lt;/code&gt; - the reason of the failure can not be identified - please forward all the details to our Customer Support team in order to get more info and also help us to eliminate the issue. | [optional] 

## Methods

### NewPayload

`func NewPayload() *Payload`

NewPayload instantiates a new Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadWithDefaults

`func NewPayloadWithDefaults() *Payload`

NewPayloadWithDefaults instantiates a new Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankConnectionId

`func (o *Payload) GetBankConnectionId() int64`

GetBankConnectionId returns the BankConnectionId field if non-nil, zero value otherwise.

### GetBankConnectionIdOk

`func (o *Payload) GetBankConnectionIdOk() (*int64, bool)`

GetBankConnectionIdOk returns a tuple with the BankConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionId

`func (o *Payload) SetBankConnectionId(v int64)`

SetBankConnectionId sets BankConnectionId field to given value.

### HasBankConnectionId

`func (o *Payload) HasBankConnectionId() bool`

HasBankConnectionId returns a boolean if a field has been set.

### SetBankConnectionIdNil

`func (o *Payload) SetBankConnectionIdNil(b bool)`

 SetBankConnectionIdNil sets the value for BankConnectionId to be an explicit nil

### UnsetBankConnectionId
`func (o *Payload) UnsetBankConnectionId()`

UnsetBankConnectionId ensures that no value is present for BankConnectionId, not even an explicit nil
### GetPaymentId

`func (o *Payload) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *Payload) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *Payload) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *Payload) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *Payload) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *Payload) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetStandingOrderId

`func (o *Payload) GetStandingOrderId() int64`

GetStandingOrderId returns the StandingOrderId field if non-nil, zero value otherwise.

### GetStandingOrderIdOk

`func (o *Payload) GetStandingOrderIdOk() (*int64, bool)`

GetStandingOrderIdOk returns a tuple with the StandingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandingOrderId

`func (o *Payload) SetStandingOrderId(v int64)`

SetStandingOrderId sets StandingOrderId field to given value.

### HasStandingOrderId

`func (o *Payload) HasStandingOrderId() bool`

HasStandingOrderId returns a boolean if a field has been set.

### SetStandingOrderIdNil

`func (o *Payload) SetStandingOrderIdNil(b bool)`

 SetStandingOrderIdNil sets the value for StandingOrderId to be an explicit nil

### UnsetStandingOrderId
`func (o *Payload) UnsetStandingOrderId()`

UnsetStandingOrderId ensures that no value is present for StandingOrderId, not even an explicit nil
### GetErrorCode

`func (o *Payload) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Payload) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Payload) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Payload) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *Payload) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *Payload) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


