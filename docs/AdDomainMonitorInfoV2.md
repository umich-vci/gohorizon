# ADDomainMonitorInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | Pointer to [**[]ADDomainMonitorConnectionServerV2**](ADDomainMonitorConnectionServerV2.md) | Information about the AD Domain connections from each of the connection servers. | [optional] 
**DnsName** | Pointer to **string** | The DNS name for the domain. | [optional] 
**NetbiosName** | Pointer to **string** | The NetBIOS name for the domain. | [optional] 
**Nt4Domain** | Pointer to **bool** | If this is an NT4 domain or not. | [optional] 

## Methods

### NewADDomainMonitorInfoV2

`func NewADDomainMonitorInfoV2() *ADDomainMonitorInfoV2`

NewADDomainMonitorInfoV2 instantiates a new ADDomainMonitorInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainMonitorInfoV2WithDefaults

`func NewADDomainMonitorInfoV2WithDefaults() *ADDomainMonitorInfoV2`

NewADDomainMonitorInfoV2WithDefaults instantiates a new ADDomainMonitorInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *ADDomainMonitorInfoV2) GetConnectionServers() []ADDomainMonitorConnectionServerV2`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *ADDomainMonitorInfoV2) GetConnectionServersOk() (*[]ADDomainMonitorConnectionServerV2, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *ADDomainMonitorInfoV2) SetConnectionServers(v []ADDomainMonitorConnectionServerV2)`

SetConnectionServers sets ConnectionServers field to given value.

### HasConnectionServers

`func (o *ADDomainMonitorInfoV2) HasConnectionServers() bool`

HasConnectionServers returns a boolean if a field has been set.

### GetDnsName

`func (o *ADDomainMonitorInfoV2) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ADDomainMonitorInfoV2) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ADDomainMonitorInfoV2) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *ADDomainMonitorInfoV2) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetNetbiosName

`func (o *ADDomainMonitorInfoV2) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *ADDomainMonitorInfoV2) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *ADDomainMonitorInfoV2) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *ADDomainMonitorInfoV2) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### GetNt4Domain

`func (o *ADDomainMonitorInfoV2) GetNt4Domain() bool`

GetNt4Domain returns the Nt4Domain field if non-nil, zero value otherwise.

### GetNt4DomainOk

`func (o *ADDomainMonitorInfoV2) GetNt4DomainOk() (*bool, bool)`

GetNt4DomainOk returns a tuple with the Nt4Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNt4Domain

`func (o *ADDomainMonitorInfoV2) SetNt4Domain(v bool)`

SetNt4Domain sets Nt4Domain field to given value.

### HasNt4Domain

`func (o *ADDomainMonitorInfoV2) HasNt4Domain() bool`

HasNt4Domain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


