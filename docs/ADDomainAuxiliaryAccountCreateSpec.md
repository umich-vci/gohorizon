# ADDomainAuxiliaryAccountCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuxiliaryAccounts** | [**[]ServiceAccountCredentials**](ServiceAccountCredentials.md) | Auxiliary service account credentials. | 

## Methods

### NewADDomainAuxiliaryAccountCreateSpec

`func NewADDomainAuxiliaryAccountCreateSpec(auxiliaryAccounts []ServiceAccountCredentials, ) *ADDomainAuxiliaryAccountCreateSpec`

NewADDomainAuxiliaryAccountCreateSpec instantiates a new ADDomainAuxiliaryAccountCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainAuxiliaryAccountCreateSpecWithDefaults

`func NewADDomainAuxiliaryAccountCreateSpecWithDefaults() *ADDomainAuxiliaryAccountCreateSpec`

NewADDomainAuxiliaryAccountCreateSpecWithDefaults instantiates a new ADDomainAuxiliaryAccountCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxiliaryAccounts

`func (o *ADDomainAuxiliaryAccountCreateSpec) GetAuxiliaryAccounts() []ServiceAccountCredentials`

GetAuxiliaryAccounts returns the AuxiliaryAccounts field if non-nil, zero value otherwise.

### GetAuxiliaryAccountsOk

`func (o *ADDomainAuxiliaryAccountCreateSpec) GetAuxiliaryAccountsOk() (*[]ServiceAccountCredentials, bool)`

GetAuxiliaryAccountsOk returns a tuple with the AuxiliaryAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryAccounts

`func (o *ADDomainAuxiliaryAccountCreateSpec) SetAuxiliaryAccounts(v []ServiceAccountCredentials)`

SetAuxiliaryAccounts sets AuxiliaryAccounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


