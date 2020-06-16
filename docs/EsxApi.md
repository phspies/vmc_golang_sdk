# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsSddcEsxsPost**](EsxApi.md#OrgsOrgSddcsSddcEsxsPost) | **Post** /orgs/{org}/sddcs/{sddc}/esxs | Add/Remove one or more ESX hosts in the target cloud

# **OrgsOrgSddcsSddcEsxsPost**
> Task OrgsOrgSddcsSddcEsxsPost(ctx, body, org, sddc, optional)
Add/Remove one or more ESX hosts in the target cloud

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EsxConfig**](EsxConfig.md)| esxConfig | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
 **optional** | ***EsxApiOrgsOrgSddcsSddcEsxsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EsxApiOrgsOrgSddcsSddcEsxsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **action** | **optional.**| If &#x3D; &#x27;add&#x27;, will add the esx. If &#x3D; &#x27;remove&#x27;, will delete the esx/esxs bound to a single cluster (Cluster Id is mandatory for non cluster 1 esx remove). If &#x3D; &#x27;force-remove&#x27;, will delete the esx even if it can lead to data loss (This is an privileged operation). If &#x3D; &#x27;addToAll&#x27;, will add esxs to all clusters in the SDDC (This is an privileged operation). If &#x3D; &#x27;removeFromAll&#x27;, will delete the esxs from all clusters in the SDDC (This is an privileged operation). If &#x3D; &#x27;attach-diskgroup&#x27;, will attach the provided diskgroups to a given host (privileged). If &#x3D; &#x27;detach-diskgroup&#x27;, will detach the diskgroups of a given host (privileged). Default behaviour is &#x27;add&#x27;  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

