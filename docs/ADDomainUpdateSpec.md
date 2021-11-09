# ADDomainUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainAdvancedSettings** | Pointer to [**ADDomainAdvancedSettings**](ADDomainAdvancedSettings.md) |  | [optional] 
**PrimaryAccount** | Pointer to [**ServiceAccountCredentials**](ServiceAccountCredentials.md) |  | [optional] 

## Methods

### NewADDomainUpdateSpec

`func NewADDomainUpdateSpec() *ADDomainUpdateSpec`

NewADDomainUpdateSpec instantiates a new ADDomainUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainUpdateSpecWithDefaults

`func NewADDomainUpdateSpecWithDefaults() *ADDomainUpdateSpec`

NewADDomainUpdateSpecWithDefaults instantiates a new ADDomainUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainAdvancedSettings

`func (o *ADDomainUpdateSpec) GetAdDomainAdvancedSettings() ADDomainAdvancedSettings`

GetAdDomainAdvancedSettings returns the AdDomainAdvancedSettings field if non-nil, zero value otherwise.

### GetAdDomainAdvancedSettingsOk

`func (o *ADDomainUpdateSpec) GetAdDomainAdvancedSettingsOk() (*ADDomainAdvancedSettings, bool)`

GetAdDomainAdvancedSettingsOk returns a tuple with the AdDomainAdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainAdvancedSettings

`func (o *ADDomainUpdateSpec) SetAdDomainAdvancedSettings(v ADDomainAdvancedSettings)`

SetAdDomainAdvancedSettings sets AdDomainAdvancedSettings field to given value.

### HasAdDomainAdvancedSettings

`func (o *ADDomainUpdateSpec) HasAdDomainAdvancedSettings() bool`

HasAdDomainAdvancedSettings returns a boolean if a field has been set.

### GetPrimaryAccount

`func (o *ADDomainUpdateSpec) GetPrimaryAccount() ServiceAccountCredentials`

GetPrimaryAccount returns the PrimaryAccount field if non-nil, zero value otherwise.

### GetPrimaryAccountOk

`func (o *ADDomainUpdateSpec) GetPrimaryAccountOk() (*ServiceAccountCredentials, bool)`

GetPrimaryAccountOk returns a tuple with the PrimaryAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccount

`func (o *ADDomainUpdateSpec) SetPrimaryAccount(v ServiceAccountCredentials)`

SetPrimaryAccount sets PrimaryAccount field to given value.

### HasPrimaryAccount

`func (o *ADDomainUpdateSpec) HasPrimaryAccount() bool`

HasPrimaryAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


