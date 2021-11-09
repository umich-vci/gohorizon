# NetworkLabelAssignmentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not this specification is enabled. While this specification is disabled, automatic network label assigment for this desktop pool will skip over the network label in this spec. | [optional] 
**MaxLabel** | Pointer to **int32** | The maximum number of times this label can be assigned to a machine. Note this count only encompasses this spec. That is, this label may be used for other NICs and in other Desktop pools, but those assignments will not be counted towards this total. This count also does not include assignments of this label to machines not under the control of View. | [optional] 
**MaxLabelType** | Pointer to **string** | This type specifies whether or not there is a maximum limit to the number of times this label may be assigned to machines within this spec. While this specification is enabled and unlimited, specs after this one in the NIC&#39;s network label specification list will never be used. * UNLIMITED: The network label assignment specification has no limit on the number of labels to assign. * LIMITED: The network label assignment specification has a limited number of labels to assign. | [optional] 
**NetworkLabelName** | Pointer to **string** | The network label id for this spec. This network label must not have any incompatibility reasons that would preclude it from automatic machine assignment. | [optional] 

## Methods

### NewNetworkLabelAssignmentSettings

`func NewNetworkLabelAssignmentSettings() *NetworkLabelAssignmentSettings`

NewNetworkLabelAssignmentSettings instantiates a new NetworkLabelAssignmentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLabelAssignmentSettingsWithDefaults

`func NewNetworkLabelAssignmentSettingsWithDefaults() *NetworkLabelAssignmentSettings`

NewNetworkLabelAssignmentSettingsWithDefaults instantiates a new NetworkLabelAssignmentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworkLabelAssignmentSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkLabelAssignmentSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkLabelAssignmentSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkLabelAssignmentSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxLabel

`func (o *NetworkLabelAssignmentSettings) GetMaxLabel() int32`

GetMaxLabel returns the MaxLabel field if non-nil, zero value otherwise.

### GetMaxLabelOk

`func (o *NetworkLabelAssignmentSettings) GetMaxLabelOk() (*int32, bool)`

GetMaxLabelOk returns a tuple with the MaxLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLabel

`func (o *NetworkLabelAssignmentSettings) SetMaxLabel(v int32)`

SetMaxLabel sets MaxLabel field to given value.

### HasMaxLabel

`func (o *NetworkLabelAssignmentSettings) HasMaxLabel() bool`

HasMaxLabel returns a boolean if a field has been set.

### GetMaxLabelType

`func (o *NetworkLabelAssignmentSettings) GetMaxLabelType() string`

GetMaxLabelType returns the MaxLabelType field if non-nil, zero value otherwise.

### GetMaxLabelTypeOk

`func (o *NetworkLabelAssignmentSettings) GetMaxLabelTypeOk() (*string, bool)`

GetMaxLabelTypeOk returns a tuple with the MaxLabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLabelType

`func (o *NetworkLabelAssignmentSettings) SetMaxLabelType(v string)`

SetMaxLabelType sets MaxLabelType field to given value.

### HasMaxLabelType

`func (o *NetworkLabelAssignmentSettings) HasMaxLabelType() bool`

HasMaxLabelType returns a boolean if a field has been set.

### GetNetworkLabelName

`func (o *NetworkLabelAssignmentSettings) GetNetworkLabelName() string`

GetNetworkLabelName returns the NetworkLabelName field if non-nil, zero value otherwise.

### GetNetworkLabelNameOk

`func (o *NetworkLabelAssignmentSettings) GetNetworkLabelNameOk() (*string, bool)`

GetNetworkLabelNameOk returns a tuple with the NetworkLabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLabelName

`func (o *NetworkLabelAssignmentSettings) SetNetworkLabelName(v string)`

SetNetworkLabelName sets NetworkLabelName field to given value.

### HasNetworkLabelName

`func (o *NetworkLabelAssignmentSettings) HasNetworkLabelName() bool`

HasNetworkLabelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


