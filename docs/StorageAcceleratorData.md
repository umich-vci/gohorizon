# StorageAcceleratorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultCacheSizeMb** | Pointer to **int32** | Default size of the cache in megabytes. This property has a default value of 1024. This property has a minimum value of 100. This property has a maximum value of 2048.  | [optional] 
**Enabled** | Pointer to **bool** | Is View Storage Accelerator enabled? This property has a default value of false. | [optional] 
**HostOverrides** | Pointer to [**[]HostOverrideData**](HostOverrideData.md) | Cache size overrides for hosts which support View Storage Accelerator. | [optional] 

## Methods

### NewStorageAcceleratorData

`func NewStorageAcceleratorData() *StorageAcceleratorData`

NewStorageAcceleratorData instantiates a new StorageAcceleratorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageAcceleratorDataWithDefaults

`func NewStorageAcceleratorDataWithDefaults() *StorageAcceleratorData`

NewStorageAcceleratorDataWithDefaults instantiates a new StorageAcceleratorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultCacheSizeMb

`func (o *StorageAcceleratorData) GetDefaultCacheSizeMb() int32`

GetDefaultCacheSizeMb returns the DefaultCacheSizeMb field if non-nil, zero value otherwise.

### GetDefaultCacheSizeMbOk

`func (o *StorageAcceleratorData) GetDefaultCacheSizeMbOk() (*int32, bool)`

GetDefaultCacheSizeMbOk returns a tuple with the DefaultCacheSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCacheSizeMb

`func (o *StorageAcceleratorData) SetDefaultCacheSizeMb(v int32)`

SetDefaultCacheSizeMb sets DefaultCacheSizeMb field to given value.

### HasDefaultCacheSizeMb

`func (o *StorageAcceleratorData) HasDefaultCacheSizeMb() bool`

HasDefaultCacheSizeMb returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageAcceleratorData) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageAcceleratorData) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageAcceleratorData) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageAcceleratorData) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHostOverrides

`func (o *StorageAcceleratorData) GetHostOverrides() []HostOverrideData`

GetHostOverrides returns the HostOverrides field if non-nil, zero value otherwise.

### GetHostOverridesOk

`func (o *StorageAcceleratorData) GetHostOverridesOk() (*[]HostOverrideData, bool)`

GetHostOverridesOk returns a tuple with the HostOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOverrides

`func (o *StorageAcceleratorData) SetHostOverrides(v []HostOverrideData)`

SetHostOverrides sets HostOverrides field to given value.

### HasHostOverrides

`func (o *StorageAcceleratorData) HasHostOverrides() bool`

HasHostOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


