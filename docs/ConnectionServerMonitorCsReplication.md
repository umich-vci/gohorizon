# ConnectionServerMonitorCSReplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Connection Server host name or IP address. | 
**Status** | **string** | LDAP replication status. * OK: The connection to the Connection Server is working properly. * ERROR: There is a problem with LDAP replication to the Connection Server. | 

## Methods

### NewConnectionServerMonitorCSReplication

`func NewConnectionServerMonitorCSReplication(serverName string, status string, ) *ConnectionServerMonitorCSReplication`

NewConnectionServerMonitorCSReplication instantiates a new ConnectionServerMonitorCSReplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServerMonitorCSReplicationWithDefaults

`func NewConnectionServerMonitorCSReplicationWithDefaults() *ConnectionServerMonitorCSReplication`

NewConnectionServerMonitorCSReplicationWithDefaults instantiates a new ConnectionServerMonitorCSReplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *ConnectionServerMonitorCSReplication) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ConnectionServerMonitorCSReplication) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ConnectionServerMonitorCSReplication) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetStatus

`func (o *ConnectionServerMonitorCSReplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionServerMonitorCSReplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionServerMonitorCSReplication) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


