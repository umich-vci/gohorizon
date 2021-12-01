# DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumUnassignedMachinesKeptPoweredOn** | **int32** | Number of unassigned machines kept powered on. This value must be lesser than or equal to the number of specified names. | 
**SpecifiedNames** | Pointer to [**[]MachineSpecifiedName**](MachineSpecifiedName.md) | Initial specified names of machines in the desktop pool. | [optional] 
**StartMachinesInMaintenanceMode** | **bool** | Allows virtual machines to be customized manually before users can log in and access them. This mode must be exited manually. | 

## Methods

### NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec

`func NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec(numUnassignedMachinesKeptPoweredOn int32, startMachinesInMaintenanceMode bool, ) *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec`

NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec instantiates a new DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpecWithDefaults

`func NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpecWithDefaults() *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec`

NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumUnassignedMachinesKeptPoweredOn

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetNumUnassignedMachinesKeptPoweredOn() int32`

GetNumUnassignedMachinesKeptPoweredOn returns the NumUnassignedMachinesKeptPoweredOn field if non-nil, zero value otherwise.

### GetNumUnassignedMachinesKeptPoweredOnOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetNumUnassignedMachinesKeptPoweredOnOk() (*int32, bool)`

GetNumUnassignedMachinesKeptPoweredOnOk returns a tuple with the NumUnassignedMachinesKeptPoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnassignedMachinesKeptPoweredOn

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) SetNumUnassignedMachinesKeptPoweredOn(v int32)`

SetNumUnassignedMachinesKeptPoweredOn sets NumUnassignedMachinesKeptPoweredOn field to given value.


### GetSpecifiedNames

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetSpecifiedNames() []MachineSpecifiedName`

GetSpecifiedNames returns the SpecifiedNames field if non-nil, zero value otherwise.

### GetSpecifiedNamesOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetSpecifiedNamesOk() (*[]MachineSpecifiedName, bool)`

GetSpecifiedNamesOk returns a tuple with the SpecifiedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedNames

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) SetSpecifiedNames(v []MachineSpecifiedName)`

SetSpecifiedNames sets SpecifiedNames field to given value.

### HasSpecifiedNames

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) HasSpecifiedNames() bool`

HasSpecifiedNames returns a boolean if a field has been set.

### GetStartMachinesInMaintenanceMode

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetStartMachinesInMaintenanceMode() bool`

GetStartMachinesInMaintenanceMode returns the StartMachinesInMaintenanceMode field if non-nil, zero value otherwise.

### GetStartMachinesInMaintenanceModeOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetStartMachinesInMaintenanceModeOk() (*bool, bool)`

GetStartMachinesInMaintenanceModeOk returns a tuple with the StartMachinesInMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMachinesInMaintenanceMode

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) SetStartMachinesInMaintenanceMode(v bool)`

SetStartMachinesInMaintenanceMode sets StartMachinesInMaintenanceMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


