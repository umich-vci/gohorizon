# FarmStorageSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | Pointer to [**[]FarmDatastoreSettingsInfo**](FarmDatastoreSettingsInfo.md) | List of IDs of the datastore used to store the RDS Server. | [optional] 
**ReplicaDiskDatastoreId** | Pointer to **string** | Datastore to store replica disks for instant clone machines. This is set when use_separate_datastores_replica_and_os_disks is true. | [optional] 
**UseSeparateDatastoresReplicaAndOsDisks** | Pointer to **bool** | Indicates whether to use separate datastores for replica and OS disks. | [optional] 
**UseViewStorageAccelerator** | Pointer to **bool** | Indicates whether to use view storage accelerator. | [optional] 
**UseVsan** | Pointer to **bool** | Indicates whether to use vSphere VSAN. | [optional] 

## Methods

### NewFarmStorageSettingsInfo

`func NewFarmStorageSettingsInfo() *FarmStorageSettingsInfo`

NewFarmStorageSettingsInfo instantiates a new FarmStorageSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmStorageSettingsInfoWithDefaults

`func NewFarmStorageSettingsInfoWithDefaults() *FarmStorageSettingsInfo`

NewFarmStorageSettingsInfoWithDefaults instantiates a new FarmStorageSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *FarmStorageSettingsInfo) GetDatastores() []FarmDatastoreSettingsInfo`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *FarmStorageSettingsInfo) GetDatastoresOk() (*[]FarmDatastoreSettingsInfo, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *FarmStorageSettingsInfo) SetDatastores(v []FarmDatastoreSettingsInfo)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *FarmStorageSettingsInfo) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetReplicaDiskDatastoreId

`func (o *FarmStorageSettingsInfo) GetReplicaDiskDatastoreId() string`

GetReplicaDiskDatastoreId returns the ReplicaDiskDatastoreId field if non-nil, zero value otherwise.

### GetReplicaDiskDatastoreIdOk

`func (o *FarmStorageSettingsInfo) GetReplicaDiskDatastoreIdOk() (*string, bool)`

GetReplicaDiskDatastoreIdOk returns a tuple with the ReplicaDiskDatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaDiskDatastoreId

`func (o *FarmStorageSettingsInfo) SetReplicaDiskDatastoreId(v string)`

SetReplicaDiskDatastoreId sets ReplicaDiskDatastoreId field to given value.

### HasReplicaDiskDatastoreId

`func (o *FarmStorageSettingsInfo) HasReplicaDiskDatastoreId() bool`

HasReplicaDiskDatastoreId returns a boolean if a field has been set.

### GetUseSeparateDatastoresReplicaAndOsDisks

`func (o *FarmStorageSettingsInfo) GetUseSeparateDatastoresReplicaAndOsDisks() bool`

GetUseSeparateDatastoresReplicaAndOsDisks returns the UseSeparateDatastoresReplicaAndOsDisks field if non-nil, zero value otherwise.

### GetUseSeparateDatastoresReplicaAndOsDisksOk

`func (o *FarmStorageSettingsInfo) GetUseSeparateDatastoresReplicaAndOsDisksOk() (*bool, bool)`

GetUseSeparateDatastoresReplicaAndOsDisksOk returns a tuple with the UseSeparateDatastoresReplicaAndOsDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSeparateDatastoresReplicaAndOsDisks

`func (o *FarmStorageSettingsInfo) SetUseSeparateDatastoresReplicaAndOsDisks(v bool)`

SetUseSeparateDatastoresReplicaAndOsDisks sets UseSeparateDatastoresReplicaAndOsDisks field to given value.

### HasUseSeparateDatastoresReplicaAndOsDisks

`func (o *FarmStorageSettingsInfo) HasUseSeparateDatastoresReplicaAndOsDisks() bool`

HasUseSeparateDatastoresReplicaAndOsDisks returns a boolean if a field has been set.

### GetUseViewStorageAccelerator

`func (o *FarmStorageSettingsInfo) GetUseViewStorageAccelerator() bool`

GetUseViewStorageAccelerator returns the UseViewStorageAccelerator field if non-nil, zero value otherwise.

### GetUseViewStorageAcceleratorOk

`func (o *FarmStorageSettingsInfo) GetUseViewStorageAcceleratorOk() (*bool, bool)`

GetUseViewStorageAcceleratorOk returns a tuple with the UseViewStorageAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseViewStorageAccelerator

`func (o *FarmStorageSettingsInfo) SetUseViewStorageAccelerator(v bool)`

SetUseViewStorageAccelerator sets UseViewStorageAccelerator field to given value.

### HasUseViewStorageAccelerator

`func (o *FarmStorageSettingsInfo) HasUseViewStorageAccelerator() bool`

HasUseViewStorageAccelerator returns a boolean if a field has been set.

### GetUseVsan

`func (o *FarmStorageSettingsInfo) GetUseVsan() bool`

GetUseVsan returns the UseVsan field if non-nil, zero value otherwise.

### GetUseVsanOk

`func (o *FarmStorageSettingsInfo) GetUseVsanOk() (*bool, bool)`

GetUseVsanOk returns a tuple with the UseVsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVsan

`func (o *FarmStorageSettingsInfo) SetUseVsan(v bool)`

SetUseVsan sets UseVsan field to given value.

### HasUseVsan

`func (o *FarmStorageSettingsInfo) HasUseVsan() bool`

HasUseVsan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


