# FarmCloneprepCustomizationSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostSynchronizationScriptName** | Pointer to **string** | Post synchronization script. ClonePrep can run a customization script on instant-clone machines after they are created or recovered or a new image is pushed. Provide the path to the script on the parent virtual machine. | [optional] 
**PostSynchronizationScriptParameters** | Pointer to **string** | Post synchronization script parameters. | [optional] 
**PowerOffScriptName** | Pointer to **string** | Power off script. ClonePrep can run a customization script on instant-clone machines before they are powered off. Provide the path to the script on the parent virtual machine. | [optional] 
**PowerOffScriptParameters** | Pointer to **string** | Power off script parameters. | [optional] 
**PrimingComputerAccount** | Pointer to **string** | Instant Clone publishing needs an additional computer account in the same AD domain as the clones. This field accepts the pre-created computer accounts. This field accepts the pre-created computer accounts. This property is ignored when reuse_pre_existing_accounts is set to false. | [optional] 

## Methods

### NewFarmCloneprepCustomizationSettingsUpdateSpec

`func NewFarmCloneprepCustomizationSettingsUpdateSpec() *FarmCloneprepCustomizationSettingsUpdateSpec`

NewFarmCloneprepCustomizationSettingsUpdateSpec instantiates a new FarmCloneprepCustomizationSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmCloneprepCustomizationSettingsUpdateSpecWithDefaults

`func NewFarmCloneprepCustomizationSettingsUpdateSpecWithDefaults() *FarmCloneprepCustomizationSettingsUpdateSpec`

NewFarmCloneprepCustomizationSettingsUpdateSpecWithDefaults instantiates a new FarmCloneprepCustomizationSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostSynchronizationScriptName

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPostSynchronizationScriptName() string`

GetPostSynchronizationScriptName returns the PostSynchronizationScriptName field if non-nil, zero value otherwise.

### GetPostSynchronizationScriptNameOk

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPostSynchronizationScriptNameOk() (*string, bool)`

GetPostSynchronizationScriptNameOk returns a tuple with the PostSynchronizationScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSynchronizationScriptName

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) SetPostSynchronizationScriptName(v string)`

SetPostSynchronizationScriptName sets PostSynchronizationScriptName field to given value.

### HasPostSynchronizationScriptName

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) HasPostSynchronizationScriptName() bool`

HasPostSynchronizationScriptName returns a boolean if a field has been set.

### GetPostSynchronizationScriptParameters

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPostSynchronizationScriptParameters() string`

GetPostSynchronizationScriptParameters returns the PostSynchronizationScriptParameters field if non-nil, zero value otherwise.

### GetPostSynchronizationScriptParametersOk

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPostSynchronizationScriptParametersOk() (*string, bool)`

GetPostSynchronizationScriptParametersOk returns a tuple with the PostSynchronizationScriptParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSynchronizationScriptParameters

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) SetPostSynchronizationScriptParameters(v string)`

SetPostSynchronizationScriptParameters sets PostSynchronizationScriptParameters field to given value.

### HasPostSynchronizationScriptParameters

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) HasPostSynchronizationScriptParameters() bool`

HasPostSynchronizationScriptParameters returns a boolean if a field has been set.

### GetPowerOffScriptName

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPowerOffScriptName() string`

GetPowerOffScriptName returns the PowerOffScriptName field if non-nil, zero value otherwise.

### GetPowerOffScriptNameOk

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPowerOffScriptNameOk() (*string, bool)`

GetPowerOffScriptNameOk returns a tuple with the PowerOffScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffScriptName

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) SetPowerOffScriptName(v string)`

SetPowerOffScriptName sets PowerOffScriptName field to given value.

### HasPowerOffScriptName

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) HasPowerOffScriptName() bool`

HasPowerOffScriptName returns a boolean if a field has been set.

### GetPowerOffScriptParameters

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPowerOffScriptParameters() string`

GetPowerOffScriptParameters returns the PowerOffScriptParameters field if non-nil, zero value otherwise.

### GetPowerOffScriptParametersOk

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPowerOffScriptParametersOk() (*string, bool)`

GetPowerOffScriptParametersOk returns a tuple with the PowerOffScriptParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffScriptParameters

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) SetPowerOffScriptParameters(v string)`

SetPowerOffScriptParameters sets PowerOffScriptParameters field to given value.

### HasPowerOffScriptParameters

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) HasPowerOffScriptParameters() bool`

HasPowerOffScriptParameters returns a boolean if a field has been set.

### GetPrimingComputerAccount

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPrimingComputerAccount() string`

GetPrimingComputerAccount returns the PrimingComputerAccount field if non-nil, zero value otherwise.

### GetPrimingComputerAccountOk

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) GetPrimingComputerAccountOk() (*string, bool)`

GetPrimingComputerAccountOk returns a tuple with the PrimingComputerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimingComputerAccount

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) SetPrimingComputerAccount(v string)`

SetPrimingComputerAccount sets PrimingComputerAccount field to given value.

### HasPrimingComputerAccount

`func (o *FarmCloneprepCustomizationSettingsUpdateSpec) HasPrimingComputerAccount() bool`

HasPrimingComputerAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


