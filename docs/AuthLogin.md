# AuthLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain | 
**Password** | **string** | User password | 
**Username** | **string** | User Name | 

## Methods

### NewAuthLogin

`func NewAuthLogin(domain string, password string, username string, ) *AuthLogin`

NewAuthLogin instantiates a new AuthLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLoginWithDefaults

`func NewAuthLoginWithDefaults() *AuthLogin`

NewAuthLoginWithDefaults instantiates a new AuthLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *AuthLogin) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AuthLogin) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AuthLogin) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPassword

`func (o *AuthLogin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthLogin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthLogin) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *AuthLogin) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthLogin) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthLogin) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


