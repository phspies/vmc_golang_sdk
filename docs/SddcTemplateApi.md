# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcTemplatesGet**](SddcTemplateApi.md#OrgsOrgSddcTemplatesGet) | **Get** /orgs/{org}/sddc-templates | List all available SDDC configuration templates in an organization
[**OrgsOrgSddcTemplatesTemplateIdDelete**](SddcTemplateApi.md#OrgsOrgSddcTemplatesTemplateIdDelete) | **Delete** /orgs/{org}/sddc-templates/{templateId} | Delete SDDC template identified by given id.
[**OrgsOrgSddcTemplatesTemplateIdGet**](SddcTemplateApi.md#OrgsOrgSddcTemplatesTemplateIdGet) | **Get** /orgs/{org}/sddc-templates/{templateId} | Get configuration template by given template id.
[**OrgsOrgSddcsSddcSddcTemplateGet**](SddcTemplateApi.md#OrgsOrgSddcsSddcSddcTemplateGet) | **Get** /orgs/{org}/sddcs/{sddc}/sddc-template | Get configuration template for an SDDC

# **OrgsOrgSddcTemplatesGet**
> []SddcTemplate OrgsOrgSddcTemplatesGet(ctx, org)
List all available SDDC configuration templates in an organization

List all available SDDC configuration templates in an organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 

### Return type

[**[]SddcTemplate**](SddcTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcTemplatesTemplateIdDelete**
> Task OrgsOrgSddcTemplatesTemplateIdDelete(ctx, org, templateId)
Delete SDDC template identified by given id.

Delete SDDC template identified by given id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **templateId** | **string**| SDDC Template identifier | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcTemplatesTemplateIdGet**
> SddcTemplate OrgsOrgSddcTemplatesTemplateIdGet(ctx, org, templateId)
Get configuration template by given template id.

Get configuration template by given template id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **templateId** | **string**| SDDC Template identifier | 

### Return type

[**SddcTemplate**](SddcTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcSddcTemplateGet**
> SddcTemplate OrgsOrgSddcsSddcSddcTemplateGet(ctx, org, sddc)
Get configuration template for an SDDC

Get configuration template for  an SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**SddcTemplate**](SddcTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

