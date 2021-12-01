# SAMLAuthenticatorMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | Pointer to [**[]SAMLAuthenticatorMonitorConnectionServer**](SAMLAuthenticatorMonitorConnectionServer.md) | Information about the SAML authenticator connections from each of the connection servers. | [optional] 
**Details** | Pointer to [**SAMLAuthMonitorDetails**](SAMLAuthMonitorDetails.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the SAML Authenticator. | [optional] 

## Methods

### NewSAMLAuthenticatorMonitorInfo

`func NewSAMLAuthenticatorMonitorInfo() *SAMLAuthenticatorMonitorInfo`

NewSAMLAuthenticatorMonitorInfo instantiates a new SAMLAuthenticatorMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLAuthenticatorMonitorInfoWithDefaults

`func NewSAMLAuthenticatorMonitorInfoWithDefaults() *SAMLAuthenticatorMonitorInfo`

NewSAMLAuthenticatorMonitorInfoWithDefaults instantiates a new SAMLAuthenticatorMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *SAMLAuthenticatorMonitorInfo) GetConnectionServers() []SAMLAuthenticatorMonitorConnectionServer`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *SAMLAuthenticatorMonitorInfo) GetConnectionServersOk() (*[]SAMLAuthenticatorMonitorConnectionServer, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *SAMLAuthenticatorMonitorInfo) SetConnectionServers(v []SAMLAuthenticatorMonitorConnectionServer)`

SetConnectionServers sets ConnectionServers field to given value.

### HasConnectionServers

`func (o *SAMLAuthenticatorMonitorInfo) HasConnectionServers() bool`

HasConnectionServers returns a boolean if a field has been set.

### GetDetails

`func (o *SAMLAuthenticatorMonitorInfo) GetDetails() SAMLAuthMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SAMLAuthenticatorMonitorInfo) GetDetailsOk() (*SAMLAuthMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SAMLAuthenticatorMonitorInfo) SetDetails(v SAMLAuthMonitorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SAMLAuthenticatorMonitorInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *SAMLAuthenticatorMonitorInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SAMLAuthenticatorMonitorInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SAMLAuthenticatorMonitorInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SAMLAuthenticatorMonitorInfo) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


