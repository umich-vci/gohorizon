# ApplicationAntiAffinityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiAffinityCount** | **int32** | Maximum number of other applications that can be running on the RDS Server before the RDS Server is rejected for new application sessions. | 
**AntiAffinityPatterns** | **[]string** | Set of pattern strings to match against process names on a RDS Server when attempting to launch a session for this Application. For each application running on an RDSServer that matches one of the patterns, the tally against the count threshold is incremented.&lt;br&gt;Pattern strings may contain wildcard characters. &#39;*&#39; matches zero or more characters. &#39;?&#39; matches exactly one character. | 

## Methods

### NewApplicationAntiAffinityData

`func NewApplicationAntiAffinityData(antiAffinityCount int32, antiAffinityPatterns []string, ) *ApplicationAntiAffinityData`

NewApplicationAntiAffinityData instantiates a new ApplicationAntiAffinityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAntiAffinityDataWithDefaults

`func NewApplicationAntiAffinityDataWithDefaults() *ApplicationAntiAffinityData`

NewApplicationAntiAffinityDataWithDefaults instantiates a new ApplicationAntiAffinityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiAffinityCount

`func (o *ApplicationAntiAffinityData) GetAntiAffinityCount() int32`

GetAntiAffinityCount returns the AntiAffinityCount field if non-nil, zero value otherwise.

### GetAntiAffinityCountOk

`func (o *ApplicationAntiAffinityData) GetAntiAffinityCountOk() (*int32, bool)`

GetAntiAffinityCountOk returns a tuple with the AntiAffinityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinityCount

`func (o *ApplicationAntiAffinityData) SetAntiAffinityCount(v int32)`

SetAntiAffinityCount sets AntiAffinityCount field to given value.


### GetAntiAffinityPatterns

`func (o *ApplicationAntiAffinityData) GetAntiAffinityPatterns() []string`

GetAntiAffinityPatterns returns the AntiAffinityPatterns field if non-nil, zero value otherwise.

### GetAntiAffinityPatternsOk

`func (o *ApplicationAntiAffinityData) GetAntiAffinityPatternsOk() (*[]string, bool)`

GetAntiAffinityPatternsOk returns a tuple with the AntiAffinityPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinityPatterns

`func (o *ApplicationAntiAffinityData) SetAntiAffinityPatterns(v []string)`

SetAntiAffinityPatterns sets AntiAffinityPatterns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


