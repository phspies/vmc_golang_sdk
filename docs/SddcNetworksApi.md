# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsSddcNetworks40SddcNetworksGet**](SddcNetworksApi.md#OrgsOrgSddcsSddcNetworks40SddcNetworksGet) | **Get** /orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks | 
[**OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdDelete**](SddcNetworksApi.md#OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdDelete) | **Delete** /orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId} | 
[**OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdGet**](SddcNetworksApi.md#OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdGet) | **Get** /orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId} | 
[**OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdPut**](SddcNetworksApi.md#OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdPut) | **Put** /orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId} | 
[**OrgsOrgSddcsSddcNetworks40SddcNetworksPost**](SddcNetworksApi.md#OrgsOrgSddcsSddcNetworks40SddcNetworksPost) | **Post** /orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks | 

# **OrgsOrgSddcsSddcNetworks40SddcNetworksGet**
> DataPageSddcNetwork OrgsOrgSddcsSddcNetworks40SddcNetworksGet(ctx, org, sddc, optional)


Retrieve all networks in an SDDC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
 **optional** | ***SddcNetworksApiOrgsOrgSddcsSddcNetworks40SddcNetworksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcNetworksApiOrgsOrgSddcsSddcNetworks40SddcNetworksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| Page size for pagination. | 
 **startIndex** | **optional.Int32**| Start index of page. | 
 **prevSddcNetworkId** | **optional.String**| Previous logical network id. | 
 **sortOrderAscending** | **optional.Bool**| Sort order ascending. | 

### Return type

[**DataPageSddcNetwork**](DataPageSddcNetwork.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdDelete**
> OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdDelete(ctx, org, sddc, networkId)


Delete a network in an SDDC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **networkId** | **string**| Logical Network Identifier | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdGet**
> SddcNetwork OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdGet(ctx, org, sddc, networkId)


Retrieve information about a network in an SDDC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **networkId** | **string**| Logical Network Identifier | 

### Return type

[**SddcNetwork**](sddcNetwork.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdPut**
> OrgsOrgSddcsSddcNetworks40SddcNetworksNetworkIdPut(ctx, body, org, sddc, networkId)


Modify a network in an SDDC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SddcNetwork**](SddcNetwork.md)|  | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **networkId** | **string**| Logical Network Identifier | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcNetworks40SddcNetworksPost**
> OrgsOrgSddcsSddcNetworks40SddcNetworksPost(ctx, body, org, sddc)


Create a network in an SDDC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SddcNetwork**](SddcNetwork.md)|  | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

