# DesktopPoolMetricsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the desktop pool. | [optional] 
**NumConnectedSessions** | Pointer to **int32** | Number of connected sessions of the desktop pool. | [optional] 
**NumMachines** | Pointer to **int32** | Number of machines in the desktop pool. | [optional] 
**OccupancyCount** | Pointer to **int32** | Occupancy count for the desktop pool. * For dedicated assignment desktop, it is the number of assigned machine count. * For floating assignment desktop, it is the summation of the connected and disconnected sessions. | [optional] 

## Methods

### NewDesktopPoolMetricsInfo

`func NewDesktopPoolMetricsInfo() *DesktopPoolMetricsInfo`

NewDesktopPoolMetricsInfo instantiates a new DesktopPoolMetricsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolMetricsInfoWithDefaults

`func NewDesktopPoolMetricsInfoWithDefaults() *DesktopPoolMetricsInfo`

NewDesktopPoolMetricsInfoWithDefaults instantiates a new DesktopPoolMetricsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DesktopPoolMetricsInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DesktopPoolMetricsInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DesktopPoolMetricsInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DesktopPoolMetricsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumConnectedSessions

`func (o *DesktopPoolMetricsInfo) GetNumConnectedSessions() int32`

GetNumConnectedSessions returns the NumConnectedSessions field if non-nil, zero value otherwise.

### GetNumConnectedSessionsOk

`func (o *DesktopPoolMetricsInfo) GetNumConnectedSessionsOk() (*int32, bool)`

GetNumConnectedSessionsOk returns a tuple with the NumConnectedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumConnectedSessions

`func (o *DesktopPoolMetricsInfo) SetNumConnectedSessions(v int32)`

SetNumConnectedSessions sets NumConnectedSessions field to given value.

### HasNumConnectedSessions

`func (o *DesktopPoolMetricsInfo) HasNumConnectedSessions() bool`

HasNumConnectedSessions returns a boolean if a field has been set.

### GetNumMachines

`func (o *DesktopPoolMetricsInfo) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *DesktopPoolMetricsInfo) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *DesktopPoolMetricsInfo) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.

### HasNumMachines

`func (o *DesktopPoolMetricsInfo) HasNumMachines() bool`

HasNumMachines returns a boolean if a field has been set.

### GetOccupancyCount

`func (o *DesktopPoolMetricsInfo) GetOccupancyCount() int32`

GetOccupancyCount returns the OccupancyCount field if non-nil, zero value otherwise.

### GetOccupancyCountOk

`func (o *DesktopPoolMetricsInfo) GetOccupancyCountOk() (*int32, bool)`

GetOccupancyCountOk returns a tuple with the OccupancyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancyCount

`func (o *DesktopPoolMetricsInfo) SetOccupancyCount(v int32)`

SetOccupancyCount sets OccupancyCount field to given value.

### HasOccupancyCount

`func (o *DesktopPoolMetricsInfo) HasOccupancyCount() bool`

HasOccupancyCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


