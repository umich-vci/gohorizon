# ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **[]string** | List of days for a given range of time. | 
**EndTime** | **string** | Ending time for the blackout in 24-hour format. | 
**StartTime** | **string** | Starting time for the blackout in 24-hour format. | 

## Methods

### NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpec

`func NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpec(days []string, endTime string, startTime string, ) *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec`

NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpec instantiates a new ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpecWithDefaults

`func NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpecWithDefaults() *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec`

NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpecWithDefaults instantiates a new ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) SetDays(v []string)`

SetDays sets Days field to given value.


### GetEndTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetStartTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


