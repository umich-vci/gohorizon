# RCXServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID representing RCX server. | [optional] 
**Name** | Pointer to **string** | FQDN/IP address of the RCX server. | [optional] 
**Port** | Pointer to **int32** | RCX server&#39;s port. | [optional] 
**Status** | Pointer to **string** | This indicates the current status of RCX server. * UP: RCX server is running. * DOWN: RCX server is down. * UNKNOWN: RCX server status is unknown. | [optional] 
**Thumbprints** | Pointer to [**[]CertificateThumbprint**](CertificateThumbprint.md) | Thumbprints of the RCX server certificates. | [optional] 
**Version** | Pointer to **string** | Version information of RCX server. | [optional] 

## Methods

### NewRCXServerInfo

`func NewRCXServerInfo() *RCXServerInfo`

NewRCXServerInfo instantiates a new RCXServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCXServerInfoWithDefaults

`func NewRCXServerInfoWithDefaults() *RCXServerInfo`

NewRCXServerInfoWithDefaults instantiates a new RCXServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RCXServerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RCXServerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RCXServerInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RCXServerInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RCXServerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RCXServerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RCXServerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RCXServerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *RCXServerInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RCXServerInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RCXServerInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RCXServerInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *RCXServerInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RCXServerInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RCXServerInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RCXServerInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThumbprints

`func (o *RCXServerInfo) GetThumbprints() []CertificateThumbprint`

GetThumbprints returns the Thumbprints field if non-nil, zero value otherwise.

### GetThumbprintsOk

`func (o *RCXServerInfo) GetThumbprintsOk() (*[]CertificateThumbprint, bool)`

GetThumbprintsOk returns a tuple with the Thumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprints

`func (o *RCXServerInfo) SetThumbprints(v []CertificateThumbprint)`

SetThumbprints sets Thumbprints field to given value.

### HasThumbprints

`func (o *RCXServerInfo) HasThumbprints() bool`

HasThumbprints returns a boolean if a field has been set.

### GetVersion

`func (o *RCXServerInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RCXServerInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RCXServerInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RCXServerInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


