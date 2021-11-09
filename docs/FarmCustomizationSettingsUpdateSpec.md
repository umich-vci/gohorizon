# FarmCustomizationSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdContainerRdn** | **string** | Instant Clone Engine Active Directory container for clone prep. | 
**CloneprepCustomizationSettings** | [**FarmCloneprepCustomizationSettingsUpdateSpec**](FarmCloneprepCustomizationSettingsUpdateSpec.md) |  | 
**ReusePreExistingAccounts** | **bool** | Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. | 

## Methods

### NewFarmCustomizationSettingsUpdateSpec

`func NewFarmCustomizationSettingsUpdateSpec(adContainerRdn string, cloneprepCustomizationSettings FarmCloneprepCustomizationSettingsUpdateSpec, reusePreExistingAccounts bool, ) *FarmCustomizationSettingsUpdateSpec`

NewFarmCustomizationSettingsUpdateSpec instantiates a new FarmCustomizationSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmCustomizationSettingsUpdateSpecWithDefaults

`func NewFarmCustomizationSettingsUpdateSpecWithDefaults() *FarmCustomizationSettingsUpdateSpec`

NewFarmCustomizationSettingsUpdateSpecWithDefaults instantiates a new FarmCustomizationSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdContainerRdn

`func (o *FarmCustomizationSettingsUpdateSpec) GetAdContainerRdn() string`

GetAdContainerRdn returns the AdContainerRdn field if non-nil, zero value otherwise.

### GetAdContainerRdnOk

`func (o *FarmCustomizationSettingsUpdateSpec) GetAdContainerRdnOk() (*string, bool)`

GetAdContainerRdnOk returns a tuple with the AdContainerRdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdContainerRdn

`func (o *FarmCustomizationSettingsUpdateSpec) SetAdContainerRdn(v string)`

SetAdContainerRdn sets AdContainerRdn field to given value.


### GetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsUpdateSpec) GetCloneprepCustomizationSettings() FarmCloneprepCustomizationSettingsUpdateSpec`

GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field if non-nil, zero value otherwise.

### GetCloneprepCustomizationSettingsOk

`func (o *FarmCustomizationSettingsUpdateSpec) GetCloneprepCustomizationSettingsOk() (*FarmCloneprepCustomizationSettingsUpdateSpec, bool)`

GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneprepCustomizationSettings

`func (o *FarmCustomizationSettingsUpdateSpec) SetCloneprepCustomizationSettings(v FarmCloneprepCustomizationSettingsUpdateSpec)`

SetCloneprepCustomizationSettings sets CloneprepCustomizationSettings field to given value.


### GetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsUpdateSpec) GetReusePreExistingAccounts() bool`

GetReusePreExistingAccounts returns the ReusePreExistingAccounts field if non-nil, zero value otherwise.

### GetReusePreExistingAccountsOk

`func (o *FarmCustomizationSettingsUpdateSpec) GetReusePreExistingAccountsOk() (*bool, bool)`

GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusePreExistingAccounts

`func (o *FarmCustomizationSettingsUpdateSpec) SetReusePreExistingAccounts(v bool)`

SetReusePreExistingAccounts sets ReusePreExistingAccounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


