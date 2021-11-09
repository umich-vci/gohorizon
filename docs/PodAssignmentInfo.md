# PodAssignmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalApplicationEntitlementId** | Pointer to **string** | ID of the Global Application Entitlement associated with this pod assignment.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**GlobalDesktopEntitlementId** | Pointer to **string** | ID of the Global Desktop Entitlement associated with this pod assignment.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this pod assignment.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**PodId** | Pointer to **string** | ID representing the pod associated with this pod assignment.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**UserId** | Pointer to **string** | SID of the user associated with this pod assignment.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 

## Methods

### NewPodAssignmentInfo

`func NewPodAssignmentInfo() *PodAssignmentInfo`

NewPodAssignmentInfo instantiates a new PodAssignmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodAssignmentInfoWithDefaults

`func NewPodAssignmentInfoWithDefaults() *PodAssignmentInfo`

NewPodAssignmentInfoWithDefaults instantiates a new PodAssignmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalApplicationEntitlementId

`func (o *PodAssignmentInfo) GetGlobalApplicationEntitlementId() string`

GetGlobalApplicationEntitlementId returns the GlobalApplicationEntitlementId field if non-nil, zero value otherwise.

### GetGlobalApplicationEntitlementIdOk

`func (o *PodAssignmentInfo) GetGlobalApplicationEntitlementIdOk() (*string, bool)`

GetGlobalApplicationEntitlementIdOk returns a tuple with the GlobalApplicationEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApplicationEntitlementId

`func (o *PodAssignmentInfo) SetGlobalApplicationEntitlementId(v string)`

SetGlobalApplicationEntitlementId sets GlobalApplicationEntitlementId field to given value.

### HasGlobalApplicationEntitlementId

`func (o *PodAssignmentInfo) HasGlobalApplicationEntitlementId() bool`

HasGlobalApplicationEntitlementId returns a boolean if a field has been set.

### GetGlobalDesktopEntitlementId

`func (o *PodAssignmentInfo) GetGlobalDesktopEntitlementId() string`

GetGlobalDesktopEntitlementId returns the GlobalDesktopEntitlementId field if non-nil, zero value otherwise.

### GetGlobalDesktopEntitlementIdOk

`func (o *PodAssignmentInfo) GetGlobalDesktopEntitlementIdOk() (*string, bool)`

GetGlobalDesktopEntitlementIdOk returns a tuple with the GlobalDesktopEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDesktopEntitlementId

`func (o *PodAssignmentInfo) SetGlobalDesktopEntitlementId(v string)`

SetGlobalDesktopEntitlementId sets GlobalDesktopEntitlementId field to given value.

### HasGlobalDesktopEntitlementId

`func (o *PodAssignmentInfo) HasGlobalDesktopEntitlementId() bool`

HasGlobalDesktopEntitlementId returns a boolean if a field has been set.

### GetId

`func (o *PodAssignmentInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PodAssignmentInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PodAssignmentInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PodAssignmentInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPodId

`func (o *PodAssignmentInfo) GetPodId() string`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *PodAssignmentInfo) GetPodIdOk() (*string, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *PodAssignmentInfo) SetPodId(v string)`

SetPodId sets PodId field to given value.

### HasPodId

`func (o *PodAssignmentInfo) HasPodId() bool`

HasPodId returns a boolean if a field has been set.

### GetUserId

`func (o *PodAssignmentInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PodAssignmentInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PodAssignmentInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PodAssignmentInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


