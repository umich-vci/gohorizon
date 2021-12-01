# ADDomainMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | Pointer to [**[]ADDomainMonitorConnectionServer**](ADDomainMonitorConnectionServer.md) | Information about the AD Domain connections from each of the connection servers. | [optional] 
**DnsName** | Pointer to **string** | The DNS name for the domain. | [optional] 
**NetbiosName** | Pointer to **string** | The NetBIOS name for the domain. | [optional] 
**Nt4Domain** | Pointer to **bool** | If this is an NT4 domain or not. | [optional] 

## Methods

### NewADDomainMonitorInfo

`func NewADDomainMonitorInfo() *ADDomainMonitorInfo`

NewADDomainMonitorInfo instantiates a new ADDomainMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainMonitorInfoWithDefaults

`func NewADDomainMonitorInfoWithDefaults() *ADDomainMonitorInfo`

NewADDomainMonitorInfoWithDefaults instantiates a new ADDomainMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *ADDomainMonitorInfo) GetConnectionServers() []ADDomainMonitorConnectionServer`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *ADDomainMonitorInfo) GetConnectionServersOk() (*[]ADDomainMonitorConnectionServer, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *ADDomainMonitorInfo) SetConnectionServers(v []ADDomainMonitorConnectionServer)`

SetConnectionServers sets ConnectionServers field to given value.

### HasConnectionServers

`func (o *ADDomainMonitorInfo) HasConnectionServers() bool`

HasConnectionServers returns a boolean if a field has been set.

### GetDnsName

`func (o *ADDomainMonitorInfo) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ADDomainMonitorInfo) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ADDomainMonitorInfo) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *ADDomainMonitorInfo) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetNetbiosName

`func (o *ADDomainMonitorInfo) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *ADDomainMonitorInfo) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *ADDomainMonitorInfo) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *ADDomainMonitorInfo) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### GetNt4Domain

`func (o *ADDomainMonitorInfo) GetNt4Domain() bool`

GetNt4Domain returns the Nt4Domain field if non-nil, zero value otherwise.

### GetNt4DomainOk

`func (o *ADDomainMonitorInfo) GetNt4DomainOk() (*bool, bool)`

GetNt4DomainOk returns a tuple with the Nt4Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNt4Domain

`func (o *ADDomainMonitorInfo) SetNt4Domain(v bool)`

SetNt4Domain sets Nt4Domain field to given value.

### HasNt4Domain

`func (o *ADDomainMonitorInfo) HasNt4Domain() bool`

HasNt4Domain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


