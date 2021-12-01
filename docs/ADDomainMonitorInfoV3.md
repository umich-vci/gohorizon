# ADDomainMonitorInfoV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | Pointer to [**[]ADDomainMonitorConnectionServerV2**](ADDomainMonitorConnectionServerV2.md) | Information about the AD Domain connections from each of the connection servers. | [optional] 
**DnsName** | Pointer to **string** | The DNS name for the domain. | [optional] 
**DomainType** | Pointer to **string** | AD Domain Type. * CONNECTION_SERVER_DOMAIN: The domain having trust with connection server domain. * NO_TRUST_DOMAIN: The domain not having any trust with connection server domain. | [optional] 
**NetbiosName** | Pointer to **string** | The NetBIOS name for the domain. | [optional] 
**Nt4Domain** | Pointer to **bool** | If this is an NT4 domain or not. | [optional] 
**ServiceAccounts** | Pointer to [**[]ServiceAccount**](ServiceAccount.md) | Service accounts for the domain. | [optional] 

## Methods

### NewADDomainMonitorInfoV3

`func NewADDomainMonitorInfoV3() *ADDomainMonitorInfoV3`

NewADDomainMonitorInfoV3 instantiates a new ADDomainMonitorInfoV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainMonitorInfoV3WithDefaults

`func NewADDomainMonitorInfoV3WithDefaults() *ADDomainMonitorInfoV3`

NewADDomainMonitorInfoV3WithDefaults instantiates a new ADDomainMonitorInfoV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *ADDomainMonitorInfoV3) GetConnectionServers() []ADDomainMonitorConnectionServerV2`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *ADDomainMonitorInfoV3) GetConnectionServersOk() (*[]ADDomainMonitorConnectionServerV2, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *ADDomainMonitorInfoV3) SetConnectionServers(v []ADDomainMonitorConnectionServerV2)`

SetConnectionServers sets ConnectionServers field to given value.

### HasConnectionServers

`func (o *ADDomainMonitorInfoV3) HasConnectionServers() bool`

HasConnectionServers returns a boolean if a field has been set.

### GetDnsName

`func (o *ADDomainMonitorInfoV3) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ADDomainMonitorInfoV3) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ADDomainMonitorInfoV3) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *ADDomainMonitorInfoV3) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDomainType

`func (o *ADDomainMonitorInfoV3) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *ADDomainMonitorInfoV3) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *ADDomainMonitorInfoV3) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.

### HasDomainType

`func (o *ADDomainMonitorInfoV3) HasDomainType() bool`

HasDomainType returns a boolean if a field has been set.

### GetNetbiosName

`func (o *ADDomainMonitorInfoV3) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *ADDomainMonitorInfoV3) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *ADDomainMonitorInfoV3) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *ADDomainMonitorInfoV3) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### GetNt4Domain

`func (o *ADDomainMonitorInfoV3) GetNt4Domain() bool`

GetNt4Domain returns the Nt4Domain field if non-nil, zero value otherwise.

### GetNt4DomainOk

`func (o *ADDomainMonitorInfoV3) GetNt4DomainOk() (*bool, bool)`

GetNt4DomainOk returns a tuple with the Nt4Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNt4Domain

`func (o *ADDomainMonitorInfoV3) SetNt4Domain(v bool)`

SetNt4Domain sets Nt4Domain field to given value.

### HasNt4Domain

`func (o *ADDomainMonitorInfoV3) HasNt4Domain() bool`

HasNt4Domain returns a boolean if a field has been set.

### GetServiceAccounts

`func (o *ADDomainMonitorInfoV3) GetServiceAccounts() []ServiceAccount`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ADDomainMonitorInfoV3) GetServiceAccountsOk() (*[]ServiceAccount, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ADDomainMonitorInfoV3) SetServiceAccounts(v []ServiceAccount)`

SetServiceAccounts sets ServiceAccounts field to given value.

### HasServiceAccounts

`func (o *ADDomainMonitorInfoV3) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


