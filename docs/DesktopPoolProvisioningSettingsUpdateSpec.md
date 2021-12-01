# DesktopPoolProvisioningSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostOrClusterId** | **string** | Host or cluster where the machines are deployed in. | 
**ImStreamId** | Pointer to **string** | Applicable To: Full clone desktop pools.&lt;br&gt; This is required when vm_template_id is not set. | [optional] 
**ImTagId** | Pointer to **string** | Applicable To: Full clone desktop pools.&lt;br&gt; This is required when im_stream_id is set. | [optional] 
**ResourcePoolId** | **string** | Resource pool to deploy the machines. | 
**VmTemplateId** | Pointer to **string** | Applicable To: Full clone desktop pools.&lt;br&gt;Template from which full clone machines are deployed. | [optional] 

## Methods

### NewDesktopPoolProvisioningSettingsUpdateSpec

`func NewDesktopPoolProvisioningSettingsUpdateSpec(hostOrClusterId string, resourcePoolId string, ) *DesktopPoolProvisioningSettingsUpdateSpec`

NewDesktopPoolProvisioningSettingsUpdateSpec instantiates a new DesktopPoolProvisioningSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolProvisioningSettingsUpdateSpecWithDefaults

`func NewDesktopPoolProvisioningSettingsUpdateSpecWithDefaults() *DesktopPoolProvisioningSettingsUpdateSpec`

NewDesktopPoolProvisioningSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolProvisioningSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostOrClusterId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetHostOrClusterId() string`

GetHostOrClusterId returns the HostOrClusterId field if non-nil, zero value otherwise.

### GetHostOrClusterIdOk

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetHostOrClusterIdOk() (*string, bool)`

GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrClusterId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) SetHostOrClusterId(v string)`

SetHostOrClusterId sets HostOrClusterId field to given value.


### GetImStreamId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.

### HasImStreamId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) HasImStreamId() bool`

HasImStreamId returns a boolean if a field has been set.

### GetImTagId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetImTagId() string`

GetImTagId returns the ImTagId field if non-nil, zero value otherwise.

### GetImTagIdOk

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetImTagIdOk() (*string, bool)`

GetImTagIdOk returns a tuple with the ImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImTagId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) SetImTagId(v string)`

SetImTagId sets ImTagId field to given value.

### HasImTagId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) HasImTagId() bool`

HasImTagId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.


### GetVmTemplateId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.

### HasVmTemplateId

`func (o *DesktopPoolProvisioningSettingsUpdateSpec) HasVmTemplateId() bool`

HasVmTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


