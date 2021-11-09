# ViewComposerMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | [**[]ViewComposerMonitorConnectionServer**](ViewComposerMonitorConnectionServer.md) | Information about the View Composer connections from each of the connection servers. | 
**Details** | [**ViewComposerMonitorDetails**](ViewComposerMonitorDetails.md) |  | 
**Name** | **string** | View Composer server host name or IP address. | 
**Port** | **int32** | View Composer server port number. | 

## Methods

### NewViewComposerMonitorInfo

`func NewViewComposerMonitorInfo(connectionServers []ViewComposerMonitorConnectionServer, details ViewComposerMonitorDetails, name string, port int32, ) *ViewComposerMonitorInfo`

NewViewComposerMonitorInfo instantiates a new ViewComposerMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewComposerMonitorInfoWithDefaults

`func NewViewComposerMonitorInfoWithDefaults() *ViewComposerMonitorInfo`

NewViewComposerMonitorInfoWithDefaults instantiates a new ViewComposerMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *ViewComposerMonitorInfo) GetConnectionServers() []ViewComposerMonitorConnectionServer`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *ViewComposerMonitorInfo) GetConnectionServersOk() (*[]ViewComposerMonitorConnectionServer, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *ViewComposerMonitorInfo) SetConnectionServers(v []ViewComposerMonitorConnectionServer)`

SetConnectionServers sets ConnectionServers field to given value.


### GetDetails

`func (o *ViewComposerMonitorInfo) GetDetails() ViewComposerMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ViewComposerMonitorInfo) GetDetailsOk() (*ViewComposerMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ViewComposerMonitorInfo) SetDetails(v ViewComposerMonitorDetails)`

SetDetails sets Details field to given value.


### GetName

`func (o *ViewComposerMonitorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewComposerMonitorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewComposerMonitorInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *ViewComposerMonitorInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ViewComposerMonitorInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ViewComposerMonitorInfo) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


