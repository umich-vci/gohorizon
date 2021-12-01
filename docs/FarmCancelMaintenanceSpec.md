# FarmCancelMaintenanceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceMode** | **string** | Scheduled maintenance mode to be cancelled. * IMMEDIATE: All server VMs will be refreshed once, immediately or at user scheduled time. * RECURRING: All server VMs will be periodically refreshed based on FarmInstantCloneRecurringMaintenancePeriod and StartTime | 

## Methods

### NewFarmCancelMaintenanceSpec

`func NewFarmCancelMaintenanceSpec(maintenanceMode string, ) *FarmCancelMaintenanceSpec`

NewFarmCancelMaintenanceSpec instantiates a new FarmCancelMaintenanceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmCancelMaintenanceSpecWithDefaults

`func NewFarmCancelMaintenanceSpecWithDefaults() *FarmCancelMaintenanceSpec`

NewFarmCancelMaintenanceSpecWithDefaults instantiates a new FarmCancelMaintenanceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceMode

`func (o *FarmCancelMaintenanceSpec) GetMaintenanceMode() string`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *FarmCancelMaintenanceSpec) GetMaintenanceModeOk() (*string, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *FarmCancelMaintenanceSpec) SetMaintenanceMode(v string)`

SetMaintenanceMode sets MaintenanceMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


