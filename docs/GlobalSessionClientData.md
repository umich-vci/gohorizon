# GlobalSessionClientData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | IP Address of the client machine for the session. | [optional] 
**LocationId** | Pointer to **string** | Client location for the session. | [optional] 
**Name** | Pointer to **string** | Client machine name for the session. | [optional] 
**Type** | Pointer to **string** | Client type for the session. * WINDOWS: Client type is Windows client. * MAC: Client type is Mac client. * HTMLACCESS: Client type is Web client. * LINUX: Client type is Linux client. * IOS: Client type is iOS client. * ANDROID: Client type is Android client. * OTHER: Client type is other. | [optional] 
**Version** | Pointer to **string** | Client version for the session. | [optional] 

## Methods

### NewGlobalSessionClientData

`func NewGlobalSessionClientData() *GlobalSessionClientData`

NewGlobalSessionClientData instantiates a new GlobalSessionClientData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSessionClientDataWithDefaults

`func NewGlobalSessionClientDataWithDefaults() *GlobalSessionClientData`

NewGlobalSessionClientDataWithDefaults instantiates a new GlobalSessionClientData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GlobalSessionClientData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GlobalSessionClientData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GlobalSessionClientData) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GlobalSessionClientData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLocationId

`func (o *GlobalSessionClientData) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GlobalSessionClientData) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GlobalSessionClientData) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GlobalSessionClientData) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetName

`func (o *GlobalSessionClientData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalSessionClientData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalSessionClientData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalSessionClientData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GlobalSessionClientData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalSessionClientData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalSessionClientData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GlobalSessionClientData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *GlobalSessionClientData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GlobalSessionClientData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GlobalSessionClientData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GlobalSessionClientData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


