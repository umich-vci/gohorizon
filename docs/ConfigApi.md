# \ConfigApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFederationAccessGroup**](ConfigApi.md#CreateFederationAccessGroup) | **Post** /config/v1/federation-access-groups | Creates federation access group.
[**CreateICDomainAccount**](ConfigApi.md#CreateICDomainAccount) | **Post** /config/v1/ic-domain-accounts | Creates instant clone domain account.
[**CreateIMAsset**](ConfigApi.md#CreateIMAsset) | **Post** /config/v1/im-assets | Creates image management asset.
[**CreateIMAssets**](ConfigApi.md#CreateIMAssets) | **Post** /config/v1/im-assets/action/bulk-create | Creates one or more image management assets. Each of the index of result in the response, correspond to the index of the original asset.
[**CreateIMStream**](ConfigApi.md#CreateIMStream) | **Post** /config/v1/im-streams | Creates image management stream.
[**CreateIMStreams**](ConfigApi.md#CreateIMStreams) | **Post** /config/v1/im-streams/action/bulk-create | Creates one or more image management streams. Each of the index of result in the response, correspond to the index of the original stream.
[**CreateIMTag**](ConfigApi.md#CreateIMTag) | **Post** /config/v1/im-tags | Creates image management tag.
[**CreateIMTags**](ConfigApi.md#CreateIMTags) | **Post** /config/v1/im-tags/action/bulk-create | Creates one or more image management tags. Each of the index of result in the response, correspond to the index of the original tag.
[**CreateIMVersion**](ConfigApi.md#CreateIMVersion) | **Post** /config/v1/im-versions | Creates image management version.
[**CreateIMVersions**](ConfigApi.md#CreateIMVersions) | **Post** /config/v1/im-versions/action/bulk-create | Creates one or more image management versions. Each of the index of result in the response, correspond to the index of the original version.
[**DeleteFederationAccessGroup**](ConfigApi.md#DeleteFederationAccessGroup) | **Delete** /config/v1/federation-access-groups/{id} | Deletes a federation access group.
[**DeleteICDomainAccount**](ConfigApi.md#DeleteICDomainAccount) | **Delete** /config/v1/ic-domain-accounts/{id} | Deletes instant clone domain account.
[**DeleteIMAsset**](ConfigApi.md#DeleteIMAsset) | **Delete** /config/v1/im-assets/{id} | Deletes image management asset.
[**DeleteIMStream**](ConfigApi.md#DeleteIMStream) | **Delete** /config/v1/im-streams/{id} | Deletes image management stream.
[**DeleteIMTag**](ConfigApi.md#DeleteIMTag) | **Delete** /config/v1/im-tags/{id} | Deletes image management tag.
[**DeleteIMVersion**](ConfigApi.md#DeleteIMVersion) | **Delete** /config/v1/im-versions/{id} | Deletes image management version.
[**GetEnvironment**](ConfigApi.md#GetEnvironment) | **Get** /config/v1/environment-properties | Retrieves the environment settings.
[**GetEnvironmentV2**](ConfigApi.md#GetEnvironmentV2) | **Get** /config/v2/environment-properties | Retrieves the environment settings.
[**GetFeatureSettings**](ConfigApi.md#GetFeatureSettings) | **Get** /config/v1/settings/feature | Retrieves the feature settings.
[**GetFederationAccessGroup**](ConfigApi.md#GetFederationAccessGroup) | **Get** /config/v1/federation-access-groups/{id} | Retrieves a federation access group.
[**GetGeneralSettings**](ConfigApi.md#GetGeneralSettings) | **Get** /config/v1/settings/general | Retrieves the general settings.
[**GetICDomainAccount**](ConfigApi.md#GetICDomainAccount) | **Get** /config/v1/ic-domain-accounts/{id} | Gets instant clone domain account.
[**GetIMAsset**](ConfigApi.md#GetIMAsset) | **Get** /config/v1/im-assets/{id} | Gets image management asset.
[**GetIMStream**](ConfigApi.md#GetIMStream) | **Get** /config/v1/im-streams/{id} | Gets image management stream.
[**GetIMTag**](ConfigApi.md#GetIMTag) | **Get** /config/v1/im-tags/{id} | Gets image management tag.
[**GetIMVersion**](ConfigApi.md#GetIMVersion) | **Get** /config/v1/im-versions/{id} | Gets image management version.
[**GetLocalAccessGroup**](ConfigApi.md#GetLocalAccessGroup) | **Get** /config/v1/local-access-groups/{id} | Retrieves a local access group.
[**GetSecuritySettings**](ConfigApi.md#GetSecuritySettings) | **Get** /config/v1/settings/security | Retrieves the security settings.
[**GetSettings**](ConfigApi.md#GetSettings) | **Get** /config/v1/settings | Retrieves the configuration settings.
[**ListFederationAccessGroups**](ConfigApi.md#ListFederationAccessGroups) | **Get** /config/v1/federation-access-groups | Lists all federation access groups.
[**ListICDomainAccounts**](ConfigApi.md#ListICDomainAccounts) | **Get** /config/v1/ic-domain-accounts | Lists instant clone domain accounts of the environment.
[**ListIMAssets**](ConfigApi.md#ListIMAssets) | **Get** /config/v1/im-assets | Lists image management assets.
[**ListIMStreams**](ConfigApi.md#ListIMStreams) | **Get** /config/v1/im-streams | Lists image management streams.
[**ListIMTags**](ConfigApi.md#ListIMTags) | **Get** /config/v1/im-tags | Lists image management tags.
[**ListIMVersions**](ConfigApi.md#ListIMVersions) | **Get** /config/v1/im-versions | Lists image management versions.
[**ListLocalAccessGroups**](ConfigApi.md#ListLocalAccessGroups) | **Get** /config/v1/local-access-groups | Lists all local access groups.
[**ListRCXServers**](ConfigApi.md#ListRCXServers) | **Get** /config/v1/rcx/servers | Lists RCX servers of the cluster.
[**ListVCInfo**](ConfigApi.md#ListVCInfo) | **Get** /config/v1/virtual-centers | Lists Virtual Centers configured in the environment.
[**ListVCInfoV2**](ConfigApi.md#ListVCInfoV2) | **Get** /config/v2/virtual-centers | Lists Virtual Centers configured in the environment.
[**RegisterRCXClient**](ConfigApi.md#RegisterRCXClient) | **Post** /config/v1/rcx/clients | Registers the RCX client
[**UnregisterRCXClient**](ConfigApi.md#UnregisterRCXClient) | **Delete** /config/v1/rcx/clients/{id} | Unregisters the given RCX Client
[**UpdateFeatureSettings**](ConfigApi.md#UpdateFeatureSettings) | **Put** /config/v1/settings/feature | Updates the feature settings.
[**UpdateGeneralSettings**](ConfigApi.md#UpdateGeneralSettings) | **Put** /config/v1/settings/general | Updates the general settings.
[**UpdateICDomainAccount**](ConfigApi.md#UpdateICDomainAccount) | **Put** /config/v1/ic-domain-accounts/{id} | Updates instant clone domain account.
[**UpdateIMAsset**](ConfigApi.md#UpdateIMAsset) | **Put** /config/v1/im-assets/{id} | Updates image management asset.
[**UpdateIMStream**](ConfigApi.md#UpdateIMStream) | **Put** /config/v1/im-streams/{id} | Updates image management stream.
[**UpdateIMTag**](ConfigApi.md#UpdateIMTag) | **Put** /config/v1/im-tags/{id} | Updates image management tag.
[**UpdateIMVersion**](ConfigApi.md#UpdateIMVersion) | **Put** /config/v1/im-versions/{id} | Updates image management version.
[**UpdateRCXClient**](ConfigApi.md#UpdateRCXClient) | **Put** /config/v1/rcx/clients/{id} | Updates the given RCX client.
[**UpdateSecuritySettings**](ConfigApi.md#UpdateSecuritySettings) | **Put** /config/v1/settings/security | Updates the security settings.
[**UpdateSettings**](ConfigApi.md#UpdateSettings) | **Put** /config/v1/settings | Updates the configuration settings.



## CreateFederationAccessGroup

> CreateFederationAccessGroup(ctx).Body(body).Execute()

Creates federation access group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewFederationAccessGroupCreateSpec("Sales") // FederationAccessGroupCreateSpec | Federation access group object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateFederationAccessGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateFederationAccessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFederationAccessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FederationAccessGroupCreateSpec**](FederationAccessGroupCreateSpec.md) | Federation access group object to be created. | 

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


## CreateICDomainAccount

> CreateICDomainAccount(ctx).Body(body).Execute()

Creates instant clone domain account.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewInstantCloneDomainAccountCreateSpec("S-1-5-21-1085031214-1563985344-725345543", []string{"Password_example"}, "testuser") // InstantCloneDomainAccountCreateSpec | Instant clone domain account object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateICDomainAccount(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateICDomainAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateICDomainAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InstantCloneDomainAccountCreateSpec**](InstantCloneDomainAccountCreateSpec.md) | Instant clone domain account object to be created. | 

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

> CreateIMAsset(ctx).Body(body).Execute()

Creates image management asset.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewImageManagementAssetCreateSpec("INSTANT_CLONE", "abc16e8f-9ba0-4789-a5dd-6880f32c52df", "6f85b3a5-e7d0-4ad6-a1e3-37168dd1ed51", "RDSH_APPS", "AVAILABLE", "f148f3e8-db0e-4abb-9c33-7e5205ccd360") // ImageManagementAssetCreateSpec | Image management asset object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateIMAsset(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateIMAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIMAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ImageManagementAssetCreateSpec**](ImageManagementAssetCreateSpec.md) | Image management asset object to be created. | 

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

> []BulkItemResponseInfo CreateIMAssets(ctx).Body(body).Execute()

Creates one or more image management assets. Each of the index of result in the response, correspond to the index of the original asset.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.ImageManagementAssetCreateSpec{*openapiclient.NewImageManagementAssetCreateSpec("INSTANT_CLONE", "abc16e8f-9ba0-4789-a5dd-6880f32c52df", "6f85b3a5-e7d0-4ad6-a1e3-37168dd1ed51", "RDSH_APPS", "AVAILABLE", "f148f3e8-db0e-4abb-9c33-7e5205ccd360")} // []ImageManagementAssetCreateSpec | List of Image management asset object to be created in bulk.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateIMAssets(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateIMAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIMAssets`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.CreateIMAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIMAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]ImageManagementAssetCreateSpec**](ImageManagementAssetCreateSpec.md) | List of Image management asset object to be created in bulk. | 

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

> CreateIMStream(ctx).Body(body).Execute()

Creates image management stream.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewImageManagementStreamCreateSpec("Win10", "WINDOWS_10", "MARKET_PLACE", "AVAILABLE") // ImageManagementStreamCreateSpec | Image management stream object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateIMStream(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateIMStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIMStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ImageManagementStreamCreateSpec**](ImageManagementStreamCreateSpec.md) | Image management stream object to be created. | 

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

> []BulkItemResponseInfo CreateIMStreams(ctx).Body(body).Execute()

Creates one or more image management streams. Each of the index of result in the response, correspond to the index of the original stream.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.ImageManagementStreamCreateSpec{*openapiclient.NewImageManagementStreamCreateSpec("Win10", "WINDOWS_10", "MARKET_PLACE", "AVAILABLE")} // []ImageManagementStreamCreateSpec | List of Image management stream object to be created in bulk.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateIMStreams(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateIMStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIMStreams`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.CreateIMStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIMStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]ImageManagementStreamCreateSpec**](ImageManagementStreamCreateSpec.md) | List of Image management stream object to be created in bulk. | 

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

> CreateIMTag(ctx).Body(body).Execute()

Creates image management tag.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewImageManagementTagCreateSpec("abc16e8f-9ba0-4789-a5dd-6880f32c52df", "6f85b3a5-e7d0-4ad6-a1e3-37168dd1ed51", "PROD") // ImageManagementTagCreateSpec | Image management tag object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateIMTag(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateIMTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIMTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ImageManagementTagCreateSpec**](ImageManagementTagCreateSpec.md) | Image management tag object to be created. | 

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

> []BulkItemResponseInfo CreateIMTags(ctx).Body(body).Execute()

Creates one or more image management tags. Each of the index of result in the response, correspond to the index of the original tag.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.ImageManagementTagCreateSpec{*openapiclient.NewImageManagementTagCreateSpec("abc16e8f-9ba0-4789-a5dd-6880f32c52df", "6f85b3a5-e7d0-4ad6-a1e3-37168dd1ed51", "PROD")} // []ImageManagementTagCreateSpec | List of Image management tag object to be created in bulk.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateIMTags(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateIMTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIMTags`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.CreateIMTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIMTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]ImageManagementTagCreateSpec**](ImageManagementTagCreateSpec.md) | List of Image management tag object to be created in bulk. | 

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

> CreateIMVersion(ctx).Body(body).Execute()

Creates image management version.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewImageManagementVersionCreateSpec("abc16e8f-9ba0-4789-a5dd-6880f32c52df", "v1", "AVAILABLE") // ImageManagementVersionCreateSpec | Image management version object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateIMVersion(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateIMVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIMVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ImageManagementVersionCreateSpec**](ImageManagementVersionCreateSpec.md) | Image management version object to be created. | 

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

> []BulkItemResponseInfo CreateIMVersions(ctx).Body(body).Execute()

Creates one or more image management versions. Each of the index of result in the response, correspond to the index of the original version.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.ImageManagementVersionCreateSpec{*openapiclient.NewImageManagementVersionCreateSpec("abc16e8f-9ba0-4789-a5dd-6880f32c52df", "v1", "AVAILABLE")} // []ImageManagementVersionCreateSpec | List of Image management version object to be created in bulk.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.CreateIMVersions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateIMVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIMVersions`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.CreateIMVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIMVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]ImageManagementVersionCreateSpec**](ImageManagementVersionCreateSpec.md) | List of Image management version object to be created in bulk. | 

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


## DeleteFederationAccessGroup

> DeleteFederationAccessGroup(ctx, id).Execute()

Deletes a federation access group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.DeleteFederationAccessGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteFederationAccessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFederationAccessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteICDomainAccount

> DeleteICDomainAccount(ctx, id).Execute()

Deletes instant clone domain account.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.DeleteICDomainAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteICDomainAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteICDomainAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> DeleteIMAsset(ctx, id).Execute()

Deletes image management asset.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.DeleteIMAsset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteIMAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIMAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> DeleteIMStream(ctx, id).Execute()

Deletes image management stream.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.DeleteIMStream(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteIMStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIMStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> DeleteIMTag(ctx, id).Execute()

Deletes image management tag.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.DeleteIMTag(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteIMTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIMTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> DeleteIMVersion(ctx, id).Execute()

Deletes image management version.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.DeleteIMVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteIMVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIMVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetEnvironment

> EnvironmentInfo GetEnvironment(ctx).Execute()

Retrieves the environment settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetEnvironment(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironment`: EnvironmentInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


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


## GetEnvironmentV2

> EnvironmentInfoV2 GetEnvironmentV2(ctx).Execute()

Retrieves the environment settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetEnvironmentV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetEnvironmentV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentV2`: EnvironmentInfoV2
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetEnvironmentV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentV2Request struct via the builder pattern


### Return type

[**EnvironmentInfoV2**](EnvironmentInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureSettings

> FeatureSettings GetFeatureSettings(ctx).Execute()

Retrieves the feature settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetFeatureSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetFeatureSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureSettings`: FeatureSettings
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetFeatureSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureSettingsRequest struct via the builder pattern


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


## GetFederationAccessGroup

> FederationAccessGroupInfo GetFederationAccessGroup(ctx, id).Execute()

Retrieves a federation access group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetFederationAccessGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetFederationAccessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFederationAccessGroup`: FederationAccessGroupInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetFederationAccessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFederationAccessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FederationAccessGroupInfo**](FederationAccessGroupInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneralSettings

> GeneralSettings GetGeneralSettings(ctx).Execute()

Retrieves the general settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetGeneralSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetGeneralSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeneralSettings`: GeneralSettings
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetGeneralSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneralSettingsRequest struct via the builder pattern


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

> InstantCloneDomainAccountInfo GetICDomainAccount(ctx, id).Execute()

Gets instant clone domain account.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetICDomainAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetICDomainAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetICDomainAccount`: InstantCloneDomainAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetICDomainAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetICDomainAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ImageManagementAssetInfo GetIMAsset(ctx, id).Execute()

Gets image management asset.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetIMAsset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetIMAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIMAsset`: ImageManagementAssetInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetIMAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIMAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ImageManagementStreamInfo GetIMStream(ctx, id).Execute()

Gets image management stream.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetIMStream(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetIMStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIMStream`: ImageManagementStreamInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetIMStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIMStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ImageManagementTagInfo GetIMTag(ctx, id).Execute()

Gets image management tag.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetIMTag(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetIMTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIMTag`: ImageManagementTagInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetIMTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIMTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ImageManagementVersionInfo GetIMVersion(ctx, id).Execute()

Gets image management version.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetIMVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetIMVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIMVersion`: ImageManagementVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetIMVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIMVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetLocalAccessGroup

> LocalAccessGroupInfo GetLocalAccessGroup(ctx, id).Execute()

Retrieves a local access group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetLocalAccessGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetLocalAccessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalAccessGroup`: LocalAccessGroupInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetLocalAccessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalAccessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocalAccessGroupInfo**](LocalAccessGroupInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecuritySettings

> SecuritySettings GetSecuritySettings(ctx).Execute()

Retrieves the security settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetSecuritySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetSecuritySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecuritySettings`: SecuritySettings
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetSecuritySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecuritySettingsRequest struct via the builder pattern


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


## GetSettings

> SettingsInfo GetSettings(ctx).Execute()

Retrieves the configuration settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.GetSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettings`: SettingsInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


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


## ListFederationAccessGroups

> []FederationAccessGroupInfo ListFederationAccessGroups(ctx).Execute()

Lists all federation access groups.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListFederationAccessGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListFederationAccessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFederationAccessGroups`: []FederationAccessGroupInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListFederationAccessGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFederationAccessGroupsRequest struct via the builder pattern


### Return type

[**[]FederationAccessGroupInfo**](FederationAccessGroupInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListICDomainAccounts

> []InstantCloneDomainAccountInfo ListICDomainAccounts(ctx).Page(page).Size(size).Execute()

Lists instant clone domain accounts of the environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListICDomainAccounts(context.Background()).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListICDomainAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListICDomainAccounts`: []InstantCloneDomainAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListICDomainAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListICDomainAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

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

> []ImageManagementAssetInfo ListIMAssets(ctx).ImVersionId(imVersionId).Execute()

Lists image management assets.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    imVersionId := "imVersionId_example" // string | Image management version ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListIMAssets(context.Background()).ImVersionId(imVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListIMAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIMAssets`: []ImageManagementAssetInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListIMAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIMAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imVersionId** | **string** | Image management version ID | 

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

> []ImageManagementStreamInfo ListIMStreams(ctx).Execute()

Lists image management streams.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListIMStreams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListIMStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIMStreams`: []ImageManagementStreamInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListIMStreams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIMStreamsRequest struct via the builder pattern


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

> []ImageManagementTagInfo ListIMTags(ctx).ImStreamId(imStreamId).Execute()

Lists image management tags.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    imStreamId := "imStreamId_example" // string | Image management stream ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListIMTags(context.Background()).ImStreamId(imStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListIMTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIMTags`: []ImageManagementTagInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListIMTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIMTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imStreamId** | **string** | Image management stream ID | 

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

> []ImageManagementVersionInfo ListIMVersions(ctx).ImStreamId(imStreamId).Execute()

Lists image management versions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    imStreamId := "imStreamId_example" // string | Image management stream ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListIMVersions(context.Background()).ImStreamId(imStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListIMVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIMVersions`: []ImageManagementVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListIMVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIMVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imStreamId** | **string** | Image management stream ID | 

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


## ListLocalAccessGroups

> []LocalAccessGroupInfo ListLocalAccessGroups(ctx).Execute()

Lists all local access groups.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListLocalAccessGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListLocalAccessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocalAccessGroups`: []LocalAccessGroupInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListLocalAccessGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLocalAccessGroupsRequest struct via the builder pattern


### Return type

[**[]LocalAccessGroupInfo**](LocalAccessGroupInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRCXServers

> []RCXServerInfo ListRCXServers(ctx).Execute()

Lists RCX servers of the cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListRCXServers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListRCXServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRCXServers`: []RCXServerInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListRCXServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRCXServersRequest struct via the builder pattern


### Return type

[**[]RCXServerInfo**](RCXServerInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVCInfo

> []VirtualCenterInfo ListVCInfo(ctx).Execute()

Lists Virtual Centers configured in the environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListVCInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListVCInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVCInfo`: []VirtualCenterInfo
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListVCInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVCInfoRequest struct via the builder pattern


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


## ListVCInfoV2

> []VirtualCenterInfoV2 ListVCInfoV2(ctx).Execute()

Lists Virtual Centers configured in the environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ListVCInfoV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ListVCInfoV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVCInfoV2`: []VirtualCenterInfoV2
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ListVCInfoV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVCInfoV2Request struct via the builder pattern


### Return type

[**[]VirtualCenterInfoV2**](VirtualCenterInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterRCXClient

> RegisterRCXClient(ctx).Body(body).Execute()

Registers the RCX client



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewRCXClientRegisterSpec("example.com", []openapiclient.CertificateThumbprint{*openapiclient.NewCertificateThumbprint("8f:92:9d:3b:a7:85:55:88:60:cd:e1:c8:1e:70:9a:8b:37:6d:a6:e6", "SHA_1")}) // RCXClientRegisterSpec | RCX client object to be registered.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.RegisterRCXClient(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.RegisterRCXClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRCXClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RCXClientRegisterSpec**](RCXClientRegisterSpec.md) | RCX client object to be registered. | 

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

> UnregisterRCXClient(ctx, id).Execute()

Unregisters the given RCX Client



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UnregisterRCXClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UnregisterRCXClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterRCXClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateFeatureSettings

> UpdateFeatureSettings(ctx).Body(body).Execute()

Updates the feature settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewFeatureSettingsUpdateSpec() // FeatureSettingsUpdateSpec | Feature settings object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateFeatureSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateFeatureSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FeatureSettingsUpdateSpec**](FeatureSettingsUpdateSpec.md) | Feature settings object to be updated. | 

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


## UpdateGeneralSettings

> UpdateGeneralSettings(ctx).Body(body).Execute()

Updates the general settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewGeneralSettingsUpdateSpec("ENABLED", "NEVER", "TIMEOUT_AFTER", int32(300), "DISABLED_AFTER", false, false) // GeneralSettingsUpdateSpec | General settings object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateGeneralSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateGeneralSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGeneralSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GeneralSettingsUpdateSpec**](GeneralSettingsUpdateSpec.md) | General settings object to be updated. | 

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

> UpdateICDomainAccount(ctx, id).Body(body).Execute()

Updates instant clone domain account.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id
    body := *openapiclient.NewInstantCloneDomainAccountUpdateSpec([]string{"Password_example"}) // InstantCloneDomainAccountUpdateSpec | Instant clone domain account object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateICDomainAccount(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateICDomainAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateICDomainAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InstantCloneDomainAccountUpdateSpec**](InstantCloneDomainAccountUpdateSpec.md) | Instant clone domain account object to be updated. | 

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

> UpdateIMAsset(ctx, id).Body(body).Execute()

Updates image management asset.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id
    body := *openapiclient.NewImageManagementAssetUpdateSpec("INSTANT_CLONE", "RDSH_APPS", "AVAILABLE") // ImageManagementAssetUpdateSpec | Image management asset object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateIMAsset(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateIMAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIMAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ImageManagementAssetUpdateSpec**](ImageManagementAssetUpdateSpec.md) | Image management asset object to be updated. | 

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

> UpdateIMStream(ctx, id).Body(body).Execute()

Updates image management stream.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id
    body := *openapiclient.NewImageManagementStreamUpdateSpec("Win10", "WINDOWS_10", "MARKET_PLACE", "AVAILABLE") // ImageManagementStreamUpdateSpec | Image management stream object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateIMStream(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateIMStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIMStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ImageManagementStreamUpdateSpec**](ImageManagementStreamUpdateSpec.md) | Image management stream object to be updated. | 

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

> UpdateIMTag(ctx, id).Body(body).Execute()

Updates image management tag.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id
    body := *openapiclient.NewImageManagementTagUpdateSpec("7e85b3a5-e7d0-4ad6-a1e3-37168dd1ed62", "PROD") // ImageManagementTagUpdateSpec | Image management tag object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateIMTag(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateIMTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIMTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ImageManagementTagUpdateSpec**](ImageManagementTagUpdateSpec.md) | Image management tag object to be updated. | 

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

> UpdateIMVersion(ctx, id).Body(body).Execute()

Updates image management version.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id
    body := *openapiclient.NewImageManagementVersionUpdateSpec("v1", "AVAILABLE") // ImageManagementVersionUpdateSpec | Image management version object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateIMVersion(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateIMVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIMVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ImageManagementVersionUpdateSpec**](ImageManagementVersionUpdateSpec.md) | Image management version object to be updated. | 

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

> UpdateRCXClient(ctx, id).Body(body).Execute()

Updates the given RCX client.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id
    body := *openapiclient.NewRCXClientUpdateSpec([]openapiclient.CertificateThumbprint{*openapiclient.NewCertificateThumbprint("8f:92:9d:3b:a7:85:55:88:60:cd:e1:c8:1e:70:9a:8b:37:6d:a6:e6", "SHA_1")}) // RCXClientUpdateSpec | RCX client object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateRCXClient(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateRCXClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRCXClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RCXClientUpdateSpec**](RCXClientUpdateSpec.md) | RCX client object to be updated. | 

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


## UpdateSecuritySettings

> UpdateSecuritySettings(ctx).Body(body).Execute()

Updates the security settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSecuritySettingsUpdateSpec("ENABLED") // SecuritySettingsUpdateSpec | Security settings object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateSecuritySettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateSecuritySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecuritySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SecuritySettingsUpdateSpec**](SecuritySettingsUpdateSpec.md) | Security settings object to be updated. | 

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


## UpdateSettings

> UpdateSettings(ctx).Body(body).Execute()

Updates the configuration settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSettingsUpdateSpec() // SettingsUpdateSpec | Configuration settings object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.UpdateSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SettingsUpdateSpec**](SettingsUpdateSpec.md) | Configuration settings object to be updated. | 

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

