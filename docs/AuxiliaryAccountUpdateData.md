# AuxiliaryAccountUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Auxiliary Service account ID. | 
**Password** | **[]string** | Service account user password. | 

## Methods

### NewAuxiliaryAccountUpdateData

`func NewAuxiliaryAccountUpdateData(id string, password []string, ) *AuxiliaryAccountUpdateData`

NewAuxiliaryAccountUpdateData instantiates a new AuxiliaryAccountUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuxiliaryAccountUpdateDataWithDefaults

`func NewAuxiliaryAccountUpdateDataWithDefaults() *AuxiliaryAccountUpdateData`

NewAuxiliaryAccountUpdateDataWithDefaults instantiates a new AuxiliaryAccountUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuxiliaryAccountUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuxiliaryAccountUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuxiliaryAccountUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetPassword

`func (o *AuxiliaryAccountUpdateData) GetPassword() []string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuxiliaryAccountUpdateData) GetPasswordOk() (*[]string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuxiliaryAccountUpdateData) SetPassword(v []string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


