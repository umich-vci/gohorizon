# PodUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudManaged** | Pointer to **bool** | Indicates whether this pod is managed from cloud. Default value is false. | [optional] 
**Description** | Pointer to **string** | Description of this pod. | [optional] 
**Name** | **string** | Name of this pod. | 
**SiteId** | **string** | ID of the site this pod belongs to. | 

## Methods

### NewPodUpdateSpec

`func NewPodUpdateSpec(name string, siteId string, ) *PodUpdateSpec`

NewPodUpdateSpec instantiates a new PodUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodUpdateSpecWithDefaults

`func NewPodUpdateSpecWithDefaults() *PodUpdateSpec`

NewPodUpdateSpecWithDefaults instantiates a new PodUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudManaged

`func (o *PodUpdateSpec) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *PodUpdateSpec) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *PodUpdateSpec) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *PodUpdateSpec) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetDescription

`func (o *PodUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PodUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PodUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PodUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PodUpdateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodUpdateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodUpdateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetSiteId

`func (o *PodUpdateSpec) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PodUpdateSpec) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PodUpdateSpec) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


