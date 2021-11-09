# NotFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**BaseFilter**](BaseFilter.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewNotFilter

`func NewNotFilter() *NotFilter`

NewNotFilter instantiates a new NotFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFilterWithDefaults

`func NewNotFilterWithDefaults() *NotFilter`

NewNotFilterWithDefaults instantiates a new NotFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *NotFilter) GetFilter() BaseFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *NotFilter) GetFilterOk() (*BaseFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *NotFilter) SetFilter(v BaseFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *NotFilter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetType

`func (o *NotFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotFilter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


