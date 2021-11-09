# ConnectionServerMonitorInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**CertificateMonitorInfo**](CertificateMonitorInfo.md) |  | 
**ConnectionCount** | Pointer to **int32** | Number of connections to this Connection Server. | [optional] 
**CsReplications** | Pointer to [**[]ConnectionServerMonitorCSReplication**](ConnectionServerMonitorCSReplication.md) | Connection Server replication status with respect to the Peer Connection Servers in the same cluster. | [optional] 
**DefaultCertificate** | Pointer to **bool** | Indicates whether server has the default certificate. | [optional] 
**Details** | [**ConnectionServerMonitorDetails**](ConnectionServerMonitorDetails.md) |  | 
**Id** | **string** | Unique ID of the Connection Server. | 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp in milliseconds when the last update was obtained. Measured as epoch time. | [optional] 
**Name** | **string** | Connection Server host name or IP address. | 
**Services** | Pointer to [**[]ConnectionServerMonitorServiceStatus**](ConnectionServerMonitorServiceStatus.md) | Connection Server related Windows services information. | [optional] 
**SessionProtocolData** | Pointer to [**[]ConnectionServerSessionProtocolData**](ConnectionServerSessionProtocolData.md) | PCoIP, RDP or BLAST protocol sessions details when clients connect directly to the connection server. | [optional] 
**SessionThreshold** | Pointer to **int32** | The maximum number of connections allowed for the connection server through the Horizon client. If all of the secure gateways (HTTP(S)/PCOIP/BLAST) are enabled, this field denotes the maximum number of connections allowed for the connection server.If none of the secure gateways(HTTP(S)/PCOIP/BLAST) are enabled, sessionThreshold value will not be set. | [optional] 
**Status** | **string** | Status of the Connection Server. * OK: The Connection Server is working properly. * ERROR: Error occurred when connecting to Connection Server. * NOT_RESPONDING: The Connection Server is not responding. * UNKNOWN: Status of Connection Server is unknown. | 
**TunnelConnectionCount** | Pointer to **int32** | Number of connections tunneled through this Connection Server. | [optional] 

## Methods

### NewConnectionServerMonitorInfoV2

`func NewConnectionServerMonitorInfoV2(certificate CertificateMonitorInfo, details ConnectionServerMonitorDetails, id string, name string, status string, ) *ConnectionServerMonitorInfoV2`

NewConnectionServerMonitorInfoV2 instantiates a new ConnectionServerMonitorInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServerMonitorInfoV2WithDefaults

`func NewConnectionServerMonitorInfoV2WithDefaults() *ConnectionServerMonitorInfoV2`

NewConnectionServerMonitorInfoV2WithDefaults instantiates a new ConnectionServerMonitorInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *ConnectionServerMonitorInfoV2) GetCertificate() CertificateMonitorInfo`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ConnectionServerMonitorInfoV2) GetCertificateOk() (*CertificateMonitorInfo, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ConnectionServerMonitorInfoV2) SetCertificate(v CertificateMonitorInfo)`

SetCertificate sets Certificate field to given value.


### GetConnectionCount

`func (o *ConnectionServerMonitorInfoV2) GetConnectionCount() int32`

GetConnectionCount returns the ConnectionCount field if non-nil, zero value otherwise.

### GetConnectionCountOk

`func (o *ConnectionServerMonitorInfoV2) GetConnectionCountOk() (*int32, bool)`

GetConnectionCountOk returns a tuple with the ConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCount

`func (o *ConnectionServerMonitorInfoV2) SetConnectionCount(v int32)`

SetConnectionCount sets ConnectionCount field to given value.

### HasConnectionCount

`func (o *ConnectionServerMonitorInfoV2) HasConnectionCount() bool`

HasConnectionCount returns a boolean if a field has been set.

### GetCsReplications

`func (o *ConnectionServerMonitorInfoV2) GetCsReplications() []ConnectionServerMonitorCSReplication`

GetCsReplications returns the CsReplications field if non-nil, zero value otherwise.

### GetCsReplicationsOk

`func (o *ConnectionServerMonitorInfoV2) GetCsReplicationsOk() (*[]ConnectionServerMonitorCSReplication, bool)`

GetCsReplicationsOk returns a tuple with the CsReplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsReplications

`func (o *ConnectionServerMonitorInfoV2) SetCsReplications(v []ConnectionServerMonitorCSReplication)`

SetCsReplications sets CsReplications field to given value.

### HasCsReplications

`func (o *ConnectionServerMonitorInfoV2) HasCsReplications() bool`

HasCsReplications returns a boolean if a field has been set.

### GetDefaultCertificate

`func (o *ConnectionServerMonitorInfoV2) GetDefaultCertificate() bool`

GetDefaultCertificate returns the DefaultCertificate field if non-nil, zero value otherwise.

### GetDefaultCertificateOk

`func (o *ConnectionServerMonitorInfoV2) GetDefaultCertificateOk() (*bool, bool)`

GetDefaultCertificateOk returns a tuple with the DefaultCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificate

`func (o *ConnectionServerMonitorInfoV2) SetDefaultCertificate(v bool)`

SetDefaultCertificate sets DefaultCertificate field to given value.

### HasDefaultCertificate

`func (o *ConnectionServerMonitorInfoV2) HasDefaultCertificate() bool`

HasDefaultCertificate returns a boolean if a field has been set.

### GetDetails

`func (o *ConnectionServerMonitorInfoV2) GetDetails() ConnectionServerMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConnectionServerMonitorInfoV2) GetDetailsOk() (*ConnectionServerMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConnectionServerMonitorInfoV2) SetDetails(v ConnectionServerMonitorDetails)`

SetDetails sets Details field to given value.


### GetId

`func (o *ConnectionServerMonitorInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionServerMonitorInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionServerMonitorInfoV2) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTimestamp

`func (o *ConnectionServerMonitorInfoV2) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *ConnectionServerMonitorInfoV2) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *ConnectionServerMonitorInfoV2) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *ConnectionServerMonitorInfoV2) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ConnectionServerMonitorInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionServerMonitorInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionServerMonitorInfoV2) SetName(v string)`

SetName sets Name field to given value.


### GetServices

`func (o *ConnectionServerMonitorInfoV2) GetServices() []ConnectionServerMonitorServiceStatus`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ConnectionServerMonitorInfoV2) GetServicesOk() (*[]ConnectionServerMonitorServiceStatus, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ConnectionServerMonitorInfoV2) SetServices(v []ConnectionServerMonitorServiceStatus)`

SetServices sets Services field to given value.

### HasServices

`func (o *ConnectionServerMonitorInfoV2) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSessionProtocolData

`func (o *ConnectionServerMonitorInfoV2) GetSessionProtocolData() []ConnectionServerSessionProtocolData`

GetSessionProtocolData returns the SessionProtocolData field if non-nil, zero value otherwise.

### GetSessionProtocolDataOk

`func (o *ConnectionServerMonitorInfoV2) GetSessionProtocolDataOk() (*[]ConnectionServerSessionProtocolData, bool)`

GetSessionProtocolDataOk returns a tuple with the SessionProtocolData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocolData

`func (o *ConnectionServerMonitorInfoV2) SetSessionProtocolData(v []ConnectionServerSessionProtocolData)`

SetSessionProtocolData sets SessionProtocolData field to given value.

### HasSessionProtocolData

`func (o *ConnectionServerMonitorInfoV2) HasSessionProtocolData() bool`

HasSessionProtocolData returns a boolean if a field has been set.

### GetSessionThreshold

`func (o *ConnectionServerMonitorInfoV2) GetSessionThreshold() int32`

GetSessionThreshold returns the SessionThreshold field if non-nil, zero value otherwise.

### GetSessionThresholdOk

`func (o *ConnectionServerMonitorInfoV2) GetSessionThresholdOk() (*int32, bool)`

GetSessionThresholdOk returns a tuple with the SessionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionThreshold

`func (o *ConnectionServerMonitorInfoV2) SetSessionThreshold(v int32)`

SetSessionThreshold sets SessionThreshold field to given value.

### HasSessionThreshold

`func (o *ConnectionServerMonitorInfoV2) HasSessionThreshold() bool`

HasSessionThreshold returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectionServerMonitorInfoV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionServerMonitorInfoV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionServerMonitorInfoV2) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTunnelConnectionCount

`func (o *ConnectionServerMonitorInfoV2) GetTunnelConnectionCount() int32`

GetTunnelConnectionCount returns the TunnelConnectionCount field if non-nil, zero value otherwise.

### GetTunnelConnectionCountOk

`func (o *ConnectionServerMonitorInfoV2) GetTunnelConnectionCountOk() (*int32, bool)`

GetTunnelConnectionCountOk returns a tuple with the TunnelConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelConnectionCount

`func (o *ConnectionServerMonitorInfoV2) SetTunnelConnectionCount(v int32)`

SetTunnelConnectionCount sets TunnelConnectionCount field to given value.

### HasTunnelConnectionCount

`func (o *ConnectionServerMonitorInfoV2) HasTunnelConnectionCount() bool`

HasTunnelConnectionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


