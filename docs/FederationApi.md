# \FederationApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHomeSites**](FederationApi.md#CreateHomeSites) | **Post** /federation/v1/home-sites | Creates the given home sites in the pod federation.
[**CreateSite**](FederationApi.md#CreateSite) | **Post** /federation/v1/sites | Creates a site.
[**DeleteHomeSites**](FederationApi.md#DeleteHomeSites) | **Delete** /federation/v1/home-sites | Deletes the given home sites from the pod federation.
[**DeleteSite**](FederationApi.md#DeleteSite) | **Delete** /federation/v1/sites/{id} | Deletes a site.
[**EjectPod**](FederationApi.md#EjectPod) | **Post** /federation/v1/cpa/action/eject | Removes a pod from Cloud Pod Federation.
[**GetHomeSite**](FederationApi.md#GetHomeSite) | **Get** /federation/v1/home-sites/{id} | Retrieves a given home site in the pod federation.
[**GetPod**](FederationApi.md#GetPod) | **Get** /federation/v1/pods/{id} | Retrieves a given pod from the pod federation.
[**GetPodAssignment**](FederationApi.md#GetPodAssignment) | **Get** /federation/v1/pod-assignments/{id} | Retrieves a given pod assignment from the pod federation.
[**GetPodEndpoint**](FederationApi.md#GetPodEndpoint) | **Get** /federation/v1/pods/{id}/endpoints/{endpointId} | Retrieves pod endpoint details for the given pod endpoint id in the given pod.
[**GetPodFederation**](FederationApi.md#GetPodFederation) | **Get** /federation/v1/cpa | Retrieves the pod federation details.
[**GetSite**](FederationApi.md#GetSite) | **Get** /federation/v1/sites/{id} | Retrives a given site.
[**GetTask**](FederationApi.md#GetTask) | **Get** /federation/v1/cpa/tasks/{id} | Retrieves the information for a given task.
[**InitializeCPA**](FederationApi.md#InitializeCPA) | **Post** /federation/v1/cpa/action/initialize | Initialize Cloud Pod Federation.
[**JoinCPA**](FederationApi.md#JoinCPA) | **Post** /federation/v1/cpa/action/join | Join Cloud Pod Federation.
[**ListHomeSites**](FederationApi.md#ListHomeSites) | **Get** /federation/v1/home-sites | Lists all the home sites in the pod federation.
[**ListPodAssignments**](FederationApi.md#ListPodAssignments) | **Get** /federation/v1/pod-assignments | Lists all the pod assignments in the pod federation.
[**ListPodEndpoint**](FederationApi.md#ListPodEndpoint) | **Get** /federation/v1/pods/{id}/endpoints | Lists all the pod endpoints for the given pod.
[**ListPods**](FederationApi.md#ListPods) | **Get** /federation/v1/pods | Lists all the pods in the pod federation.
[**ListSites**](FederationApi.md#ListSites) | **Get** /federation/v1/sites | Lists all the sites in the pod federation.
[**ListTasks**](FederationApi.md#ListTasks) | **Get** /federation/v1/cpa/tasks | Lists all the CPA tasks in the pod federation.
[**ResolveHomeSites**](FederationApi.md#ResolveHomeSites) | **Post** /federation/v1/home-sites/action/resolve | Resolves home sites for a user in the pod federation.
[**UninitializeCPA**](FederationApi.md#UninitializeCPA) | **Post** /federation/v1/cpa/action/uninitialize | Uninitialize Cloud Pod Federation.
[**UnjoinCPA**](FederationApi.md#UnjoinCPA) | **Post** /federation/v1/cpa/action/unjoin | Unjoin from Cloud Pod Federation.
[**UpdatePod**](FederationApi.md#UpdatePod) | **Put** /federation/v1/pods/{id} | Updates the given pod in the pod federation.
[**UpdatePodFederation**](FederationApi.md#UpdatePodFederation) | **Put** /federation/v1/cpa | Updates a Pod Federation.
[**UpdateSite**](FederationApi.md#UpdateSite) | **Put** /federation/v1/sites/{id} | Updates a site.



## CreateHomeSites

> []BulkItemResponseInfo CreateHomeSites(ctx).Body(body).Execute()

Creates the given home sites in the pod federation.



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
    body := []openapiclient.HomeSiteCreateSpec{*openapiclient.NewHomeSiteCreateSpec("S-1-5-32-551", "32a5ea06-cd09-4609-b3e5-df8379e99c13")} // []HomeSiteCreateSpec | List of home site objects to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.CreateHomeSites(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.CreateHomeSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHomeSites`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.CreateHomeSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHomeSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]HomeSiteCreateSpec**](HomeSiteCreateSpec.md) | List of home site objects to be created. | 

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


## CreateSite

> CreateSite(ctx).Body(body).Execute()

Creates a site.



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
    body := *openapiclient.NewSiteCreateSpec("US Site") // SiteCreateSpec | Site object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.CreateSite(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.CreateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SiteCreateSpec**](SiteCreateSpec.md) | Site object to be created. | 

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


## DeleteHomeSites

> []BulkItemResponseInfo DeleteHomeSites(ctx).Body(body).Execute()

Deletes the given home sites from the pod federation.



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
    body := []string{"Property_example"} // []string | List of home site IDs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.DeleteHomeSites(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.DeleteHomeSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHomeSites`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.DeleteHomeSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHomeSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of home site IDs to be deleted. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSite

> DeleteSite(ctx, id).Execute()

Deletes a site.



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
    resp, r, err := api_client.FederationApi.DeleteSite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.DeleteSite``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteRequest struct via the builder pattern


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


## EjectPod

> EjectPod(ctx).Body(body).Execute()

Removes a pod from Cloud Pod Federation.



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
    body := *openapiclient.NewCPAEjectSpec("9e94a90d-e7c2-40b6-a702-bd781512408d") // CPAEjectSpec | The specification for removing a pod from pod federation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.EjectPod(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.EjectPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEjectPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CPAEjectSpec**](CPAEjectSpec.md) | The specification for removing a pod from pod federation. | 

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


## GetHomeSite

> HomeSiteInfo GetHomeSite(ctx, id).Execute()

Retrieves a given home site in the pod federation.



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
    resp, r, err := api_client.FederationApi.GetHomeSite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.GetHomeSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHomeSite`: HomeSiteInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.GetHomeSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHomeSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HomeSiteInfo**](HomeSiteInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPod

> PodInfo GetPod(ctx, id).Execute()

Retrieves a given pod from the pod federation.



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
    id := "id_example" // string | ID of the Pod.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.GetPod(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.GetPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPod`: PodInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.GetPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Pod. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PodInfo**](PodInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPodAssignment

> PodAssignmentInfo GetPodAssignment(ctx, id).Execute()

Retrieves a given pod assignment from the pod federation.



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
    resp, r, err := api_client.FederationApi.GetPodAssignment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.GetPodAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPodAssignment`: PodAssignmentInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.GetPodAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPodAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PodAssignmentInfo**](PodAssignmentInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPodEndpoint

> PodEndpointInfo GetPodEndpoint(ctx, endpointId, id).Execute()

Retrieves pod endpoint details for the given pod endpoint id in the given pod.



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
    endpointId := "endpointId_example" // string | Pod endpoint ID
    id := "id_example" // string | Pod ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.GetPodEndpoint(context.Background(), endpointId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.GetPodEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPodEndpoint`: PodEndpointInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.GetPodEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **string** | Pod endpoint ID | 
**id** | **string** | Pod ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPodEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PodEndpointInfo**](PodEndpointInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPodFederation

> PodFederationInfo GetPodFederation(ctx).Execute()

Retrieves the pod federation details.



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
    resp, r, err := api_client.FederationApi.GetPodFederation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.GetPodFederation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPodFederation`: PodFederationInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.GetPodFederation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPodFederationRequest struct via the builder pattern


### Return type

[**PodFederationInfo**](PodFederationInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSite

> SiteInfo GetSite(ctx, id).Execute()

Retrives a given site.



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
    resp, r, err := api_client.FederationApi.GetSite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.GetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSite`: SiteInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.GetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteInfo**](SiteInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> TaskInfo GetTask(ctx, id).Execute()

Retrieves the information for a given task.



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
    resp, r, err := api_client.FederationApi.GetTask(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: TaskInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskInfo**](TaskInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeCPA

> CPATaskResponseInfo InitializeCPA(ctx).Execute()

Initialize Cloud Pod Federation.



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
    resp, r, err := api_client.FederationApi.InitializeCPA(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.InitializeCPA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeCPA`: CPATaskResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.InitializeCPA`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeCPARequest struct via the builder pattern


### Return type

[**CPATaskResponseInfo**](CPATaskResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinCPA

> CPATaskResponseInfo JoinCPA(ctx).Body(body).Execute()

Join Cloud Pod Federation.



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
    body := *openapiclient.NewCPAJoinSpec([]string{"Password_example"}, "cs1.example.com", "AD-TEST-DOMAIN\Administrator") // CPAJoinSpec | The specification for joining the pod federation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.JoinCPA(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.JoinCPA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoinCPA`: CPATaskResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.JoinCPA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJoinCPARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CPAJoinSpec**](CPAJoinSpec.md) | The specification for joining the pod federation. | 

### Return type

[**CPATaskResponseInfo**](CPATaskResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHomeSites

> []HomeSiteInfo ListHomeSites(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists all the home sites in the pod federation.



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
    filter := "{"type":"And", "filters":[{"type":"Equals", "name":"<>", "value":"<>"}] }" // string | filter expression built using fields with <b>'supported filters'</b> as described in output <b>model</b> schema of this API. (optional)
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    size := int32(10) // int32 | size, if passed should be > 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.ListHomeSites(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.ListHomeSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHomeSites`: []HomeSiteInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.ListHomeSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHomeSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0 | 

### Return type

[**[]HomeSiteInfo**](HomeSiteInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPodAssignments

> []PodAssignmentInfo ListPodAssignments(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists all the pod assignments in the pod federation.



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
    filter := "{"type":"And", "filters":[{"type":"Equals", "name":"<>", "value":"<>"}] }" // string | filter expression built using fields with <b>'supported filters'</b> as described in output <b>model</b> schema of this API. (optional)
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    size := int32(10) // int32 | size, if passed should be > 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.ListPodAssignments(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.ListPodAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPodAssignments`: []PodAssignmentInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.ListPodAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPodAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0 | 

### Return type

[**[]PodAssignmentInfo**](PodAssignmentInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPodEndpoint

> []PodEndpointInfo ListPodEndpoint(ctx, id).Execute()

Lists all the pod endpoints for the given pod.



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
    id := "id_example" // string | Pod ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.ListPodEndpoint(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.ListPodEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPodEndpoint`: []PodEndpointInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.ListPodEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Pod ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPodEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PodEndpointInfo**](PodEndpointInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPods

> []PodInfo ListPods(ctx).Execute()

Lists all the pods in the pod federation.



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
    resp, r, err := api_client.FederationApi.ListPods(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.ListPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPods`: []PodInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.ListPods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPodsRequest struct via the builder pattern


### Return type

[**[]PodInfo**](PodInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSites

> []SiteInfo ListSites(ctx).Execute()

Lists all the sites in the pod federation.



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
    resp, r, err := api_client.FederationApi.ListSites(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.ListSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSites`: []SiteInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.ListSites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSitesRequest struct via the builder pattern


### Return type

[**[]SiteInfo**](SiteInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> []TaskInfo ListTasks(ctx).Execute()

Lists all the CPA tasks in the pod federation.



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
    resp, r, err := api_client.FederationApi.ListTasks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.ListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTasks`: []TaskInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.ListTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


### Return type

[**[]TaskInfo**](TaskInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveHomeSites

> []HomeSiteResolutionInfo ResolveHomeSites(ctx).Body(body).Execute()

Resolves home sites for a user in the pod federation.



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
    body := *openapiclient.NewHomeSiteResolutionSpec("S-1-5-21-3623811015-3361044348") // HomeSiteResolutionSpec | Home site specification to be resolved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.ResolveHomeSites(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.ResolveHomeSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveHomeSites`: []HomeSiteResolutionInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.ResolveHomeSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveHomeSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HomeSiteResolutionSpec**](HomeSiteResolutionSpec.md) | Home site specification to be resolved. | 

### Return type

[**[]HomeSiteResolutionInfo**](HomeSiteResolutionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninitializeCPA

> CPATaskResponseInfo UninitializeCPA(ctx).Execute()

Uninitialize Cloud Pod Federation.



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
    resp, r, err := api_client.FederationApi.UninitializeCPA(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.UninitializeCPA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UninitializeCPA`: CPATaskResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.UninitializeCPA`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUninitializeCPARequest struct via the builder pattern


### Return type

[**CPATaskResponseInfo**](CPATaskResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnjoinCPA

> CPATaskResponseInfo UnjoinCPA(ctx).Execute()

Unjoin from Cloud Pod Federation.



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
    resp, r, err := api_client.FederationApi.UnjoinCPA(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.UnjoinCPA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnjoinCPA`: CPATaskResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `FederationApi.UnjoinCPA`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUnjoinCPARequest struct via the builder pattern


### Return type

[**CPATaskResponseInfo**](CPATaskResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePod

> UpdatePod(ctx, id).PodUpdateSpec(podUpdateSpec).Execute()

Updates the given pod in the pod federation.



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
    id := "id_example" // string | ID of the Pod.
    podUpdateSpec := *openapiclient.NewPodUpdateSpec("Cluster-CS-1", "9a892821-8c3d-4e61-9d65-69dfec7b70dc") // PodUpdateSpec | Pod object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.UpdatePod(context.Background(), id).PodUpdateSpec(podUpdateSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.UpdatePod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Pod. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **podUpdateSpec** | [**PodUpdateSpec**](PodUpdateSpec.md) | Pod object to be updated. | 

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


## UpdatePodFederation

> UpdatePodFederation(ctx).Body(body).Execute()

Updates a Pod Federation.



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
    body := *openapiclient.NewCPAUpdateSpec("Horizon Cloud Pod Federation") // CPAUpdateSpec | Pod Federation object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.UpdatePodFederation(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.UpdatePodFederation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePodFederationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CPAUpdateSpec**](CPAUpdateSpec.md) | Pod Federation object to be updated. | 

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


## UpdateSite

> UpdateSite(ctx, id).Body(body).Execute()

Updates a site.



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
    body := *openapiclient.NewSiteUpdateSpec("US Site") // SiteUpdateSpec | Site object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FederationApi.UpdateSite(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationApi.UpdateSite``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SiteUpdateSpec**](SiteUpdateSpec.md) | Site object to be updated. | 

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

