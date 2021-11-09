# PodFederationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServerStatuses** | Pointer to [**[]ConnectionServerStatus**](ConnectionServerStatus.md) | Individual connection server CPA status for this pod. | [optional] 
**Guid** | Pointer to **string** | GUID representing the pod federation. | [optional] 
**LocalConnectionServerStatus** | Pointer to **string** | CPA status of the current connection server in the pod. * ENABLED: CPA is enabled. * DISABLED: CPA is disabled. * PENDING: CPA is undergoing an operation related to initialization, uninitialization, joining, or unjoining. * ENABLE_ERROR: The connection server has failed to reach the ENABLED status in a timely manner. This may also indicate the current connection server was recently installed. * DISABLE_ERROR: The connection server has failed to reach the DISABLED status in a timely manner. | [optional] 
**Name** | Pointer to **string** | Name of the pod federation. | [optional] 
**Sites** | Pointer to **[]string** | Member sites in the pod federation. | [optional] 

## Methods

### NewPodFederationInfo

`func NewPodFederationInfo() *PodFederationInfo`

NewPodFederationInfo instantiates a new PodFederationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodFederationInfoWithDefaults

`func NewPodFederationInfoWithDefaults() *PodFederationInfo`

NewPodFederationInfoWithDefaults instantiates a new PodFederationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServerStatuses

`func (o *PodFederationInfo) GetConnectionServerStatuses() []ConnectionServerStatus`

GetConnectionServerStatuses returns the ConnectionServerStatuses field if non-nil, zero value otherwise.

### GetConnectionServerStatusesOk

`func (o *PodFederationInfo) GetConnectionServerStatusesOk() (*[]ConnectionServerStatus, bool)`

GetConnectionServerStatusesOk returns a tuple with the ConnectionServerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServerStatuses

`func (o *PodFederationInfo) SetConnectionServerStatuses(v []ConnectionServerStatus)`

SetConnectionServerStatuses sets ConnectionServerStatuses field to given value.

### HasConnectionServerStatuses

`func (o *PodFederationInfo) HasConnectionServerStatuses() bool`

HasConnectionServerStatuses returns a boolean if a field has been set.

### GetGuid

`func (o *PodFederationInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *PodFederationInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *PodFederationInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *PodFederationInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLocalConnectionServerStatus

`func (o *PodFederationInfo) GetLocalConnectionServerStatus() string`

GetLocalConnectionServerStatus returns the LocalConnectionServerStatus field if non-nil, zero value otherwise.

### GetLocalConnectionServerStatusOk

`func (o *PodFederationInfo) GetLocalConnectionServerStatusOk() (*string, bool)`

GetLocalConnectionServerStatusOk returns a tuple with the LocalConnectionServerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConnectionServerStatus

`func (o *PodFederationInfo) SetLocalConnectionServerStatus(v string)`

SetLocalConnectionServerStatus sets LocalConnectionServerStatus field to given value.

### HasLocalConnectionServerStatus

`func (o *PodFederationInfo) HasLocalConnectionServerStatus() bool`

HasLocalConnectionServerStatus returns a boolean if a field has been set.

### GetName

`func (o *PodFederationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodFederationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodFederationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PodFederationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSites

`func (o *PodFederationInfo) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PodFederationInfo) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PodFederationInfo) SetSites(v []string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *PodFederationInfo) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


