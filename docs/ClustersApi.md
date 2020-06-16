# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsSddcClustersClusterReconfigurePost**](ClustersApi.md#OrgsOrgSddcsSddcClustersClusterReconfigurePost) | **Post** /orgs/{org}/sddcs/{sddc}/clusters/{cluster}/reconfigure | Initiate a reconfiguration for the cluster.

# **OrgsOrgSddcsSddcClustersClusterReconfigurePost**
> Task OrgsOrgSddcsSddcClustersClusterReconfigurePost(ctx, body, org, sddc, cluster)
Initiate a reconfiguration for the cluster.

This reconfiguration will handle changing both the number of hosts and the cluster storage capacity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterReconfigureParams**](ClusterReconfigureParams.md)| clusterReconfigureParams | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **cluster** | **string**| cluster identifier | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

