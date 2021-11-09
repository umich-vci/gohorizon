# HomeSiteCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdUserOrGroupId** | **string** | SID of the user or group for whom the home site is to be configured. | 
**GlobalApplicationEntitlementId** | Pointer to **string** | ID of the Global Application Entitlement for which this site is the overriding home site. | [optional] 
**GlobalDesktopEntitlementId** | Pointer to **string** | ID of the Global Desktop Entitlement for which this site is the overriding home site. | [optional] 
**SiteId** | **string** | ID of the site for this home site configuration. | 

## Methods

### NewHomeSiteCreateSpec

`func NewHomeSiteCreateSpec(adUserOrGroupId string, siteId string, ) *HomeSiteCreateSpec`

NewHomeSiteCreateSpec instantiates a new HomeSiteCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHomeSiteCreateSpecWithDefaults

`func NewHomeSiteCreateSpecWithDefaults() *HomeSiteCreateSpec`

NewHomeSiteCreateSpecWithDefaults instantiates a new HomeSiteCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdUserOrGroupId

`func (o *HomeSiteCreateSpec) GetAdUserOrGroupId() string`

GetAdUserOrGroupId returns the AdUserOrGroupId field if non-nil, zero value otherwise.

### GetAdUserOrGroupIdOk

`func (o *HomeSiteCreateSpec) GetAdUserOrGroupIdOk() (*string, bool)`

GetAdUserOrGroupIdOk returns a tuple with the AdUserOrGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserOrGroupId

`func (o *HomeSiteCreateSpec) SetAdUserOrGroupId(v string)`

SetAdUserOrGroupId sets AdUserOrGroupId field to given value.


### GetGlobalApplicationEntitlementId

`func (o *HomeSiteCreateSpec) GetGlobalApplicationEntitlementId() string`

GetGlobalApplicationEntitlementId returns the GlobalApplicationEntitlementId field if non-nil, zero value otherwise.

### GetGlobalApplicationEntitlementIdOk

`func (o *HomeSiteCreateSpec) GetGlobalApplicationEntitlementIdOk() (*string, bool)`

GetGlobalApplicationEntitlementIdOk returns a tuple with the GlobalApplicationEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApplicationEntitlementId

`func (o *HomeSiteCreateSpec) SetGlobalApplicationEntitlementId(v string)`

SetGlobalApplicationEntitlementId sets GlobalApplicationEntitlementId field to given value.

### HasGlobalApplicationEntitlementId

`func (o *HomeSiteCreateSpec) HasGlobalApplicationEntitlementId() bool`

HasGlobalApplicationEntitlementId returns a boolean if a field has been set.

### GetGlobalDesktopEntitlementId

`func (o *HomeSiteCreateSpec) GetGlobalDesktopEntitlementId() string`

GetGlobalDesktopEntitlementId returns the GlobalDesktopEntitlementId field if non-nil, zero value otherwise.

### GetGlobalDesktopEntitlementIdOk

`func (o *HomeSiteCreateSpec) GetGlobalDesktopEntitlementIdOk() (*string, bool)`

GetGlobalDesktopEntitlementIdOk returns a tuple with the GlobalDesktopEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDesktopEntitlementId

`func (o *HomeSiteCreateSpec) SetGlobalDesktopEntitlementId(v string)`

SetGlobalDesktopEntitlementId sets GlobalDesktopEntitlementId field to given value.

### HasGlobalDesktopEntitlementId

`func (o *HomeSiteCreateSpec) HasGlobalDesktopEntitlementId() bool`

HasGlobalDesktopEntitlementId returns a boolean if a field has been set.

### GetSiteId

`func (o *HomeSiteCreateSpec) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *HomeSiteCreateSpec) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *HomeSiteCreateSpec) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


