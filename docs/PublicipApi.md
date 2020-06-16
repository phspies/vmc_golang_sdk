# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsSddcPublicipsGet**](PublicipApi.md#OrgsOrgSddcsSddcPublicipsGet) | **Get** /orgs/{org}/sddcs/{sddc}/publicips | List all public IPs for a SDDC
[**OrgsOrgSddcsSddcPublicipsIdDelete**](PublicipApi.md#OrgsOrgSddcsSddcPublicipsIdDelete) | **Delete** /orgs/{org}/sddcs/{sddc}/publicips/{id} | Free one public IP for a SDDC
[**OrgsOrgSddcsSddcPublicipsIdGet**](PublicipApi.md#OrgsOrgSddcsSddcPublicipsIdGet) | **Get** /orgs/{org}/sddcs/{sddc}/publicips/{id} | Get one public IP for a SDDC
[**OrgsOrgSddcsSddcPublicipsIdPatch**](PublicipApi.md#OrgsOrgSddcsSddcPublicipsIdPatch) | **Patch** /orgs/{org}/sddcs/{sddc}/publicips/{id} | Attach or detach a public IP to workload VM for a SDDC
[**OrgsOrgSddcsSddcPublicipsPost**](PublicipApi.md#OrgsOrgSddcsSddcPublicipsPost) | **Post** /orgs/{org}/sddcs/{sddc}/publicips | Allocate public IPs for a SDDC

# **OrgsOrgSddcsSddcPublicipsGet**
> []SddcPublicIp OrgsOrgSddcsSddcPublicipsGet(ctx, org, sddc)
List all public IPs for a SDDC

list all public IPs for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**[]SddcPublicIp**](SddcPublicIp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsIdDelete**
> Task OrgsOrgSddcsSddcPublicipsIdDelete(ctx, org, sddc, id)
Free one public IP for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **id** | **string**| ip allocation id | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsIdGet**
> SddcPublicIp OrgsOrgSddcsSddcPublicipsIdGet(ctx, org, sddc, id)
Get one public IP for a SDDC

Get one public IP for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **id** | **string**| ip allocation id | 

### Return type

[**SddcPublicIp**](SddcPublicIp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsIdPatch**
> Task OrgsOrgSddcsSddcPublicipsIdPatch(ctx, body, org, sddc, id, action)
Attach or detach a public IP to workload VM for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SddcPublicIp**](SddcPublicIp.md)| SddcPublicIp object to update | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **id** | **string**| ip allocation id | 
  **action** | **string**| Type of action as &#x27;attach&#x27;, &#x27;detach&#x27;, &#x27;reattach&#x27;, or &#x27;rename&#x27;. For &#x27;attch&#x27;, the public IP must not be attached and &#x27;associated_private_ip&#x27; in the payload needs to be set with a workload VM private IP. For &#x27;detach&#x27;, the public IP must be attached and &#x27;associated_private_ip&#x27; in the payload should not be set with any value. For &#x27;reattach&#x27;, the public IP must be attached and &#x27;associated_private_ip&#x27; in the payload needs to be set with a new workload VM private IP. For &#x27;rename&#x27;, the &#x27;name&#x27; in the payload needs to have a new name string.  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsPost**
> Task OrgsOrgSddcsSddcPublicipsPost(ctx, body, org, sddc)
Allocate public IPs for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SddcAllocatePublicIpSpec**](SddcAllocatePublicIpSpec.md)| allocation spec | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

