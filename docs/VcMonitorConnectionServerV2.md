# VCMonitorConnectionServerV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**CertificateMonitorInfo**](CertificateMonitorInfo.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the Connection Server. | [optional] 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp in milliseconds when the last update was obtained. Measured as epoch time. | [optional] 
**Name** | Pointer to **string** | Connection server host name or IP address. | [optional] 
**Status** | Pointer to **string** | Status of the Virtual Center Connection with respect to this Connection Server. * OK: The connection to Virtual Center server is working properly. * DOWN: The connection to Virtual Center server is down. * RECONNECTING: The connection to Virtual Center server was lost and is being reconnected to. * UNKNOWN: Connection state to Virtual Center server is unknown. * INVALID_CREDENTIALS: The supplied credentials cannot be used to authenticate to the Virtual Center server. * CANNOT_LOGIN: The connection server cannot login to the Virtual Center server. * NOT_YET_CONNECTED: Connection server has not yet connected to Virtual Center server. | [optional] 
**ThumbprintAccepted** | Pointer to **bool** | Indicates if the thumbprints of the Virtual Center was accepted. | [optional] 

## Methods

### NewVCMonitorConnectionServerV2

`func NewVCMonitorConnectionServerV2() *VCMonitorConnectionServerV2`

NewVCMonitorConnectionServerV2 instantiates a new VCMonitorConnectionServerV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCMonitorConnectionServerV2WithDefaults

`func NewVCMonitorConnectionServerV2WithDefaults() *VCMonitorConnectionServerV2`

NewVCMonitorConnectionServerV2WithDefaults instantiates a new VCMonitorConnectionServerV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *VCMonitorConnectionServerV2) GetCertificate() CertificateMonitorInfo`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *VCMonitorConnectionServerV2) GetCertificateOk() (*CertificateMonitorInfo, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *VCMonitorConnectionServerV2) SetCertificate(v CertificateMonitorInfo)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *VCMonitorConnectionServerV2) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetId

`func (o *VCMonitorConnectionServerV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VCMonitorConnectionServerV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VCMonitorConnectionServerV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VCMonitorConnectionServerV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *VCMonitorConnectionServerV2) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *VCMonitorConnectionServerV2) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *VCMonitorConnectionServerV2) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *VCMonitorConnectionServerV2) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *VCMonitorConnectionServerV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VCMonitorConnectionServerV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VCMonitorConnectionServerV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VCMonitorConnectionServerV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VCMonitorConnectionServerV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VCMonitorConnectionServerV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VCMonitorConnectionServerV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VCMonitorConnectionServerV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThumbprintAccepted

`func (o *VCMonitorConnectionServerV2) GetThumbprintAccepted() bool`

GetThumbprintAccepted returns the ThumbprintAccepted field if non-nil, zero value otherwise.

### GetThumbprintAcceptedOk

`func (o *VCMonitorConnectionServerV2) GetThumbprintAcceptedOk() (*bool, bool)`

GetThumbprintAcceptedOk returns a tuple with the ThumbprintAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprintAccepted

`func (o *VCMonitorConnectionServerV2) SetThumbprintAccepted(v bool)`

SetThumbprintAccepted sets ThumbprintAccepted field to given value.

### HasThumbprintAccepted

`func (o *VCMonitorConnectionServerV2) HasThumbprintAccepted() bool`

HasThumbprintAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


