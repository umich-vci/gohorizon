# DesktopPoolCloneprepCustomizationSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostSynchronizationScriptName** | Pointer to **string** | Post synchronization script. ClonePrep can run a customization script on instant-clone machines after they are created or recovered or a new image is pushed. Provide the path to the script on the parent virtual machine. | [optional] 
**PostSynchronizationScriptParameters** | Pointer to **string** | Post synchronization script parameters. | [optional] 
**PowerOffScriptName** | Pointer to **string** | Power off script. ClonePrep can run a customization script on instant-clone machines before they are powered off. Provide the path to the script on the parent virtual machine. | [optional] 
**PowerOffScriptParameters** | Pointer to **string** | Power off script parameters. | [optional] 
**PrimingComputerAccount** | Pointer to **string** | Instant Clone publishing needs an additional computer account in the same AD domain as the clones. This field accepts the pre-created computer accounts. | [optional] 

## Methods

### NewDesktopPoolCloneprepCustomizationSettingsUpdateSpec

`func NewDesktopPoolCloneprepCustomizationSettingsUpdateSpec() *DesktopPoolCloneprepCustomizationSettingsUpdateSpec`

NewDesktopPoolCloneprepCustomizationSettingsUpdateSpec instantiates a new DesktopPoolCloneprepCustomizationSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolCloneprepCustomizationSettingsUpdateSpecWithDefaults

`func NewDesktopPoolCloneprepCustomizationSettingsUpdateSpecWithDefaults() *DesktopPoolCloneprepCustomizationSettingsUpdateSpec`

NewDesktopPoolCloneprepCustomizationSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolCloneprepCustomizationSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostSynchronizationScriptName

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPostSynchronizationScriptName() string`

GetPostSynchronizationScriptName returns the PostSynchronizationScriptName field if non-nil, zero value otherwise.

### GetPostSynchronizationScriptNameOk

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPostSynchronizationScriptNameOk() (*string, bool)`

GetPostSynchronizationScriptNameOk returns a tuple with the PostSynchronizationScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSynchronizationScriptName

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) SetPostSynchronizationScriptName(v string)`

SetPostSynchronizationScriptName sets PostSynchronizationScriptName field to given value.

### HasPostSynchronizationScriptName

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) HasPostSynchronizationScriptName() bool`

HasPostSynchronizationScriptName returns a boolean if a field has been set.

### GetPostSynchronizationScriptParameters

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPostSynchronizationScriptParameters() string`

GetPostSynchronizationScriptParameters returns the PostSynchronizationScriptParameters field if non-nil, zero value otherwise.

### GetPostSynchronizationScriptParametersOk

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPostSynchronizationScriptParametersOk() (*string, bool)`

GetPostSynchronizationScriptParametersOk returns a tuple with the PostSynchronizationScriptParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSynchronizationScriptParameters

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) SetPostSynchronizationScriptParameters(v string)`

SetPostSynchronizationScriptParameters sets PostSynchronizationScriptParameters field to given value.

### HasPostSynchronizationScriptParameters

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) HasPostSynchronizationScriptParameters() bool`

HasPostSynchronizationScriptParameters returns a boolean if a field has been set.

### GetPowerOffScriptName

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPowerOffScriptName() string`

GetPowerOffScriptName returns the PowerOffScriptName field if non-nil, zero value otherwise.

### GetPowerOffScriptNameOk

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPowerOffScriptNameOk() (*string, bool)`

GetPowerOffScriptNameOk returns a tuple with the PowerOffScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffScriptName

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) SetPowerOffScriptName(v string)`

SetPowerOffScriptName sets PowerOffScriptName field to given value.

### HasPowerOffScriptName

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) HasPowerOffScriptName() bool`

HasPowerOffScriptName returns a boolean if a field has been set.

### GetPowerOffScriptParameters

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPowerOffScriptParameters() string`

GetPowerOffScriptParameters returns the PowerOffScriptParameters field if non-nil, zero value otherwise.

### GetPowerOffScriptParametersOk

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPowerOffScriptParametersOk() (*string, bool)`

GetPowerOffScriptParametersOk returns a tuple with the PowerOffScriptParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffScriptParameters

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) SetPowerOffScriptParameters(v string)`

SetPowerOffScriptParameters sets PowerOffScriptParameters field to given value.

### HasPowerOffScriptParameters

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) HasPowerOffScriptParameters() bool`

HasPowerOffScriptParameters returns a boolean if a field has been set.

### GetPrimingComputerAccount

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPrimingComputerAccount() string`

GetPrimingComputerAccount returns the PrimingComputerAccount field if non-nil, zero value otherwise.

### GetPrimingComputerAccountOk

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) GetPrimingComputerAccountOk() (*string, bool)`

GetPrimingComputerAccountOk returns a tuple with the PrimingComputerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimingComputerAccount

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) SetPrimingComputerAccount(v string)`

SetPrimingComputerAccount sets PrimingComputerAccount field to given value.

### HasPrimingComputerAccount

`func (o *DesktopPoolCloneprepCustomizationSettingsUpdateSpec) HasPrimingComputerAccount() bool`

HasPrimingComputerAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


