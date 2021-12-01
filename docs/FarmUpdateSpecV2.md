# FarmUpdateSpecV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | **string** | Access groups can organize the entities such as farms in the organization. They can also be used for delegated administration. | 
**AutomatedFarmSettings** | Pointer to [**FarmAutomatedSettingsUpdateSpecV2**](FarmAutomatedSettingsUpdateSpecV2.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the farm. | [optional] 
**DisplayName** | **string** | Display name of the farm. | 
**DisplayProtocolSettings** | [**FarmDisplayProtocolSettingsUpdateSpec**](FarmDisplayProtocolSettingsUpdateSpec.md) |  | 
**Enabled** | **bool** | Indicates whether the farm is enabled for brokering. | 
**LoadBalancerSettings** | Pointer to [**RDSHLoadBalancerSettingsUpdateSpec**](RDSHLoadBalancerSettingsUpdateSpec.md) |  | [optional] 
**ServerErrorThreshold** | **int32** | The minimum number of machines that must be fully operational in order to avoid showing the farm in an error state.  | 
**SessionSettings** | [**FarmSessionSettingsUpdateSpec**](FarmSessionSettingsUpdateSpec.md) |  | 
**UseCustomScriptForLoadBalancing** | **bool** | Indicates whether to use custom scripts for load balancing or not.  | 

## Methods

### NewFarmUpdateSpecV2

`func NewFarmUpdateSpecV2(accessGroupId string, displayName string, displayProtocolSettings FarmDisplayProtocolSettingsUpdateSpec, enabled bool, serverErrorThreshold int32, sessionSettings FarmSessionSettingsUpdateSpec, useCustomScriptForLoadBalancing bool, ) *FarmUpdateSpecV2`

NewFarmUpdateSpecV2 instantiates a new FarmUpdateSpecV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmUpdateSpecV2WithDefaults

`func NewFarmUpdateSpecV2WithDefaults() *FarmUpdateSpecV2`

NewFarmUpdateSpecV2WithDefaults instantiates a new FarmUpdateSpecV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *FarmUpdateSpecV2) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *FarmUpdateSpecV2) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *FarmUpdateSpecV2) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.


### GetAutomatedFarmSettings

`func (o *FarmUpdateSpecV2) GetAutomatedFarmSettings() FarmAutomatedSettingsUpdateSpecV2`

GetAutomatedFarmSettings returns the AutomatedFarmSettings field if non-nil, zero value otherwise.

### GetAutomatedFarmSettingsOk

`func (o *FarmUpdateSpecV2) GetAutomatedFarmSettingsOk() (*FarmAutomatedSettingsUpdateSpecV2, bool)`

GetAutomatedFarmSettingsOk returns a tuple with the AutomatedFarmSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedFarmSettings

`func (o *FarmUpdateSpecV2) SetAutomatedFarmSettings(v FarmAutomatedSettingsUpdateSpecV2)`

SetAutomatedFarmSettings sets AutomatedFarmSettings field to given value.

### HasAutomatedFarmSettings

`func (o *FarmUpdateSpecV2) HasAutomatedFarmSettings() bool`

HasAutomatedFarmSettings returns a boolean if a field has been set.

### GetDescription

`func (o *FarmUpdateSpecV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FarmUpdateSpecV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FarmUpdateSpecV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FarmUpdateSpecV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FarmUpdateSpecV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FarmUpdateSpecV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FarmUpdateSpecV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDisplayProtocolSettings

`func (o *FarmUpdateSpecV2) GetDisplayProtocolSettings() FarmDisplayProtocolSettingsUpdateSpec`

GetDisplayProtocolSettings returns the DisplayProtocolSettings field if non-nil, zero value otherwise.

### GetDisplayProtocolSettingsOk

`func (o *FarmUpdateSpecV2) GetDisplayProtocolSettingsOk() (*FarmDisplayProtocolSettingsUpdateSpec, bool)`

GetDisplayProtocolSettingsOk returns a tuple with the DisplayProtocolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocolSettings

`func (o *FarmUpdateSpecV2) SetDisplayProtocolSettings(v FarmDisplayProtocolSettingsUpdateSpec)`

SetDisplayProtocolSettings sets DisplayProtocolSettings field to given value.


### GetEnabled

`func (o *FarmUpdateSpecV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FarmUpdateSpecV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FarmUpdateSpecV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoadBalancerSettings

`func (o *FarmUpdateSpecV2) GetLoadBalancerSettings() RDSHLoadBalancerSettingsUpdateSpec`

GetLoadBalancerSettings returns the LoadBalancerSettings field if non-nil, zero value otherwise.

### GetLoadBalancerSettingsOk

`func (o *FarmUpdateSpecV2) GetLoadBalancerSettingsOk() (*RDSHLoadBalancerSettingsUpdateSpec, bool)`

GetLoadBalancerSettingsOk returns a tuple with the LoadBalancerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSettings

`func (o *FarmUpdateSpecV2) SetLoadBalancerSettings(v RDSHLoadBalancerSettingsUpdateSpec)`

SetLoadBalancerSettings sets LoadBalancerSettings field to given value.

### HasLoadBalancerSettings

`func (o *FarmUpdateSpecV2) HasLoadBalancerSettings() bool`

HasLoadBalancerSettings returns a boolean if a field has been set.

### GetServerErrorThreshold

`func (o *FarmUpdateSpecV2) GetServerErrorThreshold() int32`

GetServerErrorThreshold returns the ServerErrorThreshold field if non-nil, zero value otherwise.

### GetServerErrorThresholdOk

`func (o *FarmUpdateSpecV2) GetServerErrorThresholdOk() (*int32, bool)`

GetServerErrorThresholdOk returns a tuple with the ServerErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerErrorThreshold

`func (o *FarmUpdateSpecV2) SetServerErrorThreshold(v int32)`

SetServerErrorThreshold sets ServerErrorThreshold field to given value.


### GetSessionSettings

`func (o *FarmUpdateSpecV2) GetSessionSettings() FarmSessionSettingsUpdateSpec`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *FarmUpdateSpecV2) GetSessionSettingsOk() (*FarmSessionSettingsUpdateSpec, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *FarmUpdateSpecV2) SetSessionSettings(v FarmSessionSettingsUpdateSpec)`

SetSessionSettings sets SessionSettings field to given value.


### GetUseCustomScriptForLoadBalancing

`func (o *FarmUpdateSpecV2) GetUseCustomScriptForLoadBalancing() bool`

GetUseCustomScriptForLoadBalancing returns the UseCustomScriptForLoadBalancing field if non-nil, zero value otherwise.

### GetUseCustomScriptForLoadBalancingOk

`func (o *FarmUpdateSpecV2) GetUseCustomScriptForLoadBalancingOk() (*bool, bool)`

GetUseCustomScriptForLoadBalancingOk returns a tuple with the UseCustomScriptForLoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomScriptForLoadBalancing

`func (o *FarmUpdateSpecV2) SetUseCustomScriptForLoadBalancing(v bool)`

SetUseCustomScriptForLoadBalancing sets UseCustomScriptForLoadBalancing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


