# VMFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]VMFolderInfo**](VMFolderInfo.md) | Child nodes of the VM folder. | [optional] 
**DatacenterId** | Pointer to **string** | Datacenter id for this VM folder. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the VM folder. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this VM folder from being used in desktop pool or farm. | [optional] 
**Name** | Pointer to **string** | VM folder name. | [optional] 
**Path** | Pointer to **string** | VM folder path. | [optional] 
**Type** | Pointer to **string** | VM folder type. * DATACENTER: A datacenter that serves as a folder suitable for use in desktop pool/farm. * FOLDER: A regular folder suitable for use in desktop pool/farm. * OTHER: Other folder type that cannot be used in desktop pool/farm. | [optional] 
**VcenterId** | Pointer to **string** | Virtual Center id for this VM folder. | [optional] 

## Methods

### NewVMFolderInfo

`func NewVMFolderInfo() *VMFolderInfo`

NewVMFolderInfo instantiates a new VMFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMFolderInfoWithDefaults

`func NewVMFolderInfoWithDefaults() *VMFolderInfo`

NewVMFolderInfoWithDefaults instantiates a new VMFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *VMFolderInfo) GetChildren() []VMFolderInfo`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *VMFolderInfo) GetChildrenOk() (*[]VMFolderInfo, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *VMFolderInfo) SetChildren(v []VMFolderInfo)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *VMFolderInfo) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDatacenterId

`func (o *VMFolderInfo) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *VMFolderInfo) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *VMFolderInfo) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *VMFolderInfo) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetId

`func (o *VMFolderInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMFolderInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMFolderInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VMFolderInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *VMFolderInfo) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *VMFolderInfo) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *VMFolderInfo) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *VMFolderInfo) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetName

`func (o *VMFolderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMFolderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMFolderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VMFolderInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *VMFolderInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VMFolderInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VMFolderInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VMFolderInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *VMFolderInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VMFolderInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VMFolderInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VMFolderInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVcenterId

`func (o *VMFolderInfo) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *VMFolderInfo) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *VMFolderInfo) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *VMFolderInfo) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


