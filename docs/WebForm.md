# WebForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Globally unique web form&#39;s identifier | 
**Url** | **string** | Full web form&#39;s URL (including the hostname).&lt;br/&gt;You can enhance the given URL with the following query parameters: &lt;code&gt;redirectUrl&lt;/code&gt;, &lt;code&gt;errorRedirectUrl&lt;/code&gt;, &lt;code&gt;customerSupportUrl&lt;/code&gt;.&lt;br/&gt;Find more info in the &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/For-best-results!.2477654019.html#Forbestresults!-Enhanceend-userexperience!&#39; target&#x3D;&#39;_blank&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | 
**CreatedAt** | **string** | The timestamp when the web form was created in the format &lt;code&gt;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt;. | 
**ExpiresAt** | **string** | The timestamp when the web form expires in the format &lt;code&gt;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt;. | 
**Type** | [**WebFormType**](WebFormType.md) |  | 
**Status** | [**WebFormStatus**](WebFormStatus.md) |  | 
**Payload** | [**Payload**](Payload.md) |  | 

## Methods

### NewWebForm

`func NewWebForm(id string, url string, createdAt string, expiresAt string, type_ WebFormType, status WebFormStatus, payload Payload, ) *WebForm`

NewWebForm instantiates a new WebForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebFormWithDefaults

`func NewWebFormWithDefaults() *WebForm`

NewWebFormWithDefaults instantiates a new WebForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebForm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebForm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebForm) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WebForm) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebForm) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebForm) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *WebForm) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebForm) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebForm) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *WebForm) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *WebForm) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *WebForm) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetType

`func (o *WebForm) GetType() WebFormType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebForm) GetTypeOk() (*WebFormType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebForm) SetType(v WebFormType)`

SetType sets Type field to given value.


### GetStatus

`func (o *WebForm) GetStatus() WebFormStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebForm) GetStatusOk() (*WebFormStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebForm) SetStatus(v WebFormStatus)`

SetStatus sets Status field to given value.


### GetPayload

`func (o *WebForm) GetPayload() Payload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebForm) GetPayloadOk() (*Payload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebForm) SetPayload(v Payload)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


