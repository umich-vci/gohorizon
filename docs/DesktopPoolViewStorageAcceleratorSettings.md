# DesktopPoolViewStorageAcceleratorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutTimes** | Pointer to [**[]ViewStorageAcceleratorBlackoutTimeSettings**](ViewStorageAcceleratorBlackoutTimeSettings.md) | A list of blackout times. Storage accelerator regeneration and machine disk space reclamation do not occur during blackout times. The same blackout policy applies to both operations. If unset, no blackout times are used. | [optional] 
**RegenerateViewStorageAcceleratorDays** | Pointer to **int32** | How often to regenerate the View Storage Accelerator cache. Measured in Days. | [optional] 
**UseViewStorageAccelerator** | Pointer to **bool** | Whether to use View Storage Accelerator. | [optional] 
**ViewStorageAcceleratorDiskTypes** | Pointer to **string** | Disk types to enable for the View Storage Accelerator feature. This is only applicable to linked clone desktop pools. * OS_DISKS: OS disks. * OS_AND_PERSISTENT_DISKS: OS and persistent disks. | [optional] 

## Methods

### NewDesktopPoolViewStorageAcceleratorSettings

`func NewDesktopPoolViewStorageAcceleratorSettings() *DesktopPoolViewStorageAcceleratorSettings`

NewDesktopPoolViewStorageAcceleratorSettings instantiates a new DesktopPoolViewStorageAcceleratorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolViewStorageAcceleratorSettingsWithDefaults

`func NewDesktopPoolViewStorageAcceleratorSettingsWithDefaults() *DesktopPoolViewStorageAcceleratorSettings`

NewDesktopPoolViewStorageAcceleratorSettingsWithDefaults instantiates a new DesktopPoolViewStorageAcceleratorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettings) GetBlackoutTimes() []ViewStorageAcceleratorBlackoutTimeSettings`

GetBlackoutTimes returns the BlackoutTimes field if non-nil, zero value otherwise.

### GetBlackoutTimesOk

`func (o *DesktopPoolViewStorageAcceleratorSettings) GetBlackoutTimesOk() (*[]ViewStorageAcceleratorBlackoutTimeSettings, bool)`

GetBlackoutTimesOk returns a tuple with the BlackoutTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettings) SetBlackoutTimes(v []ViewStorageAcceleratorBlackoutTimeSettings)`

SetBlackoutTimes sets BlackoutTimes field to given value.

### HasBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettings) HasBlackoutTimes() bool`

HasBlackoutTimes returns a boolean if a field has been set.

### GetRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettings) GetRegenerateViewStorageAcceleratorDays() int32`

GetRegenerateViewStorageAcceleratorDays returns the RegenerateViewStorageAcceleratorDays field if non-nil, zero value otherwise.

### GetRegenerateViewStorageAcceleratorDaysOk

`func (o *DesktopPoolViewStorageAcceleratorSettings) GetRegenerateViewStorageAcceleratorDaysOk() (*int32, bool)`

GetRegenerateViewStorageAcceleratorDaysOk returns a tuple with the RegenerateViewStorageAcceleratorDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettings) SetRegenerateViewStorageAcceleratorDays(v int32)`

SetRegenerateViewStorageAcceleratorDays sets RegenerateViewStorageAcceleratorDays field to given value.

### HasRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettings) HasRegenerateViewStorageAcceleratorDays() bool`

HasRegenerateViewStorageAcceleratorDays returns a boolean if a field has been set.

### GetUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettings) GetUseViewStorageAccelerator() bool`

GetUseViewStorageAccelerator returns the UseViewStorageAccelerator field if non-nil, zero value otherwise.

### GetUseViewStorageAcceleratorOk

`func (o *DesktopPoolViewStorageAcceleratorSettings) GetUseViewStorageAcceleratorOk() (*bool, bool)`

GetUseViewStorageAcceleratorOk returns a tuple with the UseViewStorageAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettings) SetUseViewStorageAccelerator(v bool)`

SetUseViewStorageAccelerator sets UseViewStorageAccelerator field to given value.

### HasUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettings) HasUseViewStorageAccelerator() bool`

HasUseViewStorageAccelerator returns a boolean if a field has been set.

### GetViewStorageAcceleratorDiskTypes

`func (o *DesktopPoolViewStorageAcceleratorSettings) GetViewStorageAcceleratorDiskTypes() string`

GetViewStorageAcceleratorDiskTypes returns the ViewStorageAcceleratorDiskTypes field if non-nil, zero value otherwise.

### GetViewStorageAcceleratorDiskTypesOk

`func (o *DesktopPoolViewStorageAcceleratorSettings) GetViewStorageAcceleratorDiskTypesOk() (*string, bool)`

GetViewStorageAcceleratorDiskTypesOk returns a tuple with the ViewStorageAcceleratorDiskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStorageAcceleratorDiskTypes

`func (o *DesktopPoolViewStorageAcceleratorSettings) SetViewStorageAcceleratorDiskTypes(v string)`

SetViewStorageAcceleratorDiskTypes sets ViewStorageAcceleratorDiskTypes field to given value.

### HasViewStorageAcceleratorDiskTypes

`func (o *DesktopPoolViewStorageAcceleratorSettings) HasViewStorageAcceleratorDiskTypes() bool`

HasViewStorageAcceleratorDiskTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


