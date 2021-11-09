# FarmDisplayProtocolSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | **bool** | Indicates whether the users can choose the protocol. | 
**DefaultDisplayProtocol** | **string** | Indicates default server display protocol, when user is not allowed to choose protocol.Farms support PCOIP, RDP and BLAST. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | 
**SessionCollaborationEnabled** | **bool** | Indicates whether session collaboration feature is enabled. Session collaboration allows a user to share their remote session with other users. | 

## Methods

### NewFarmDisplayProtocolSettingsUpdateSpec

`func NewFarmDisplayProtocolSettingsUpdateSpec(allowUsersToChooseProtocol bool, defaultDisplayProtocol string, sessionCollaborationEnabled bool, ) *FarmDisplayProtocolSettingsUpdateSpec`

NewFarmDisplayProtocolSettingsUpdateSpec instantiates a new FarmDisplayProtocolSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmDisplayProtocolSettingsUpdateSpecWithDefaults

`func NewFarmDisplayProtocolSettingsUpdateSpecWithDefaults() *FarmDisplayProtocolSettingsUpdateSpec`

NewFarmDisplayProtocolSettingsUpdateSpecWithDefaults instantiates a new FarmDisplayProtocolSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsersToChooseProtocol

`func (o *FarmDisplayProtocolSettingsUpdateSpec) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *FarmDisplayProtocolSettingsUpdateSpec) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *FarmDisplayProtocolSettingsUpdateSpec) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.


### GetDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettingsUpdateSpec) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *FarmDisplayProtocolSettingsUpdateSpec) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettingsUpdateSpec) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.


### GetSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettingsUpdateSpec) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *FarmDisplayProtocolSettingsUpdateSpec) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettingsUpdateSpec) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


