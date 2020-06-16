# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgStorageClusterConstraintsGet**](StorageApi.md#OrgsOrgStorageClusterConstraintsGet) | **Get** /orgs/{org}/storage/cluster-constraints | Get constraints on cluster storage size for EBS-backed clusters.

# **OrgsOrgStorageClusterConstraintsGet**
> VsanConfigConstraints OrgsOrgStorageClusterConstraintsGet(ctx, org, provider, numHosts, optional)
Get constraints on cluster storage size for EBS-backed clusters.

Get constraints on cluster storage size for EBS-backed clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **provider** | **string**| Cloud storage provider ID (example AWS) | 
  **numHosts** | **int32**| Number of hosts in cluster | 
 **optional** | ***StorageApiOrgsOrgStorageClusterConstraintsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiOrgsOrgStorageClusterConstraintsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oneNodeReducedCapacity** | **optional.Bool**| Whether this sddc is reduced capacity 1NODE. | 

### Return type

[**VsanConfigConstraints**](VsanConfigConstraints.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

