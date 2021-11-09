# DesktopPoolVirtualMachineSpecifiedNamingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumMachines** | Pointer to **int32** | Number of machines in the pool. | [optional] 
**NumUnassignedMachinesKeptPoweredOn** | Pointer to **int32** | Number of unassigned machines kept powered on. This value must be less than or equal to the number of specified names. | [optional] 
**StartMachinesInMaintenanceMode** | Pointer to **bool** | Allows virtual machines to be customized manually before users can log in and access them. This mode must be exited manually. | [optional] 

## Methods

### NewDesktopPoolVirtualMachineSpecifiedNamingSettings

`func NewDesktopPoolVirtualMachineSpecifiedNamingSettings() *DesktopPoolVirtualMachineSpecifiedNamingSettings`

NewDesktopPoolVirtualMachineSpecifiedNamingSettings instantiates a new DesktopPoolVirtualMachineSpecifiedNamingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolVirtualMachineSpecifiedNamingSettingsWithDefaults

`func NewDesktopPoolVirtualMachineSpecifiedNamingSettingsWithDefaults() *DesktopPoolVirtualMachineSpecifiedNamingSettings`

NewDesktopPoolVirtualMachineSpecifiedNamingSettingsWithDefaults instantiates a new DesktopPoolVirtualMachineSpecifiedNamingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumMachines

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.

### HasNumMachines

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) HasNumMachines() bool`

HasNumMachines returns a boolean if a field has been set.

### GetNumUnassignedMachinesKeptPoweredOn

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) GetNumUnassignedMachinesKeptPoweredOn() int32`

GetNumUnassignedMachinesKeptPoweredOn returns the NumUnassignedMachinesKeptPoweredOn field if non-nil, zero value otherwise.

### GetNumUnassignedMachinesKeptPoweredOnOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) GetNumUnassignedMachinesKeptPoweredOnOk() (*int32, bool)`

GetNumUnassignedMachinesKeptPoweredOnOk returns a tuple with the NumUnassignedMachinesKeptPoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnassignedMachinesKeptPoweredOn

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) SetNumUnassignedMachinesKeptPoweredOn(v int32)`

SetNumUnassignedMachinesKeptPoweredOn sets NumUnassignedMachinesKeptPoweredOn field to given value.

### HasNumUnassignedMachinesKeptPoweredOn

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) HasNumUnassignedMachinesKeptPoweredOn() bool`

HasNumUnassignedMachinesKeptPoweredOn returns a boolean if a field has been set.

### GetStartMachinesInMaintenanceMode

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) GetStartMachinesInMaintenanceMode() bool`

GetStartMachinesInMaintenanceMode returns the StartMachinesInMaintenanceMode field if non-nil, zero value otherwise.

### GetStartMachinesInMaintenanceModeOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) GetStartMachinesInMaintenanceModeOk() (*bool, bool)`

GetStartMachinesInMaintenanceModeOk returns a tuple with the StartMachinesInMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMachinesInMaintenanceMode

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) SetStartMachinesInMaintenanceMode(v bool)`

SetStartMachinesInMaintenanceMode sets StartMachinesInMaintenanceMode field to given value.

### HasStartMachinesInMaintenanceMode

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettings) HasStartMachinesInMaintenanceMode() bool`

HasStartMachinesInMaintenanceMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


