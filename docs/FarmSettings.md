# FarmSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteInProgess** | **bool** | Indicates whether the Farm is in the process of being deleted. Default value is false. | 
**DesktopId** | **string** | Desktop pool Id representing the RDS Desktop pool to which this Farm belongs. | [optional] 
**DisplayProtocolSettings** | [**FarmDisplayProtocolSettings**](FarmDisplayProtocolSettings.md) |  | [optional] 
**LoadBalancerSettings** | [**FarmLoadBalancerSettings**](FarmLoadBalancerSettings.md) |  | [optional] 
**ServerErrorThreshold** | **int32** | The minimum number of machines that must be fully operational in order toavoid showing the farm in an error state. Default value is 0. | [optional] 
**SessionSettings** | [**FarmSessionSettings**](FarmSessionSettings.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


