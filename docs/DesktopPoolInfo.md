# DesktopPoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the Desktop Pool. The maximum length is 1024 characters. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the Desktop Pool. The maximum length is 256 characters. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Desktop Pool is enabled for brokering. | [optional] 
**Id** | Pointer to **string** | Unique ID representing Desktop Pool. | [optional] 
**Name** | Pointer to **string** | Name of the Desktop Pool. The maximum length is 64 characters. | [optional] 
**Settings** | Pointer to [**DesktopPoolSettings**](DesktopPoolSettings.md) |  | [optional] 
**Source** | Pointer to **string** | Source of the Machines in this Desktop Pool. * INSTANT_CLONE: The Desktop Pool uses instant clone technology for provisioning the machines. Applicable for AUTOMATED type desktop pools. * LINKED_CLONE: The Desktop Pool uses linked clone technology for provisioning the machines. Applicable for AUTOMATED type desktop pools. * VIRTUAL_CENTER: The Desktop Pool uses Virtual Center as source for provisioning the machines. Applicable for AUTOMATED and MANUAL type desktop pools. * RDS: The Desktop Pool is backed by Farm. The Farm used in this Desktop Pool can be of any Source. * UNMANAGED: The Desktop Pool holds the non-vCenter source machines that includes physical computers, blade PCs and non-vCenter servers. Applicable for MANUAL type desktop pools. | [optional] 
**Type** | Pointer to **string** | Type of the Desktop Pool. * AUTOMATED: Automated Desktop Pool. * MANUAL: Manual Desktop Pool. * RDS: RDS Desktop Pool. | [optional] 

## Methods

### NewDesktopPoolInfo

`func NewDesktopPoolInfo() *DesktopPoolInfo`

NewDesktopPoolInfo instantiates a new DesktopPoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolInfoWithDefaults

`func NewDesktopPoolInfoWithDefaults() *DesktopPoolInfo`

NewDesktopPoolInfoWithDefaults instantiates a new DesktopPoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DesktopPoolInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DesktopPoolInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DesktopPoolInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DesktopPoolInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DesktopPoolInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DesktopPoolInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DesktopPoolInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DesktopPoolInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *DesktopPoolInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DesktopPoolInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DesktopPoolInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DesktopPoolInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DesktopPoolInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DesktopPoolInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DesktopPoolInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DesktopPoolInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DesktopPoolInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DesktopPoolInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DesktopPoolInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DesktopPoolInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *DesktopPoolInfo) GetSettings() DesktopPoolSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DesktopPoolInfo) GetSettingsOk() (*DesktopPoolSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DesktopPoolInfo) SetSettings(v DesktopPoolSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DesktopPoolInfo) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSource

`func (o *DesktopPoolInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DesktopPoolInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DesktopPoolInfo) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DesktopPoolInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *DesktopPoolInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DesktopPoolInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DesktopPoolInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DesktopPoolInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


