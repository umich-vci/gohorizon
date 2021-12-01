# DesktopPoolDeleteSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFromDisk** | Pointer to **bool** | Determines whether the machine VMs should be deleted from vCenter Server.&lt;br&gt; This must be false for RDS and unmanaged desktop pools and true for Instant Clone desktop pools.&lt;br&gt;Default value is true for IC pools and false for all other types of desktop pools. &lt;br&gt; | [optional] 

## Methods

### NewDesktopPoolDeleteSpec

`func NewDesktopPoolDeleteSpec() *DesktopPoolDeleteSpec`

NewDesktopPoolDeleteSpec instantiates a new DesktopPoolDeleteSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolDeleteSpecWithDefaults

`func NewDesktopPoolDeleteSpecWithDefaults() *DesktopPoolDeleteSpec`

NewDesktopPoolDeleteSpecWithDefaults instantiates a new DesktopPoolDeleteSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFromDisk

`func (o *DesktopPoolDeleteSpec) GetDeleteFromDisk() bool`

GetDeleteFromDisk returns the DeleteFromDisk field if non-nil, zero value otherwise.

### GetDeleteFromDiskOk

`func (o *DesktopPoolDeleteSpec) GetDeleteFromDiskOk() (*bool, bool)`

GetDeleteFromDiskOk returns a tuple with the DeleteFromDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFromDisk

`func (o *DesktopPoolDeleteSpec) SetDeleteFromDisk(v bool)`

SetDeleteFromDisk sets DeleteFromDisk field to given value.

### HasDeleteFromDisk

`func (o *DesktopPoolDeleteSpec) HasDeleteFromDisk() bool`

HasDeleteFromDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


