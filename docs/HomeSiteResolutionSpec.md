# HomeSiteResolutionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalApplicationEntitlementId** | Pointer to **string** | ID of the Global Application Entitlement this home site resolution is for. | [optional] 
**GlobalDesktopEntitlementId** | Pointer to **string** | ID of the Global Desktop Entitlement this home site resolution is for. | [optional] 
**UserId** | **string** | SID of the user for whom home site is to be resolved. | 

## Methods

### NewHomeSiteResolutionSpec

`func NewHomeSiteResolutionSpec(userId string, ) *HomeSiteResolutionSpec`

NewHomeSiteResolutionSpec instantiates a new HomeSiteResolutionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHomeSiteResolutionSpecWithDefaults

`func NewHomeSiteResolutionSpecWithDefaults() *HomeSiteResolutionSpec`

NewHomeSiteResolutionSpecWithDefaults instantiates a new HomeSiteResolutionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalApplicationEntitlementId

`func (o *HomeSiteResolutionSpec) GetGlobalApplicationEntitlementId() string`

GetGlobalApplicationEntitlementId returns the GlobalApplicationEntitlementId field if non-nil, zero value otherwise.

### GetGlobalApplicationEntitlementIdOk

`func (o *HomeSiteResolutionSpec) GetGlobalApplicationEntitlementIdOk() (*string, bool)`

GetGlobalApplicationEntitlementIdOk returns a tuple with the GlobalApplicationEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApplicationEntitlementId

`func (o *HomeSiteResolutionSpec) SetGlobalApplicationEntitlementId(v string)`

SetGlobalApplicationEntitlementId sets GlobalApplicationEntitlementId field to given value.

### HasGlobalApplicationEntitlementId

`func (o *HomeSiteResolutionSpec) HasGlobalApplicationEntitlementId() bool`

HasGlobalApplicationEntitlementId returns a boolean if a field has been set.

### GetGlobalDesktopEntitlementId

`func (o *HomeSiteResolutionSpec) GetGlobalDesktopEntitlementId() string`

GetGlobalDesktopEntitlementId returns the GlobalDesktopEntitlementId field if non-nil, zero value otherwise.

### GetGlobalDesktopEntitlementIdOk

`func (o *HomeSiteResolutionSpec) GetGlobalDesktopEntitlementIdOk() (*string, bool)`

GetGlobalDesktopEntitlementIdOk returns a tuple with the GlobalDesktopEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDesktopEntitlementId

`func (o *HomeSiteResolutionSpec) SetGlobalDesktopEntitlementId(v string)`

SetGlobalDesktopEntitlementId sets GlobalDesktopEntitlementId field to given value.

### HasGlobalDesktopEntitlementId

`func (o *HomeSiteResolutionSpec) HasGlobalDesktopEntitlementId() bool`

HasGlobalDesktopEntitlementId returns a boolean if a field has been set.

### GetUserId

`func (o *HomeSiteResolutionSpec) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HomeSiteResolutionSpec) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HomeSiteResolutionSpec) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


