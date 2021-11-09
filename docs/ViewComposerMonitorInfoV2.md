# ViewComposerMonitorInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | [**[]ViewComposerMonitorConnectionServerV2**](ViewComposerMonitorConnectionServerV2.md) | Information about the View Composer connections from each of the connection servers. | 
**Details** | [**ViewComposerMonitorDetails**](ViewComposerMonitorDetails.md) |  | 
**Name** | **string** | View Composer server host name or IP address. | 
**Port** | **int32** | View Composer server port number. | 

## Methods

### NewViewComposerMonitorInfoV2

`func NewViewComposerMonitorInfoV2(connectionServers []ViewComposerMonitorConnectionServerV2, details ViewComposerMonitorDetails, name string, port int32, ) *ViewComposerMonitorInfoV2`

NewViewComposerMonitorInfoV2 instantiates a new ViewComposerMonitorInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewComposerMonitorInfoV2WithDefaults

`func NewViewComposerMonitorInfoV2WithDefaults() *ViewComposerMonitorInfoV2`

NewViewComposerMonitorInfoV2WithDefaults instantiates a new ViewComposerMonitorInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *ViewComposerMonitorInfoV2) GetConnectionServers() []ViewComposerMonitorConnectionServerV2`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *ViewComposerMonitorInfoV2) GetConnectionServersOk() (*[]ViewComposerMonitorConnectionServerV2, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *ViewComposerMonitorInfoV2) SetConnectionServers(v []ViewComposerMonitorConnectionServerV2)`

SetConnectionServers sets ConnectionServers field to given value.


### GetDetails

`func (o *ViewComposerMonitorInfoV2) GetDetails() ViewComposerMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ViewComposerMonitorInfoV2) GetDetailsOk() (*ViewComposerMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ViewComposerMonitorInfoV2) SetDetails(v ViewComposerMonitorDetails)`

SetDetails sets Details field to given value.


### GetName

`func (o *ViewComposerMonitorInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewComposerMonitorInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewComposerMonitorInfoV2) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *ViewComposerMonitorInfoV2) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ViewComposerMonitorInfoV2) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ViewComposerMonitorInfoV2) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


