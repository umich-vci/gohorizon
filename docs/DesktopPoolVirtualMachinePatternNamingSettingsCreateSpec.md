# DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberOfMachines** | Pointer to **int32** | Maximum number of machines in the desktop pool. &lt;br&gt; Default value is 1. | [optional] 
**MinNumberOfMachines** | Pointer to **int32** | This is applicable if provisioning_time is set to ON_DEMAND with default value of 0. &lt;br&gt; | [optional] 
**NamingPattern** | **string** | Virtual machines will be named according to the specified naming pattern.&lt;br&gt; By default, view manager appends a unique number to the specified pattern to provide a unique name for each virtual machine. To place this unique number elsewhere in the pattern, use &#39;{n}&#39;. (For example: vm-{n}-sales.) The unique number can also be made a fixed length. (For example: vm-{n:fixed&#x3D;3}-sales will name VMs from vm-001-sales to vm-999-sales). &lt;br&gt;Machine names are constrained to a maximum size of 15 characters including the unique number. Therefore, care must be taken when choosing a pattern. If the maximum desktop size is 9 machines, the pattern must be at most 14 characters. For 99 machines, 13 characters, for 999 machines, 12 characters. For 9999 machines, 11 characters. If using a fixed size token, use a maximum of 14 characters for \&quot;n&#x3D;1\&quot;, 13 characters for \&quot;n&#x3D;2\&quot;, 12 characters for \&quot;n&#x3D;3\&quot;, and 11 characters for \&quot;n&#x3D;4\&quot;.&lt;br&gt;If {n} is specified with no size, a size of 2 is automatically used and if no {} is specified, {n&#x3D;2} is automatically appended to the end of the pattern. | 
**NumberOfSpareMachines** | Pointer to **int32** | Number of spare powered on machines. &lt;br&gt; Default value is 1. | [optional] 
**ProvisioningTime** | Pointer to **string** | Determines when the machines are provisioned. &lt;br&gt; Default value is UP_FRONT * ON_DEMAND: Provision machines on demand. * UP_FRONT: Provision all machines up-front. | [optional] 

## Methods

### NewDesktopPoolVirtualMachinePatternNamingSettingsCreateSpec

`func NewDesktopPoolVirtualMachinePatternNamingSettingsCreateSpec(namingPattern string, ) *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec`

NewDesktopPoolVirtualMachinePatternNamingSettingsCreateSpec instantiates a new DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolVirtualMachinePatternNamingSettingsCreateSpecWithDefaults

`func NewDesktopPoolVirtualMachinePatternNamingSettingsCreateSpecWithDefaults() *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec`

NewDesktopPoolVirtualMachinePatternNamingSettingsCreateSpecWithDefaults instantiates a new DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetMaxNumberOfMachines() int32`

GetMaxNumberOfMachines returns the MaxNumberOfMachines field if non-nil, zero value otherwise.

### GetMaxNumberOfMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetMaxNumberOfMachinesOk() (*int32, bool)`

GetMaxNumberOfMachinesOk returns a tuple with the MaxNumberOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) SetMaxNumberOfMachines(v int32)`

SetMaxNumberOfMachines sets MaxNumberOfMachines field to given value.

### HasMaxNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) HasMaxNumberOfMachines() bool`

HasMaxNumberOfMachines returns a boolean if a field has been set.

### GetMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetMinNumberOfMachines() int32`

GetMinNumberOfMachines returns the MinNumberOfMachines field if non-nil, zero value otherwise.

### GetMinNumberOfMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetMinNumberOfMachinesOk() (*int32, bool)`

GetMinNumberOfMachinesOk returns a tuple with the MinNumberOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) SetMinNumberOfMachines(v int32)`

SetMinNumberOfMachines sets MinNumberOfMachines field to given value.

### HasMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) HasMinNumberOfMachines() bool`

HasMinNumberOfMachines returns a boolean if a field has been set.

### GetNamingPattern

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.


### GetNumberOfSpareMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetNumberOfSpareMachines() int32`

GetNumberOfSpareMachines returns the NumberOfSpareMachines field if non-nil, zero value otherwise.

### GetNumberOfSpareMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetNumberOfSpareMachinesOk() (*int32, bool)`

GetNumberOfSpareMachinesOk returns a tuple with the NumberOfSpareMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpareMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) SetNumberOfSpareMachines(v int32)`

SetNumberOfSpareMachines sets NumberOfSpareMachines field to given value.

### HasNumberOfSpareMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) HasNumberOfSpareMachines() bool`

HasNumberOfSpareMachines returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetProvisioningTime() string`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) GetProvisioningTimeOk() (*string, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) SetProvisioningTime(v string)`

SetProvisioningTime sets ProvisioningTime field to given value.

### HasProvisioningTime

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsCreateSpec) HasProvisioningTime() bool`

HasProvisioningTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


