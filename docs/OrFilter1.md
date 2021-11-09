# OrFilter1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]BaseFilter**](BaseFilter.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOrFilter1

`func NewOrFilter1() *OrFilter1`

NewOrFilter1 instantiates a new OrFilter1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrFilter1WithDefaults

`func NewOrFilter1WithDefaults() *OrFilter1`

NewOrFilter1WithDefaults instantiates a new OrFilter1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *OrFilter1) GetFilters() []BaseFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OrFilter1) GetFiltersOk() (*[]BaseFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OrFilter1) SetFilters(v []BaseFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OrFilter1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetType

`func (o *OrFilter1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrFilter1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrFilter1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrFilter1) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


