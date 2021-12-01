# FarmCustomizationSettingsInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdContainerRdn** | Pointer to **string** | Instant Clone Engine Active Directory container for clone prep. | [optional] 
**CloneprepCustomizationSettings** | Pointer to [**FarmCloneprepCustomizationSettingsInfo**](FarmCloneprepCustomizationSettingsInfo.md) |  | [optional] 
**CustomizationType** | Pointer to **string** | Type of customization to use. * SYS_PREP: Applicable To: Instant clone automated Farms.&lt;br&gt;Microsoft Sysprep is a tool to deploy the configured operating system installation from a base image. The machine can then be customized based on an answer script. * CLONE_PREP: Applicable To: Instant clone automated Farms.&lt;br&gt;ClonePrep is a VMware system tool executed by Instant Clone Engine during a instant clone machine deployment. ClonePrep personalizes each machine created from the Master image. | [optional] 
**InstantCloneDomainAccountId** | Pointer to **string** | Instant clone domain account. This is the administrator which will add the machines to its domain upon creation. | [optional] 
**ReusePreExistingAccounts** | Pointer to **bool** | Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. | [optional] 
**SysprepCustomizationSpecId** | Pointer to **string** | Customization specification used when Sysprep customization is requested. This will be set if customization_type is set to SYS_PREP  | [optional] 

## Methods

### NewFarmCustomizationSettingsInfoV2

`func NewFarmCustomizationSettingsInfoV2() *FarmCustomizationSettingsInfoV2`

NewFarmCustomizationSettingsInfoV2 instantiates a new FarmCustomizationSettingsInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmCustomizationSettingsInfoV2WithDefaults

`func NewFarmCustomizationSettingsInfoV2WithDefaults() *FarmCustomizationSettingsInfoV2`

NewFarmCustomizationSettingsInfoV2WithDefaults instantiates a new FarmCustomizationSettingsInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdContainerRdn

`func (o *FarmCustomizationSettingsInfoV2) GetAdContainerRdn() string`

GetAdContainerRdn returns the AdContainerRdn field if non-nil, zero value otherwise.

### GetAdContainerRdnOk

`func (o *FarmCustomizationSettingsInfoV2) GetAdContainerRdnOk() (*string, bool)`

GetAdContainerRdnOk returns a tuple with the AdContainerRdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdContainerRdn

`func (o *FarmCustomizationSettingsInfoV2) SetAdContainerRdn(v string)`

SetAdContainerRdn sets AdContainerRdn field to given value.

### HasAdContainerRdn

`func (o *FarmCustomizationSettingsInfoV2) HasAdContainerRdn() bool`

HasAdContainerRdn returns a boolean if a field has been set.

### GetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsInfoV2) GetCloneprepCustomizationSettings() FarmCloneprepCustomizationSettingsInfo`

GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field if non-nil, zero value otherwise.

### GetCloneprepCustomizationSettingsOk

`func (o *FarmCustomizationSettingsInfoV2) GetCloneprepCustomizationSettingsOk() (*FarmCloneprepCustomizationSettingsInfo, bool)`

GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsInfoV2) SetCloneprepCustomizationSettings(v FarmCloneprepCustomizationSettingsInfo)`

SetCloneprepCustomizationSettings sets CloneprepCustomizationSettings field to given value.

### HasCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsInfoV2) HasCloneprepCustomizationSettings() bool`

HasCloneprepCustomizationSettings returns a boolean if a field has been set.

### GetCustomizationType

`func (o *FarmCustomizationSettingsInfoV2) GetCustomizationType() string`

GetCustomizationType returns the CustomizationType field if non-nil, zero value otherwise.

### GetCustomizationTypeOk

`func (o *FarmCustomizationSettingsInfoV2) GetCustomizationTypeOk() (*string, bool)`

GetCustomizationTypeOk returns a tuple with the CustomizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationType

`func (o *FarmCustomizationSettingsInfoV2) SetCustomizationType(v string)`

SetCustomizationType sets CustomizationType field to given value.

### HasCustomizationType

`func (o *FarmCustomizationSettingsInfoV2) HasCustomizationType() bool`

HasCustomizationType returns a boolean if a field has been set.

### GetInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsInfoV2) GetInstantCloneDomainAccountId() string`

GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field if non-nil, zero value otherwise.

### GetInstantCloneDomainAccountIdOk

`func (o *FarmCustomizationSettingsInfoV2) GetInstantCloneDomainAccountIdOk() (*string, bool)`

GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsInfoV2) SetInstantCloneDomainAccountId(v string)`

SetInstantCloneDomainAccountId sets InstantCloneDomainAccountId field to given value.

### HasInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsInfoV2) HasInstantCloneDomainAccountId() bool`

HasInstantCloneDomainAccountId returns a boolean if a field has been set.

### GetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsInfoV2) GetReusePreExistingAccounts() bool`

GetReusePreExistingAccounts returns the ReusePreExistingAccounts field if non-nil, zero value otherwise.

### GetReusePreExistingAccountsOk

`func (o *FarmCustomizationSettingsInfoV2) GetReusePreExistingAccountsOk() (*bool, bool)`

GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsInfoV2) SetReusePreExistingAccounts(v bool)`

SetReusePreExistingAccounts sets ReusePreExistingAccounts field to given value.

### HasReusePreExistingAccounts

`func (o *FarmCustomizationSettingsInfoV2) HasReusePreExistingAccounts() bool`

HasReusePreExistingAccounts returns a boolean if a field has been set.

### GetSysprepCustomizationSpecId

`func (o *FarmCustomizationSettingsInfoV2) GetSysprepCustomizationSpecId() string`

GetSysprepCustomizationSpecId returns the SysprepCustomizationSpecId field if non-nil, zero value otherwise.

### GetSysprepCustomizationSpecIdOk

`func (o *FarmCustomizationSettingsInfoV2) GetSysprepCustomizationSpecIdOk() (*string, bool)`

GetSysprepCustomizationSpecIdOk returns a tuple with the SysprepCustomizationSpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprepCustomizationSpecId

`func (o *FarmCustomizationSettingsInfoV2) SetSysprepCustomizationSpecId(v string)`

SetSysprepCustomizationSpecId sets SysprepCustomizationSpecId field to given value.

### HasSysprepCustomizationSpecId

`func (o *FarmCustomizationSettingsInfoV2) HasSysprepCustomizationSpecId() bool`

HasSysprepCustomizationSpecId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


