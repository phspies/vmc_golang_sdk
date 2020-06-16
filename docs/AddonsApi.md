# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsGet**](AddonsApi.md#OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsGet) | **Get** /orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials | List all the credentials assoicated with an addon type with in a SDDC
[**OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNameGet**](AddonsApi.md#OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNameGet) | **Get** /orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name} | Get credential details by name
[**OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNamePut**](AddonsApi.md#OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNamePut) | **Put** /orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name} | Update credential details
[**OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsPost**](AddonsApi.md#OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsPost) | **Post** /orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials | Associate a new add on credentials with the SDDC such as HCX

# **OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsGet**
> []NewCredentials OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsGet(ctx, org, sddcId, addonType)
List all the credentials assoicated with an addon type with in a SDDC

List all the credentials assoicated with an addon type with in a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | [**string**](.md)| Org id of the associated SDDC | 
  **sddcId** | [**string**](.md)| Id of the SDDC | 
  **addonType** | **string**| Add on type | 

### Return type

[**[]NewCredentials**](NewCredentials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNameGet**
> NewCredentials OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNameGet(ctx, org, sddcId, addonType, name)
Get credential details by name

Get credential details by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | [**string**](.md)| Org id of the associated SDDC | 
  **sddcId** | [**string**](.md)| Id of the SDDC | 
  **addonType** | **string**| Add on type | 
  **name** | **string**| name of the credentials | 

### Return type

[**NewCredentials**](NewCredentials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNamePut**
> NewCredentials OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNamePut(ctx, body, org, sddcId, addonType, name)
Update credential details

Update credential details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCredentials**](UpdateCredentials.md)| Credentials update payload | 
  **org** | [**string**](.md)| Org id of the associated SDDC | 
  **sddcId** | [**string**](.md)| Id of the SDDC | 
  **addonType** | **string**| Add on type | 
  **name** | **string**| name of the credentials | 

### Return type

[**NewCredentials**](NewCredentials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsPost**
> NewCredentials OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsPost(ctx, body, org, sddcId, addonType)
Associate a new add on credentials with the SDDC such as HCX

Associated a new add on credentials with the SDDC such as HCX

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewCredentials**](NewCredentials.md)| Credentials creation payload | 
  **org** | [**string**](.md)| Org id of the associated SDDC | 
  **sddcId** | [**string**](.md)| Id of the SDDC | 
  **addonType** | **string**| Add on type | 

### Return type

[**NewCredentials**](NewCredentials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

