# DatastoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityMb** | Pointer to **int64** | Maximum capacity of this datastore, in MB. | [optional] 
**DatacenterId** | Pointer to **string** | Datacenter id for this datastore. | [optional] 
**DiskType** | Pointer to **string** | Disk type of the datastore. * SSD: Solid State Drive disk type. * NON_SSD: NON-Solid State Drive disk type. * UNKNOWN: Unknown disk type. * NON_VMFS: NON-VMFS disk type. | [optional] 
**FileSystemType** | Pointer to **string** | File system type of the datastore. * VMFS: Virtual Machine File System. * NFS: Network File System. * VSAN: vSAN File System. * VVOL: Virtual Volumes. | [optional] 
**FreeSpaceMb** | Pointer to **int64** | Available capacity of this datastore, in MB. | [optional] 
**HostOrClusterId** | Pointer to **string** | Host or Cluster id for this datastore. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the datastore. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this Datastore from being used in desktop pool/farm. | [optional] 
**LocalDatastore** | Pointer to **bool** | Indicates if this datastore is local to a single host. | [optional] 
**Name** | Pointer to **string** | Datastore name. | [optional] 
**NumberOfVms** | Pointer to **int32** | Indicates the number of virtual machines the datastore has for desktop pool/farm when applicable | [optional] 
**Path** | Pointer to **string** | Datastore path. | [optional] 
**VcenterId** | Pointer to **string** | Virtual Center id for this datastore. | [optional] 
**VmfsMajorVersion** | Pointer to **string** | The VMFS major version number. | [optional] 

## Methods

### NewDatastoreInfo

`func NewDatastoreInfo() *DatastoreInfo`

NewDatastoreInfo instantiates a new DatastoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreInfoWithDefaults

`func NewDatastoreInfoWithDefaults() *DatastoreInfo`

NewDatastoreInfoWithDefaults instantiates a new DatastoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityMb

`func (o *DatastoreInfo) GetCapacityMb() int64`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *DatastoreInfo) GetCapacityMbOk() (*int64, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *DatastoreInfo) SetCapacityMb(v int64)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *DatastoreInfo) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetDatacenterId

`func (o *DatastoreInfo) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *DatastoreInfo) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *DatastoreInfo) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *DatastoreInfo) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetDiskType

`func (o *DatastoreInfo) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *DatastoreInfo) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *DatastoreInfo) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *DatastoreInfo) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetFileSystemType

`func (o *DatastoreInfo) GetFileSystemType() string`

GetFileSystemType returns the FileSystemType field if non-nil, zero value otherwise.

### GetFileSystemTypeOk

`func (o *DatastoreInfo) GetFileSystemTypeOk() (*string, bool)`

GetFileSystemTypeOk returns a tuple with the FileSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemType

`func (o *DatastoreInfo) SetFileSystemType(v string)`

SetFileSystemType sets FileSystemType field to given value.

### HasFileSystemType

`func (o *DatastoreInfo) HasFileSystemType() bool`

HasFileSystemType returns a boolean if a field has been set.

### GetFreeSpaceMb

`func (o *DatastoreInfo) GetFreeSpaceMb() int64`

GetFreeSpaceMb returns the FreeSpaceMb field if non-nil, zero value otherwise.

### GetFreeSpaceMbOk

`func (o *DatastoreInfo) GetFreeSpaceMbOk() (*int64, bool)`

GetFreeSpaceMbOk returns a tuple with the FreeSpaceMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceMb

`func (o *DatastoreInfo) SetFreeSpaceMb(v int64)`

SetFreeSpaceMb sets FreeSpaceMb field to given value.

### HasFreeSpaceMb

`func (o *DatastoreInfo) HasFreeSpaceMb() bool`

HasFreeSpaceMb returns a boolean if a field has been set.

### GetHostOrClusterId

`func (o *DatastoreInfo) GetHostOrClusterId() string`

GetHostOrClusterId returns the HostOrClusterId field if non-nil, zero value otherwise.

### GetHostOrClusterIdOk

`func (o *DatastoreInfo) GetHostOrClusterIdOk() (*string, bool)`

GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrClusterId

`func (o *DatastoreInfo) SetHostOrClusterId(v string)`

SetHostOrClusterId sets HostOrClusterId field to given value.

### HasHostOrClusterId

`func (o *DatastoreInfo) HasHostOrClusterId() bool`

HasHostOrClusterId returns a boolean if a field has been set.

### GetId

`func (o *DatastoreInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatastoreInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatastoreInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatastoreInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *DatastoreInfo) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *DatastoreInfo) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *DatastoreInfo) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *DatastoreInfo) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetLocalDatastore

`func (o *DatastoreInfo) GetLocalDatastore() bool`

GetLocalDatastore returns the LocalDatastore field if non-nil, zero value otherwise.

### GetLocalDatastoreOk

`func (o *DatastoreInfo) GetLocalDatastoreOk() (*bool, bool)`

GetLocalDatastoreOk returns a tuple with the LocalDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDatastore

`func (o *DatastoreInfo) SetLocalDatastore(v bool)`

SetLocalDatastore sets LocalDatastore field to given value.

### HasLocalDatastore

`func (o *DatastoreInfo) HasLocalDatastore() bool`

HasLocalDatastore returns a boolean if a field has been set.

### GetName

`func (o *DatastoreInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatastoreInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatastoreInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatastoreInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfVms

`func (o *DatastoreInfo) GetNumberOfVms() int32`

GetNumberOfVms returns the NumberOfVms field if non-nil, zero value otherwise.

### GetNumberOfVmsOk

`func (o *DatastoreInfo) GetNumberOfVmsOk() (*int32, bool)`

GetNumberOfVmsOk returns a tuple with the NumberOfVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVms

`func (o *DatastoreInfo) SetNumberOfVms(v int32)`

SetNumberOfVms sets NumberOfVms field to given value.

### HasNumberOfVms

`func (o *DatastoreInfo) HasNumberOfVms() bool`

HasNumberOfVms returns a boolean if a field has been set.

### GetPath

`func (o *DatastoreInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DatastoreInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DatastoreInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DatastoreInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVcenterId

`func (o *DatastoreInfo) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *DatastoreInfo) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *DatastoreInfo) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *DatastoreInfo) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.

### GetVmfsMajorVersion

`func (o *DatastoreInfo) GetVmfsMajorVersion() string`

GetVmfsMajorVersion returns the VmfsMajorVersion field if non-nil, zero value otherwise.

### GetVmfsMajorVersionOk

`func (o *DatastoreInfo) GetVmfsMajorVersionOk() (*string, bool)`

GetVmfsMajorVersionOk returns a tuple with the VmfsMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmfsMajorVersion

`func (o *DatastoreInfo) SetVmfsMajorVersion(v string)`

SetVmfsMajorVersion sets VmfsMajorVersion field to given value.

### HasVmfsMajorVersion

`func (o *DatastoreInfo) HasVmfsMajorVersion() bool`

HasVmfsMajorVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


