# SAMLAuthenticatorMonitorConnectionServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the Connection Server. | [optional] 
**Name** | Pointer to **string** | Connection server host name or IP address. | [optional] 
**Status** | Pointer to **string** | Status of the SAML authenticator with respect to this Connection Server. * OK: The connection to SAML authenticator is working properly. * ERROR: Error occurred when connecting to SAML authenticator. * WARN: The connection to SAML authenticator has minor issues. * UNKNOWN: State of SAML authenticator is unknown. | [optional] 
**ThumbprintAccepted** | Pointer to **bool** | Indicates if the thumbprint of the SAML authenticator was accepted. | [optional] 

## Methods

### NewSAMLAuthenticatorMonitorConnectionServer

`func NewSAMLAuthenticatorMonitorConnectionServer() *SAMLAuthenticatorMonitorConnectionServer`

NewSAMLAuthenticatorMonitorConnectionServer instantiates a new SAMLAuthenticatorMonitorConnectionServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLAuthenticatorMonitorConnectionServerWithDefaults

`func NewSAMLAuthenticatorMonitorConnectionServerWithDefaults() *SAMLAuthenticatorMonitorConnectionServer`

NewSAMLAuthenticatorMonitorConnectionServerWithDefaults instantiates a new SAMLAuthenticatorMonitorConnectionServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SAMLAuthenticatorMonitorConnectionServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SAMLAuthenticatorMonitorConnectionServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SAMLAuthenticatorMonitorConnectionServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SAMLAuthenticatorMonitorConnectionServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SAMLAuthenticatorMonitorConnectionServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SAMLAuthenticatorMonitorConnectionServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SAMLAuthenticatorMonitorConnectionServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SAMLAuthenticatorMonitorConnectionServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SAMLAuthenticatorMonitorConnectionServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SAMLAuthenticatorMonitorConnectionServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SAMLAuthenticatorMonitorConnectionServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SAMLAuthenticatorMonitorConnectionServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThumbprintAccepted

`func (o *SAMLAuthenticatorMonitorConnectionServer) GetThumbprintAccepted() bool`

GetThumbprintAccepted returns the ThumbprintAccepted field if non-nil, zero value otherwise.

### GetThumbprintAcceptedOk

`func (o *SAMLAuthenticatorMonitorConnectionServer) GetThumbprintAcceptedOk() (*bool, bool)`

GetThumbprintAcceptedOk returns a tuple with the ThumbprintAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprintAccepted

`func (o *SAMLAuthenticatorMonitorConnectionServer) SetThumbprintAccepted(v bool)`

SetThumbprintAccepted sets ThumbprintAccepted field to given value.

### HasThumbprintAccepted

`func (o *SAMLAuthenticatorMonitorConnectionServer) HasThumbprintAccepted() bool`

HasThumbprintAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


