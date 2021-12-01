# FarmCustomizationSettingsUpdateSpecV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdContainerRdn** | **string** | Instant Clone Engine Active Directory container for clone prep. | 
**CloneprepCustomizationSettings** | [**FarmCloneprepCustomizationSettingsUpdateSpec**](FarmCloneprepCustomizationSettingsUpdateSpec.md) |  | 
**CustomizationType** | **string** | Type of customization to use. * SYS_PREP: Applicable To: Instant clone automated Farms.&lt;br&gt;Microsoft Sysprep is a tool to deploy the configured operating system installation from a base image. The machine can then be customized based on an answer script. * CLONE_PREP: Applicable To: Instant clone automated Farms.&lt;br&gt;ClonePrep is a VMware system tool executed by Instant Clone Engine during a instant clone machine deployment. ClonePrep personalizes each machine created from the Master image. | 
**InstantCloneDomainAccountId** | **string** | Instant clone domain account. This is the administrator which will add the machines to its domain. | 
**ReusePreExistingAccounts** | **bool** | Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. | 
**SysprepCustomizationSpecId** | Pointer to **string** | Customization specification to use when Sysprep customization is requested. This is required if customization_type is set to SYS_PREP | [optional] 

## Methods

### NewFarmCustomizationSettingsUpdateSpecV2

`func NewFarmCustomizationSettingsUpdateSpecV2(adContainerRdn string, cloneprepCustomizationSettings FarmCloneprepCustomizationSettingsUpdateSpec, customizationType string, instantCloneDomainAccountId string, reusePreExistingAccounts bool, ) *FarmCustomizationSettingsUpdateSpecV2`

NewFarmCustomizationSettingsUpdateSpecV2 instantiates a new FarmCustomizationSettingsUpdateSpecV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmCustomizationSettingsUpdateSpecV2WithDefaults

`func NewFarmCustomizationSettingsUpdateSpecV2WithDefaults() *FarmCustomizationSettingsUpdateSpecV2`

NewFarmCustomizationSettingsUpdateSpecV2WithDefaults instantiates a new FarmCustomizationSettingsUpdateSpecV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdContainerRdn

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetAdContainerRdn() string`

GetAdContainerRdn returns the AdContainerRdn field if non-nil, zero value otherwise.

### GetAdContainerRdnOk

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetAdContainerRdnOk() (*string, bool)`

GetAdContainerRdnOk returns a tuple with the AdContainerRdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdContainerRdn

`func (o *FarmCustomizationSettingsUpdateSpecV2) SetAdContainerRdn(v string)`

SetAdContainerRdn sets AdContainerRdn field to given value.


### GetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetCloneprepCustomizationSettings() FarmCloneprepCustomizationSettingsUpdateSpec`

GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field if non-nil, zero value otherwise.

### GetCloneprepCustomizationSettingsOk

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetCloneprepCustomizationSettingsOk() (*FarmCloneprepCustomizationSettingsUpdateSpec, bool)`

GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsUpdateSpecV2) SetCloneprepCustomizationSettings(v FarmCloneprepCustomizationSettingsUpdateSpec)`

SetCloneprepCustomizationSettings sets CloneprepCustomizationSettings field to given value.


### GetCustomizationType

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetCustomizationType() string`

GetCustomizationType returns the CustomizationType field if non-nil, zero value otherwise.

### GetCustomizationTypeOk

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetCustomizationTypeOk() (*string, bool)`

GetCustomizationTypeOk returns a tuple with the CustomizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationType

`func (o *FarmCustomizationSettingsUpdateSpecV2) SetCustomizationType(v string)`

SetCustomizationType sets CustomizationType field to given value.


### GetInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetInstantCloneDomainAccountId() string`

GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field if non-nil, zero value otherwise.

### GetInstantCloneDomainAccountIdOk

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetInstantCloneDomainAccountIdOk() (*string, bool)`

GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsUpdateSpecV2) SetInstantCloneDomainAccountId(v string)`

SetInstantCloneDomainAccountId sets InstantCloneDomainAccountId field to given value.


### GetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetReusePreExistingAccounts() bool`

GetReusePreExistingAccounts returns the ReusePreExistingAccounts field if non-nil, zero value otherwise.

### GetReusePreExistingAccountsOk

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetReusePreExistingAccountsOk() (*bool, bool)`

GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsUpdateSpecV2) SetReusePreExistingAccounts(v bool)`

SetReusePreExistingAccounts sets ReusePreExistingAccounts field to given value.


### GetSysprepCustomizationSpecId

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetSysprepCustomizationSpecId() string`

GetSysprepCustomizationSpecId returns the SysprepCustomizationSpecId field if non-nil, zero value otherwise.

### GetSysprepCustomizationSpecIdOk

`func (o *FarmCustomizationSettingsUpdateSpecV2) GetSysprepCustomizationSpecIdOk() (*string, bool)`

GetSysprepCustomizationSpecIdOk returns a tuple with the SysprepCustomizationSpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprepCustomizationSpecId

`func (o *FarmCustomizationSettingsUpdateSpecV2) SetSysprepCustomizationSpecId(v string)`

SetSysprepCustomizationSpecId sets SysprepCustomizationSpecId field to given value.

### HasSysprepCustomizationSpecId

`func (o *FarmCustomizationSettingsUpdateSpecV2) HasSysprepCustomizationSpecId() bool`

HasSysprepCustomizationSpecId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


