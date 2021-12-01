# DesktopPoolStorageSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | [**[]DesktopPoolDatastoreSettingsUpdateSpec**](DesktopPoolDatastoreSettingsUpdateSpec.md) | Datastores to store the machine (or the OS disk using other options for instant clone machine storage) | 
**ReclaimVmDiskSpace** | Pointer to **bool** | With vSphere 5.x, virtual machines can be configured to use a space efficient disk format that supports reclamation of unused disk space (such as deleted files). This option reclaims unused disk space on each virtual machine. The operation is initiated when an estimate of used disk space exceeds the specified threshold. | [optional] 
**ReclamationThresholdMb** | Pointer to **int64** | Initiate reclamation when unused space on virtual machine exceeds the threshold in MB. Default value is 1. | [optional] 
**ReplicaDiskDatastoreId** | Pointer to **string** | Datastore to store replica disks for instant clone machines. | [optional] 
**UseSeparateDatastoresReplicaAndOsDisks** | Pointer to **bool** | Indicates whether to use separate datastores for replica and OS disks. | [optional] 
**UseVsan** | **bool** | Indicates whether to use vSphere vSAN. | 

## Methods

### NewDesktopPoolStorageSettingsUpdateSpec

`func NewDesktopPoolStorageSettingsUpdateSpec(datastores []DesktopPoolDatastoreSettingsUpdateSpec, useVsan bool, ) *DesktopPoolStorageSettingsUpdateSpec`

NewDesktopPoolStorageSettingsUpdateSpec instantiates a new DesktopPoolStorageSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolStorageSettingsUpdateSpecWithDefaults

`func NewDesktopPoolStorageSettingsUpdateSpecWithDefaults() *DesktopPoolStorageSettingsUpdateSpec`

NewDesktopPoolStorageSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolStorageSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetDatastores() []DesktopPoolDatastoreSettingsUpdateSpec`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetDatastoresOk() (*[]DesktopPoolDatastoreSettingsUpdateSpec, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *DesktopPoolStorageSettingsUpdateSpec) SetDatastores(v []DesktopPoolDatastoreSettingsUpdateSpec)`

SetDatastores sets Datastores field to given value.


### GetReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetReclaimVmDiskSpace() bool`

GetReclaimVmDiskSpace returns the ReclaimVmDiskSpace field if non-nil, zero value otherwise.

### GetReclaimVmDiskSpaceOk

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetReclaimVmDiskSpaceOk() (*bool, bool)`

GetReclaimVmDiskSpaceOk returns a tuple with the ReclaimVmDiskSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettingsUpdateSpec) SetReclaimVmDiskSpace(v bool)`

SetReclaimVmDiskSpace sets ReclaimVmDiskSpace field to given value.

### HasReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettingsUpdateSpec) HasReclaimVmDiskSpace() bool`

HasReclaimVmDiskSpace returns a boolean if a field has been set.

### GetReclamationThresholdMb

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetReclamationThresholdMb() int64`

GetReclamationThresholdMb returns the ReclamationThresholdMb field if non-nil, zero value otherwise.

### GetReclamationThresholdMbOk

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetReclamationThresholdMbOk() (*int64, bool)`

GetReclamationThresholdMbOk returns a tuple with the ReclamationThresholdMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclamationThresholdMb

`func (o *DesktopPoolStorageSettingsUpdateSpec) SetReclamationThresholdMb(v int64)`

SetReclamationThresholdMb sets ReclamationThresholdMb field to given value.

### HasReclamationThresholdMb

`func (o *DesktopPoolStorageSettingsUpdateSpec) HasReclamationThresholdMb() bool`

HasReclamationThresholdMb returns a boolean if a field has been set.

### GetReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetReplicaDiskDatastoreId() string`

GetReplicaDiskDatastoreId returns the ReplicaDiskDatastoreId field if non-nil, zero value otherwise.

### GetReplicaDiskDatastoreIdOk

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetReplicaDiskDatastoreIdOk() (*string, bool)`

GetReplicaDiskDatastoreIdOk returns a tuple with the ReplicaDiskDatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettingsUpdateSpec) SetReplicaDiskDatastoreId(v string)`

SetReplicaDiskDatastoreId sets ReplicaDiskDatastoreId field to given value.

### HasReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettingsUpdateSpec) HasReplicaDiskDatastoreId() bool`

HasReplicaDiskDatastoreId returns a boolean if a field has been set.

### GetUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetUseSeparateDatastoresReplicaAndOsDisks() bool`

GetUseSeparateDatastoresReplicaAndOsDisks returns the UseSeparateDatastoresReplicaAndOsDisks field if non-nil, zero value otherwise.

### GetUseSeparateDatastoresReplicaAndOsDisksOk

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetUseSeparateDatastoresReplicaAndOsDisksOk() (*bool, bool)`

GetUseSeparateDatastoresReplicaAndOsDisksOk returns a tuple with the UseSeparateDatastoresReplicaAndOsDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettingsUpdateSpec) SetUseSeparateDatastoresReplicaAndOsDisks(v bool)`

SetUseSeparateDatastoresReplicaAndOsDisks sets UseSeparateDatastoresReplicaAndOsDisks field to given value.

### HasUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettingsUpdateSpec) HasUseSeparateDatastoresReplicaAndOsDisks() bool`

HasUseSeparateDatastoresReplicaAndOsDisks returns a boolean if a field has been set.

### GetUseVsan

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetUseVsan() bool`

GetUseVsan returns the UseVsan field if non-nil, zero value otherwise.

### GetUseVsanOk

`func (o *DesktopPoolStorageSettingsUpdateSpec) GetUseVsanOk() (*bool, bool)`

GetUseVsanOk returns a tuple with the UseVsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVsan

`func (o *DesktopPoolStorageSettingsUpdateSpec) SetUseVsan(v bool)`

SetUseVsan sets UseVsan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


