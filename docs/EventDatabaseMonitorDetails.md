# EventDatabaseMonitorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | Pointer to **string** | The name of the database. | [optional] 
**Port** | Pointer to **int32** | The port of the database server. | [optional] 
**Prefix** | Pointer to **string** | The prefix for event tables in the database. | [optional] 
**ServerName** | Pointer to **string** | The name or ip address of the database server. | [optional] 
**Type** | Pointer to **string** | The type of the database. * ORACLE: An Oracle database. * SQL_SERVER: A SQL server database. * POSTGRESQL: A PostgreSQL database. | [optional] 
**UserName** | Pointer to **string** | The username used to connect to the database. | [optional] 

## Methods

### NewEventDatabaseMonitorDetails

`func NewEventDatabaseMonitorDetails() *EventDatabaseMonitorDetails`

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

### HasDatabaseName

`func (o *EventDatabaseMonitorDetails) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

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

### HasPort

`func (o *EventDatabaseMonitorDetails) HasPort() bool`

HasPort returns a boolean if a field has been set.

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

### HasPrefix

`func (o *EventDatabaseMonitorDetails) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

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

### HasServerName

`func (o *EventDatabaseMonitorDetails) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

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

### HasType

`func (o *EventDatabaseMonitorDetails) HasType() bool`

HasType returns a boolean if a field has been set.

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

### HasUserName

`func (o *EventDatabaseMonitorDetails) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


