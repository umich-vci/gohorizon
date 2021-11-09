# DesktopPoolStorageSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | Pointer to [**[]DesktopPoolDatastoreSettings**](DesktopPoolDatastoreSettings.md) | Datastores to store the machine (or the OS disk using other options for linked clone or instant clone machine storage) | [optional] 
**NonPersistentDiskDriveLetter** | Pointer to **string** | Non persistent disk drive letter. | [optional] 
**NonPersistentDiskSizeMb** | Pointer to **int32** | Size of the non persistent disk in MB. | [optional] 
**NonPersistentRedirectDisposableFiles** | Pointer to **bool** | Redirect disposable files to a non-persistent disk that will be deleted automatically when a user&#39;s session ends. | [optional] 
**PersistentDiskDatastores** | Pointer to [**[]DesktopPoolDatastoreSettings**](DesktopPoolDatastoreSettings.md) | Datastores to store persistent disks for linked clone machines. | [optional] 
**PersistentDiskDriveLetter** | Pointer to **string** | Persistent disk drive letter. | [optional] 
**PersistentDiskSizeMb** | Pointer to **int32** | Size of the persistent disk in MB. | [optional] 
**ReclaimVmDiskSpace** | Pointer to **bool** | With vSphere 5.x, virtual machines can be configured to use a space efficient disk format that supports reclamation of unused disk space (such as deleted files). This option reclaims unused disk space on each virtual machine. The operation is initiated when an estimate of used disk space exceeds the specified threshold. | [optional] 
**ReclamationThresholdMb** | Pointer to **int64** | Initiate reclamation when unused space on virtual machine exceeds the threshold in MB. | [optional] 
**RedirectWindowsProfile** | Pointer to **bool** | Windows profiles will be redirected to persistent disks, which are not affected by View Composer operations such as refresh, recompose and rebalance. | [optional] 
**ReplicaDiskDatastoreId** | Pointer to **string** | Datastore to store replica disks for linked clone and instant clone machines. | [optional] 
**UseNativeSnapshots** | Pointer to **bool** | Applicable To: Linked/instant clone automated desktop pool.&lt;br&gt;Native NFS Snapshots (VAAI - vStorage API for Array Integration) is a hardware feature of certain storage arrays. It uses native snapshotting technology to provide linked clone functionality. | [optional] 
**UseSeparateDatastoresPersistentAndOsDisks** | Pointer to **bool** | Whether to use separate datastores for persistent and OS disks. | [optional] 
**UseSeparateDatastoresReplicaAndOsDisks** | Pointer to **bool** | Whether to use separate datastores for replica and OS disks. | [optional] 
**UseVsan** | Pointer to **bool** | Whether to use vSphere vSAN. | [optional] 

## Methods

### NewDesktopPoolStorageSettings

`func NewDesktopPoolStorageSettings() *DesktopPoolStorageSettings`

NewDesktopPoolStorageSettings instantiates a new DesktopPoolStorageSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolStorageSettingsWithDefaults

`func NewDesktopPoolStorageSettingsWithDefaults() *DesktopPoolStorageSettings`

NewDesktopPoolStorageSettingsWithDefaults instantiates a new DesktopPoolStorageSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *DesktopPoolStorageSettings) GetDatastores() []DesktopPoolDatastoreSettings`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *DesktopPoolStorageSettings) GetDatastoresOk() (*[]DesktopPoolDatastoreSettings, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *DesktopPoolStorageSettings) SetDatastores(v []DesktopPoolDatastoreSettings)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *DesktopPoolStorageSettings) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetNonPersistentDiskDriveLetter

`func (o *DesktopPoolStorageSettings) GetNonPersistentDiskDriveLetter() string`

GetNonPersistentDiskDriveLetter returns the NonPersistentDiskDriveLetter field if non-nil, zero value otherwise.

### GetNonPersistentDiskDriveLetterOk

`func (o *DesktopPoolStorageSettings) GetNonPersistentDiskDriveLetterOk() (*string, bool)`

GetNonPersistentDiskDriveLetterOk returns a tuple with the NonPersistentDiskDriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPersistentDiskDriveLetter

`func (o *DesktopPoolStorageSettings) SetNonPersistentDiskDriveLetter(v string)`

SetNonPersistentDiskDriveLetter sets NonPersistentDiskDriveLetter field to given value.

### HasNonPersistentDiskDriveLetter

`func (o *DesktopPoolStorageSettings) HasNonPersistentDiskDriveLetter() bool`

HasNonPersistentDiskDriveLetter returns a boolean if a field has been set.

### GetNonPersistentDiskSizeMb

`func (o *DesktopPoolStorageSettings) GetNonPersistentDiskSizeMb() int32`

GetNonPersistentDiskSizeMb returns the NonPersistentDiskSizeMb field if non-nil, zero value otherwise.

### GetNonPersistentDiskSizeMbOk

`func (o *DesktopPoolStorageSettings) GetNonPersistentDiskSizeMbOk() (*int32, bool)`

GetNonPersistentDiskSizeMbOk returns a tuple with the NonPersistentDiskSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPersistentDiskSizeMb

`func (o *DesktopPoolStorageSettings) SetNonPersistentDiskSizeMb(v int32)`

SetNonPersistentDiskSizeMb sets NonPersistentDiskSizeMb field to given value.

### HasNonPersistentDiskSizeMb

`func (o *DesktopPoolStorageSettings) HasNonPersistentDiskSizeMb() bool`

HasNonPersistentDiskSizeMb returns a boolean if a field has been set.

### GetNonPersistentRedirectDisposableFiles

`func (o *DesktopPoolStorageSettings) GetNonPersistentRedirectDisposableFiles() bool`

GetNonPersistentRedirectDisposableFiles returns the NonPersistentRedirectDisposableFiles field if non-nil, zero value otherwise.

### GetNonPersistentRedirectDisposableFilesOk

`func (o *DesktopPoolStorageSettings) GetNonPersistentRedirectDisposableFilesOk() (*bool, bool)`

GetNonPersistentRedirectDisposableFilesOk returns a tuple with the NonPersistentRedirectDisposableFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPersistentRedirectDisposableFiles

`func (o *DesktopPoolStorageSettings) SetNonPersistentRedirectDisposableFiles(v bool)`

SetNonPersistentRedirectDisposableFiles sets NonPersistentRedirectDisposableFiles field to given value.

### HasNonPersistentRedirectDisposableFiles

`func (o *DesktopPoolStorageSettings) HasNonPersistentRedirectDisposableFiles() bool`

HasNonPersistentRedirectDisposableFiles returns a boolean if a field has been set.

### GetPersistentDiskDatastores

`func (o *DesktopPoolStorageSettings) GetPersistentDiskDatastores() []DesktopPoolDatastoreSettings`

GetPersistentDiskDatastores returns the PersistentDiskDatastores field if non-nil, zero value otherwise.

### GetPersistentDiskDatastoresOk

`func (o *DesktopPoolStorageSettings) GetPersistentDiskDatastoresOk() (*[]DesktopPoolDatastoreSettings, bool)`

GetPersistentDiskDatastoresOk returns a tuple with the PersistentDiskDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentDiskDatastores

`func (o *DesktopPoolStorageSettings) SetPersistentDiskDatastores(v []DesktopPoolDatastoreSettings)`

SetPersistentDiskDatastores sets PersistentDiskDatastores field to given value.

### HasPersistentDiskDatastores

`func (o *DesktopPoolStorageSettings) HasPersistentDiskDatastores() bool`

HasPersistentDiskDatastores returns a boolean if a field has been set.

### GetPersistentDiskDriveLetter

`func (o *DesktopPoolStorageSettings) GetPersistentDiskDriveLetter() string`

GetPersistentDiskDriveLetter returns the PersistentDiskDriveLetter field if non-nil, zero value otherwise.

### GetPersistentDiskDriveLetterOk

`func (o *DesktopPoolStorageSettings) GetPersistentDiskDriveLetterOk() (*string, bool)`

GetPersistentDiskDriveLetterOk returns a tuple with the PersistentDiskDriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentDiskDriveLetter

`func (o *DesktopPoolStorageSettings) SetPersistentDiskDriveLetter(v string)`

SetPersistentDiskDriveLetter sets PersistentDiskDriveLetter field to given value.

### HasPersistentDiskDriveLetter

`func (o *DesktopPoolStorageSettings) HasPersistentDiskDriveLetter() bool`

HasPersistentDiskDriveLetter returns a boolean if a field has been set.

### GetPersistentDiskSizeMb

`func (o *DesktopPoolStorageSettings) GetPersistentDiskSizeMb() int32`

GetPersistentDiskSizeMb returns the PersistentDiskSizeMb field if non-nil, zero value otherwise.

### GetPersistentDiskSizeMbOk

`func (o *DesktopPoolStorageSettings) GetPersistentDiskSizeMbOk() (*int32, bool)`

GetPersistentDiskSizeMbOk returns a tuple with the PersistentDiskSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentDiskSizeMb

`func (o *DesktopPoolStorageSettings) SetPersistentDiskSizeMb(v int32)`

SetPersistentDiskSizeMb sets PersistentDiskSizeMb field to given value.

### HasPersistentDiskSizeMb

`func (o *DesktopPoolStorageSettings) HasPersistentDiskSizeMb() bool`

HasPersistentDiskSizeMb returns a boolean if a field has been set.

### GetReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettings) GetReclaimVmDiskSpace() bool`

GetReclaimVmDiskSpace returns the ReclaimVmDiskSpace field if non-nil, zero value otherwise.

### GetReclaimVmDiskSpaceOk

`func (o *DesktopPoolStorageSettings) GetReclaimVmDiskSpaceOk() (*bool, bool)`

GetReclaimVmDiskSpaceOk returns a tuple with the ReclaimVmDiskSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettings) SetReclaimVmDiskSpace(v bool)`

SetReclaimVmDiskSpace sets ReclaimVmDiskSpace field to given value.

### HasReclaimVmDiskSpace

`func (o *DesktopPoolStorageSettings) HasReclaimVmDiskSpace() bool`

HasReclaimVmDiskSpace returns a boolean if a field has been set.

### GetReclamationThresholdMb

`func (o *DesktopPoolStorageSettings) GetReclamationThresholdMb() int64`

GetReclamationThresholdMb returns the ReclamationThresholdMb field if non-nil, zero value otherwise.

### GetReclamationThresholdMbOk

`func (o *DesktopPoolStorageSettings) GetReclamationThresholdMbOk() (*int64, bool)`

GetReclamationThresholdMbOk returns a tuple with the ReclamationThresholdMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclamationThresholdMb

`func (o *DesktopPoolStorageSettings) SetReclamationThresholdMb(v int64)`

SetReclamationThresholdMb sets ReclamationThresholdMb field to given value.

### HasReclamationThresholdMb

`func (o *DesktopPoolStorageSettings) HasReclamationThresholdMb() bool`

HasReclamationThresholdMb returns a boolean if a field has been set.

### GetRedirectWindowsProfile

`func (o *DesktopPoolStorageSettings) GetRedirectWindowsProfile() bool`

GetRedirectWindowsProfile returns the RedirectWindowsProfile field if non-nil, zero value otherwise.

### GetRedirectWindowsProfileOk

`func (o *DesktopPoolStorageSettings) GetRedirectWindowsProfileOk() (*bool, bool)`

GetRedirectWindowsProfileOk returns a tuple with the RedirectWindowsProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectWindowsProfile

`func (o *DesktopPoolStorageSettings) SetRedirectWindowsProfile(v bool)`

SetRedirectWindowsProfile sets RedirectWindowsProfile field to given value.

### HasRedirectWindowsProfile

`func (o *DesktopPoolStorageSettings) HasRedirectWindowsProfile() bool`

HasRedirectWindowsProfile returns a boolean if a field has been set.

### GetReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettings) GetReplicaDiskDatastoreId() string`

GetReplicaDiskDatastoreId returns the ReplicaDiskDatastoreId field if non-nil, zero value otherwise.

### GetReplicaDiskDatastoreIdOk

`func (o *DesktopPoolStorageSettings) GetReplicaDiskDatastoreIdOk() (*string, bool)`

GetReplicaDiskDatastoreIdOk returns a tuple with the ReplicaDiskDatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettings) SetReplicaDiskDatastoreId(v string)`

SetReplicaDiskDatastoreId sets ReplicaDiskDatastoreId field to given value.

### HasReplicaDiskDatastoreId

`func (o *DesktopPoolStorageSettings) HasReplicaDiskDatastoreId() bool`

HasReplicaDiskDatastoreId returns a boolean if a field has been set.

### GetUseNativeSnapshots

`func (o *DesktopPoolStorageSettings) GetUseNativeSnapshots() bool`

GetUseNativeSnapshots returns the UseNativeSnapshots field if non-nil, zero value otherwise.

### GetUseNativeSnapshotsOk

`func (o *DesktopPoolStorageSettings) GetUseNativeSnapshotsOk() (*bool, bool)`

GetUseNativeSnapshotsOk returns a tuple with the UseNativeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNativeSnapshots

`func (o *DesktopPoolStorageSettings) SetUseNativeSnapshots(v bool)`

SetUseNativeSnapshots sets UseNativeSnapshots field to given value.

### HasUseNativeSnapshots

`func (o *DesktopPoolStorageSettings) HasUseNativeSnapshots() bool`

HasUseNativeSnapshots returns a boolean if a field has been set.

### GetUseSeparateDatastoresPersistentAndOsDisks

`func (o *DesktopPoolStorageSettings) GetUseSeparateDatastoresPersistentAndOsDisks() bool`

GetUseSeparateDatastoresPersistentAndOsDisks returns the UseSeparateDatastoresPersistentAndOsDisks field if non-nil, zero value otherwise.

### GetUseSeparateDatastoresPersistentAndOsDisksOk

`func (o *DesktopPoolStorageSettings) GetUseSeparateDatastoresPersistentAndOsDisksOk() (*bool, bool)`

GetUseSeparateDatastoresPersistentAndOsDisksOk returns a tuple with the UseSeparateDatastoresPersistentAndOsDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSeparateDatastoresPersistentAndOsDisks

`func (o *DesktopPoolStorageSettings) SetUseSeparateDatastoresPersistentAndOsDisks(v bool)`

SetUseSeparateDatastoresPersistentAndOsDisks sets UseSeparateDatastoresPersistentAndOsDisks field to given value.

### HasUseSeparateDatastoresPersistentAndOsDisks

`func (o *DesktopPoolStorageSettings) HasUseSeparateDatastoresPersistentAndOsDisks() bool`

HasUseSeparateDatastoresPersistentAndOsDisks returns a boolean if a field has been set.

### GetUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettings) GetUseSeparateDatastoresReplicaAndOsDisks() bool`

GetUseSeparateDatastoresReplicaAndOsDisks returns the UseSeparateDatastoresReplicaAndOsDisks field if non-nil, zero value otherwise.

### GetUseSeparateDatastoresReplicaAndOsDisksOk

`func (o *DesktopPoolStorageSettings) GetUseSeparateDatastoresReplicaAndOsDisksOk() (*bool, bool)`

GetUseSeparateDatastoresReplicaAndOsDisksOk returns a tuple with the UseSeparateDatastoresReplicaAndOsDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettings) SetUseSeparateDatastoresReplicaAndOsDisks(v bool)`

SetUseSeparateDatastoresReplicaAndOsDisks sets UseSeparateDatastoresReplicaAndOsDisks field to given value.

### HasUseSeparateDatastoresReplicaAndOsDisks

`func (o *DesktopPoolStorageSettings) HasUseSeparateDatastoresReplicaAndOsDisks() bool`

HasUseSeparateDatastoresReplicaAndOsDisks returns a boolean if a field has been set.

### GetUseVsan

`func (o *DesktopPoolStorageSettings) GetUseVsan() bool`

GetUseVsan returns the UseVsan field if non-nil, zero value otherwise.

### GetUseVsanOk

`func (o *DesktopPoolStorageSettings) GetUseVsanOk() (*bool, bool)`

GetUseVsanOk returns a tuple with the UseVsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVsan

`func (o *DesktopPoolStorageSettings) SetUseVsan(v bool)`

SetUseVsan sets UseVsan field to given value.

### HasUseVsan

`func (o *DesktopPoolStorageSettings) HasUseVsan() bool`

HasUseVsan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


