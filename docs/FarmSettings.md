# FarmSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteInProgess** | **bool** | Indicates whether the Farm is in the process of being deleted. Default value is false. | 
**DesktopId** | Pointer to **string** | Desktop pool Id representing the RDS Desktop pool to which this Farm belongs. | [optional] 
**DisplayProtocolSettings** | Pointer to [**FarmDisplayProtocolSettings**](FarmDisplayProtocolSettings.md) |  | [optional] 
**LoadBalancerSettings** | Pointer to [**FarmLoadBalancerSettings**](FarmLoadBalancerSettings.md) |  | [optional] 
**ServerErrorThreshold** | Pointer to **int32** | The minimum number of machines that must be fully operational in order toavoid showing the farm in an error state. Default value is 0. | [optional] 
**SessionSettings** | Pointer to [**FarmSessionSettings**](FarmSessionSettings.md) |  | [optional] 

## Methods

### NewFarmSettings

`func NewFarmSettings(deleteInProgess bool, ) *FarmSettings`

NewFarmSettings instantiates a new FarmSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmSettingsWithDefaults

`func NewFarmSettingsWithDefaults() *FarmSettings`

NewFarmSettingsWithDefaults instantiates a new FarmSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteInProgess

`func (o *FarmSettings) GetDeleteInProgess() bool`

GetDeleteInProgess returns the DeleteInProgess field if non-nil, zero value otherwise.

### GetDeleteInProgessOk

`func (o *FarmSettings) GetDeleteInProgessOk() (*bool, bool)`

GetDeleteInProgessOk returns a tuple with the DeleteInProgess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInProgess

`func (o *FarmSettings) SetDeleteInProgess(v bool)`

SetDeleteInProgess sets DeleteInProgess field to given value.


### GetDesktopId

`func (o *FarmSettings) GetDesktopId() string`

GetDesktopId returns the DesktopId field if non-nil, zero value otherwise.

### GetDesktopIdOk

`func (o *FarmSettings) GetDesktopIdOk() (*string, bool)`

GetDesktopIdOk returns a tuple with the DesktopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopId

`func (o *FarmSettings) SetDesktopId(v string)`

SetDesktopId sets DesktopId field to given value.

### HasDesktopId

`func (o *FarmSettings) HasDesktopId() bool`

HasDesktopId returns a boolean if a field has been set.

### GetDisplayProtocolSettings

`func (o *FarmSettings) GetDisplayProtocolSettings() FarmDisplayProtocolSettings`

GetDisplayProtocolSettings returns the DisplayProtocolSettings field if non-nil, zero value otherwise.

### GetDisplayProtocolSettingsOk

`func (o *FarmSettings) GetDisplayProtocolSettingsOk() (*FarmDisplayProtocolSettings, bool)`

GetDisplayProtocolSettingsOk returns a tuple with the DisplayProtocolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocolSettings

`func (o *FarmSettings) SetDisplayProtocolSettings(v FarmDisplayProtocolSettings)`

SetDisplayProtocolSettings sets DisplayProtocolSettings field to given value.

### HasDisplayProtocolSettings

`func (o *FarmSettings) HasDisplayProtocolSettings() bool`

HasDisplayProtocolSettings returns a boolean if a field has been set.

### GetLoadBalancerSettings

`func (o *FarmSettings) GetLoadBalancerSettings() FarmLoadBalancerSettings`

GetLoadBalancerSettings returns the LoadBalancerSettings field if non-nil, zero value otherwise.

### GetLoadBalancerSettingsOk

`func (o *FarmSettings) GetLoadBalancerSettingsOk() (*FarmLoadBalancerSettings, bool)`

GetLoadBalancerSettingsOk returns a tuple with the LoadBalancerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSettings

`func (o *FarmSettings) SetLoadBalancerSettings(v FarmLoadBalancerSettings)`

SetLoadBalancerSettings sets LoadBalancerSettings field to given value.

### HasLoadBalancerSettings

`func (o *FarmSettings) HasLoadBalancerSettings() bool`

HasLoadBalancerSettings returns a boolean if a field has been set.

### GetServerErrorThreshold

`func (o *FarmSettings) GetServerErrorThreshold() int32`

GetServerErrorThreshold returns the ServerErrorThreshold field if non-nil, zero value otherwise.

### GetServerErrorThresholdOk

`func (o *FarmSettings) GetServerErrorThresholdOk() (*int32, bool)`

GetServerErrorThresholdOk returns a tuple with the ServerErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerErrorThreshold

`func (o *FarmSettings) SetServerErrorThreshold(v int32)`

SetServerErrorThreshold sets ServerErrorThreshold field to given value.

### HasServerErrorThreshold

`func (o *FarmSettings) HasServerErrorThreshold() bool`

HasServerErrorThreshold returns a boolean if a field has been set.

### GetSessionSettings

`func (o *FarmSettings) GetSessionSettings() FarmSessionSettings`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *FarmSettings) GetSessionSettingsOk() (*FarmSessionSettings, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *FarmSettings) SetSessionSettings(v FarmSessionSettings)`

SetSessionSettings sets SessionSettings field to given value.

### HasSessionSettings

`func (o *FarmSettings) HasSessionSettings() bool`

HasSessionSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


