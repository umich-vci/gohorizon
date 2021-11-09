# SiteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Detailed description of the site. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the site. | [optional] 
**Name** | Pointer to **string** | Name of the site. | [optional] 
**Pods** | Pointer to **[]string** | Member pods for this site. | [optional] 

## Methods

### NewSiteInfo

`func NewSiteInfo() *SiteInfo`

NewSiteInfo instantiates a new SiteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteInfoWithDefaults

`func NewSiteInfoWithDefaults() *SiteInfo`

NewSiteInfoWithDefaults instantiates a new SiteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SiteInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SiteInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SiteInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SiteInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SiteInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SiteInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPods

`func (o *SiteInfo) GetPods() []string`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *SiteInfo) GetPodsOk() (*[]string, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *SiteInfo) SetPods(v []string)`

SetPods sets Pods field to given value.

### HasPods

`func (o *SiteInfo) HasPods() bool`

HasPods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


