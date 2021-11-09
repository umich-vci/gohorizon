# PhysicalMachineRegisterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | An optional string to describe how and why this physical machine was registered. | [optional] 
**DnsName** | **string** | The DNS name for the physical machine. | 
**OperatingSystem** | **string** | The Operating System of the physical machine. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS) | 

## Methods

### NewPhysicalMachineRegisterSpec

`func NewPhysicalMachineRegisterSpec(dnsName string, operatingSystem string, ) *PhysicalMachineRegisterSpec`

NewPhysicalMachineRegisterSpec instantiates a new PhysicalMachineRegisterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalMachineRegisterSpecWithDefaults

`func NewPhysicalMachineRegisterSpecWithDefaults() *PhysicalMachineRegisterSpec`

NewPhysicalMachineRegisterSpecWithDefaults instantiates a new PhysicalMachineRegisterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PhysicalMachineRegisterSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PhysicalMachineRegisterSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PhysicalMachineRegisterSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PhysicalMachineRegisterSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsName

`func (o *PhysicalMachineRegisterSpec) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *PhysicalMachineRegisterSpec) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *PhysicalMachineRegisterSpec) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetOperatingSystem

`func (o *PhysicalMachineRegisterSpec) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *PhysicalMachineRegisterSpec) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *PhysicalMachineRegisterSpec) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


