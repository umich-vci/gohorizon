# FarmInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | Access groups can organize the entities such as farms in the organization. They can also be used for delegated administration.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**AutomatedFarmSettings** | Pointer to [**FarmAutomatedSettingsInfo**](FarmAutomatedSettingsInfo.md) |  | [optional] 
**DeleteInProgress** | Pointer to **bool** | Indicates whether the farm is in the process of being deleted.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Description** | Pointer to **string** | Description of the farm. The maximum length is 1024 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DesktopPoolId** | Pointer to **string** | ID of RDS desktop pool associated with the farm.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the farm. The maximum length is 256 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DisplayProtocolSettings** | Pointer to [**FarmDisplayProtocolSettingsInfo**](FarmDisplayProtocolSettingsInfo.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the farm is enabled for brokering.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing farm.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**LoadBalancerSettings** | Pointer to [**RDSHLoadBalancerSettingsInfo**](RDSHLoadBalancerSettingsInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the farm. The maximum length is 64 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**ServerErrorThreshold** | Pointer to **int32** | The minimum number of machines that must be fully operational in order to avoid showing the farm in an error state.  | [optional] 
**SessionSettings** | Pointer to [**FarmSessionSettingsInfo**](FarmSessionSettingsInfo.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the farm.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * AUTOMATED: Automated Farm. * MANUAL: Manual Farm. | [optional] 
**UseCustomScriptForLoadBalancing** | Pointer to **bool** | Indicates whether to use custom scripts for load balancing or not. | [optional] 

## Methods

### NewFarmInfoV2

`func NewFarmInfoV2() *FarmInfoV2`

NewFarmInfoV2 instantiates a new FarmInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmInfoV2WithDefaults

`func NewFarmInfoV2WithDefaults() *FarmInfoV2`

NewFarmInfoV2WithDefaults instantiates a new FarmInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *FarmInfoV2) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *FarmInfoV2) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *FarmInfoV2) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *FarmInfoV2) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAutomatedFarmSettings

`func (o *FarmInfoV2) GetAutomatedFarmSettings() FarmAutomatedSettingsInfo`

GetAutomatedFarmSettings returns the AutomatedFarmSettings field if non-nil, zero value otherwise.

### GetAutomatedFarmSettingsOk

`func (o *FarmInfoV2) GetAutomatedFarmSettingsOk() (*FarmAutomatedSettingsInfo, bool)`

GetAutomatedFarmSettingsOk returns a tuple with the AutomatedFarmSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedFarmSettings

`func (o *FarmInfoV2) SetAutomatedFarmSettings(v FarmAutomatedSettingsInfo)`

SetAutomatedFarmSettings sets AutomatedFarmSettings field to given value.

### HasAutomatedFarmSettings

`func (o *FarmInfoV2) HasAutomatedFarmSettings() bool`

HasAutomatedFarmSettings returns a boolean if a field has been set.

### GetDeleteInProgress

`func (o *FarmInfoV2) GetDeleteInProgress() bool`

GetDeleteInProgress returns the DeleteInProgress field if non-nil, zero value otherwise.

### GetDeleteInProgressOk

`func (o *FarmInfoV2) GetDeleteInProgressOk() (*bool, bool)`

GetDeleteInProgressOk returns a tuple with the DeleteInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInProgress

`func (o *FarmInfoV2) SetDeleteInProgress(v bool)`

SetDeleteInProgress sets DeleteInProgress field to given value.

### HasDeleteInProgress

`func (o *FarmInfoV2) HasDeleteInProgress() bool`

HasDeleteInProgress returns a boolean if a field has been set.

### GetDescription

`func (o *FarmInfoV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FarmInfoV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FarmInfoV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FarmInfoV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesktopPoolId

`func (o *FarmInfoV2) GetDesktopPoolId() string`

GetDesktopPoolId returns the DesktopPoolId field if non-nil, zero value otherwise.

### GetDesktopPoolIdOk

`func (o *FarmInfoV2) GetDesktopPoolIdOk() (*string, bool)`

GetDesktopPoolIdOk returns a tuple with the DesktopPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolId

`func (o *FarmInfoV2) SetDesktopPoolId(v string)`

SetDesktopPoolId sets DesktopPoolId field to given value.

### HasDesktopPoolId

`func (o *FarmInfoV2) HasDesktopPoolId() bool`

HasDesktopPoolId returns a boolean if a field has been set.

### GetDisplayName

`func (o *FarmInfoV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FarmInfoV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FarmInfoV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FarmInfoV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayProtocolSettings

`func (o *FarmInfoV2) GetDisplayProtocolSettings() FarmDisplayProtocolSettingsInfo`

GetDisplayProtocolSettings returns the DisplayProtocolSettings field if non-nil, zero value otherwise.

### GetDisplayProtocolSettingsOk

`func (o *FarmInfoV2) GetDisplayProtocolSettingsOk() (*FarmDisplayProtocolSettingsInfo, bool)`

GetDisplayProtocolSettingsOk returns a tuple with the DisplayProtocolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocolSettings

`func (o *FarmInfoV2) SetDisplayProtocolSettings(v FarmDisplayProtocolSettingsInfo)`

SetDisplayProtocolSettings sets DisplayProtocolSettings field to given value.

### HasDisplayProtocolSettings

`func (o *FarmInfoV2) HasDisplayProtocolSettings() bool`

HasDisplayProtocolSettings returns a boolean if a field has been set.

### GetEnabled

`func (o *FarmInfoV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FarmInfoV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FarmInfoV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FarmInfoV2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *FarmInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FarmInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FarmInfoV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FarmInfoV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancerSettings

`func (o *FarmInfoV2) GetLoadBalancerSettings() RDSHLoadBalancerSettingsInfo`

GetLoadBalancerSettings returns the LoadBalancerSettings field if non-nil, zero value otherwise.

### GetLoadBalancerSettingsOk

`func (o *FarmInfoV2) GetLoadBalancerSettingsOk() (*RDSHLoadBalancerSettingsInfo, bool)`

GetLoadBalancerSettingsOk returns a tuple with the LoadBalancerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSettings

`func (o *FarmInfoV2) SetLoadBalancerSettings(v RDSHLoadBalancerSettingsInfo)`

SetLoadBalancerSettings sets LoadBalancerSettings field to given value.

### HasLoadBalancerSettings

`func (o *FarmInfoV2) HasLoadBalancerSettings() bool`

HasLoadBalancerSettings returns a boolean if a field has been set.

### GetName

`func (o *FarmInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FarmInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FarmInfoV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FarmInfoV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerErrorThreshold

`func (o *FarmInfoV2) GetServerErrorThreshold() int32`

GetServerErrorThreshold returns the ServerErrorThreshold field if non-nil, zero value otherwise.

### GetServerErrorThresholdOk

`func (o *FarmInfoV2) GetServerErrorThresholdOk() (*int32, bool)`

GetServerErrorThresholdOk returns a tuple with the ServerErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerErrorThreshold

`func (o *FarmInfoV2) SetServerErrorThreshold(v int32)`

SetServerErrorThreshold sets ServerErrorThreshold field to given value.

### HasServerErrorThreshold

`func (o *FarmInfoV2) HasServerErrorThreshold() bool`

HasServerErrorThreshold returns a boolean if a field has been set.

### GetSessionSettings

`func (o *FarmInfoV2) GetSessionSettings() FarmSessionSettingsInfo`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *FarmInfoV2) GetSessionSettingsOk() (*FarmSessionSettingsInfo, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *FarmInfoV2) SetSessionSettings(v FarmSessionSettingsInfo)`

SetSessionSettings sets SessionSettings field to given value.

### HasSessionSettings

`func (o *FarmInfoV2) HasSessionSettings() bool`

HasSessionSettings returns a boolean if a field has been set.

### GetType

`func (o *FarmInfoV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FarmInfoV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FarmInfoV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FarmInfoV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseCustomScriptForLoadBalancing

`func (o *FarmInfoV2) GetUseCustomScriptForLoadBalancing() bool`

GetUseCustomScriptForLoadBalancing returns the UseCustomScriptForLoadBalancing field if non-nil, zero value otherwise.

### GetUseCustomScriptForLoadBalancingOk

`func (o *FarmInfoV2) GetUseCustomScriptForLoadBalancingOk() (*bool, bool)`

GetUseCustomScriptForLoadBalancingOk returns a tuple with the UseCustomScriptForLoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomScriptForLoadBalancing

`func (o *FarmInfoV2) SetUseCustomScriptForLoadBalancing(v bool)`

SetUseCustomScriptForLoadBalancing sets UseCustomScriptForLoadBalancing field to given value.

### HasUseCustomScriptForLoadBalancing

`func (o *FarmInfoV2) HasUseCustomScriptForLoadBalancing() bool`

HasUseCustomScriptForLoadBalancing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


