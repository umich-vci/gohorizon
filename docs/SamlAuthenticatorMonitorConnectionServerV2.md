# SAMLAuthenticatorMonitorConnectionServerV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the Connection Server. | [optional] 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp in milliseconds when the last update was obtained. Measured as epoch time. | [optional] 
**Name** | Pointer to **string** | Connection server host name or IP address. | [optional] 
**Status** | Pointer to **string** | Status of the SAML authenticator with respect to this Connection Server. * OK: The connection to SAML authenticator is working properly. * ERROR: Error occurred when connecting to SAML authenticator. * WARN: The connection to SAML authenticator has minor issues. * UNKNOWN: State of SAML authenticator is unknown. | [optional] 
**ThumbprintAccepted** | Pointer to **bool** | Indicates if the thumbprint of the SAML authenticator was accepted. | [optional] 

## Methods

### NewSAMLAuthenticatorMonitorConnectionServerV2

`func NewSAMLAuthenticatorMonitorConnectionServerV2() *SAMLAuthenticatorMonitorConnectionServerV2`

NewSAMLAuthenticatorMonitorConnectionServerV2 instantiates a new SAMLAuthenticatorMonitorConnectionServerV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLAuthenticatorMonitorConnectionServerV2WithDefaults

`func NewSAMLAuthenticatorMonitorConnectionServerV2WithDefaults() *SAMLAuthenticatorMonitorConnectionServerV2`

NewSAMLAuthenticatorMonitorConnectionServerV2WithDefaults instantiates a new SAMLAuthenticatorMonitorConnectionServerV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThumbprintAccepted

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetThumbprintAccepted() bool`

GetThumbprintAccepted returns the ThumbprintAccepted field if non-nil, zero value otherwise.

### GetThumbprintAcceptedOk

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetThumbprintAcceptedOk() (*bool, bool)`

GetThumbprintAcceptedOk returns a tuple with the ThumbprintAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprintAccepted

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetThumbprintAccepted(v bool)`

SetThumbprintAccepted sets ThumbprintAccepted field to given value.

### HasThumbprintAccepted

`func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasThumbprintAccepted() bool`

HasThumbprintAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


