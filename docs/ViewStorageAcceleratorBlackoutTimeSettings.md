# ViewStorageAcceleratorBlackoutTimeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **[]string** | List of days for a given range of time. | [optional] 
**EndTime** | Pointer to **string** | Ending time for the blackout in 24-hour format. | [optional] 
**StartTime** | Pointer to **string** | Starting time for the blackout in 24-hour format. | [optional] 

## Methods

### NewViewStorageAcceleratorBlackoutTimeSettings

`func NewViewStorageAcceleratorBlackoutTimeSettings() *ViewStorageAcceleratorBlackoutTimeSettings`

NewViewStorageAcceleratorBlackoutTimeSettings instantiates a new ViewStorageAcceleratorBlackoutTimeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStorageAcceleratorBlackoutTimeSettingsWithDefaults

`func NewViewStorageAcceleratorBlackoutTimeSettingsWithDefaults() *ViewStorageAcceleratorBlackoutTimeSettings`

NewViewStorageAcceleratorBlackoutTimeSettingsWithDefaults instantiates a new ViewStorageAcceleratorBlackoutTimeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetEndTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ViewStorageAcceleratorBlackoutTimeSettings) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


