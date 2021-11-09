# ADUserOrGroupSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | AD container for this user or group. | [optional] 
**Description** | Pointer to **string** | Description number of this user or group. Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**DisplayName** | Pointer to **string** | Login name with domain of this user or group. | [optional] 
**DistinguishedName** | Pointer to **string** | Active Directory distinguished name for this user or group. | [optional] 
**Domain** | **string** | DNS name of the domain in which this user or group belongs. Supported Filters : &#39;Equals&#39;.  Also, if &#39;Or&#39; filter is used anywhere in filter string for this model class, then that &#39;Or&#39; filter should nest only &#39;Equals&#39; filter on &#39;domain&#39; or &#39;id&#39; field. | 
**Email** | Pointer to **string** | Email address of this user or group. Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**FirstName** | Pointer to **string** | First name of this user or group. | [optional] 
**Group** | **bool** | Indicates if this object represents a group. This field is NOT supported in filter string. To use any filter on &#39;group&#39;, use &#39;group_only&#39; query param directly. | 
**Guid** | **string** | GUID of the user or group in RFC 4122 format. Supported Filters : &#39;Equals&#39;. | 
**Id** | **string** | Unique SID representing this AD User or Group. Supported Filters : &#39;Equals&#39;.&#39;Or&#39; filter chain of &#39;Equals&#39; filter can be used to query for more than one id. For this model, if &#39;Or&#39; filter is used, then it should nest only &#39;Equals&#39; filter on &#39;domain&#39; or &#39;id&#39; field. | 
**KioskUser** | Pointer to **bool** | Indicates if this user or group is a \&quot;kiosk user\&quot; that supports client authentication. Client authentication is the process of supporting client devices directly logging into resources. | [optional] 
**LastName** | Pointer to **string** | Last name of this user or group. | [optional] 
**LoginName** | Pointer to **string** | Login name of this user or group.  Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**LongDisplayName** | Pointer to **string** | Login name, domain and name for this user or group, else display name | [optional] 
**Name** | Pointer to **string** | Name of this user or group.  Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**Phone** | Pointer to **string** | Phone number of this user. Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 
**UserDisplayName** | Pointer to **string** | User or group&#39;s display name. This corresponds with displayName attribute in AD. | [optional] 
**UserPrincipalName** | Pointer to **string** | User Principal name(UPN) of this user. Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;Contains&#39;. | [optional] 

## Methods

### NewADUserOrGroupSummary

`func NewADUserOrGroupSummary(domain string, group bool, guid string, id string, ) *ADUserOrGroupSummary`

NewADUserOrGroupSummary instantiates a new ADUserOrGroupSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADUserOrGroupSummaryWithDefaults

`func NewADUserOrGroupSummaryWithDefaults() *ADUserOrGroupSummary`

NewADUserOrGroupSummaryWithDefaults instantiates a new ADUserOrGroupSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *ADUserOrGroupSummary) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ADUserOrGroupSummary) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ADUserOrGroupSummary) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ADUserOrGroupSummary) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDescription

`func (o *ADUserOrGroupSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ADUserOrGroupSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ADUserOrGroupSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ADUserOrGroupSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ADUserOrGroupSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ADUserOrGroupSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ADUserOrGroupSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ADUserOrGroupSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *ADUserOrGroupSummary) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *ADUserOrGroupSummary) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *ADUserOrGroupSummary) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *ADUserOrGroupSummary) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetDomain

`func (o *ADUserOrGroupSummary) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ADUserOrGroupSummary) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ADUserOrGroupSummary) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEmail

`func (o *ADUserOrGroupSummary) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ADUserOrGroupSummary) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ADUserOrGroupSummary) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ADUserOrGroupSummary) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *ADUserOrGroupSummary) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ADUserOrGroupSummary) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ADUserOrGroupSummary) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ADUserOrGroupSummary) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGroup

`func (o *ADUserOrGroupSummary) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ADUserOrGroupSummary) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ADUserOrGroupSummary) SetGroup(v bool)`

SetGroup sets Group field to given value.


### GetGuid

`func (o *ADUserOrGroupSummary) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ADUserOrGroupSummary) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ADUserOrGroupSummary) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetId

`func (o *ADUserOrGroupSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ADUserOrGroupSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ADUserOrGroupSummary) SetId(v string)`

SetId sets Id field to given value.


### GetKioskUser

`func (o *ADUserOrGroupSummary) GetKioskUser() bool`

GetKioskUser returns the KioskUser field if non-nil, zero value otherwise.

### GetKioskUserOk

`func (o *ADUserOrGroupSummary) GetKioskUserOk() (*bool, bool)`

GetKioskUserOk returns a tuple with the KioskUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKioskUser

`func (o *ADUserOrGroupSummary) SetKioskUser(v bool)`

SetKioskUser sets KioskUser field to given value.

### HasKioskUser

`func (o *ADUserOrGroupSummary) HasKioskUser() bool`

HasKioskUser returns a boolean if a field has been set.

### GetLastName

`func (o *ADUserOrGroupSummary) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ADUserOrGroupSummary) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ADUserOrGroupSummary) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ADUserOrGroupSummary) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLoginName

`func (o *ADUserOrGroupSummary) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *ADUserOrGroupSummary) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *ADUserOrGroupSummary) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *ADUserOrGroupSummary) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLongDisplayName

`func (o *ADUserOrGroupSummary) GetLongDisplayName() string`

GetLongDisplayName returns the LongDisplayName field if non-nil, zero value otherwise.

### GetLongDisplayNameOk

`func (o *ADUserOrGroupSummary) GetLongDisplayNameOk() (*string, bool)`

GetLongDisplayNameOk returns a tuple with the LongDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDisplayName

`func (o *ADUserOrGroupSummary) SetLongDisplayName(v string)`

SetLongDisplayName sets LongDisplayName field to given value.

### HasLongDisplayName

`func (o *ADUserOrGroupSummary) HasLongDisplayName() bool`

HasLongDisplayName returns a boolean if a field has been set.

### GetName

`func (o *ADUserOrGroupSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ADUserOrGroupSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ADUserOrGroupSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ADUserOrGroupSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ADUserOrGroupSummary) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ADUserOrGroupSummary) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ADUserOrGroupSummary) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ADUserOrGroupSummary) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *ADUserOrGroupSummary) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *ADUserOrGroupSummary) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *ADUserOrGroupSummary) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *ADUserOrGroupSummary) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *ADUserOrGroupSummary) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *ADUserOrGroupSummary) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *ADUserOrGroupSummary) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *ADUserOrGroupSummary) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


