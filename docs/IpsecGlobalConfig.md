# IpsecGlobalConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Psk** | **string** | IPsec Global Pre Shared Key. Maximum characters is 128. Required when peerIp is configured as &#x27;any&#x27; in NSX Edge IPsec Site configuration. | [optional] [default to null]
**CaCertificates** | [***CaCertificates**](caCertificates.md) |  | [optional] [default to null]
**ServiceCertificate** | **string** | Certificate name or identifier. Required when x.509 is selected as the authentication mode. | [optional] [default to null]
**CrlCertificates** | [***CrlCertificates**](crlCertificates.md) |  | [optional] [default to null]
**Extension** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

