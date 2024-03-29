# HostOrClusterDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **bool** | Whether or not this is a cluster or a host. | [optional] 
**DatacenterId** | Pointer to **string** | Datacenter id for this host or cluster. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this Host Or Cluster from being used in desktop pool creation. | [optional] 
**Name** | Pointer to **string** | Host or cluster display name. | [optional] 
**Path** | Pointer to **string** | Host or cluster path. | [optional] 
**VcenterId** | Pointer to **string** | Virtual Center id for this host or cluster. | [optional] 
**VgpuTypes** | Pointer to **[]string** | Types of NVIDIA GRID vGPUs supported by this host or at least one host on this cluster. If unset, this host or cluster does not support NVIDIA GRID vGPUs and cannot be used for desktop creation with NVIDIA GRID vGPU support enabled. | [optional] 

## Methods

### NewHostOrClusterDetails

`func NewHostOrClusterDetails() *HostOrClusterDetails`

NewHostOrClusterDetails instantiates a new HostOrClusterDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOrClusterDetailsWithDefaults

`func NewHostOrClusterDetailsWithDefaults() *HostOrClusterDetails`

NewHostOrClusterDetailsWithDefaults instantiates a new HostOrClusterDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *HostOrClusterDetails) GetCluster() bool`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HostOrClusterDetails) GetClusterOk() (*bool, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HostOrClusterDetails) SetCluster(v bool)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HostOrClusterDetails) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatacenterId

`func (o *HostOrClusterDetails) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *HostOrClusterDetails) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *HostOrClusterDetails) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *HostOrClusterDetails) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *HostOrClusterDetails) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *HostOrClusterDetails) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *HostOrClusterDetails) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *HostOrClusterDetails) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetName

`func (o *HostOrClusterDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostOrClusterDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostOrClusterDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostOrClusterDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *HostOrClusterDetails) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HostOrClusterDetails) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HostOrClusterDetails) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HostOrClusterDetails) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVcenterId

`func (o *HostOrClusterDetails) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *HostOrClusterDetails) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *HostOrClusterDetails) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *HostOrClusterDetails) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.

### GetVgpuTypes

`func (o *HostOrClusterDetails) GetVgpuTypes() []string`

GetVgpuTypes returns the VgpuTypes field if non-nil, zero value otherwise.

### GetVgpuTypesOk

`func (o *HostOrClusterDetails) GetVgpuTypesOk() (*[]string, bool)`

GetVgpuTypesOk returns a tuple with the VgpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuTypes

`func (o *HostOrClusterDetails) SetVgpuTypes(v []string)`

SetVgpuTypes sets VgpuTypes field to given value.

### HasVgpuTypes

`func (o *HostOrClusterDetails) HasVgpuTypes() bool`

HasVgpuTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


