# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgReservationsGet**](ReservationsApi.md#OrgsOrgReservationsGet) | **Get** /orgs/{org}/reservations | Get all reservations for this org
[**OrgsOrgReservationsReservationMwGet**](ReservationsApi.md#OrgsOrgReservationsReservationMwGet) | **Get** /orgs/{org}/reservations/{reservation}/mw | get the maintenance window for this SDDC
[**OrgsOrgReservationsReservationMwPut**](ReservationsApi.md#OrgsOrgReservationsReservationMwPut) | **Put** /orgs/{org}/reservations/{reservation}/mw | update the maintenance window for this SDDC

# **OrgsOrgReservationsGet**
> []MaintenanceWindowEntry OrgsOrgReservationsGet(ctx, org)
Get all reservations for this org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 

### Return type

[**[]MaintenanceWindowEntry**](MaintenanceWindowEntry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgReservationsReservationMwGet**
> MaintenanceWindowGet OrgsOrgReservationsReservationMwGet(ctx, org, reservation)
get the maintenance window for this SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **reservation** | **string**| Reservation Identifier | 

### Return type

[**MaintenanceWindowGet**](MaintenanceWindowGet.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgReservationsReservationMwPut**
> MaintenanceWindow OrgsOrgReservationsReservationMwPut(ctx, body, org, reservation)
update the maintenance window for this SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MaintenanceWindow**](MaintenanceWindow.md)| Maintenance Window | 
  **org** | **string**| Organization identifier. | 
  **reservation** | **string**| Reservation Identifier | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

