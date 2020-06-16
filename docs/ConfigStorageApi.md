# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsSddcClustersClusterConfigConstraintsGet**](ConfigStorageApi.md#OrgsOrgSddcsSddcClustersClusterConfigConstraintsGet) | **Get** /orgs/{org}/sddcs/{sddc}/clusters/{cluster}/config/constraints | Get storage size constraints on existing cluster.

# **OrgsOrgSddcsSddcClustersClusterConfigConstraintsGet**
> VsanClusterReconfigConstraints OrgsOrgSddcsSddcClustersClusterConfigConstraintsGet(ctx, org, sddc, cluster, optional)
Get storage size constraints on existing cluster.

This API provides the storage choices available when reconfiguring storage in a cluster. The constraints returned will give vSAN reconfiguration biases and available vSAN capacities per reconfiguration bias. The constraints also indicate the default vSAN capacity per reconfiguration biases as well as the default reconfiguration bias. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **cluster** | **string**| cluster identifier | 
 **optional** | ***ConfigStorageApiOrgsOrgSddcsSddcClustersClusterConfigConstraintsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigStorageApiOrgsOrgSddcsSddcClustersClusterConfigConstraintsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expectedNumHosts** | **optional.Int32**| The expected number of hosts in the cluster. If not specified, returned is based on current number of hosts in the cluster.  | 

### Return type

[**VsanClusterReconfigConstraints**](VsanClusterReconfigConstraints.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

