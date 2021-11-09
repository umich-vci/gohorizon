# FeatureSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudManaged** | Pointer to **bool** | Indicates whether this cluster/pod is managed by Horizon Cloud Services. This will be false only when there are no cloud managed machines. | [optional] 
**EnableHelpdesk** | Pointer to **bool** | Determines whether Helpdesk feature is enabled or not. By default Helpdesk is enabled. | [optional] 
**EnableImageManagement** | Pointer to **bool** | Determines whether Image Management feature is enabled or not. By default Image Management is disabled. | [optional] 

## Methods

### NewFeatureSettings

`func NewFeatureSettings() *FeatureSettings`

NewFeatureSettings instantiates a new FeatureSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSettingsWithDefaults

`func NewFeatureSettingsWithDefaults() *FeatureSettings`

NewFeatureSettingsWithDefaults instantiates a new FeatureSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudManaged

`func (o *FeatureSettings) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *FeatureSettings) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *FeatureSettings) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *FeatureSettings) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetEnableHelpdesk

`func (o *FeatureSettings) GetEnableHelpdesk() bool`

GetEnableHelpdesk returns the EnableHelpdesk field if non-nil, zero value otherwise.

### GetEnableHelpdeskOk

`func (o *FeatureSettings) GetEnableHelpdeskOk() (*bool, bool)`

GetEnableHelpdeskOk returns a tuple with the EnableHelpdesk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHelpdesk

`func (o *FeatureSettings) SetEnableHelpdesk(v bool)`

SetEnableHelpdesk sets EnableHelpdesk field to given value.

### HasEnableHelpdesk

`func (o *FeatureSettings) HasEnableHelpdesk() bool`

HasEnableHelpdesk returns a boolean if a field has been set.

### GetEnableImageManagement

`func (o *FeatureSettings) GetEnableImageManagement() bool`

GetEnableImageManagement returns the EnableImageManagement field if non-nil, zero value otherwise.

### GetEnableImageManagementOk

`func (o *FeatureSettings) GetEnableImageManagementOk() (*bool, bool)`

GetEnableImageManagementOk returns a tuple with the EnableImageManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImageManagement

`func (o *FeatureSettings) SetEnableImageManagement(v bool)`

SetEnableImageManagement sets EnableImageManagement field to given value.

### HasEnableImageManagement

`func (o *FeatureSettings) HasEnableImageManagement() bool`

HasEnableImageManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


