# FarmCustomizationSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdContainerRdn** | Pointer to **string** | Instant Clone Engine Active Directory container for clone prep. | [optional] 
**CloneprepCustomizationSettings** | Pointer to [**FarmCloneprepCustomizationSettingsInfo**](FarmCloneprepCustomizationSettingsInfo.md) |  | [optional] 
**InstantCloneDomainAccountId** | Pointer to **string** | Instant clone domain account. This is the administrator which will add the machines to its domain upon creation. | [optional] 
**ReusePreExistingAccounts** | Pointer to **bool** | Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. | [optional] 

## Methods

### NewFarmCustomizationSettingsInfo

`func NewFarmCustomizationSettingsInfo() *FarmCustomizationSettingsInfo`

NewFarmCustomizationSettingsInfo instantiates a new FarmCustomizationSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmCustomizationSettingsInfoWithDefaults

`func NewFarmCustomizationSettingsInfoWithDefaults() *FarmCustomizationSettingsInfo`

NewFarmCustomizationSettingsInfoWithDefaults instantiates a new FarmCustomizationSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdContainerRdn

`func (o *FarmCustomizationSettingsInfo) GetAdContainerRdn() string`

GetAdContainerRdn returns the AdContainerRdn field if non-nil, zero value otherwise.

### GetAdContainerRdnOk

`func (o *FarmCustomizationSettingsInfo) GetAdContainerRdnOk() (*string, bool)`

GetAdContainerRdnOk returns a tuple with the AdContainerRdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdContainerRdn

`func (o *FarmCustomizationSettingsInfo) SetAdContainerRdn(v string)`

SetAdContainerRdn sets AdContainerRdn field to given value.

### HasAdContainerRdn

`func (o *FarmCustomizationSettingsInfo) HasAdContainerRdn() bool`

HasAdContainerRdn returns a boolean if a field has been set.

### GetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsInfo) GetCloneprepCustomizationSettings() FarmCloneprepCustomizationSettingsInfo`

GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field if non-nil, zero value otherwise.

### GetCloneprepCustomizationSettingsOk

`func (o *FarmCustomizationSettingsInfo) GetCloneprepCustomizationSettingsOk() (*FarmCloneprepCustomizationSettingsInfo, bool)`

GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsInfo) SetCloneprepCustomizationSettings(v FarmCloneprepCustomizationSettingsInfo)`

SetCloneprepCustomizationSettings sets CloneprepCustomizationSettings field to given value.

### HasCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsInfo) HasCloneprepCustomizationSettings() bool`

HasCloneprepCustomizationSettings returns a boolean if a field has been set.

### GetInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsInfo) GetInstantCloneDomainAccountId() string`

GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field if non-nil, zero value otherwise.

### GetInstantCloneDomainAccountIdOk

`func (o *FarmCustomizationSettingsInfo) GetInstantCloneDomainAccountIdOk() (*string, bool)`

GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsInfo) SetInstantCloneDomainAccountId(v string)`

SetInstantCloneDomainAccountId sets InstantCloneDomainAccountId field to given value.

### HasInstantCloneDomainAccountId

`func (o *FarmCustomizationSettingsInfo) HasInstantCloneDomainAccountId() bool`

HasInstantCloneDomainAccountId returns a boolean if a field has been set.

### GetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsInfo) GetReusePreExistingAccounts() bool`

GetReusePreExistingAccounts returns the ReusePreExistingAccounts field if non-nil, zero value otherwise.

### GetReusePreExistingAccountsOk

`func (o *FarmCustomizationSettingsInfo) GetReusePreExistingAccountsOk() (*bool, bool)`

GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsInfo) SetReusePreExistingAccounts(v bool)`

SetReusePreExistingAccounts sets ReusePreExistingAccounts field to given value.

### HasReusePreExistingAccounts

`func (o *FarmCustomizationSettingsInfo) HasReusePreExistingAccounts() bool`

HasReusePreExistingAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


