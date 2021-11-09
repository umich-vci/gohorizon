# HomeSiteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdUserOrGroupId** | Pointer to **string** | SID of the user or group for whom this is the home site.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**GlobalApplicationEntitlementId** | Pointer to **string** | ID of the Global Application Entitlement for which this site is the overriding home site.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**GlobalDesktopEntitlementId** | Pointer to **string** | ID of the Global Desktop Entitlement for which this site is the overriding home site.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this home site assignment.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**SiteId** | Pointer to **string** | ID representing this home site.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 

## Methods

### NewHomeSiteInfo

`func NewHomeSiteInfo() *HomeSiteInfo`

NewHomeSiteInfo instantiates a new HomeSiteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHomeSiteInfoWithDefaults

`func NewHomeSiteInfoWithDefaults() *HomeSiteInfo`

NewHomeSiteInfoWithDefaults instantiates a new HomeSiteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdUserOrGroupId

`func (o *HomeSiteInfo) GetAdUserOrGroupId() string`

GetAdUserOrGroupId returns the AdUserOrGroupId field if non-nil, zero value otherwise.

### GetAdUserOrGroupIdOk

`func (o *HomeSiteInfo) GetAdUserOrGroupIdOk() (*string, bool)`

GetAdUserOrGroupIdOk returns a tuple with the AdUserOrGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserOrGroupId

`func (o *HomeSiteInfo) SetAdUserOrGroupId(v string)`

SetAdUserOrGroupId sets AdUserOrGroupId field to given value.

### HasAdUserOrGroupId

`func (o *HomeSiteInfo) HasAdUserOrGroupId() bool`

HasAdUserOrGroupId returns a boolean if a field has been set.

### GetGlobalApplicationEntitlementId

`func (o *HomeSiteInfo) GetGlobalApplicationEntitlementId() string`

GetGlobalApplicationEntitlementId returns the GlobalApplicationEntitlementId field if non-nil, zero value otherwise.

### GetGlobalApplicationEntitlementIdOk

`func (o *HomeSiteInfo) GetGlobalApplicationEntitlementIdOk() (*string, bool)`

GetGlobalApplicationEntitlementIdOk returns a tuple with the GlobalApplicationEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApplicationEntitlementId

`func (o *HomeSiteInfo) SetGlobalApplicationEntitlementId(v string)`

SetGlobalApplicationEntitlementId sets GlobalApplicationEntitlementId field to given value.

### HasGlobalApplicationEntitlementId

`func (o *HomeSiteInfo) HasGlobalApplicationEntitlementId() bool`

HasGlobalApplicationEntitlementId returns a boolean if a field has been set.

### GetGlobalDesktopEntitlementId

`func (o *HomeSiteInfo) GetGlobalDesktopEntitlementId() string`

GetGlobalDesktopEntitlementId returns the GlobalDesktopEntitlementId field if non-nil, zero value otherwise.

### GetGlobalDesktopEntitlementIdOk

`func (o *HomeSiteInfo) GetGlobalDesktopEntitlementIdOk() (*string, bool)`

GetGlobalDesktopEntitlementIdOk returns a tuple with the GlobalDesktopEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDesktopEntitlementId

`func (o *HomeSiteInfo) SetGlobalDesktopEntitlementId(v string)`

SetGlobalDesktopEntitlementId sets GlobalDesktopEntitlementId field to given value.

### HasGlobalDesktopEntitlementId

`func (o *HomeSiteInfo) HasGlobalDesktopEntitlementId() bool`

HasGlobalDesktopEntitlementId returns a boolean if a field has been set.

### GetId

`func (o *HomeSiteInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HomeSiteInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HomeSiteInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HomeSiteInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSiteId

`func (o *HomeSiteInfo) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *HomeSiteInfo) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *HomeSiteInfo) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *HomeSiteInfo) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


