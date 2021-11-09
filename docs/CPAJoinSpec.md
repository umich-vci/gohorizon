# CPAJoinSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **[]string** | The password for the user. | 
**RemotePodAddress** | **string** | The IP address or hostname for the remote pod. | 
**Username** | **string** | The user name, along with domain name, with sufficient privilege to perform a global LDAP join against the remote pod. The down-level logon name format (domain\\username) is allowed. | 

## Methods

### NewCPAJoinSpec

`func NewCPAJoinSpec(password []string, remotePodAddress string, username string, ) *CPAJoinSpec`

NewCPAJoinSpec instantiates a new CPAJoinSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCPAJoinSpecWithDefaults

`func NewCPAJoinSpecWithDefaults() *CPAJoinSpec`

NewCPAJoinSpecWithDefaults instantiates a new CPAJoinSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *CPAJoinSpec) GetPassword() []string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CPAJoinSpec) GetPasswordOk() (*[]string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CPAJoinSpec) SetPassword(v []string)`

SetPassword sets Password field to given value.


### GetRemotePodAddress

`func (o *CPAJoinSpec) GetRemotePodAddress() string`

GetRemotePodAddress returns the RemotePodAddress field if non-nil, zero value otherwise.

### GetRemotePodAddressOk

`func (o *CPAJoinSpec) GetRemotePodAddressOk() (*string, bool)`

GetRemotePodAddressOk returns a tuple with the RemotePodAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePodAddress

`func (o *CPAJoinSpec) SetRemotePodAddress(v string)`

SetRemotePodAddress sets RemotePodAddress field to given value.


### GetUsername

`func (o *CPAJoinSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CPAJoinSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CPAJoinSpec) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


