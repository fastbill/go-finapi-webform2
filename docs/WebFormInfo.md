# WebFormInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Globally unique web form&#39;s identifier | 
**Url** | **string** | Full web form&#39;s URL (including the hostname).&lt;br/&gt;You can enhance the given URL with the following query parameters: &lt;code&gt;redirectUrl&lt;/code&gt;, &lt;code&gt;errorRedirectUrl&lt;/code&gt;, &lt;code&gt;customerSupportUrl&lt;/code&gt;.&lt;br/&gt;Find more info in the &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/For-best-results!.2477654019.html#Forbestresults!-Enhanceend-userexperience!&#39; target&#x3D;&#39;_blank&#39;&gt;Web Form 2.0 Public Documentation&lt;/a&gt;. | 
**Status** | [**WebFormStatus**](WebFormStatus.md) |  | 

## Methods

### NewWebFormInfo

`func NewWebFormInfo(id string, url string, status WebFormStatus, ) *WebFormInfo`

NewWebFormInfo instantiates a new WebFormInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebFormInfoWithDefaults

`func NewWebFormInfoWithDefaults() *WebFormInfo`

NewWebFormInfoWithDefaults instantiates a new WebFormInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebFormInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebFormInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebFormInfo) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WebFormInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebFormInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebFormInfo) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *WebFormInfo) GetStatus() WebFormStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebFormInfo) GetStatusOk() (*WebFormStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebFormInfo) SetStatus(v WebFormStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


