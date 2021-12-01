# DesktopPoolViewStorageAcceleratorSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutTimes** | Pointer to [**[]ViewStorageAcceleratorBlackoutTimeSettingsUpdateSpec**](ViewStorageAcceleratorBlackoutTimeSettingsUpdateSpec.md) | A list of blackout times. | [optional] 
**RegenerateViewStorageAcceleratorDays** | Pointer to **int32** | How often to regenerate the View Storage Accelerator cache. Measured in Days. This property has a default value of 7. | [optional] 
**UseViewStorageAccelerator** | Pointer to **bool** | Indicates whether to use View Storage Accelerator. | [optional] 

## Methods

### NewDesktopPoolViewStorageAcceleratorSettingsUpdateSpec

`func NewDesktopPoolViewStorageAcceleratorSettingsUpdateSpec() *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec`

NewDesktopPoolViewStorageAcceleratorSettingsUpdateSpec instantiates a new DesktopPoolViewStorageAcceleratorSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolViewStorageAcceleratorSettingsUpdateSpecWithDefaults

`func NewDesktopPoolViewStorageAcceleratorSettingsUpdateSpecWithDefaults() *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec`

NewDesktopPoolViewStorageAcceleratorSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolViewStorageAcceleratorSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) GetBlackoutTimes() []ViewStorageAcceleratorBlackoutTimeSettingsUpdateSpec`

GetBlackoutTimes returns the BlackoutTimes field if non-nil, zero value otherwise.

### GetBlackoutTimesOk

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) GetBlackoutTimesOk() (*[]ViewStorageAcceleratorBlackoutTimeSettingsUpdateSpec, bool)`

GetBlackoutTimesOk returns a tuple with the BlackoutTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) SetBlackoutTimes(v []ViewStorageAcceleratorBlackoutTimeSettingsUpdateSpec)`

SetBlackoutTimes sets BlackoutTimes field to given value.

### HasBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) HasBlackoutTimes() bool`

HasBlackoutTimes returns a boolean if a field has been set.

### GetRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) GetRegenerateViewStorageAcceleratorDays() int32`

GetRegenerateViewStorageAcceleratorDays returns the RegenerateViewStorageAcceleratorDays field if non-nil, zero value otherwise.

### GetRegenerateViewStorageAcceleratorDaysOk

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) GetRegenerateViewStorageAcceleratorDaysOk() (*int32, bool)`

GetRegenerateViewStorageAcceleratorDaysOk returns a tuple with the RegenerateViewStorageAcceleratorDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) SetRegenerateViewStorageAcceleratorDays(v int32)`

SetRegenerateViewStorageAcceleratorDays sets RegenerateViewStorageAcceleratorDays field to given value.

### HasRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) HasRegenerateViewStorageAcceleratorDays() bool`

HasRegenerateViewStorageAcceleratorDays returns a boolean if a field has been set.

### GetUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) GetUseViewStorageAccelerator() bool`

GetUseViewStorageAccelerator returns the UseViewStorageAccelerator field if non-nil, zero value otherwise.

### GetUseViewStorageAcceleratorOk

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) GetUseViewStorageAcceleratorOk() (*bool, bool)`

GetUseViewStorageAcceleratorOk returns a tuple with the UseViewStorageAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) SetUseViewStorageAccelerator(v bool)`

SetUseViewStorageAccelerator sets UseViewStorageAccelerator field to given value.

### HasUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettingsUpdateSpec) HasUseViewStorageAccelerator() bool`

HasUseViewStorageAccelerator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


