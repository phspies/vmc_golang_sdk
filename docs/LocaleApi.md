# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalePost**](LocaleApi.md#LocalePost) | **Post** /locale | Set locale for the session

# **LocalePost**
> VmcLocale LocalePost(ctx, body)
Set locale for the session

Sets the locale for the session which is used for translating responses.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VmcLocale**](VmcLocale.md)| The locale to be set. | 

### Return type

[**VmcLocale**](VmcLocale.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

