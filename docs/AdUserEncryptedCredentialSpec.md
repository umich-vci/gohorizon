# ADUserEncryptedCredentialSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The domain of user. Note that domain is optional if UPN is supplied. | [optional] 
**EncryptedPassword** | **string** | Encrypted password for the user. | 
**KeyId** | **string** | The keyId of the cluster&#39;s SSO KeyPair used to encrypt the protectedPasswordKey. | 
**ProtectedPasswordKey** | **string** | Decryption key for the password. This key is itself encrypted with cluster&#39;s SSO keypair. | 
**Username** | **string** | The username or UPN. | 

## Methods

### NewADUserEncryptedCredentialSpec

`func NewADUserEncryptedCredentialSpec(encryptedPassword string, keyId string, protectedPasswordKey string, username string, ) *ADUserEncryptedCredentialSpec`

NewADUserEncryptedCredentialSpec instantiates a new ADUserEncryptedCredentialSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADUserEncryptedCredentialSpecWithDefaults

`func NewADUserEncryptedCredentialSpecWithDefaults() *ADUserEncryptedCredentialSpec`

NewADUserEncryptedCredentialSpecWithDefaults instantiates a new ADUserEncryptedCredentialSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ADUserEncryptedCredentialSpec) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ADUserEncryptedCredentialSpec) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ADUserEncryptedCredentialSpec) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ADUserEncryptedCredentialSpec) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *ADUserEncryptedCredentialSpec) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *ADUserEncryptedCredentialSpec) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *ADUserEncryptedCredentialSpec) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.


### GetKeyId

`func (o *ADUserEncryptedCredentialSpec) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ADUserEncryptedCredentialSpec) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ADUserEncryptedCredentialSpec) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetProtectedPasswordKey

`func (o *ADUserEncryptedCredentialSpec) GetProtectedPasswordKey() string`

GetProtectedPasswordKey returns the ProtectedPasswordKey field if non-nil, zero value otherwise.

### GetProtectedPasswordKeyOk

`func (o *ADUserEncryptedCredentialSpec) GetProtectedPasswordKeyOk() (*string, bool)`

GetProtectedPasswordKeyOk returns a tuple with the ProtectedPasswordKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedPasswordKey

`func (o *ADUserEncryptedCredentialSpec) SetProtectedPasswordKey(v string)`

SetProtectedPasswordKey sets ProtectedPasswordKey field to given value.


### GetUsername

`func (o *ADUserEncryptedCredentialSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ADUserEncryptedCredentialSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ADUserEncryptedCredentialSpec) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


