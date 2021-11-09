# ViewComposerMonitorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVerion** | **string** | The version of the View Composer API used to communicate with the View Composer server. | 
**Build** | **string** | The build of the View Composer server. | 
**MinEsxVersion** | **string** | The minimum ESX version required for compatibility with this View Composer server. | 
**MinVcVersion** | **string** | The minimum Virtual Center version required for compatibility with this View Composer server. | 
**ReferencedVcs** | **[]string** | The Virtual Center servers referencing to this View Composer. | 
**Version** | **string** | The version of the View Composer server. | 

## Methods

### NewViewComposerMonitorDetails

`func NewViewComposerMonitorDetails(apiVerion string, build string, minEsxVersion string, minVcVersion string, referencedVcs []string, version string, ) *ViewComposerMonitorDetails`

NewViewComposerMonitorDetails instantiates a new ViewComposerMonitorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewComposerMonitorDetailsWithDefaults

`func NewViewComposerMonitorDetailsWithDefaults() *ViewComposerMonitorDetails`

NewViewComposerMonitorDetailsWithDefaults instantiates a new ViewComposerMonitorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVerion

`func (o *ViewComposerMonitorDetails) GetApiVerion() string`

GetApiVerion returns the ApiVerion field if non-nil, zero value otherwise.

### GetApiVerionOk

`func (o *ViewComposerMonitorDetails) GetApiVerionOk() (*string, bool)`

GetApiVerionOk returns a tuple with the ApiVerion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVerion

`func (o *ViewComposerMonitorDetails) SetApiVerion(v string)`

SetApiVerion sets ApiVerion field to given value.


### GetBuild

`func (o *ViewComposerMonitorDetails) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ViewComposerMonitorDetails) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ViewComposerMonitorDetails) SetBuild(v string)`

SetBuild sets Build field to given value.


### GetMinEsxVersion

`func (o *ViewComposerMonitorDetails) GetMinEsxVersion() string`

GetMinEsxVersion returns the MinEsxVersion field if non-nil, zero value otherwise.

### GetMinEsxVersionOk

`func (o *ViewComposerMonitorDetails) GetMinEsxVersionOk() (*string, bool)`

GetMinEsxVersionOk returns a tuple with the MinEsxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEsxVersion

`func (o *ViewComposerMonitorDetails) SetMinEsxVersion(v string)`

SetMinEsxVersion sets MinEsxVersion field to given value.


### GetMinVcVersion

`func (o *ViewComposerMonitorDetails) GetMinVcVersion() string`

GetMinVcVersion returns the MinVcVersion field if non-nil, zero value otherwise.

### GetMinVcVersionOk

`func (o *ViewComposerMonitorDetails) GetMinVcVersionOk() (*string, bool)`

GetMinVcVersionOk returns a tuple with the MinVcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVcVersion

`func (o *ViewComposerMonitorDetails) SetMinVcVersion(v string)`

SetMinVcVersion sets MinVcVersion field to given value.


### GetReferencedVcs

`func (o *ViewComposerMonitorDetails) GetReferencedVcs() []string`

GetReferencedVcs returns the ReferencedVcs field if non-nil, zero value otherwise.

### GetReferencedVcsOk

`func (o *ViewComposerMonitorDetails) GetReferencedVcsOk() (*[]string, bool)`

GetReferencedVcsOk returns a tuple with the ReferencedVcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedVcs

`func (o *ViewComposerMonitorDetails) SetReferencedVcs(v []string)`

SetReferencedVcs sets ReferencedVcs field to given value.


### GetVersion

`func (o *ViewComposerMonitorDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ViewComposerMonitorDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ViewComposerMonitorDetails) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


