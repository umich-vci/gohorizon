/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2106
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// ADUserOrGroupInfo Information related to AD User or Group.
type ADUserOrGroupInfo struct {
	// AD container for this user or group.
	Container *string `json:"container,omitempty"`
	// Description number of this user or group. Supported Filters : 'Equals', 'StartsWith', 'Contains'.
	Description *string `json:"description,omitempty"`
	// Login name with domain of this user or group.
	DisplayName *string `json:"display_name,omitempty"`
	// Active Directory distinguished name for this user or group.
	DistinguishedName *string `json:"distinguished_name,omitempty"`
	// DNS name of the domain in which this user or group belongs. Supported Filters : 'Equals'.  Also, if 'Or' filter is used anywhere in filter string for this model class, then that 'Or' filter should nest only 'Equals' filter on 'domain' or 'id' field.
	Domain string `json:"domain"`
	// Email address of this user or group. Supported Filters : 'Equals', 'StartsWith', 'Contains'.
	Email *string `json:"email,omitempty"`
	// First name of this user or group.
	FirstName *string `json:"first_name,omitempty"`
	// Indicates if this object represents a group. This field is NOT supported in filter string. To use any filter on 'group', use 'group_only' query param directly.
	Group bool `json:"group"`
	// List of unique SIDs of the groups, this user or group belongs to.
	GroupMembershipIds *[]string `json:"group_membership_ids,omitempty"`
	// GUID of the user or group in RFC 4122 format. Supported Filters : 'Equals'.
	Guid string `json:"guid"`
	// Unique SID representing this AD User or Group. Supported Filters : 'Equals'.'Or' filter chain of 'Equals' filter can be used to query for more than one id. For this model, if 'Or' filter is used, then it should nest only 'Equals' filter on 'domain' or 'id' field.
	Id string `json:"id"`
	// Indicates if this user or group is a \"kiosk user\" that supports client authentication. Client authentication is the process of supporting client devices directly logging into resources.
	KioskUser *bool `json:"kiosk_user,omitempty"`
	// Last name of this user or group.
	LastName *string `json:"last_name,omitempty"`
	// Login name of this user or group.  Supported Filters : 'Equals', 'StartsWith', 'Contains'.
	LoginName *string `json:"login_name,omitempty"`
	// Login name, domain and name for this user or group, else display name
	LongDisplayName *string `json:"long_display_name,omitempty"`
	// Name of this user or group.  Supported Filters : 'Equals', 'StartsWith', 'Contains'.
	Name *string `json:"name,omitempty"`
	// Phone number of this user. Supported Filters : 'Equals', 'StartsWith', 'Contains'.
	Phone *string `json:"phone,omitempty"`
	// Number of subgroups in this group, or 0 if not a group.
	SubGroupCount *int32 `json:"sub_group_count,omitempty"`
	// Number of users in this group, or 0 if not a group.
	UserCount *int32 `json:"user_count,omitempty"`
	// User or group's display name. This corresponds with displayName attribute in AD.
	UserDisplayName *string `json:"user_display_name,omitempty"`
	// User Principal name(UPN) of this user. Supported Filters : 'Equals', 'StartsWith', 'Contains'.
	UserPrincipalName *string `json:"user_principal_name,omitempty"`
}

// NewADUserOrGroupInfo instantiates a new ADUserOrGroupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewADUserOrGroupInfo(domain string, group bool, guid string, id string) *ADUserOrGroupInfo {
	this := ADUserOrGroupInfo{}
	this.Domain = domain
	this.Group = group
	this.Guid = guid
	this.Id = id
	return &this
}

// NewADUserOrGroupInfoWithDefaults instantiates a new ADUserOrGroupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewADUserOrGroupInfoWithDefaults() *ADUserOrGroupInfo {
	this := ADUserOrGroupInfo{}
	return &this
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetContainerOk() (*string, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *ADUserOrGroupInfo) SetContainer(v string) {
	o.Container = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ADUserOrGroupInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ADUserOrGroupInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetDistinguishedName() string {
	if o == nil || o.DistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || o.DistinguishedName == nil {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasDistinguishedName() bool {
	if o != nil && o.DistinguishedName != nil {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *ADUserOrGroupInfo) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetDomain returns the Domain field value
func (o *ADUserOrGroupInfo) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ADUserOrGroupInfo) SetDomain(v string) {
	o.Domain = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ADUserOrGroupInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ADUserOrGroupInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetGroup returns the Group field value
func (o *ADUserOrGroupInfo) GetGroup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetGroupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ADUserOrGroupInfo) SetGroup(v bool) {
	o.Group = v
}

// GetGroupMembershipIds returns the GroupMembershipIds field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetGroupMembershipIds() []string {
	if o == nil || o.GroupMembershipIds == nil {
		var ret []string
		return ret
	}
	return *o.GroupMembershipIds
}

// GetGroupMembershipIdsOk returns a tuple with the GroupMembershipIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetGroupMembershipIdsOk() (*[]string, bool) {
	if o == nil || o.GroupMembershipIds == nil {
		return nil, false
	}
	return o.GroupMembershipIds, true
}

// HasGroupMembershipIds returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasGroupMembershipIds() bool {
	if o != nil && o.GroupMembershipIds != nil {
		return true
	}

	return false
}

// SetGroupMembershipIds gets a reference to the given []string and assigns it to the GroupMembershipIds field.
func (o *ADUserOrGroupInfo) SetGroupMembershipIds(v []string) {
	o.GroupMembershipIds = &v
}

// GetGuid returns the Guid field value
func (o *ADUserOrGroupInfo) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *ADUserOrGroupInfo) SetGuid(v string) {
	o.Guid = v
}

// GetId returns the Id field value
func (o *ADUserOrGroupInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ADUserOrGroupInfo) SetId(v string) {
	o.Id = v
}

// GetKioskUser returns the KioskUser field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetKioskUser() bool {
	if o == nil || o.KioskUser == nil {
		var ret bool
		return ret
	}
	return *o.KioskUser
}

// GetKioskUserOk returns a tuple with the KioskUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetKioskUserOk() (*bool, bool) {
	if o == nil || o.KioskUser == nil {
		return nil, false
	}
	return o.KioskUser, true
}

// HasKioskUser returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasKioskUser() bool {
	if o != nil && o.KioskUser != nil {
		return true
	}

	return false
}

// SetKioskUser gets a reference to the given bool and assigns it to the KioskUser field.
func (o *ADUserOrGroupInfo) SetKioskUser(v bool) {
	o.KioskUser = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ADUserOrGroupInfo) SetLastName(v string) {
	o.LastName = &v
}

// GetLoginName returns the LoginName field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetLoginName() string {
	if o == nil || o.LoginName == nil {
		var ret string
		return ret
	}
	return *o.LoginName
}

// GetLoginNameOk returns a tuple with the LoginName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetLoginNameOk() (*string, bool) {
	if o == nil || o.LoginName == nil {
		return nil, false
	}
	return o.LoginName, true
}

// HasLoginName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasLoginName() bool {
	if o != nil && o.LoginName != nil {
		return true
	}

	return false
}

// SetLoginName gets a reference to the given string and assigns it to the LoginName field.
func (o *ADUserOrGroupInfo) SetLoginName(v string) {
	o.LoginName = &v
}

// GetLongDisplayName returns the LongDisplayName field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetLongDisplayName() string {
	if o == nil || o.LongDisplayName == nil {
		var ret string
		return ret
	}
	return *o.LongDisplayName
}

// GetLongDisplayNameOk returns a tuple with the LongDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetLongDisplayNameOk() (*string, bool) {
	if o == nil || o.LongDisplayName == nil {
		return nil, false
	}
	return o.LongDisplayName, true
}

// HasLongDisplayName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasLongDisplayName() bool {
	if o != nil && o.LongDisplayName != nil {
		return true
	}

	return false
}

// SetLongDisplayName gets a reference to the given string and assigns it to the LongDisplayName field.
func (o *ADUserOrGroupInfo) SetLongDisplayName(v string) {
	o.LongDisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ADUserOrGroupInfo) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ADUserOrGroupInfo) SetPhone(v string) {
	o.Phone = &v
}

// GetSubGroupCount returns the SubGroupCount field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetSubGroupCount() int32 {
	if o == nil || o.SubGroupCount == nil {
		var ret int32
		return ret
	}
	return *o.SubGroupCount
}

// GetSubGroupCountOk returns a tuple with the SubGroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetSubGroupCountOk() (*int32, bool) {
	if o == nil || o.SubGroupCount == nil {
		return nil, false
	}
	return o.SubGroupCount, true
}

// HasSubGroupCount returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasSubGroupCount() bool {
	if o != nil && o.SubGroupCount != nil {
		return true
	}

	return false
}

// SetSubGroupCount gets a reference to the given int32 and assigns it to the SubGroupCount field.
func (o *ADUserOrGroupInfo) SetSubGroupCount(v int32) {
	o.SubGroupCount = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetUserCount() int32 {
	if o == nil || o.UserCount == nil {
		var ret int32
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetUserCountOk() (*int32, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasUserCount() bool {
	if o != nil && o.UserCount != nil {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.
func (o *ADUserOrGroupInfo) SetUserCount(v int32) {
	o.UserCount = &v
}

// GetUserDisplayName returns the UserDisplayName field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetUserDisplayName() string {
	if o == nil || o.UserDisplayName == nil {
		var ret string
		return ret
	}
	return *o.UserDisplayName
}

// GetUserDisplayNameOk returns a tuple with the UserDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetUserDisplayNameOk() (*string, bool) {
	if o == nil || o.UserDisplayName == nil {
		return nil, false
	}
	return o.UserDisplayName, true
}

// HasUserDisplayName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasUserDisplayName() bool {
	if o != nil && o.UserDisplayName != nil {
		return true
	}

	return false
}

// SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.
func (o *ADUserOrGroupInfo) SetUserDisplayName(v string) {
	o.UserDisplayName = &v
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise.
func (o *ADUserOrGroupInfo) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUserOrGroupInfo) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil || o.UserPrincipalName == nil {
		return nil, false
	}
	return o.UserPrincipalName, true
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *ADUserOrGroupInfo) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName != nil {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.
func (o *ADUserOrGroupInfo) SetUserPrincipalName(v string) {
	o.UserPrincipalName = &v
}

func (o ADUserOrGroupInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.DistinguishedName != nil {
		toSerialize["distinguished_name"] = o.DistinguishedName
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if o.GroupMembershipIds != nil {
		toSerialize["group_membership_ids"] = o.GroupMembershipIds
	}
	if true {
		toSerialize["guid"] = o.Guid
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.KioskUser != nil {
		toSerialize["kiosk_user"] = o.KioskUser
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.LoginName != nil {
		toSerialize["login_name"] = o.LoginName
	}
	if o.LongDisplayName != nil {
		toSerialize["long_display_name"] = o.LongDisplayName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.SubGroupCount != nil {
		toSerialize["sub_group_count"] = o.SubGroupCount
	}
	if o.UserCount != nil {
		toSerialize["user_count"] = o.UserCount
	}
	if o.UserDisplayName != nil {
		toSerialize["user_display_name"] = o.UserDisplayName
	}
	if o.UserPrincipalName != nil {
		toSerialize["user_principal_name"] = o.UserPrincipalName
	}
	return json.Marshal(toSerialize)
}

type NullableADUserOrGroupInfo struct {
	value *ADUserOrGroupInfo
	isSet bool
}

func (v NullableADUserOrGroupInfo) Get() *ADUserOrGroupInfo {
	return v.value
}

func (v *NullableADUserOrGroupInfo) Set(val *ADUserOrGroupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableADUserOrGroupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableADUserOrGroupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableADUserOrGroupInfo(val *ADUserOrGroupInfo) *NullableADUserOrGroupInfo {
	return &NullableADUserOrGroupInfo{value: val, isSet: true}
}

func (v NullableADUserOrGroupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableADUserOrGroupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
