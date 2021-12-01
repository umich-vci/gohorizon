# SAMLAuthenticatorMonitorInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | Pointer to [**[]SAMLAuthenticatorMonitorConnectionServerV2**](SAMLAuthenticatorMonitorConnectionServerV2.md) | Information about the SAML authenticator connections from each of the connection servers. | [optional] 
**Details** | Pointer to [**SAMLAuthMonitorDetails**](SAMLAuthMonitorDetails.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the SAML Authenticator. | [optional] 

## Methods

### NewSAMLAuthenticatorMonitorInfoV2

`func NewSAMLAuthenticatorMonitorInfoV2() *SAMLAuthenticatorMonitorInfoV2`

NewSAMLAuthenticatorMonitorInfoV2 instantiates a new SAMLAuthenticatorMonitorInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLAuthenticatorMonitorInfoV2WithDefaults

`func NewSAMLAuthenticatorMonitorInfoV2WithDefaults() *SAMLAuthenticatorMonitorInfoV2`

NewSAMLAuthenticatorMonitorInfoV2WithDefaults instantiates a new SAMLAuthenticatorMonitorInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *SAMLAuthenticatorMonitorInfoV2) GetConnectionServers() []SAMLAuthenticatorMonitorConnectionServerV2`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *SAMLAuthenticatorMonitorInfoV2) GetConnectionServersOk() (*[]SAMLAuthenticatorMonitorConnectionServerV2, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *SAMLAuthenticatorMonitorInfoV2) SetConnectionServers(v []SAMLAuthenticatorMonitorConnectionServerV2)`

SetConnectionServers sets ConnectionServers field to given value.

### HasConnectionServers

`func (o *SAMLAuthenticatorMonitorInfoV2) HasConnectionServers() bool`

HasConnectionServers returns a boolean if a field has been set.

### GetDetails

`func (o *SAMLAuthenticatorMonitorInfoV2) GetDetails() SAMLAuthMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SAMLAuthenticatorMonitorInfoV2) GetDetailsOk() (*SAMLAuthMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SAMLAuthenticatorMonitorInfoV2) SetDetails(v SAMLAuthMonitorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SAMLAuthenticatorMonitorInfoV2) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *SAMLAuthenticatorMonitorInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SAMLAuthenticatorMonitorInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SAMLAuthenticatorMonitorInfoV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SAMLAuthenticatorMonitorInfoV2) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


