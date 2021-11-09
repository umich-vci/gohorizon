# DatastoreSpaceRequirementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskType** | Pointer to **string** | Indicates the type of disk used for storage. * OS: Disk to store operating system related data. * REPLICA: Disk for placement of replica VMs created by instant clone engine. | [optional] 
**MaxSizeDiskGb** | Pointer to **float64** | Indicates maximum recommended disk space, in GB. | [optional] 
**MidSizeDiskGb** | Pointer to **float64** | Indicates recommended disk space with 50% utilization, in GB. | [optional] 
**MinSizeDiskGb** | Pointer to **float64** | Indicates minimum recommended disk space, in GB. | [optional] 

## Methods

### NewDatastoreSpaceRequirementInfo

`func NewDatastoreSpaceRequirementInfo() *DatastoreSpaceRequirementInfo`

NewDatastoreSpaceRequirementInfo instantiates a new DatastoreSpaceRequirementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreSpaceRequirementInfoWithDefaults

`func NewDatastoreSpaceRequirementInfoWithDefaults() *DatastoreSpaceRequirementInfo`

NewDatastoreSpaceRequirementInfoWithDefaults instantiates a new DatastoreSpaceRequirementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskType

`func (o *DatastoreSpaceRequirementInfo) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *DatastoreSpaceRequirementInfo) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *DatastoreSpaceRequirementInfo) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *DatastoreSpaceRequirementInfo) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetMaxSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) GetMaxSizeDiskGb() float64`

GetMaxSizeDiskGb returns the MaxSizeDiskGb field if non-nil, zero value otherwise.

### GetMaxSizeDiskGbOk

`func (o *DatastoreSpaceRequirementInfo) GetMaxSizeDiskGbOk() (*float64, bool)`

GetMaxSizeDiskGbOk returns a tuple with the MaxSizeDiskGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) SetMaxSizeDiskGb(v float64)`

SetMaxSizeDiskGb sets MaxSizeDiskGb field to given value.

### HasMaxSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) HasMaxSizeDiskGb() bool`

HasMaxSizeDiskGb returns a boolean if a field has been set.

### GetMidSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) GetMidSizeDiskGb() float64`

GetMidSizeDiskGb returns the MidSizeDiskGb field if non-nil, zero value otherwise.

### GetMidSizeDiskGbOk

`func (o *DatastoreSpaceRequirementInfo) GetMidSizeDiskGbOk() (*float64, bool)`

GetMidSizeDiskGbOk returns a tuple with the MidSizeDiskGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) SetMidSizeDiskGb(v float64)`

SetMidSizeDiskGb sets MidSizeDiskGb field to given value.

### HasMidSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) HasMidSizeDiskGb() bool`

HasMidSizeDiskGb returns a boolean if a field has been set.

### GetMinSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) GetMinSizeDiskGb() float64`

GetMinSizeDiskGb returns the MinSizeDiskGb field if non-nil, zero value otherwise.

### GetMinSizeDiskGbOk

`func (o *DatastoreSpaceRequirementInfo) GetMinSizeDiskGbOk() (*float64, bool)`

GetMinSizeDiskGbOk returns a tuple with the MinSizeDiskGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) SetMinSizeDiskGb(v float64)`

SetMinSizeDiskGb sets MinSizeDiskGb field to given value.

### HasMinSizeDiskGb

`func (o *DatastoreSpaceRequirementInfo) HasMinSizeDiskGb() bool`

HasMinSizeDiskGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


