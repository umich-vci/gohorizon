# FarmRecurringMaintenanceSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenancePeriod** | Pointer to **string** | Settings for recurring maintenance operations. * DAILY: Daily recurring maintenance * WEEKLY: Weekly recurring maintenance * MONTHLY: Monthly recurring maintenance | [optional] 
**MaintenancePeriodFrequency** | Pointer to **int32** | Indicates frequency of repeating maintenance and is expressed as a multiple of the maintenance_period.  | [optional] 
**StartIndex** | Pointer to **int32** | Start index for weekly or monthly maintenance. Weekly: 1-7 (Sun-Sat), Monthly: 1-31. This is set when maintenance_period is WEEKLY or MONTHLY. | [optional] 
**StartTime** | Pointer to **string** | Start time configured for the recurring maintenance. This is in the form hh:mm in 24 hours format | [optional] 

## Methods

### NewFarmRecurringMaintenanceSettingsInfo

`func NewFarmRecurringMaintenanceSettingsInfo() *FarmRecurringMaintenanceSettingsInfo`

NewFarmRecurringMaintenanceSettingsInfo instantiates a new FarmRecurringMaintenanceSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmRecurringMaintenanceSettingsInfoWithDefaults

`func NewFarmRecurringMaintenanceSettingsInfoWithDefaults() *FarmRecurringMaintenanceSettingsInfo`

NewFarmRecurringMaintenanceSettingsInfoWithDefaults instantiates a new FarmRecurringMaintenanceSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenancePeriod

`func (o *FarmRecurringMaintenanceSettingsInfo) GetMaintenancePeriod() string`

GetMaintenancePeriod returns the MaintenancePeriod field if non-nil, zero value otherwise.

### GetMaintenancePeriodOk

`func (o *FarmRecurringMaintenanceSettingsInfo) GetMaintenancePeriodOk() (*string, bool)`

GetMaintenancePeriodOk returns a tuple with the MaintenancePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePeriod

`func (o *FarmRecurringMaintenanceSettingsInfo) SetMaintenancePeriod(v string)`

SetMaintenancePeriod sets MaintenancePeriod field to given value.

### HasMaintenancePeriod

`func (o *FarmRecurringMaintenanceSettingsInfo) HasMaintenancePeriod() bool`

HasMaintenancePeriod returns a boolean if a field has been set.

### GetMaintenancePeriodFrequency

`func (o *FarmRecurringMaintenanceSettingsInfo) GetMaintenancePeriodFrequency() int32`

GetMaintenancePeriodFrequency returns the MaintenancePeriodFrequency field if non-nil, zero value otherwise.

### GetMaintenancePeriodFrequencyOk

`func (o *FarmRecurringMaintenanceSettingsInfo) GetMaintenancePeriodFrequencyOk() (*int32, bool)`

GetMaintenancePeriodFrequencyOk returns a tuple with the MaintenancePeriodFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePeriodFrequency

`func (o *FarmRecurringMaintenanceSettingsInfo) SetMaintenancePeriodFrequency(v int32)`

SetMaintenancePeriodFrequency sets MaintenancePeriodFrequency field to given value.

### HasMaintenancePeriodFrequency

`func (o *FarmRecurringMaintenanceSettingsInfo) HasMaintenancePeriodFrequency() bool`

HasMaintenancePeriodFrequency returns a boolean if a field has been set.

### GetStartIndex

`func (o *FarmRecurringMaintenanceSettingsInfo) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *FarmRecurringMaintenanceSettingsInfo) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *FarmRecurringMaintenanceSettingsInfo) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *FarmRecurringMaintenanceSettingsInfo) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetStartTime

`func (o *FarmRecurringMaintenanceSettingsInfo) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FarmRecurringMaintenanceSettingsInfo) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FarmRecurringMaintenanceSettingsInfo) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *FarmRecurringMaintenanceSettingsInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


