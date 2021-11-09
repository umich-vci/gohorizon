# SettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureSettings** | Pointer to [**FeatureSettingsUpdateSpec**](FeatureSettingsUpdateSpec.md) |  | [optional] 
**GeneralSettings** | Pointer to [**GeneralSettingsUpdateSpec**](GeneralSettingsUpdateSpec.md) |  | [optional] 
**SecuritySettings** | Pointer to [**SecuritySettingsUpdateSpec**](SecuritySettingsUpdateSpec.md) |  | [optional] 

## Methods

### NewSettingsUpdateSpec

`func NewSettingsUpdateSpec() *SettingsUpdateSpec`

NewSettingsUpdateSpec instantiates a new SettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsUpdateSpecWithDefaults

`func NewSettingsUpdateSpecWithDefaults() *SettingsUpdateSpec`

NewSettingsUpdateSpecWithDefaults instantiates a new SettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureSettings

`func (o *SettingsUpdateSpec) GetFeatureSettings() FeatureSettingsUpdateSpec`

GetFeatureSettings returns the FeatureSettings field if non-nil, zero value otherwise.

### GetFeatureSettingsOk

`func (o *SettingsUpdateSpec) GetFeatureSettingsOk() (*FeatureSettingsUpdateSpec, bool)`

GetFeatureSettingsOk returns a tuple with the FeatureSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSettings

`func (o *SettingsUpdateSpec) SetFeatureSettings(v FeatureSettingsUpdateSpec)`

SetFeatureSettings sets FeatureSettings field to given value.

### HasFeatureSettings

`func (o *SettingsUpdateSpec) HasFeatureSettings() bool`

HasFeatureSettings returns a boolean if a field has been set.

### GetGeneralSettings

`func (o *SettingsUpdateSpec) GetGeneralSettings() GeneralSettingsUpdateSpec`

GetGeneralSettings returns the GeneralSettings field if non-nil, zero value otherwise.

### GetGeneralSettingsOk

`func (o *SettingsUpdateSpec) GetGeneralSettingsOk() (*GeneralSettingsUpdateSpec, bool)`

GetGeneralSettingsOk returns a tuple with the GeneralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralSettings

`func (o *SettingsUpdateSpec) SetGeneralSettings(v GeneralSettingsUpdateSpec)`

SetGeneralSettings sets GeneralSettings field to given value.

### HasGeneralSettings

`func (o *SettingsUpdateSpec) HasGeneralSettings() bool`

HasGeneralSettings returns a boolean if a field has been set.

### GetSecuritySettings

`func (o *SettingsUpdateSpec) GetSecuritySettings() SecuritySettingsUpdateSpec`

GetSecuritySettings returns the SecuritySettings field if non-nil, zero value otherwise.

### GetSecuritySettingsOk

`func (o *SettingsUpdateSpec) GetSecuritySettingsOk() (*SecuritySettingsUpdateSpec, bool)`

GetSecuritySettingsOk returns a tuple with the SecuritySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySettings

`func (o *SettingsUpdateSpec) SetSecuritySettings(v SecuritySettingsUpdateSpec)`

SetSecuritySettings sets SecuritySettings field to given value.

### HasSecuritySettings

`func (o *SettingsUpdateSpec) HasSecuritySettings() bool`

HasSecuritySettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


