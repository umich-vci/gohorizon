# FarmCustomizationSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdContainerRdn** | Pointer to **string** | Instant Clone Engine Active Directory container for clone prep. Default value is CN&#x3D;Computers | [optional] 
**CloneprepCustomizationSettings** | Pointer to [**FarmCloneprepCustomizationSettingsCreateSpec**](FarmCloneprepCustomizationSettingsCreateSpec.md) |  | [optional] 
**InstantCloneDomainAccountId** | **string** | Instant clone domain account. This is the administrator which will add the machines to its domain upon creation. | 
**ReusePreExistingAccounts** | Pointer to **bool** | Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. Default value is false. | [optional] 

## Methods

### NewFarmCustomizationSettingsCreateSpec

`func NewFarmCustomizationSettingsCreateSpec(instantCloneDomainAccountId string, ) *FarmCustomizationSettingsCreateSpec`

NewFarmCustomizationSettingsCreateSpec instantiates a new FarmCustomizationSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmCustomizationSettingsCreateSpecWithDefaults

`func NewFarmCustomizationSettingsCreateSpecWithDefaults() *FarmCustomizationSettingsCreateSpec`

NewFarmCustomizationSettingsCreateSpecWithDefaults instantiates a new FarmCustomizationSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdContainerRdn

`func (o *FarmCustomizationSettingsCreateSpec) GetAdContainerRdn() string`

GetAdContainerRdn returns the AdContainerRdn field if non-nil, zero value otherwise.

### GetAdContainerRdnOk

`func (o *FarmCustomizationSettingsCreateSpec) GetAdContainerRdnOk() (*string, bool)`

GetAdContainerRdnOk returns a tuple with the AdContainerRdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdContainerRdn

`func (o *FarmCustomizationSettingsCreateSpec) SetAdContainerRdn(v string)`

SetAdContainerRdn sets AdContainerRdn field to given value.

### HasAdContainerRdn

`func (o *FarmCustomizationSettingsCreateSpec) HasAdContainerRdn() bool`

HasAdContainerRdn returns a boolean if a field has been set.

### GetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsCreateSpec) GetCloneprepCustomizationSettings() FarmCloneprepCustomizationSettingsCreateSpec`

GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field if non-nil, zero value otherwise.

### GetCloneprepCustomizationSettingsOk

`func (o *FarmCustomizationSettingsCreateSpec) GetCloneprepCustomizationSettingsOk() (*FarmCloneprepCustomizationSettingsCreateSpec, bool)`

GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsCreateSpec) SetCloneprepCustomizationSettings(v FarmCloneprepCustomizationSettingsCreateSpec)`

SetCloneprepCustomizationSettings sets CloneprepCustomizationSettings field to given value.

### HasCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsCreateSpec) HasCloneprepCustomizationSettings() bool`

HasCloneprepCustomizationSettings returns a boolean if a field has been set.

### GetInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsCreateSpec) GetInstantCloneDomainAccountId() string`

GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field if non-nil, zero value otherwise.

### GetInstantCloneDomainAccountIdOk

`func (o *FarmCustomizationSettingsCreateSpec) GetInstantCloneDomainAccountIdOk() (*string, bool)`

GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsCreateSpec) SetInstantCloneDomainAccountId(v string)`

SetInstantCloneDomainAccountId sets InstantCloneDomainAccountId field to given value.


### GetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsCreateSpec) GetReusePreExistingAccounts() bool`

GetReusePreExistingAccounts returns the ReusePreExistingAccounts field if non-nil, zero value otherwise.

### GetReusePreExistingAccountsOk

`func (o *FarmCustomizationSettingsCreateSpec) GetReusePreExistingAccountsOk() (*bool, bool)`

GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsCreateSpec) SetReusePreExistingAccounts(v bool)`

SetReusePreExistingAccounts sets ReusePreExistingAccounts field to given value.

### HasReusePreExistingAccounts

`func (o *FarmCustomizationSettingsCreateSpec) HasReusePreExistingAccounts() bool`

HasReusePreExistingAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


