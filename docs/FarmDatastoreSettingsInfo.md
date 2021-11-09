# FarmDatastoreSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreId** | Pointer to **string** | Id of the datastore. | [optional] 
**StorageOvercommit** | Pointer to **string** | Storage overcommit determines how view places new RDS Servers on the selected datastores. With an aggressive overcommit level, view reserves less space for sparse disk growth, but fits more RDS Servers on the datastore. * NONE: No overcommit. * CONSERVATIVE: Conservative. * MODERATE: Moderate. * AGGRESSIVE: Aggressive. * UNBOUNDED: Unbounded. | [optional] 

## Methods

### NewFarmDatastoreSettingsInfo

`func NewFarmDatastoreSettingsInfo() *FarmDatastoreSettingsInfo`

NewFarmDatastoreSettingsInfo instantiates a new FarmDatastoreSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmDatastoreSettingsInfoWithDefaults

`func NewFarmDatastoreSettingsInfoWithDefaults() *FarmDatastoreSettingsInfo`

NewFarmDatastoreSettingsInfoWithDefaults instantiates a new FarmDatastoreSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreId

`func (o *FarmDatastoreSettingsInfo) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *FarmDatastoreSettingsInfo) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *FarmDatastoreSettingsInfo) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *FarmDatastoreSettingsInfo) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetStorageOvercommit

`func (o *FarmDatastoreSettingsInfo) GetStorageOvercommit() string`

GetStorageOvercommit returns the StorageOvercommit field if non-nil, zero value otherwise.

### GetStorageOvercommitOk

`func (o *FarmDatastoreSettingsInfo) GetStorageOvercommitOk() (*string, bool)`

GetStorageOvercommitOk returns a tuple with the StorageOvercommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOvercommit

`func (o *FarmDatastoreSettingsInfo) SetStorageOvercommit(v string)`

SetStorageOvercommit sets StorageOvercommit field to given value.

### HasStorageOvercommit

`func (o *FarmDatastoreSettingsInfo) HasStorageOvercommit() bool`

HasStorageOvercommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


