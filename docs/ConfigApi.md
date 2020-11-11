# \ConfigApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateICDomainAccount**](ConfigApi.md#CreateICDomainAccount) | **Post** /config/v1/ic-domain-accounts | Creates instant clone domain account.
[**CreateIMAsset**](ConfigApi.md#CreateIMAsset) | **Post** /config/v1/im-assets | Creates image management asset.
[**CreateIMAssets**](ConfigApi.md#CreateIMAssets) | **Post** /config/v1/im-assets/action/bulk-create | Creates one or more image management assets. Each of the index of result in the response, correspond to the index of the original asset.
[**CreateIMStream**](ConfigApi.md#CreateIMStream) | **Post** /config/v1/im-streams | Creates image management stream.
[**CreateIMStreams**](ConfigApi.md#CreateIMStreams) | **Post** /config/v1/im-streams/action/bulk-create | Creates one or more image management streams. Each of the index of result in the response, correspond to the index of the original stream.
[**CreateIMTag**](ConfigApi.md#CreateIMTag) | **Post** /config/v1/im-tags | Creates image management tag.
[**CreateIMTags**](ConfigApi.md#CreateIMTags) | **Post** /config/v1/im-tags/action/bulk-create | Creates one or more image management tags. Each of the index of result in the response, correspond to the index of the original tag.
[**CreateIMVersion**](ConfigApi.md#CreateIMVersion) | **Post** /config/v1/im-versions | Creates image management version.
[**CreateIMVersions**](ConfigApi.md#CreateIMVersions) | **Post** /config/v1/im-versions/action/bulk-create | Creates one or more image management versions. Each of the index of result in the response, correspond to the index of the original version.
[**DeleteICDomainAccount**](ConfigApi.md#DeleteICDomainAccount) | **Delete** /config/v1/ic-domain-accounts/{id} | Deletes instant clone domain account.
[**DeleteIMAsset**](ConfigApi.md#DeleteIMAsset) | **Delete** /config/v1/im-assets/{id} | Deletes image management asset.
[**DeleteIMStream**](ConfigApi.md#DeleteIMStream) | **Delete** /config/v1/im-streams/{id} | Deletes image management stream.
[**DeleteIMTag**](ConfigApi.md#DeleteIMTag) | **Delete** /config/v1/im-tags/{id} | Deletes image management tag.
[**DeleteIMVersion**](ConfigApi.md#DeleteIMVersion) | **Delete** /config/v1/im-versions/{id} | Deletes image management version.
[**GetEnvironmentUsingGET**](ConfigApi.md#GetEnvironmentUsingGET) | **Get** /config/v1/environment-properties | Retrieves the environment settings.
[**GetFeatureSettingsUsingGET**](ConfigApi.md#GetFeatureSettingsUsingGET) | **Get** /config/v1/settings/feature | Retrieves the feature settings.
[**GetGeneralSettingsUsingGET**](ConfigApi.md#GetGeneralSettingsUsingGET) | **Get** /config/v1/settings/general | Retrieves the general settings.
[**GetICDomainAccount**](ConfigApi.md#GetICDomainAccount) | **Get** /config/v1/ic-domain-accounts/{id} | Gets instant clone domain account.
[**GetIMAsset**](ConfigApi.md#GetIMAsset) | **Get** /config/v1/im-assets/{id} | Gets image management asset.
[**GetIMStream**](ConfigApi.md#GetIMStream) | **Get** /config/v1/im-streams/{id} | Gets image management stream.
[**GetIMTag**](ConfigApi.md#GetIMTag) | **Get** /config/v1/im-tags/{id} | Gets image management tag.
[**GetIMVersion**](ConfigApi.md#GetIMVersion) | **Get** /config/v1/im-versions/{id} | Gets image management version.
[**GetSecuritySettingsUsingGET**](ConfigApi.md#GetSecuritySettingsUsingGET) | **Get** /config/v1/settings/security | Retrieves the security settings.
[**GetSettingsUsingGET**](ConfigApi.md#GetSettingsUsingGET) | **Get** /config/v1/settings | Retrieves the configuration settings.
[**ListICDomainAccounts**](ConfigApi.md#ListICDomainAccounts) | **Get** /config/v1/ic-domain-accounts | Lists instant clone domain accounts of the environment.
[**ListIMAssets**](ConfigApi.md#ListIMAssets) | **Get** /config/v1/im-assets | Lists image management assets.
[**ListIMStreams**](ConfigApi.md#ListIMStreams) | **Get** /config/v1/im-streams | Lists image management streams.
[**ListIMTags**](ConfigApi.md#ListIMTags) | **Get** /config/v1/im-tags | Lists image management tags.
[**ListIMVersions**](ConfigApi.md#ListIMVersions) | **Get** /config/v1/im-versions | Lists image management versions.
[**ListRCXServers**](ConfigApi.md#ListRCXServers) | **Get** /config/v1/rcx/servers | Lists RCX servers of the cluster.
[**ListVCInfo**](ConfigApi.md#ListVCInfo) | **Get** /config/v1/virtual-centers | Lists Virtual Centers configured in the environment.
[**RegisterRCXClient**](ConfigApi.md#RegisterRCXClient) | **Post** /config/v1/rcx/clients | Registers the RCX client
[**UnregisterRCXClient**](ConfigApi.md#UnregisterRCXClient) | **Delete** /config/v1/rcx/clients/{id} | Unregisters the given RCX Client
[**UpdateFeatureSettingsUsingPUT**](ConfigApi.md#UpdateFeatureSettingsUsingPUT) | **Put** /config/v1/settings/feature | Updates the feature settings.
[**UpdateGeneralSettingsUsingPUT**](ConfigApi.md#UpdateGeneralSettingsUsingPUT) | **Put** /config/v1/settings/general | Updates the general settings.
[**UpdateICDomainAccount**](ConfigApi.md#UpdateICDomainAccount) | **Put** /config/v1/ic-domain-accounts/{id} | Updates instant clone domain account.
[**UpdateIMAsset**](ConfigApi.md#UpdateIMAsset) | **Put** /config/v1/im-assets/{id} | Updates image management asset.
[**UpdateIMStream**](ConfigApi.md#UpdateIMStream) | **Put** /config/v1/im-streams/{id} | Updates image management stream.
[**UpdateIMTag**](ConfigApi.md#UpdateIMTag) | **Put** /config/v1/im-tags/{id} | Updates image management tag.
[**UpdateIMVersion**](ConfigApi.md#UpdateIMVersion) | **Put** /config/v1/im-versions/{id} | Updates image management version.
[**UpdateRCXClient**](ConfigApi.md#UpdateRCXClient) | **Put** /config/v1/rcx/clients/{id} | Updates the given RCX client.
[**UpdateSecuritySettingsUsingPUT**](ConfigApi.md#UpdateSecuritySettingsUsingPUT) | **Put** /config/v1/settings/security | Updates the security settings.
[**UpdateSettingsUsingPUT**](ConfigApi.md#UpdateSettingsUsingPUT) | **Put** /config/v1/settings | Updates the configuration settings.



## CreateICDomainAccount

> CreateICDomainAccount(ctx, body)

Creates instant clone domain account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InstantCloneDomainAccountCreateSpec**](InstantCloneDomainAccountCreateSpec.md)| Instant clone domain account object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIMAsset

> CreateIMAsset(ctx, body)

Creates image management asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ImageManagementAssetCreateSpec**](ImageManagementAssetCreateSpec.md)| Image management asset object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIMAssets

> []BulkItemResponseInfo CreateIMAssets(ctx, body)

Creates one or more image management assets. Each of the index of result in the response, correspond to the index of the original asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]ImageManagementAssetCreateSpec**](ImageManagementAssetCreateSpec.md)| List of Image management asset object to be created in bulk. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIMStream

> CreateIMStream(ctx, body)

Creates image management stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ImageManagementStreamCreateSpec**](ImageManagementStreamCreateSpec.md)| Image management stream object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIMStreams

> []BulkItemResponseInfo CreateIMStreams(ctx, body)

Creates one or more image management streams. Each of the index of result in the response, correspond to the index of the original stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]ImageManagementStreamCreateSpec**](ImageManagementStreamCreateSpec.md)| List of Image management stream object to be created in bulk. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIMTag

> CreateIMTag(ctx, body)

Creates image management tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ImageManagementTagCreateSpec**](ImageManagementTagCreateSpec.md)| Image management tag object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIMTags

> []BulkItemResponseInfo CreateIMTags(ctx, body)

Creates one or more image management tags. Each of the index of result in the response, correspond to the index of the original tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]ImageManagementTagCreateSpec**](ImageManagementTagCreateSpec.md)| List of Image management tag object to be created in bulk. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIMVersion

> CreateIMVersion(ctx, body)

Creates image management version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ImageManagementVersionCreateSpec**](ImageManagementVersionCreateSpec.md)| Image management version object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIMVersions

> []BulkItemResponseInfo CreateIMVersions(ctx, body)

Creates one or more image management versions. Each of the index of result in the response, correspond to the index of the original version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]ImageManagementVersionCreateSpec**](ImageManagementVersionCreateSpec.md)| List of Image management version object to be created in bulk. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteICDomainAccount

> DeleteICDomainAccount(ctx, id)

Deletes instant clone domain account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIMAsset

> DeleteIMAsset(ctx, id)

Deletes image management asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIMStream

> DeleteIMStream(ctx, id)

Deletes image management stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIMTag

> DeleteIMTag(ctx, id)

Deletes image management tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIMVersion

> DeleteIMVersion(ctx, id)

Deletes image management version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentUsingGET

> EnvironmentInfo GetEnvironmentUsingGET(ctx, )

Retrieves the environment settings.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EnvironmentInfo**](EnvironmentInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureSettingsUsingGET

> FeatureSettings GetFeatureSettingsUsingGET(ctx, )

Retrieves the feature settings.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**FeatureSettings**](FeatureSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneralSettingsUsingGET

> GeneralSettings GetGeneralSettingsUsingGET(ctx, )

Retrieves the general settings.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GeneralSettings**](GeneralSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetICDomainAccount

> InstantCloneDomainAccountInfo GetICDomainAccount(ctx, id)

Gets instant clone domain account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**InstantCloneDomainAccountInfo**](InstantCloneDomainAccountInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIMAsset

> ImageManagementAssetInfo GetIMAsset(ctx, id)

Gets image management asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**ImageManagementAssetInfo**](ImageManagementAssetInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIMStream

> ImageManagementStreamInfo GetIMStream(ctx, id)

Gets image management stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**ImageManagementStreamInfo**](ImageManagementStreamInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIMTag

> ImageManagementTagInfo GetIMTag(ctx, id)

Gets image management tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**ImageManagementTagInfo**](ImageManagementTagInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIMVersion

> ImageManagementVersionInfo GetIMVersion(ctx, id)

Gets image management version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**ImageManagementVersionInfo**](ImageManagementVersionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecuritySettingsUsingGET

> SecuritySettings GetSecuritySettingsUsingGET(ctx, )

Retrieves the security settings.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SecuritySettings**](SecuritySettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsUsingGET

> SettingsInfo GetSettingsUsingGET(ctx, )

Retrieves the configuration settings.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SettingsInfo**](SettingsInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListICDomainAccounts

> []InstantCloneDomainAccountInfo ListICDomainAccounts(ctx, )

Lists instant clone domain accounts of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]InstantCloneDomainAccountInfo**](InstantCloneDomainAccountInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIMAssets

> []ImageManagementAssetInfo ListIMAssets(ctx, imVersionId)

Lists image management assets.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imVersionId** | **string**| Image management version ID | 

### Return type

[**[]ImageManagementAssetInfo**](ImageManagementAssetInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIMStreams

> []ImageManagementStreamInfo ListIMStreams(ctx, )

Lists image management streams.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ImageManagementStreamInfo**](ImageManagementStreamInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIMTags

> []ImageManagementTagInfo ListIMTags(ctx, imStreamId)

Lists image management tags.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imStreamId** | **string**| Image management stream ID | 

### Return type

[**[]ImageManagementTagInfo**](ImageManagementTagInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIMVersions

> []ImageManagementVersionInfo ListIMVersions(ctx, imStreamId)

Lists image management versions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imStreamId** | **string**| Image management stream ID | 

### Return type

[**[]ImageManagementVersionInfo**](ImageManagementVersionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRCXServers

> []RcxServerInfo ListRCXServers(ctx, )

Lists RCX servers of the cluster.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]RcxServerInfo**](RCXServerInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVCInfo

> []VirtualCenterInfo ListVCInfo(ctx, )

Lists Virtual Centers configured in the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]VirtualCenterInfo**](VirtualCenterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterRCXClient

> RegisterRCXClient(ctx, body)

Registers the RCX client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RcxClientRegisterSpec**](RcxClientRegisterSpec.md)| RCX client object to be registered. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterRCXClient

> UnregisterRCXClient(ctx, id)

Unregisters the given RCX Client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureSettingsUsingPUT

> UpdateFeatureSettingsUsingPUT(ctx, body)

Updates the feature settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**FeatureSettingsUpdateSpec**](FeatureSettingsUpdateSpec.md)| Feature settings object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGeneralSettingsUsingPUT

> UpdateGeneralSettingsUsingPUT(ctx, body)

Updates the general settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GeneralSettingsUpdateSpec**](GeneralSettingsUpdateSpec.md)| General settings object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateICDomainAccount

> UpdateICDomainAccount(ctx, id, body)

Updates instant clone domain account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**InstantCloneDomainAccountUpdateSpec**](InstantCloneDomainAccountUpdateSpec.md)| Instant clone domain account object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIMAsset

> UpdateIMAsset(ctx, id, body)

Updates image management asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**ImageManagementAssetUpdateSpec**](ImageManagementAssetUpdateSpec.md)| Image management asset object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIMStream

> UpdateIMStream(ctx, id, body)

Updates image management stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**ImageManagementStreamUpdateSpec**](ImageManagementStreamUpdateSpec.md)| Image management stream object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIMTag

> UpdateIMTag(ctx, id, body)

Updates image management tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**ImageManagementTagUpdateSpec**](ImageManagementTagUpdateSpec.md)| Image management tag object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIMVersion

> UpdateIMVersion(ctx, id, body)

Updates image management version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**ImageManagementVersionUpdateSpec**](ImageManagementVersionUpdateSpec.md)| Image management version object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRCXClient

> UpdateRCXClient(ctx, id, body)

Updates the given RCX client.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**RcxClientUpdateSpec**](RcxClientUpdateSpec.md)| RCX client object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecuritySettingsUsingPUT

> UpdateSecuritySettingsUsingPUT(ctx, body)

Updates the security settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SecuritySettingsUpdateSpec**](SecuritySettingsUpdateSpec.md)| Security settings object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettingsUsingPUT

> UpdateSettingsUsingPUT(ctx, body)

Updates the configuration settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SettingsUpdateSpec**](SettingsUpdateSpec.md)| Configuration settings object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

