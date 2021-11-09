# RestrictedClientData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of Horizon Client. * WINDOWS: The client is the Horizon Client for Windows. * MAC: The client is the Horizon Client for MacOS. * HTMLACCESS: The client is a Web client. * LINUX: The client is the Horizon Client for Linux. * IOS: The client is the Horizon Client for IOS devices. * ANDROID: The client is the Horizon Client for Android. * WINSTORE: The client is the Horizon Client for Windows 10 UWP. * CHROME: The client is the Horizon Client for Chrome Native OS. | [optional] 
**Version** | Pointer to **string** | The version of Horizon Client. | [optional] 

## Methods

### NewRestrictedClientData

`func NewRestrictedClientData() *RestrictedClientData`

NewRestrictedClientData instantiates a new RestrictedClientData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedClientDataWithDefaults

`func NewRestrictedClientDataWithDefaults() *RestrictedClientData`

NewRestrictedClientDataWithDefaults instantiates a new RestrictedClientData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RestrictedClientData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestrictedClientData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestrictedClientData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RestrictedClientData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *RestrictedClientData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RestrictedClientData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RestrictedClientData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RestrictedClientData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


