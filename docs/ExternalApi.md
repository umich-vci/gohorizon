# \ExternalApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuxiliaryAccounts**](ExternalApi.md#AddAuxiliaryAccounts) | **Post** /external/v1/ad-domains/{id}/action/add-auxiliary-accounts | Add auxiliary accounts to the untrusted domain
[**AuditEventsExtendedAttributes**](ExternalApi.md#AuditEventsExtendedAttributes) | **Get** /external/v1/audit-events/extended-attributes | Get extended attributes of audit events
[**Bind**](ExternalApi.md#Bind) | **Post** /external/v1/ad-domains/action/bind | Bind untrusted domain to the connection server.
[**ChangeUserPassword**](ExternalApi.md#ChangeUserPassword) | **Post** /external/v1/ad-users-or-groups/action/change-user-password | Changes the password of AD User
[**ComputeSpaceRequirements**](ExternalApi.md#ComputeSpaceRequirements) | **Post** /external/v1/datastores/action/compute-requirements | Computes the datastore space requirements for inventory resources.
[**DeleteAuxiliaryAccounts**](ExternalApi.md#DeleteAuxiliaryAccounts) | **Post** /external/v1/ad-domains/{id}/action/delete-auxiliary-accounts | Specification to delete auxiliary accounts from the untrusted domain
[**GetADUserOrGroupInfo**](ExternalApi.md#GetADUserOrGroupInfo) | **Get** /external/v1/ad-users-or-groups/{id} | Get information related to AD User or Group
[**ListADDomains**](ExternalApi.md#ListADDomains) | **Get** /external/v1/ad-domains | Lists information related to AD Domains of the environment.
[**ListADDomainsV2**](ExternalApi.md#ListADDomainsV2) | **Get** /external/v2/ad-domains | Lists information related to AD Domains of the environment.
[**ListADDomainsV3**](ExternalApi.md#ListADDomainsV3) | **Get** /external/v3/ad-domains | Lists information related to AD Domains of the environment.
[**ListADUserOrGroupSummary**](ExternalApi.md#ListADUserOrGroupSummary) | **Get** /external/v1/ad-users-or-groups | Lists AD users or groups information.
[**ListAuditEvents**](ExternalApi.md#ListAuditEvents) | **Get** /external/v1/audit-events | Lists the audit events.
[**ListBaseSnapshots**](ExternalApi.md#ListBaseSnapshots) | **Get** /external/v1/base-snapshots | Lists all the VM snapshots from the vCenter for a given VM.
[**ListBaseVMs**](ExternalApi.md#ListBaseVMs) | **Get** /external/v1/base-vms | Lists all the VMs from a vCenter or a datacenter in that vCenter which may be suitable as snapshots for instant clone desktop pool or farm creation.
[**ListCustomizationSpecs**](ExternalApi.md#ListCustomizationSpecs) | **Get** /external/v1/customization-specifications | Lists all the customization specifications from the vCenter.
[**ListDatacenters**](ExternalApi.md#ListDatacenters) | **Get** /external/v1/datacenters | Lists all the datacenters of a vCenter.
[**ListDatastoreClusters**](ExternalApi.md#ListDatastoreClusters) | **Get** /external/v1/datastore-clusters | Lists all the datastore clusters from the vCenter for the given host or cluster.
[**ListDatastorePaths**](ExternalApi.md#ListDatastorePaths) | **Get** /external/v1/datastore-paths | Lists all the folder paths within a Datastore from vCenter.
[**ListHostsOrClusters**](ExternalApi.md#ListHostsOrClusters) | **Get** /external/v1/hosts-or-clusters | Lists all the hosts or clusters of the datacenter.
[**ListNetworkInterfaceCards**](ExternalApi.md#ListNetworkInterfaceCards) | **Get** /external/v1/network-interface-cards | Returns a list of network interface cards (NICs) suitable for configuration on a desktop pool/farm.
[**ListNetworkLabels**](ExternalApi.md#ListNetworkLabels) | **Get** /external/v1/network-labels | Retrieves all network labels on the given host or cluster
[**ListResourcePools**](ExternalApi.md#ListResourcePools) | **Get** /external/v1/resource-pools | Lists all the resource pools from the vCenter for the given host or cluster.
[**ListVMFolders**](ExternalApi.md#ListVMFolders) | **Get** /external/v1/vm-folders | Lists all the VM folders from the vCenter for the given datacenter.
[**ListVMTemplates**](ExternalApi.md#ListVMTemplates) | **Get** /external/v1/vm-templates | Lists all the VM templates from a vCenter or a datacenter for the given vCenter which may be suitable for full clone desktop pool creation.
[**ListVirtualMachines**](ExternalApi.md#ListVirtualMachines) | **Get** /external/v1/virtual-machines | Lists all the VMs from a vCenter.
[**Listdatastores**](ExternalApi.md#Listdatastores) | **Get** /external/v1/datastores | Lists all the datastores from the vCenter for the given host or cluster.
[**Unbind**](ExternalApi.md#Unbind) | **Post** /external/v1/ad-domains/{id}/action/unbind | Unbind untrusted domain from the connection server.
[**Update**](ExternalApi.md#Update) | **Post** /external/v1/ad-domains/{id}/action/update | Updates untrusted domain.
[**UpdateAuxiliaryAccounts**](ExternalApi.md#UpdateAuxiliaryAccounts) | **Post** /external/v1/ad-domains/action/update-auxiliary-accounts | Update auxiliary accounts of the untrusted domain
[**ValidateADUserEncryptedCredentials**](ExternalApi.md#ValidateADUserEncryptedCredentials) | **Post** /external/v1/ad-users-or-groups/action/validate-user-encrypted-credentials | Validates the encrypted credentials of AD User



## AddAuxiliaryAccounts

> []BulkItemResponseInfo AddAuxiliaryAccounts(ctx, id).Body(body).Execute()

Add auxiliary accounts to the untrusted domain



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
    body := *openapiclient.NewADDomainAuxiliaryAccountCreateSpec([]openapiclient.ServiceAccountCredentials{*openapiclient.NewServiceAccountCredentials([]string{"Password_example"}, "Administrator")}) // ADDomainAuxiliaryAccountCreateSpec | Specification of auxiliary accounts.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.AddAuxiliaryAccounts(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.AddAuxiliaryAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAuxiliaryAccounts`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.AddAuxiliaryAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAuxiliaryAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ADDomainAuxiliaryAccountCreateSpec**](ADDomainAuxiliaryAccountCreateSpec.md) | Specification of auxiliary accounts. | 

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


## AuditEventsExtendedAttributes

> []AuditEventAttributeInfo AuditEventsExtendedAttributes(ctx).Ids(ids).Execute()

Get extended attributes of audit events

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
    ids := []int64{int64(123)} // []int64 | Audit Event IDs

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.AuditEventsExtendedAttributes(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.AuditEventsExtendedAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditEventsExtendedAttributes`: []AuditEventAttributeInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.AuditEventsExtendedAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditEventsExtendedAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Audit Event IDs | 

### Return type

[**[]AuditEventAttributeInfo**](AuditEventAttributeInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bind

> ADDomainBindInfo Bind(ctx).Body(body).Execute()

Bind untrusted domain to the connection server.



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
    body := *openapiclient.NewADDomainSpec("example.com", "EXAMPLE") // ADDomainSpec | Specification of untrusted domain.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.Bind(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.Bind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Bind`: ADDomainBindInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.Bind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ADDomainSpec**](ADDomainSpec.md) | Specification of untrusted domain. | 

### Return type

[**ADDomainBindInfo**](ADDomainBindInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeUserPassword

> ADUserInfo ChangeUserPassword(ctx).Body(body).Execute()

Changes the password of AD User



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
    body := *openapiclient.NewADUserChangePasswordSpec("ut0yGq7CbhtBUGfW3ngjkvjQ2cc=", "4TGdfggfwrrGUPSg4/JK/tYub7lUI8pGtyHd/ty5g8h5=", "9QYxpdXrcrOGPSSz1/K/pJu8QlYT7pDkaKg/rb3hlw4=", "Qvvjglg5iZinyuldroueo/hQFyqydMMDJPmfYGgIebqxbU9chJ9I8iM9SCBRHSkSW9y+RMQOfC", "testuser or testuser@example.com") // ADUserChangePasswordSpec | AD user password object to be changed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ChangeUserPassword(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ChangeUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeUserPassword`: ADUserInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ChangeUserPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ADUserChangePasswordSpec**](ADUserChangePasswordSpec.md) | AD user password object to be changed. | 

### Return type

[**ADUserInfo**](ADUserInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeSpaceRequirements

> []DatastoreSpaceRequirementInfo ComputeSpaceRequirements(ctx).Body(body).Execute()

Computes the datastore space requirements for inventory resources.



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
    body := *openapiclient.NewDatastoreSpaceRequirementSpec(int32(123), "INSTANT_CLONE", "FARM", "ed3f92f3-0eef-4bf1-a405-de69f138d382") // DatastoreSpaceRequirementSpec | Datastore space requirement to be computed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ComputeSpaceRequirements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ComputeSpaceRequirements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeSpaceRequirements`: []DatastoreSpaceRequirementInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ComputeSpaceRequirements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComputeSpaceRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DatastoreSpaceRequirementSpec**](DatastoreSpaceRequirementSpec.md) | Datastore space requirement to be computed. | 

### Return type

[**[]DatastoreSpaceRequirementInfo**](DatastoreSpaceRequirementInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuxiliaryAccounts

> []BulkItemResponseInfo DeleteAuxiliaryAccounts(ctx, id).Body(body).Execute()

Specification to delete auxiliary accounts from the untrusted domain



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
    body := *openapiclient.NewADDomainAuxiliaryAccountDeleteSpec([]string{"AuxiliaryAccountIds_example"}) // ADDomainAuxiliaryAccountDeleteSpec | Auxiliary accounts to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.DeleteAuxiliaryAccounts(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.DeleteAuxiliaryAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAuxiliaryAccounts`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.DeleteAuxiliaryAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuxiliaryAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ADDomainAuxiliaryAccountDeleteSpec**](ADDomainAuxiliaryAccountDeleteSpec.md) | Auxiliary accounts to delete. | 

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


## GetADUserOrGroupInfo

> ADUserOrGroupInfo GetADUserOrGroupInfo(ctx, id).Execute()

Get information related to AD User or Group

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
    resp, r, err := api_client.ExternalApi.GetADUserOrGroupInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.GetADUserOrGroupInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetADUserOrGroupInfo`: ADUserOrGroupInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.GetADUserOrGroupInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetADUserOrGroupInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ADUserOrGroupInfo**](ADUserOrGroupInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADDomains

> []ADDomainInfo ListADDomains(ctx).Execute()

Lists information related to AD Domains of the environment.

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
    resp, r, err := api_client.ExternalApi.ListADDomains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListADDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListADDomains`: []ADDomainInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListADDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListADDomainsRequest struct via the builder pattern


### Return type

[**[]ADDomainInfo**](ADDomainInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADDomainsV2

> []ADDomainInfoV2 ListADDomainsV2(ctx).Execute()

Lists information related to AD Domains of the environment.

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
    resp, r, err := api_client.ExternalApi.ListADDomainsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListADDomainsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListADDomainsV2`: []ADDomainInfoV2
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListADDomainsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListADDomainsV2Request struct via the builder pattern


### Return type

[**[]ADDomainInfoV2**](ADDomainInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADDomainsV3

> []ADDomainInfoV3 ListADDomainsV3(ctx).Execute()

Lists information related to AD Domains of the environment.

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
    resp, r, err := api_client.ExternalApi.ListADDomainsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListADDomainsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListADDomainsV3`: []ADDomainInfoV3
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListADDomainsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListADDomainsV3Request struct via the builder pattern


### Return type

[**[]ADDomainInfoV3**](ADDomainInfoV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADUserOrGroupSummary

> []ADUserOrGroupSummary ListADUserOrGroupSummary(ctx).Filter(filter).GroupOnly(groupOnly).Page(page).Size(size).Execute()

Lists AD users or groups information.



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
    groupOnly := "groupOnly_example" // string | Presence of this query param indicates to filter only groups or only users.   If passed as \"true\", then only groups are returned.  If passed as \"false\", then only users are returned.  If not passed passed at all, then both types are returned. (optional)
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    size := int32(10) // int32 | size, if passed should be > 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListADUserOrGroupSummary(context.Background()).Filter(filter).GroupOnly(groupOnly).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListADUserOrGroupSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListADUserOrGroupSummary`: []ADUserOrGroupSummary
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListADUserOrGroupSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListADUserOrGroupSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **groupOnly** | **string** | Presence of this query param indicates to filter only groups or only users.   If passed as \&quot;true\&quot;, then only groups are returned.  If passed as \&quot;false\&quot;, then only users are returned.  If not passed passed at all, then both types are returned. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0 | 

### Return type

[**[]ADUserOrGroupSummary**](ADUserOrGroupSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuditEvents

> []AuditEventSummary ListAuditEvents(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the audit events.



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
    resp, r, err := api_client.ExternalApi.ListAuditEvents(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListAuditEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditEvents`: []AuditEventSummary
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListAuditEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0 | 

### Return type

[**[]AuditEventSummary**](AuditEventSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBaseSnapshots

> []BaseSnapshotInfo ListBaseSnapshots(ctx).BaseVmId(baseVmId).VcenterId(vcenterId).Execute()

Lists all the VM snapshots from the vCenter for a given VM.



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
    baseVmId := "vm-1" // string | VM ID
    vcenterId := "d0325b13-2bf1-4fa4-b027-e780004f2d1e" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListBaseSnapshots(context.Background()).BaseVmId(baseVmId).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListBaseSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBaseSnapshots`: []BaseSnapshotInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListBaseSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBaseSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseVmId** | **string** | VM ID | 
 **vcenterId** | **string** | Virtual Center ID | 

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

> []BaseVMInfo ListBaseVMs(ctx).VcenterId(vcenterId).DatacenterId(datacenterId).FilterIncompatibleVms(filterIncompatibleVms).Execute()

Lists all the VMs from a vCenter or a datacenter in that vCenter which may be suitable as snapshots for instant clone desktop pool or farm creation.



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
    vcenterId := "vcenterId_example" // string | Virtual Center ID
    datacenterId := "datacenterId_example" // string | Datacenter ID (optional)
    filterIncompatibleVms := false // bool | Whether to filter out incompatible VMs (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListBaseVMs(context.Background()).VcenterId(vcenterId).DatacenterId(datacenterId).FilterIncompatibleVms(filterIncompatibleVms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListBaseVMs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBaseVMs`: []BaseVMInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListBaseVMs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBaseVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vcenterId** | **string** | Virtual Center ID | 
 **datacenterId** | **string** | Datacenter ID | 
 **filterIncompatibleVms** | **bool** | Whether to filter out incompatible VMs | [default to false]

### Return type

[**[]BaseVMInfo**](BaseVMInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomizationSpecs

> []CustomizationSpecInfo ListCustomizationSpecs(ctx).VcenterId(vcenterId).Execute()

Lists all the customization specifications from the vCenter.



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
    vcenterId := "vcenterId_example" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListCustomizationSpecs(context.Background()).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListCustomizationSpecs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomizationSpecs`: []CustomizationSpecInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListCustomizationSpecs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomizationSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vcenterId** | **string** | Virtual Center ID | 

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

> []DatacenterInfo ListDatacenters(ctx).VcenterId(vcenterId).Execute()

Lists all the datacenters of a vCenter.



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
    vcenterId := "vcenterId_example" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListDatacenters(context.Background()).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListDatacenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatacenters`: []DatacenterInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListDatacenters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatacentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vcenterId** | **string** | Virtual Center ID | 

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


## ListDatastoreClusters

> []DatastoreClusterInfo ListDatastoreClusters(ctx).HostOrClusterId(hostOrClusterId).VcenterId(vcenterId).Execute()

Lists all the datastore clusters from the vCenter for the given host or cluster.



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
    hostOrClusterId := "hostOrClusterId_example" // string | Host or Cluster ID
    vcenterId := "vcenterId_example" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListDatastoreClusters(context.Background()).HostOrClusterId(hostOrClusterId).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListDatastoreClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatastoreClusters`: []DatastoreClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListDatastoreClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatastoreClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostOrClusterId** | **string** | Host or Cluster ID | 
 **vcenterId** | **string** | Virtual Center ID | 

### Return type

[**[]DatastoreClusterInfo**](DatastoreClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatastorePaths

> []DatastorePathInfo ListDatastorePaths(ctx).DatastoreId(datastoreId).VcenterId(vcenterId).Execute()

Lists all the folder paths within a Datastore from vCenter.



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
    datastoreId := "datastore-29" // string | Datastore ID
    vcenterId := "d0325b13-2bf1-4fa4-b027-e780004f2d1e" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListDatastorePaths(context.Background()).DatastoreId(datastoreId).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListDatastorePaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatastorePaths`: []DatastorePathInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListDatastorePaths`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatastorePathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datastoreId** | **string** | Datastore ID | 
 **vcenterId** | **string** | Virtual Center ID | 

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

> []HostOrClusterInfo ListHostsOrClusters(ctx).DatacenterId(datacenterId).VcenterId(vcenterId).Execute()

Lists all the hosts or clusters of the datacenter.



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
    datacenterId := "datacenterId_example" // string | Datacenter ID
    vcenterId := "vcenterId_example" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListHostsOrClusters(context.Background()).DatacenterId(datacenterId).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListHostsOrClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHostsOrClusters`: []HostOrClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListHostsOrClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostsOrClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datacenterId** | **string** | Datacenter ID | 
 **vcenterId** | **string** | Virtual Center ID | 

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

> []NetworkInterfaceCardInfo ListNetworkInterfaceCards(ctx).VcenterId(vcenterId).BaseSnapshotId(baseSnapshotId).BaseVmId(baseVmId).VmTemplateId(vmTemplateId).Execute()

Returns a list of network interface cards (NICs) suitable for configuration on a desktop pool/farm.



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
    vcenterId := "vcenterId_example" // string | Virtual Center ID
    baseSnapshotId := "baseSnapshotId_example" // string | Base Snapshot ID (optional)
    baseVmId := "baseVmId_example" // string | Base VM ID (optional)
    vmTemplateId := "vmTemplateId_example" // string | VM Template ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListNetworkInterfaceCards(context.Background()).VcenterId(vcenterId).BaseSnapshotId(baseSnapshotId).BaseVmId(baseVmId).VmTemplateId(vmTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListNetworkInterfaceCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkInterfaceCards`: []NetworkInterfaceCardInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListNetworkInterfaceCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkInterfaceCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vcenterId** | **string** | Virtual Center ID | 
 **baseSnapshotId** | **string** | Base Snapshot ID | 
 **baseVmId** | **string** | Base VM ID | 
 **vmTemplateId** | **string** | VM Template ID | 

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

> []NetworkLabelInfo ListNetworkLabels(ctx).HostOrClusterId(hostOrClusterId).VcenterId(vcenterId).NetworkType(networkType).Execute()

Retrieves all network labels on the given host or cluster



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
    hostOrClusterId := "hostOrClusterId_example" // string | Host or Cluster ID
    vcenterId := "vcenterId_example" // string | Virtual Center ID
    networkType := "networkType_example" // string | Network Type * NETWORK: Standard network. * OPAQUE_NETWORK: Opaque network. * DISTRUBUTED_VIRTUAL_PORT_GROUP: DVS Port group. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListNetworkLabels(context.Background()).HostOrClusterId(hostOrClusterId).VcenterId(vcenterId).NetworkType(networkType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListNetworkLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkLabels`: []NetworkLabelInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListNetworkLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostOrClusterId** | **string** | Host or Cluster ID | 
 **vcenterId** | **string** | Virtual Center ID | 
 **networkType** | **string** | Network Type * NETWORK: Standard network. * OPAQUE_NETWORK: Opaque network. * DISTRUBUTED_VIRTUAL_PORT_GROUP: DVS Port group. | 

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

> []ResourcePoolInfo ListResourcePools(ctx).HostOrClusterId(hostOrClusterId).VcenterId(vcenterId).Execute()

Lists all the resource pools from the vCenter for the given host or cluster.



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
    hostOrClusterId := "domain-c85" // string | Host or Cluster ID
    vcenterId := "d0325b13-2bf1-4fa4-b027-e780004f2d1e" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListResourcePools(context.Background()).HostOrClusterId(hostOrClusterId).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResourcePools`: []ResourcePoolInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListResourcePools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostOrClusterId** | **string** | Host or Cluster ID | 
 **vcenterId** | **string** | Virtual Center ID | 

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

> []VMFolderInfo ListVMFolders(ctx).DatacenterId(datacenterId).VcenterId(vcenterId).Execute()

Lists all the VM folders from the vCenter for the given datacenter.



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
    datacenterId := "datacenter-1" // string | Datacenter ID
    vcenterId := "d0325b13-2bf1-4fa4-b027-e780004f2d1e" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListVMFolders(context.Background()).DatacenterId(datacenterId).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListVMFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVMFolders`: []VMFolderInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListVMFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVMFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datacenterId** | **string** | Datacenter ID | 
 **vcenterId** | **string** | Virtual Center ID | 

### Return type

[**[]VMFolderInfo**](VMFolderInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVMTemplates

> []VMTemplateInfo ListVMTemplates(ctx).VcenterId(vcenterId).DatacenterId(datacenterId).Execute()

Lists all the VM templates from a vCenter or a datacenter for the given vCenter which may be suitable for full clone desktop pool creation.



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
    vcenterId := "vcenterId_example" // string | Virtual Center ID
    datacenterId := "datacenterId_example" // string | Datacenter ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListVMTemplates(context.Background()).VcenterId(vcenterId).DatacenterId(datacenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListVMTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVMTemplates`: []VMTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListVMTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVMTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vcenterId** | **string** | Virtual Center ID | 
 **datacenterId** | **string** | Datacenter ID | 

### Return type

[**[]VMTemplateInfo**](VMTemplateInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachines

> []VirtualMachineInfo ListVirtualMachines(ctx).VcenterId(vcenterId).Execute()

Lists all the VMs from a vCenter.



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
    vcenterId := "vcenterId_example" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ListVirtualMachines(context.Background()).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ListVirtualMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualMachines`: []VirtualMachineInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ListVirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vcenterId** | **string** | Virtual Center ID | 

### Return type

[**[]VirtualMachineInfo**](VirtualMachineInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listdatastores

> []DatastoreInfo Listdatastores(ctx).HostOrClusterId(hostOrClusterId).VcenterId(vcenterId).Execute()

Lists all the datastores from the vCenter for the given host or cluster.



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
    hostOrClusterId := "domain-c85" // string | Host or Cluster ID
    vcenterId := "d0325b13-2bf1-4fa4-b027-e780004f2d1e" // string | Virtual Center ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.Listdatastores(context.Background()).HostOrClusterId(hostOrClusterId).VcenterId(vcenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.Listdatastores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Listdatastores`: []DatastoreInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.Listdatastores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListdatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostOrClusterId** | **string** | Host or Cluster ID | 
 **vcenterId** | **string** | Virtual Center ID | 

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


## Unbind

> Unbind(ctx, id).Execute()

Unbind untrusted domain from the connection server.



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
    resp, r, err := api_client.ExternalApi.Unbind(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.Unbind``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnbindRequest struct via the builder pattern


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


## Update

> Update(ctx, id).Body(body).Execute()

Updates untrusted domain.



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
    body := *openapiclient.NewADDomainUpdateSpec() // ADDomainUpdateSpec | Untrusted domain object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.Update(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.Update``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ADDomainUpdateSpec**](ADDomainUpdateSpec.md) | Untrusted domain object to be updated. | 

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


## UpdateAuxiliaryAccounts

> []BulkItemResponseInfo UpdateAuxiliaryAccounts(ctx).Body(body).Execute()

Update auxiliary accounts of the untrusted domain



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
    body := *openapiclient.NewADDomainAuxiliaryAccountUpdateSpec([]openapiclient.AuxiliaryAccountUpdateData{*openapiclient.NewAuxiliaryAccountUpdateData("1f95a15c-a7a5-4584-963f-2c3f5355b49f", []string{"Password_example"})}) // ADDomainAuxiliaryAccountUpdateSpec | Specification to update auxiliary accounts.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.UpdateAuxiliaryAccounts(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.UpdateAuxiliaryAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuxiliaryAccounts`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.UpdateAuxiliaryAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuxiliaryAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ADDomainAuxiliaryAccountUpdateSpec**](ADDomainAuxiliaryAccountUpdateSpec.md) | Specification to update auxiliary accounts. | 

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


## ValidateADUserEncryptedCredentials

> ADUserInfo ValidateADUserEncryptedCredentials(ctx).Body(body).Execute()

Validates the encrypted credentials of AD User



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
    body := *openapiclient.NewADUserEncryptedCredentialSpec("9QYxpdXrcrOGPSSz1/K/pJu8QlYT7pDkaKg/rb3hlw4=", "ut0yGq7CbhtBUGfW3ngjkvjQ2cc=", "Qvvjglg5iZinyuldroueo/hQFyqydMMDJPmfYGgIebqxbU9chJ9I8iM9SCBRHSkSW9y+RM", "testuser or testuser@example.com") // ADUserEncryptedCredentialSpec | AD user encrypted credentials object to be validated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalApi.ValidateADUserEncryptedCredentials(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalApi.ValidateADUserEncryptedCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateADUserEncryptedCredentials`: ADUserInfo
    fmt.Fprintf(os.Stdout, "Response from `ExternalApi.ValidateADUserEncryptedCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateADUserEncryptedCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ADUserEncryptedCredentialSpec**](ADUserEncryptedCredentialSpec.md) | AD user encrypted credentials object to be validated. | 

### Return type

[**ADUserInfo**](ADUserInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

