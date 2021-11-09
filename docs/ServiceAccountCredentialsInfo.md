# ServiceAccountCredentialsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique SID representing auxiliary account. | [optional] 
**Username** | Pointer to **string** | Auxiliary Service account username. | [optional] 

## Methods

### NewServiceAccountCredentialsInfo

`func NewServiceAccountCredentialsInfo() *ServiceAccountCredentialsInfo`

NewServiceAccountCredentialsInfo instantiates a new ServiceAccountCredentialsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountCredentialsInfoWithDefaults

`func NewServiceAccountCredentialsInfoWithDefaults() *ServiceAccountCredentialsInfo`

NewServiceAccountCredentialsInfoWithDefaults instantiates a new ServiceAccountCredentialsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAccountCredentialsInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountCredentialsInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountCredentialsInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccountCredentialsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ServiceAccountCredentialsInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceAccountCredentialsInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceAccountCredentialsInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServiceAccountCredentialsInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


