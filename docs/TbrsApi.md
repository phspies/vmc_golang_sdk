# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgTbrsReservationPost**](TbrsApi.md#OrgsOrgTbrsReservationPost) | **Post** /orgs/{org}/tbrs/reservation | 
[**OrgsOrgTbrsSupportWindowGet**](TbrsApi.md#OrgsOrgTbrsSupportWindowGet) | **Get** /orgs/{org}/tbrs/support-window | 
[**OrgsOrgTbrsSupportWindowIdPut**](TbrsApi.md#OrgsOrgTbrsSupportWindowIdPut) | **Put** /orgs/{org}/tbrs/support-window/{id} | 

# **OrgsOrgTbrsReservationPost**
> map[string][]ReservationWindow OrgsOrgTbrsReservationPost(ctx, org, optional)


Retreive all reservations for all SDDCs in this org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***TbrsApiOrgsOrgTbrsReservationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TbrsApiOrgsOrgTbrsReservationPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SddcStateRequest**](SddcStateRequest.md)| SDDCs and/or states to query | 

### Return type

[**map[string][]ReservationWindow**](array.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgTbrsSupportWindowGet**
> []SupportWindow OrgsOrgTbrsSupportWindowGet(ctx, org, optional)


Get all available support windows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***TbrsApiOrgsOrgTbrsSupportWindowGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TbrsApiOrgsOrgTbrsSupportWindowGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minimumSeatsAvailable** | **optional.Int32**| minimum seats available (used as a filter) | 
 **createdBy** | **optional.String**| user name which was used to create the support window (used as a filter) | 

### Return type

[**[]SupportWindow**](SupportWindow.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgTbrsSupportWindowIdPut**
> SupportWindowId OrgsOrgTbrsSupportWindowIdPut(ctx, body, org, id)


Move Sddc to new support window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SddcId**](SddcId.md)| SDDC to move | 
  **org** | **string**| Organization identifier. | 
  **id** | **string**| Target Support Window ID | 

### Return type

[**SupportWindowId**](SupportWindowId.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

