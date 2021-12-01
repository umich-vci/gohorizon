# VCMonitorConnectionServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**CertificateMonitorInfo**](CertificateMonitorInfo.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the Connection Server. | [optional] 
**Name** | Pointer to **string** | Connection server host name or IP address. | [optional] 
**Status** | Pointer to **string** | Status of the Virtual Center Connection with respect to this Connection Server. * OK: The connection to Virtual Center server is working properly. * DOWN: The connection to Virtual Center server is down. * RECONNECTING: The connection to Virtual Center server was lost and is being reconnected to. * UNKNOWN: Connection state to Virtual Center server is unknown. * INVALID_CREDENTIALS: The supplied credentials cannot be used to authenticate to the Virtual Center server. * CANNOT_LOGIN: The connection server cannot login to the Virtual Center server. * NOT_YET_CONNECTED: Connection server has not yet connected to Virtual Center server. | [optional] 
**ThumbprintAccepted** | Pointer to **bool** | Indicates if the thumbprints of the Virtual Center was accepted. | [optional] 

## Methods

### NewVCMonitorConnectionServer

`func NewVCMonitorConnectionServer() *VCMonitorConnectionServer`

NewVCMonitorConnectionServer instantiates a new VCMonitorConnectionServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCMonitorConnectionServerWithDefaults

`func NewVCMonitorConnectionServerWithDefaults() *VCMonitorConnectionServer`

NewVCMonitorConnectionServerWithDefaults instantiates a new VCMonitorConnectionServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *VCMonitorConnectionServer) GetCertificate() CertificateMonitorInfo`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *VCMonitorConnectionServer) GetCertificateOk() (*CertificateMonitorInfo, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *VCMonitorConnectionServer) SetCertificate(v CertificateMonitorInfo)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *VCMonitorConnectionServer) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetId

`func (o *VCMonitorConnectionServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VCMonitorConnectionServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VCMonitorConnectionServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VCMonitorConnectionServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VCMonitorConnectionServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VCMonitorConnectionServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VCMonitorConnectionServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VCMonitorConnectionServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VCMonitorConnectionServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VCMonitorConnectionServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VCMonitorConnectionServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VCMonitorConnectionServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThumbprintAccepted

`func (o *VCMonitorConnectionServer) GetThumbprintAccepted() bool`

GetThumbprintAccepted returns the ThumbprintAccepted field if non-nil, zero value otherwise.

### GetThumbprintAcceptedOk

`func (o *VCMonitorConnectionServer) GetThumbprintAcceptedOk() (*bool, bool)`

GetThumbprintAcceptedOk returns a tuple with the ThumbprintAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprintAccepted

`func (o *VCMonitorConnectionServer) SetThumbprintAccepted(v bool)`

SetThumbprintAccepted sets ThumbprintAccepted field to given value.

### HasThumbprintAccepted

`func (o *VCMonitorConnectionServer) HasThumbprintAccepted() bool`

HasThumbprintAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


