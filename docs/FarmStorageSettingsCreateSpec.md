# FarmStorageSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | [**[]FarmDatastoreSettingsCreateSpec**](FarmDatastoreSettingsCreateSpec.md) | List of IDs of the datastore used to store the RDS Server. | 
**ReplicaDiskDatastoreId** | Pointer to **string** | Datastore to store replica disks for instant clone machines. This is required if use_separate_datastores_replica_and_os_disks is set to true. | [optional] 
**UseSeparateDatastoresReplicaAndOsDisks** | Pointer to **bool** | Indicates whether to use separate datastores for replica and OS disks. Default value is false. | [optional] 
**UseViewStorageAccelerator** | Pointer to **bool** | Indicates whether to use view storage accelerator. Default value is false. | [optional] 
**UseVsan** | Pointer to **bool** | Indicates whether to use vSphere VSAN. Default value is false. | [optional] 

## Methods

### NewFarmStorageSettingsCreateSpec

`func NewFarmStorageSettingsCreateSpec(datastores []FarmDatastoreSettingsCreateSpec, ) *FarmStorageSettingsCreateSpec`

NewFarmStorageSettingsCreateSpec instantiates a new FarmStorageSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmStorageSettingsCreateSpecWithDefaults

`func NewFarmStorageSettingsCreateSpecWithDefaults() *FarmStorageSettingsCreateSpec`

NewFarmStorageSettingsCreateSpecWithDefaults instantiates a new FarmStorageSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *FarmStorageSettingsCreateSpec) GetDatastores() []FarmDatastoreSettingsCreateSpec`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *FarmStorageSettingsCreateSpec) GetDatastoresOk() (*[]FarmDatastoreSettingsCreateSpec, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *FarmStorageSettingsCreateSpec) SetDatastores(v []FarmDatastoreSettingsCreateSpec)`

SetDatastores sets Datastores field to given value.


### GetReplicaDiskDatastoreId

`func (o *FarmStorageSettingsCreateSpec) GetReplicaDiskDatastoreId() string`

GetReplicaDiskDatastoreId returns the ReplicaDiskDatastoreId field if non-nil, zero value otherwise.

### GetReplicaDiskDatastoreIdOk

`func (o *FarmStorageSettingsCreateSpec) GetReplicaDiskDatastoreIdOk() (*string, bool)`

GetReplicaDiskDatastoreIdOk returns a tuple with the ReplicaDiskDatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaDiskDatastoreId

`func (o *FarmStorageSettingsCreateSpec) SetReplicaDiskDatastoreId(v string)`

SetReplicaDiskDatastoreId sets ReplicaDiskDatastoreId field to given value.

### HasReplicaDiskDatastoreId

`func (o *FarmStorageSettingsCreateSpec) HasReplicaDiskDatastoreId() bool`

HasReplicaDiskDatastoreId returns a boolean if a field has been set.

### GetUseSeparateDatastoresReplicaAndOsDisks

`func (o *FarmStorageSettingsCreateSpec) GetUseSeparateDatastoresReplicaAndOsDisks() bool`

GetUseSeparateDatastoresReplicaAndOsDisks returns the UseSeparateDatastoresReplicaAndOsDisks field if non-nil, zero value otherwise.

### GetUseSeparateDatastoresReplicaAndOsDisksOk

`func (o *FarmStorageSettingsCreateSpec) GetUseSeparateDatastoresReplicaAndOsDisksOk() (*bool, bool)`

GetUseSeparateDatastoresReplicaAndOsDisksOk returns a tuple with the UseSeparateDatastoresReplicaAndOsDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSeparateDatastoresReplicaAndOsDisks

`func (o *FarmStorageSettingsCreateSpec) SetUseSeparateDatastoresReplicaAndOsDisks(v bool)`

SetUseSeparateDatastoresReplicaAndOsDisks sets UseSeparateDatastoresReplicaAndOsDisks field to given value.

### HasUseSeparateDatastoresReplicaAndOsDisks

`func (o *FarmStorageSettingsCreateSpec) HasUseSeparateDatastoresReplicaAndOsDisks() bool`

HasUseSeparateDatastoresReplicaAndOsDisks returns a boolean if a field has been set.

### GetUseViewStorageAccelerator

`func (o *FarmStorageSettingsCreateSpec) GetUseViewStorageAccelerator() bool`

GetUseViewStorageAccelerator returns the UseViewStorageAccelerator field if non-nil, zero value otherwise.

### GetUseViewStorageAcceleratorOk

`func (o *FarmStorageSettingsCreateSpec) GetUseViewStorageAcceleratorOk() (*bool, bool)`

GetUseViewStorageAcceleratorOk returns a tuple with the UseViewStorageAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseViewStorageAccelerator

`func (o *FarmStorageSettingsCreateSpec) SetUseViewStorageAccelerator(v bool)`

SetUseViewStorageAccelerator sets UseViewStorageAccelerator field to given value.

### HasUseViewStorageAccelerator

`func (o *FarmStorageSettingsCreateSpec) HasUseViewStorageAccelerator() bool`

HasUseViewStorageAccelerator returns a boolean if a field has been set.

### GetUseVsan

`func (o *FarmStorageSettingsCreateSpec) GetUseVsan() bool`

GetUseVsan returns the UseVsan field if non-nil, zero value otherwise.

### GetUseVsanOk

`func (o *FarmStorageSettingsCreateSpec) GetUseVsanOk() (*bool, bool)`

GetUseVsanOk returns a tuple with the UseVsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVsan

`func (o *FarmStorageSettingsCreateSpec) SetUseVsan(v bool)`

SetUseVsan sets UseVsan field to given value.

### HasUseVsan

`func (o *FarmStorageSettingsCreateSpec) HasUseVsan() bool`

HasUseVsan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


