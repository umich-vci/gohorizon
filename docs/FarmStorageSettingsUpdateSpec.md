# FarmStorageSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | [**[]FarmDatastoreSettingsUpdateSpec**](FarmDatastoreSettingsUpdateSpec.md) | List of IDs of the datastore used to store the RDS Server. This can be modified only if there are no current operations ( operation is NONE). | 
**ReplicaDiskDatastoreId** | Pointer to **string** | Datastore to store replica disks for instant clone machines. This is required if use_separate_datastores_replica_and_os_disks is set to true. This can be modified only if there are no current operations ( operation is NONE). | [optional] 

## Methods

### NewFarmStorageSettingsUpdateSpec

`func NewFarmStorageSettingsUpdateSpec(datastores []FarmDatastoreSettingsUpdateSpec, ) *FarmStorageSettingsUpdateSpec`

NewFarmStorageSettingsUpdateSpec instantiates a new FarmStorageSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmStorageSettingsUpdateSpecWithDefaults

`func NewFarmStorageSettingsUpdateSpecWithDefaults() *FarmStorageSettingsUpdateSpec`

NewFarmStorageSettingsUpdateSpecWithDefaults instantiates a new FarmStorageSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *FarmStorageSettingsUpdateSpec) GetDatastores() []FarmDatastoreSettingsUpdateSpec`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *FarmStorageSettingsUpdateSpec) GetDatastoresOk() (*[]FarmDatastoreSettingsUpdateSpec, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *FarmStorageSettingsUpdateSpec) SetDatastores(v []FarmDatastoreSettingsUpdateSpec)`

SetDatastores sets Datastores field to given value.


### GetReplicaDiskDatastoreId

`func (o *FarmStorageSettingsUpdateSpec) GetReplicaDiskDatastoreId() string`

GetReplicaDiskDatastoreId returns the ReplicaDiskDatastoreId field if non-nil, zero value otherwise.

### GetReplicaDiskDatastoreIdOk

`func (o *FarmStorageSettingsUpdateSpec) GetReplicaDiskDatastoreIdOk() (*string, bool)`

GetReplicaDiskDatastoreIdOk returns a tuple with the ReplicaDiskDatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaDiskDatastoreId

`func (o *FarmStorageSettingsUpdateSpec) SetReplicaDiskDatastoreId(v string)`

SetReplicaDiskDatastoreId sets ReplicaDiskDatastoreId field to given value.

### HasReplicaDiskDatastoreId

`func (o *FarmStorageSettingsUpdateSpec) HasReplicaDiskDatastoreId() bool`

HasReplicaDiskDatastoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


