# RDSServerRegisterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | An optional string to describe how and why this RDS Server was registered. | [optional] 
**DnsName** | **string** | The DNS name for the RDS Server. | 
**FarmId** | Pointer to **string** | The farm to which the RDS Server will be added. | [optional] 
**OperatingSystem** | **string** | The Operating System of the RDS Server. * UNKNOWN: Unknown * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_SERVER_OTHER: Linux Server (other) | 

## Methods

### NewRDSServerRegisterSpec

`func NewRDSServerRegisterSpec(dnsName string, operatingSystem string, ) *RDSServerRegisterSpec`

NewRDSServerRegisterSpec instantiates a new RDSServerRegisterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSServerRegisterSpecWithDefaults

`func NewRDSServerRegisterSpecWithDefaults() *RDSServerRegisterSpec`

NewRDSServerRegisterSpecWithDefaults instantiates a new RDSServerRegisterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RDSServerRegisterSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RDSServerRegisterSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RDSServerRegisterSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RDSServerRegisterSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsName

`func (o *RDSServerRegisterSpec) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RDSServerRegisterSpec) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RDSServerRegisterSpec) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetFarmId

`func (o *RDSServerRegisterSpec) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *RDSServerRegisterSpec) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *RDSServerRegisterSpec) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *RDSServerRegisterSpec) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *RDSServerRegisterSpec) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *RDSServerRegisterSpec) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *RDSServerRegisterSpec) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


