# DesktopPoolCustomizationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdContainerRdn** | Pointer to **string** | Applicable To: Linked/instant clone automated desktop pools.&lt;br&gt;View Composer and Instant Clone Engine Active Directory container for QuickPrep and ClonePrep. | [optional] 
**CloneprepCustomizationSettings** | Pointer to [**DesktopPoolCloneprepCustomizationSettings**](DesktopPoolCloneprepCustomizationSettings.md) |  | [optional] 
**CustomizationType** | Pointer to **string** | Type of customization to use. * NONE: Applicable To: Full clone desktop pools.&lt;br&gt;No customization. * SYS_PREP: Applicable To: Full clone desktop pools.&lt;br&gt;Microsoft Sysprep is a tool to deploy the configured operating system installation from a base image. The machine can then be customized based on an answer script. Sysprep can modify a larger number of configurable parameters than QuickPrep. * CLONE_PREP: Applicable To: Instant clone desktop pools.&lt;br&gt;ClonePrep is a VMware system tool executed by Instant Clone Engine during a instant clone machine deployment. ClonePrep personalizes each machine created from the Master image. | [optional] 
**DoNotPowerOnVmsAfterCreation** | Pointer to **bool** | Whether to power on VMs after creation. This is the settings when customization will be done manually. | [optional] 
**InstantCloneDomainAccountId** | Pointer to **string** | Applicable To: Instant clone automated desktop pools.&lt;br&gt;Instant clone domain account. This is the administrator which will add the machines to its domain upon creation. | [optional] 
**QuickprepCustomizationSettings** | Pointer to [**DesktopPoolQuickprepCustomizationSettings**](DesktopPoolQuickprepCustomizationSettings.md) |  | [optional] 
**ReusePreExistingAccounts** | Pointer to **bool** | Applicable To: Manual and automated desktop pools.&lt;br&gt;Whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. This is applicable only for automated desktop pools. | [optional] 
**SysprepCustomizationSpecId** | Pointer to **string** | Customization specification to use when Sysprep customization is requested. | [optional] 
**ViewComposerDomainAccountId** | Pointer to **string** | Applicable To: Linked clone automated desktop pools.&lt;br&gt;View Composer domain account. This is the administrator which will add the machines to its domain upon creation. This must be set for linked-clone automated desktop pools. | [optional] 

## Methods

### NewDesktopPoolCustomizationSettings

`func NewDesktopPoolCustomizationSettings() *DesktopPoolCustomizationSettings`

NewDesktopPoolCustomizationSettings instantiates a new DesktopPoolCustomizationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolCustomizationSettingsWithDefaults

`func NewDesktopPoolCustomizationSettingsWithDefaults() *DesktopPoolCustomizationSettings`

NewDesktopPoolCustomizationSettingsWithDefaults instantiates a new DesktopPoolCustomizationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdContainerRdn

`func (o *DesktopPoolCustomizationSettings) GetAdContainerRdn() string`

GetAdContainerRdn returns the AdContainerRdn field if non-nil, zero value otherwise.

### GetAdContainerRdnOk

`func (o *DesktopPoolCustomizationSettings) GetAdContainerRdnOk() (*string, bool)`

GetAdContainerRdnOk returns a tuple with the AdContainerRdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdContainerRdn

`func (o *DesktopPoolCustomizationSettings) SetAdContainerRdn(v string)`

SetAdContainerRdn sets AdContainerRdn field to given value.

### HasAdContainerRdn

`func (o *DesktopPoolCustomizationSettings) HasAdContainerRdn() bool`

HasAdContainerRdn returns a boolean if a field has been set.

### GetCloneprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettings) GetCloneprepCustomizationSettings() DesktopPoolCloneprepCustomizationSettings`

GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field if non-nil, zero value otherwise.

### GetCloneprepCustomizationSettingsOk

`func (o *DesktopPoolCustomizationSettings) GetCloneprepCustomizationSettingsOk() (*DesktopPoolCloneprepCustomizationSettings, bool)`

GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettings) SetCloneprepCustomizationSettings(v DesktopPoolCloneprepCustomizationSettings)`

SetCloneprepCustomizationSettings sets CloneprepCustomizationSettings field to given value.

### HasCloneprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettings) HasCloneprepCustomizationSettings() bool`

HasCloneprepCustomizationSettings returns a boolean if a field has been set.

### GetCustomizationType

`func (o *DesktopPoolCustomizationSettings) GetCustomizationType() string`

GetCustomizationType returns the CustomizationType field if non-nil, zero value otherwise.

### GetCustomizationTypeOk

`func (o *DesktopPoolCustomizationSettings) GetCustomizationTypeOk() (*string, bool)`

GetCustomizationTypeOk returns a tuple with the CustomizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationType

`func (o *DesktopPoolCustomizationSettings) SetCustomizationType(v string)`

SetCustomizationType sets CustomizationType field to given value.

### HasCustomizationType

`func (o *DesktopPoolCustomizationSettings) HasCustomizationType() bool`

HasCustomizationType returns a boolean if a field has been set.

### GetDoNotPowerOnVmsAfterCreation

`func (o *DesktopPoolCustomizationSettings) GetDoNotPowerOnVmsAfterCreation() bool`

GetDoNotPowerOnVmsAfterCreation returns the DoNotPowerOnVmsAfterCreation field if non-nil, zero value otherwise.

### GetDoNotPowerOnVmsAfterCreationOk

`func (o *DesktopPoolCustomizationSettings) GetDoNotPowerOnVmsAfterCreationOk() (*bool, bool)`

GetDoNotPowerOnVmsAfterCreationOk returns a tuple with the DoNotPowerOnVmsAfterCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotPowerOnVmsAfterCreation

`func (o *DesktopPoolCustomizationSettings) SetDoNotPowerOnVmsAfterCreation(v bool)`

SetDoNotPowerOnVmsAfterCreation sets DoNotPowerOnVmsAfterCreation field to given value.

### HasDoNotPowerOnVmsAfterCreation

`func (o *DesktopPoolCustomizationSettings) HasDoNotPowerOnVmsAfterCreation() bool`

HasDoNotPowerOnVmsAfterCreation returns a boolean if a field has been set.

### GetInstantCloneDomainAccountId

`func (o *DesktopPoolCustomizationSettings) GetInstantCloneDomainAccountId() string`

GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field if non-nil, zero value otherwise.

### GetInstantCloneDomainAccountIdOk

`func (o *DesktopPoolCustomizationSettings) GetInstantCloneDomainAccountIdOk() (*string, bool)`

GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneDomainAccountId

`func (o *DesktopPoolCustomizationSettings) SetInstantCloneDomainAccountId(v string)`

SetInstantCloneDomainAccountId sets InstantCloneDomainAccountId field to given value.

### HasInstantCloneDomainAccountId

`func (o *DesktopPoolCustomizationSettings) HasInstantCloneDomainAccountId() bool`

HasInstantCloneDomainAccountId returns a boolean if a field has been set.

### GetQuickprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettings) GetQuickprepCustomizationSettings() DesktopPoolQuickprepCustomizationSettings`

GetQuickprepCustomizationSettings returns the QuickprepCustomizationSettings field if non-nil, zero value otherwise.

### GetQuickprepCustomizationSettingsOk

`func (o *DesktopPoolCustomizationSettings) GetQuickprepCustomizationSettingsOk() (*DesktopPoolQuickprepCustomizationSettings, bool)`

GetQuickprepCustomizationSettingsOk returns a tuple with the QuickprepCustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettings) SetQuickprepCustomizationSettings(v DesktopPoolQuickprepCustomizationSettings)`

SetQuickprepCustomizationSettings sets QuickprepCustomizationSettings field to given value.

### HasQuickprepCustomizationSettings

`func (o *DesktopPoolCustomizationSettings) HasQuickprepCustomizationSettings() bool`

HasQuickprepCustomizationSettings returns a boolean if a field has been set.

### GetReusePreExistingAccounts

`func (o *DesktopPoolCustomizationSettings) GetReusePreExistingAccounts() bool`

GetReusePreExistingAccounts returns the ReusePreExistingAccounts field if non-nil, zero value otherwise.

### GetReusePreExistingAccountsOk

`func (o *DesktopPoolCustomizationSettings) GetReusePreExistingAccountsOk() (*bool, bool)`

GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusePreExistingAccounts

`func (o *DesktopPoolCustomizationSettings) SetReusePreExistingAccounts(v bool)`

SetReusePreExistingAccounts sets ReusePreExistingAccounts field to given value.

### HasReusePreExistingAccounts

`func (o *DesktopPoolCustomizationSettings) HasReusePreExistingAccounts() bool`

HasReusePreExistingAccounts returns a boolean if a field has been set.

### GetSysprepCustomizationSpecId

`func (o *DesktopPoolCustomizationSettings) GetSysprepCustomizationSpecId() string`

GetSysprepCustomizationSpecId returns the SysprepCustomizationSpecId field if non-nil, zero value otherwise.

### GetSysprepCustomizationSpecIdOk

`func (o *DesktopPoolCustomizationSettings) GetSysprepCustomizationSpecIdOk() (*string, bool)`

GetSysprepCustomizationSpecIdOk returns a tuple with the SysprepCustomizationSpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprepCustomizationSpecId

`func (o *DesktopPoolCustomizationSettings) SetSysprepCustomizationSpecId(v string)`

SetSysprepCustomizationSpecId sets SysprepCustomizationSpecId field to given value.

### HasSysprepCustomizationSpecId

`func (o *DesktopPoolCustomizationSettings) HasSysprepCustomizationSpecId() bool`

HasSysprepCustomizationSpecId returns a boolean if a field has been set.

### GetViewComposerDomainAccountId

`func (o *DesktopPoolCustomizationSettings) GetViewComposerDomainAccountId() string`

GetViewComposerDomainAccountId returns the ViewComposerDomainAccountId field if non-nil, zero value otherwise.

### GetViewComposerDomainAccountIdOk

`func (o *DesktopPoolCustomizationSettings) GetViewComposerDomainAccountIdOk() (*string, bool)`

GetViewComposerDomainAccountIdOk returns a tuple with the ViewComposerDomainAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewComposerDomainAccountId

`func (o *DesktopPoolCustomizationSettings) SetViewComposerDomainAccountId(v string)`

SetViewComposerDomainAccountId sets ViewComposerDomainAccountId field to given value.

### HasViewComposerDomainAccountId

`func (o *DesktopPoolCustomizationSettings) HasViewComposerDomainAccountId() bool`

HasViewComposerDomainAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


