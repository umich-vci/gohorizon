# NetworkLabelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailablePorts** | Pointer to **int32** | Available ports in network label. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the network label. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this Network Label from being used in desktoppool/farm configuration. | [optional] 
**LabelType** | Pointer to **string** | Network label type. * EARLY_BINDING: A free Distributed Virtual Port will be selected and assigned to a Virtual Machine when the Virtual Machine is reconfigured to connect to the portgroup. Instant clones desktop pools/farms only support port group type of early binding. * EPHEMERAL: A Distributed Virtual Port will be created and assigned to a Virtual Machine when the Virtual Machine is powered on, and will be deleted when the Virtual Machine is powered off. An ephemeral portgroup has no limit on the number of ports that can be a part of this portgroup. In cases where the vCenter Server is unavailable the host can create conflict ports in this portgroup to be used by a Virtual Machine at power on. * LATE_BINDING: Deprecated as of vSphere API 5.0 A free DistributedVirtualPort will be selected and assigned to a Virtual Machine when the Virtual Machine is powered on. | [optional] 
**MaxPorts** | Pointer to **int32** | The total number of ports present. | [optional] 
**Name** | Pointer to **string** | Network label name. | [optional] 
**SwitchType** | Pointer to **string** | Network label switch type. * STANDARD_SWITCH: Standard Switch. * DISTRIBUTED_VIRTUAL_SWITCH: Distributed Virtual Switch. * NSX_NETWORK_SWITCH: NSX network Switch. | [optional] 

## Methods

### NewNetworkLabelInfo

`func NewNetworkLabelInfo() *NetworkLabelInfo`

NewNetworkLabelInfo instantiates a new NetworkLabelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLabelInfoWithDefaults

`func NewNetworkLabelInfoWithDefaults() *NetworkLabelInfo`

NewNetworkLabelInfoWithDefaults instantiates a new NetworkLabelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailablePorts

`func (o *NetworkLabelInfo) GetAvailablePorts() int32`

GetAvailablePorts returns the AvailablePorts field if non-nil, zero value otherwise.

### GetAvailablePortsOk

`func (o *NetworkLabelInfo) GetAvailablePortsOk() (*int32, bool)`

GetAvailablePortsOk returns a tuple with the AvailablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePorts

`func (o *NetworkLabelInfo) SetAvailablePorts(v int32)`

SetAvailablePorts sets AvailablePorts field to given value.

### HasAvailablePorts

`func (o *NetworkLabelInfo) HasAvailablePorts() bool`

HasAvailablePorts returns a boolean if a field has been set.

### GetId

`func (o *NetworkLabelInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkLabelInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkLabelInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkLabelInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *NetworkLabelInfo) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *NetworkLabelInfo) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *NetworkLabelInfo) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *NetworkLabelInfo) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetLabelType

`func (o *NetworkLabelInfo) GetLabelType() string`

GetLabelType returns the LabelType field if non-nil, zero value otherwise.

### GetLabelTypeOk

`func (o *NetworkLabelInfo) GetLabelTypeOk() (*string, bool)`

GetLabelTypeOk returns a tuple with the LabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelType

`func (o *NetworkLabelInfo) SetLabelType(v string)`

SetLabelType sets LabelType field to given value.

### HasLabelType

`func (o *NetworkLabelInfo) HasLabelType() bool`

HasLabelType returns a boolean if a field has been set.

### GetMaxPorts

`func (o *NetworkLabelInfo) GetMaxPorts() int32`

GetMaxPorts returns the MaxPorts field if non-nil, zero value otherwise.

### GetMaxPortsOk

`func (o *NetworkLabelInfo) GetMaxPortsOk() (*int32, bool)`

GetMaxPortsOk returns a tuple with the MaxPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPorts

`func (o *NetworkLabelInfo) SetMaxPorts(v int32)`

SetMaxPorts sets MaxPorts field to given value.

### HasMaxPorts

`func (o *NetworkLabelInfo) HasMaxPorts() bool`

HasMaxPorts returns a boolean if a field has been set.

### GetName

`func (o *NetworkLabelInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkLabelInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkLabelInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkLabelInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSwitchType

`func (o *NetworkLabelInfo) GetSwitchType() string`

GetSwitchType returns the SwitchType field if non-nil, zero value otherwise.

### GetSwitchTypeOk

`func (o *NetworkLabelInfo) GetSwitchTypeOk() (*string, bool)`

GetSwitchTypeOk returns a tuple with the SwitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchType

`func (o *NetworkLabelInfo) SetSwitchType(v string)`

SetSwitchType sets SwitchType field to given value.

### HasSwitchType

`func (o *NetworkLabelInfo) HasSwitchType() bool`

HasSwitchType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


