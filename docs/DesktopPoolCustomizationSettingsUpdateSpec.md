# DesktopPoolCustomizationSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdContainerRdn** | Pointer to **string** | This is required for instant clone desktop pools.&lt;br&gt;Instant Clone Engine Active Directory container for ClonePrep. | [optional] 
**CloneprepCustomizationSettings** | Pointer to [**DesktopPoolCloneprepCustomizationSettingsUpdateSpec**](DesktopPoolCloneprepCustomizationSettingsUpdateSpec.md) |  | [optional] 
**CustomizationType** | **string** | Type of customization to use. * NONE: Applicable To: Full clone desktop pools.&lt;br&gt;No customization. * SYS_PREP: Applicable To: Full clone desktop pools.&lt;br&gt;Microsoft Sysprep is a tool to deploy the configured operating system installation from a base image. The machine can then be customized based on an answer script. Sysprep can modify a larger number of configurable parameters than QuickPrep. * CLONE_PREP: Applicable To: Instant clone desktop pools.&lt;br&gt;ClonePrep is a VMware system tool executed by Instant Clone Engine during a instant clone machine deployment. ClonePrep personalizes each machine created from the Master image. | 
**DoNotPowerOnVmsAfterCreation** | Pointer to **bool** | Indicates whether to power on VMs after creation. This is the settings when customization will be done manually. This property is required if customization_type is set NONE. | [optional] 
**InstantCloneDomainAccountId** | Pointer to **string** | This is required for instant clone desktop pools.&lt;br&gt;This is the administrator which will add the machines to its domain upon creation. | [optional] 
**ReusePreExistingAccounts** | Pointer to **bool** | Applicable To: Automated desktop pools.&lt;br&gt;Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. | [optional] 
**SysprepCustomizationSpecId** | Pointer to **string** | This is required when customization_type is set as SYS_PREP.&lt;br&gt;Customization specification to use when Sysprep customization is requested. | [optional] 

## Methods

### NewDesktopPoolCustomizationSettingsUpdateSpec

`func NewDesktopPoolCustomizationSettingsUpdateSpec(customizationType string, ) *DesktopPoolCustomizationSettingsUpdateSpec`

NewDesktopPoolCustomizationSettingsUpdateSpec instantiates a new DesktopPoolCustomizationSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolCustomizationSettingsUpdateSpecWithDefaults

`func NewDesktopPoolCustomizationSettingsUpdateSpecWithDefaults() *DesktopPoolCustomizationSettingsUpdateSpec`

NewDesktopPoolCustomizationSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolCustomizationSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdContainerRdn

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetAdContainerRdn() string`

GetAdContainerRdn returns the AdContainerRdn field if non-nil, zero value otherwise.

### GetAdContainerRdnOk

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetAdContainerRdnOk() (*string, bool)`

GetAdContainerRdnOk returns a tuple with the AdContainerRdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdContainerRdn

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetAdContainerRdn(v string)`

SetAdContainerRdn sets AdContainerRdn field to given value.

### HasAdContainerRdn

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasAdContainerRdn() bool`

HasAdContainerRdn returns a boolean if a field has been set.

### GetCloneprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetCloneprepCustomizationSettings() DesktopPoolCloneprepCustomizationSettingsUpdateSpec`

GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field if non-nil, zero value otherwise.

### GetCloneprepCustomizationSettingsOk

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetCloneprepCustomizationSettingsOk() (*DesktopPoolCloneprepCustomizationSettingsUpdateSpec, bool)`

GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetCloneprepCustomizationSettings(v DesktopPoolCloneprepCustomizationSettingsUpdateSpec)`

SetCloneprepCustomizationSettings sets CloneprepCustomizationSettings field to given value.

### HasCloneprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasCloneprepCustomizationSettings() bool`

HasCloneprepCustomizationSettings returns a boolean if a field has been set.

### GetCustomizationType

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetCustomizationType() string`

GetCustomizationType returns the CustomizationType field if non-nil, zero value otherwise.

### GetCustomizationTypeOk

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetCustomizationTypeOk() (*string, bool)`

GetCustomizationTypeOk returns a tuple with the CustomizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationType

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetCustomizationType(v string)`

SetCustomizationType sets CustomizationType field to given value.


### GetDoNotPowerOnVmsAfterCreation

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetDoNotPowerOnVmsAfterCreation() bool`

GetDoNotPowerOnVmsAfterCreation returns the DoNotPowerOnVmsAfterCreation field if non-nil, zero value otherwise.

### GetDoNotPowerOnVmsAfterCreationOk

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetDoNotPowerOnVmsAfterCreationOk() (*bool, bool)`

GetDoNotPowerOnVmsAfterCreationOk returns a tuple with the DoNotPowerOnVmsAfterCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotPowerOnVmsAfterCreation

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetDoNotPowerOnVmsAfterCreation(v bool)`

SetDoNotPowerOnVmsAfterCreation sets DoNotPowerOnVmsAfterCreation field to given value.

### HasDoNotPowerOnVmsAfterCreation

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasDoNotPowerOnVmsAfterCreation() bool`

HasDoNotPowerOnVmsAfterCreation returns a boolean if a field has been set.

### GetInstantCloneDomainAccountId

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetInstantCloneDomainAccountId() string`

GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field if non-nil, zero value otherwise.

### GetInstantCloneDomainAccountIdOk

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetInstantCloneDomainAccountIdOk() (*string, bool)`

GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneDomainAccountId

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetInstantCloneDomainAccountId(v string)`

SetInstantCloneDomainAccountId sets InstantCloneDomainAccountId field to given value.

### HasInstantCloneDomainAccountId

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasInstantCloneDomainAccountId() bool`

HasInstantCloneDomainAccountId returns a boolean if a field has been set.

### GetReusePreExistingAccounts

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetReusePreExistingAccounts() bool`

GetReusePreExistingAccounts returns the ReusePreExistingAccounts field if non-nil, zero value otherwise.

### GetReusePreExistingAccountsOk

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetReusePreExistingAccountsOk() (*bool, bool)`

GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusePreExistingAccounts

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetReusePreExistingAccounts(v bool)`

SetReusePreExistingAccounts sets ReusePreExistingAccounts field to given value.

### HasReusePreExistingAccounts

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasReusePreExistingAccounts() bool`

HasReusePreExistingAccounts returns a boolean if a field has been set.

### GetSysprepCustomizationSpecId

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetSysprepCustomizationSpecId() string`

GetSysprepCustomizationSpecId returns the SysprepCustomizationSpecId field if non-nil, zero value otherwise.

### GetSysprepCustomizationSpecIdOk

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetSysprepCustomizationSpecIdOk() (*string, bool)`

GetSysprepCustomizationSpecIdOk returns a tuple with the SysprepCustomizationSpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprepCustomizationSpecId

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetSysprepCustomizationSpecId(v string)`

SetSysprepCustomizationSpecId sets SysprepCustomizationSpecId field to given value.

### HasSysprepCustomizationSpecId

`func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasSysprepCustomizationSpecId() bool`

HasSysprepCustomizationSpecId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


