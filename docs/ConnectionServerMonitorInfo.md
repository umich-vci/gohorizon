# ConnectionServerMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**CertificateMonitorInfo**](CertificateMonitorInfo.md) |  | [optional] 
**ConnectionCount** | Pointer to **int32** | Number of connections to this Connection Server. | [optional] 
**CsReplications** | Pointer to [**[]ConnectionServerMonitorCSReplication**](ConnectionServerMonitorCSReplication.md) | Connection Server replication status with respect to the Peer Connection Servers in the same cluster. | [optional] 
**DefaultCertificate** | Pointer to **bool** | Indicates whether server has the default certificate. | [optional] 
**Details** | Pointer to [**ConnectionServerMonitorDetails**](ConnectionServerMonitorDetails.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the Connection Server. | [optional] 
**Name** | Pointer to **string** | Connection Server host name or IP address. | [optional] 
**Services** | Pointer to [**[]ConnectionServerMonitorServiceStatus**](ConnectionServerMonitorServiceStatus.md) | Connection Server related Windows services information. | [optional] 
**SessionProtocolData** | Pointer to [**[]ConnectionServerSessionProtocolData**](ConnectionServerSessionProtocolData.md) | PCoIP, RDP or BLAST protocol sessions details when clients connect directly to the connection server. | [optional] 
**SessionThreshold** | Pointer to **int32** | The maximum number of connections allowed for the connection server through the Horizon client. If all of the secure gateways (HTTP(S)/PCOIP/BLAST) are enabled, this field denotes the maximum number of connections allowed for the connection server.If none of the secure gateways(HTTP(S)/PCOIP/BLAST) are enabled, sessionThreshold value will not be set. | [optional] 
**Status** | Pointer to **string** | Status of the Connection Server. * OK: The Connection Server is working properly. * ERROR: Error occurred when connecting to Connection Server. * NOT_RESPONDING: The Connection Server is not responding. * UNKNOWN: Status of Connection Server is unknown. | [optional] 
**TunnelConnectionCount** | Pointer to **int32** | Number of connections tunneled through this Connection Server. | [optional] 

## Methods

### NewConnectionServerMonitorInfo

`func NewConnectionServerMonitorInfo() *ConnectionServerMonitorInfo`

NewConnectionServerMonitorInfo instantiates a new ConnectionServerMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServerMonitorInfoWithDefaults

`func NewConnectionServerMonitorInfoWithDefaults() *ConnectionServerMonitorInfo`

NewConnectionServerMonitorInfoWithDefaults instantiates a new ConnectionServerMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *ConnectionServerMonitorInfo) GetCertificate() CertificateMonitorInfo`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ConnectionServerMonitorInfo) GetCertificateOk() (*CertificateMonitorInfo, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ConnectionServerMonitorInfo) SetCertificate(v CertificateMonitorInfo)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ConnectionServerMonitorInfo) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetConnectionCount

`func (o *ConnectionServerMonitorInfo) GetConnectionCount() int32`

GetConnectionCount returns the ConnectionCount field if non-nil, zero value otherwise.

### GetConnectionCountOk

`func (o *ConnectionServerMonitorInfo) GetConnectionCountOk() (*int32, bool)`

GetConnectionCountOk returns a tuple with the ConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCount

`func (o *ConnectionServerMonitorInfo) SetConnectionCount(v int32)`

SetConnectionCount sets ConnectionCount field to given value.

### HasConnectionCount

`func (o *ConnectionServerMonitorInfo) HasConnectionCount() bool`

HasConnectionCount returns a boolean if a field has been set.

### GetCsReplications

`func (o *ConnectionServerMonitorInfo) GetCsReplications() []ConnectionServerMonitorCSReplication`

GetCsReplications returns the CsReplications field if non-nil, zero value otherwise.

### GetCsReplicationsOk

`func (o *ConnectionServerMonitorInfo) GetCsReplicationsOk() (*[]ConnectionServerMonitorCSReplication, bool)`

GetCsReplicationsOk returns a tuple with the CsReplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsReplications

`func (o *ConnectionServerMonitorInfo) SetCsReplications(v []ConnectionServerMonitorCSReplication)`

SetCsReplications sets CsReplications field to given value.

### HasCsReplications

`func (o *ConnectionServerMonitorInfo) HasCsReplications() bool`

HasCsReplications returns a boolean if a field has been set.

### GetDefaultCertificate

`func (o *ConnectionServerMonitorInfo) GetDefaultCertificate() bool`

GetDefaultCertificate returns the DefaultCertificate field if non-nil, zero value otherwise.

### GetDefaultCertificateOk

`func (o *ConnectionServerMonitorInfo) GetDefaultCertificateOk() (*bool, bool)`

GetDefaultCertificateOk returns a tuple with the DefaultCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificate

`func (o *ConnectionServerMonitorInfo) SetDefaultCertificate(v bool)`

SetDefaultCertificate sets DefaultCertificate field to given value.

### HasDefaultCertificate

`func (o *ConnectionServerMonitorInfo) HasDefaultCertificate() bool`

HasDefaultCertificate returns a boolean if a field has been set.

### GetDetails

`func (o *ConnectionServerMonitorInfo) GetDetails() ConnectionServerMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConnectionServerMonitorInfo) GetDetailsOk() (*ConnectionServerMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConnectionServerMonitorInfo) SetDetails(v ConnectionServerMonitorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConnectionServerMonitorInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *ConnectionServerMonitorInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionServerMonitorInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionServerMonitorInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionServerMonitorInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConnectionServerMonitorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionServerMonitorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionServerMonitorInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionServerMonitorInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServices

`func (o *ConnectionServerMonitorInfo) GetServices() []ConnectionServerMonitorServiceStatus`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ConnectionServerMonitorInfo) GetServicesOk() (*[]ConnectionServerMonitorServiceStatus, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ConnectionServerMonitorInfo) SetServices(v []ConnectionServerMonitorServiceStatus)`

SetServices sets Services field to given value.

### HasServices

`func (o *ConnectionServerMonitorInfo) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSessionProtocolData

`func (o *ConnectionServerMonitorInfo) GetSessionProtocolData() []ConnectionServerSessionProtocolData`

GetSessionProtocolData returns the SessionProtocolData field if non-nil, zero value otherwise.

### GetSessionProtocolDataOk

`func (o *ConnectionServerMonitorInfo) GetSessionProtocolDataOk() (*[]ConnectionServerSessionProtocolData, bool)`

GetSessionProtocolDataOk returns a tuple with the SessionProtocolData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocolData

`func (o *ConnectionServerMonitorInfo) SetSessionProtocolData(v []ConnectionServerSessionProtocolData)`

SetSessionProtocolData sets SessionProtocolData field to given value.

### HasSessionProtocolData

`func (o *ConnectionServerMonitorInfo) HasSessionProtocolData() bool`

HasSessionProtocolData returns a boolean if a field has been set.

### GetSessionThreshold

`func (o *ConnectionServerMonitorInfo) GetSessionThreshold() int32`

GetSessionThreshold returns the SessionThreshold field if non-nil, zero value otherwise.

### GetSessionThresholdOk

`func (o *ConnectionServerMonitorInfo) GetSessionThresholdOk() (*int32, bool)`

GetSessionThresholdOk returns a tuple with the SessionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionThreshold

`func (o *ConnectionServerMonitorInfo) SetSessionThreshold(v int32)`

SetSessionThreshold sets SessionThreshold field to given value.

### HasSessionThreshold

`func (o *ConnectionServerMonitorInfo) HasSessionThreshold() bool`

HasSessionThreshold returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectionServerMonitorInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionServerMonitorInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionServerMonitorInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectionServerMonitorInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnelConnectionCount

`func (o *ConnectionServerMonitorInfo) GetTunnelConnectionCount() int32`

GetTunnelConnectionCount returns the TunnelConnectionCount field if non-nil, zero value otherwise.

### GetTunnelConnectionCountOk

`func (o *ConnectionServerMonitorInfo) GetTunnelConnectionCountOk() (*int32, bool)`

GetTunnelConnectionCountOk returns a tuple with the TunnelConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelConnectionCount

`func (o *ConnectionServerMonitorInfo) SetTunnelConnectionCount(v int32)`

SetTunnelConnectionCount sets TunnelConnectionCount field to given value.

### HasTunnelConnectionCount

`func (o *ConnectionServerMonitorInfo) HasTunnelConnectionCount() bool`

HasTunnelConnectionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


