# AuthTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Access Token to be used in API calls. | [optional] 
**RefreshToken** | Pointer to **string** | Refresh Token to be used to get a new Access token. | [optional] 

## Methods

### NewAuthTokens

`func NewAuthTokens() *AuthTokens`

NewAuthTokens instantiates a new AuthTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokensWithDefaults

`func NewAuthTokensWithDefaults() *AuthTokens`

NewAuthTokensWithDefaults instantiates a new AuthTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AuthTokens) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthTokens) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthTokens) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthTokens) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *AuthTokens) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AuthTokens) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AuthTokens) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AuthTokens) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


