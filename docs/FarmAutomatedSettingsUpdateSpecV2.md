# FarmAutomatedSettingsUpdateSpecV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomizationSettings** | [**FarmCustomizationSettingsUpdateSpecV2**](FarmCustomizationSettingsUpdateSpecV2.md) |  | 
**EnableProvisioning** | **bool** | Indicates whether to enable provisioning immediately. | 
**MaxSessionType** | **string** | RDS Server type for max sessions. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions. | 
**MaxSessions** | Pointer to **int32** | Maximum number of sessions allowed for RDS Server. This is required if max_session_type is set to LIMITED. | [optional] 
**MinReadyVms** | **int32** | Minimum number of ready (provisioned) RDS Servers during Instant clone maintenance operations. Use this setting to perform machine maintenance operations in a rolling fashion. Increasing this count may decrease the concurrency for Instant clone maintenance operations for the automated farm.  | 
**Nics** | Pointer to [**[]FarmNetworkInterfaceCardSettingsUpdateSpec**](FarmNetworkInterfaceCardSettingsUpdateSpec.md) | Network interface card settings for RDS Servers provisioned for this farm. A NIC may appear at most once in these settings and must be present on this RDS Server&#39;s parent&#39;s snapshot. Not all NICs need be configured. Any that are not will use default settings. | [optional] 
**PatternNamingSettings** | [**FarmRDSServersPatternNamingSettingsUpdateSpec**](FarmRDSServersPatternNamingSettingsUpdateSpec.md) |  | 
**ProvisioningSettings** | [**FarmProvisioningSettingsUpdateSpec**](FarmProvisioningSettingsUpdateSpec.md) |  | 
**StopProvisioningOnError** | **bool** | Indicates whether provisioning on all VMs stops on error. | 
**StorageSettings** | [**FarmStorageSettingsUpdateSpec**](FarmStorageSettingsUpdateSpec.md) |  | 
**TransparentPageSharingScope** | **string** | Transparent page sharing scope for the farm. * VM: Inter-VM page sharing is not permitted. * FARM: Inter-VM page sharing among VMs belonging to the same automated farm is permitted. * POD: Inter-VM page sharing among VMs belonging to the same Pod is permitted. * GLOBAL: Inter-VM page sharing among all VMs on the same host is permitted. | 

## Methods

### NewFarmAutomatedSettingsUpdateSpecV2

`func NewFarmAutomatedSettingsUpdateSpecV2(customizationSettings FarmCustomizationSettingsUpdateSpecV2, enableProvisioning bool, maxSessionType string, minReadyVms int32, patternNamingSettings FarmRDSServersPatternNamingSettingsUpdateSpec, provisioningSettings FarmProvisioningSettingsUpdateSpec, stopProvisioningOnError bool, storageSettings FarmStorageSettingsUpdateSpec, transparentPageSharingScope string, ) *FarmAutomatedSettingsUpdateSpecV2`

NewFarmAutomatedSettingsUpdateSpecV2 instantiates a new FarmAutomatedSettingsUpdateSpecV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmAutomatedSettingsUpdateSpecV2WithDefaults

`func NewFarmAutomatedSettingsUpdateSpecV2WithDefaults() *FarmAutomatedSettingsUpdateSpecV2`

NewFarmAutomatedSettingsUpdateSpecV2WithDefaults instantiates a new FarmAutomatedSettingsUpdateSpecV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomizationSettings

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetCustomizationSettings() FarmCustomizationSettingsUpdateSpecV2`

GetCustomizationSettings returns the CustomizationSettings field if non-nil, zero value otherwise.

### GetCustomizationSettingsOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetCustomizationSettingsOk() (*FarmCustomizationSettingsUpdateSpecV2, bool)`

GetCustomizationSettingsOk returns a tuple with the CustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationSettings

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetCustomizationSettings(v FarmCustomizationSettingsUpdateSpecV2)`

SetCustomizationSettings sets CustomizationSettings field to given value.


### GetEnableProvisioning

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetEnableProvisioning() bool`

GetEnableProvisioning returns the EnableProvisioning field if non-nil, zero value otherwise.

### GetEnableProvisioningOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetEnableProvisioningOk() (*bool, bool)`

GetEnableProvisioningOk returns a tuple with the EnableProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProvisioning

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetEnableProvisioning(v bool)`

SetEnableProvisioning sets EnableProvisioning field to given value.


### GetMaxSessionType

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetMaxSessionType() string`

GetMaxSessionType returns the MaxSessionType field if non-nil, zero value otherwise.

### GetMaxSessionTypeOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetMaxSessionTypeOk() (*string, bool)`

GetMaxSessionTypeOk returns a tuple with the MaxSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionType

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetMaxSessionType(v string)`

SetMaxSessionType sets MaxSessionType field to given value.


### GetMaxSessions

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetMaxSessions() int32`

GetMaxSessions returns the MaxSessions field if non-nil, zero value otherwise.

### GetMaxSessionsOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetMaxSessionsOk() (*int32, bool)`

GetMaxSessionsOk returns a tuple with the MaxSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessions

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetMaxSessions(v int32)`

SetMaxSessions sets MaxSessions field to given value.

### HasMaxSessions

`func (o *FarmAutomatedSettingsUpdateSpecV2) HasMaxSessions() bool`

HasMaxSessions returns a boolean if a field has been set.

### GetMinReadyVms

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetMinReadyVms() int32`

GetMinReadyVms returns the MinReadyVms field if non-nil, zero value otherwise.

### GetMinReadyVmsOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetMinReadyVmsOk() (*int32, bool)`

GetMinReadyVmsOk returns a tuple with the MinReadyVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadyVms

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetMinReadyVms(v int32)`

SetMinReadyVms sets MinReadyVms field to given value.


### GetNics

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetNics() []FarmNetworkInterfaceCardSettingsUpdateSpec`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetNicsOk() (*[]FarmNetworkInterfaceCardSettingsUpdateSpec, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetNics(v []FarmNetworkInterfaceCardSettingsUpdateSpec)`

SetNics sets Nics field to given value.

### HasNics

`func (o *FarmAutomatedSettingsUpdateSpecV2) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetPatternNamingSettings

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetPatternNamingSettings() FarmRDSServersPatternNamingSettingsUpdateSpec`

GetPatternNamingSettings returns the PatternNamingSettings field if non-nil, zero value otherwise.

### GetPatternNamingSettingsOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetPatternNamingSettingsOk() (*FarmRDSServersPatternNamingSettingsUpdateSpec, bool)`

GetPatternNamingSettingsOk returns a tuple with the PatternNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternNamingSettings

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetPatternNamingSettings(v FarmRDSServersPatternNamingSettingsUpdateSpec)`

SetPatternNamingSettings sets PatternNamingSettings field to given value.


### GetProvisioningSettings

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetProvisioningSettings() FarmProvisioningSettingsUpdateSpec`

GetProvisioningSettings returns the ProvisioningSettings field if non-nil, zero value otherwise.

### GetProvisioningSettingsOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetProvisioningSettingsOk() (*FarmProvisioningSettingsUpdateSpec, bool)`

GetProvisioningSettingsOk returns a tuple with the ProvisioningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSettings

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetProvisioningSettings(v FarmProvisioningSettingsUpdateSpec)`

SetProvisioningSettings sets ProvisioningSettings field to given value.


### GetStopProvisioningOnError

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetStopProvisioningOnError() bool`

GetStopProvisioningOnError returns the StopProvisioningOnError field if non-nil, zero value otherwise.

### GetStopProvisioningOnErrorOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetStopProvisioningOnErrorOk() (*bool, bool)`

GetStopProvisioningOnErrorOk returns a tuple with the StopProvisioningOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProvisioningOnError

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetStopProvisioningOnError(v bool)`

SetStopProvisioningOnError sets StopProvisioningOnError field to given value.


### GetStorageSettings

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetStorageSettings() FarmStorageSettingsUpdateSpec`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetStorageSettingsOk() (*FarmStorageSettingsUpdateSpec, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetStorageSettings(v FarmStorageSettingsUpdateSpec)`

SetStorageSettings sets StorageSettings field to given value.


### GetTransparentPageSharingScope

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetTransparentPageSharingScope() string`

GetTransparentPageSharingScope returns the TransparentPageSharingScope field if non-nil, zero value otherwise.

### GetTransparentPageSharingScopeOk

`func (o *FarmAutomatedSettingsUpdateSpecV2) GetTransparentPageSharingScopeOk() (*string, bool)`

GetTransparentPageSharingScopeOk returns a tuple with the TransparentPageSharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentPageSharingScope

`func (o *FarmAutomatedSettingsUpdateSpecV2) SetTransparentPageSharingScope(v string)`

SetTransparentPageSharingScope sets TransparentPageSharingScope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


