# ConnectivityValidationInput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | input value type, i.e. HOSTNAME_OR_IP, HOST_IP, HOSTNAME. Accept FQDN or IP address as input value when id &#x3D; HOSTNAME_OR_IP, accept FQDN ONLY when id &#x3D; HOSTNAME, accept IP address ONLY when id &#x3D; HOST_IP. | [optional] [default to null]
**Value** | **string** | the FQDN or IP address to run the test against, use \\#primary-dns or \\#secondary-dns as the on-prem primary/secondary DNS server IP. | [optional] [default to null]
**Label** | **string** | (Optional, for UI display only) input value label. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

