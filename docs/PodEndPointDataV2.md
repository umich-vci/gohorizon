# PodEndPointDataV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates whether an endpoint is enabled. A disabled endpoint will be excluded from participating inter-pod communication. | 
**Id** | **string** | Unique ID for a pod endpoint. | 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp in milliseconds when the last update was obtained. Measured as epoch time. | [optional] 
**Name** | **string** | Name for the pod endpoint. | 
**RoundtripTime** | Pointer to **int64** | Round trip time (in milliseconds) for ping request between the local pod endpoint and the remote pod. | [optional] 
**Status** | **string** | Status of the pod endpoint. * ONLINE: Pod endpoint is online and functional. * UNCHECKED: Pod endpoint was offline and it just came back online but the system has not verified that it is functional. * OFFLINE: Pod endpoint is offline or unreachable. | 
**Url** | **string** | The URL for the pod endpoint. This address and special port will be used for inter-pod communication. | 

## Methods

### NewPodEndPointDataV2

`func NewPodEndPointDataV2(enabled bool, id string, name string, status string, url string, ) *PodEndPointDataV2`

NewPodEndPointDataV2 instantiates a new PodEndPointDataV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodEndPointDataV2WithDefaults

`func NewPodEndPointDataV2WithDefaults() *PodEndPointDataV2`

NewPodEndPointDataV2WithDefaults instantiates a new PodEndPointDataV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PodEndPointDataV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PodEndPointDataV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PodEndPointDataV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *PodEndPointDataV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PodEndPointDataV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PodEndPointDataV2) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTimestamp

`func (o *PodEndPointDataV2) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *PodEndPointDataV2) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *PodEndPointDataV2) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *PodEndPointDataV2) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *PodEndPointDataV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodEndPointDataV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodEndPointDataV2) SetName(v string)`

SetName sets Name field to given value.


### GetRoundtripTime

`func (o *PodEndPointDataV2) GetRoundtripTime() int64`

GetRoundtripTime returns the RoundtripTime field if non-nil, zero value otherwise.

### GetRoundtripTimeOk

`func (o *PodEndPointDataV2) GetRoundtripTimeOk() (*int64, bool)`

GetRoundtripTimeOk returns a tuple with the RoundtripTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundtripTime

`func (o *PodEndPointDataV2) SetRoundtripTime(v int64)`

SetRoundtripTime sets RoundtripTime field to given value.

### HasRoundtripTime

`func (o *PodEndPointDataV2) HasRoundtripTime() bool`

HasRoundtripTime returns a boolean if a field has been set.

### GetStatus

`func (o *PodEndPointDataV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PodEndPointDataV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PodEndPointDataV2) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUrl

`func (o *PodEndPointDataV2) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PodEndPointDataV2) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PodEndPointDataV2) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


