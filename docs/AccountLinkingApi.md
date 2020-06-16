# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgAccountLinkCompatibleSubnetsAsyncGet**](AccountLinkingApi.md#OrgsOrgAccountLinkCompatibleSubnetsAsyncGet) | **Get** /orgs/{org}/account-link/compatible-subnets-async | Gets a customer&#x27;s compatible subnets for account linking via a task. The information is returned as a member of the task (found in task.params[&#x27;subnet_list_result&#x27;] when you are notified it is complete), and it&#x27;s documented under ref /definitions/AwsCompatibleSubnets
[**OrgsOrgAccountLinkCompatibleSubnetsAsyncPost**](AccountLinkingApi.md#OrgsOrgAccountLinkCompatibleSubnetsAsyncPost) | **Post** /orgs/{org}/account-link/compatible-subnets-async | Sets which subnet to use to link accounts and finishes the linking process via a task
[**OrgsOrgAccountLinkCompatibleSubnetsGet**](AccountLinkingApi.md#OrgsOrgAccountLinkCompatibleSubnetsGet) | **Get** /orgs/{org}/account-link/compatible-subnets | Gets a customer&#x27;s compatible subnets for account linking
[**OrgsOrgAccountLinkCompatibleSubnetsPost**](AccountLinkingApi.md#OrgsOrgAccountLinkCompatibleSubnetsPost) | **Post** /orgs/{org}/account-link/compatible-subnets | Sets which subnet to use to link accounts and finishes the linking process
[**OrgsOrgAccountLinkConnectedAccountsGet**](AccountLinkingApi.md#OrgsOrgAccountLinkConnectedAccountsGet) | **Get** /orgs/{org}/account-link/connected-accounts | Get a list of connected accounts
[**OrgsOrgAccountLinkConnectedAccountsLinkedAccountPathIdDelete**](AccountLinkingApi.md#OrgsOrgAccountLinkConnectedAccountsLinkedAccountPathIdDelete) | **Delete** /orgs/{org}/account-link/connected-accounts/{linkedAccountPathId} | Delete a particular connected (linked) account.
[**OrgsOrgAccountLinkGet**](AccountLinkingApi.md#OrgsOrgAccountLinkGet) | **Get** /orgs/{org}/account-link | Gets a link that can be used on a customer&#x27;s account to start the linking process.
[**OrgsOrgAccountLinkMapCustomerZonesPost**](AccountLinkingApi.md#OrgsOrgAccountLinkMapCustomerZonesPost) | **Post** /orgs/{org}/account-link/map-customer-zones | Creates a task to re-map customer&#x27;s datacenters across zones.
[**OrgsOrgAccountLinkSddcConnectionsGet**](AccountLinkingApi.md#OrgsOrgAccountLinkSddcConnectionsGet) | **Get** /orgs/{org}/account-link/sddc-connections | Get a list of SDDC connections currently setup for the customer&#x27;s organization.

# **OrgsOrgAccountLinkCompatibleSubnetsAsyncGet**
> Task OrgsOrgAccountLinkCompatibleSubnetsAsyncGet(ctx, org, optional)
Gets a customer's compatible subnets for account linking via a task. The information is returned as a member of the task (found in task.params['subnet_list_result'] when you are notified it is complete), and it's documented under ref /definitions/AwsCompatibleSubnets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***AccountLinkingApiOrgsOrgAccountLinkCompatibleSubnetsAsyncGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLinkingApiOrgsOrgAccountLinkCompatibleSubnetsAsyncGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkedAccountId** | **optional.String**| The linked connected account identifier | 
 **region** | **optional.String**| The region of the cloud resources to work in | 
 **sddc** | **optional.String**| sddc | 
 **instanceType** | **optional.String**| The server instance type to be used. | 
 **sddcType** | **optional.String**| The sddc type to be used. (1NODE, SingleAZ, MultiAZ) | 
 **numOfHosts** | **optional.Int32**| The number of hosts | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgAccountLinkCompatibleSubnetsAsyncPost**
> Task OrgsOrgAccountLinkCompatibleSubnetsAsyncPost(ctx, body, org)
Sets which subnet to use to link accounts and finishes the linking process via a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AwsSubnet**](AwsSubnet.md)| The subnet chosen by the customer | 
  **org** | **string**| Organization identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgAccountLinkCompatibleSubnetsGet**
> AwsCompatibleSubnets OrgsOrgAccountLinkCompatibleSubnetsGet(ctx, org, optional)
Gets a customer's compatible subnets for account linking

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***AccountLinkingApiOrgsOrgAccountLinkCompatibleSubnetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLinkingApiOrgsOrgAccountLinkCompatibleSubnetsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkedAccountId** | **optional.String**| The linked connected account identifier | 
 **region** | **optional.String**| The region of the cloud resources to work in | 
 **sddc** | **optional.String**| sddc | 
 **forceRefresh** | **optional.Bool**| When true, forces the mappings for datacenters to be refreshed for the connected account. | 
 **instanceType** | **optional.String**| The server instance type to be used. | 
 **sddcType** | **optional.String**| The sddc type to be used. (1NODE, SingleAZ, MultiAZ) | 
 **numOfHosts** | **optional.Int32**| The number of hosts | 

### Return type

[**AwsCompatibleSubnets**](AwsCompatibleSubnets.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgAccountLinkCompatibleSubnetsPost**
> AwsSubnet OrgsOrgAccountLinkCompatibleSubnetsPost(ctx, org)
Sets which subnet to use to link accounts and finishes the linking process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 

### Return type

[**AwsSubnet**](AwsSubnet.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgAccountLinkConnectedAccountsGet**
> []AwsCustomerConnectedAccount OrgsOrgAccountLinkConnectedAccountsGet(ctx, org, optional)
Get a list of connected accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***AccountLinkingApiOrgsOrgAccountLinkConnectedAccountsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLinkingApiOrgsOrgAccountLinkConnectedAccountsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | **optional.String**| The cloud provider of the SDDC (AWS or ZeroCloud). Default value is AWS. | 

### Return type

[**[]AwsCustomerConnectedAccount**](AwsCustomerConnectedAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgAccountLinkConnectedAccountsLinkedAccountPathIdDelete**
> AwsCustomerConnectedAccount OrgsOrgAccountLinkConnectedAccountsLinkedAccountPathIdDelete(ctx, org, linkedAccountPathId, optional)
Delete a particular connected (linked) account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **linkedAccountPathId** | **string**| The linked connected account identifier | 
 **optional** | ***AccountLinkingApiOrgsOrgAccountLinkConnectedAccountsLinkedAccountPathIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLinkingApiOrgsOrgAccountLinkConnectedAccountsLinkedAccountPathIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceEvenWhenSddcPresent** | **optional.Bool**| When true, forcibly removes a connected account even when SDDC&#x27;s are still linked to it. | 

### Return type

[**AwsCustomerConnectedAccount**](AwsCustomerConnectedAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgAccountLinkGet**
> OrgsOrgAccountLinkGet(ctx, org)
Gets a link that can be used on a customer's account to start the linking process.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgAccountLinkMapCustomerZonesPost**
> Task OrgsOrgAccountLinkMapCustomerZonesPost(ctx, body, org)
Creates a task to re-map customer's datacenters across zones.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MapZonesRequest**](MapZonesRequest.md)| The zones request information about who to map and what to map. | 
  **org** | **string**| Organization identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgAccountLinkSddcConnectionsGet**
> []AwsSddcConnection OrgsOrgAccountLinkSddcConnectionsGet(ctx, org, optional)
Get a list of SDDC connections currently setup for the customer's organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***AccountLinkingApiOrgsOrgAccountLinkSddcConnectionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLinkingApiOrgsOrgAccountLinkSddcConnectionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sddc** | **optional.String**| sddc | 

### Return type

[**[]AwsSddcConnection**](AwsSddcConnection.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

