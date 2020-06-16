# IpsecSite

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Psk** | **string** | Pre Shared Key for the IPsec Site. Required if Site peerIp is not &#x27;any&#x27;. Global PSK is used when Authentication mode is PSK and Site peerIp is &#x27;any&#x27;. | [optional] [default to null]
**LocalId** | **string** | Local ID of the IPsec Site. Defaults to the local IP. | [optional] [default to null]
**EnablePfs** | **bool** | Enable/disable Perfect Forward Secrecy. Default is true. | [optional] [default to null]
**AuthenticationMode** | **string** | Authentication mode for the IPsec Site. Valid values are psk and x.509, with psk as default. | [optional] [default to null]
**PeerSubnets** | [***Subnets**](Subnets.md) |  | [optional] [default to null]
**DhGroup** | **string** | Diffie-Hellman algorithm group. Defaults to DH14 for FIPS enabled NSX Edge. DH2 and DH5 are not supported when FIPS is enabled on NSX Edge. Valid values are DH2, DH5, DH14, DH15, DH16. | [optional] [default to null]
**SiteId** | **string** | ID of the IPsec Site configuration provided by NSX Manager. | [optional] [default to null]
**Description** | **string** | Description of the IPsec Site. | [optional] [default to null]
**PeerIp** | **string** | IP (IPv4) address or FQDN of the Peer. Can also be specified as &#x27;any&#x27;. Required. | [optional] [default to null]
**Name** | **string** | Name of the IPsec Site. | [optional] [default to null]
**Certificate** | **string** |  | [optional] [default to null]
**LocalIp** | **string** | Local IP of the IPsec Site. Should be one of the IP addresses configured on the uplink interfaces of the NSX Edge. Required. | [optional] [default to null]
**EncryptionAlgorithm** | **string** | IPsec encryption algorithm with default as aes256. Valid values are &#x27;aes&#x27;, &#x27;aes256&#x27;, &#x27;3des&#x27;, &#x27;aes-gcm&#x27;. | [optional] [default to null]
**Enabled** | **bool** | Enable/disable IPsec Site. | [optional] [default to null]
**Mtu** | **int32** | MTU for the IPsec site. Defaults to the mtu of the NSX Edge vnic specified by the localIp. Optional. | [optional] [default to null]
**Extension** | **string** |  | [optional] [default to null]
**PeerId** | **string** | Peer ID. Should be unique for all IPsec Site&#x27;s configured for an NSX Edge. | [optional] [default to null]
**LocalSubnets** | [***Subnets**](Subnets.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

