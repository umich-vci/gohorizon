# HostOverrideData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheSizeMb** | Pointer to **int32** | Size of the cache in megabytes. This property has a minimum value of 100. This property has a maximum value of 2048. | [optional] 
**Path** | Pointer to **string** | The path of the host that supports View Storage Accelerator. | [optional] 

## Methods

### NewHostOverrideData

`func NewHostOverrideData() *HostOverrideData`

NewHostOverrideData instantiates a new HostOverrideData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOverrideDataWithDefaults

`func NewHostOverrideDataWithDefaults() *HostOverrideData`

NewHostOverrideDataWithDefaults instantiates a new HostOverrideData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheSizeMb

`func (o *HostOverrideData) GetCacheSizeMb() int32`

GetCacheSizeMb returns the CacheSizeMb field if non-nil, zero value otherwise.

### GetCacheSizeMbOk

`func (o *HostOverrideData) GetCacheSizeMbOk() (*int32, bool)`

GetCacheSizeMbOk returns a tuple with the CacheSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSizeMb

`func (o *HostOverrideData) SetCacheSizeMb(v int32)`

SetCacheSizeMb sets CacheSizeMb field to given value.

### HasCacheSizeMb

`func (o *HostOverrideData) HasCacheSizeMb() bool`

HasCacheSizeMb returns a boolean if a field has been set.

### GetPath

`func (o *HostOverrideData) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HostOverrideData) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HostOverrideData) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HostOverrideData) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


