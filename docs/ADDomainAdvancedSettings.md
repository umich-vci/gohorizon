# ADDomainAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainAutoDiscovery** | Pointer to **bool** | Auto discovers domain controllers. Auto discovery, AD domain controllers and preferred site name are mutually exclusive. Only one of them can be defined at a time. Default value is true. | [optional] 
**AdDomainContext** | **string** | Active directory domain Context. | 
**AdDomainControllers** | Pointer to **[]string** | One or more AD domain controllers. Auto discovery, AD domain controllers and preferred site name are mutually exclusive.  Only one of them can be defined at a time. | [optional] 
**AdDomainPreferredSite** | Pointer to **string** | ADDomain preferred domain site. Auto discovery, AD domain controllers and preferred site name are mutually exclusive. Only one of them can be defined at a time. | [optional] 
**Port** | **int32** | Port of the server to connect to. | 

## Methods

### NewADDomainAdvancedSettings

`func NewADDomainAdvancedSettings(adDomainContext string, port int32, ) *ADDomainAdvancedSettings`

NewADDomainAdvancedSettings instantiates a new ADDomainAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainAdvancedSettingsWithDefaults

`func NewADDomainAdvancedSettingsWithDefaults() *ADDomainAdvancedSettings`

NewADDomainAdvancedSettingsWithDefaults instantiates a new ADDomainAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainAutoDiscovery

`func (o *ADDomainAdvancedSettings) GetAdDomainAutoDiscovery() bool`

GetAdDomainAutoDiscovery returns the AdDomainAutoDiscovery field if non-nil, zero value otherwise.

### GetAdDomainAutoDiscoveryOk

`func (o *ADDomainAdvancedSettings) GetAdDomainAutoDiscoveryOk() (*bool, bool)`

GetAdDomainAutoDiscoveryOk returns a tuple with the AdDomainAutoDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainAutoDiscovery

`func (o *ADDomainAdvancedSettings) SetAdDomainAutoDiscovery(v bool)`

SetAdDomainAutoDiscovery sets AdDomainAutoDiscovery field to given value.

### HasAdDomainAutoDiscovery

`func (o *ADDomainAdvancedSettings) HasAdDomainAutoDiscovery() bool`

HasAdDomainAutoDiscovery returns a boolean if a field has been set.

### GetAdDomainContext

`func (o *ADDomainAdvancedSettings) GetAdDomainContext() string`

GetAdDomainContext returns the AdDomainContext field if non-nil, zero value otherwise.

### GetAdDomainContextOk

`func (o *ADDomainAdvancedSettings) GetAdDomainContextOk() (*string, bool)`

GetAdDomainContextOk returns a tuple with the AdDomainContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainContext

`func (o *ADDomainAdvancedSettings) SetAdDomainContext(v string)`

SetAdDomainContext sets AdDomainContext field to given value.


### GetAdDomainControllers

`func (o *ADDomainAdvancedSettings) GetAdDomainControllers() []string`

GetAdDomainControllers returns the AdDomainControllers field if non-nil, zero value otherwise.

### GetAdDomainControllersOk

`func (o *ADDomainAdvancedSettings) GetAdDomainControllersOk() (*[]string, bool)`

GetAdDomainControllersOk returns a tuple with the AdDomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainControllers

`func (o *ADDomainAdvancedSettings) SetAdDomainControllers(v []string)`

SetAdDomainControllers sets AdDomainControllers field to given value.

### HasAdDomainControllers

`func (o *ADDomainAdvancedSettings) HasAdDomainControllers() bool`

HasAdDomainControllers returns a boolean if a field has been set.

### GetAdDomainPreferredSite

`func (o *ADDomainAdvancedSettings) GetAdDomainPreferredSite() string`

GetAdDomainPreferredSite returns the AdDomainPreferredSite field if non-nil, zero value otherwise.

### GetAdDomainPreferredSiteOk

`func (o *ADDomainAdvancedSettings) GetAdDomainPreferredSiteOk() (*string, bool)`

GetAdDomainPreferredSiteOk returns a tuple with the AdDomainPreferredSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainPreferredSite

`func (o *ADDomainAdvancedSettings) SetAdDomainPreferredSite(v string)`

SetAdDomainPreferredSite sets AdDomainPreferredSite field to given value.

### HasAdDomainPreferredSite

`func (o *ADDomainAdvancedSettings) HasAdDomainPreferredSite() bool`

HasAdDomainPreferredSite returns a boolean if a field has been set.

### GetPort

`func (o *ADDomainAdvancedSettings) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ADDomainAdvancedSettings) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ADDomainAdvancedSettings) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


