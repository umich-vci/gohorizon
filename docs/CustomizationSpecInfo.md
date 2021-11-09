# CustomizationSpecInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Customization specification description. | [optional] 
**GuestOs** | Pointer to **string** | Guest Operating system. * UNKNOWN: Unknown * WINDOWS: Windows * LINUX: Linux | [optional] 
**Id** | Pointer to **string** | Unique ID representing the customization specification. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this customization specification from being used in desktop pool creation. | [optional] 
**Name** | Pointer to **string** | Name of the customization specification. | [optional] 

## Methods

### NewCustomizationSpecInfo

`func NewCustomizationSpecInfo() *CustomizationSpecInfo`

NewCustomizationSpecInfo instantiates a new CustomizationSpecInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomizationSpecInfoWithDefaults

`func NewCustomizationSpecInfoWithDefaults() *CustomizationSpecInfo`

NewCustomizationSpecInfoWithDefaults instantiates a new CustomizationSpecInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomizationSpecInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomizationSpecInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomizationSpecInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomizationSpecInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGuestOs

`func (o *CustomizationSpecInfo) GetGuestOs() string`

GetGuestOs returns the GuestOs field if non-nil, zero value otherwise.

### GetGuestOsOk

`func (o *CustomizationSpecInfo) GetGuestOsOk() (*string, bool)`

GetGuestOsOk returns a tuple with the GuestOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOs

`func (o *CustomizationSpecInfo) SetGuestOs(v string)`

SetGuestOs sets GuestOs field to given value.

### HasGuestOs

`func (o *CustomizationSpecInfo) HasGuestOs() bool`

HasGuestOs returns a boolean if a field has been set.

### GetId

`func (o *CustomizationSpecInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomizationSpecInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomizationSpecInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomizationSpecInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *CustomizationSpecInfo) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *CustomizationSpecInfo) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *CustomizationSpecInfo) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *CustomizationSpecInfo) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetName

`func (o *CustomizationSpecInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomizationSpecInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomizationSpecInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomizationSpecInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


