# BaseVMInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterId** | Pointer to **string** | Datacenter id for this VM. | [optional] 
**Id** | Pointer to **string** | Unique ID representing a VM. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this BaseVM from having its snapshots used in linked or instant clone desktop or farm creation. | [optional] 
**Name** | Pointer to **string** | VM name. | [optional] 
**NetworkType** | Pointer to **string** | Type of network base VM belongs to. * STANDARD_NETWORK: Standard network. * OPAQUE_NETWORK: Opaque network. * DISTRUBUTED_VIRTUAL_PORT_GROUP: DVS port group. | [optional] 
**OperatingSystem** | Pointer to **string** | Operating system. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS) | [optional] 
**OperatingSystemDisplayName** | Pointer to **string** | Operating system display name from Virtual Center. | [optional] 
**Path** | Pointer to **string** | VM path. | [optional] 
**VcenterId** | Pointer to **string** | Virtual Center id for this VM. | [optional] 

## Methods

### NewBaseVMInfo

`func NewBaseVMInfo() *BaseVMInfo`

NewBaseVMInfo instantiates a new BaseVMInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseVMInfoWithDefaults

`func NewBaseVMInfoWithDefaults() *BaseVMInfo`

NewBaseVMInfoWithDefaults instantiates a new BaseVMInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *BaseVMInfo) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *BaseVMInfo) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *BaseVMInfo) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *BaseVMInfo) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetId

`func (o *BaseVMInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseVMInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseVMInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseVMInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *BaseVMInfo) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *BaseVMInfo) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *BaseVMInfo) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *BaseVMInfo) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetName

`func (o *BaseVMInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseVMInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseVMInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseVMInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkType

`func (o *BaseVMInfo) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *BaseVMInfo) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *BaseVMInfo) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *BaseVMInfo) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *BaseVMInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *BaseVMInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *BaseVMInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *BaseVMInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperatingSystemDisplayName

`func (o *BaseVMInfo) GetOperatingSystemDisplayName() string`

GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field if non-nil, zero value otherwise.

### GetOperatingSystemDisplayNameOk

`func (o *BaseVMInfo) GetOperatingSystemDisplayNameOk() (*string, bool)`

GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemDisplayName

`func (o *BaseVMInfo) SetOperatingSystemDisplayName(v string)`

SetOperatingSystemDisplayName sets OperatingSystemDisplayName field to given value.

### HasOperatingSystemDisplayName

`func (o *BaseVMInfo) HasOperatingSystemDisplayName() bool`

HasOperatingSystemDisplayName returns a boolean if a field has been set.

### GetPath

`func (o *BaseVMInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BaseVMInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BaseVMInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BaseVMInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVcenterId

`func (o *BaseVMInfo) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *BaseVMInfo) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *BaseVMInfo) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *BaseVMInfo) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


