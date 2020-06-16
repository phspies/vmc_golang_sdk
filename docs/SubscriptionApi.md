# {{classname}}

All URIs are relative to *https://vmc.vmware.com/vmc/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsOrgSubscriptionsGet**](SubscriptionApi.md#OrgsOrgSubscriptionsGet) | **Get** /orgs/{org}/subscriptions | Get all subscriptions
[**OrgsOrgSubscriptionsOfferInstancesGet**](SubscriptionApi.md#OrgsOrgSubscriptionsOfferInstancesGet) | **Get** /orgs/{org}/subscriptions/offer-instances | List all offers available for the specific product type in the specific region
[**OrgsOrgSubscriptionsPost**](SubscriptionApi.md#OrgsOrgSubscriptionsPost) | **Post** /orgs/{org}/subscriptions | Create a subscription
[**OrgsOrgSubscriptionsProductsGet**](SubscriptionApi.md#OrgsOrgSubscriptionsProductsGet) | **Get** /orgs/{org}/subscriptions/products | List of all the products that are available for purchase.
[**OrgsOrgSubscriptionsSubscriptionGet**](SubscriptionApi.md#OrgsOrgSubscriptionsSubscriptionGet) | **Get** /orgs/{org}/subscriptions/{subscription} | Get a subscription

# **OrgsOrgSubscriptionsGet**
> []SubscriptionDetails OrgsOrgSubscriptionsGet(ctx, org, optional)
Get all subscriptions

Returns all subscriptions for a given org id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
 **optional** | ***SubscriptionApiOrgsOrgSubscriptionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionApiOrgsOrgSubscriptionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offerType** | **optional.String**| Offer Type  * &#x60;ON_DEMAND&#x60; - on-demand subscription  * &#x60;TERM&#x60; - term subscription  * All subscriptions if not specified  | 

### Return type

[**[]SubscriptionDetails**](SubscriptionDetails.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSubscriptionsOfferInstancesGet**
> OfferInstancesHolder OrgsOrgSubscriptionsOfferInstancesGet(ctx, org, region, productType, optional)
List all offers available for the specific product type in the specific region

List all offers available for the specific product type in the specific region 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **region** | **string**| Region for the offer | 
  **productType** | **string**| Type of the product in offers. *This has been deprecated*. Please use product &amp; type parameters | 
 **optional** | ***SubscriptionApiOrgsOrgSubscriptionsOfferInstancesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionApiOrgsOrgSubscriptionsOfferInstancesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **product** | **optional.String**| The product that you are trying to purchase, eg. host. This needs to be accompanied by the type parameter | 
 **type_** | **optional.String**| The type/flavor of the product you are trying it purchase,eg. an &#x60;r5.metal&#x60; host. This needs to be accompanied by the product parameter.  | 

### Return type

[**OfferInstancesHolder**](OfferInstancesHolder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSubscriptionsPost**
> Task OrgsOrgSubscriptionsPost(ctx, body, org)
Create a subscription

Initiates the creation of a subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionRequest**](SubscriptionRequest.md)| subscriptionRequest | 
  **org** | **string**| Organization identifier. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSubscriptionsProductsGet**
> []SubscriptionProducts OrgsOrgSubscriptionsProductsGet(ctx, org)
List of all the products that are available for purchase.

List of all the products that are available for purchase. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 

### Return type

[**[]SubscriptionProducts**](SubscriptionProducts.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrgsOrgSubscriptionsSubscriptionGet**
> SubscriptionDetails OrgsOrgSubscriptionsSubscriptionGet(ctx, org, subscription)
Get a subscription

Get subscription details for a given subscription id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier. | 
  **subscription** | **string**| SubscriptionId for an sddc. | 

### Return type

[**SubscriptionDetails**](SubscriptionDetails.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

