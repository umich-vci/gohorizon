# DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumUnassignedMachinesKeptPoweredOn** | Pointer to **int32** | Number of unassigned machines kept powered on. This property must be lesser than or equal to the number of specified names.&lt;br&gt; Default value is 1. | [optional] 
**SpecifiedNames** | Pointer to [**[]MachineSpecifiedName**](MachineSpecifiedName.md) | Specified names for the virtual machines. | [optional] 
**StartMachinesInMaintenanceMode** | Pointer to **bool** | Allows virtual machines to be customized manually before users can log in and access them. This mode must be exited manually. &lt;br&gt; Default value is false. | [optional] 

## Methods

### NewDesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec

`func NewDesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec() *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec`

NewDesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec instantiates a new DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpecWithDefaults

`func NewDesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpecWithDefaults() *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec`

NewDesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpecWithDefaults instantiates a new DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumUnassignedMachinesKeptPoweredOn

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) GetNumUnassignedMachinesKeptPoweredOn() int32`

GetNumUnassignedMachinesKeptPoweredOn returns the NumUnassignedMachinesKeptPoweredOn field if non-nil, zero value otherwise.

### GetNumUnassignedMachinesKeptPoweredOnOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) GetNumUnassignedMachinesKeptPoweredOnOk() (*int32, bool)`

GetNumUnassignedMachinesKeptPoweredOnOk returns a tuple with the NumUnassignedMachinesKeptPoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnassignedMachinesKeptPoweredOn

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) SetNumUnassignedMachinesKeptPoweredOn(v int32)`

SetNumUnassignedMachinesKeptPoweredOn sets NumUnassignedMachinesKeptPoweredOn field to given value.

### HasNumUnassignedMachinesKeptPoweredOn

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) HasNumUnassignedMachinesKeptPoweredOn() bool`

HasNumUnassignedMachinesKeptPoweredOn returns a boolean if a field has been set.

### GetSpecifiedNames

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) GetSpecifiedNames() []MachineSpecifiedName`

GetSpecifiedNames returns the SpecifiedNames field if non-nil, zero value otherwise.

### GetSpecifiedNamesOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) GetSpecifiedNamesOk() (*[]MachineSpecifiedName, bool)`

GetSpecifiedNamesOk returns a tuple with the SpecifiedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedNames

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) SetSpecifiedNames(v []MachineSpecifiedName)`

SetSpecifiedNames sets SpecifiedNames field to given value.

### HasSpecifiedNames

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) HasSpecifiedNames() bool`

HasSpecifiedNames returns a boolean if a field has been set.

### GetStartMachinesInMaintenanceMode

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) GetStartMachinesInMaintenanceMode() bool`

GetStartMachinesInMaintenanceMode returns the StartMachinesInMaintenanceMode field if non-nil, zero value otherwise.

### GetStartMachinesInMaintenanceModeOk

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) GetStartMachinesInMaintenanceModeOk() (*bool, bool)`

GetStartMachinesInMaintenanceModeOk returns a tuple with the StartMachinesInMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMachinesInMaintenanceMode

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) SetStartMachinesInMaintenanceMode(v bool)`

SetStartMachinesInMaintenanceMode sets StartMachinesInMaintenanceMode field to given value.

### HasStartMachinesInMaintenanceMode

`func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsCreateSpec) HasStartMachinesInMaintenanceMode() bool`

HasStartMachinesInMaintenanceMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


