# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgTasksGet**](TaskApi.md#OrgsOrgTasksGet) | **Get** /orgs/{org}/tasks | List all tasks for organization
[**OrgsOrgTasksTaskGet**](TaskApi.md#OrgsOrgTasksTaskGet) | **Get** /orgs/{org}/tasks/{task} | Get task details
[**OrgsOrgTasksTaskPost**](TaskApi.md#OrgsOrgTasksTaskPost) | **Post** /orgs/{org}/tasks/{task} | Modify an existing task

# **OrgsOrgTasksGet**
> []Task OrgsOrgTasksGet(ctx, org, optional)
List all tasks for organization

List all tasks with optional filtering. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***TaskApiOrgsOrgTasksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskApiOrgsOrgTasksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter expression  Binary Operators: &#x27;eq&#x27;, &#x27;ne&#x27;, &#x27;lt&#x27;, &#x27;gt&#x27;, &#x27;le&#x27;, &#x27;ge&#x27;, &#x27;mul&#x27;, &#x27;div&#x27;, &#x27;mod&#x27;, &#x27;sub&#x27;, &#x27;add&#x27; Unary Operators: &#x27;not&#x27;, &#x27;-&#x27; (minus) String Operators: &#x27;startswith&#x27;, &#x27;endswith&#x27;, &#x27;length&#x27;, &#x27;contains&#x27;, &#x27;tolower&#x27;, &#x27;toupper&#x27;,  Nested attributes are composed using &#x27;.&#x27;  Dates must be formatted as yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss[.SSS]Z  Strings should enclosed in single quotes, escape single quote with two single quotes  The special literal &#x27;created&#x27; will be mapped to the time the resource was first created.  Examples:    - $filter&#x3D;(updated gt 2016-08-09T13:00:00Z) and (org_id eq 278710ff4e-6b6d-4d4e-aefb-ca637f38609e)   - $filter&#x3D;(created eq 2016-08-09)   - $filter&#x3D;(created gt 2016-08-09) and (sddc.status eq &#x27;READY&#x27;)  | 

### Return type

[**[]Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgTasksTaskGet**
> Task OrgsOrgTasksTaskGet(ctx, org, task)
Get task details

Retrieve details of a task. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **task** | **string**| Task identifier | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgTasksTaskPost**
> Task OrgsOrgTasksTaskPost(ctx, org, task, optional)
Modify an existing task

Request that a running task be canceled. This is advisory only, some tasks may not be cancelable, and some tasks might take an arbitrary amount of time to respond to a cancelation request. The task must be monitored to determine subsequent status. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **task** | **string**| Task identifier | 
 **optional** | ***TaskApiOrgsOrgTasksTaskPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskApiOrgsOrgTasksTaskPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **action** | **optional.String**| If &#x3D; &#x27;cancel&#x27;, task will be canceled | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

