# FarmUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | **string** | Access groups can organize the entities such as farms in the organization. They can also be used for delegated administration. | 
**AutomatedFarmSettings** | Pointer to [**FarmAutomatedSettingsUpdateSpec**](FarmAutomatedSettingsUpdateSpec.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the farm. | [optional] 
**DisplayName** | **string** | Display name of the farm. | 
**DisplayProtocolSettings** | [**FarmDisplayProtocolSettingsUpdateSpec**](FarmDisplayProtocolSettingsUpdateSpec.md) |  | 
**Enabled** | **bool** | Indicates whether the farm is enabled for brokering. | 
**LoadBalancerSettings** | Pointer to [**RDSHLoadBalancerSettingsUpdateSpec**](RDSHLoadBalancerSettingsUpdateSpec.md) |  | [optional] 
**ServerErrorThreshold** | **int32** | The minimum number of machines that must be fully operational in order to avoid showing the farm in an error state.  | 
**SessionSettings** | [**FarmSessionSettingsUpdateSpec**](FarmSessionSettingsUpdateSpec.md) |  | 
**UseCustomScriptForLoadBalancing** | **bool** | Indicates whether to use custom scripts for load balancing or not.  | 

## Methods

### NewFarmUpdateSpec

`func NewFarmUpdateSpec(accessGroupId string, displayName string, displayProtocolSettings FarmDisplayProtocolSettingsUpdateSpec, enabled bool, serverErrorThreshold int32, sessionSettings FarmSessionSettingsUpdateSpec, useCustomScriptForLoadBalancing bool, ) *FarmUpdateSpec`

NewFarmUpdateSpec instantiates a new FarmUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmUpdateSpecWithDefaults

`func NewFarmUpdateSpecWithDefaults() *FarmUpdateSpec`

NewFarmUpdateSpecWithDefaults instantiates a new FarmUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *FarmUpdateSpec) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *FarmUpdateSpec) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *FarmUpdateSpec) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.


### GetAutomatedFarmSettings

`func (o *FarmUpdateSpec) GetAutomatedFarmSettings() FarmAutomatedSettingsUpdateSpec`

GetAutomatedFarmSettings returns the AutomatedFarmSettings field if non-nil, zero value otherwise.

### GetAutomatedFarmSettingsOk

`func (o *FarmUpdateSpec) GetAutomatedFarmSettingsOk() (*FarmAutomatedSettingsUpdateSpec, bool)`

GetAutomatedFarmSettingsOk returns a tuple with the AutomatedFarmSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedFarmSettings

`func (o *FarmUpdateSpec) SetAutomatedFarmSettings(v FarmAutomatedSettingsUpdateSpec)`

SetAutomatedFarmSettings sets AutomatedFarmSettings field to given value.

### HasAutomatedFarmSettings

`func (o *FarmUpdateSpec) HasAutomatedFarmSettings() bool`

HasAutomatedFarmSettings returns a boolean if a field has been set.

### GetDescription

`func (o *FarmUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FarmUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FarmUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FarmUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FarmUpdateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FarmUpdateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FarmUpdateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDisplayProtocolSettings

`func (o *FarmUpdateSpec) GetDisplayProtocolSettings() FarmDisplayProtocolSettingsUpdateSpec`

GetDisplayProtocolSettings returns the DisplayProtocolSettings field if non-nil, zero value otherwise.

### GetDisplayProtocolSettingsOk

`func (o *FarmUpdateSpec) GetDisplayProtocolSettingsOk() (*FarmDisplayProtocolSettingsUpdateSpec, bool)`

GetDisplayProtocolSettingsOk returns a tuple with the DisplayProtocolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocolSettings

`func (o *FarmUpdateSpec) SetDisplayProtocolSettings(v FarmDisplayProtocolSettingsUpdateSpec)`

SetDisplayProtocolSettings sets DisplayProtocolSettings field to given value.


### GetEnabled

`func (o *FarmUpdateSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FarmUpdateSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FarmUpdateSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoadBalancerSettings

`func (o *FarmUpdateSpec) GetLoadBalancerSettings() RDSHLoadBalancerSettingsUpdateSpec`

GetLoadBalancerSettings returns the LoadBalancerSettings field if non-nil, zero value otherwise.

### GetLoadBalancerSettingsOk

`func (o *FarmUpdateSpec) GetLoadBalancerSettingsOk() (*RDSHLoadBalancerSettingsUpdateSpec, bool)`

GetLoadBalancerSettingsOk returns a tuple with the LoadBalancerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSettings

`func (o *FarmUpdateSpec) SetLoadBalancerSettings(v RDSHLoadBalancerSettingsUpdateSpec)`

SetLoadBalancerSettings sets LoadBalancerSettings field to given value.

### HasLoadBalancerSettings

`func (o *FarmUpdateSpec) HasLoadBalancerSettings() bool`

HasLoadBalancerSettings returns a boolean if a field has been set.

### GetServerErrorThreshold

`func (o *FarmUpdateSpec) GetServerErrorThreshold() int32`

GetServerErrorThreshold returns the ServerErrorThreshold field if non-nil, zero value otherwise.

### GetServerErrorThresholdOk

`func (o *FarmUpdateSpec) GetServerErrorThresholdOk() (*int32, bool)`

GetServerErrorThresholdOk returns a tuple with the ServerErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerErrorThreshold

`func (o *FarmUpdateSpec) SetServerErrorThreshold(v int32)`

SetServerErrorThreshold sets ServerErrorThreshold field to given value.


### GetSessionSettings

`func (o *FarmUpdateSpec) GetSessionSettings() FarmSessionSettingsUpdateSpec`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *FarmUpdateSpec) GetSessionSettingsOk() (*FarmSessionSettingsUpdateSpec, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *FarmUpdateSpec) SetSessionSettings(v FarmSessionSettingsUpdateSpec)`

SetSessionSettings sets SessionSettings field to given value.


### GetUseCustomScriptForLoadBalancing

`func (o *FarmUpdateSpec) GetUseCustomScriptForLoadBalancing() bool`

GetUseCustomScriptForLoadBalancing returns the UseCustomScriptForLoadBalancing field if non-nil, zero value otherwise.

### GetUseCustomScriptForLoadBalancingOk

`func (o *FarmUpdateSpec) GetUseCustomScriptForLoadBalancingOk() (*bool, bool)`

GetUseCustomScriptForLoadBalancingOk returns a tuple with the UseCustomScriptForLoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomScriptForLoadBalancing

`func (o *FarmUpdateSpec) SetUseCustomScriptForLoadBalancing(v bool)`

SetUseCustomScriptForLoadBalancing sets UseCustomScriptForLoadBalancing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


