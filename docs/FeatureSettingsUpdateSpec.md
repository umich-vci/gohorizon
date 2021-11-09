# FeatureSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudManaged** | Pointer to **bool** | Indicates whether this cluster/pod is managed by Horizon Cloud Services. This will be false only when there are no cloud managed machines. | [optional] 

## Methods

### NewFeatureSettingsUpdateSpec

`func NewFeatureSettingsUpdateSpec() *FeatureSettingsUpdateSpec`

NewFeatureSettingsUpdateSpec instantiates a new FeatureSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSettingsUpdateSpecWithDefaults

`func NewFeatureSettingsUpdateSpecWithDefaults() *FeatureSettingsUpdateSpec`

NewFeatureSettingsUpdateSpecWithDefaults instantiates a new FeatureSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudManaged

`func (o *FeatureSettingsUpdateSpec) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *FeatureSettingsUpdateSpec) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *FeatureSettingsUpdateSpec) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *FeatureSettingsUpdateSpec) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


