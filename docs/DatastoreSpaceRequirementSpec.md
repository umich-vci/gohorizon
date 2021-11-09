# DatastoreSpaceRequirementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseSnapshotId** | Pointer to **string** | Parent VM snapshot ID. Must be set if source is INSTANT_CLONE. | [optional] 
**BaseVmId** | Pointer to **string** | Parent VM ID. Must be set if source is INSTANT_CLONE. | [optional] 
**Id** | Pointer to **string** | Id of inventory resource for which space requirement is to be found. Can be desktop pool or farm id. | [optional] 
**PoolSize** | **int32** | Desired size of the desktop pool or farm. | 
**Source** | **string** | Source or provisioning type of machines. * FULL_CLONE: Virtual Machines created from a vCenter Server template. * INSTANT_CLONE: Virtual Machines created by instant clone engine. | 
**Type** | **string** | Type of inventory resource for which space requirement is to be found. * DESKTOP_POOL: Desktop pool inventory resource. * FARM: Farm inventory resource. | 
**UseSeparateReplicaAndOsDisk** | Pointer to **bool** | Indicates whether separate datastores are to be used for OS and replica disks. Will be ignored if source is FULL_CLONE or vSAN is to be configured. Default value is false. | [optional] 
**UseVsan** | Pointer to **bool** | Indicates whether vSAN is to be configured for the desktop pool or farm. Default value is false. vSAN should be configured if set to true. | [optional] 
**UserAssignment** | Pointer to **string** | User assignment of the desktop pool. Will be ignored if type is FARM. Default value is FLOATING. * DEDICATED: Dedicated user assignment. * FLOATING: Floating user assignment. | [optional] 
**VcenterId** | **string** | ID of virtual center where parent VM or master image is present. | 
**VmTemplateId** | Pointer to **string** | VM template ID. Must be set if source is FULL_CLONE. | [optional] 

## Methods

### NewDatastoreSpaceRequirementSpec

`func NewDatastoreSpaceRequirementSpec(poolSize int32, source string, type_ string, vcenterId string, ) *DatastoreSpaceRequirementSpec`

NewDatastoreSpaceRequirementSpec instantiates a new DatastoreSpaceRequirementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreSpaceRequirementSpecWithDefaults

`func NewDatastoreSpaceRequirementSpecWithDefaults() *DatastoreSpaceRequirementSpec`

NewDatastoreSpaceRequirementSpecWithDefaults instantiates a new DatastoreSpaceRequirementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseSnapshotId

`func (o *DatastoreSpaceRequirementSpec) GetBaseSnapshotId() string`

GetBaseSnapshotId returns the BaseSnapshotId field if non-nil, zero value otherwise.

### GetBaseSnapshotIdOk

`func (o *DatastoreSpaceRequirementSpec) GetBaseSnapshotIdOk() (*string, bool)`

GetBaseSnapshotIdOk returns a tuple with the BaseSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotId

`func (o *DatastoreSpaceRequirementSpec) SetBaseSnapshotId(v string)`

SetBaseSnapshotId sets BaseSnapshotId field to given value.

### HasBaseSnapshotId

`func (o *DatastoreSpaceRequirementSpec) HasBaseSnapshotId() bool`

HasBaseSnapshotId returns a boolean if a field has been set.

### GetBaseVmId

`func (o *DatastoreSpaceRequirementSpec) GetBaseVmId() string`

GetBaseVmId returns the BaseVmId field if non-nil, zero value otherwise.

### GetBaseVmIdOk

`func (o *DatastoreSpaceRequirementSpec) GetBaseVmIdOk() (*string, bool)`

GetBaseVmIdOk returns a tuple with the BaseVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVmId

`func (o *DatastoreSpaceRequirementSpec) SetBaseVmId(v string)`

SetBaseVmId sets BaseVmId field to given value.

### HasBaseVmId

`func (o *DatastoreSpaceRequirementSpec) HasBaseVmId() bool`

HasBaseVmId returns a boolean if a field has been set.

### GetId

`func (o *DatastoreSpaceRequirementSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatastoreSpaceRequirementSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatastoreSpaceRequirementSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatastoreSpaceRequirementSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPoolSize

`func (o *DatastoreSpaceRequirementSpec) GetPoolSize() int32`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *DatastoreSpaceRequirementSpec) GetPoolSizeOk() (*int32, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *DatastoreSpaceRequirementSpec) SetPoolSize(v int32)`

SetPoolSize sets PoolSize field to given value.


### GetSource

`func (o *DatastoreSpaceRequirementSpec) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DatastoreSpaceRequirementSpec) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DatastoreSpaceRequirementSpec) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *DatastoreSpaceRequirementSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatastoreSpaceRequirementSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatastoreSpaceRequirementSpec) SetType(v string)`

SetType sets Type field to given value.


### GetUseSeparateReplicaAndOsDisk

`func (o *DatastoreSpaceRequirementSpec) GetUseSeparateReplicaAndOsDisk() bool`

GetUseSeparateReplicaAndOsDisk returns the UseSeparateReplicaAndOsDisk field if non-nil, zero value otherwise.

### GetUseSeparateReplicaAndOsDiskOk

`func (o *DatastoreSpaceRequirementSpec) GetUseSeparateReplicaAndOsDiskOk() (*bool, bool)`

GetUseSeparateReplicaAndOsDiskOk returns a tuple with the UseSeparateReplicaAndOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSeparateReplicaAndOsDisk

`func (o *DatastoreSpaceRequirementSpec) SetUseSeparateReplicaAndOsDisk(v bool)`

SetUseSeparateReplicaAndOsDisk sets UseSeparateReplicaAndOsDisk field to given value.

### HasUseSeparateReplicaAndOsDisk

`func (o *DatastoreSpaceRequirementSpec) HasUseSeparateReplicaAndOsDisk() bool`

HasUseSeparateReplicaAndOsDisk returns a boolean if a field has been set.

### GetUseVsan

`func (o *DatastoreSpaceRequirementSpec) GetUseVsan() bool`

GetUseVsan returns the UseVsan field if non-nil, zero value otherwise.

### GetUseVsanOk

`func (o *DatastoreSpaceRequirementSpec) GetUseVsanOk() (*bool, bool)`

GetUseVsanOk returns a tuple with the UseVsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVsan

`func (o *DatastoreSpaceRequirementSpec) SetUseVsan(v bool)`

SetUseVsan sets UseVsan field to given value.

### HasUseVsan

`func (o *DatastoreSpaceRequirementSpec) HasUseVsan() bool`

HasUseVsan returns a boolean if a field has been set.

### GetUserAssignment

`func (o *DatastoreSpaceRequirementSpec) GetUserAssignment() string`

GetUserAssignment returns the UserAssignment field if non-nil, zero value otherwise.

### GetUserAssignmentOk

`func (o *DatastoreSpaceRequirementSpec) GetUserAssignmentOk() (*string, bool)`

GetUserAssignmentOk returns a tuple with the UserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignment

`func (o *DatastoreSpaceRequirementSpec) SetUserAssignment(v string)`

SetUserAssignment sets UserAssignment field to given value.

### HasUserAssignment

`func (o *DatastoreSpaceRequirementSpec) HasUserAssignment() bool`

HasUserAssignment returns a boolean if a field has been set.

### GetVcenterId

`func (o *DatastoreSpaceRequirementSpec) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *DatastoreSpaceRequirementSpec) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *DatastoreSpaceRequirementSpec) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.


### GetVmTemplateId

`func (o *DatastoreSpaceRequirementSpec) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *DatastoreSpaceRequirementSpec) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *DatastoreSpaceRequirementSpec) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.

### HasVmTemplateId

`func (o *DatastoreSpaceRequirementSpec) HasVmTemplateId() bool`

HasVmTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


