# ADDomainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | Pointer to **string** | DNS name of the AD Domain. | [optional] 
**Id** | Pointer to **string** | Unique SID representing AD Domain. | [optional] 
**NetbiosName** | Pointer to **string** | NetBIOS name of the AD Domain. | [optional] 

## Methods

### NewADDomainInfo

`func NewADDomainInfo() *ADDomainInfo`

NewADDomainInfo instantiates a new ADDomainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainInfoWithDefaults

`func NewADDomainInfoWithDefaults() *ADDomainInfo`

NewADDomainInfoWithDefaults instantiates a new ADDomainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *ADDomainInfo) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ADDomainInfo) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ADDomainInfo) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *ADDomainInfo) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetId

`func (o *ADDomainInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ADDomainInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ADDomainInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ADDomainInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetbiosName

`func (o *ADDomainInfo) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *ADDomainInfo) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *ADDomainInfo) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *ADDomainInfo) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


