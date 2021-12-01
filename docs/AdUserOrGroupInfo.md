# ADUserOrGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | AD container for this user or group. | [optional] 
**Description** | Pointer to **string** | Description number of this user or group. Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**DisplayName** | Pointer to **string** | Login name with domain of this user or group. | [optional] 
**DistinguishedName** | Pointer to **string** | Active Directory distinguished name for this user or group. | [optional] 
**Domain** | Pointer to **string** | DNS name of the domain in which this user or group belongs. Supported Filters : &#39;Equals&#39;.  Also, if &#39;Or&#39; filter is used anywhere in filter string for this model class, then that &#39;Or&#39; filter should nest only &#39;Equals&#39; filter on &#39;domain&#39; or &#39;id&#39; field. | [optional] 
**Email** | Pointer to **string** | Email address of this user or group. Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**FirstName** | Pointer to **string** | First name of this user or group. | [optional] 
**Group** | Pointer to **bool** | Indicates if this object represents a group. This field is NOT supported in filter string. To use any filter on &#39;group&#39;, use &#39;group_only&#39; query param directly. | [optional] 
**GroupMembershipIds** | Pointer to **[]string** | List of unique SIDs of the groups, this user or group belongs to. | [optional] 
**Guid** | Pointer to **string** | GUID of the user or group in RFC 4122 format. Supported Filters : &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **string** | Unique SID representing this AD User or Group. Supported Filters : &#39;Equals&#39;.&#39;Or&#39; filter chain of &#39;Equals&#39; filter can be used to query for more than one id. For this model, if &#39;Or&#39; filter is used, then it should nest only &#39;Equals&#39; filter on &#39;domain&#39; or &#39;id&#39; field. | [optional] 
**KioskUser** | Pointer to **bool** | Indicates if this user or group is a \&quot;kiosk user\&quot; that supports client authentication. Client authentication is the process of supporting client devices directly logging into resources. | [optional] 
**LastName** | Pointer to **string** | Last name of this user or group. | [optional] 
**LoginName** | Pointer to **string** | Login name of this user or group.  Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**LongDisplayName** | Pointer to **string** | Login name, domain and name for this user or group, else display name | [optional] 
**Name** | Pointer to **string** | Name of this user or group.  Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**Phone** | Pointer to **string** | Phone number of this user. Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**SubGroupCount** | Pointer to **int32** | Number of subgroups in this group, or 0 if not a group. | [optional] 
**UserCount** | Pointer to **int32** | Number of users in this group, or 0 if not a group. | [optional] 
**UserDisplayName** | Pointer to **string** | User or group&#39;s display name. This corresponds with displayName attribute in AD. | [optional] 
**UserPrincipalName** | Pointer to **string** | User Principal name(UPN) of this user. Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 

## Methods

### NewADUserOrGroupInfo

`func NewADUserOrGroupInfo() *ADUserOrGroupInfo`

NewADUserOrGroupInfo instantiates a new ADUserOrGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADUserOrGroupInfoWithDefaults

`func NewADUserOrGroupInfoWithDefaults() *ADUserOrGroupInfo`

NewADUserOrGroupInfoWithDefaults instantiates a new ADUserOrGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *ADUserOrGroupInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ADUserOrGroupInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ADUserOrGroupInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ADUserOrGroupInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDescription

`func (o *ADUserOrGroupInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ADUserOrGroupInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ADUserOrGroupInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ADUserOrGroupInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ADUserOrGroupInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ADUserOrGroupInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ADUserOrGroupInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ADUserOrGroupInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *ADUserOrGroupInfo) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *ADUserOrGroupInfo) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *ADUserOrGroupInfo) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *ADUserOrGroupInfo) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetDomain

`func (o *ADUserOrGroupInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ADUserOrGroupInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ADUserOrGroupInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ADUserOrGroupInfo) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEmail

`func (o *ADUserOrGroupInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ADUserOrGroupInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ADUserOrGroupInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ADUserOrGroupInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *ADUserOrGroupInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ADUserOrGroupInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ADUserOrGroupInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ADUserOrGroupInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGroup

`func (o *ADUserOrGroupInfo) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ADUserOrGroupInfo) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ADUserOrGroupInfo) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ADUserOrGroupInfo) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupMembershipIds

`func (o *ADUserOrGroupInfo) GetGroupMembershipIds() []string`

GetGroupMembershipIds returns the GroupMembershipIds field if non-nil, zero value otherwise.

### GetGroupMembershipIdsOk

`func (o *ADUserOrGroupInfo) GetGroupMembershipIdsOk() (*[]string, bool)`

GetGroupMembershipIdsOk returns a tuple with the GroupMembershipIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipIds

`func (o *ADUserOrGroupInfo) SetGroupMembershipIds(v []string)`

SetGroupMembershipIds sets GroupMembershipIds field to given value.

### HasGroupMembershipIds

`func (o *ADUserOrGroupInfo) HasGroupMembershipIds() bool`

HasGroupMembershipIds returns a boolean if a field has been set.

### GetGuid

`func (o *ADUserOrGroupInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ADUserOrGroupInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ADUserOrGroupInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ADUserOrGroupInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetId

`func (o *ADUserOrGroupInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ADUserOrGroupInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ADUserOrGroupInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ADUserOrGroupInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKioskUser

`func (o *ADUserOrGroupInfo) GetKioskUser() bool`

GetKioskUser returns the KioskUser field if non-nil, zero value otherwise.

### GetKioskUserOk

`func (o *ADUserOrGroupInfo) GetKioskUserOk() (*bool, bool)`

GetKioskUserOk returns a tuple with the KioskUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKioskUser

`func (o *ADUserOrGroupInfo) SetKioskUser(v bool)`

SetKioskUser sets KioskUser field to given value.

### HasKioskUser

`func (o *ADUserOrGroupInfo) HasKioskUser() bool`

HasKioskUser returns a boolean if a field has been set.

### GetLastName

`func (o *ADUserOrGroupInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ADUserOrGroupInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ADUserOrGroupInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ADUserOrGroupInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLoginName

`func (o *ADUserOrGroupInfo) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *ADUserOrGroupInfo) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *ADUserOrGroupInfo) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *ADUserOrGroupInfo) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLongDisplayName

`func (o *ADUserOrGroupInfo) GetLongDisplayName() string`

GetLongDisplayName returns the LongDisplayName field if non-nil, zero value otherwise.

### GetLongDisplayNameOk

`func (o *ADUserOrGroupInfo) GetLongDisplayNameOk() (*string, bool)`

GetLongDisplayNameOk returns a tuple with the LongDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDisplayName

`func (o *ADUserOrGroupInfo) SetLongDisplayName(v string)`

SetLongDisplayName sets LongDisplayName field to given value.

### HasLongDisplayName

`func (o *ADUserOrGroupInfo) HasLongDisplayName() bool`

HasLongDisplayName returns a boolean if a field has been set.

### GetName

`func (o *ADUserOrGroupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ADUserOrGroupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ADUserOrGroupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ADUserOrGroupInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ADUserOrGroupInfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ADUserOrGroupInfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ADUserOrGroupInfo) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ADUserOrGroupInfo) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSubGroupCount

`func (o *ADUserOrGroupInfo) GetSubGroupCount() int32`

GetSubGroupCount returns the SubGroupCount field if non-nil, zero value otherwise.

### GetSubGroupCountOk

`func (o *ADUserOrGroupInfo) GetSubGroupCountOk() (*int32, bool)`

GetSubGroupCountOk returns a tuple with the SubGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGroupCount

`func (o *ADUserOrGroupInfo) SetSubGroupCount(v int32)`

SetSubGroupCount sets SubGroupCount field to given value.

### HasSubGroupCount

`func (o *ADUserOrGroupInfo) HasSubGroupCount() bool`

HasSubGroupCount returns a boolean if a field has been set.

### GetUserCount

`func (o *ADUserOrGroupInfo) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *ADUserOrGroupInfo) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *ADUserOrGroupInfo) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *ADUserOrGroupInfo) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *ADUserOrGroupInfo) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *ADUserOrGroupInfo) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *ADUserOrGroupInfo) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *ADUserOrGroupInfo) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *ADUserOrGroupInfo) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *ADUserOrGroupInfo) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *ADUserOrGroupInfo) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *ADUserOrGroupInfo) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


