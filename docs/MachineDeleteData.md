# MachineDeleteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDeleteFromMultiDesktopPools** | Pointer to **bool** | Determines whether the machines from different desktop pools can be deleted. This defaults to false in which case only machines belonging to single desktop pool can be deleted. If true, machines from different desktop pools can be deleted. | [optional] 
**ArchiveDatastoreId** | Pointer to **string** | Determines the datastore where the persistent user disk will be saved for future use. Both this as well as the archiveDatastorePathId need to be set. If this is unset and archivePersistentDisk is specified, the persistent disk is archived in place. | [optional] 
**ArchiveDatastorePathId** | Pointer to **string** | Determines the location in the datastore where the persistent user disk will be saved for future use. If this is set, then archiveDatastoreId also needs to be specified.If this is unset and archivePersistentDisk is specified, the persistent disk is archived in place. | [optional] 
**ArchivePersistentDisk** | Pointer to **bool** | Determines whether to detach the persistent user disk and save it for future use. This can only be specified for linked-clone desktops with redirectWindowsProfile enabled, in which case it defaults to true.  | [optional] 
**DeleteFromDisk** | Pointer to **bool** | Determines whether the Machine VM should be deleted from vCenter Server. This is only applicable for managed machines. This must always be true for machines in linked and instant clone desktops. This defaults to true for linked and instant clone machines and false for all other types. If this is true, then machine being deleted must not have any active user session, otherwise delete operation would fail. | [optional] 
**ForceLogoffSession** | Pointer to **bool** | Determines whether active session on the machine to be logged off before deletion. This is only applicable for managed machines. If true, active session on the machine will be logged off before Machine delete. Otherwise,it will result in an exception.  | [optional] 

## Methods

### NewMachineDeleteData

`func NewMachineDeleteData() *MachineDeleteData`

NewMachineDeleteData instantiates a new MachineDeleteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineDeleteDataWithDefaults

`func NewMachineDeleteDataWithDefaults() *MachineDeleteData`

NewMachineDeleteDataWithDefaults instantiates a new MachineDeleteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDeleteFromMultiDesktopPools

`func (o *MachineDeleteData) GetAllowDeleteFromMultiDesktopPools() bool`

GetAllowDeleteFromMultiDesktopPools returns the AllowDeleteFromMultiDesktopPools field if non-nil, zero value otherwise.

### GetAllowDeleteFromMultiDesktopPoolsOk

`func (o *MachineDeleteData) GetAllowDeleteFromMultiDesktopPoolsOk() (*bool, bool)`

GetAllowDeleteFromMultiDesktopPoolsOk returns a tuple with the AllowDeleteFromMultiDesktopPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeleteFromMultiDesktopPools

`func (o *MachineDeleteData) SetAllowDeleteFromMultiDesktopPools(v bool)`

SetAllowDeleteFromMultiDesktopPools sets AllowDeleteFromMultiDesktopPools field to given value.

### HasAllowDeleteFromMultiDesktopPools

`func (o *MachineDeleteData) HasAllowDeleteFromMultiDesktopPools() bool`

HasAllowDeleteFromMultiDesktopPools returns a boolean if a field has been set.

### GetArchiveDatastoreId

`func (o *MachineDeleteData) GetArchiveDatastoreId() string`

GetArchiveDatastoreId returns the ArchiveDatastoreId field if non-nil, zero value otherwise.

### GetArchiveDatastoreIdOk

`func (o *MachineDeleteData) GetArchiveDatastoreIdOk() (*string, bool)`

GetArchiveDatastoreIdOk returns a tuple with the ArchiveDatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveDatastoreId

`func (o *MachineDeleteData) SetArchiveDatastoreId(v string)`

SetArchiveDatastoreId sets ArchiveDatastoreId field to given value.

### HasArchiveDatastoreId

`func (o *MachineDeleteData) HasArchiveDatastoreId() bool`

HasArchiveDatastoreId returns a boolean if a field has been set.

### GetArchiveDatastorePathId

`func (o *MachineDeleteData) GetArchiveDatastorePathId() string`

GetArchiveDatastorePathId returns the ArchiveDatastorePathId field if non-nil, zero value otherwise.

### GetArchiveDatastorePathIdOk

`func (o *MachineDeleteData) GetArchiveDatastorePathIdOk() (*string, bool)`

GetArchiveDatastorePathIdOk returns a tuple with the ArchiveDatastorePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveDatastorePathId

`func (o *MachineDeleteData) SetArchiveDatastorePathId(v string)`

SetArchiveDatastorePathId sets ArchiveDatastorePathId field to given value.

### HasArchiveDatastorePathId

`func (o *MachineDeleteData) HasArchiveDatastorePathId() bool`

HasArchiveDatastorePathId returns a boolean if a field has been set.

### GetArchivePersistentDisk

`func (o *MachineDeleteData) GetArchivePersistentDisk() bool`

GetArchivePersistentDisk returns the ArchivePersistentDisk field if non-nil, zero value otherwise.

### GetArchivePersistentDiskOk

`func (o *MachineDeleteData) GetArchivePersistentDiskOk() (*bool, bool)`

GetArchivePersistentDiskOk returns a tuple with the ArchivePersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivePersistentDisk

`func (o *MachineDeleteData) SetArchivePersistentDisk(v bool)`

SetArchivePersistentDisk sets ArchivePersistentDisk field to given value.

### HasArchivePersistentDisk

`func (o *MachineDeleteData) HasArchivePersistentDisk() bool`

HasArchivePersistentDisk returns a boolean if a field has been set.

### GetDeleteFromDisk

`func (o *MachineDeleteData) GetDeleteFromDisk() bool`

GetDeleteFromDisk returns the DeleteFromDisk field if non-nil, zero value otherwise.

### GetDeleteFromDiskOk

`func (o *MachineDeleteData) GetDeleteFromDiskOk() (*bool, bool)`

GetDeleteFromDiskOk returns a tuple with the DeleteFromDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFromDisk

`func (o *MachineDeleteData) SetDeleteFromDisk(v bool)`

SetDeleteFromDisk sets DeleteFromDisk field to given value.

### HasDeleteFromDisk

`func (o *MachineDeleteData) HasDeleteFromDisk() bool`

HasDeleteFromDisk returns a boolean if a field has been set.

### GetForceLogoffSession

`func (o *MachineDeleteData) GetForceLogoffSession() bool`

GetForceLogoffSession returns the ForceLogoffSession field if non-nil, zero value otherwise.

### GetForceLogoffSessionOk

`func (o *MachineDeleteData) GetForceLogoffSessionOk() (*bool, bool)`

GetForceLogoffSessionOk returns a tuple with the ForceLogoffSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceLogoffSession

`func (o *MachineDeleteData) SetForceLogoffSession(v bool)`

SetForceLogoffSession sets ForceLogoffSession field to given value.

### HasForceLogoffSession

`func (o *MachineDeleteData) HasForceLogoffSession() bool`

HasForceLogoffSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


