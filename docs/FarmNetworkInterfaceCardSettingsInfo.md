# FarmNetworkInterfaceCardSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaceCardId** | Pointer to **string** | ID of the network interface card for these settings. | [optional] 
**NetworkInterfaceCardName** | Pointer to **string** | Name of the network interface card for these settings. | [optional] 
**NetworkLabelAssignmentSpecs** | Pointer to [**[]NetworkLabelAssignmentSettingsInfo**](NetworkLabelAssignmentSettingsInfo.md) | Automatic network label assignment feature settings for this NIC. | [optional] 

## Methods

### NewFarmNetworkInterfaceCardSettingsInfo

`func NewFarmNetworkInterfaceCardSettingsInfo() *FarmNetworkInterfaceCardSettingsInfo`

NewFarmNetworkInterfaceCardSettingsInfo instantiates a new FarmNetworkInterfaceCardSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmNetworkInterfaceCardSettingsInfoWithDefaults

`func NewFarmNetworkInterfaceCardSettingsInfoWithDefaults() *FarmNetworkInterfaceCardSettingsInfo`

NewFarmNetworkInterfaceCardSettingsInfoWithDefaults instantiates a new FarmNetworkInterfaceCardSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterfaceCardId

`func (o *FarmNetworkInterfaceCardSettingsInfo) GetNetworkInterfaceCardId() string`

GetNetworkInterfaceCardId returns the NetworkInterfaceCardId field if non-nil, zero value otherwise.

### GetNetworkInterfaceCardIdOk

`func (o *FarmNetworkInterfaceCardSettingsInfo) GetNetworkInterfaceCardIdOk() (*string, bool)`

GetNetworkInterfaceCardIdOk returns a tuple with the NetworkInterfaceCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceCardId

`func (o *FarmNetworkInterfaceCardSettingsInfo) SetNetworkInterfaceCardId(v string)`

SetNetworkInterfaceCardId sets NetworkInterfaceCardId field to given value.

### HasNetworkInterfaceCardId

`func (o *FarmNetworkInterfaceCardSettingsInfo) HasNetworkInterfaceCardId() bool`

HasNetworkInterfaceCardId returns a boolean if a field has been set.

### GetNetworkInterfaceCardName

`func (o *FarmNetworkInterfaceCardSettingsInfo) GetNetworkInterfaceCardName() string`

GetNetworkInterfaceCardName returns the NetworkInterfaceCardName field if non-nil, zero value otherwise.

### GetNetworkInterfaceCardNameOk

`func (o *FarmNetworkInterfaceCardSettingsInfo) GetNetworkInterfaceCardNameOk() (*string, bool)`

GetNetworkInterfaceCardNameOk returns a tuple with the NetworkInterfaceCardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceCardName

`func (o *FarmNetworkInterfaceCardSettingsInfo) SetNetworkInterfaceCardName(v string)`

SetNetworkInterfaceCardName sets NetworkInterfaceCardName field to given value.

### HasNetworkInterfaceCardName

`func (o *FarmNetworkInterfaceCardSettingsInfo) HasNetworkInterfaceCardName() bool`

HasNetworkInterfaceCardName returns a boolean if a field has been set.

### GetNetworkLabelAssignmentSpecs

`func (o *FarmNetworkInterfaceCardSettingsInfo) GetNetworkLabelAssignmentSpecs() []NetworkLabelAssignmentSettingsInfo`

GetNetworkLabelAssignmentSpecs returns the NetworkLabelAssignmentSpecs field if non-nil, zero value otherwise.

### GetNetworkLabelAssignmentSpecsOk

`func (o *FarmNetworkInterfaceCardSettingsInfo) GetNetworkLabelAssignmentSpecsOk() (*[]NetworkLabelAssignmentSettingsInfo, bool)`

GetNetworkLabelAssignmentSpecsOk returns a tuple with the NetworkLabelAssignmentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLabelAssignmentSpecs

`func (o *FarmNetworkInterfaceCardSettingsInfo) SetNetworkLabelAssignmentSpecs(v []NetworkLabelAssignmentSettingsInfo)`

SetNetworkLabelAssignmentSpecs sets NetworkLabelAssignmentSpecs field to given value.

### HasNetworkLabelAssignmentSpecs

`func (o *FarmNetworkInterfaceCardSettingsInfo) HasNetworkLabelAssignmentSpecs() bool`

HasNetworkLabelAssignmentSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


