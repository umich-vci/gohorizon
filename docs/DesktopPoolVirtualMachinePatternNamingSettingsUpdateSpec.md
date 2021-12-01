# DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberOfMachines** | **int32** | Maximum number of machines in the desktop pool. | 
**MinNumberOfMachines** | Pointer to **int32** | The minimum number of machines to have provisioned if on demand provisioning is selected. This property is required if provisioning_time is set to \&quot;ON_DEMAND\&quot; . | [optional] 
**NamingPattern** | **string** | Virtual machines will be named according to the specified naming pattern. Horizon appends a unique number to the specified pattern to provide a unique name for each virtual machine. To place this unique number elsewhere in the pattern, use &#39;{n}&#39;. (For example: vm-{n}-sales.) The unique number can also be made a fixed length. (For example: vm-{n:fixed&#x3D;3}-sales will name VMs from vm-001-sales to vm-999-sales). Machine names are constrained to a maximum size of 15 characters including the unique number. Therefore, care must be taken when choosing a pattern. If the maximum desktop pool size is 9 machines, the pattern must be at most 14 characters. For 99 machines, 13 characters, for 999 machines, 12 characters. For 9999 machines, 11 characters. If using a fixed size token, use a maximum of 14 characters for \&quot;n&#x3D;1\&quot;, 13 characters for \&quot;n&#x3D;2\&quot;, 12 characters for \&quot;n&#x3D;3\&quot;, and 11 characters for \&quot;n&#x3D;4\&quot;.If {n} is specified with no size, a size of 2 is automatically used and if no {} is specified, {n&#x3D;2} is automatically appended to the end of the pattern. This property must contain only alphanumerics and dashes. It must contain at least one alpha character. It may also optionally contain a numeric placement token {n} or {n:fixed&#x3D;#}. If the pattern does not specify the numeric placement token, the maximum length is 14 characters. | 
**NumberOfSpareMachines** | **int32** | Number of spare powered on machines. | 
**ProvisioningTime** | **string** | Determines when the machines are provisioned. * ON_DEMAND: Provision machines on demand. * UP_FRONT: Provision all machines up-front. | 

## Methods

### NewDesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec

`func NewDesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec(maxNumberOfMachines int32, namingPattern string, numberOfSpareMachines int32, provisioningTime string, ) *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec`

NewDesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec instantiates a new DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolVirtualMachinePatternNamingSettingsUpdateSpecWithDefaults

`func NewDesktopPoolVirtualMachinePatternNamingSettingsUpdateSpecWithDefaults() *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec`

NewDesktopPoolVirtualMachinePatternNamingSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetMaxNumberOfMachines() int32`

GetMaxNumberOfMachines returns the MaxNumberOfMachines field if non-nil, zero value otherwise.

### GetMaxNumberOfMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetMaxNumberOfMachinesOk() (*int32, bool)`

GetMaxNumberOfMachinesOk returns a tuple with the MaxNumberOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) SetMaxNumberOfMachines(v int32)`

SetMaxNumberOfMachines sets MaxNumberOfMachines field to given value.


### GetMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetMinNumberOfMachines() int32`

GetMinNumberOfMachines returns the MinNumberOfMachines field if non-nil, zero value otherwise.

### GetMinNumberOfMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetMinNumberOfMachinesOk() (*int32, bool)`

GetMinNumberOfMachinesOk returns a tuple with the MinNumberOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) SetMinNumberOfMachines(v int32)`

SetMinNumberOfMachines sets MinNumberOfMachines field to given value.

### HasMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) HasMinNumberOfMachines() bool`

HasMinNumberOfMachines returns a boolean if a field has been set.

### GetNamingPattern

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.


### GetNumberOfSpareMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetNumberOfSpareMachines() int32`

GetNumberOfSpareMachines returns the NumberOfSpareMachines field if non-nil, zero value otherwise.

### GetNumberOfSpareMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetNumberOfSpareMachinesOk() (*int32, bool)`

GetNumberOfSpareMachinesOk returns a tuple with the NumberOfSpareMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpareMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) SetNumberOfSpareMachines(v int32)`

SetNumberOfSpareMachines sets NumberOfSpareMachines field to given value.


### GetProvisioningTime

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetProvisioningTime() string`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) GetProvisioningTimeOk() (*string, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec) SetProvisioningTime(v string)`

SetProvisioningTime sets ProvisioningTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


