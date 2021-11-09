# EnvironmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterGuid** | **string** | The GUID of a group of Connection Servers sharing the same configuration. | 
**ClusterName** | **string** | The name of a group of Connection Servers sharing the same configuration. | 
**FipsModeEnabled** | **bool** | Indicates if FIPS mode is enabled. | 
**IpMode** | **string** | Indicates the IP mode of the environment. * IPv4: The ip mode is IPv4. * IPv6: The ip mode is IPv6. | 
**LocalPodName** | Pointer to **string** | The name of the current pod in the Multi-DataCenter Horizon Pod, the value will be null when PodFederation is not initialized. | [optional] 
**TimezoneOffset** | **int32** | Represents the clusters time zone offset from UTC in seconds. | 

## Methods

### NewEnvironmentInfo

`func NewEnvironmentInfo(clusterGuid string, clusterName string, fipsModeEnabled bool, ipMode string, timezoneOffset int32, ) *EnvironmentInfo`

NewEnvironmentInfo instantiates a new EnvironmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentInfoWithDefaults

`func NewEnvironmentInfoWithDefaults() *EnvironmentInfo`

NewEnvironmentInfoWithDefaults instantiates a new EnvironmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterGuid

`func (o *EnvironmentInfo) GetClusterGuid() string`

GetClusterGuid returns the ClusterGuid field if non-nil, zero value otherwise.

### GetClusterGuidOk

`func (o *EnvironmentInfo) GetClusterGuidOk() (*string, bool)`

GetClusterGuidOk returns a tuple with the ClusterGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGuid

`func (o *EnvironmentInfo) SetClusterGuid(v string)`

SetClusterGuid sets ClusterGuid field to given value.


### GetClusterName

`func (o *EnvironmentInfo) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *EnvironmentInfo) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *EnvironmentInfo) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetFipsModeEnabled

`func (o *EnvironmentInfo) GetFipsModeEnabled() bool`

GetFipsModeEnabled returns the FipsModeEnabled field if non-nil, zero value otherwise.

### GetFipsModeEnabledOk

`func (o *EnvironmentInfo) GetFipsModeEnabledOk() (*bool, bool)`

GetFipsModeEnabledOk returns a tuple with the FipsModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsModeEnabled

`func (o *EnvironmentInfo) SetFipsModeEnabled(v bool)`

SetFipsModeEnabled sets FipsModeEnabled field to given value.


### GetIpMode

`func (o *EnvironmentInfo) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *EnvironmentInfo) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *EnvironmentInfo) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.


### GetLocalPodName

`func (o *EnvironmentInfo) GetLocalPodName() string`

GetLocalPodName returns the LocalPodName field if non-nil, zero value otherwise.

### GetLocalPodNameOk

`func (o *EnvironmentInfo) GetLocalPodNameOk() (*string, bool)`

GetLocalPodNameOk returns a tuple with the LocalPodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPodName

`func (o *EnvironmentInfo) SetLocalPodName(v string)`

SetLocalPodName sets LocalPodName field to given value.

### HasLocalPodName

`func (o *EnvironmentInfo) HasLocalPodName() bool`

HasLocalPodName returns a boolean if a field has been set.

### GetTimezoneOffset

`func (o *EnvironmentInfo) GetTimezoneOffset() int32`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *EnvironmentInfo) GetTimezoneOffsetOk() (*int32, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *EnvironmentInfo) SetTimezoneOffset(v int32)`

SetTimezoneOffset sets TimezoneOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


