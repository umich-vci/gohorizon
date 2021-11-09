# ImageManagementStreamInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **map[string]string** | Additional details about image management stream. | [optional] 
**Description** | Pointer to **string** | Image management stream description. | [optional] 
**Id** | **string** | Unique ID representing image management stream. | 
**Name** | **string** | Image management stream name. | 
**OperatingSystem** | **string** | Operating system. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS) | 
**Publisher** | Pointer to **string** | Image management stream publisher | [optional] 
**Source** | **string** | Image management stream source. * MARKET_PLACE: Image management stream is from market place. * UPLOADED: Image management stream is uploaded. * COPIED_FROM_STREAM: Image management stream is copied from another stream. * COPIED_FROM_VERSION: Image management stream is copied from a version. | 
**Status** | **string** | Image management stream status. * AVAILABLE: Image management stream is available for desktop pools/farms to be created. * DELETED: Image management stream is deleted. * DISABLED: Image management stream is disabled and no further desktop pools/farms can be created using the same. * FAILED: Image management stream creation has failed. * IN_PROGRESS: Image management stream creation is in progress. * PARTIALLY_AVAILABLE: Image management version for this stream could not be created in one or more environments. * PENDING: Image management stream is in pending state. | 
**Usable** | Pointer to **bool** | Specifies whether the image management stream can be used in desktop pool or farm. This will be set to true when: &lt;ul&gt;&lt;li&gt;Image management stream is in AVAILABLE or PARTIALLY_AVAILABLE state. &lt;/li&gt;&lt;li&gt;There is at least one image management version in AVAILABLE or PARTIALLY_AVAILABLE state for this stream. &lt;/li&gt;&lt;li&gt;There is at least one image management tag associated with the image management version.&lt;/li&gt; &lt;/ul&gt;For a specific virtual center, image management tag information will be retrieved. | [optional] 

## Methods

### NewImageManagementStreamInfo

`func NewImageManagementStreamInfo(id string, name string, operatingSystem string, source string, status string, ) *ImageManagementStreamInfo`

NewImageManagementStreamInfo instantiates a new ImageManagementStreamInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageManagementStreamInfoWithDefaults

`func NewImageManagementStreamInfoWithDefaults() *ImageManagementStreamInfo`

NewImageManagementStreamInfoWithDefaults instantiates a new ImageManagementStreamInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *ImageManagementStreamInfo) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ImageManagementStreamInfo) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ImageManagementStreamInfo) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ImageManagementStreamInfo) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetDescription

`func (o *ImageManagementStreamInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageManagementStreamInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageManagementStreamInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageManagementStreamInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ImageManagementStreamInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageManagementStreamInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageManagementStreamInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ImageManagementStreamInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageManagementStreamInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageManagementStreamInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOperatingSystem

`func (o *ImageManagementStreamInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ImageManagementStreamInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ImageManagementStreamInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetPublisher

`func (o *ImageManagementStreamInfo) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ImageManagementStreamInfo) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ImageManagementStreamInfo) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ImageManagementStreamInfo) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetSource

`func (o *ImageManagementStreamInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ImageManagementStreamInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ImageManagementStreamInfo) SetSource(v string)`

SetSource sets Source field to given value.


### GetStatus

`func (o *ImageManagementStreamInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageManagementStreamInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageManagementStreamInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUsable

`func (o *ImageManagementStreamInfo) GetUsable() bool`

GetUsable returns the Usable field if non-nil, zero value otherwise.

### GetUsableOk

`func (o *ImageManagementStreamInfo) GetUsableOk() (*bool, bool)`

GetUsableOk returns a tuple with the Usable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsable

`func (o *ImageManagementStreamInfo) SetUsable(v bool)`

SetUsable sets Usable field to given value.

### HasUsable

`func (o *ImageManagementStreamInfo) HasUsable() bool`

HasUsable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


