# ApplicationIconCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Icon image data in PNG format. This is the Base64 encoded binary data of the image. | 
**Height** | **int64** | Icon height. | 
**Width** | **int64** | Icon width. | 

## Methods

### NewApplicationIconCreateSpec

`func NewApplicationIconCreateSpec(data string, height int64, width int64, ) *ApplicationIconCreateSpec`

NewApplicationIconCreateSpec instantiates a new ApplicationIconCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationIconCreateSpecWithDefaults

`func NewApplicationIconCreateSpecWithDefaults() *ApplicationIconCreateSpec`

NewApplicationIconCreateSpecWithDefaults instantiates a new ApplicationIconCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApplicationIconCreateSpec) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationIconCreateSpec) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationIconCreateSpec) SetData(v string)`

SetData sets Data field to given value.


### GetHeight

`func (o *ApplicationIconCreateSpec) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ApplicationIconCreateSpec) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ApplicationIconCreateSpec) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetWidth

`func (o *ApplicationIconCreateSpec) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ApplicationIconCreateSpec) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ApplicationIconCreateSpec) SetWidth(v int64)`

SetWidth sets Width field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


