# NetworkLabelAssignmentSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether or not this specification is enabled. While this specification is disabled, automatic network label assignment for this automated farm will skip over the network label in this spec. Default value is true. | [optional] 
**MaxLabel** | Pointer to **int32** | The maximum number of times this label can be assigned to a RDS server. Note this count only encompasses this spec. That is, this label may be used for other NICs and in other farms, but those assignments will not be counted towards this total. This count also does not include assignments of this label to RDS server not under the control of View. This is required if max_label_type is set to LIMITED. | [optional] 
**MaxLabelType** | **string** | This type specifies whether or not there is a maximum limit to the number of times this label may be assigned to machines within this spec. While this specification is enabled and unlimited, specs after this one in the NIC&#39;s network label specification list will never be used. * UNLIMITED: The network label assignment specification has no limit on the number of labels to assign. * LIMITED: The network label assignment specification has a limited number of labels to assign. | 
**NetworkLabelName** | **string** | The network label id for this spec. This network label must not have any incompatibility reasons that would preclude it from automatic machine assignment. | 

## Methods

### NewNetworkLabelAssignmentSettingsCreateSpec

`func NewNetworkLabelAssignmentSettingsCreateSpec(maxLabelType string, networkLabelName string, ) *NetworkLabelAssignmentSettingsCreateSpec`

NewNetworkLabelAssignmentSettingsCreateSpec instantiates a new NetworkLabelAssignmentSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLabelAssignmentSettingsCreateSpecWithDefaults

`func NewNetworkLabelAssignmentSettingsCreateSpecWithDefaults() *NetworkLabelAssignmentSettingsCreateSpec`

NewNetworkLabelAssignmentSettingsCreateSpecWithDefaults instantiates a new NetworkLabelAssignmentSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworkLabelAssignmentSettingsCreateSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkLabelAssignmentSettingsCreateSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkLabelAssignmentSettingsCreateSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkLabelAssignmentSettingsCreateSpec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxLabel

`func (o *NetworkLabelAssignmentSettingsCreateSpec) GetMaxLabel() int32`

GetMaxLabel returns the MaxLabel field if non-nil, zero value otherwise.

### GetMaxLabelOk

`func (o *NetworkLabelAssignmentSettingsCreateSpec) GetMaxLabelOk() (*int32, bool)`

GetMaxLabelOk returns a tuple with the MaxLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLabel

`func (o *NetworkLabelAssignmentSettingsCreateSpec) SetMaxLabel(v int32)`

SetMaxLabel sets MaxLabel field to given value.

### HasMaxLabel

`func (o *NetworkLabelAssignmentSettingsCreateSpec) HasMaxLabel() bool`

HasMaxLabel returns a boolean if a field has been set.

### GetMaxLabelType

`func (o *NetworkLabelAssignmentSettingsCreateSpec) GetMaxLabelType() string`

GetMaxLabelType returns the MaxLabelType field if non-nil, zero value otherwise.

### GetMaxLabelTypeOk

`func (o *NetworkLabelAssignmentSettingsCreateSpec) GetMaxLabelTypeOk() (*string, bool)`

GetMaxLabelTypeOk returns a tuple with the MaxLabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLabelType

`func (o *NetworkLabelAssignmentSettingsCreateSpec) SetMaxLabelType(v string)`

SetMaxLabelType sets MaxLabelType field to given value.


### GetNetworkLabelName

`func (o *NetworkLabelAssignmentSettingsCreateSpec) GetNetworkLabelName() string`

GetNetworkLabelName returns the NetworkLabelName field if non-nil, zero value otherwise.

### GetNetworkLabelNameOk

`func (o *NetworkLabelAssignmentSettingsCreateSpec) GetNetworkLabelNameOk() (*string, bool)`

GetNetworkLabelNameOk returns a tuple with the NetworkLabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLabelName

`func (o *NetworkLabelAssignmentSettingsCreateSpec) SetNetworkLabelName(v string)`

SetNetworkLabelName sets NetworkLabelName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


