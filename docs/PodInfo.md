# PodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveGlobalApplicationEntitlements** | Pointer to **[]string** | List of Global Application Entitlements with mappings to application pools within this pod. | [optional] 
**ActiveGlobalDesktopEntitlements** | Pointer to **[]string** | List of Global Desktop Entitlements with mappings to desktop pools within this pod. | [optional] 
**CloudManaged** | Pointer to **bool** | Indicates whether this pod is managed from cloud. | [optional] 
**Description** | Pointer to **string** | Description of this pod. | [optional] 
**Endpoints** | Pointer to **[]string** | List of endpoints within this pod. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this pod. | [optional] 
**LocalPod** | Pointer to **bool** | Indicates whether this is the local pod the request was made from.&lt;br&gt;Only one pod in the pod federation will return true. | [optional] 
**Name** | Pointer to **string** | Name of this pod. | [optional] 
**SiteId** | Pointer to **string** | ID of the site this pod belongs to. | [optional] 

## Methods

### NewPodInfo

`func NewPodInfo() *PodInfo`

NewPodInfo instantiates a new PodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodInfoWithDefaults

`func NewPodInfoWithDefaults() *PodInfo`

NewPodInfoWithDefaults instantiates a new PodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveGlobalApplicationEntitlements

`func (o *PodInfo) GetActiveGlobalApplicationEntitlements() []string`

GetActiveGlobalApplicationEntitlements returns the ActiveGlobalApplicationEntitlements field if non-nil, zero value otherwise.

### GetActiveGlobalApplicationEntitlementsOk

`func (o *PodInfo) GetActiveGlobalApplicationEntitlementsOk() (*[]string, bool)`

GetActiveGlobalApplicationEntitlementsOk returns a tuple with the ActiveGlobalApplicationEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGlobalApplicationEntitlements

`func (o *PodInfo) SetActiveGlobalApplicationEntitlements(v []string)`

SetActiveGlobalApplicationEntitlements sets ActiveGlobalApplicationEntitlements field to given value.

### HasActiveGlobalApplicationEntitlements

`func (o *PodInfo) HasActiveGlobalApplicationEntitlements() bool`

HasActiveGlobalApplicationEntitlements returns a boolean if a field has been set.

### GetActiveGlobalDesktopEntitlements

`func (o *PodInfo) GetActiveGlobalDesktopEntitlements() []string`

GetActiveGlobalDesktopEntitlements returns the ActiveGlobalDesktopEntitlements field if non-nil, zero value otherwise.

### GetActiveGlobalDesktopEntitlementsOk

`func (o *PodInfo) GetActiveGlobalDesktopEntitlementsOk() (*[]string, bool)`

GetActiveGlobalDesktopEntitlementsOk returns a tuple with the ActiveGlobalDesktopEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGlobalDesktopEntitlements

`func (o *PodInfo) SetActiveGlobalDesktopEntitlements(v []string)`

SetActiveGlobalDesktopEntitlements sets ActiveGlobalDesktopEntitlements field to given value.

### HasActiveGlobalDesktopEntitlements

`func (o *PodInfo) HasActiveGlobalDesktopEntitlements() bool`

HasActiveGlobalDesktopEntitlements returns a boolean if a field has been set.

### GetCloudManaged

`func (o *PodInfo) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *PodInfo) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *PodInfo) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *PodInfo) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetDescription

`func (o *PodInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PodInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PodInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PodInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoints

`func (o *PodInfo) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PodInfo) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PodInfo) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *PodInfo) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetId

`func (o *PodInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PodInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PodInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PodInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalPod

`func (o *PodInfo) GetLocalPod() bool`

GetLocalPod returns the LocalPod field if non-nil, zero value otherwise.

### GetLocalPodOk

`func (o *PodInfo) GetLocalPodOk() (*bool, bool)`

GetLocalPodOk returns a tuple with the LocalPod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPod

`func (o *PodInfo) SetLocalPod(v bool)`

SetLocalPod sets LocalPod field to given value.

### HasLocalPod

`func (o *PodInfo) HasLocalPod() bool`

HasLocalPod returns a boolean if a field has been set.

### GetName

`func (o *PodInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PodInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteId

`func (o *PodInfo) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PodInfo) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PodInfo) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PodInfo) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


