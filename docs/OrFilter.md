# OrFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]BaseFilter**](BaseFilter.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOrFilter

`func NewOrFilter() *OrFilter`

NewOrFilter instantiates a new OrFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrFilterWithDefaults

`func NewOrFilterWithDefaults() *OrFilter`

NewOrFilterWithDefaults instantiates a new OrFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *OrFilter) GetFilters() []BaseFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OrFilter) GetFiltersOk() (*[]BaseFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OrFilter) SetFilters(v []BaseFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OrFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetType

`func (o *OrFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrFilter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


