# ADDomainInfoV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainAdvancedSettings** | Pointer to [**ADDomainAdvancedSettings**](ADDomainAdvancedSettings.md) |  | [optional] 
**AuxiliaryAccounts** | Pointer to [**[]ServiceAccountCredentialsInfo**](ServiceAccountCredentialsInfo.md) | Auxiliary service accounts information of untrusted domain. | [optional] 
**DnsName** | Pointer to **string** | DNS name of the AD Domain. | [optional] 
**DomainType** | Pointer to **string** | AD Domain Type. * CONNECTION_SERVER_DOMAIN: The domain having trust with connection server domain. * NO_TRUST_DOMAIN: The domain not having any trust with connection server domain. | [optional] 
**Id** | Pointer to **string** | Unique SID representing AD Domain. | [optional] 
**NetbiosName** | Pointer to **string** | NetBIOS name of the AD Domain. | [optional] 
**PrimaryAccount** | Pointer to [**ServiceAccountCredentials**](ServiceAccountCredentials.md) |  | [optional] 

## Methods

### NewADDomainInfoV3

`func NewADDomainInfoV3() *ADDomainInfoV3`

NewADDomainInfoV3 instantiates a new ADDomainInfoV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainInfoV3WithDefaults

`func NewADDomainInfoV3WithDefaults() *ADDomainInfoV3`

NewADDomainInfoV3WithDefaults instantiates a new ADDomainInfoV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainAdvancedSettings

`func (o *ADDomainInfoV3) GetAdDomainAdvancedSettings() ADDomainAdvancedSettings`

GetAdDomainAdvancedSettings returns the AdDomainAdvancedSettings field if non-nil, zero value otherwise.

### GetAdDomainAdvancedSettingsOk

`func (o *ADDomainInfoV3) GetAdDomainAdvancedSettingsOk() (*ADDomainAdvancedSettings, bool)`

GetAdDomainAdvancedSettingsOk returns a tuple with the AdDomainAdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainAdvancedSettings

`func (o *ADDomainInfoV3) SetAdDomainAdvancedSettings(v ADDomainAdvancedSettings)`

SetAdDomainAdvancedSettings sets AdDomainAdvancedSettings field to given value.

### HasAdDomainAdvancedSettings

`func (o *ADDomainInfoV3) HasAdDomainAdvancedSettings() bool`

HasAdDomainAdvancedSettings returns a boolean if a field has been set.

### GetAuxiliaryAccounts

`func (o *ADDomainInfoV3) GetAuxiliaryAccounts() []ServiceAccountCredentialsInfo`

GetAuxiliaryAccounts returns the AuxiliaryAccounts field if non-nil, zero value otherwise.

### GetAuxiliaryAccountsOk

`func (o *ADDomainInfoV3) GetAuxiliaryAccountsOk() (*[]ServiceAccountCredentialsInfo, bool)`

GetAuxiliaryAccountsOk returns a tuple with the AuxiliaryAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryAccounts

`func (o *ADDomainInfoV3) SetAuxiliaryAccounts(v []ServiceAccountCredentialsInfo)`

SetAuxiliaryAccounts sets AuxiliaryAccounts field to given value.

### HasAuxiliaryAccounts

`func (o *ADDomainInfoV3) HasAuxiliaryAccounts() bool`

HasAuxiliaryAccounts returns a boolean if a field has been set.

### GetDnsName

`func (o *ADDomainInfoV3) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ADDomainInfoV3) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ADDomainInfoV3) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *ADDomainInfoV3) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDomainType

`func (o *ADDomainInfoV3) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *ADDomainInfoV3) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *ADDomainInfoV3) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.

### HasDomainType

`func (o *ADDomainInfoV3) HasDomainType() bool`

HasDomainType returns a boolean if a field has been set.

### GetId

`func (o *ADDomainInfoV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ADDomainInfoV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ADDomainInfoV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ADDomainInfoV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetbiosName

`func (o *ADDomainInfoV3) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *ADDomainInfoV3) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *ADDomainInfoV3) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *ADDomainInfoV3) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### GetPrimaryAccount

`func (o *ADDomainInfoV3) GetPrimaryAccount() ServiceAccountCredentials`

GetPrimaryAccount returns the PrimaryAccount field if non-nil, zero value otherwise.

### GetPrimaryAccountOk

`func (o *ADDomainInfoV3) GetPrimaryAccountOk() (*ServiceAccountCredentials, bool)`

GetPrimaryAccountOk returns a tuple with the PrimaryAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccount

`func (o *ADDomainInfoV3) SetPrimaryAccount(v ServiceAccountCredentials)`

SetPrimaryAccount sets PrimaryAccount field to given value.

### HasPrimaryAccount

`func (o *ADDomainInfoV3) HasPrimaryAccount() bool`

HasPrimaryAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


