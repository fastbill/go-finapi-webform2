# TaskPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionId** | **int64** | Identifier of the bank connection in the Access API. Initialized as soon as the task process is started. Use those ID to gather Bank Connection data from Access endpoints like, \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#get-/api/v1/bankConnections/-id-&#39;&gt;Get a bank connection&lt;/a&gt;\&quot; or \&quot;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.finapi.io/?product&#x3D;access#get-/api/v1/accounts&#39;&gt;Get and search all accounts&lt;/a&gt;\&quot;. | 
**WebForm** | Pointer to [**NullableWebFormInfo**](WebFormInfo.md) |  | [optional] 
**ErrorCode** | Pointer to **NullableString** | Reason of the task failure.&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; This enum can be extended in the future as new cases arise!&lt;br/&gt;&lt;br/&gt;Codes can be interpreted as follows:&lt;br/&gt;&amp;bull; &lt;code&gt;BANK_SERVER_REJECTION&lt;/code&gt; - the flow has been terminated on the bank side, e.g., in case of incorrect credentials;&lt;br/&gt;&amp;bull; &lt;code&gt;INVALID_TOKEN&lt;/code&gt; - the given access token expired or became invalid during the flow; &lt;br/&gt;&amp;bull; &lt;code&gt;UNEXPECTED_ACCESS_RESPONSE&lt;/code&gt; - an unexpected response has been received from the Access API - similarly to the &lt;code&gt;INTERNAL_ERROR&lt;/code&gt; code, please forward all details to our Customer Support team; &lt;br/&gt;&amp;bull; &lt;code&gt;INTERNAL_ERROR&lt;/code&gt; - the reason of the failure can not be identified - please forward all the details to our Customer Support team in order to get more info and also help us to eliminate the issue. | [optional] 

## Methods

### NewTaskPayload

`func NewTaskPayload(bankConnectionId int64, ) *TaskPayload`

NewTaskPayload instantiates a new TaskPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPayloadWithDefaults

`func NewTaskPayloadWithDefaults() *TaskPayload`

NewTaskPayloadWithDefaults instantiates a new TaskPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankConnectionId

`func (o *TaskPayload) GetBankConnectionId() int64`

GetBankConnectionId returns the BankConnectionId field if non-nil, zero value otherwise.

### GetBankConnectionIdOk

`func (o *TaskPayload) GetBankConnectionIdOk() (*int64, bool)`

GetBankConnectionIdOk returns a tuple with the BankConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionId

`func (o *TaskPayload) SetBankConnectionId(v int64)`

SetBankConnectionId sets BankConnectionId field to given value.


### GetWebForm

`func (o *TaskPayload) GetWebForm() WebFormInfo`

GetWebForm returns the WebForm field if non-nil, zero value otherwise.

### GetWebFormOk

`func (o *TaskPayload) GetWebFormOk() (*WebFormInfo, bool)`

GetWebFormOk returns a tuple with the WebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebForm

`func (o *TaskPayload) SetWebForm(v WebFormInfo)`

SetWebForm sets WebForm field to given value.

### HasWebForm

`func (o *TaskPayload) HasWebForm() bool`

HasWebForm returns a boolean if a field has been set.

### SetWebFormNil

`func (o *TaskPayload) SetWebFormNil(b bool)`

 SetWebFormNil sets the value for WebForm to be an explicit nil

### UnsetWebForm
`func (o *TaskPayload) UnsetWebForm()`

UnsetWebForm ensures that no value is present for WebForm, not even an explicit nil
### GetErrorCode

`func (o *TaskPayload) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TaskPayload) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TaskPayload) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TaskPayload) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *TaskPayload) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *TaskPayload) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


