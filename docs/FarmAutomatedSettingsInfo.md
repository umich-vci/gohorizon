# FarmAutomatedSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomizationSettings** | Pointer to [**FarmCustomizationSettingsInfo**](FarmCustomizationSettingsInfo.md) |  | [optional] 
**EnableProvisioning** | Pointer to **bool** | Indicates whether to enable provisioning immediately.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ImageSource** | Pointer to **string** | Source of image used in the farm.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * VIRTUAL_CENTER: Image was created in virtual center. * IMAGE_CATALOG: Image was created in image catalog. | [optional] 
**MaxSessionType** | Pointer to **string** | RDS Server type for max sessions. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions. | [optional] 
**MaxSessions** | Pointer to **int32** | Maximum number of sessions allowed for RDS Server. This is set when max_session_type is LIMITED. | [optional] 
**MinReadyVms** | Pointer to **int32** | Minimum number of ready (provisioned) RDS Servers during Instant clone maintenance operations. Use this setting to perform machine maintenance operations in a rolling fashion. Increasing this count may decrease the concurrency for Instant clone maintenance operations for the automated farm. | [optional] 
**Nics** | Pointer to [**[]FarmNetworkInterfaceCardSettingsInfo**](FarmNetworkInterfaceCardSettingsInfo.md) | Network interface card settings for RDS Servers provisioned for this farm. A NIC may appear at most once in these settings and must be present on this RDS Server&#39;s parent&#39;s snapshot. Not all NICs need be configured. Any that are not will use default settings. | [optional] 
**OperatingSystem** | Pointer to **string** | The guest operating system. * UNKNOWN: Unknown * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_SERVER_OTHER: Linux Server (other) | [optional] 
**OperatingSystemArchitecture** | Pointer to **string** | The guest operating system architecture. * UNKNOWN: Operating System cannot be determined. * BIT_32: 32 bit Operating System Architecture. * BIT_64: 64 bit Operating System Architecture. | [optional] 
**PatternNamingSettings** | Pointer to [**FarmRDSServersPatternNamingSettingsInfo**](FarmRDSServersPatternNamingSettingsInfo.md) |  | [optional] 
**ProvisioningSettings** | Pointer to [**FarmProvisioningSettingsInfo**](FarmProvisioningSettingsInfo.md) |  | [optional] 
**ProvisioningStatusData** | Pointer to [**FarmProvisioningStatusInfo**](FarmProvisioningStatusInfo.md) |  | [optional] 
**StopProvisioningOnError** | Pointer to **bool** | Indicates whether provisioning on all VMs stops on error. | [optional] 
**StorageSettings** | Pointer to [**FarmStorageSettingsInfo**](FarmStorageSettingsInfo.md) |  | [optional] 
**TransparentPageSharingScope** | Pointer to **string** | Transparent page sharing scope for the farm. * VM: Inter-VM page sharing is not permitted. * FARM: Inter-VM page sharing among VMs belonging to the same automated farm is permitted. * POD: Inter-VM page sharing among VMs belonging to the same Pod is permitted. * GLOBAL: Inter-VM page sharing among all VMs on the same host is permitted. | [optional] 
**VcenterId** | Pointer to **string** | ID of the virtual center server.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 

## Methods

### NewFarmAutomatedSettingsInfo

`func NewFarmAutomatedSettingsInfo() *FarmAutomatedSettingsInfo`

NewFarmAutomatedSettingsInfo instantiates a new FarmAutomatedSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmAutomatedSettingsInfoWithDefaults

`func NewFarmAutomatedSettingsInfoWithDefaults() *FarmAutomatedSettingsInfo`

NewFarmAutomatedSettingsInfoWithDefaults instantiates a new FarmAutomatedSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomizationSettings

`func (o *FarmAutomatedSettingsInfo) GetCustomizationSettings() FarmCustomizationSettingsInfo`

GetCustomizationSettings returns the CustomizationSettings field if non-nil, zero value otherwise.

### GetCustomizationSettingsOk

`func (o *FarmAutomatedSettingsInfo) GetCustomizationSettingsOk() (*FarmCustomizationSettingsInfo, bool)`

GetCustomizationSettingsOk returns a tuple with the CustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationSettings

`func (o *FarmAutomatedSettingsInfo) SetCustomizationSettings(v FarmCustomizationSettingsInfo)`

SetCustomizationSettings sets CustomizationSettings field to given value.

### HasCustomizationSettings

`func (o *FarmAutomatedSettingsInfo) HasCustomizationSettings() bool`

HasCustomizationSettings returns a boolean if a field has been set.

### GetEnableProvisioning

`func (o *FarmAutomatedSettingsInfo) GetEnableProvisioning() bool`

GetEnableProvisioning returns the EnableProvisioning field if non-nil, zero value otherwise.

### GetEnableProvisioningOk

`func (o *FarmAutomatedSettingsInfo) GetEnableProvisioningOk() (*bool, bool)`

GetEnableProvisioningOk returns a tuple with the EnableProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProvisioning

`func (o *FarmAutomatedSettingsInfo) SetEnableProvisioning(v bool)`

SetEnableProvisioning sets EnableProvisioning field to given value.

### HasEnableProvisioning

`func (o *FarmAutomatedSettingsInfo) HasEnableProvisioning() bool`

HasEnableProvisioning returns a boolean if a field has been set.

### GetImageSource

`func (o *FarmAutomatedSettingsInfo) GetImageSource() string`

GetImageSource returns the ImageSource field if non-nil, zero value otherwise.

### GetImageSourceOk

`func (o *FarmAutomatedSettingsInfo) GetImageSourceOk() (*string, bool)`

GetImageSourceOk returns a tuple with the ImageSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSource

`func (o *FarmAutomatedSettingsInfo) SetImageSource(v string)`

SetImageSource sets ImageSource field to given value.

### HasImageSource

`func (o *FarmAutomatedSettingsInfo) HasImageSource() bool`

HasImageSource returns a boolean if a field has been set.

### GetMaxSessionType

`func (o *FarmAutomatedSettingsInfo) GetMaxSessionType() string`

GetMaxSessionType returns the MaxSessionType field if non-nil, zero value otherwise.

### GetMaxSessionTypeOk

`func (o *FarmAutomatedSettingsInfo) GetMaxSessionTypeOk() (*string, bool)`

GetMaxSessionTypeOk returns a tuple with the MaxSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionType

`func (o *FarmAutomatedSettingsInfo) SetMaxSessionType(v string)`

SetMaxSessionType sets MaxSessionType field to given value.

### HasMaxSessionType

`func (o *FarmAutomatedSettingsInfo) HasMaxSessionType() bool`

HasMaxSessionType returns a boolean if a field has been set.

### GetMaxSessions

`func (o *FarmAutomatedSettingsInfo) GetMaxSessions() int32`

GetMaxSessions returns the MaxSessions field if non-nil, zero value otherwise.

### GetMaxSessionsOk

`func (o *FarmAutomatedSettingsInfo) GetMaxSessionsOk() (*int32, bool)`

GetMaxSessionsOk returns a tuple with the MaxSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessions

`func (o *FarmAutomatedSettingsInfo) SetMaxSessions(v int32)`

SetMaxSessions sets MaxSessions field to given value.

### HasMaxSessions

`func (o *FarmAutomatedSettingsInfo) HasMaxSessions() bool`

HasMaxSessions returns a boolean if a field has been set.

### GetMinReadyVms

`func (o *FarmAutomatedSettingsInfo) GetMinReadyVms() int32`

GetMinReadyVms returns the MinReadyVms field if non-nil, zero value otherwise.

### GetMinReadyVmsOk

`func (o *FarmAutomatedSettingsInfo) GetMinReadyVmsOk() (*int32, bool)`

GetMinReadyVmsOk returns a tuple with the MinReadyVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadyVms

`func (o *FarmAutomatedSettingsInfo) SetMinReadyVms(v int32)`

SetMinReadyVms sets MinReadyVms field to given value.

### HasMinReadyVms

`func (o *FarmAutomatedSettingsInfo) HasMinReadyVms() bool`

HasMinReadyVms returns a boolean if a field has been set.

### GetNics

`func (o *FarmAutomatedSettingsInfo) GetNics() []FarmNetworkInterfaceCardSettingsInfo`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *FarmAutomatedSettingsInfo) GetNicsOk() (*[]FarmNetworkInterfaceCardSettingsInfo, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *FarmAutomatedSettingsInfo) SetNics(v []FarmNetworkInterfaceCardSettingsInfo)`

SetNics sets Nics field to given value.

### HasNics

`func (o *FarmAutomatedSettingsInfo) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *FarmAutomatedSettingsInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *FarmAutomatedSettingsInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *FarmAutomatedSettingsInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *FarmAutomatedSettingsInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperatingSystemArchitecture

`func (o *FarmAutomatedSettingsInfo) GetOperatingSystemArchitecture() string`

GetOperatingSystemArchitecture returns the OperatingSystemArchitecture field if non-nil, zero value otherwise.

### GetOperatingSystemArchitectureOk

`func (o *FarmAutomatedSettingsInfo) GetOperatingSystemArchitectureOk() (*string, bool)`

GetOperatingSystemArchitectureOk returns a tuple with the OperatingSystemArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemArchitecture

`func (o *FarmAutomatedSettingsInfo) SetOperatingSystemArchitecture(v string)`

SetOperatingSystemArchitecture sets OperatingSystemArchitecture field to given value.

### HasOperatingSystemArchitecture

`func (o *FarmAutomatedSettingsInfo) HasOperatingSystemArchitecture() bool`

HasOperatingSystemArchitecture returns a boolean if a field has been set.

### GetPatternNamingSettings

`func (o *FarmAutomatedSettingsInfo) GetPatternNamingSettings() FarmRDSServersPatternNamingSettingsInfo`

GetPatternNamingSettings returns the PatternNamingSettings field if non-nil, zero value otherwise.

### GetPatternNamingSettingsOk

`func (o *FarmAutomatedSettingsInfo) GetPatternNamingSettingsOk() (*FarmRDSServersPatternNamingSettingsInfo, bool)`

GetPatternNamingSettingsOk returns a tuple with the PatternNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternNamingSettings

`func (o *FarmAutomatedSettingsInfo) SetPatternNamingSettings(v FarmRDSServersPatternNamingSettingsInfo)`

SetPatternNamingSettings sets PatternNamingSettings field to given value.

### HasPatternNamingSettings

`func (o *FarmAutomatedSettingsInfo) HasPatternNamingSettings() bool`

HasPatternNamingSettings returns a boolean if a field has been set.

### GetProvisioningSettings

`func (o *FarmAutomatedSettingsInfo) GetProvisioningSettings() FarmProvisioningSettingsInfo`

GetProvisioningSettings returns the ProvisioningSettings field if non-nil, zero value otherwise.

### GetProvisioningSettingsOk

`func (o *FarmAutomatedSettingsInfo) GetProvisioningSettingsOk() (*FarmProvisioningSettingsInfo, bool)`

GetProvisioningSettingsOk returns a tuple with the ProvisioningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSettings

`func (o *FarmAutomatedSettingsInfo) SetProvisioningSettings(v FarmProvisioningSettingsInfo)`

SetProvisioningSettings sets ProvisioningSettings field to given value.

### HasProvisioningSettings

`func (o *FarmAutomatedSettingsInfo) HasProvisioningSettings() bool`

HasProvisioningSettings returns a boolean if a field has been set.

### GetProvisioningStatusData

`func (o *FarmAutomatedSettingsInfo) GetProvisioningStatusData() FarmProvisioningStatusInfo`

GetProvisioningStatusData returns the ProvisioningStatusData field if non-nil, zero value otherwise.

### GetProvisioningStatusDataOk

`func (o *FarmAutomatedSettingsInfo) GetProvisioningStatusDataOk() (*FarmProvisioningStatusInfo, bool)`

GetProvisioningStatusDataOk returns a tuple with the ProvisioningStatusData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatusData

`func (o *FarmAutomatedSettingsInfo) SetProvisioningStatusData(v FarmProvisioningStatusInfo)`

SetProvisioningStatusData sets ProvisioningStatusData field to given value.

### HasProvisioningStatusData

`func (o *FarmAutomatedSettingsInfo) HasProvisioningStatusData() bool`

HasProvisioningStatusData returns a boolean if a field has been set.

### GetStopProvisioningOnError

`func (o *FarmAutomatedSettingsInfo) GetStopProvisioningOnError() bool`

GetStopProvisioningOnError returns the StopProvisioningOnError field if non-nil, zero value otherwise.

### GetStopProvisioningOnErrorOk

`func (o *FarmAutomatedSettingsInfo) GetStopProvisioningOnErrorOk() (*bool, bool)`

GetStopProvisioningOnErrorOk returns a tuple with the StopProvisioningOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProvisioningOnError

`func (o *FarmAutomatedSettingsInfo) SetStopProvisioningOnError(v bool)`

SetStopProvisioningOnError sets StopProvisioningOnError field to given value.

### HasStopProvisioningOnError

`func (o *FarmAutomatedSettingsInfo) HasStopProvisioningOnError() bool`

HasStopProvisioningOnError returns a boolean if a field has been set.

### GetStorageSettings

`func (o *FarmAutomatedSettingsInfo) GetStorageSettings() FarmStorageSettingsInfo`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *FarmAutomatedSettingsInfo) GetStorageSettingsOk() (*FarmStorageSettingsInfo, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *FarmAutomatedSettingsInfo) SetStorageSettings(v FarmStorageSettingsInfo)`

SetStorageSettings sets StorageSettings field to given value.

### HasStorageSettings

`func (o *FarmAutomatedSettingsInfo) HasStorageSettings() bool`

HasStorageSettings returns a boolean if a field has been set.

### GetTransparentPageSharingScope

`func (o *FarmAutomatedSettingsInfo) GetTransparentPageSharingScope() string`

GetTransparentPageSharingScope returns the TransparentPageSharingScope field if non-nil, zero value otherwise.

### GetTransparentPageSharingScopeOk

`func (o *FarmAutomatedSettingsInfo) GetTransparentPageSharingScopeOk() (*string, bool)`

GetTransparentPageSharingScopeOk returns a tuple with the TransparentPageSharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentPageSharingScope

`func (o *FarmAutomatedSettingsInfo) SetTransparentPageSharingScope(v string)`

SetTransparentPageSharingScope sets TransparentPageSharingScope field to given value.

### HasTransparentPageSharingScope

`func (o *FarmAutomatedSettingsInfo) HasTransparentPageSharingScope() bool`

HasTransparentPageSharingScope returns a boolean if a field has been set.

### GetVcenterId

`func (o *FarmAutomatedSettingsInfo) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *FarmAutomatedSettingsInfo) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *FarmAutomatedSettingsInfo) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *FarmAutomatedSettingsInfo) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


