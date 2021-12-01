# GlobalSessionActionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | IDs of the sessions on which action is to be performed. | 
**PodId** | **string** | ID of the hosting pod for the sessions. | 

## Methods

### NewGlobalSessionActionSpec

`func NewGlobalSessionActionSpec(ids []string, podId string, ) *GlobalSessionActionSpec`

NewGlobalSessionActionSpec instantiates a new GlobalSessionActionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSessionActionSpecWithDefaults

`func NewGlobalSessionActionSpecWithDefaults() *GlobalSessionActionSpec`

NewGlobalSessionActionSpecWithDefaults instantiates a new GlobalSessionActionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *GlobalSessionActionSpec) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *GlobalSessionActionSpec) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *GlobalSessionActionSpec) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetPodId

`func (o *GlobalSessionActionSpec) GetPodId() string`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *GlobalSessionActionSpec) GetPodIdOk() (*string, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *GlobalSessionActionSpec) SetPodId(v string)`

SetPodId sets PodId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


