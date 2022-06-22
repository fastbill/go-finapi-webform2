# UserMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | The IP address of the user&#39;s device | 
**DeviceOs** | **string** | The user&#39;s device and/or operating system identification | 
**UserAgent** | **string** | The user&#39;s web browser or other client device identification | 

## Methods

### NewUserMetadata

`func NewUserMetadata(ipAddress string, deviceOs string, userAgent string, ) *UserMetadata`

NewUserMetadata instantiates a new UserMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMetadataWithDefaults

`func NewUserMetadataWithDefaults() *UserMetadata`

NewUserMetadataWithDefaults instantiates a new UserMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *UserMetadata) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UserMetadata) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UserMetadata) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetDeviceOs

`func (o *UserMetadata) GetDeviceOs() string`

GetDeviceOs returns the DeviceOs field if non-nil, zero value otherwise.

### GetDeviceOsOk

`func (o *UserMetadata) GetDeviceOsOk() (*string, bool)`

GetDeviceOsOk returns a tuple with the DeviceOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOs

`func (o *UserMetadata) SetDeviceOs(v string)`

SetDeviceOs sets DeviceOs field to given value.


### GetUserAgent

`func (o *UserMetadata) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *UserMetadata) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *UserMetadata) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


