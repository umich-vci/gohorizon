# HomeSiteResolutionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalApplicationEntitlementId** | Pointer to **string** | ID of the Global Application Entitlement this home site resolution is for. | [optional] 
**GlobalDesktopEntitlementId** | Pointer to **string** | ID of the Global Desktop Entitlement this home site resolution is for. | [optional] 
**ResolutionData** | Pointer to [**[]HomeSiteResolutionData**](HomeSiteResolutionData.md) | Home Site Resolution Data for this Global Entitlement. | [optional] 

## Methods

### NewHomeSiteResolutionInfo

`func NewHomeSiteResolutionInfo() *HomeSiteResolutionInfo`

NewHomeSiteResolutionInfo instantiates a new HomeSiteResolutionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHomeSiteResolutionInfoWithDefaults

`func NewHomeSiteResolutionInfoWithDefaults() *HomeSiteResolutionInfo`

NewHomeSiteResolutionInfoWithDefaults instantiates a new HomeSiteResolutionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalApplicationEntitlementId

`func (o *HomeSiteResolutionInfo) GetGlobalApplicationEntitlementId() string`

GetGlobalApplicationEntitlementId returns the GlobalApplicationEntitlementId field if non-nil, zero value otherwise.

### GetGlobalApplicationEntitlementIdOk

`func (o *HomeSiteResolutionInfo) GetGlobalApplicationEntitlementIdOk() (*string, bool)`

GetGlobalApplicationEntitlementIdOk returns a tuple with the GlobalApplicationEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApplicationEntitlementId

`func (o *HomeSiteResolutionInfo) SetGlobalApplicationEntitlementId(v string)`

SetGlobalApplicationEntitlementId sets GlobalApplicationEntitlementId field to given value.

### HasGlobalApplicationEntitlementId

`func (o *HomeSiteResolutionInfo) HasGlobalApplicationEntitlementId() bool`

HasGlobalApplicationEntitlementId returns a boolean if a field has been set.

### GetGlobalDesktopEntitlementId

`func (o *HomeSiteResolutionInfo) GetGlobalDesktopEntitlementId() string`

GetGlobalDesktopEntitlementId returns the GlobalDesktopEntitlementId field if non-nil, zero value otherwise.

### GetGlobalDesktopEntitlementIdOk

`func (o *HomeSiteResolutionInfo) GetGlobalDesktopEntitlementIdOk() (*string, bool)`

GetGlobalDesktopEntitlementIdOk returns a tuple with the GlobalDesktopEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDesktopEntitlementId

`func (o *HomeSiteResolutionInfo) SetGlobalDesktopEntitlementId(v string)`

SetGlobalDesktopEntitlementId sets GlobalDesktopEntitlementId field to given value.

### HasGlobalDesktopEntitlementId

`func (o *HomeSiteResolutionInfo) HasGlobalDesktopEntitlementId() bool`

HasGlobalDesktopEntitlementId returns a boolean if a field has been set.

### GetResolutionData

`func (o *HomeSiteResolutionInfo) GetResolutionData() []HomeSiteResolutionData`

GetResolutionData returns the ResolutionData field if non-nil, zero value otherwise.

### GetResolutionDataOk

`func (o *HomeSiteResolutionInfo) GetResolutionDataOk() (*[]HomeSiteResolutionData, bool)`

GetResolutionDataOk returns a tuple with the ResolutionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionData

`func (o *HomeSiteResolutionInfo) SetResolutionData(v []HomeSiteResolutionData)`

SetResolutionData sets ResolutionData field to given value.

### HasResolutionData

`func (o *HomeSiteResolutionInfo) HasResolutionData() bool`

HasResolutionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


