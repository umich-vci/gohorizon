# \ExternalApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeUserPassword**](ExternalApi.md#ChangeUserPassword) | **Post** /external/v1/ad-users-or-groups/action/change-user-password | Changes the password of AD User
[**GetADUserOrGroupInfo**](ExternalApi.md#GetADUserOrGroupInfo) | **Get** /external/v1/ad-users-or-groups/{id} | Get information related to AD User or Group
[**ListADDomains**](ExternalApi.md#ListADDomains) | **Get** /external/v1/ad-domains | Lists information related to AD Domains of the environment.
[**ListADUserOrGroupSummary**](ExternalApi.md#ListADUserOrGroupSummary) | **Get** /external/v1/ad-users-or-groups | Lists AD users or groups information.
[**ListBaseSnapshots**](ExternalApi.md#ListBaseSnapshots) | **Get** /external/v1/base-snapshots | Lists all the VM snapshots from the vCenter for a given VM.
[**ListBaseVMs**](ExternalApi.md#ListBaseVMs) | **Get** /external/v1/base-vms | Lists all the VMs from a vCenter or a datacenter in that vCenter which may be suitable as snapshots for instant/linked clone desktop or farm creation.
[**ListCustomizationSpecs**](ExternalApi.md#ListCustomizationSpecs) | **Get** /external/v1/customization-specifications | Lists all the customization specifications from the vCenter.
[**ListDatacenters**](ExternalApi.md#ListDatacenters) | **Get** /external/v1/datacenters | Lists all the datacenters of a vCenter.
[**ListDatastorePaths**](ExternalApi.md#ListDatastorePaths) | **Get** /external/v1/datastore-paths | Lists all the folder paths within a Datastore from vCenter.
[**ListHostsOrClusters**](ExternalApi.md#ListHostsOrClusters) | **Get** /external/v1/hosts-or-clusters | Lists all the hosts or clusters of the datacenter.
[**ListNetworkInterfaceCards**](ExternalApi.md#ListNetworkInterfaceCards) | **Get** /external/v1/network-interface-cards | Returns a list of network interface cards (NICs) suitable for configuration on a desktop pool/farm.
[**ListNetworkLabels**](ExternalApi.md#ListNetworkLabels) | **Get** /external/v1/network-labels | Retrieves all network labels on the given host or cluster
[**ListResourcePools**](ExternalApi.md#ListResourcePools) | **Get** /external/v1/resource-pools | Lists all the resource pools from the vCenter for the given host or cluster.
[**ListVMFolders**](ExternalApi.md#ListVMFolders) | **Get** /external/v1/vm-folders | Lists all the VM folders from the vCenter for the given datacenter.
[**ListVMTemplates**](ExternalApi.md#ListVMTemplates) | **Get** /external/v1/vm-templates | Lists all the VM templates from a vCenter or a datacenter for the given vCenter which may be suitable for full clone desktop pool creation.
[**Listdatastores**](ExternalApi.md#Listdatastores) | **Get** /external/v1/datastores | Lists all the datastoress from the vCenter for the given host or cluster.
[**ValidateADUserEncryptedCredentials**](ExternalApi.md#ValidateADUserEncryptedCredentials) | **Post** /external/v1/ad-users-or-groups/action/validate-user-encrypted-credentials | Validates the encrypted credentials of AD User



## ChangeUserPassword

> AdUserInfo ChangeUserPassword(ctx, body)

Changes the password of AD User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AdUserChangePasswordSpec**](AdUserChangePasswordSpec.md)| AD user password object to be changed. | 

### Return type

[**AdUserInfo**](ADUserInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetADUserOrGroupInfo

> AdUserOrGroupInfo GetADUserOrGroupInfo(ctx, id)

Get information related to AD User or Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**AdUserOrGroupInfo**](ADUserOrGroupInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADDomains

> []AdDomainInfo ListADDomains(ctx, )

Lists information related to AD Domains of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AdDomainInfo**](ADDomainInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADUserOrGroupSummary

> []AdUserOrGroupSummary ListADUserOrGroupSummary(ctx, optional)

Lists AD users or groups information.

This API supports <b>Pagination</b> and <b>Filters.</b><br/>For Pagination, optional query params of 'page' and 'size' needs to be send.<br/>For Filters, refer to 'ADUserOrGroupSummary' model description to find supported filters on specific field.<br/>For full information on using Filters, refer to 'Horizon Server REST Pagination and Filter Guide' of 'VMware Horizon Server API' in code.vmware.com

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListADUserOrGroupSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListADUserOrGroupSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupOnly** | **optional.String**| Presence of this query param indicates to filter only groups or only users.   If passed as \&quot;true\&quot;, then only groups are returned.  If passed as \&quot;false\&quot;, then only users are returned.  If not passed passed at all, then both types are returned. | 

### Return type

[**[]AdUserOrGroupSummary**](ADUserOrGroupSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBaseSnapshots

> []BaseSnapshotInfo ListBaseSnapshots(ctx, baseVmId, vcenterId)

Lists all the VM snapshots from the vCenter for a given VM.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**baseVmId** | **string**| VM ID | 
**vcenterId** | **string**| Virtual Center ID | 

### Return type

[**[]BaseSnapshotInfo**](BaseSnapshotInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBaseVMs

> []BaseVmInfo ListBaseVMs(ctx, vcenterId, optional)

Lists all the VMs from a vCenter or a datacenter in that vCenter which may be suitable as snapshots for instant/linked clone desktop or farm creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vcenterId** | **string**| Virtual Center ID | 
 **optional** | ***ListBaseVMsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBaseVMsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datacenterId** | **optional.String**| Datacenter ID | 

### Return type

[**[]BaseVmInfo**](BaseVMInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomizationSpecs

> []CustomizationSpecInfo ListCustomizationSpecs(ctx, vcenterId)

Lists all the customization specifications from the vCenter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vcenterId** | **string**| Virtual Center ID | 

### Return type

[**[]CustomizationSpecInfo**](CustomizationSpecInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatacenters

> []DatacenterInfo ListDatacenters(ctx, vcenterId)

Lists all the datacenters of a vCenter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vcenterId** | **string**| Virtual Center ID | 

### Return type

[**[]DatacenterInfo**](DatacenterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatastorePaths

> []DatastorePathInfo ListDatastorePaths(ctx, datastoreId, vcenterId)

Lists all the folder paths within a Datastore from vCenter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datastoreId** | **string**| Datastore ID | 
**vcenterId** | **string**| Virtual Center ID | 

### Return type

[**[]DatastorePathInfo**](DatastorePathInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostsOrClusters

> []HostOrClusterInfo ListHostsOrClusters(ctx, datacenterId, vcenterId)

Lists all the hosts or clusters of the datacenter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| Datacenter ID | 
**vcenterId** | **string**| Virtual Center ID | 

### Return type

[**[]HostOrClusterInfo**](HostOrClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkInterfaceCards

> []NetworkInterfaceCardInfo ListNetworkInterfaceCards(ctx, vcenterId, optional)

Returns a list of network interface cards (NICs) suitable for configuration on a desktop pool/farm.

If the base VM and snapshot are specified, then the NICs  present on the given snapshot are listed.<br/> If the template is specified then the NICs present on the given template are listed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vcenterId** | **string**| Virtual Center ID | 
 **optional** | ***ListNetworkInterfaceCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNetworkInterfaceCardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseSnapshotId** | **optional.String**| Base Snapshot ID | 
 **baseVmId** | **optional.String**| Base VM ID | 
 **vmTemplateId** | **optional.String**| VM Template ID | 

### Return type

[**[]NetworkInterfaceCardInfo**](NetworkInterfaceCardInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkLabels

> []NetworkLabelInfo ListNetworkLabels(ctx, hostOrClusterId, vcenterId, optional)

Retrieves all network labels on the given host or cluster

API retrieves by filtering on the network type (if specified) that may be suitable for configuration with a desktop pool/farm's network interface card.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostOrClusterId** | **string**| Host or Cluster ID | 
**vcenterId** | **string**| Virtual Center ID | 
 **optional** | ***ListNetworkLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNetworkLabelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkType** | **optional.String**| Network Type * NETWORK: Standard network. * OPAQUE_NETWORK: Opaque network. * DISTRUBUTED_VIRTUAL_PORT_GROUP: DVS Port group. | 

### Return type

[**[]NetworkLabelInfo**](NetworkLabelInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourcePools

> []ResourcePoolInfo ListResourcePools(ctx, hostOrClusterId, vcenterId)

Lists all the resource pools from the vCenter for the given host or cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostOrClusterId** | **string**| Host or Cluster ID | 
**vcenterId** | **string**| Virtual Center ID | 

### Return type

[**[]ResourcePoolInfo**](ResourcePoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVMFolders

> []VmFolderInfo ListVMFolders(ctx, datacenterId, vcenterId)

Lists all the VM folders from the vCenter for the given datacenter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| Datacenter ID | 
**vcenterId** | **string**| Virtual Center ID | 

### Return type

[**[]VmFolderInfo**](VMFolderInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVMTemplates

> []VmTemplateInfo ListVMTemplates(ctx, vcenterId, optional)

Lists all the VM templates from a vCenter or a datacenter for the given vCenter which may be suitable for full clone desktop pool creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vcenterId** | **string**| Virtual Center ID | 
 **optional** | ***ListVMTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVMTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datacenterId** | **optional.String**| Datacenter ID | 

### Return type

[**[]VmTemplateInfo**](VMTemplateInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listdatastores

> []DatastoreInfo Listdatastores(ctx, hostOrClusterId, vcenterId)

Lists all the datastoress from the vCenter for the given host or cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostOrClusterId** | **string**| Host or Cluster ID | 
**vcenterId** | **string**| Virtual Center ID | 

### Return type

[**[]DatastoreInfo**](DatastoreInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateADUserEncryptedCredentials

> AdUserInfo ValidateADUserEncryptedCredentials(ctx, body)

Validates the encrypted credentials of AD User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AdUserEncryptedCredentialSpec**](AdUserEncryptedCredentialSpec.md)| AD user encrypted credentials object to be validated. | 

### Return type

[**AdUserInfo**](ADUserInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

