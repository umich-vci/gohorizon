# EntitlementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdUserOrGroupIds** | Pointer to **[]string** | List of ad-user-or-group SIDs for the entitlement operations on the given resource. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the resource. | [optional] 

## Methods

### NewEntitlementSpec

`func NewEntitlementSpec() *EntitlementSpec`

NewEntitlementSpec instantiates a new EntitlementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementSpecWithDefaults

`func NewEntitlementSpecWithDefaults() *EntitlementSpec`

NewEntitlementSpecWithDefaults instantiates a new EntitlementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdUserOrGroupIds

`func (o *EntitlementSpec) GetAdUserOrGroupIds() []string`

GetAdUserOrGroupIds returns the AdUserOrGroupIds field if non-nil, zero value otherwise.

### GetAdUserOrGroupIdsOk

`func (o *EntitlementSpec) GetAdUserOrGroupIdsOk() (*[]string, bool)`

GetAdUserOrGroupIdsOk returns a tuple with the AdUserOrGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserOrGroupIds

`func (o *EntitlementSpec) SetAdUserOrGroupIds(v []string)`

SetAdUserOrGroupIds sets AdUserOrGroupIds field to given value.

### HasAdUserOrGroupIds

`func (o *EntitlementSpec) HasAdUserOrGroupIds() bool`

HasAdUserOrGroupIds returns a boolean if a field has been set.

### GetId

`func (o *EntitlementSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementSpec) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


