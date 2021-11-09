# ServiceAccountCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **[]string** | Service account user password. | 
**Username** | **string** | Service account username. | 

## Methods

### NewServiceAccountCredentials

`func NewServiceAccountCredentials(password []string, username string, ) *ServiceAccountCredentials`

NewServiceAccountCredentials instantiates a new ServiceAccountCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountCredentialsWithDefaults

`func NewServiceAccountCredentialsWithDefaults() *ServiceAccountCredentials`

NewServiceAccountCredentialsWithDefaults instantiates a new ServiceAccountCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ServiceAccountCredentials) GetPassword() []string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServiceAccountCredentials) GetPasswordOk() (*[]string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServiceAccountCredentials) SetPassword(v []string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *ServiceAccountCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceAccountCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceAccountCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


