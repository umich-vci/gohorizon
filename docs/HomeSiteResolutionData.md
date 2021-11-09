# HomeSiteResolutionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | SID of the group through which the user gets this home site. | [optional] 
**Id** | Pointer to **string** | ID of the Home Site Configuration. | [optional] 
**Resolved** | Pointer to **bool** | Indicates whether this is the resolved home site. | [optional] 
**SiteId** | Pointer to **string** | ID of the Site configured as Home Site. | [optional] 
**Type** | Pointer to **string** | Indicates whether the home site is directly assigned to the user or group and whether it is associated with a global entitlement. * USER_OVERRIDE: Indicates that a home site override is associated with a global entitlement for the user. * GROUP_OVERRIDE: Indicates that a home site override is associated with a global entitlement for a group that the user belongs to. * USER_DEFAULT: Indicates that a home site is directly assigned to the user. * GROUP_DEFAULT: Indicates that a home site is directly assigned to a group that the user belongs to. | [optional] 

## Methods

### NewHomeSiteResolutionData

`func NewHomeSiteResolutionData() *HomeSiteResolutionData`

NewHomeSiteResolutionData instantiates a new HomeSiteResolutionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHomeSiteResolutionDataWithDefaults

`func NewHomeSiteResolutionDataWithDefaults() *HomeSiteResolutionData`

NewHomeSiteResolutionDataWithDefaults instantiates a new HomeSiteResolutionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *HomeSiteResolutionData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *HomeSiteResolutionData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *HomeSiteResolutionData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *HomeSiteResolutionData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *HomeSiteResolutionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HomeSiteResolutionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HomeSiteResolutionData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HomeSiteResolutionData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResolved

`func (o *HomeSiteResolutionData) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *HomeSiteResolutionData) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *HomeSiteResolutionData) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *HomeSiteResolutionData) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetSiteId

`func (o *HomeSiteResolutionData) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *HomeSiteResolutionData) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *HomeSiteResolutionData) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *HomeSiteResolutionData) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetType

`func (o *HomeSiteResolutionData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HomeSiteResolutionData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HomeSiteResolutionData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HomeSiteResolutionData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


