# EntitlementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdUserOrGroupIds** | Pointer to **[]string** | List of ad-user-or-group SIDs which are entitled to the given resource. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the resource.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 

## Methods

### NewEntitlementInfo

`func NewEntitlementInfo() *EntitlementInfo`

NewEntitlementInfo instantiates a new EntitlementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementInfoWithDefaults

`func NewEntitlementInfoWithDefaults() *EntitlementInfo`

NewEntitlementInfoWithDefaults instantiates a new EntitlementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdUserOrGroupIds

`func (o *EntitlementInfo) GetAdUserOrGroupIds() []string`

GetAdUserOrGroupIds returns the AdUserOrGroupIds field if non-nil, zero value otherwise.

### GetAdUserOrGroupIdsOk

`func (o *EntitlementInfo) GetAdUserOrGroupIdsOk() (*[]string, bool)`

GetAdUserOrGroupIdsOk returns a tuple with the AdUserOrGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserOrGroupIds

`func (o *EntitlementInfo) SetAdUserOrGroupIds(v []string)`

SetAdUserOrGroupIds sets AdUserOrGroupIds field to given value.

### HasAdUserOrGroupIds

`func (o *EntitlementInfo) HasAdUserOrGroupIds() bool`

HasAdUserOrGroupIds returns a boolean if a field has been set.

### GetId

`func (o *EntitlementInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementInfo) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


