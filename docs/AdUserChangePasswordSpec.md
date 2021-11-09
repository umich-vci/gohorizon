# ADUserChangePasswordSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The domain of user. Note that domain is optional if UPN is supplied. | [optional] 
**KeyId** | **string** | The keyId of the cluster&#39;s SSO KeyPair used to encrypt the password key. | 
**NewEncryptedPassword** | **string** | New password for the user in encrypted format. | 
**OldEncryptedPassword** | **string** | Old password for the user in encrypted format. | 
**ProtectedPasswordKey** | **string** | Decryption key for the password. This key is itself encrypted with cluster&#39;s SSO keypair. | 
**Username** | **string** | The username or UPN. | 

## Methods

### NewADUserChangePasswordSpec

`func NewADUserChangePasswordSpec(keyId string, newEncryptedPassword string, oldEncryptedPassword string, protectedPasswordKey string, username string, ) *ADUserChangePasswordSpec`

NewADUserChangePasswordSpec instantiates a new ADUserChangePasswordSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADUserChangePasswordSpecWithDefaults

`func NewADUserChangePasswordSpecWithDefaults() *ADUserChangePasswordSpec`

NewADUserChangePasswordSpecWithDefaults instantiates a new ADUserChangePasswordSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ADUserChangePasswordSpec) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ADUserChangePasswordSpec) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ADUserChangePasswordSpec) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ADUserChangePasswordSpec) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetKeyId

`func (o *ADUserChangePasswordSpec) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ADUserChangePasswordSpec) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ADUserChangePasswordSpec) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetNewEncryptedPassword

`func (o *ADUserChangePasswordSpec) GetNewEncryptedPassword() string`

GetNewEncryptedPassword returns the NewEncryptedPassword field if non-nil, zero value otherwise.

### GetNewEncryptedPasswordOk

`func (o *ADUserChangePasswordSpec) GetNewEncryptedPasswordOk() (*string, bool)`

GetNewEncryptedPasswordOk returns a tuple with the NewEncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEncryptedPassword

`func (o *ADUserChangePasswordSpec) SetNewEncryptedPassword(v string)`

SetNewEncryptedPassword sets NewEncryptedPassword field to given value.


### GetOldEncryptedPassword

`func (o *ADUserChangePasswordSpec) GetOldEncryptedPassword() string`

GetOldEncryptedPassword returns the OldEncryptedPassword field if non-nil, zero value otherwise.

### GetOldEncryptedPasswordOk

`func (o *ADUserChangePasswordSpec) GetOldEncryptedPasswordOk() (*string, bool)`

GetOldEncryptedPasswordOk returns a tuple with the OldEncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldEncryptedPassword

`func (o *ADUserChangePasswordSpec) SetOldEncryptedPassword(v string)`

SetOldEncryptedPassword sets OldEncryptedPassword field to given value.


### GetProtectedPasswordKey

`func (o *ADUserChangePasswordSpec) GetProtectedPasswordKey() string`

GetProtectedPasswordKey returns the ProtectedPasswordKey field if non-nil, zero value otherwise.

### GetProtectedPasswordKeyOk

`func (o *ADUserChangePasswordSpec) GetProtectedPasswordKeyOk() (*string, bool)`

GetProtectedPasswordKeyOk returns a tuple with the ProtectedPasswordKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedPasswordKey

`func (o *ADUserChangePasswordSpec) SetProtectedPasswordKey(v string)`

SetProtectedPasswordKey sets ProtectedPasswordKey field to given value.


### GetUsername

`func (o *ADUserChangePasswordSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ADUserChangePasswordSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ADUserChangePasswordSpec) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


