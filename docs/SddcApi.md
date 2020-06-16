# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSddcsGet**](SddcApi.md#OrgsOrgSddcsGet) | **Get** /orgs/{org}/sddcs | List all the SDDCs of an organization
[**OrgsOrgSddcsPost**](SddcApi.md#OrgsOrgSddcsPost) | **Post** /orgs/{org}/sddcs | Provision SDDC
[**OrgsOrgSddcsProvisionSpecGet**](SddcApi.md#OrgsOrgSddcsProvisionSpecGet) | **Get** /orgs/{org}/sddcs/provision-spec | Get sddc provision spec for an org
[**OrgsOrgSddcsSddcClustersClusterDelete**](SddcApi.md#OrgsOrgSddcsSddcClustersClusterDelete) | **Delete** /orgs/{org}/sddcs/{sddc}/clusters/{cluster} | Delete a cluster.
[**OrgsOrgSddcsSddcClustersClusterReconfigurePost**](SddcApi.md#OrgsOrgSddcsSddcClustersClusterReconfigurePost) | **Post** /orgs/{org}/sddcs/{sddc}/clusters/{cluster}/reconfigure | Initiate a reconfiguration for the cluster.
[**OrgsOrgSddcsSddcClustersPost**](SddcApi.md#OrgsOrgSddcsSddcClustersPost) | **Post** /orgs/{org}/sddcs/{sddc}/clusters | Add a cluster in the target sddc.
[**OrgsOrgSddcsSddcConvertPost**](SddcApi.md#OrgsOrgSddcsSddcConvertPost) | **Post** /orgs/{org}/sddcs/{sddc}/convert | Converts a one host SDDC to a four node DEFAULT SDDC.
[**OrgsOrgSddcsSddcDelete**](SddcApi.md#OrgsOrgSddcsSddcDelete) | **Delete** /orgs/{org}/sddcs/{sddc} | Delete SDDC
[**OrgsOrgSddcsSddcDnsPrivatePut**](SddcApi.md#OrgsOrgSddcsSddcDnsPrivatePut) | **Put** /orgs/{org}/sddcs/{sddc}/dns/private | Update the DNS records of management VMs to use private IP addresses
[**OrgsOrgSddcsSddcDnsPublicPut**](SddcApi.md#OrgsOrgSddcsSddcDnsPublicPut) | **Put** /orgs/{org}/sddcs/{sddc}/dns/public | Update the DNS records of management VMs to use public IP addresses
[**OrgsOrgSddcsSddcEsxsPost**](SddcApi.md#OrgsOrgSddcsSddcEsxsPost) | **Post** /orgs/{org}/sddcs/{sddc}/esxs | Add/Remove one or more ESX hosts in the target cloud
[**OrgsOrgSddcsSddcGet**](SddcApi.md#OrgsOrgSddcsSddcGet) | **Get** /orgs/{org}/sddcs/{sddc} | Get SDDC
[**OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsGet**](SddcApi.md#OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsGet) | **Get** /orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials | List all the credentials assoicated with an addon type with in a SDDC
[**OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNameGet**](SddcApi.md#OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNameGet) | **Get** /orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name} | Get credential details by name
[**OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNamePut**](SddcApi.md#OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNamePut) | **Put** /orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name} | Update credential details
[**OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsPost**](SddcApi.md#OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsPost) | **Post** /orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials | Associate a new add on credentials with the SDDC such as HCX
[**OrgsOrgSddcsSddcManagementVmsManagementVmIdDnsIpTypePut**](SddcApi.md#OrgsOrgSddcsSddcManagementVmsManagementVmIdDnsIpTypePut) | **Put** /orgs/{org}/sddcs/{sddc}/management-vms/{managementVmId}/dns/{ipType} | Update the DNS records of management VMs to use public or private IP addresses
[**OrgsOrgSddcsSddcPatch**](SddcApi.md#OrgsOrgSddcsSddcPatch) | **Patch** /orgs/{org}/sddcs/{sddc} | Patch SDDC
[**OrgsOrgSddcsSddcPublicipsGet**](SddcApi.md#OrgsOrgSddcsSddcPublicipsGet) | **Get** /orgs/{org}/sddcs/{sddc}/publicips | List all public IPs for a SDDC
[**OrgsOrgSddcsSddcPublicipsIdDelete**](SddcApi.md#OrgsOrgSddcsSddcPublicipsIdDelete) | **Delete** /orgs/{org}/sddcs/{sddc}/publicips/{id} | Free one public IP for a SDDC
[**OrgsOrgSddcsSddcPublicipsIdGet**](SddcApi.md#OrgsOrgSddcsSddcPublicipsIdGet) | **Get** /orgs/{org}/sddcs/{sddc}/publicips/{id} | Get one public IP for a SDDC
[**OrgsOrgSddcsSddcPublicipsIdPatch**](SddcApi.md#OrgsOrgSddcsSddcPublicipsIdPatch) | **Patch** /orgs/{org}/sddcs/{sddc}/publicips/{id} | Attach or detach a public IP to workload VM for a SDDC
[**OrgsOrgSddcsSddcPublicipsPost**](SddcApi.md#OrgsOrgSddcsSddcPublicipsPost) | **Post** /orgs/{org}/sddcs/{sddc}/publicips | Allocate public IPs for a SDDC

# **OrgsOrgSddcsGet**
> []Sddc OrgsOrgSddcsGet(ctx, org, optional)
List all the SDDCs of an organization

List all the SDDCs of an organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***SddcApiOrgsOrgSddcsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcApiOrgsOrgSddcsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **optional.Bool**| When true, forces the result to also include deleted SDDCs. | 

### Return type

[**[]Sddc**](Sddc.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsPost**
> Task OrgsOrgSddcsPost(ctx, body, org, optional)
Provision SDDC

Provision an SDDC in target cloud

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AwsSddcConfig**](AwsSddcConfig.md)| sddcConfig | 
  **org** | **string**| Organization identifier. | 
 **optional** | ***SddcApiOrgsOrgSddcsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcApiOrgsOrgSddcsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.**| When true, only validates the given sddc configuration without provisioning. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsProvisionSpecGet**
> ProvisionSpec OrgsOrgSddcsProvisionSpecGet(ctx, org)
Get sddc provision spec for an org

Get sddc provision spec for an org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 

### Return type

[**ProvisionSpec**](ProvisionSpec.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcClustersClusterDelete**
> Task OrgsOrgSddcsSddcClustersClusterDelete(ctx, org, sddc, cluster)
Delete a cluster.

This is a force operation which will delete the cluster even if there can be a data loss. Before calling this operation, all the VMs should be powered off.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **cluster** | **string**| cluster identifier | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **OrgsOrgSddcsSddcClustersPost**
> Task OrgsOrgSddcsSddcClustersPost(ctx, body, org, sddc)
Add a cluster in the target sddc.

Creates a new cluster in customers sddcs with passed clusterConfig.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterConfig**](ClusterConfig.md)| clusterConfig | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcConvertPost**
> Task OrgsOrgSddcsSddcConvertPost(ctx, org, sddc)
Converts a one host SDDC to a four node DEFAULT SDDC.

This API converts a one host SDDC to a four node DEFAULT SDDC. It takes care of configuring and upgrading the vCenter configurations on the SDDC for high availability and data redundancy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcDelete**
> Task OrgsOrgSddcsSddcDelete(ctx, org, sddc, optional)
Delete SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
 **optional** | ***SddcApiOrgsOrgSddcsSddcDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcApiOrgsOrgSddcsSddcDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **retainConfiguration** | **optional.Bool**| If &#x3D; &#x27;true&#x27;, the SDDC&#x27;s configuration is retained as a template for later use. This flag is applicable only to SDDCs in ACTIVE state.  | 
 **templateName** | **optional.String**| Only applicable when retainConfiguration is also set to &#x27;true&#x27;. When set, this value will be used as the name of the SDDC configuration template generated.  | 
 **force** | **optional.Bool**| If &#x3D; true, will delete forcefully. Beware: do not use the force flag if there is a chance an active provisioning or deleting task is running against this SDDC. This option is restricted.  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcDnsPrivatePut**
> Task OrgsOrgSddcsSddcDnsPrivatePut(ctx, org, sddc)
Update the DNS records of management VMs to use private IP addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcDnsPublicPut**
> Task OrgsOrgSddcsSddcDnsPublicPut(ctx, org, sddc)
Update the DNS records of management VMs to use public IP addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 **optional** | ***SddcApiOrgsOrgSddcsSddcEsxsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcApiOrgsOrgSddcsSddcEsxsPostOpts struct
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

# **OrgsOrgSddcsSddcGet**
> Sddc OrgsOrgSddcsSddcGet(ctx, org, sddc)
Get SDDC

Get SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Sddc**](Sddc.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsGet**
> []NewCredentials OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsGet(ctx, org, sddcId, addonType)
List all the credentials assoicated with an addon type with in a SDDC

List all the credentials assoicated with an addon type with in a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | [**string**](.md)| Org id of the associated SDDC | 
  **sddcId** | [**string**](.md)| Id of the SDDC | 
  **addonType** | **string**| Add on type | 

### Return type

[**[]NewCredentials**](NewCredentials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNameGet**
> NewCredentials OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNameGet(ctx, org, sddcId, addonType, name)
Get credential details by name

Get credential details by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | [**string**](.md)| Org id of the associated SDDC | 
  **sddcId** | [**string**](.md)| Id of the SDDC | 
  **addonType** | **string**| Add on type | 
  **name** | **string**| name of the credentials | 

### Return type

[**NewCredentials**](NewCredentials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNamePut**
> NewCredentials OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsNamePut(ctx, body, org, sddcId, addonType, name)
Update credential details

Update credential details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCredentials**](UpdateCredentials.md)| Credentials update payload | 
  **org** | [**string**](.md)| Org id of the associated SDDC | 
  **sddcId** | [**string**](.md)| Id of the SDDC | 
  **addonType** | **string**| Add on type | 
  **name** | **string**| name of the credentials | 

### Return type

[**NewCredentials**](NewCredentials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsPost**
> NewCredentials OrgsOrgSddcsSddcIdAddonsAddonTypeCredentialsPost(ctx, body, org, sddcId, addonType)
Associate a new add on credentials with the SDDC such as HCX

Associated a new add on credentials with the SDDC such as HCX

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewCredentials**](NewCredentials.md)| Credentials creation payload | 
  **org** | [**string**](.md)| Org id of the associated SDDC | 
  **sddcId** | [**string**](.md)| Id of the SDDC | 
  **addonType** | **string**| Add on type | 

### Return type

[**NewCredentials**](NewCredentials.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcManagementVmsManagementVmIdDnsIpTypePut**
> Task OrgsOrgSddcsSddcManagementVmsManagementVmIdDnsIpTypePut(ctx, org, sddc, managementVmId, ipType)
Update the DNS records of management VMs to use public or private IP addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **managementVmId** | **string**| the management VM ID | 
  **ipType** | **string**| the ip type to associate with FQDN in DNS | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPatch**
> Sddc OrgsOrgSddcsSddcPatch(ctx, body, org, sddc)
Patch SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SddcPatchRequest**](SddcPatchRequest.md)| Patch request for the SDDC | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Sddc**](Sddc.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsGet**
> []SddcPublicIp OrgsOrgSddcsSddcPublicipsGet(ctx, org, sddc)
List all public IPs for a SDDC

list all public IPs for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**[]SddcPublicIp**](SddcPublicIp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsIdDelete**
> Task OrgsOrgSddcsSddcPublicipsIdDelete(ctx, org, sddc, id)
Free one public IP for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **id** | **string**| ip allocation id | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsIdGet**
> SddcPublicIp OrgsOrgSddcsSddcPublicipsIdGet(ctx, org, sddc, id)
Get one public IP for a SDDC

Get one public IP for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **id** | **string**| ip allocation id | 

### Return type

[**SddcPublicIp**](SddcPublicIp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsIdPatch**
> Task OrgsOrgSddcsSddcPublicipsIdPatch(ctx, body, org, sddc, id, action)
Attach or detach a public IP to workload VM for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SddcPublicIp**](SddcPublicIp.md)| SddcPublicIp object to update | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 
  **id** | **string**| ip allocation id | 
  **action** | **string**| Type of action as &#x27;attach&#x27;, &#x27;detach&#x27;, &#x27;reattach&#x27;, or &#x27;rename&#x27;. For &#x27;attch&#x27;, the public IP must not be attached and &#x27;associated_private_ip&#x27; in the payload needs to be set with a workload VM private IP. For &#x27;detach&#x27;, the public IP must be attached and &#x27;associated_private_ip&#x27; in the payload should not be set with any value. For &#x27;reattach&#x27;, the public IP must be attached and &#x27;associated_private_ip&#x27; in the payload needs to be set with a new workload VM private IP. For &#x27;rename&#x27;, the &#x27;name&#x27; in the payload needs to have a new name string.  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSddcsSddcPublicipsPost**
> Task OrgsOrgSddcsSddcPublicipsPost(ctx, body, org, sddc)
Allocate public IPs for a SDDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SddcAllocatePublicIpSpec**](SddcAllocatePublicIpSpec.md)| allocation spec | 
  **org** | **string**| Organization identifier. | 
  **sddc** | **string**| Sddc Identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

