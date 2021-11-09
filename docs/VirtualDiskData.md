# VirtualDiskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityMb** | Pointer to **int64** | The disk capacity in MB. | [optional] 
**DatastorePath** | Pointer to **string** | The virtual disk&#39;s datastore. | [optional] 
**Path** | Pointer to **string** | The virtual disk path. | [optional] 

## Methods

### NewVirtualDiskData

`func NewVirtualDiskData() *VirtualDiskData`

NewVirtualDiskData instantiates a new VirtualDiskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskDataWithDefaults

`func NewVirtualDiskDataWithDefaults() *VirtualDiskData`

NewVirtualDiskDataWithDefaults instantiates a new VirtualDiskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityMb

`func (o *VirtualDiskData) GetCapacityMb() int64`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *VirtualDiskData) GetCapacityMbOk() (*int64, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *VirtualDiskData) SetCapacityMb(v int64)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *VirtualDiskData) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetDatastorePath

`func (o *VirtualDiskData) GetDatastorePath() string`

GetDatastorePath returns the DatastorePath field if non-nil, zero value otherwise.

### GetDatastorePathOk

`func (o *VirtualDiskData) GetDatastorePathOk() (*string, bool)`

GetDatastorePathOk returns a tuple with the DatastorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastorePath

`func (o *VirtualDiskData) SetDatastorePath(v string)`

SetDatastorePath sets DatastorePath field to given value.

### HasDatastorePath

`func (o *VirtualDiskData) HasDatastorePath() bool`

HasDatastorePath returns a boolean if a field has been set.

### GetPath

`func (o *VirtualDiskData) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VirtualDiskData) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VirtualDiskData) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VirtualDiskData) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


