# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsSddcDnsPrivatePut**](DnsApi.md#OrgsOrgSddcsSddcDnsPrivatePut) | **Put** /orgs/{org}/sddcs/{sddc}/dns/private | Update the DNS records of management VMs to use private IP addresses
[**OrgsOrgSddcsSddcDnsPublicPut**](DnsApi.md#OrgsOrgSddcsSddcDnsPublicPut) | **Put** /orgs/{org}/sddcs/{sddc}/dns/public | Update the DNS records of management VMs to use public IP addresses
[**OrgsOrgSddcsSddcManagementVmsManagementVmIdDnsIpTypePut**](DnsApi.md#OrgsOrgSddcsSddcManagementVmsManagementVmIdDnsIpTypePut) | **Put** /orgs/{org}/sddcs/{sddc}/management-vms/{managementVmId}/dns/{ipType} | Update the DNS records of management VMs to use public or private IP addresses

# **OrgsOrgSddcsSddcDnsPrivatePut**
> Task OrgsOrgSddcsSddcDnsPrivatePut(ctx, org, sddc)
Update the DNS records of management VMs to use private IP addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcDnsPublicPut**
> Task OrgsOrgSddcsSddcDnsPublicPut(ctx, org, sddc)
Update the DNS records of management VMs to use public IP addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcManagementVmsManagementVmIdDnsIpTypePut**
> Task OrgsOrgSddcsSddcManagementVmsManagementVmIdDnsIpTypePut(ctx, org, sddc, managementVmId, ipType)
Update the DNS records of management VMs to use public or private IP addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **managementVmId** | **string**| the management VM ID | 
  **ipType** | **string**| the ip type to associate with FQDN in DNS | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

