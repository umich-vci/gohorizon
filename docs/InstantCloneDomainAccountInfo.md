# InstantCloneDomainAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainId** | **string** | SID of the AD Domain that this account user belongs to. | 
**Id** | **string** | Unique ID representing instant clone domain account. | 
**Username** | **string** | User name of the account. | 

## Methods

### NewInstantCloneDomainAccountInfo

`func NewInstantCloneDomainAccountInfo(adDomainId string, id string, username string, ) *InstantCloneDomainAccountInfo`

NewInstantCloneDomainAccountInfo instantiates a new InstantCloneDomainAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantCloneDomainAccountInfoWithDefaults

`func NewInstantCloneDomainAccountInfoWithDefaults() *InstantCloneDomainAccountInfo`

NewInstantCloneDomainAccountInfoWithDefaults instantiates a new InstantCloneDomainAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainId

`func (o *InstantCloneDomainAccountInfo) GetAdDomainId() string`

GetAdDomainId returns the AdDomainId field if non-nil, zero value otherwise.

### GetAdDomainIdOk

`func (o *InstantCloneDomainAccountInfo) GetAdDomainIdOk() (*string, bool)`

GetAdDomainIdOk returns a tuple with the AdDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainId

`func (o *InstantCloneDomainAccountInfo) SetAdDomainId(v string)`

SetAdDomainId sets AdDomainId field to given value.


### GetId

`func (o *InstantCloneDomainAccountInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstantCloneDomainAccountInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstantCloneDomainAccountInfo) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *InstantCloneDomainAccountInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InstantCloneDomainAccountInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InstantCloneDomainAccountInfo) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


