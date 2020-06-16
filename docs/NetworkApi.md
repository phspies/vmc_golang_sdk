# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsSddcNetworkingConnectivityTestsGet**](NetworkApi.md#OrgsOrgSddcsSddcNetworkingConnectivityTestsGet) | **Get** /orgs/{org}/sddcs/{sddc}/networking/connectivity-tests | Retrieve metadata for connectivity tests.
[**OrgsOrgSddcsSddcNetworkingConnectivityTestsPost**](NetworkApi.md#OrgsOrgSddcsSddcNetworkingConnectivityTestsPost) | **Post** /orgs/{org}/sddcs/{sddc}/networking/connectivity-tests | ConnectivityValidationGroupResultWrapper will be  available at task.params[&#x27;test_result&#x27;].

# **OrgsOrgSddcsSddcNetworkingConnectivityTestsGet**
> ConnectivityValidationGroups OrgsOrgSddcsSddcNetworkingConnectivityTestsGet(ctx, org, sddc)
Retrieve metadata for connectivity tests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**ConnectivityValidationGroups**](ConnectivityValidationGroups.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcNetworkingConnectivityTestsPost**
> Task OrgsOrgSddcsSddcNetworkingConnectivityTestsPost(ctx, body, org, sddc, action)
ConnectivityValidationGroupResultWrapper will be  available at task.params['test_result'].

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectivityValidationGroup**](ConnectivityValidationGroup.md)| request information | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **action** | **string**| If &#x3D; &#x27;start&#x27;, start connectivity tests.  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

