# InstantCloneDomainAccountCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainId** | **string** | SID of the AD Domain that this account user belongs to. | 
**Password** | **[]string** | Password of the account. | 
**Username** | **string** | User name of the account. | 

## Methods

### NewInstantCloneDomainAccountCreateSpec

`func NewInstantCloneDomainAccountCreateSpec(adDomainId string, password []string, username string, ) *InstantCloneDomainAccountCreateSpec`

NewInstantCloneDomainAccountCreateSpec instantiates a new InstantCloneDomainAccountCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantCloneDomainAccountCreateSpecWithDefaults

`func NewInstantCloneDomainAccountCreateSpecWithDefaults() *InstantCloneDomainAccountCreateSpec`

NewInstantCloneDomainAccountCreateSpecWithDefaults instantiates a new InstantCloneDomainAccountCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainId

`func (o *InstantCloneDomainAccountCreateSpec) GetAdDomainId() string`

GetAdDomainId returns the AdDomainId field if non-nil, zero value otherwise.

### GetAdDomainIdOk

`func (o *InstantCloneDomainAccountCreateSpec) GetAdDomainIdOk() (*string, bool)`

GetAdDomainIdOk returns a tuple with the AdDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainId

`func (o *InstantCloneDomainAccountCreateSpec) SetAdDomainId(v string)`

SetAdDomainId sets AdDomainId field to given value.


### GetPassword

`func (o *InstantCloneDomainAccountCreateSpec) GetPassword() []string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InstantCloneDomainAccountCreateSpec) GetPasswordOk() (*[]string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InstantCloneDomainAccountCreateSpec) SetPassword(v []string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *InstantCloneDomainAccountCreateSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InstantCloneDomainAccountCreateSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InstantCloneDomainAccountCreateSpec) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


