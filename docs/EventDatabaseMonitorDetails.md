# EventDatabaseMonitorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** | The name of the database. | 
**Port** | **int32** | The port of the database server. | 
**Prefix** | **string** | The prefix for event tables in the database. | 
**ServerName** | **string** | The name or ip address of the database server. | 
**Type** | **string** | The type of the database. * ORACLE: An Oracle database. * SQL_SERVER: A SQL server database. * POSTGRESQL: A PostgreSQL database. | 
**UserName** | **string** | The username used to connect to the database. | 

## Methods

### NewEventDatabaseMonitorDetails

`func NewEventDatabaseMonitorDetails(databaseName string, port int32, prefix string, serverName string, type_ string, userName string, ) *EventDatabaseMonitorDetails`

NewEventDatabaseMonitorDetails instantiates a new EventDatabaseMonitorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDatabaseMonitorDetailsWithDefaults

`func NewEventDatabaseMonitorDetailsWithDefaults() *EventDatabaseMonitorDetails`

NewEventDatabaseMonitorDetailsWithDefaults instantiates a new EventDatabaseMonitorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *EventDatabaseMonitorDetails) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *EventDatabaseMonitorDetails) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *EventDatabaseMonitorDetails) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetPort

`func (o *EventDatabaseMonitorDetails) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EventDatabaseMonitorDetails) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EventDatabaseMonitorDetails) SetPort(v int32)`

SetPort sets Port field to given value.


### GetPrefix

`func (o *EventDatabaseMonitorDetails) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *EventDatabaseMonitorDetails) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *EventDatabaseMonitorDetails) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetServerName

`func (o *EventDatabaseMonitorDetails) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *EventDatabaseMonitorDetails) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *EventDatabaseMonitorDetails) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetType

`func (o *EventDatabaseMonitorDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventDatabaseMonitorDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventDatabaseMonitorDetails) SetType(v string)`

SetType sets Type field to given value.


### GetUserName

`func (o *EventDatabaseMonitorDetails) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EventDatabaseMonitorDetails) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EventDatabaseMonitorDetails) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


