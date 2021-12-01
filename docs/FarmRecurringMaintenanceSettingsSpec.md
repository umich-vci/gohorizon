# FarmRecurringMaintenanceSettingsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenancePeriod** | **string** | This represents the frequency at which to perform recurring maintenance. * DAILY: Daily recurring maintenance * WEEKLY: Weekly recurring maintenance * MONTHLY: Monthly recurring maintenance | 
**MaintenancePeriodFrequency** | Pointer to **int32** | Indicates how frequently to repeat maintenance, expressed as multiple of the maintenance period. e.g. Every 2 weeks. Default value is 1. | [optional] 
**StartIndex** | Pointer to **int32** | Start index for weekly or monthly maintenance.This required if maintenance_period is set to either WEEKLY or MONTHLY. Weekly: 1-7 (Sun-Sat), Monthly: 1-31 | [optional] 
**StartTime** | **string** | Start time configured for the recurring maintenance. This should be in the form hh:mm in 24 hours format. | 

## Methods

### NewFarmRecurringMaintenanceSettingsSpec

`func NewFarmRecurringMaintenanceSettingsSpec(maintenancePeriod string, startTime string, ) *FarmRecurringMaintenanceSettingsSpec`

NewFarmRecurringMaintenanceSettingsSpec instantiates a new FarmRecurringMaintenanceSettingsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmRecurringMaintenanceSettingsSpecWithDefaults

`func NewFarmRecurringMaintenanceSettingsSpecWithDefaults() *FarmRecurringMaintenanceSettingsSpec`

NewFarmRecurringMaintenanceSettingsSpecWithDefaults instantiates a new FarmRecurringMaintenanceSettingsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenancePeriod

`func (o *FarmRecurringMaintenanceSettingsSpec) GetMaintenancePeriod() string`

GetMaintenancePeriod returns the MaintenancePeriod field if non-nil, zero value otherwise.

### GetMaintenancePeriodOk

`func (o *FarmRecurringMaintenanceSettingsSpec) GetMaintenancePeriodOk() (*string, bool)`

GetMaintenancePeriodOk returns a tuple with the MaintenancePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePeriod

`func (o *FarmRecurringMaintenanceSettingsSpec) SetMaintenancePeriod(v string)`

SetMaintenancePeriod sets MaintenancePeriod field to given value.


### GetMaintenancePeriodFrequency

`func (o *FarmRecurringMaintenanceSettingsSpec) GetMaintenancePeriodFrequency() int32`

GetMaintenancePeriodFrequency returns the MaintenancePeriodFrequency field if non-nil, zero value otherwise.

### GetMaintenancePeriodFrequencyOk

`func (o *FarmRecurringMaintenanceSettingsSpec) GetMaintenancePeriodFrequencyOk() (*int32, bool)`

GetMaintenancePeriodFrequencyOk returns a tuple with the MaintenancePeriodFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePeriodFrequency

`func (o *FarmRecurringMaintenanceSettingsSpec) SetMaintenancePeriodFrequency(v int32)`

SetMaintenancePeriodFrequency sets MaintenancePeriodFrequency field to given value.

### HasMaintenancePeriodFrequency

`func (o *FarmRecurringMaintenanceSettingsSpec) HasMaintenancePeriodFrequency() bool`

HasMaintenancePeriodFrequency returns a boolean if a field has been set.

### GetStartIndex

`func (o *FarmRecurringMaintenanceSettingsSpec) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *FarmRecurringMaintenanceSettingsSpec) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *FarmRecurringMaintenanceSettingsSpec) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *FarmRecurringMaintenanceSettingsSpec) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetStartTime

`func (o *FarmRecurringMaintenanceSettingsSpec) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FarmRecurringMaintenanceSettingsSpec) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FarmRecurringMaintenanceSettingsSpec) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


