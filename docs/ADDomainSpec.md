# ADDomainSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainAdvancedSettings** | Pointer to [**ADDomainAdvancedSettings**](ADDomainAdvancedSettings.md) |  | [optional] 
**DnsName** | **string** | The DNS name of the domain. | 
**NetbiosName** | **string** | The NetBIOS name of the domain. | 
**PrimaryAccount** | Pointer to [**ServiceAccountCredentials**](ServiceAccountCredentials.md) |  | [optional] 

## Methods

### NewADDomainSpec

`func NewADDomainSpec(dnsName string, netbiosName string, ) *ADDomainSpec`

NewADDomainSpec instantiates a new ADDomainSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainSpecWithDefaults

`func NewADDomainSpecWithDefaults() *ADDomainSpec`

NewADDomainSpecWithDefaults instantiates a new ADDomainSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainAdvancedSettings

`func (o *ADDomainSpec) GetAdDomainAdvancedSettings() ADDomainAdvancedSettings`

GetAdDomainAdvancedSettings returns the AdDomainAdvancedSettings field if non-nil, zero value otherwise.

### GetAdDomainAdvancedSettingsOk

`func (o *ADDomainSpec) GetAdDomainAdvancedSettingsOk() (*ADDomainAdvancedSettings, bool)`

GetAdDomainAdvancedSettingsOk returns a tuple with the AdDomainAdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainAdvancedSettings

`func (o *ADDomainSpec) SetAdDomainAdvancedSettings(v ADDomainAdvancedSettings)`

SetAdDomainAdvancedSettings sets AdDomainAdvancedSettings field to given value.

### HasAdDomainAdvancedSettings

`func (o *ADDomainSpec) HasAdDomainAdvancedSettings() bool`

HasAdDomainAdvancedSettings returns a boolean if a field has been set.

### GetDnsName

`func (o *ADDomainSpec) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ADDomainSpec) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ADDomainSpec) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetNetbiosName

`func (o *ADDomainSpec) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *ADDomainSpec) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *ADDomainSpec) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.


### GetPrimaryAccount

`func (o *ADDomainSpec) GetPrimaryAccount() ServiceAccountCredentials`

GetPrimaryAccount returns the PrimaryAccount field if non-nil, zero value otherwise.

### GetPrimaryAccountOk

`func (o *ADDomainSpec) GetPrimaryAccountOk() (*ServiceAccountCredentials, bool)`

GetPrimaryAccountOk returns a tuple with the PrimaryAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccount

`func (o *ADDomainSpec) SetPrimaryAccount(v ServiceAccountCredentials)`

SetPrimaryAccount sets PrimaryAccount field to given value.

### HasPrimaryAccount

`func (o *ADDomainSpec) HasPrimaryAccount() bool`

HasPrimaryAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


