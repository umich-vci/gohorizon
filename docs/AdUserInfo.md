# ADUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | DNS name of the domain in which this user or group belongs to. | [optional] 
**GroupGuids** | Pointer to **[]string** | Guids of the user&#39;s groups in RFC 4122 format. | [optional] 
**GroupSids** | Pointer to **[]string** | List of unique SIDs of the groups, this user or group belongs to. | [optional] 
**UserGuid** | Pointer to **string** | GUID of the user in RFC 4122 format. | [optional] 
**UserPrincipalName** | Pointer to **string** | User Principal name(UPN) of this user. | [optional] 
**UserSid** | Pointer to **string** | Unique SID representing this AD User. | [optional] 
**Username** | Pointer to **string** | Username of this user. | [optional] 

## Methods

### NewADUserInfo

`func NewADUserInfo() *ADUserInfo`

NewADUserInfo instantiates a new ADUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADUserInfoWithDefaults

`func NewADUserInfoWithDefaults() *ADUserInfo`

NewADUserInfoWithDefaults instantiates a new ADUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ADUserInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ADUserInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ADUserInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ADUserInfo) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGroupGuids

`func (o *ADUserInfo) GetGroupGuids() []string`

GetGroupGuids returns the GroupGuids field if non-nil, zero value otherwise.

### GetGroupGuidsOk

`func (o *ADUserInfo) GetGroupGuidsOk() (*[]string, bool)`

GetGroupGuidsOk returns a tuple with the GroupGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuids

`func (o *ADUserInfo) SetGroupGuids(v []string)`

SetGroupGuids sets GroupGuids field to given value.

### HasGroupGuids

`func (o *ADUserInfo) HasGroupGuids() bool`

HasGroupGuids returns a boolean if a field has been set.

### GetGroupSids

`func (o *ADUserInfo) GetGroupSids() []string`

GetGroupSids returns the GroupSids field if non-nil, zero value otherwise.

### GetGroupSidsOk

`func (o *ADUserInfo) GetGroupSidsOk() (*[]string, bool)`

GetGroupSidsOk returns a tuple with the GroupSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSids

`func (o *ADUserInfo) SetGroupSids(v []string)`

SetGroupSids sets GroupSids field to given value.

### HasGroupSids

`func (o *ADUserInfo) HasGroupSids() bool`

HasGroupSids returns a boolean if a field has been set.

### GetUserGuid

`func (o *ADUserInfo) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *ADUserInfo) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *ADUserInfo) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *ADUserInfo) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *ADUserInfo) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *ADUserInfo) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *ADUserInfo) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *ADUserInfo) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetUserSid

`func (o *ADUserInfo) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *ADUserInfo) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *ADUserInfo) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *ADUserInfo) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### GetUsername

`func (o *ADUserInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ADUserInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ADUserInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ADUserInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


