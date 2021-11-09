# DesktopPoolVirtualMachinePatternNamingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberOfMachines** | Pointer to **int32** | Maximum number of machines in the desktop pool. | [optional] 
**MinNumberOfMachines** | Pointer to **int32** | The minimum number of machines to have provisioned if on demand provisioning is selected. | [optional] 
**NamingPattern** | Pointer to **string** | Virtual machines will be named according to the specified naming pattern.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**NumberOfSpareMachines** | Pointer to **int32** | Number of spare powered on machines. | [optional] 
**ProvisioningTime** | Pointer to **string** | Determines when the machines are provisioned. * ON_DEMAND: Provision machines on demand. * UP_FRONT: Provision all machines up-front. | [optional] 

## Methods

### NewDesktopPoolVirtualMachinePatternNamingSettings

`func NewDesktopPoolVirtualMachinePatternNamingSettings() *DesktopPoolVirtualMachinePatternNamingSettings`

NewDesktopPoolVirtualMachinePatternNamingSettings instantiates a new DesktopPoolVirtualMachinePatternNamingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolVirtualMachinePatternNamingSettingsWithDefaults

`func NewDesktopPoolVirtualMachinePatternNamingSettingsWithDefaults() *DesktopPoolVirtualMachinePatternNamingSettings`

NewDesktopPoolVirtualMachinePatternNamingSettingsWithDefaults instantiates a new DesktopPoolVirtualMachinePatternNamingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetMaxNumberOfMachines() int32`

GetMaxNumberOfMachines returns the MaxNumberOfMachines field if non-nil, zero value otherwise.

### GetMaxNumberOfMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetMaxNumberOfMachinesOk() (*int32, bool)`

GetMaxNumberOfMachinesOk returns a tuple with the MaxNumberOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetMaxNumberOfMachines(v int32)`

SetMaxNumberOfMachines sets MaxNumberOfMachines field to given value.

### HasMaxNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasMaxNumberOfMachines() bool`

HasMaxNumberOfMachines returns a boolean if a field has been set.

### GetMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetMinNumberOfMachines() int32`

GetMinNumberOfMachines returns the MinNumberOfMachines field if non-nil, zero value otherwise.

### GetMinNumberOfMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetMinNumberOfMachinesOk() (*int32, bool)`

GetMinNumberOfMachinesOk returns a tuple with the MinNumberOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetMinNumberOfMachines(v int32)`

SetMinNumberOfMachines sets MinNumberOfMachines field to given value.

### HasMinNumberOfMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasMinNumberOfMachines() bool`

HasMinNumberOfMachines returns a boolean if a field has been set.

### GetNamingPattern

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.

### GetNumberOfSpareMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetNumberOfSpareMachines() int32`

GetNumberOfSpareMachines returns the NumberOfSpareMachines field if non-nil, zero value otherwise.

### GetNumberOfSpareMachinesOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetNumberOfSpareMachinesOk() (*int32, bool)`

GetNumberOfSpareMachinesOk returns a tuple with the NumberOfSpareMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpareMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetNumberOfSpareMachines(v int32)`

SetNumberOfSpareMachines sets NumberOfSpareMachines field to given value.

### HasNumberOfSpareMachines

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasNumberOfSpareMachines() bool`

HasNumberOfSpareMachines returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetProvisioningTime() string`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetProvisioningTimeOk() (*string, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetProvisioningTime(v string)`

SetProvisioningTime sets ProvisioningTime field to given value.

### HasProvisioningTime

`func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasProvisioningTime() bool`

HasProvisioningTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


