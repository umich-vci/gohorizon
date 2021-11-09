# ClientData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | IP address of the client machine for the session.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**LocationId** | Pointer to **string** | Client location for the session.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**Name** | Pointer to **string** | Client machine hostname for the session.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**Type** | Pointer to **string** | Client type for the session.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * WINDOWS: Client type is Windows client. * MAC: Client type is Mac client. * HTMLACCESS: Client type is Web client. * LINUX: Client type is Linux client. * IOS: Client type is iOS client. * ANDROID: Client type is Android client. * OTHER: Client type is other. | [optional] 
**Version** | Pointer to **string** | Client version for the session.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 

## Methods

### NewClientData

`func NewClientData() *ClientData`

NewClientData instantiates a new ClientData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDataWithDefaults

`func NewClientDataWithDefaults() *ClientData`

NewClientDataWithDefaults instantiates a new ClientData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ClientData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClientData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClientData) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ClientData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLocationId

`func (o *ClientData) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ClientData) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ClientData) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ClientData) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetName

`func (o *ClientData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ClientData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *ClientData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClientData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClientData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClientData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


