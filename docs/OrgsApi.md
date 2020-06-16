# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsGet**](OrgsApi.md#OrgsGet) | **Get** /orgs | Get organizations associated with calling user.
[**OrgsOrgGet**](OrgsApi.md#OrgsOrgGet) | **Get** /orgs/{org} | Get details of organization
[**OrgsOrgPaymentMethodsGet**](OrgsApi.md#OrgsOrgPaymentMethodsGet) | **Get** /orgs/{org}/payment-methods | Get payment methods of organization
[**OrgsOrgProvidersGet**](OrgsApi.md#OrgsOrgProvidersGet) | **Get** /orgs/{org}/providers | Get enabled cloud providers for an organization

# **OrgsGet**
> []Organization OrgsGet(ctx, )
Get organizations associated with calling user.

Return a list of all organizations the calling user (based on credential) is authorized on. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Organization**](Organization.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgGet**
> Organization OrgsOrgGet(ctx, org)
Get details of organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 

### Return type

[**Organization**](Organization.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgPaymentMethodsGet**
> []PaymentMethodInfo OrgsOrgPaymentMethodsGet(ctx, org, optional)
Get payment methods of organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***OrgsApiOrgsOrgPaymentMethodsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsApiOrgsOrgPaymentMethodsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultFlag** | **optional.Bool**| When true, will only return default payment methods. | [default to false]

### Return type

[**[]PaymentMethodInfo**](PaymentMethodInfo.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgProvidersGet**
> []AwsCloudProvider OrgsOrgProvidersGet(ctx, org)
Get enabled cloud providers for an organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 

### Return type

[**[]AwsCloudProvider**](AwsCloudProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

