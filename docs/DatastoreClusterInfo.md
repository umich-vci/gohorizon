# DatastoreClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityMb** | Pointer to **int64** | Maximum capacity of this datastore cluster, in MB. | [optional] 
**DatacenterId** | Pointer to **string** | ID of the datacenter for this datastore cluster. | [optional] 
**DatastoreIds** | Pointer to **[]string** | IDs of datastores which are a part of this datastore cluster. | [optional] 
**FreeSpaceMb** | Pointer to **int64** | Available capacity of this datastore cluster, in MB. | [optional] 
**HostOrClusterId** | Pointer to **string** | ID of the host or cluster for this datastore cluster. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this datastore cluster. | [optional] 
**Name** | Pointer to **string** | Datastore cluster name. | [optional] 
**Path** | Pointer to **string** | Datastore cluster path. | [optional] 
**VcenterId** | Pointer to **string** | ID of the virtual center for this datastore cluster. | [optional] 

## Methods

### NewDatastoreClusterInfo

`func NewDatastoreClusterInfo() *DatastoreClusterInfo`

NewDatastoreClusterInfo instantiates a new DatastoreClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreClusterInfoWithDefaults

`func NewDatastoreClusterInfoWithDefaults() *DatastoreClusterInfo`

NewDatastoreClusterInfoWithDefaults instantiates a new DatastoreClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityMb

`func (o *DatastoreClusterInfo) GetCapacityMb() int64`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *DatastoreClusterInfo) GetCapacityMbOk() (*int64, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *DatastoreClusterInfo) SetCapacityMb(v int64)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *DatastoreClusterInfo) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetDatacenterId

`func (o *DatastoreClusterInfo) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *DatastoreClusterInfo) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *DatastoreClusterInfo) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *DatastoreClusterInfo) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetDatastoreIds

`func (o *DatastoreClusterInfo) GetDatastoreIds() []string`

GetDatastoreIds returns the DatastoreIds field if non-nil, zero value otherwise.

### GetDatastoreIdsOk

`func (o *DatastoreClusterInfo) GetDatastoreIdsOk() (*[]string, bool)`

GetDatastoreIdsOk returns a tuple with the DatastoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreIds

`func (o *DatastoreClusterInfo) SetDatastoreIds(v []string)`

SetDatastoreIds sets DatastoreIds field to given value.

### HasDatastoreIds

`func (o *DatastoreClusterInfo) HasDatastoreIds() bool`

HasDatastoreIds returns a boolean if a field has been set.

### GetFreeSpaceMb

`func (o *DatastoreClusterInfo) GetFreeSpaceMb() int64`

GetFreeSpaceMb returns the FreeSpaceMb field if non-nil, zero value otherwise.

### GetFreeSpaceMbOk

`func (o *DatastoreClusterInfo) GetFreeSpaceMbOk() (*int64, bool)`

GetFreeSpaceMbOk returns a tuple with the FreeSpaceMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceMb

`func (o *DatastoreClusterInfo) SetFreeSpaceMb(v int64)`

SetFreeSpaceMb sets FreeSpaceMb field to given value.

### HasFreeSpaceMb

`func (o *DatastoreClusterInfo) HasFreeSpaceMb() bool`

HasFreeSpaceMb returns a boolean if a field has been set.

### GetHostOrClusterId

`func (o *DatastoreClusterInfo) GetHostOrClusterId() string`

GetHostOrClusterId returns the HostOrClusterId field if non-nil, zero value otherwise.

### GetHostOrClusterIdOk

`func (o *DatastoreClusterInfo) GetHostOrClusterIdOk() (*string, bool)`

GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrClusterId

`func (o *DatastoreClusterInfo) SetHostOrClusterId(v string)`

SetHostOrClusterId sets HostOrClusterId field to given value.

### HasHostOrClusterId

`func (o *DatastoreClusterInfo) HasHostOrClusterId() bool`

HasHostOrClusterId returns a boolean if a field has been set.

### GetId

`func (o *DatastoreClusterInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatastoreClusterInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatastoreClusterInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatastoreClusterInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DatastoreClusterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatastoreClusterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatastoreClusterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatastoreClusterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *DatastoreClusterInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DatastoreClusterInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DatastoreClusterInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DatastoreClusterInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVcenterId

`func (o *DatastoreClusterInfo) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *DatastoreClusterInfo) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *DatastoreClusterInfo) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *DatastoreClusterInfo) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


