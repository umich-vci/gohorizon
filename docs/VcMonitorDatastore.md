# VCMonitorDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityMb** | Pointer to **int32** | The capacity of the datastore in megabytes. | [optional] 
**Details** | Pointer to [**VCMonitorDatastoreDetails**](VCMonitorDatastoreDetails.md) |  | [optional] 
**FreeSpaceMb** | Pointer to **int32** | The free space on the datastore in megabytes. | [optional] 
**Status** | Pointer to **string** | Status of the datastore. * ACCESSIBLE: The datastore is accessible. * NOT_ACCESSIBLE: The datastore is not accessible. | [optional] 
**Type** | Pointer to **string** | Type of the datastore. * VMFS: VMFS datastore. * VSAN: VSAN datastore. * VVOL: VVOL datastore. | [optional] 

## Methods

### NewVCMonitorDatastore

`func NewVCMonitorDatastore() *VCMonitorDatastore`

NewVCMonitorDatastore instantiates a new VCMonitorDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCMonitorDatastoreWithDefaults

`func NewVCMonitorDatastoreWithDefaults() *VCMonitorDatastore`

NewVCMonitorDatastoreWithDefaults instantiates a new VCMonitorDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityMb

`func (o *VCMonitorDatastore) GetCapacityMb() int32`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *VCMonitorDatastore) GetCapacityMbOk() (*int32, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *VCMonitorDatastore) SetCapacityMb(v int32)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *VCMonitorDatastore) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetDetails

`func (o *VCMonitorDatastore) GetDetails() VCMonitorDatastoreDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VCMonitorDatastore) GetDetailsOk() (*VCMonitorDatastoreDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VCMonitorDatastore) SetDetails(v VCMonitorDatastoreDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VCMonitorDatastore) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFreeSpaceMb

`func (o *VCMonitorDatastore) GetFreeSpaceMb() int32`

GetFreeSpaceMb returns the FreeSpaceMb field if non-nil, zero value otherwise.

### GetFreeSpaceMbOk

`func (o *VCMonitorDatastore) GetFreeSpaceMbOk() (*int32, bool)`

GetFreeSpaceMbOk returns a tuple with the FreeSpaceMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceMb

`func (o *VCMonitorDatastore) SetFreeSpaceMb(v int32)`

SetFreeSpaceMb sets FreeSpaceMb field to given value.

### HasFreeSpaceMb

`func (o *VCMonitorDatastore) HasFreeSpaceMb() bool`

HasFreeSpaceMb returns a boolean if a field has been set.

### GetStatus

`func (o *VCMonitorDatastore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VCMonitorDatastore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VCMonitorDatastore) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VCMonitorDatastore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *VCMonitorDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VCMonitorDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VCMonitorDatastore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VCMonitorDatastore) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


