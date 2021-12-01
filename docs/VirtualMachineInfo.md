# VirtualMachineInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardwareVersion** | Pointer to **int32** | Hardware version for this VM. | [optional] 
**Id** | Pointer to **string** | Unique ID representing a VM. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this Virtual Machine from having its snapshots used in linked or instant clone desktop or farm. | [optional] 
**Name** | Pointer to **string** | VM name. | [optional] 
**OperatingSystem** | Pointer to **string** | Operating system. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS) | [optional] 
**OperatingSystemDisplayName** | Pointer to **string** | Operating system display name from Virtual Center. | [optional] 
**Path** | Pointer to **string** | VM path. | [optional] 
**VGputype** | Pointer to **string** | Virtual GPU type. | [optional] 

## Methods

### NewVirtualMachineInfo

`func NewVirtualMachineInfo() *VirtualMachineInfo`

NewVirtualMachineInfo instantiates a new VirtualMachineInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineInfoWithDefaults

`func NewVirtualMachineInfoWithDefaults() *VirtualMachineInfo`

NewVirtualMachineInfoWithDefaults instantiates a new VirtualMachineInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardwareVersion

`func (o *VirtualMachineInfo) GetHardwareVersion() int32`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *VirtualMachineInfo) GetHardwareVersionOk() (*int32, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *VirtualMachineInfo) SetHardwareVersion(v int32)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *VirtualMachineInfo) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetId

`func (o *VirtualMachineInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualMachineInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualMachineInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualMachineInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *VirtualMachineInfo) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *VirtualMachineInfo) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *VirtualMachineInfo) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *VirtualMachineInfo) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetName

`func (o *VirtualMachineInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachineInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachineInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualMachineInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *VirtualMachineInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *VirtualMachineInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *VirtualMachineInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *VirtualMachineInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperatingSystemDisplayName

`func (o *VirtualMachineInfo) GetOperatingSystemDisplayName() string`

GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field if non-nil, zero value otherwise.

### GetOperatingSystemDisplayNameOk

`func (o *VirtualMachineInfo) GetOperatingSystemDisplayNameOk() (*string, bool)`

GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemDisplayName

`func (o *VirtualMachineInfo) SetOperatingSystemDisplayName(v string)`

SetOperatingSystemDisplayName sets OperatingSystemDisplayName field to given value.

### HasOperatingSystemDisplayName

`func (o *VirtualMachineInfo) HasOperatingSystemDisplayName() bool`

HasOperatingSystemDisplayName returns a boolean if a field has been set.

### GetPath

`func (o *VirtualMachineInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VirtualMachineInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VirtualMachineInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VirtualMachineInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVGputype

`func (o *VirtualMachineInfo) GetVGputype() string`

GetVGputype returns the VGputype field if non-nil, zero value otherwise.

### GetVGputypeOk

`func (o *VirtualMachineInfo) GetVGputypeOk() (*string, bool)`

GetVGputypeOk returns a tuple with the VGputype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVGputype

`func (o *VirtualMachineInfo) SetVGputype(v string)`

SetVGputype sets VGputype field to given value.

### HasVGputype

`func (o *VirtualMachineInfo) HasVGputype() bool`

HasVGputype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


