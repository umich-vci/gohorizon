# \InventoryApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMachines**](InventoryApi.md#AddMachines) | **Post** /inventory/v1/desktop-pools/{id}/action/add-machines | Adds machines to the given manual desktop pool.
[**AddMachinesByName**](InventoryApi.md#AddMachinesByName) | **Post** /inventory/v1/desktop-pools/{id}/action/add-machines-by-name | Adds the named machines to the given desktop pool.
[**AssignUsers**](InventoryApi.md#AssignUsers) | **Post** /inventory/v1/machines/{id}/action/assign-users | Assigns the specified users to the machine.
[**CreateApplicationPool**](InventoryApi.md#CreateApplicationPool) | **Post** /inventory/v1/application-pools | Creates an application pool.
[**DeleteApplicationPool**](InventoryApi.md#DeleteApplicationPool) | **Delete** /inventory/v1/application-pools/{id} | Deletes application pool.
[**DeleteMachine**](InventoryApi.md#DeleteMachine) | **Delete** /inventory/v1/machines/{id} | Deletes the machine.
[**DeleteMachines**](InventoryApi.md#DeleteMachines) | **Delete** /inventory/v1/machines | Deletes the specified machines.
[**DisconnectSessions**](InventoryApi.md#DisconnectSessions) | **Post** /inventory/v1/sessions/action/disconnect | Disconnects user sessions
[**EnterMaintenance**](InventoryApi.md#EnterMaintenance) | **Post** /inventory/v1/machines/action/enter-maintenance | Puts the machines into maintenance mode.
[**ExitMaintenance**](InventoryApi.md#ExitMaintenance) | **Post** /inventory/v1/machines/action/exit-maintenance | Puts the machines out of maintenance mode.
[**GetApplicationIcon**](InventoryApi.md#GetApplicationIcon) | **Get** /inventory/v1/application-icons/{id} | Gets application icon.
[**GetApplicationPool**](InventoryApi.md#GetApplicationPool) | **Get** /inventory/v1/application-pools/{id} | Gets application pool.
[**GetDesktopPool**](InventoryApi.md#GetDesktopPool) | **Get** /inventory/v1/desktop-pools/{id} | Gets the Desktop Pool information.
[**GetDesktopPoolV2**](InventoryApi.md#GetDesktopPoolV2) | **Get** /inventory/v2/desktop-pools/{id} | Gets the desktop pool information.
[**GetFarm**](InventoryApi.md#GetFarm) | **Get** /inventory/v1/farms/{id} | Gets the Farm information.
[**GetMachineUsingGET**](InventoryApi.md#GetMachineUsingGET) | **Get** /inventory/v1/machines/{id} | Gets the Machine information.
[**GetSessionInfo**](InventoryApi.md#GetSessionInfo) | **Get** /inventory/v1/sessions/{id} | Gets the Session information.
[**ListApplicationIcons**](InventoryApi.md#ListApplicationIcons) | **Get** /inventory/v1/application-icons | Lists the application icons for the given application pool.
[**ListApplicationPools**](InventoryApi.md#ListApplicationPools) | **Get** /inventory/v1/application-pools | Lists the application pools in the environment.
[**ListDesktopPools**](InventoryApi.md#ListDesktopPools) | **Get** /inventory/v1/desktop-pools | Lists the Desktop Pools in the environment.
[**ListDesktopPoolsV2**](InventoryApi.md#ListDesktopPoolsV2) | **Get** /inventory/v2/desktop-pools | Lists the desktop pools in the environment.
[**ListFarms**](InventoryApi.md#ListFarms) | **Get** /inventory/v1/farms | Lists the Farms in the environment.
[**ListInstalledApplications**](InventoryApi.md#ListInstalledApplications) | **Get** /inventory/v1/desktop-pools/{id}/installed-applications | Lists the installed applications on the given desktop pool.
[**ListInstalledApplications1**](InventoryApi.md#ListInstalledApplications1) | **Get** /inventory/v1/farms/{id}/installed-applications | Lists the installed applications on the given farm.
[**ListMachinesUsingGET**](InventoryApi.md#ListMachinesUsingGET) | **Get** /inventory/v1/machines | Lists the Machines in the environment.
[**ListSessionInfo**](InventoryApi.md#ListSessionInfo) | **Get** /inventory/v1/sessions | Lists the Session information in the environment.
[**LogOffSessions**](InventoryApi.md#LogOffSessions) | **Post** /inventory/v1/sessions/action/logoff | Logs off user sessions, if they are not locked.
[**RebuildMachines**](InventoryApi.md#RebuildMachines) | **Post** /inventory/v1/machines/action/rebuild | Rebuilds the specified machines.
[**RecoverMachines**](InventoryApi.md#RecoverMachines) | **Post** /inventory/v1/machines/action/recover | Recovers the specified machines.
[**RemoveMachines**](InventoryApi.md#RemoveMachines) | **Post** /inventory/v1/desktop-pools/{id}/action/remove-machines | Removes machines from the given manual desktop pool.
[**ResetMachines**](InventoryApi.md#ResetMachines) | **Post** /inventory/v1/machines/action/reset | Resets the specified machines.
[**ResetSessions**](InventoryApi.md#ResetSessions) | **Post** /inventory/v1/sessions/action/reset | Resets machine of user sessions. The machine must be managed by Virtual Center and the session cannot be an application or an RDS desktop session.
[**RestartMachines**](InventoryApi.md#RestartMachines) | **Post** /inventory/v1/machines/action/restart | Restarts the specified machines.
[**RestartSessions**](InventoryApi.md#RestartSessions) | **Post** /inventory/v1/sessions/action/restart | Restarts machine of user sessions. The machine must be managed by Virtual Center and the session cannot be an application or an RDS desktop session.
[**SendMessageToSessions**](InventoryApi.md#SendMessageToSessions) | **Post** /inventory/v1/sessions/action/send-message | Sends the message to user sessions
[**UnassignUsers**](InventoryApi.md#UnassignUsers) | **Post** /inventory/v1/machines/{id}/action/unassign-users | Un-assigns the specified users from the machine.
[**UpdateApplicationPool**](InventoryApi.md#UpdateApplicationPool) | **Put** /inventory/v1/application-pools/{id} | Updates application pool.



## AddMachines

> []BulkItemResponseInfo AddMachines(ctx, id, body)

Adds machines to the given manual desktop pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**[]string**](string.md)| List of Machine Ids representing the machines to be added to the desktop pool. | 

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


## AddMachinesByName

> []BulkItemResponseInfo AddMachinesByName(ctx, id, body)

Adds the named machines to the given desktop pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**[]MachineSpecifiedName**](MachineSpecifiedName.md)| List of MachineSpecifiedName representing the machines to be added to the desktop pool. | 

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


## AssignUsers

> []BulkItemResponseInfo AssignUsers(ctx, id, body)

Assigns the specified users to the machine.

Each response entity in the result corresponds to a user SID in the input.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**[]string**](string.md)| List of User SIDs representing the users to be assigned to the machine. | 

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


## CreateApplicationPool

> CreateApplicationPool(ctx, body)

Creates an application pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ApplicationPoolCreateSpec**](ApplicationPoolCreateSpec.md)| Application pool object to be created. | 

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


## DeleteApplicationPool

> DeleteApplicationPool(ctx, id)

Deletes application pool.

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


## DeleteMachine

> DeleteMachine(ctx, id, body)

Deletes the machine.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**MachineDeleteData**](MachineDeleteData.md)| The specification applicable to deleting the machine. | 

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


## DeleteMachines

> []BulkItemResponseInfo DeleteMachines(ctx, body)

Deletes the specified machines.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MachineDeleteSpec**](MachineDeleteSpec.md)| The machines and specification for deletion. | 

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


## DisconnectSessions

> []BulkItemResponseInfo DisconnectSessions(ctx, body)

Disconnects user sessions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of session ids to be disconnected. | 

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


## EnterMaintenance

> []BulkItemResponseInfo EnterMaintenance(ctx, body)

Puts the machines into maintenance mode.

When in maintenance mode, users cannot access the machines. Each response entity in the result corresponds to a machine id in the input.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of Machine Ids representing the machines to be put into maintenance mode. | 

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


## ExitMaintenance

> []BulkItemResponseInfo ExitMaintenance(ctx, body)

Puts the machines out of maintenance mode.

Each response entity in the result corresponds to a machine id in the input.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of Machine Ids representing the machines to be put out of maintenance mode. | 

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


## GetApplicationIcon

> ApplicationIconInfo GetApplicationIcon(ctx, id)

Gets application icon.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**ApplicationIconInfo**](ApplicationIconInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationPool

> ApplicationPoolInfo GetApplicationPool(ctx, id)

Gets application pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**ApplicationPoolInfo**](ApplicationPoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPool

> DesktopPoolInfo GetDesktopPool(ctx, id)

Gets the Desktop Pool information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**DesktopPoolInfo**](DesktopPoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPoolV2

> DesktopPoolInfoV2 GetDesktopPoolV2(ctx, id)

Gets the desktop pool information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**DesktopPoolInfoV2**](DesktopPoolInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFarm

> FarmInfo GetFarm(ctx, id)

Gets the Farm information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**FarmInfo**](FarmInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachineUsingGET

> MachineInfo GetMachineUsingGET(ctx, id)

Gets the Machine information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**MachineInfo**](MachineInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionInfo

> SessionInfo GetSessionInfo(ctx, id)

Gets the Session information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**SessionInfo**](SessionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationIcons

> []ApplicationIconInfo ListApplicationIcons(ctx, applicationPoolId)

Lists the application icons for the given application pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPoolId** | **string**| Application Pool ID | 

### Return type

[**[]ApplicationIconInfo**](ApplicationIconInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationPools

> []ApplicationPoolInfo ListApplicationPools(ctx, )

Lists the application pools in the environment.

This API supports <b>Pagination</b> and <b>Filters.</b><br/>For Pagination, optional query params of 'page' and 'size' needs to be send.<br/>For filters, refer to 'ApplicationPoolInfo' model description to find supported filters on specific field.<br/>For full information on using filters, refer to 'Horizon Server REST Pagination and Filter Guide' of 'VMware Horizon Server API' in code.vmware.com

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ApplicationPoolInfo**](ApplicationPoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPools

> []DesktopPoolInfo ListDesktopPools(ctx, )

Lists the Desktop Pools in the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DesktopPoolInfo**](DesktopPoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPoolsV2

> []DesktopPoolInfoV2 ListDesktopPoolsV2(ctx, )

Lists the desktop pools in the environment.

This API supports <b>Pagination</b>.<br/>For Pagination, optional query params of 'page' and 'size' needs to be send.<br/>

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DesktopPoolInfoV2**](DesktopPoolInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFarms

> []FarmInfo ListFarms(ctx, )

Lists the Farms in the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]FarmInfo**](FarmInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledApplications

> []InstalledApplicationInfo ListInstalledApplications(ctx, id)

Lists the installed applications on the given desktop pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**[]InstalledApplicationInfo**](InstalledApplicationInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledApplications1

> []InstalledApplicationInfo ListInstalledApplications1(ctx, id)

Lists the installed applications on the given farm.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**[]InstalledApplicationInfo**](InstalledApplicationInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMachinesUsingGET

> []MachineInfo ListMachinesUsingGET(ctx, )

Lists the Machines in the environment.

This API supports <b>Pagination</b> and <b>Filters.</b><br/>For Pagination, optional query params of 'page' and 'size' need to be sent.<br/>For Filters, refer to 'MachineInfo' model description to find supported filters on specific field.<br/>For full information on using Filters, refer to 'Horizon Server REST Pagination and Filter Guide' of 'VMware Horizon Server API' in code.vmware.com

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]MachineInfo**](MachineInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionInfo

> []SessionInfo ListSessionInfo(ctx, )

Lists the Session information in the environment.

This API supports <b>Pagination</b>.<br/>For Pagination, optional query params of 'page' and 'size' needs to be send.<br/>

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SessionInfo**](SessionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogOffSessions

> []BulkItemResponseInfo LogOffSessions(ctx, body, optional)

Logs off user sessions, if they are not locked.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of session ids to be logged off. | 
 **optional** | ***LogOffSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LogOffSessionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forced** | **optional.Bool**| Indicates to Log off session forcibly.  If passed as \&quot;true\&quot;, then sessions are logoff forcibly, even if they are locked.  If passed as \&quot;false\&quot; or not passed at all, then sessions will be normally logged off, if they are not locked. | [default to false]

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


## RebuildMachines

> []BulkItemResponseInfo RebuildMachines(ctx, body)

Rebuilds the specified machines.

Each response entity in the result corresponds to a machine id in the input.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of Machine Ids representing the machines to be rebuilt. | 

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


## RecoverMachines

> []BulkItemResponseInfo RecoverMachines(ctx, body)

Recovers the specified machines.

Each response entity in the result corresponds to a machine id in the input.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of Machine Ids representing the machines to be recovered. | 

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


## RemoveMachines

> []BulkItemResponseInfo RemoveMachines(ctx, id, body)

Removes machines from the given manual desktop pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**[]string**](string.md)| List of Machine Ids representing the machines to be removed from the desktop pool. | 

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


## ResetMachines

> []BulkItemResponseInfo ResetMachines(ctx, body)

Resets the specified machines.

Each response entity in the result corresponds to a machine id in the input.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of Machine Ids representing the machines to be reset. | 

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


## ResetSessions

> []BulkItemResponseInfo ResetSessions(ctx, body)

Resets machine of user sessions. The machine must be managed by Virtual Center and the session cannot be an application or an RDS desktop session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of session ids to be reset. | 

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


## RestartMachines

> []BulkItemResponseInfo RestartMachines(ctx, body)

Restarts the specified machines.

Each response entity in the result corresponds to a machine id in the input.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of Machine Ids representing the machines to be restarted. | 

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


## RestartSessions

> []BulkItemResponseInfo RestartSessions(ctx, body)

Restarts machine of user sessions. The machine must be managed by Virtual Center and the session cannot be an application or an RDS desktop session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]string**](string.md)| List of session ids to be restarted. | 

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


## SendMessageToSessions

> []BulkItemResponseInfo SendMessageToSessions(ctx, body)

Sends the message to user sessions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SessionSendMessageSpec**](SessionSendMessageSpec.md)| Message information object to be sent to sessions. | 

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


## UnassignUsers

> []BulkItemResponseInfo UnassignUsers(ctx, id, body)

Un-assigns the specified users from the machine.

Each response entity in the result corresponds to a user SID in the input.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**[]string**](string.md)| List of User SIDs representing the users to be un-assigned from the machine. | 

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


## UpdateApplicationPool

> UpdateApplicationPool(ctx, id, body)

Updates application pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
**body** | [**ApplicationPoolUpdateSpec**](ApplicationPoolUpdateSpec.md)| Application pool object to be updated. | 

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

