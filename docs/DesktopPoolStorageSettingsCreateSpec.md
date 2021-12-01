# DesktopPoolStorageSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | [**[]DesktopPoolDatastoreSettingsCreateSpec**](DesktopPoolDatastoreSettingsCreateSpec.md) | Datastores to store the machine. | 
**ReclaimVmDiskSpace** | Pointer to **bool** | With vSphere 5.x, virtual machines can be configured to use a space efficient disk format that supports reclamation of unused diskspace (such as deleted files). This option reclaims unused diskspace on each virtual machine. The operation is initiated when an estimate of used disk space exceeds the specified threshold. &lt;br&gt; Default value is false. | [optional] 
**ReclamationThresholdMb** | Pointer to **int64** | Initiate reclamation when unused space on virtual machine exceeds the threshold in MB. &lt;br&gt; This property is required if reclaim_vm_disk_space is set to true. | [optional] 
**ReplicaDiskDatastoreId** | Pointer to **string** | Datastore to store replica disks for instant clone machines. &lt;br&gt; This property is required if use_separate_datastores_replica_and_os_disks is set to true. &lt;br&gt; | [optional] 
**UseSeparateDatastoresReplicaAndOsDisks** | Pointer to **bool** | Indicates whether to use separate datastores for replica and OS disks. &lt;br&gt; Default value is false. | [optional] 
**UseVsan** | Pointer to **bool** | Indicates whether to use vSphere vSAN. &lt;br&gt; Default value is false. | [optional] 

## Methods

### NewDesktopPoolStorageSettingsCreateSpec

`func NewDesktopPoolStorageSettingsCreateSpec(datastores []DesktopPoolDatastoreSettingsCreateSpec, ) *DesktopPoolStorageSettingsCreateSpec`

NewDesktopPoolStorageSettingsCreateSpec instantiates a new DesktopPoolStorageSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolStorageSettingsCreateSpecWithDefaults

`func NewDesktopPoolStorageSettingsCreateSpecWithDefaults() *DesktopPoolStorageSettingsCreateSpec`

NewDesktopPoolStorageSettingsCreateSpecWithDefaults instantiates a new DesktopPoolStorageSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *DesktopPoolStorageSettingsCreateSpec) GetDatastores() []DesktopPoolDatastoreSettingsCreateSpec`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *DesktopPoolStorageSettingsCreateSpec) GetDatastoresOk() (*[]DesktopPoolDatastoreSettingsCreateSpec, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *DesktopPoolStorageSettingsCreateSpec) SetDatastores(v []DesktopPoolDatastoreSettingsCreateSpec)`

SetDatastores sets Datastores field to given value.


### GetReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettingsCreateSpec) GetReclaimVmDiskSpace() bool`

GetReclaimVmDiskSpace returns the ReclaimVmDiskSpace field if non-nil, zero value otherwise.

### GetReclaimVmDiskSpaceOk

`func (o *DesktopPoolStorageSettingsCreateSpec) GetReclaimVmDiskSpaceOk() (*bool, bool)`

GetReclaimVmDiskSpaceOk returns a tuple with the ReclaimVmDiskSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettingsCreateSpec) SetReclaimVmDiskSpace(v bool)`

SetReclaimVmDiskSpace sets ReclaimVmDiskSpace field to given value.

### HasReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettingsCreateSpec) HasReclaimVmDiskSpace() bool`

HasReclaimVmDiskSpace returns a boolean if a field has been set.

### GetReclamationThresholdMb

`func (o *DesktopPoolStorageSettingsCreateSpec) GetReclamationThresholdMb() int64`

GetReclamationThresholdMb returns the ReclamationThresholdMb field if non-nil, zero value otherwise.

### GetReclamationThresholdMbOk

`func (o *DesktopPoolStorageSettingsCreateSpec) GetReclamationThresholdMbOk() (*int64, bool)`

GetReclamationThresholdMbOk returns a tuple with the ReclamationThresholdMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclamationThresholdMb

`func (o *DesktopPoolStorageSettingsCreateSpec) SetReclamationThresholdMb(v int64)`

SetReclamationThresholdMb sets ReclamationThresholdMb field to given value.

### HasReclamationThresholdMb

`func (o *DesktopPoolStorageSettingsCreateSpec) HasReclamationThresholdMb() bool`

HasReclamationThresholdMb returns a boolean if a field has been set.

### GetReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettingsCreateSpec) GetReplicaDiskDatastoreId() string`

GetReplicaDiskDatastoreId returns the ReplicaDiskDatastoreId field if non-nil, zero value otherwise.

### GetReplicaDiskDatastoreIdOk

`func (o *DesktopPoolStorageSettingsCreateSpec) GetReplicaDiskDatastoreIdOk() (*string, bool)`

GetReplicaDiskDatastoreIdOk returns a tuple with the ReplicaDiskDatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettingsCreateSpec) SetReplicaDiskDatastoreId(v string)`

SetReplicaDiskDatastoreId sets ReplicaDiskDatastoreId field to given value.

### HasReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettingsCreateSpec) HasReplicaDiskDatastoreId() bool`

HasReplicaDiskDatastoreId returns a boolean if a field has been set.

### GetUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettingsCreateSpec) GetUseSeparateDatastoresReplicaAndOsDisks() bool`

GetUseSeparateDatastoresReplicaAndOsDisks returns the UseSeparateDatastoresReplicaAndOsDisks field if non-nil, zero value otherwise.

### GetUseSeparateDatastoresReplicaAndOsDisksOk

`func (o *DesktopPoolStorageSettingsCreateSpec) GetUseSeparateDatastoresReplicaAndOsDisksOk() (*bool, bool)`

GetUseSeparateDatastoresReplicaAndOsDisksOk returns a tuple with the UseSeparateDatastoresReplicaAndOsDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettingsCreateSpec) SetUseSeparateDatastoresReplicaAndOsDisks(v bool)`

SetUseSeparateDatastoresReplicaAndOsDisks sets UseSeparateDatastoresReplicaAndOsDisks field to given value.

### HasUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettingsCreateSpec) HasUseSeparateDatastoresReplicaAndOsDisks() bool`

HasUseSeparateDatastoresReplicaAndOsDisks returns a boolean if a field has been set.

### GetUseVsan

`func (o *DesktopPoolStorageSettingsCreateSpec) GetUseVsan() bool`

GetUseVsan returns the UseVsan field if non-nil, zero value otherwise.

### GetUseVsanOk

`func (o *DesktopPoolStorageSettingsCreateSpec) GetUseVsanOk() (*bool, bool)`

GetUseVsanOk returns a tuple with the UseVsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVsan

`func (o *DesktopPoolStorageSettingsCreateSpec) SetUseVsan(v bool)`

SetUseVsan sets UseVsan field to given value.

### HasUseVsan

`func (o *DesktopPoolStorageSettingsCreateSpec) HasUseVsan() bool`

HasUseVsan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


