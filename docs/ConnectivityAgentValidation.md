# ConnectivityAgentValidation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | source appliance of connectivity test, i.e. VCENTER, SRM, VR. | [optional] [default to null]
**Type_** | **string** | type of connectivity test, i.e. PING, TRACEROUTE, DNS, CONNECTIVITY, CURL. For CONNECTIVITY and CURL tests only, please specify the ports to be tested against. | [optional] [default to null]
**Ports** | **[]string** | TCP ports ONLY for CONNECTIVITY and CURL tests. | [optional] [default to null]
**Path** | **string** | URL path ONLY for CURL tests. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

