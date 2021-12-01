# DesktopPoolViewStorageAcceleratorSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutTimes** | Pointer to [**[]ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec**](ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec.md) | A list of blackout times. | [optional] 
**RegenerateViewStorageAcceleratorDays** | Pointer to **int32** | How often to regenerate the View Storage Accelerator cache. Measured in Days. &lt;br&gt; This property is required if useViewStorageAccelerator is set to true. &lt;br&gt; Default value is 7. | [optional] 
**UseViewStorageAccelerator** | Pointer to **bool** | Indicates whether to use View Storage Accelerator. &lt;br&gt; Default value is false. | [optional] 

## Methods

### NewDesktopPoolViewStorageAcceleratorSettingsCreateSpec

`func NewDesktopPoolViewStorageAcceleratorSettingsCreateSpec() *DesktopPoolViewStorageAcceleratorSettingsCreateSpec`

NewDesktopPoolViewStorageAcceleratorSettingsCreateSpec instantiates a new DesktopPoolViewStorageAcceleratorSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolViewStorageAcceleratorSettingsCreateSpecWithDefaults

`func NewDesktopPoolViewStorageAcceleratorSettingsCreateSpecWithDefaults() *DesktopPoolViewStorageAcceleratorSettingsCreateSpec`

NewDesktopPoolViewStorageAcceleratorSettingsCreateSpecWithDefaults instantiates a new DesktopPoolViewStorageAcceleratorSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetBlackoutTimes() []ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec`

GetBlackoutTimes returns the BlackoutTimes field if non-nil, zero value otherwise.

### GetBlackoutTimesOk

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetBlackoutTimesOk() (*[]ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec, bool)`

GetBlackoutTimesOk returns a tuple with the BlackoutTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) SetBlackoutTimes(v []ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec)`

SetBlackoutTimes sets BlackoutTimes field to given value.

### HasBlackoutTimes

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) HasBlackoutTimes() bool`

HasBlackoutTimes returns a boolean if a field has been set.

### GetRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetRegenerateViewStorageAcceleratorDays() int32`

GetRegenerateViewStorageAcceleratorDays returns the RegenerateViewStorageAcceleratorDays field if non-nil, zero value otherwise.

### GetRegenerateViewStorageAcceleratorDaysOk

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetRegenerateViewStorageAcceleratorDaysOk() (*int32, bool)`

GetRegenerateViewStorageAcceleratorDaysOk returns a tuple with the RegenerateViewStorageAcceleratorDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) SetRegenerateViewStorageAcceleratorDays(v int32)`

SetRegenerateViewStorageAcceleratorDays sets RegenerateViewStorageAcceleratorDays field to given value.

### HasRegenerateViewStorageAcceleratorDays

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) HasRegenerateViewStorageAcceleratorDays() bool`

HasRegenerateViewStorageAcceleratorDays returns a boolean if a field has been set.

### GetUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetUseViewStorageAccelerator() bool`

GetUseViewStorageAccelerator returns the UseViewStorageAccelerator field if non-nil, zero value otherwise.

### GetUseViewStorageAcceleratorOk

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetUseViewStorageAcceleratorOk() (*bool, bool)`

GetUseViewStorageAcceleratorOk returns a tuple with the UseViewStorageAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) SetUseViewStorageAccelerator(v bool)`

SetUseViewStorageAccelerator sets UseViewStorageAccelerator field to given value.

### HasUseViewStorageAccelerator

`func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) HasUseViewStorageAccelerator() bool`

HasUseViewStorageAccelerator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


