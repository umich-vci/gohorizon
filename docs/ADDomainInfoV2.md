# ADDomainInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainAdvancedSettings** | Pointer to [**ADDomainAdvancedSettings**](ADDomainAdvancedSettings.md) |  | [optional] 
**DnsName** | **string** | DNS name of the AD Domain. | 
**DomainType** | Pointer to **string** | AD Domain Type. * CONNECTION_SERVER_DOMAIN: The domain having trust with connection server domain. * NO_TRUST_DOMAIN: The domain not having any trust with connection server domain. | [optional] 
**Id** | **string** | Unique SID representing AD Domain. | 
**NetbiosName** | **string** | NetBIOS name of the AD Domain. | 
**PrimaryAccount** | Pointer to [**ServiceAccountCredentials**](ServiceAccountCredentials.md) |  | [optional] 

## Methods

### NewADDomainInfoV2

`func NewADDomainInfoV2(dnsName string, id string, netbiosName string, ) *ADDomainInfoV2`

NewADDomainInfoV2 instantiates a new ADDomainInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainInfoV2WithDefaults

`func NewADDomainInfoV2WithDefaults() *ADDomainInfoV2`

NewADDomainInfoV2WithDefaults instantiates a new ADDomainInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainAdvancedSettings

`func (o *ADDomainInfoV2) GetAdDomainAdvancedSettings() ADDomainAdvancedSettings`

GetAdDomainAdvancedSettings returns the AdDomainAdvancedSettings field if non-nil, zero value otherwise.

### GetAdDomainAdvancedSettingsOk

`func (o *ADDomainInfoV2) GetAdDomainAdvancedSettingsOk() (*ADDomainAdvancedSettings, bool)`

GetAdDomainAdvancedSettingsOk returns a tuple with the AdDomainAdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainAdvancedSettings

`func (o *ADDomainInfoV2) SetAdDomainAdvancedSettings(v ADDomainAdvancedSettings)`

SetAdDomainAdvancedSettings sets AdDomainAdvancedSettings field to given value.

### HasAdDomainAdvancedSettings

`func (o *ADDomainInfoV2) HasAdDomainAdvancedSettings() bool`

HasAdDomainAdvancedSettings returns a boolean if a field has been set.

### GetDnsName

`func (o *ADDomainInfoV2) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ADDomainInfoV2) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ADDomainInfoV2) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetDomainType

`func (o *ADDomainInfoV2) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *ADDomainInfoV2) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *ADDomainInfoV2) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.

### HasDomainType

`func (o *ADDomainInfoV2) HasDomainType() bool`

HasDomainType returns a boolean if a field has been set.

### GetId

`func (o *ADDomainInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ADDomainInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ADDomainInfoV2) SetId(v string)`

SetId sets Id field to given value.


### GetNetbiosName

`func (o *ADDomainInfoV2) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *ADDomainInfoV2) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *ADDomainInfoV2) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.


### GetPrimaryAccount

`func (o *ADDomainInfoV2) GetPrimaryAccount() ServiceAccountCredentials`

GetPrimaryAccount returns the PrimaryAccount field if non-nil, zero value otherwise.

### GetPrimaryAccountOk

`func (o *ADDomainInfoV2) GetPrimaryAccountOk() (*ServiceAccountCredentials, bool)`

GetPrimaryAccountOk returns a tuple with the PrimaryAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccount

`func (o *ADDomainInfoV2) SetPrimaryAccount(v ServiceAccountCredentials)`

SetPrimaryAccount sets PrimaryAccount field to given value.

### HasPrimaryAccount

`func (o *ADDomainInfoV2) HasPrimaryAccount() bool`

HasPrimaryAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


