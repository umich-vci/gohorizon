# EnvironmentInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterGuid** | Pointer to **string** | The GUID of a group of Connection Servers sharing the same configuration. | [optional] 
**ClusterName** | Pointer to **string** | The name of a group of Connection Servers sharing the same configuration. | [optional] 
**DeploymentType** | Pointer to **string** | Indicates different environments that Horizon can be deployed into. * GENERAL: Horizon is deployed on On-premises. * AZURE: Horizon is deployed on Azure. * AWS: Horizon is deployed on AWS. * DELL_EMC: Horizon is deployed on Dell EMC. * GOOGLE: Horizon is deployed on Google Cloud. * ORACLE: Horizon is deployed on Oracle Cloud. | [optional] 
**FipsModeEnabled** | Pointer to **bool** | Indicates if FIPS mode is enabled. | [optional] 
**IpMode** | Pointer to **string** | Indicates the IP mode of the environment. * IPv4: The ip mode is IPv4. * IPv6: The ip mode is IPv6. | [optional] 
**LocalConnectionServerBuild** | Pointer to **string** | Local connection Server build number. | [optional] 
**LocalConnectionServerVersion** | Pointer to **string** | Local connection Server version number. | [optional] 
**LocalPodName** | Pointer to **string** | The name of the current pod in the Multi-DataCenter Horizon Pod, the value will be null when PodFederation is not initialized. | [optional] 
**TimezoneOffset** | Pointer to **int32** | Represents the clusters time zone offset from UTC in seconds. | [optional] 

## Methods

### NewEnvironmentInfoV2

`func NewEnvironmentInfoV2() *EnvironmentInfoV2`

NewEnvironmentInfoV2 instantiates a new EnvironmentInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentInfoV2WithDefaults

`func NewEnvironmentInfoV2WithDefaults() *EnvironmentInfoV2`

NewEnvironmentInfoV2WithDefaults instantiates a new EnvironmentInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterGuid

`func (o *EnvironmentInfoV2) GetClusterGuid() string`

GetClusterGuid returns the ClusterGuid field if non-nil, zero value otherwise.

### GetClusterGuidOk

`func (o *EnvironmentInfoV2) GetClusterGuidOk() (*string, bool)`

GetClusterGuidOk returns a tuple with the ClusterGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGuid

`func (o *EnvironmentInfoV2) SetClusterGuid(v string)`

SetClusterGuid sets ClusterGuid field to given value.

### HasClusterGuid

`func (o *EnvironmentInfoV2) HasClusterGuid() bool`

HasClusterGuid returns a boolean if a field has been set.

### GetClusterName

`func (o *EnvironmentInfoV2) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *EnvironmentInfoV2) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *EnvironmentInfoV2) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *EnvironmentInfoV2) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDeploymentType

`func (o *EnvironmentInfoV2) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *EnvironmentInfoV2) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *EnvironmentInfoV2) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *EnvironmentInfoV2) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetFipsModeEnabled

`func (o *EnvironmentInfoV2) GetFipsModeEnabled() bool`

GetFipsModeEnabled returns the FipsModeEnabled field if non-nil, zero value otherwise.

### GetFipsModeEnabledOk

`func (o *EnvironmentInfoV2) GetFipsModeEnabledOk() (*bool, bool)`

GetFipsModeEnabledOk returns a tuple with the FipsModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsModeEnabled

`func (o *EnvironmentInfoV2) SetFipsModeEnabled(v bool)`

SetFipsModeEnabled sets FipsModeEnabled field to given value.

### HasFipsModeEnabled

`func (o *EnvironmentInfoV2) HasFipsModeEnabled() bool`

HasFipsModeEnabled returns a boolean if a field has been set.

### GetIpMode

`func (o *EnvironmentInfoV2) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *EnvironmentInfoV2) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *EnvironmentInfoV2) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *EnvironmentInfoV2) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetLocalConnectionServerBuild

`func (o *EnvironmentInfoV2) GetLocalConnectionServerBuild() string`

GetLocalConnectionServerBuild returns the LocalConnectionServerBuild field if non-nil, zero value otherwise.

### GetLocalConnectionServerBuildOk

`func (o *EnvironmentInfoV2) GetLocalConnectionServerBuildOk() (*string, bool)`

GetLocalConnectionServerBuildOk returns a tuple with the LocalConnectionServerBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConnectionServerBuild

`func (o *EnvironmentInfoV2) SetLocalConnectionServerBuild(v string)`

SetLocalConnectionServerBuild sets LocalConnectionServerBuild field to given value.

### HasLocalConnectionServerBuild

`func (o *EnvironmentInfoV2) HasLocalConnectionServerBuild() bool`

HasLocalConnectionServerBuild returns a boolean if a field has been set.

### GetLocalConnectionServerVersion

`func (o *EnvironmentInfoV2) GetLocalConnectionServerVersion() string`

GetLocalConnectionServerVersion returns the LocalConnectionServerVersion field if non-nil, zero value otherwise.

### GetLocalConnectionServerVersionOk

`func (o *EnvironmentInfoV2) GetLocalConnectionServerVersionOk() (*string, bool)`

GetLocalConnectionServerVersionOk returns a tuple with the LocalConnectionServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConnectionServerVersion

`func (o *EnvironmentInfoV2) SetLocalConnectionServerVersion(v string)`

SetLocalConnectionServerVersion sets LocalConnectionServerVersion field to given value.

### HasLocalConnectionServerVersion

`func (o *EnvironmentInfoV2) HasLocalConnectionServerVersion() bool`

HasLocalConnectionServerVersion returns a boolean if a field has been set.

### GetLocalPodName

`func (o *EnvironmentInfoV2) GetLocalPodName() string`

GetLocalPodName returns the LocalPodName field if non-nil, zero value otherwise.

### GetLocalPodNameOk

`func (o *EnvironmentInfoV2) GetLocalPodNameOk() (*string, bool)`

GetLocalPodNameOk returns a tuple with the LocalPodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPodName

`func (o *EnvironmentInfoV2) SetLocalPodName(v string)`

SetLocalPodName sets LocalPodName field to given value.

### HasLocalPodName

`func (o *EnvironmentInfoV2) HasLocalPodName() bool`

HasLocalPodName returns a boolean if a field has been set.

### GetTimezoneOffset

`func (o *EnvironmentInfoV2) GetTimezoneOffset() int32`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *EnvironmentInfoV2) GetTimezoneOffsetOk() (*int32, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *EnvironmentInfoV2) SetTimezoneOffset(v int32)`

SetTimezoneOffset sets TimezoneOffset field to given value.

### HasTimezoneOffset

`func (o *EnvironmentInfoV2) HasTimezoneOffset() bool`

HasTimezoneOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


