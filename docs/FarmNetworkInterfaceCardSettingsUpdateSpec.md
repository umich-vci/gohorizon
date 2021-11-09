# FarmNetworkInterfaceCardSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaceCardId** | **string** | ID of the network interface card for these settings. | 
**NetworkLabelAssignmentSpecs** | Pointer to [**[]NetworkLabelAssignmentSettingsUpdateSpec**](NetworkLabelAssignmentSettingsUpdateSpec.md) | Automatic network label assignment feature settings for this NIC.  By default, newly provisioned machines of an automated farm retain their parent image&#39;s network labels on each of their network interface cards. In certain circumstances, notably dealing with VLAN subset sizing and DHCP IP address availability, it may be desirable for the cloned VM to instead use different network labels for these newly provisioned machines. This feature allows an administrator to provide a per NIC list of network labels and their maximum availability to be automatically distributed to newly provisioned machines.&lt;br&gt;If this is not set, the feature is disabled.&lt;br&gt;Starting at the alphabetically first network label spec in the list that has not yet been assigned its maximum count for this NIC on this VM, the VM will have its next provisioned machine&#39;s NIC assigned that label. If all network labels in this list have reached their maximum count, this VM will have further provisioned machines assigned the last label in the list over capacity, and an error will be logged. Not all labels need be configured.&lt;br&gt;Any changes to these settings will be applicable only to RDS servers provisioned after the change. Already provisioned RDS servers will never have their network label assignments altered by this feature. | [optional] 

## Methods

### NewFarmNetworkInterfaceCardSettingsUpdateSpec

`func NewFarmNetworkInterfaceCardSettingsUpdateSpec(networkInterfaceCardId string, ) *FarmNetworkInterfaceCardSettingsUpdateSpec`

NewFarmNetworkInterfaceCardSettingsUpdateSpec instantiates a new FarmNetworkInterfaceCardSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmNetworkInterfaceCardSettingsUpdateSpecWithDefaults

`func NewFarmNetworkInterfaceCardSettingsUpdateSpecWithDefaults() *FarmNetworkInterfaceCardSettingsUpdateSpec`

NewFarmNetworkInterfaceCardSettingsUpdateSpecWithDefaults instantiates a new FarmNetworkInterfaceCardSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterfaceCardId

`func (o *FarmNetworkInterfaceCardSettingsUpdateSpec) GetNetworkInterfaceCardId() string`

GetNetworkInterfaceCardId returns the NetworkInterfaceCardId field if non-nil, zero value otherwise.

### GetNetworkInterfaceCardIdOk

`func (o *FarmNetworkInterfaceCardSettingsUpdateSpec) GetNetworkInterfaceCardIdOk() (*string, bool)`

GetNetworkInterfaceCardIdOk returns a tuple with the NetworkInterfaceCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceCardId

`func (o *FarmNetworkInterfaceCardSettingsUpdateSpec) SetNetworkInterfaceCardId(v string)`

SetNetworkInterfaceCardId sets NetworkInterfaceCardId field to given value.


### GetNetworkLabelAssignmentSpecs

`func (o *FarmNetworkInterfaceCardSettingsUpdateSpec) GetNetworkLabelAssignmentSpecs() []NetworkLabelAssignmentSettingsUpdateSpec`

GetNetworkLabelAssignmentSpecs returns the NetworkLabelAssignmentSpecs field if non-nil, zero value otherwise.

### GetNetworkLabelAssignmentSpecsOk

`func (o *FarmNetworkInterfaceCardSettingsUpdateSpec) GetNetworkLabelAssignmentSpecsOk() (*[]NetworkLabelAssignmentSettingsUpdateSpec, bool)`

GetNetworkLabelAssignmentSpecsOk returns a tuple with the NetworkLabelAssignmentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLabelAssignmentSpecs

`func (o *FarmNetworkInterfaceCardSettingsUpdateSpec) SetNetworkLabelAssignmentSpecs(v []NetworkLabelAssignmentSettingsUpdateSpec)`

SetNetworkLabelAssignmentSpecs sets NetworkLabelAssignmentSpecs field to given value.

### HasNetworkLabelAssignmentSpecs

`func (o *FarmNetworkInterfaceCardSettingsUpdateSpec) HasNetworkLabelAssignmentSpecs() bool`

HasNetworkLabelAssignmentSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


