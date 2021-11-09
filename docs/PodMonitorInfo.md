# PodMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]PodEndPointData**](PodEndPointData.md) | The list of pod endpoints within this pod. An endpoint is basically a connection server in that pod. | [optional] 
**Id** | **string** | Unique ID of the pod. | 
**Name** | **string** | Display name for the pod. | 
**SiteId** | **string** | The Id of the site this pod belongs to. | 

## Methods

### NewPodMonitorInfo

`func NewPodMonitorInfo(id string, name string, siteId string, ) *PodMonitorInfo`

NewPodMonitorInfo instantiates a new PodMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodMonitorInfoWithDefaults

`func NewPodMonitorInfoWithDefaults() *PodMonitorInfo`

NewPodMonitorInfoWithDefaults instantiates a new PodMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *PodMonitorInfo) GetEndpoints() []PodEndPointData`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PodMonitorInfo) GetEndpointsOk() (*[]PodEndPointData, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PodMonitorInfo) SetEndpoints(v []PodEndPointData)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *PodMonitorInfo) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetId

`func (o *PodMonitorInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PodMonitorInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PodMonitorInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PodMonitorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodMonitorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodMonitorInfo) SetName(v string)`

SetName sets Name field to given value.


### GetSiteId

`func (o *PodMonitorInfo) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PodMonitorInfo) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PodMonitorInfo) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


