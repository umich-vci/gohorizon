# FarmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Farm description. The maximum length is 1024 characters. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the Farm. The maximum length is 256 characters. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Farm is enabled for brokering. Default value is true. | [optional] 
**Id** | Pointer to **string** | Unique ID representing Farm. | [optional] 
**Name** | Pointer to **string** | Name of the Farm. The maximum length is 64 characters. | [optional] 
**Settings** | Pointer to [**FarmSettings**](FarmSettings.md) |  | [optional] 
**Source** | Pointer to **string** | Type of the Farm. * INSTANT_CLONE: The Farm uses instant clone technology for provisioning the RDS Servers.Applicable for AUTOMATED type Farms only. * LINKED_CLONE: The Farm uses linked clone technology for provisioning the RDS Servers.Applicable for AUTOMATED type Farms only. | [optional] 
**Type** | Pointer to **string** | Type of the Farm. * AUTOMATED: Automated Farm. * MANUAL: Manual Farm. | [optional] 

## Methods

### NewFarmInfo

`func NewFarmInfo() *FarmInfo`

NewFarmInfo instantiates a new FarmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmInfoWithDefaults

`func NewFarmInfoWithDefaults() *FarmInfo`

NewFarmInfoWithDefaults instantiates a new FarmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FarmInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FarmInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FarmInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FarmInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FarmInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FarmInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FarmInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FarmInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *FarmInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FarmInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FarmInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FarmInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *FarmInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FarmInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FarmInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FarmInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FarmInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FarmInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FarmInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FarmInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *FarmInfo) GetSettings() FarmSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *FarmInfo) GetSettingsOk() (*FarmSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *FarmInfo) SetSettings(v FarmSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *FarmInfo) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSource

`func (o *FarmInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FarmInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FarmInfo) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FarmInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *FarmInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FarmInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FarmInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FarmInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


