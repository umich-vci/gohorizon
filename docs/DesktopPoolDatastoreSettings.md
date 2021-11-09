# DesktopPoolDatastoreSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreId** | Pointer to **string** | Id of the datastore. | [optional] 
**SdrsCluster** | Pointer to **bool** | Whether datastore represents a Storage DRS cluster. | [optional] 
**StorageOvercommit** | Pointer to **string** | Storage overcommit determines how Horizon places new machines on the selected datastores. With an aggressive overcommit level, Horizon reserves less space for sparse disk growth, but fits more machines on the datastore. * NONE: No overcommit. * CONSERVATIVE: Conservative. * MODERATE: Moderate. * AGGRESSIVE: Aggressive. * UNBOUNDED: Unbounded. | [optional] 

## Methods

### NewDesktopPoolDatastoreSettings

`func NewDesktopPoolDatastoreSettings() *DesktopPoolDatastoreSettings`

NewDesktopPoolDatastoreSettings instantiates a new DesktopPoolDatastoreSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolDatastoreSettingsWithDefaults

`func NewDesktopPoolDatastoreSettingsWithDefaults() *DesktopPoolDatastoreSettings`

NewDesktopPoolDatastoreSettingsWithDefaults instantiates a new DesktopPoolDatastoreSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreId

`func (o *DesktopPoolDatastoreSettings) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *DesktopPoolDatastoreSettings) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *DesktopPoolDatastoreSettings) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *DesktopPoolDatastoreSettings) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetSdrsCluster

`func (o *DesktopPoolDatastoreSettings) GetSdrsCluster() bool`

GetSdrsCluster returns the SdrsCluster field if non-nil, zero value otherwise.

### GetSdrsClusterOk

`func (o *DesktopPoolDatastoreSettings) GetSdrsClusterOk() (*bool, bool)`

GetSdrsClusterOk returns a tuple with the SdrsCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdrsCluster

`func (o *DesktopPoolDatastoreSettings) SetSdrsCluster(v bool)`

SetSdrsCluster sets SdrsCluster field to given value.

### HasSdrsCluster

`func (o *DesktopPoolDatastoreSettings) HasSdrsCluster() bool`

HasSdrsCluster returns a boolean if a field has been set.

### GetStorageOvercommit

`func (o *DesktopPoolDatastoreSettings) GetStorageOvercommit() string`

GetStorageOvercommit returns the StorageOvercommit field if non-nil, zero value otherwise.

### GetStorageOvercommitOk

`func (o *DesktopPoolDatastoreSettings) GetStorageOvercommitOk() (*string, bool)`

GetStorageOvercommitOk returns a tuple with the StorageOvercommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOvercommit

`func (o *DesktopPoolDatastoreSettings) SetStorageOvercommit(v string)`

SetStorageOvercommit sets StorageOvercommit field to given value.

### HasStorageOvercommit

`func (o *DesktopPoolDatastoreSettings) HasStorageOvercommit() bool`

HasStorageOvercommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


