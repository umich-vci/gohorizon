# FarmLoadBalancerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomScriptInUse** | Pointer to **bool** | Indicates whether to use custom scripts for Load Balancing.Default is false. | [optional] 
**LbMetricSettings** | Pointer to [**LoadBalancerMetricSettings**](LoadBalancerMetricSettings.md) |  | [optional] 

## Methods

### NewFarmLoadBalancerSettings

`func NewFarmLoadBalancerSettings() *FarmLoadBalancerSettings`

NewFarmLoadBalancerSettings instantiates a new FarmLoadBalancerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmLoadBalancerSettingsWithDefaults

`func NewFarmLoadBalancerSettingsWithDefaults() *FarmLoadBalancerSettings`

NewFarmLoadBalancerSettingsWithDefaults instantiates a new FarmLoadBalancerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomScriptInUse

`func (o *FarmLoadBalancerSettings) GetCustomScriptInUse() bool`

GetCustomScriptInUse returns the CustomScriptInUse field if non-nil, zero value otherwise.

### GetCustomScriptInUseOk

`func (o *FarmLoadBalancerSettings) GetCustomScriptInUseOk() (*bool, bool)`

GetCustomScriptInUseOk returns a tuple with the CustomScriptInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomScriptInUse

`func (o *FarmLoadBalancerSettings) SetCustomScriptInUse(v bool)`

SetCustomScriptInUse sets CustomScriptInUse field to given value.

### HasCustomScriptInUse

`func (o *FarmLoadBalancerSettings) HasCustomScriptInUse() bool`

HasCustomScriptInUse returns a boolean if a field has been set.

### GetLbMetricSettings

`func (o *FarmLoadBalancerSettings) GetLbMetricSettings() LoadBalancerMetricSettings`

GetLbMetricSettings returns the LbMetricSettings field if non-nil, zero value otherwise.

### GetLbMetricSettingsOk

`func (o *FarmLoadBalancerSettings) GetLbMetricSettingsOk() (*LoadBalancerMetricSettings, bool)`

GetLbMetricSettingsOk returns a tuple with the LbMetricSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMetricSettings

`func (o *FarmLoadBalancerSettings) SetLbMetricSettings(v LoadBalancerMetricSettings)`

SetLbMetricSettings sets LbMetricSettings field to given value.

### HasLbMetricSettings

`func (o *FarmLoadBalancerSettings) HasLbMetricSettings() bool`

HasLbMetricSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


