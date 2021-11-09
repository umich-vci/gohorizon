# ApplicationIconInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationPoolIds** | Pointer to **[]string** | An application icon could be shared by multiple application pools. This is a set of application pool IDs this icon represents. Caller should have permission to POOL_VIEW privilege on the access group of any of the associated application pools for this field to be populated.&lt;br&gt; | [optional] 
**Data** | Pointer to **string** | Icon image data in PNG format. | [optional] 
**Height** | Pointer to **int32** | Icon height. | [optional] 
**IconHash** | Pointer to **string** | MD5 hash of icon image data, to enable quick icon lookup. | [optional] 
**IconSource** | Pointer to **string** | Source of the application icon. The icon can be from Machine/RDS Agent or custom icon. * HORIZON_AGENT: When the icon is from machine/RDS agent. * HORIZON_CONNECTION_SERVER: When the icon is uploaded by the administrator through connection server. | [optional] 
**Id** | Pointer to **string** | Unique ID representing application icon. | [optional] 
**Width** | Pointer to **int32** | Icon width. | [optional] 

## Methods

### NewApplicationIconInfo

`func NewApplicationIconInfo() *ApplicationIconInfo`

NewApplicationIconInfo instantiates a new ApplicationIconInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationIconInfoWithDefaults

`func NewApplicationIconInfoWithDefaults() *ApplicationIconInfo`

NewApplicationIconInfoWithDefaults instantiates a new ApplicationIconInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationPoolIds

`func (o *ApplicationIconInfo) GetApplicationPoolIds() []string`

GetApplicationPoolIds returns the ApplicationPoolIds field if non-nil, zero value otherwise.

### GetApplicationPoolIdsOk

`func (o *ApplicationIconInfo) GetApplicationPoolIdsOk() (*[]string, bool)`

GetApplicationPoolIdsOk returns a tuple with the ApplicationPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPoolIds

`func (o *ApplicationIconInfo) SetApplicationPoolIds(v []string)`

SetApplicationPoolIds sets ApplicationPoolIds field to given value.

### HasApplicationPoolIds

`func (o *ApplicationIconInfo) HasApplicationPoolIds() bool`

HasApplicationPoolIds returns a boolean if a field has been set.

### GetData

`func (o *ApplicationIconInfo) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationIconInfo) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationIconInfo) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ApplicationIconInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeight

`func (o *ApplicationIconInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ApplicationIconInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ApplicationIconInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ApplicationIconInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIconHash

`func (o *ApplicationIconInfo) GetIconHash() string`

GetIconHash returns the IconHash field if non-nil, zero value otherwise.

### GetIconHashOk

`func (o *ApplicationIconInfo) GetIconHashOk() (*string, bool)`

GetIconHashOk returns a tuple with the IconHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconHash

`func (o *ApplicationIconInfo) SetIconHash(v string)`

SetIconHash sets IconHash field to given value.

### HasIconHash

`func (o *ApplicationIconInfo) HasIconHash() bool`

HasIconHash returns a boolean if a field has been set.

### GetIconSource

`func (o *ApplicationIconInfo) GetIconSource() string`

GetIconSource returns the IconSource field if non-nil, zero value otherwise.

### GetIconSourceOk

`func (o *ApplicationIconInfo) GetIconSourceOk() (*string, bool)`

GetIconSourceOk returns a tuple with the IconSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSource

`func (o *ApplicationIconInfo) SetIconSource(v string)`

SetIconSource sets IconSource field to given value.

### HasIconSource

`func (o *ApplicationIconInfo) HasIconSource() bool`

HasIconSource returns a boolean if a field has been set.

### GetId

`func (o *ApplicationIconInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationIconInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationIconInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationIconInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWidth

`func (o *ApplicationIconInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ApplicationIconInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ApplicationIconInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ApplicationIconInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


