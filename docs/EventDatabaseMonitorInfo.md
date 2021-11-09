# EventDatabaseMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**EventDatabaseMonitorDetails**](EventDatabaseMonitorDetails.md) |  | [optional] 
**EventCount** | Pointer to **int32** | Number of events recorded in the database | [optional] 
**Status** | Pointer to **string** | State of the database. * CONNECTED: Connection Server is connected to the database. * CONNECTING: Connection Server is connecting to the database. * DISCONNECTED: Connection Server is disconnected from the database. * RECONNECTING: Connection Server is reconnecting to the database. * ERROR: Error in connecting to the database from Connection Server. * NOT_CONFIGURED: Database is not configured. | [optional] 

## Methods

### NewEventDatabaseMonitorInfo

`func NewEventDatabaseMonitorInfo() *EventDatabaseMonitorInfo`

NewEventDatabaseMonitorInfo instantiates a new EventDatabaseMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDatabaseMonitorInfoWithDefaults

`func NewEventDatabaseMonitorInfoWithDefaults() *EventDatabaseMonitorInfo`

NewEventDatabaseMonitorInfoWithDefaults instantiates a new EventDatabaseMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *EventDatabaseMonitorInfo) GetDetails() EventDatabaseMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EventDatabaseMonitorInfo) GetDetailsOk() (*EventDatabaseMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EventDatabaseMonitorInfo) SetDetails(v EventDatabaseMonitorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EventDatabaseMonitorInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEventCount

`func (o *EventDatabaseMonitorInfo) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *EventDatabaseMonitorInfo) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *EventDatabaseMonitorInfo) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.

### HasEventCount

`func (o *EventDatabaseMonitorInfo) HasEventCount() bool`

HasEventCount returns a boolean if a field has been set.

### GetStatus

`func (o *EventDatabaseMonitorInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventDatabaseMonitorInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventDatabaseMonitorInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventDatabaseMonitorInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


